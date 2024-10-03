package handler

import (
	"net/http"

	"github.com/friday1602/kitchen/services/common/genproto/orders"
	"github.com/friday1602/kitchen/services/common/utils"
	"github.com/friday1602/kitchen/services/orders/types"
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

func (h *OrdersHttpHandler) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("POST /orders", h.CreateOrder)
}

func (h *OrdersHttpHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req orders.CreateOrderRequest
	if err := utils.ParseJson(r, &req); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	order := orders.Order{
		OrderID: 55,
		CustomerID: req.CustomerID,
		ProductID: req.ProductID,
		Quantity: req.Quantity,
	}

	if err := h.orderService.CreateOrder(r.Context(), &order); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	res := &orders.CreateOrderResponse{Status: "success"}
	utils.WriteJson(w, http.StatusOK, res)

}
