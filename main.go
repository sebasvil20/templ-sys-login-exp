package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"
	"github.com/sebasvil20/templ-sys-login-exp/users"
	"github.com/sebasvil20/templ-sys-login-exp/utils"
	"github.com/sebasvil20/templ-sys-login-exp/views/pages"
)

var store = sessions.NewCookieStore([]byte("test"))
var decoder = schema.NewDecoder()

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/users", func(r chi.Router) {
		r.Use(AuthMiddleware)
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html")
			pages.ListUser(users.GetUsers()).Render(r.Context(), w)
		})
	})

	r.Get("/accounts", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		pages.Accounts().Render(r.Context(), w)
	})

	r.Route("/api", func(r chi.Router) {
		r.Use(ValidatorMiddleware)
		r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
			validate, _ := r.Context().Value("validator").(*validator.Validate)
			reqUser := users.UserCredentials{}
			err := r.ParseMultipartForm(4096)
			if err != nil {
				utils.HandleReturnWithStatusCode(w, 400, map[string]string{"error": "Bad request body"})
				return
			}

			err = decoder.Decode(&reqUser, r.PostForm)
			if err != nil {
				utils.HandleReturnWithStatusCode(w, 400, map[string]string{"error": "Bad request body"})
				return
			}

			err = validate.Struct(&reqUser)
			if err != nil {
				utils.HandleReturnWithStatusCode(w, 400, map[string]string{"error": fmt.Sprintf("Validation error: %s", err.Error())})
				return
			}

			if !users.Authenticate(reqUser) {
				utils.HandleReturnWithStatusCode(w, 401, map[string]string{"error": "Invalid credentials"})
				return
			}
			session, _ := store.Get(r, "session")
			session.Values["authenticated"] = true
			err = session.Save(r, w)
			if err != nil {
				utils.HandleReturnWithStatusCode(w, 500, map[string]string{"error": "Error saving session"})
				return
			}
			utils.HandleReturnWithStatusCode(w, 200, map[string]any{"data": nil})
		})

		r.Post("/signin", func(w http.ResponseWriter, r *http.Request) {
			validate, _ := r.Context().Value("validator").(*validator.Validate)
			reqUser := users.User{}

			err := r.ParseMultipartForm(4096)
			if err != nil {
				utils.HandleReturnWithStatusCode(w, 400, map[string]string{"error": "Bad request body"})
				return
			}

			err = decoder.Decode(&reqUser, r.PostForm)
			if err != nil {
				utils.HandleReturnWithStatusCode(w, 400, map[string]string{"error": "Bad request body"})
				return
			}

			err = validate.Struct(&reqUser)
			if err != nil {
				utils.HandleReturnWithStatusCode(w, 400, map[string]string{"error": fmt.Sprintf("Validation error: %s", err.Error())})
				return
			}

			users.AddUser(reqUser)

			session, _ := store.Get(r, "session")
			session.Values["authenticated"] = true
			err = session.Save(r, w)
			if err != nil {
				utils.HandleReturnWithStatusCode(w, 500, map[string]string{"error": "Error saving session"})
				return
			}
			utils.HandleReturnWithStatusCode(w, 201, map[string]any{"data": nil})
		})
	})

	http.ListenAndServe(":3000", r)
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session")
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			pages.Forbidden().Render(r.Context(), w)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func ValidatorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, "validator", validator.New())
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
