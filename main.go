package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"
	"github.com/sebasvil20/templ-sys-login-exp/middleware"
	"github.com/sebasvil20/templ-sys-login-exp/users"
)

func main() {
	decoder := schema.NewDecoder()
	store := sessions.NewCookieStore([]byte("secret"))
	instUserHandler := users.InitUserHandler()
	instUserController := users.InitUserController(instUserHandler, decoder, store)

	r := chi.NewRouter()
	r.Use(chiMiddleware.Logger)

	r.Route("/users", func(r chi.Router) {
		r.Use(middleware.Auth(store))
		r.Get("/", instUserController.ListUserPage)
	})

	r.Get("/accounts", instUserController.AccountsPage)

	r.Route("/api", func(r chi.Router) {
		r.Use(middleware.Validator)
		r.Post("/login", instUserController.LoginAPI)
		r.Post("/signin", instUserController.SigninAPI)
	})

	http.ListenAndServe(":3000", r)
}
