package query

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/jialechen7/gorder-v2/common/decorator"
	"github.com/jialechen7/gorder-v2/common/genproto/orderpb"
	stockDomain "github.com/jialechen7/gorder-v2/stock/domain"
	domain "github.com/jialechen7/gorder-v2/stock/domain/stock"
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
	var ids []string
	for _, item := range query.Items {
		ids = append(ids, item.ID)
	}
	repoItems, err := c.stockRepo.GetItems(ctx, ids)
	if err != nil {
		return nil, err
	}
	// Check if the items are in stock
	itemMap := make(map[string]*orderpb.Item)
	for _, item := range repoItems {
		itemMap[item.ID] = item
	}
	var results []*orderpb.Item
	for _, item := range query.Items {
		if stockItem, ok := itemMap[item.ID]; ok {
			if stockItem.Quantity < item.Quantity {
				return nil, stockDomain.NewInsufficientStockError(item.ID, item.Quantity, stockItem.Quantity)
			}
		} else {
			return nil, stockDomain.NewItemNotFoundError(item.ID)
		}
		results = append(results, &orderpb.Item{
			ID:       item.ID,
			Quantity: item.Quantity,
			PriceID:  getPriceID(item.ID),
		})
	}
	return results, nil
}
