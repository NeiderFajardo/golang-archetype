package injection

import (
	"github.com/NeiderFajardo/internal/products/api/fxProductApi"
	"github.com/NeiderFajardo/internal/products/application/fxProductApplication"
	"github.com/NeiderFajardo/internal/products/infrastructure/fxProductInfrastructure"
	"github.com/NeiderFajardo/internal/server"
	"go.uber.org/fx"
)

func Init() {

	// Start the application
	fx.New(
		fx.Options(
			fxProductInfrastructure.ProductInfrastructureModule,
			fxProductApplication.ProductApplicationModule,
			fxProductApi.ProductApiModule,
		),
		fx.Provide(
			server.NewServer,
			server.RegisterRoutes),
		fx.Invoke(
			server.RegisterHandlers,
			server.RunHttpServer),
	).Wait()
}
