package main

import (
	"log"
	"net/http"

	handler "github.com/friday1602/kitchen/services/orders/handler/orders"
	"github.com/friday1602/kitchen/services/orders/service"
)

type httpServer struct {
	addr string
}

func newHttpServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()
	orderService := service.NewOrderService()
	orderHandler := handler.NewHttpOrderHandler(orderService)
	orderHandler.RegisterRouter(router)

	log.Println("Starting server on", s.addr)
	return http.ListenAndServe(s.addr, router)
}
