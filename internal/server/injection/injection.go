package injection

import (
	"github.com/NeiderFajardo/internal/products/api"
	"github.com/NeiderFajardo/internal/products/application"
	"github.com/NeiderFajardo/internal/products/infrastructure"
	"github.com/NeiderFajardo/internal/server"
	"github.com/NeiderFajardo/pkg/database"
	"go.uber.org/fx"
)

func Init() {
	// Start the application
	fx.New(
		fx.Provide(
			database.NewMongoClient,
			infrastructure.NewProductRepository,
			application.NewProductService,
			api.NewProductHandler,
			server.NewServer,
			server.RegisterRoutes),
		fx.Invoke(
			server.RegisterHandlers,
			server.Run),
	).Wait()
}
