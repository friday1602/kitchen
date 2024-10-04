package main

import (
	handler "github.com/friday1602/kitchen/services/orders/handler/orders"
	"github.com/friday1602/kitchen/services/orders/service"
	"github.com/gofiber/fiber/v2"
)

type httpServer struct {
	addr string
}

func newHttpServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func (s *httpServer) Run() error {
	router := fiber.New()
	orderService := service.NewOrderService()
	orderHandler := handler.NewHttpOrderHandler(orderService)
	orderHandler.RegisterRouter(router)

	return router.Listen(s.addr)
}
