package convertor

import "sync"

var (
	orderConvertor *OrderConvertor
	orderOnce      sync.Once
)

var (
	itemConvertor *ItemConvertor
	itemOnce      sync.Once
)

var (
	itemWithQuantityConvertor *ItemWithQuantityConvertor
	itemWithQuantityOnce      sync.Once
)

func NewOrderConvertor() *OrderConvertor {
	orderOnce.Do(func() {
		orderConvertor = &OrderConvertor{}
	})
	return orderConvertor
}

func NewItemConvertor() *ItemConvertor {
	itemOnce.Do(func() {
		itemConvertor = &ItemConvertor{}
	})
	return itemConvertor
}

func NewItemWithQuantityConvertor() *ItemWithQuantityConvertor {
	itemWithQuantityOnce.Do(func() {
		itemWithQuantityConvertor = &ItemWithQuantityConvertor{}
	})
	return itemWithQuantityConvertor
}
