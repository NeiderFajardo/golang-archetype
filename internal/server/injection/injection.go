package injection

import (
	"github.com/NeiderFajardo/internal/server"
	"go.uber.org/fx"
)

func Init() {
	// Start the application
	fx.New(
		fx.Provide(
			server.RegisterMiddlewares,
			server.NewServer,
			server.RegisterRoutes),
		fx.Invoke(
			server.RegisterHandlers,
			server.Run),
	).Wait()
}
