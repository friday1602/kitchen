package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/friday1602/kitchen/services/common/genproto/orders"
	"github.com/friday1602/kitchen/services/common/utils"
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
	mux := http.NewServeMux()

	conn := newGrpcClient(":50051")
	defer conn.Close()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c := orders.NewOrderServiceClient(conn)
		ctx, cancel := context.WithTimeout(r.Context(), time.Second)
		defer cancel()
		_, err := c.CreateOrder(ctx, &orders.CreateOrderRequest{
			CustomerID: 55,
			ProductID:  55,
			Quantity:   11,
		})
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}
		resp, err := c.GetOrder(ctx, &orders.GetOrderRequest{
			CustomerID: 55,
		})
		if err = utils.WriteJson(w, http.StatusOK, resp.GetOrders()); err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}

	})
	log.Println("starting server on", s.addr)
	return http.ListenAndServe(s.addr, mux)
}

func newGrpcClient(addr string) *grpc.ClientConn {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("error connect to grpc: %v", err)
	}
	return conn
}
