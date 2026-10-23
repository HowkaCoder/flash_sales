package handlers

import (
	"context"
	"strings"
	"time"

	pb "gateway/protoc/auth"

	"github.com/gofiber/fiber/v2"
)

func RequireAuth(client pb.AuthServiceClient) fiber.Handler {
	return func(c *fiber.Ctx) error {
		header := c.Get("Authorization")
		if !strings.HasPrefix(header, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "missing token"})
		}

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		res, err := client.ValidateJWTToken(ctx, &pb.ValidateJWTTokenRequest{Token: strings.TrimPrefix(header, "Bearer ")})
		if err != nil {
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error": "auth service is unavailable"})
		}
		if !res.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token"})
		}

		c.Locals("user_id", res.UserId)
		return c.Next()
	}
}
