package main

import (
	"github.com/jialechen7/gorder-v2/common/broker"
	"github.com/jialechen7/gorder-v2/common/config"
	"github.com/jialechen7/gorder-v2/common/logging"
	"github.com/jialechen7/gorder-v2/common/server"
	"github.com/jialechen7/gorder-v2/payment/infra/consumer"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	logging.Init()
	if err := config.NewViperConfig(); err != nil {
		logrus.Fatal(err)
	}
}

func main() {
	serviceName := viper.GetString("payment.service-name")
	serverType := viper.GetString("payment.server-to-run")

	ch, closeFn := broker.Connect(
		viper.GetString("rabbitmq.user"),
		viper.GetString("rabbitmq.password"),
		viper.GetString("rabbitmq.host"),
		viper.GetString("rabbitmq.port"),
	)
	defer func() {
		_ = ch.Close()
		_ = closeFn()
	}()

	go consumer.NewConsumer().Listen(ch)

	paymentHandler := NewPaymentHandler()
	switch serverType {
	case "http":
		server.RunHTTPServer(serviceName, paymentHandler.RegisterRouters)
	case "grpc":
		logrus.Panic("unsupported type: grpc")
	default:
		logrus.Panic("unreachable code")
	}
}
