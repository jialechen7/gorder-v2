package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jialechen7/gorder-v2/common/config"
	"github.com/jialechen7/gorder-v2/common/genproto/orderpb"
	"github.com/jialechen7/gorder-v2/common/server"
	"github.com/jialechen7/gorder-v2/order/ports"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func init() {
	if err := config.NewViperConfig(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	serviceName := viper.GetString("order.service-name")

	go server.RunGRPCServer(serviceName, func(server *grpc.Server) {
		svc := ports.NewGRPCServer()
		orderpb.RegisterOrderServiceServer(server, svc)
	})

	server.RunHTTPServer(serviceName, func(router *gin.Engine) {
		svc := ports.NewHTTPServer()
		ports.RegisterHandlersWithOptions(router, svc, ports.GinServerOptions{
			BaseURL:      "/api",
			Middlewares:  nil,
			ErrorHandler: nil,
		})
	})
}
