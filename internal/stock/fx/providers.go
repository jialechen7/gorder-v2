package fx

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.uber.org/fx"

	"github.com/jialechen7/gorder-v2/common/logging"
	todoMetrics "github.com/jialechen7/gorder-v2/common/metrics"
	"github.com/jialechen7/gorder-v2/common/tracing"
	"github.com/jialechen7/gorder-v2/stock/adapters"
	"github.com/jialechen7/gorder-v2/stock/app"
	"github.com/jialechen7/gorder-v2/stock/app/query"
	domain "github.com/jialechen7/gorder-v2/stock/domain/stock"
	"github.com/jialechen7/gorder-v2/stock/ports"
)

// Config provides configuration values
type Config struct {
	ServiceName string
	ServerType  string
	JaegerURL   string
}

func NewConfig() Config {
	return Config{
		ServiceName: viper.GetString("stock.service-name"),
		ServerType:  viper.GetString("stock.server-to-run"),
		JaegerURL:   viper.GetString("jaeger.url"),
	}
}

// NewLogger provides a structured logger
func NewLogger() *logrus.Entry {
	return logrus.NewEntry(logrus.StandardLogger())
}

// NewMetrics provides metrics client
func NewMetrics() todoMetrics.TodoMetrics {
	return todoMetrics.TodoMetrics{}
}

// NewStockRepository provides stock repository implementation
// Returns the domain interface, not the concrete type
func NewStockRepository() domain.Repository {
	return adapters.NewMemoryStockRepository()
}

// QueryHandlers aggregates all query handlers for CQRS
type QueryHandlers struct {
	CheckIfItemsInStock query.CheckIfItemsInStockHandler
	GetItems            query.GetItemsHandler
}

// NewQueryHandlers directly creates query handlers without unnecessary wrapping
func NewQueryHandlers(
	repo domain.Repository,
	logger *logrus.Entry,
	metrics todoMetrics.TodoMetrics,
) QueryHandlers {
	return QueryHandlers{
		CheckIfItemsInStock: query.NewCheckIfItemsInStockHandler(repo, logger, metrics),
		GetItems:            query.NewGetItemsHandler(repo, logger, metrics),
	}
}

// NewApplication provides the main application service
func NewApplication(queries QueryHandlers) app.Application {
	return app.Application{
		Commands: app.Commands{},
		Queries: app.Queries{
			CheckIfItemsInStock: queries.CheckIfItemsInStock,
			GetItems:            queries.GetItems,
		},
	}
}

// NewTracing provides tracing shutdown function
func NewTracing(lc fx.Lifecycle, config Config) (func(context.Context) error, error) {
	shutdown, err := tracing.InitJaegerProvider(config.JaegerURL, config.ServiceName)
	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return shutdown(ctx)
		},
	})

	return shutdown, nil
}

// ServerLifecycle manages the gRPC server lifecycle
func ServerLifecycle(
	lc fx.Lifecycle,
	config Config,
	grpcServer *ports.GRPCServer,
) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logging.Init()

			// Register to Consul
			go func() {
				deregisterFunc, err := registerToConsul(ctx, config.ServiceName)
				if err != nil {
					logrus.Fatal("Failed to register to Consul: ", err)
				}

				// Store deregister function for cleanup
				lc.Append(fx.Hook{
					OnStop: func(ctx context.Context) error {
						return deregisterFunc()
					},
				})
			}()

			// Start gRPC server based on config
			if config.ServerType == "grpc" {
				go func() {
					startGRPCServer(config.ServiceName, grpcServer)
				}()
			}

			return nil
		},
	})
}
