package middlewares

import (
	"fmt"
	"net/http"
)

type AuthMiddleware struct{}

func (am *AuthMiddleware) HandlerFunc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executing AuthMiddleware")
		// Add your authentication logic here
		next.ServeHTTP(w, r)
	})
}

func (am *AuthMiddleware) Order() int {
	return 0
}
