package grpc

import (
	"context"
	"log"
	ab "order_service/protoc/auth"
	pb "order_service/protoc/warehouse"
	"time"
)

func GetUserById(client ab.AuthServiceClient, user_id uint32) *ab.GetUserByIdResponse {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &ab.GetUserByIdRequest{
		UserId: user_id,
	}

	res, err := client.GetUserById(ctx, req)
	if err != nil {
		log.Fatalf("Error with Get User By ID requet: %v", err)
	}

	log.Printf("User : %v , Msg : %v and Code : %v", res.User, res.Msg, res.Status)
	return res
}

func CheckInventory(client pb.WarehouseClient, product_id uint32) *pb.CheckInventoryResponse {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.CheckInventoryRequest{
		ProductId: product_id,
	}
	res, err := client.CheckInventory(ctx, req)
	if err != nil {
		log.Fatalf("Error with CheckInventory request : %v", err)
	}
	log.Printf("Product : %v ", res.Product)
	return res
}

func UpdateInventory(client pb.WarehouseClient, product_id, quantity uint32, mov bool) (string, uint32) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.UpdateInventoryRequest{
		ProductId: product_id,
		Quantity:  quantity,
		Mov:       mov,
	}

	res, err := client.UpdateInventory(ctx, req)
	if err != nil {
		log.Fatalf("Error with UpdateInventory request : %v", err)
	}

	log.Println("Msg : %v and Status : %v", res.Msg, res.Status)
	return res.Msg, res.Status
}
