package handlers

import (
	"context"
	"log"
	"time"

	pb "gateway/protoc/order"

	"github.com/gofiber/fiber/v2"
)

func CreateOrder(c *fiber.Ctx, client pb.OrderServiceClient) error {
	var request struct {
		UserID    uint   `json:"user_id"`
		ProductID uint   `json:"product_id"`
		Quantity  uint   `json:"quantity"`
		Status    string `json:"status"`
	}
	var req pb.CreateOrderRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req.ProductId = uint32(request.ProductID)
	req.Quantity = uint32(request.Quantity)
	req.UserId = uint32(request.UserID)
	req.Status = request.Status

	res, err := client.CreateOrder(ctx, &req)
	if err != nil {
		log.Printf("Error creating order: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(int(res.Status)).JSON(fiber.Map{"message": res.Msg, "status": res.Status})
}

func UpdateOrder(c *fiber.Ctx, client pb.OrderServiceClient) error {
	var req pb.UpdateOrderRequest
	var request struct {
		ID        uint   `json:"id"`
		UserId    uint   `json:"user_id"`
		ProductID uint   `json:"product_id"`
		Quantity  uint   `json:"quantity"`
		Status    string `json:"status"`
	}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req.Id = uint32(request.ID)
	req.ProductId = uint32(request.ProductID)
	req.Quantity = uint32(request.Quantity)
	req.UserId = uint32(request.UserId)
	req.Status = request.Status

	res, err := client.UpdateOrder(ctx, &req)
	if err != nil {
		log.Printf("Error updating order: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(int(res.Status)).JSON(fiber.Map{"message": res.Msg, "status": res.Status})
}

func GetAllOrders(c *fiber.Ctx, client pb.OrderServiceClient) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.GetAllOrders(ctx, &pb.GetAllOrdersRequest{})
	if err != nil {
		log.Printf("Error fetching orders: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"orders": res.Orders})
}

func GetOrderById(c *fiber.Ctx, client pb.OrderServiceClient) error {
	orderId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid order ID"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.GetOrderById(ctx, &pb.GetOrderByIdRequest{OrderId: uint32(orderId)})
	if err != nil {
		log.Printf("Error fetching order by ID: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(int(res.Status)).JSON(fiber.Map{"order": res.Order, "message": res.Msg, "status": res.Status})
}

func DeleteOrder(c *fiber.Ctx, client pb.OrderServiceClient) error {
	orderId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid order ID"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.DeleteOrder(ctx, &pb.DeleteOrderRequest{OrderId: uint32(orderId)})
	if err != nil {
		log.Printf("Error deleting order: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(int(res.Status)).JSON(fiber.Map{"message": res.Msg, "status": res.Status})
}
