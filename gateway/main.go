package main

import (
	"log"

	handlers "gateway/handler"
	pb "gateway/protoc/auth"
	ob "gateway/protoc/order"
	sb "gateway/protoc/sales"
	wb "gateway/protoc/warehouse"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
)

func main() {
	app := fiber.New()

	a_conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect to Auth service: %v", err)
	}
	defer a_conn.Close()

	w_conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect to Warehouse service: %v", err)
	}
	defer w_conn.Close()

	o_conn, err := grpc.Dial("localhost:50053", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect to order serviCe : %v", err)
	}
	defer o_conn.Close()

	s_conn, err := grpc.Dial("localhost:50054", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect to sAles service: %v , err")
	}
	defer s_conn.Close()

	a_client := pb.NewAuthServiceClient(a_conn)
	w_client := wb.NewWarehouseClient(w_conn)
	o_client := ob.NewOrderServiceClient(o_conn)
	s_client := sb.NewSalesClient(s_conn)

	app.Delete("/sales/:id", func(c *fiber.Ctx) error {
		return handlers.DeleteSales(c, s_client)
	})

	app.Patch("/sales", func(c *fiber.Ctx) error {
		return handlers.UpdateSales(c, s_client)
	})

	app.Post("/sales", func(c *fiber.Ctx) error {
		return handlers.CreateSales(c, s_client)
	})

	app.Get("/sales/:id", func(c *fiber.Ctx) error {
		return handlers.GetSalesById(c, s_client)
	})

	app.Get("/sales", func(c *fiber.Ctx) error {
		return handlers.GetAllSales(c, s_client)
	})

	app.Delete("/orders/:id", func(c *fiber.Ctx) error {
		return handlers.DeleteOrder(c, o_client)
	})

	app.Patch("/orders", func(c *fiber.Ctx) error {
		return handlers.UpdateOrder(c, o_client)
	})

	app.Post("/orders", func(c *fiber.Ctx) error {
		return handlers.CreateOrder(c, o_client)
	})

	app.Get("/orders/:id", func(c *fiber.Ctx) error {
		return handlers.GetOrderById(c, o_client)
	})

	app.Get("/orders", func(c *fiber.Ctx) error {
		return handlers.GetAllOrders(c, o_client)
	})

	app.Post("/register", func(c *fiber.Ctx) error {
		return handlers.RegisterUser(c, a_client)
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		return handlers.LoginUser(c, a_client)
	})

	app.Get("/users", func(c *fiber.Ctx) error {
		return handlers.GetAllUsers(c, a_client)
	})

	app.Get("/users/:id", func(c *fiber.Ctx) error {
		return handlers.GetUserById(c, a_client)
	})

	app.Patch("/users", func(c *fiber.Ctx) error {
		return handlers.UpdateUser(c, a_client)
	})

	app.Delete("/users/:id", func(c *fiber.Ctx) error {
		return handlers.DeleteUser(c, a_client)
	})

	app.Post("/validate", func(c *fiber.Ctx) error {
		return handlers.ValidateJWT(c, a_client)
	})

	app.Post("/warehouse/product", func(c *fiber.Ctx) error {
		return handlers.CreateProduct(c, w_client)
	})

	app.Patch("/warehouse/product", func(c *fiber.Ctx) error {
		return handlers.UpdateProduct(c, w_client)
	})

	app.Get("/warehouse/products", func(c *fiber.Ctx) error {
		return handlers.GetAllProducts(c, w_client)
	})

	app.Get("/warehouse/product/:id", func(c *fiber.Ctx) error {
		return handlers.GetProductById(c, w_client)
	})

	app.Delete("/warehouse/product/:id", func(c *fiber.Ctx) error {
		return handlers.DeleteProduct(c, w_client)
	})

	app.Get("/warehouse/product/:id/inventory", func(c *fiber.Ctx) error {
		return handlers.CheckInventory(c, w_client)
	})

	app.Patch("/warehouse/product/inventory", func(c *fiber.Ctx) error {
		return handlers.UpdateInventory(c, w_client)
	})

	log.Fatal(app.Listen(":3000"))
}
