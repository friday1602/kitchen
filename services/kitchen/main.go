package main

func main() {
	httpsvr := newHttpServer(":9000")
	httpsvr.Run()
}
