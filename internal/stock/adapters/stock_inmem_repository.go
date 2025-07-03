package adapters

import (
	"context"
	"sync"

	"github.com/jialechen7/gorder-v2/common/genproto/orderpb"
	domain "github.com/jialechen7/gorder-v2/stock/domain/stock"
)

type MemoryStockRepository struct {
	lock  *sync.RWMutex
	store map[string]*orderpb.Item
}

var stub = map[string]*orderpb.Item{
	"item_id": {
		ID:       "foo_item",
		Name:     "stub_item",
		Quantity: 10000,
		PriceID:  "stub_item_price_id",
	},
	"item1": {
		ID:       "item1",
		Name:     "stub_item1",
		Quantity: 10000,
		PriceID:  "stub_item1_price_id",
	},
	"item2": {
		ID:       "item2",
		Name:     "stub_item2",
		Quantity: 10000,
		PriceID:  "stub_item2_price_id",
	},
	"item3": {
		ID:       "item3",
		Name:     "stub_item3",
		Quantity: 10000,
		PriceID:  "stub_item3_price_id",
	},
}

func NewMemoryStockRepository() *MemoryStockRepository {
	return &MemoryStockRepository{
		lock:  &sync.RWMutex{},
		store: stub,
	}
}

func (m MemoryStockRepository) GetItems(ctx context.Context, ids []string) ([]*orderpb.Item, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	var (
		items   []*orderpb.Item
		missing []string
	)
	for _, id := range ids {
		if item, ok := m.store[id]; ok {
			items = append(items, item)
		} else {
			missing = append(missing, id)
		}
	}
	if len(items) == len(ids) {
		return items, nil
	}
	return nil, domain.NotFoundError{
		Missing: missing,
	}
}
