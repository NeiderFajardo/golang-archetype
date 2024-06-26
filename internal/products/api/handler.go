package api

import (
	"net/http"
	"strconv"

	"github.com/NeiderFajardo/internal/products/api/models"
	"github.com/NeiderFajardo/internal/products/application"
	"github.com/NeiderFajardo/internal/server/response"
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
		result, err := ph.productService.GetByID(r.Context(), intId)
		if err != nil {
			response.ResponseError(w, err)
			return
		}
		productResponse := models.NewProductResponse(result.ID, result.Name, result.Description, result.Price)
		errDecode := utils.Encode(w, http.StatusOK, productResponse)
		if errDecode != nil {
			response.ResponseError(w, errDecode)
		}
	}
}

func (ph *ProductHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Add your logic here
		productRequest, err := utils.Decode[models.ProductRequest](r)
		if err != nil {
			response.ResponseError(w, err)
			return
		}
		productID, err := ph.productService.Create(r.Context(), &productRequest)
		if err != nil {
			response.ResponseError(w, err)
			return
		}
		productResponse := models.NewProductResponse(productID, productRequest.Name, productRequest.Description, productRequest.Price)
		errDecode := utils.Encode(w, http.StatusCreated, productResponse)
		if errDecode != nil {
			response.ResponseError(w, errDecode)
		}
	}
}
