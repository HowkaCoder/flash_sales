package grpc

import (
	"context"
	"fmt"
	ab "order_service/protoc/auth"
	pb "order_service/protoc/warehouse"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const timeout = 3 * time.Second

func UserExists(client ab.AuthServiceClient, userID uint32) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	res, err := client.GetUserById(ctx, &ab.GetUserByIdRequest{UserId: userID})
	if status.Code(err) == codes.NotFound {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("auth service: %w", err)
	}
	return res.User != nil && res.User.Id != 0, nil
}

func ChangeInventory(client pb.WarehouseClient, productID, quantity uint32, add bool) (uint32, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	res, err := client.UpdateInventory(ctx, &pb.UpdateInventoryRequest{
		ProductId: productID,
		Quantity:  quantity,
		Mov:       add,
	})
	if err != nil {
		return 0, "", fmt.Errorf("warehouse service: %w", err)
	}
	return res.Status, res.Msg, nil
}
