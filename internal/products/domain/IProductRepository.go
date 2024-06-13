package domain

type IProductRepository interface {
	GetByID(id int) *Product
}
