package api

import (
	"context"
	"net/http"
	"strconv"

	"github.com/NeiderFajardo/internal/products/application"
	"github.com/NeiderFajardo/pkg/apierrors"
	"github.com/NeiderFajardo/pkg/response"
	"github.com/NeiderFajardo/pkg/utils"
)

type ProductStockHandler struct {
	service application.IProductStockService
}

// Request object
type SubstractStockRequest struct {
	Quantity int `json:"quantity"`
}

func (sr *SubstractStockRequest) Valid(ctx context.Context) map[string]apierrors.ApiError {
	problems := make(map[string]apierrors.ApiError)
	if sr.Quantity <= 0 {
		problems["quantity"] = *apierrors.BadRequest("quantity must be greater than 0", "invalid_quantity", "quantity")
	}
	return problems
}

func NewProductStockHandler(service application.IProductStockService) *ProductStockHandler {
	return &ProductStockHandler{
		service: service,
	}
}

func (ph *ProductStockHandler) SubstractStock() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		intId, _ := strconv.Atoi(id)
		request, err := utils.DecodeValid[*SubstractStockRequest](r)
		if err != nil {
			response.ResponseError(w, err)
			return
		}
		errResponse := ph.service.SubtractStock(r.Context(), intId, request.Quantity)
		if errResponse != nil {
			response.ResponseError(w, errResponse)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
