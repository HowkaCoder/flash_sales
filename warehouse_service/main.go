package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"warehouse_service/database"
	"warehouse_service/handler"
	pb "warehouse_service/protoc/warehouse"

	"google.golang.org/grpc"
)

func getenv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func main() {
	server := grpc.NewServer()
	pb.RegisterWarehouseServer(server, &handler.WarehouseService{DB: database.SetupDB(getenv("DB_PATH", "warehouse.db"))})

	addr := getenv("GRPC_ADDR", ":50052")
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("listen %s: %v", addr, err)
	}

	go func() {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
		<-stop
		server.GracefulStop()
	}()

	log.Printf("listening on %s", addr)
	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
