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
}

func (pr *ProductRequest) Valid(ctx context.Context) *apierrors.ApiError {
	if pr.Id <= 0 {
		return apierrors.BadRequest("The id is required", "parameter_required", "id")
	}
	return nil
}
