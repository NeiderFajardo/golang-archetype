package middlewares

import (
	"context"
	"fmt"
	"net/http"
)

func IsAuthed(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executing AuthMiddleware")
		// Add your authentication logic here
		ctx := context.WithValue(r.Context(), "user", "User")
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
