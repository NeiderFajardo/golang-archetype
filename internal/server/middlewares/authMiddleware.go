package middlewares

import (
	"fmt"
	"net/http"
)

func IsAuthed(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executing AuthMiddleware")
		// Add your authentication logic here
		next.ServeHTTP(w, r)
	})
}
