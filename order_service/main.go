package main

import (
	"log"
	"net"

	"order_service/database"
	"order_service/handler"
	ab "order_service/protoc/auth"
	pb "order_service/protoc/order"
	wb "order_service/protoc/warehouse"

	"google.golang.org/grpc"
)

func main() {
	w_conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Error with connecting Warehouse Service : %v", err)
	}

	defer w_conn.Close()

	a_conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Error with connecting Auth Service: %v", err)
	}

	a_client := ab.NewAuthServiceClient(a_conn)
	wclient := wb.NewWarehouseClient(w_conn)

	db := database.SetupDB()

	server := grpc.NewServer()
	pb.RegisterOrderServiceServer(server, &handler.OrderService{DB: db, Wclient: wclient, Aclient: a_client})

	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Error with listening server : %v ", err)
	}

	log.Println("Server listening on :50053")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Error: %v", err)
	}

}
