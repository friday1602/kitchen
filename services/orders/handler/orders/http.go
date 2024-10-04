package handler

import (
	"net/http"

	"github.com/friday1602/kitchen/services/common/genproto/orders"
	"github.com/friday1602/kitchen/services/orders/types"
	"github.com/gofiber/fiber/v2"
)

type OrdersHttpHandler struct {
	orderService types.OrderService
}

func NewHttpOrderHandler(orderService types.OrderService) *OrdersHttpHandler {
	handler := &OrdersHttpHandler{
		orderService: orderService,
	}

	return handler
}

func (h *OrdersHttpHandler) RegisterRouter(router *fiber.App) {
	router.Post("/orders", h.CreateOrder)
}

func (h *OrdersHttpHandler) CreateOrder(c *fiber.Ctx) error {
	var req orders.CreateOrderRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	order := orders.Order{
		OrderID:    55,
		CustomerID: req.CustomerID,
		ProductID:  req.ProductID,
		Quantity:   req.Quantity,
	}

	if err := h.orderService.CreateOrder(c.Context(), &order); err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	res := &orders.CreateOrderResponse{Status: "success"}
	return c.JSON(res)

}
