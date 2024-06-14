package domain

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
}

func NewProduct(name, description string, price float64) *Product {
	return &Product{
		Name:        name,
		Description: description,
		Price:       price,
	}
}
