package api

import (
	"net/http"
	"strconv"

	"github.com/NeiderFajardo/internal/products/application"
)

type ProductHandler struct {
	productService application.IProductService
}

func NewProductHandler(productService application.IProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

func (ph *ProductHandler) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Add your logic here
		id := r.PathValue("id")
		intId, _ := strconv.Atoi(id)
		result, err := ph.productService.GetByID(intId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(result.Name))
	}
}
