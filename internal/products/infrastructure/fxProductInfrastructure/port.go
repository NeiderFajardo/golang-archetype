package fxProductInfrastructure

import (
	"github.com/NeiderFajardo/config"
	"github.com/NeiderFajardo/internal/products/infrastructure"
	"github.com/NeiderFajardo/pkg/database"
	"go.uber.org/fx"
)

var ProductInfrastructureModule = fx.Module("productInfrastructure",
	fx.Provide(
		buildProductRepositoryParams,
		infrastructure.NewProductRepository,
		// This setup could be more complex, depends on the requirements
		infrastructure.NewProductEventHandler,
	),
)

func buildProductRepositoryParams() infrastructure.ProductRepositoryParams {

	fx.Populate(&infrastructure.ProductRepositoryParams{})
	return infrastructure.ProductRepositoryParams{
		DbClient: database.NewMongoClient(config.GetMongoConfig()),
	}
}
