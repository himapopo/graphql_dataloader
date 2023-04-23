package middleware

import (
	"context"
	"graphql_dataloader/usecase/user"
	"net/http"
)

func LoaderMiddleware(userLoader user.UserReader, next http.Handler) http.Handler {
	// return a middleware that injects the loader to the request context
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextCtx := context.WithValue(r.Context(), user.LoadersKey, user.NewUserLoader(userLoader))
		r = r.WithContext(nextCtx)
		next.ServeHTTP(w, r)
	})
}
