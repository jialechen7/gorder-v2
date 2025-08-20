package adapters

import (
	"context"

	"github.com/jialechen7/gorder-v2/common/genproto/orderpb"
	"github.com/jialechen7/gorder-v2/common/tracing"
	"github.com/sirupsen/logrus"
)

type OrderGRPC struct {
	client orderpb.OrderServiceClient
}

func (o OrderGRPC) UpdateOrder(ctx context.Context, order *orderpb.Order) error {
	ctx, span := tracing.Start(ctx, "order_grpc.update_order")
	defer span.End()
	_, err := o.client.UpdateOrder(ctx, order)
	if err != nil {
		logrus.Infof("payment_adapters||update_order, err=%v", err)
	}
	logrus.Infof("payment_adapters||update_order success")
	return err
}

func NewOrderGRPC(client orderpb.OrderServiceClient) *OrderGRPC {
	return &OrderGRPC{client: client}
}
