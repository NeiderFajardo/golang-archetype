package infrastructure

import (
	"fmt"

	productEvents "github.com/NeiderFajardo/internal/products/domain/events"
	"github.com/NeiderFajardo/pkg/events"
)

type ProductEventHandler struct {
	// Lista de handlers a notificar
}

func NewProductEventHandler() productEvents.IProductEventHandler {
	return &ProductEventHandler{}
}

func (h *ProductEventHandler) Notify(event events.DomainEvent) {
	fmt.Printf("ProductEventHandler.Notify: %s\n", event.Name())
}
