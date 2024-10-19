package handlers

import (
	"context"
	"log"
	"time"

	pb "gateway/protoc/auth"

	"github.com/gofiber/fiber/v2"
)

func RegisterUser(c *fiber.Ctx, client pb.AuthServiceClient) error {
	var req pb.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.Register(ctx, &req)
	if err != nil {
		log.Printf("Error during registration: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(int(res.Status)).JSON(fiber.Map{"message": res.Msg, "status": res.Status})
}

func LoginUser(c *fiber.Ctx, client pb.AuthServiceClient) error {
	var req pb.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.Login(ctx, &req)
	if err != nil {
		log.Printf("Error during login: %v\n", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(int(res.Status)).JSON(fiber.Map{"token": res.Token, "status": res.Status})
}

func GetAllUsers(c *fiber.Ctx, client pb.AuthServiceClient) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.GetAllUsers(ctx, &pb.GetAllUsersRequest{})
	if err != nil {
		log.Printf("Could not get all users: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(res.Users)
}

func GetUserById(c *fiber.Ctx, client pb.AuthServiceClient) error {
	userId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.GetUserById(ctx, &pb.GetUserByIdRequest{UserId: uint32(userId)})
	if err != nil {
		log.Printf("Could not get user by ID: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(res.User)
}

func UpdateUser(c *fiber.Ctx, client pb.AuthServiceClient) error {
	var req pb.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.UpdateUser(ctx, &req)
	if err != nil {
		log.Printf("Could not update user: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(int(res.Status)).JSON(fiber.Map{"message": res.Msg, "status": res.Status})
}

func DeleteUser(c *fiber.Ctx, client pb.AuthServiceClient) error {
	userId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.DeleteUser(ctx, &pb.DeleteUserRequest{Id: uint32(userId)})
	if err != nil {
		log.Printf("Could not delete user: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(int(res.Status)).JSON(fiber.Map{"message": res.Msg, "status": res.Status})
}

func ValidateJWT(c *fiber.Ctx, client pb.AuthServiceClient) error {
	var req pb.ValidateJWTTokenRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.ValidateJWTToken(ctx, &req)
	if err != nil {
		log.Printf("Could not validate JWT token: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"valid": res.Valid, "userId": res.UserId})
}
