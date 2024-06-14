package models

// Product response model
type ProductResponse struct {
	ID int `json:"id"`
}

// NewProductResponse creates a new ProductResponse
func NewProductResponse(id int) *ProductResponse {
	return &ProductResponse{
		ID: id,
	}
}
