package fxProductApplication

import (
	"github.com/NeiderFajardo/internal/products/application"
	"go.uber.org/fx"
)

var ProductApplicationModule = fx.Module("productApplication",
	fx.Provide(
		// Because product service depends on the IProductRepository interface, this does not violate the dependency rule from DDD
		application.NewProductService,
	),
)
