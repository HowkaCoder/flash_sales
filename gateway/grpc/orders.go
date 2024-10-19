package grpc

import (
	"context"
	ob "gateway/protoc/order"
	"log"
	"time"
)

func CreateOrder(client ob.OrderServiceClient, userId, productId, quantity uint32, status string) (string, uint32) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &ob.CreateOrderRequest{
		UserId:    userId,
		ProductId: productId,
		Quantity:  quantity,
		Status:    status,
	}

	res, err := client.CreateOrder(ctx, req)
	if err != nil {
		log.Fatalf("Error with CreateOrder request: %v", err)
	}

	log.Printf("Response Message: %v and Response code: %v", res.Msg, res.Status)
	return res.Msg, res.Status
}

func UpdateOrder(client ob.OrderServiceClient, id, userId, productId, quantity uint32, status string) (string, uint32) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &ob.UpdateOrderRequest{
		Id:        id,
		UserId:    userId,
		ProductId: productId,
		Quantity:  quantity,
		Status:    status,
	}

	res, err := client.UpdateOrder(ctx, req)
	if err != nil {
		log.Fatalf("Error with UpdateOrder request: %v", err)
	}

	log.Printf("Response Message: %v and Response code: %v", res.Msg, res.Status)
	return res.Msg, res.Status
}

func GetAllOrders(client ob.OrderServiceClient) ([]*ob.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &ob.GetAllOrdersRequest{}

	res, err := client.GetAllOrders(ctx, req)
	if err != nil {
		log.Fatalf("Error with GetAllOrders request: %v", err)
		return nil, err
	}

	log.Printf("Received %d orders", len(res.Orders))
	return res.Orders, nil
}

func GetOrderById(client ob.OrderServiceClient, orderId uint32) (*ob.Order, string, uint32) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &ob.GetOrderByIdRequest{
		OrderId: orderId,
	}

	res, err := client.GetOrderById(ctx, req)
	if err != nil {
		log.Fatalf("Error with GetOrderById request: %v", err)
	}

	log.Printf("Response Message: %v and Response code: %v", res.Msg, res.Status)
	return res.Order, res.Msg, res.Status
}

func DeleteOrder(client ob.OrderServiceClient, orderId uint32) (string, uint32) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &ob.DeleteOrderRequest{
		OrderId: orderId,
	}

	res, err := client.DeleteOrder(ctx, req)
	if err != nil {
		log.Fatalf("Error with DeleteOrder request: %v", err)
	}

	log.Printf("Response Message: %v and Response code: %v", res.Msg, res.Status)
	return res.Msg, res.Status
}
