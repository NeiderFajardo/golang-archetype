package domain

type IProductRepository interface {
	GetByID(id int) (*Product, error)
	Create(product *Product) (int, error)
}
