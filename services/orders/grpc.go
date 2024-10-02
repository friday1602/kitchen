package main

import (
	"log"
	"net"

	handler "github.com/friday1602/kitchen/services/orders/handler/orders"
	"github.com/friday1602/kitchen/services/orders/service"
	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func newGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}

func (s *gRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	// register grpc services
	orderService := service.NewOrderService()
	handler.NewGrpcOrdersService(grpcServer, orderService)

	log.Println("Starting gRPC server on", s.addr)
	return grpcServer.Serve(lis)
}
