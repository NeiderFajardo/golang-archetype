package models

import (
	"context"

	"github.com/NeiderFajardo/pkg/apierrors"
)

// Product model
type ProductRequest struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
}

func (pr *ProductRequest) Valid(ctx context.Context) map[string]apierrors.ApiError {
	problems := make(map[string]apierrors.ApiError)
	if pr.Id <= 0 {
		problems["id"] = *apierrors.BadRequest("id must be greater than 0", "invalid_id", "id")
	}
	return problems
}
