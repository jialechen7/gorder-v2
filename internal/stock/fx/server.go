package fx

import (
	"context"

	"google.golang.org/grpc"

	"github.com/jialechen7/gorder-v2/common/discovery"
	"github.com/jialechen7/gorder-v2/common/genproto/stockpb"
	"github.com/jialechen7/gorder-v2/common/server"
	"github.com/jialechen7/gorder-v2/stock/ports"
)

func registerToConsul(ctx context.Context, serviceName string) (func() error, error) {
	return discovery.RegisterToConsul(ctx, serviceName)
}

func startGRPCServer(serviceName string, grpcServer *ports.GRPCServer) {
	server.RunGRPCServer(serviceName, func(server *grpc.Server) {
		stockpb.RegisterStockServiceServer(server, grpcServer)
	})
}
