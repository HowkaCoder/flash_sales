package handlers

import (
	"context"
	"log"
	"time"

	pb "gateway/protoc/sales"

	"github.com/gofiber/fiber/v2"
)

func CreateSales(c *fiber.Ctx, client pb.SalesClient) error {
	var req pb.CreateSalesRequest
	var request struct {
		ProductID     uint   `json:"product_id"`
		DiscountPrice uint   `json:"discount_price"`
		StartTime     string `json:"start_time"`
		EndTime       string `json:"end_time"`
	}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	req.DiscountPrice = uint32(request.DiscountPrice)
	req.ProductId = uint32(request.ProductID)
	req.StartTime = request.StartTime
	req.EndTime = request.EndTime

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.CreateSales(ctx, &req)
	if err != nil {
		log.Printf("Error creating sale: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(int(res.Status)).JSON(fiber.Map{"message": res.Msg, "status": res.Status})
}

func GetAllSales(c *fiber.Ctx, client pb.SalesClient) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.GetAllSales(ctx, &pb.GetAllSalesRequest{})
	if err != nil {
		log.Printf("Error fetching sales: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"sales": res.Sales})
}

func GetSalesById(c *fiber.Ctx, client pb.SalesClient) error {
	salesId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid sales ID"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.GetSalesById(ctx, &pb.GetSalesByIdRequest{SalesId: uint32(salesId)})
	if err != nil {
		log.Printf("Error fetching sale by ID: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(int(res.Status)).JSON(fiber.Map{"sale": res.Sale, "message": res.Msg, "status": res.Status})
}

func UpdateSales(c *fiber.Ctx, client pb.SalesClient) error {
	var req pb.UpdateSalesRequest
	var request struct {
		ID            uint   `json:id`
		ProductID     uint   `json:"product_id"`
		DiscountPrice uint   `json:"discount_price"`
		StartTime     string `json:"start_time"`
		EndTime       string `json:"end_time"`
	}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	req.DiscountPrice = uint32(request.DiscountPrice)
	req.ProductId = uint32(request.ProductID)
	req.Id = uint32(request.ID)
	req.StartTime = request.StartTime
	req.EndTime = request.EndTime

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.UpdateSales(ctx, &req)
	if err != nil {
		log.Printf("Error updating sale: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(int(res.Status)).JSON(fiber.Map{"message": res.Msg, "status": res.Status})
}

func DeleteSales(c *fiber.Ctx, client pb.SalesClient) error {
	salesId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid sales ID"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.DeleteSales(ctx, &pb.DeleteSalesRequest{SalesId: uint32(salesId)})
	if err != nil {
		log.Printf("Error deleting sale: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(int(res.Status)).JSON(fiber.Map{"message": res.Msg, "status": res.Status})
}
