package service

import (
	"context"

	todoMetrics "github.com/jialechen7/gorder-v2/common/metrics"
	"github.com/jialechen7/gorder-v2/order/adapters"
	"github.com/jialechen7/gorder-v2/order/app"
	"github.com/jialechen7/gorder-v2/order/app/command"
	"github.com/jialechen7/gorder-v2/order/app/query"
	"github.com/sirupsen/logrus"
)

func NewApplication(ctx context.Context) app.Application {
	orderInmemRepo := adapters.NewMemoryOrderRepository()
	logger := logrus.NewEntry(logrus.StandardLogger())
	metricsClient := todoMetrics.TodoMetrics{}
	return app.Application{
		Commands: app.Commands{
			CreateOrder: command.NewCreateOrderHandler(orderInmemRepo, logger, metricsClient),
			UpdateOrder: command.NewUpdateOrderHandler(orderInmemRepo, logger, metricsClient),
		},
		Queries: app.Queries{
			GetCustomerOrder: query.NewGetCustomerOrderHandler(orderInmemRepo, logger, metricsClient),
		},
	}
}
