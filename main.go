package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/sessions"
	"github.com/sebasvil20/templ-sys-login-exp/users"
	"github.com/sebasvil20/templ-sys-login-exp/utils"
	"github.com/sebasvil20/templ-sys-login-exp/views/layouts"
	"github.com/sebasvil20/templ-sys-login-exp/views/pages"
)

var store = sessions.NewCookieStore([]byte("test"))

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/users", func(r chi.Router) {
		r.Use(AuthMiddleware)
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html")
			layouts.MainLayout(pages.ListUser(users.GetUsers())).Render(r.Context(), w)
		})
	})

	r.Get("/accounts", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		layouts.MainLayout(pages.Accounts()).Render(r.Context(), w)
	})

	r.Route("/api", func(r chi.Router) {
		r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
			reqUser := users.User{}
			body, _ := io.ReadAll(r.Body)
			err := json.Unmarshal(body, &reqUser)
			if err != nil {
				utils.HandleReturnWithStatusCode(w, 400, map[string]string{"error": "Bad request body"})
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
			reqUser := users.User{}
			body, _ := io.ReadAll(r.Body)
			fmt.Println(string(body))
			err := json.Unmarshal(body, &reqUser)
			if err != nil {
				utils.HandleReturnWithStatusCode(w, 400, map[string]string{"error": "Bad request body"})
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
			layouts.MainLayout(pages.Forbidden()).Render(r.Context(), w)
			return
		}
		next.ServeHTTP(w, r)
	})
}
