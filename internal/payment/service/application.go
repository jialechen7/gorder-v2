package service

import (
	"context"

	grpcClient "github.com/jialechen7/gorder-v2/common/client"
	todoMetrics "github.com/jialechen7/gorder-v2/common/metrics"
	"github.com/jialechen7/gorder-v2/payment/adapters"
	"github.com/jialechen7/gorder-v2/payment/app"
	"github.com/jialechen7/gorder-v2/payment/app/command"
	"github.com/jialechen7/gorder-v2/payment/domain"
	"github.com/jialechen7/gorder-v2/payment/infra/processor"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewApplication(ctx context.Context) (app.Application, func()) {
	orderClient, closeStockClient, err := grpcClient.NewOrderGRPCClient(ctx)
	if err != nil {
		panic(err)
	}
	orderGRPC := adapters.NewOrderGRPC(orderClient)
	stripeProcessor := processor.NewStripeProcessor(viper.GetString("stripe_key"))
	return newApplication(ctx, orderGRPC, stripeProcessor), func() {
		_ = closeStockClient()
	}
}

func newApplication(_ context.Context, orderGRPC command.OrderService, memoryProcessor domain.Processor) app.Application {
	logger := logrus.NewEntry(logrus.StandardLogger())
	metricsClient := todoMetrics.TodoMetrics{}
	return app.Application{
		Commands: app.Commands{
			CreatePayment: command.NewCreatePaymentHandler(
				memoryProcessor,
				orderGRPC,
				logger,
				metricsClient,
			),
		},
	}
}
