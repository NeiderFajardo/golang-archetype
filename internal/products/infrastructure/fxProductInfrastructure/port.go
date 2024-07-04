package fxProductInfrastructure

import (
	"github.com/NeiderFajardo/config"
	"github.com/NeiderFajardo/internal/products/infrastructure"
	"github.com/NeiderFajardo/pkg/database"
	"go.uber.org/fx"
)

var ProductInfrastructureModule = fx.Module("productInfrastructure",
	fx.Provide(
		fx.Annotate(
			buildProductRepositoryParams,
		),
		infrastructure.NewProductRepository,
	),
)

func buildProductRepositoryParams() infrastructure.ProductRepositoryParams {

	fx.Populate(&infrastructure.ProductRepositoryParams{})
	return infrastructure.ProductRepositoryParams{
		DbClient: database.NewMongoClient(config.GetMongoConfig()),
	}
}
