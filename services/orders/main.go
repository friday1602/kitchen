package main

func main() {
	grpcServer := newGRPCServer(":50051")
	grpcServer.Run()
}