package models

// Product response model
type ProductResponse struct {
	ID          int     `json:"id"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Name        string  `json:"name"`
	Stock       int     `json:"stock"`
}

// NewProductResponse creates a new ProductResponse
func NewProductResponse(id int, name, description string, price float64, stock int) *ProductResponse {
	return &ProductResponse{
		ID:          id,
		Name:        name,
		Description: description,
		Price:       price,
		Stock:       stock,
	}
}
