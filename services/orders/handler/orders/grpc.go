package handler

import (
	"context"

	"github.com/friday1602/kitchen/services/common/genproto/orders"
	"github.com/friday1602/kitchen/services/orders/types"
	"google.golang.org/grpc"
)

type OrdersGrpcHandler struct {
	orderService types.OrderService
	orders.UnimplementedOrderServiceServer
}

func NewGrpcOrdersService(grpc *grpc.Server, orderService types.OrderService) {
	grpcHandler := &OrdersGrpcHandler{
		orderService: orderService,
	}

	// register the OrderServiceServer
	orders.RegisterOrderServiceServer(grpc, grpcHandler)
}

func (h *OrdersGrpcHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	order := &orders.Order{
		OrderID:    52,
		CustomerID: 2,
		ProductID:  1,
		Quantity:   5,
	}

	if err := h.orderService.CreateOrder(ctx, order); err != nil {
		return nil, err
	}

	res := &orders.CreateOrderResponse{
		Status: "success",
	}
	return res, nil
}
