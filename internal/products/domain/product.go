package domain

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
}

func NewProduct(id int, name, description string, price float64) *Product {
	return &Product{
		ID:          id,
		Name:        name,
		Description: description,
		Price:       price,
	}
}
