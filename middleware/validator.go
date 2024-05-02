package middleware

import (
	"context"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/sessions"
	"github.com/sebasvil20/templ-sys-login-exp/views/pages"
)

func Validator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, "validator", validator.New())
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func Auth(store *sessions.CookieStore) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			session, _ := store.Get(r, "session")
			if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
				pages.Forbidden().Render(r.Context(), w)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
