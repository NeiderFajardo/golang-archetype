package domain

type Product struct {
	ID   int
	Name string
}

func NewProduct(id int, name string) *Product {
	return &Product{
		ID:   id,
		Name: name,
	}
}
