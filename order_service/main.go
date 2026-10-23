package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"order_service/database"
	"order_service/handler"
	"order_service/kafka"
	ab "order_service/protoc/auth"
	pb "order_service/protoc/order"
	wb "order_service/protoc/warehouse"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func getenv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func main() {
	wConn, err := grpc.NewClient(getenv("WAREHOUSE_ADDR", "localhost:50052"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("warehouse client: %v", err)
	}
	defer wConn.Close()

	aConn, err := grpc.NewClient(getenv("AUTH_ADDR", "localhost:50051"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("auth client: %v", err)
	}
	defer aConn.Close()

	publisher := kafka.NewPublisher(os.Getenv("KAFKA_BROKERS"), "order_service")
	defer publisher.Close()

	server := grpc.NewServer()
	pb.RegisterOrderServiceServer(server, &handler.OrderService{
		DB:        database.SetupDB(getenv("DB_PATH", "orders.db")),
		Wclient:   wb.NewWarehouseClient(wConn),
		Aclient:   ab.NewAuthServiceClient(aConn),
		Publisher: publisher,
	})

	serve(server, getenv("GRPC_ADDR", ":50053"))
}

func serve(server *grpc.Server, addr string) {
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
