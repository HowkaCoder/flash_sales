package grpc

import (
	"context"
	"fmt"
	pb "gateway/protoc/auth"
	"log"
	"time"
)

func registerUser(client pb.AuthServiceClient, name, address, email, password string) (string, uint32) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.RegisterRequest{
		Name:     name,
		Address:  address,
		Email:    email,
		Password: password,
	}
	res, err := client.Register(ctx, req)
	if err != nil {
		log.Fatalf("Error : %v\n", err)
	}

	log.Printf("Response Message: %v  and Response Status %v \n", res.Msg, res.Status)
	return res.Msg, res.Status
}

func loginUser(client pb.AuthServiceClient, email, password string) (string, uint32) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.LoginRequest{
		Email:    email,
		Password: password,
	}

	res, err := client.Login(ctx, req)
	if err != nil {
		log.Fatalf("Error :%v", err)
	}

	log.Printf("Token : %v and Status Code : %v", res.Token, res.Status)

	return res.Token, res.Status

}
func GetAllUsers(client pb.AuthServiceClient) []*pb.User {
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
	return res.Users
}

func getUserById(client pb.AuthServiceClient, userId uint32) *pb.User {
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
	return res.User
}

func updateUser(client pb.AuthServiceClient, name, email, address string) (string, uint32) {
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
	return res.Msg, res.Status
}

func deleteUser(client pb.AuthServiceClient, userId uint32) (string, uint32) {
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
	return res.Msg, res.Status
}

func validateJWT(client pb.AuthServiceClient, token string) (bool, uint32) {
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
	return res.Valid, res.UserId
}
