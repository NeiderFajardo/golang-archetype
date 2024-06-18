package models

import "context"

// Product model
type ProductRequest struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func (pr *ProductRequest) Valid(ctx context.Context) map[string]string {
	var problems = make(map[string]string)
	if pr.Id == 0 {
		problems["id"] = "id is required"
	}
	return problems
}
