package events

import (
	"github.com/NeiderFajardo/internal/products/domain"
	"github.com/NeiderFajardo/pkg/events"
	"github.com/google/uuid"
)

type ProductEvent interface {
	events.DomainEvent
	GetProduct() *domain.Product
}

type UpdatedStockEvent struct {
	id      uuid.UUID
	product *domain.Product
}

func NewUpdatedStockEvent(id uuid.UUID, product *domain.Product) UpdatedStockEvent {
	return UpdatedStockEvent{
		id:      id,
		product: product,
	}
}

func (e UpdatedStockEvent) Name() string {
	return "product.stock.updated"
}

func (e UpdatedStockEvent) EventID() uuid.UUID {
	return e.id
}

func (e UpdatedStockEvent) GetProduct() *domain.Product {
	return e.product
}
