package api

import (
	"net/http"
	"strconv"

	"github.com/NeiderFajardo/internal/products/api/models"
	"github.com/NeiderFajardo/internal/products/application"
	"github.com/NeiderFajardo/pkg/utils"
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
		response := models.NewProductResponse(result.ID, result.Name, result.Description, result.Price)
		utils.Encode(w, http.StatusOK, response)
	}
}

func (ph *ProductHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Add your logic here
		productRequest, err := utils.Decode[models.ProductRequest](r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		productID, err := ph.productService.Create(&productRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response := models.NewProductResponse(productID, productRequest.Name, productRequest.Description, productRequest.Price)
		utils.Encode(w, http.StatusCreated, response)
	}
}
