package main

import (
	"go.uber.org/fx"

	_ "github.com/jialechen7/gorder-v2/common/config"
	stockfx "github.com/jialechen7/gorder-v2/stock/fx"
	"github.com/jialechen7/gorder-v2/stock/ports"
)

func main() {
	fx.New(
		// Provide all dependencies
		fx.Provide(
			stockfx.NewConfig,
			stockfx.NewLogger,
			stockfx.NewMetrics,
			stockfx.NewStockRepository,
			// CQRS: Aggregate query handlers
			stockfx.NewQueryHandlers,
			stockfx.NewApplication,
			// Direct use of ports constructor - no unnecessary wrapper
			ports.NewGRPCServer,
			stockfx.NewTracing,
		),

		// Start server with lifecycle management
		fx.Invoke(stockfx.ServerLifecycle),
	).Run()
}
