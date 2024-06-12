package middlewares

import "net/http"

type IMiddleware interface {
	HandlerFunc(next http.Handler) http.Handler
	Order() int
}
