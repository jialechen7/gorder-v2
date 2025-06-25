package stock

import (
	"context"
	"fmt"

	"github.com/jialechen7/gorder-v2/common/genproto/orderpb"
)

type Repository interface {
	GetItems(ctx context.Context, ids []string) ([]*orderpb.Item, error)
}

type NotFoundError struct {
	Missing []string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("these stock items not found: %v", e.Missing)
}
