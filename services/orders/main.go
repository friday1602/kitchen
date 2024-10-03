package main

func main() {
	httpSrv := newHttpServer(":8000")
	go httpSrv.Run()


	grpcServer := newGRPCServer(":50051")
	grpcServer.Run()
}