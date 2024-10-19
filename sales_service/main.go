package main

import (
	"log"
	"net"
	"sales_service/database"
	"sales_service/handler"
	pb "sales_service/protoc/sales"
	wb "sales_service/protoc/warehouse"

	"google.golang.org/grpc"
)

func main() {
	w_conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Error with connecting Warehouse service: %v", err)
	}

	defer w_conn.Close()

	w_client := wb.NewWarehouseClient(w_conn)

	db := database.SetupDB()

	server := grpc.NewServer()
	pb.RegisterSalesServer(server, &handler.SaleService{DB: db, Wclient: w_client})

	lis, err := net.Listen("tcp", ":50054")
	if err != nil {
		log.Fatalf("Error with listening server: %v", err)
	}

	log.Println("Server listening on :50054")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Error %v", err)
	}
}
