package handlers

import (
	"context"
	"log"
	"time"

	pb "gateway/protoc/warehouse"

	"github.com/gofiber/fiber/v2"
)

func CreateProduct(c *fiber.Ctx, client pb.WarehouseClient) error {
	var req pb.CreateProductRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.CreateProduct(ctx, &req)
	if err != nil {
		log.Printf("Error creating product: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(int(res.Status)).JSON(fiber.Map{"message": res.Msg, "status": res.Status})
}

func UpdateProduct(c *fiber.Ctx, client pb.WarehouseClient) error {
	var req pb.UpdateProductRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.UpdateProduct(ctx, &req)
	if err != nil {
		log.Printf("Error updating product: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(int(res.Status)).JSON(fiber.Map{"message": res.Msg, "status": res.Status})
}

func GetAllProducts(c *fiber.Ctx, client pb.WarehouseClient) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.GetAllProducts(ctx, &pb.GetAllProductsRequest{})
	if err != nil {
		log.Printf("Error fetching products: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"products": res.Products})
}

func GetProductById(c *fiber.Ctx, client pb.WarehouseClient) error {
	productId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.GetProductById(ctx, &pb.GetProductByIdRequest{ProductId: uint32(productId)})
	if err != nil {
		log.Printf("Error fetching product by ID: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(int(res.Status)).JSON(fiber.Map{"product": res.Product, "message": res.Msg, "status": res.Status})
}

func DeleteProduct(c *fiber.Ctx, client pb.WarehouseClient) error {
	productId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.DeleteProduct(ctx, &pb.DeleteProductRequest{ProductId: uint32(productId)})
	if err != nil {
		log.Printf("Error deleting product: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(int(res.Status)).JSON(fiber.Map{"message": res.Msg, "status": res.Status})
}

func CheckInventory(c *fiber.Ctx, client pb.WarehouseClient) error {
	productId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.CheckInventory(ctx, &pb.CheckInventoryRequest{ProductId: uint32(productId)})
	if err != nil {
		log.Printf("Error checking inventory: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"product": res.Product})
}

func UpdateInventory(c *fiber.Ctx, client pb.WarehouseClient) error {
	var request struct {
		ProductID uint `json:"product_id"`
		Quantity  uint `json:"quantity"`
		Mov       bool `json:"mov"`
	}
	var req pb.UpdateInventoryRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	req.Move = request.Mov
	req.ProductId = uint32(request.ProductID)
	req.Quantity = uint32(request.Quantity)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.UpdateInventory(ctx, &req)
	if err != nil {
		log.Printf("Error updating inventory: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(int(res.Status)).JSON(fiber.Map{"message": res.Msg, "status": res.Status})
}
