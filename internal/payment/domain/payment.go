package domain

import (
	"context"

	"github.com/jialechen7/gorder-v2/common/genproto/orderpb"
)

type Processor interface {
	CreatePaymentLink(ctx context.Context, order *orderpb.Order) (string, error)
}
