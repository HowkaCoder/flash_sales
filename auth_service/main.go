package main

import (
	"auth_service/database"
	"auth_service/handler"
	pb "auth_service/protoc/auth"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	db := database.SetupDB()

	server := grpc.NewServer()
	jwtSecret := []byte("your_secret_key")
	pb.RegisterAuthServiceServer(server, &handler.AuthService{DB: db, JwtSecret: jwtSecret})

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Println("Listening on server :50051")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
