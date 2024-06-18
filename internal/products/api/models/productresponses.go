package models

// Product response model
type ProductResponse struct {
	ID          int     `json:"id"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Name        string  `json:"name"`
}

// NewProductResponse creates a new ProductResponse
func NewProductResponse(id int, name, description string, price float64) *ProductResponse {
	return &ProductResponse{
		ID:          id,
		Name:        name,
		Description: description,
		Price:       price,
	}
}
