package main

import (
	"context"
	"log"

	"github.com/jialechen7/gorder-v2/common/config"
	"github.com/jialechen7/gorder-v2/common/discovery"
	"github.com/jialechen7/gorder-v2/common/genproto/stockpb"
	"github.com/jialechen7/gorder-v2/common/logging"
	"github.com/jialechen7/gorder-v2/common/server"
	"github.com/jialechen7/gorder-v2/stock/ports"
	"github.com/jialechen7/gorder-v2/stock/service"
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
	serviceName := viper.GetString("stock.service-name")
	serverType := viper.GetString("stock.server-to-run")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	application := service.NewApplication(ctx)

	deregisterFunc, err := discovery.RegisterToConsul(ctx, serviceName)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = deregisterFunc()
	}()

	switch serverType {
	case "grpc":
		server.RunGRPCServer(serviceName, func(server *grpc.Server) {
			svc := ports.NewGRPCServer(application)
			stockpb.RegisterStockServiceServer(server, svc)
		})
	case "http":
		// TODO: 暂时不用
	default:
		panic("unexpected server type")
	}
}
