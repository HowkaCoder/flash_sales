package main

import (
	"net"
	"warehouse_service/database"
	"warehouse_service/handler"

	"log"
	pb "warehouse_service/protoc/warehouse"

	"google.golang.org/grpc"
)

func main() {
	Database := database.SetupDB()

	server := grpc.NewServer()

	pb.RegisterWarehouseServer(server, &handler.WarehouseService{DB: Database})

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Errot with tcp listening : %v", err)
	}
	log.Println("TCP server listening on port :50052")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Error :%v", err)
	}
}
