package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "auth_service/protoc/auth"

	"google.golang.org/grpc"
)

// Constants
const (
	address = "localhost:50051"
)

func connectGRPC() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("failed to connect: %v", err)
	}
	return conn, nil
}

func registerUser(client pb.AuthServiceClient, name, email, address, password string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.RegisterRequest{
		Name:     name,
		Email:    email,
		Address:  address,
		Password: password,
	}

	res, err := client.Register(ctx, req)
	if err != nil {
		log.Fatalf("Could not register: %v", err)
	}

	fmt.Printf("Register Response: %s (Status: %d)\n", res.Msg, res.Status)
}

func loginUser(client pb.AuthServiceClient, email, password string) string {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.LoginRequest{
		Email:    email,
		Password: password,
	}

	res, err := client.Login(ctx, req)
	if err != nil {
		log.Fatalf("Could not login: %v", err)
	}

	fmt.Printf("Login Response: Token: %s, Status: %d\n", res.Token, res.Status)
	return res.Token
}

func getAllUsers(client pb.AuthServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.GetAllUsersRequest{}

	res, err := client.GetAllUsers(ctx, req)
	if err != nil {
		log.Fatalf("Could not get all users: %v", err)
	}

	fmt.Println("All Users:")
	for _, user := range res.Users {
		fmt.Printf("ID: %d, Name: %s, Email: %s, Address: %s\n", user.Id, user.Name, user.Email, user.Address)
	}
}

func getUserById(client pb.AuthServiceClient, userId uint32) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.GetUserByIdRequest{
		UserId: userId,
	}

	res, err := client.GetUserById(ctx, req)
	if err != nil {
		log.Fatalf("Could not get user by ID: %v", err)
	}

	fmt.Printf("User Info (ID: %d): %s, Name: %s, Email: %s\n", res.User.Id, res.User.Name, res.User.Email, res.User.Address)
}

func updateUser(client pb.AuthServiceClient, name, email, address string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.UpdateUserRequest{
		Name:    name,
		Email:   email,
		Address: address,
	}

	res, err := client.UpdateUser(ctx, req)
	if err != nil {
		log.Fatalf("Could not update user: %v", err)
	}

	fmt.Printf("Update Response: %s (Status: %d)\n", res.Msg, res.Status)
}

func deleteUser(client pb.AuthServiceClient, userId uint32) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.DeleteUserRequest{
		Id: userId,
	}

	res, err := client.DeleteUser(ctx, req)
	if err != nil {
		log.Fatalf("Could not delete user: %v", err)
	}

	fmt.Printf("Delete Response: %s (Status: %d)\n", res.Msg, res.Status)
}

func validateJWT(client pb.AuthServiceClient, token string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.ValidateJWTTokenRequest{
		Token: token,
	}

	res, err := client.ValidateJWTToken(ctx, req)
	if err != nil {
		log.Fatalf("Could not validate JWT token: %v", err)
	}

	fmt.Printf("JWT Valid: %t, User ID: %d\n", res.Valid, res.UserId)
}

func main() {

	conn, err := connectGRPC()
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := pb.NewAuthServiceClient(conn)

	registerUser(client, "John Doe", "j3ohn.doe@example.com", "123 Street", "password123")

	token := loginUser(client, "j3ohn.doe@example.com", "password123")

	validateJWT(client, token)

	getAllUsers(client)

	getUserById(client, 1)

	updateUser(client, "John Doe", "john.doe@example.com", "456 New Street")

	deleteUser(client, 1)
}
