package server

import (
	"net/http"

	"github.com/NeiderFajardo/internal/products/api"
	"github.com/NeiderFajardo/internal/server/middlewares"
)

func RegisterRoutes() map[string]http.Handler {
	// Register the routes
	return map[string]http.Handler{
		"/products": api.NewProductHandler(),
	}
}

func RegisterMiddlewares() []middlewares.IMiddleware {
	return []middlewares.IMiddleware{
		&middlewares.AuthMiddleware{},
	}
}
