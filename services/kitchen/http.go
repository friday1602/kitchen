package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/friday1602/kitchen/services/common/genproto/orders"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type httpServer struct {
	addr string
}

func newHttpServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func (s *httpServer) Run() error {
	mux := fiber.New()

	conn := newGrpcClient(":50051")
	defer conn.Close()

	mux.Get("/", func(c *fiber.Ctx) error {
		cGRPC := orders.NewOrderServiceClient(conn)
		ctx, cancel := context.WithTimeout(c.Context(), time.Second)
		defer cancel()
		_, err := cGRPC.CreateOrder(ctx, &orders.CreateOrderRequest{
			CustomerID: 55,
			ProductID:  55,
			Quantity:   11,
		})
		if err != nil {
			return fiber.NewError(http.StatusInternalServerError, err.Error())
		}
		resp, err := cGRPC.GetOrder(ctx, &orders.GetOrderRequest{
			CustomerID: 55,
		})
		return c.JSON(resp)
	})
	return mux.Listen(s.addr)
}

func newGrpcClient(addr string) *grpc.ClientConn {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("error connect to grpc: %v", err)
	}
	return conn
}
