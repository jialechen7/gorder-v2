package query

import (
	"context"

	"github.com/jialechen7/gorder-v2/common/decorator"
	"github.com/jialechen7/gorder-v2/common/genproto/orderpb"
	domain "github.com/jialechen7/gorder-v2/stock/domain/stock"
	"github.com/sirupsen/logrus"
)

type CheckIfItemsInStock struct {
	Items []*orderpb.ItemWithQuantity
}

type CheckIfItemsInStockHandler decorator.QueryHandler[CheckIfItemsInStock, []*orderpb.Item]

type checkIfItemsInStockHandler struct {
	stockRepo domain.Repository
}

func NewCheckIfItemsInStockHandler(
	stockRepo domain.Repository,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,
) CheckIfItemsInStockHandler {
	if stockRepo == nil {
		panic("stockRepo is nil")
	}
	return decorator.ApplyQueryDecorators[CheckIfItemsInStock, []*orderpb.Item](
		checkIfItemsInStockHandler{stockRepo: stockRepo},
		logger,
		metricsClient,
	)
}

// TODO: 替换为真实的 PriceID
func getPriceID(_ string) string {
	return "price_1Rh9ElRxy3YYz85Em4RAlOvR"
}

func (c checkIfItemsInStockHandler) Handle(ctx context.Context, query CheckIfItemsInStock) ([]*orderpb.Item, error) {
	var results []*orderpb.Item
	for _, item := range query.Items {
		results = append(results, &orderpb.Item{
			ID:       item.ID,
			Quantity: item.Quantity,
			PriceID:  getPriceID(item.ID),
		})
	}
	return results, nil
}
