package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"
	"github.com/sebasvil20/templ-sys-login-exp/users"
	"github.com/sebasvil20/templ-sys-login-exp/views/pages"
)

func main() {
	decoder := schema.NewDecoder()
	store := sessions.NewCookieStore([]byte("secret"))
	instUserHandler := users.InitUserHandler()
	instUserController := users.InitUserController(instUserHandler, decoder, store)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/users", func(r chi.Router) {
		r.Use(AuthMiddleware)
		r.Get("/", instUserController.ListUserPage)
	})

	r.Get("/accounts", instUserController.AccountsPage)

	r.Route("/api", func(r chi.Router) {
		r.Use(ValidatorMiddleware)
		r.Post("/login", instUserController.LoginAPI)
		r.Post("/signin", instUserController.SigninAPI)
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
