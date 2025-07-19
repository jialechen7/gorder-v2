package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jialechen7/gorder-v2/common/broker"
	"github.com/jialechen7/gorder-v2/common/config"
	"github.com/jialechen7/gorder-v2/common/discovery"
	"github.com/jialechen7/gorder-v2/common/genproto/orderpb"
	"github.com/jialechen7/gorder-v2/common/logging"
	"github.com/jialechen7/gorder-v2/common/server"
	"github.com/jialechen7/gorder-v2/common/tracing"
	"github.com/jialechen7/gorder-v2/order/infra/consumer"
	"github.com/jialechen7/gorder-v2/order/ports"
	"github.com/jialechen7/gorder-v2/order/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func init() {
	logging.Init()
	if err := config.NewViperConfig(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	serviceName := viper.GetString("order.service-name")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	shutdown, err := tracing.InitJaegerProvider(viper.GetString("jaeger.url"), serviceName)
	if err != nil {
		logrus.Fatal(err)
	}
	defer shutdown(ctx)

	application, cleanup := service.NewApplication(ctx)
	defer cleanup()

	deregisterFunc, err := discovery.RegisterToConsul(ctx, serviceName)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = deregisterFunc()
	}()

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

	go consumer.NewConsumer(application).Listen(ch)

	go server.RunGRPCServer(serviceName, func(server *grpc.Server) {
		svc := ports.NewGRPCServer(application)
		orderpb.RegisterOrderServiceServer(server, svc)
	})

	server.RunHTTPServer(serviceName, func(router *gin.Engine) {
		svc := ports.NewHTTPServer(application)
		router.StaticFile("/success", "../../public/success.html")
		ports.RegisterHandlersWithOptions(router, svc, ports.GinServerOptions{
			BaseURL:      "/api",
			Middlewares:  nil,
			ErrorHandler: nil,
		})
	})
}
