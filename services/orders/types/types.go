package types

import (
	"context"

	"github.com/friday1602/kitchen/services/common/genproto/orders"
)

type OrderService interface {
	CreateOrder(context.Context, *orders.Order) error
	GetOrder(context.Context) []*orders.Order
}
