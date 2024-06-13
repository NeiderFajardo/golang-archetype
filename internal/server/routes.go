package server

import (
	"net/http"

	"github.com/NeiderFajardo/internal/products/api"
)

func RegisterRoutes(
	product *api.ProductHandler,
) map[string]http.Handler {
	// Register the routes
	return map[string]http.Handler{
		"/products/{id}": product.GetByID(),
	}
}
