package server

import (
	"net/http"

	"github.com/NeiderFajardo/internal/products/api"
)

func RegisterRoutes(
	product *api.ProductHandler,
	productStock *api.ProductStockHandler,
) map[string]http.Handler {
	// Register the routes
	return map[string]http.Handler{
		// Product routes
		"/products/{id}": product.GetByID(),
		"POST /products": product.Create(),

		// Stock routes
		"POST /products/subtract/{id}": productStock.SubstractStock(),

		// Health check
		"/health": http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}),
	}
}
