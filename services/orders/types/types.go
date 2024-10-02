package types

import (
	"context"

	"github.com/friday1602/kitchen/services/common/genproto/orders"
)

type OrderService interface {
	CreateOrder(ctx context.Context, order *orders.Order) error
}
