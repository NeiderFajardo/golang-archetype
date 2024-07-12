package fxProductApi

import (
	"github.com/NeiderFajardo/internal/products/api"
	"go.uber.org/fx"
)

var ProductApiModule = fx.Module("productApi",
	fx.Provide(
		api.NewProductHandler,
		api.NewProductStockHandler,
	),
)
