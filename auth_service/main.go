package main

import (
	"auth_service/database"
	"auth_service/handler"
	"auth_service/kafka"
	pb "auth_service/protoc/auth"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
)

func getenv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func main() {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET is not set")
	}

	publisher := kafka.NewPublisher(os.Getenv("KAFKA_BROKERS"), "auth_service")
	defer publisher.Close()

	server := grpc.NewServer()
	pb.RegisterAuthServiceServer(server, &handler.AuthService{
		DB:        database.SetupDB(getenv("DB_PATH", "auth.db")),
		JwtSecret: []byte(secret),
		Publisher: publisher,
	})

	addr := getenv("GRPC_ADDR", ":50051")
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
