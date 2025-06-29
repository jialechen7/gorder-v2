package ports

import (
	context "context"

	"github.com/jialechen7/gorder-v2/common/genproto/orderpb"
	"github.com/jialechen7/gorder-v2/common/genproto/stockpb"
	"github.com/jialechen7/gorder-v2/stock/app"
	"github.com/sirupsen/logrus"
)

type GRPCServer struct {
	app app.Application
}

func NewGRPCServer(app app.Application) *GRPCServer {
	return &GRPCServer{app: app}
}

func (G GRPCServer) GetItems(ctx context.Context, request *stockpb.GetItemsRequest) (*stockpb.GetItemsResponse, error) {
	logrus.Info("rpc_request_in, stock.GetItems")
	defer func() {
		logrus.Info("rpc_request_out, stock.GetItems")
	}()
	return &stockpb.GetItemsResponse{
		Items: []*orderpb.Item{
			{
				ID: "fake-id-from-stock-service",
			},
		},
	}, nil
}

func (G GRPCServer) CheckIfItemsInStock(ctx context.Context, request *stockpb.CheckIfItemsInStockRequest) (*stockpb.CheckIfItemsInStockResponse, error) {
	logrus.Info("rpc_request_in, stock.CheckIfItemsInStock")
	defer func() {
		logrus.Info("rpc_request_out, stock.CheckIfItemsInStock")
	}()
	return nil, nil
}
