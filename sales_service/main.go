package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"sales_service/database"
	"sales_service/handler"
	"sales_service/kafka"
	pb "sales_service/protoc/sales"
	wb "sales_service/protoc/warehouse"
	"syscall"

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

	publisher := kafka.NewPublisher(os.Getenv("KAFKA_BROKERS"), "sales_service")
	defer publisher.Close()

	server := grpc.NewServer()
	pb.RegisterSalesServer(server, &handler.SaleService{
		DB:        database.SetupDB(getenv("DB_PATH", "sales.db")),
		Wclient:   wb.NewWarehouseClient(wConn),
		Publisher: publisher,
	})

	addr := getenv("GRPC_ADDR", ":50054")
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
