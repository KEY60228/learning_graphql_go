package middleware

import (
	"context"
	"net/http"

	"gql/domain/model"
)

type contextKey string

var UserCtxKey contextKey = "userCtxKey"

func Auth(repository model.RepositoryInterface) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			currentUser := repository.UserByToken(r.Header.Get("Authorization"))
			ctx := context.WithValue(r.Context(), UserCtxKey, currentUser)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
