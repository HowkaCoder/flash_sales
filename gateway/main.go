package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	handlers "gateway/handler"
	pb "gateway/protoc/auth"
	ob "gateway/protoc/order"
	sb "gateway/protoc/sales"
	wb "gateway/protoc/warehouse"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func getenv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func dial(addr string) *grpc.ClientConn {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("grpc client %s: %v", addr, err)
	}
	return conn
}

func main() {
	aConn := dial(getenv("AUTH_ADDR", "localhost:50051"))
	wConn := dial(getenv("WAREHOUSE_ADDR", "localhost:50052"))
	oConn := dial(getenv("ORDER_ADDR", "localhost:50053"))
	sConn := dial(getenv("SALES_ADDR", "localhost:50054"))
	defer aConn.Close()
	defer wConn.Close()
	defer oConn.Close()
	defer sConn.Close()

	authClient := pb.NewAuthServiceClient(aConn)
	warehouseClient := wb.NewWarehouseClient(wConn)
	orderClient := ob.NewOrderServiceClient(oConn)
	salesClient := sb.NewSalesClient(sConn)

	app := fiber.New()
	app.Use(recover.New())
	app.Use(logger.New())

	auth := handlers.RequireAuth(authClient)

	app.Post("/register", func(c *fiber.Ctx) error { return handlers.RegisterUser(c, authClient) })
	app.Post("/login", func(c *fiber.Ctx) error { return handlers.LoginUser(c, authClient) })
	app.Post("/validate", func(c *fiber.Ctx) error { return handlers.ValidateJWT(c, authClient) })
	app.Get("/users", auth, func(c *fiber.Ctx) error { return handlers.GetAllUsers(c, authClient) })
	app.Get("/users/:id", auth, func(c *fiber.Ctx) error { return handlers.GetUserById(c, authClient) })
	app.Patch("/users", auth, func(c *fiber.Ctx) error { return handlers.UpdateUser(c, authClient) })
	app.Delete("/users/:id", auth, func(c *fiber.Ctx) error { return handlers.DeleteUser(c, authClient) })

	app.Get("/warehouse/products", func(c *fiber.Ctx) error { return handlers.GetAllProducts(c, warehouseClient) })
	app.Get("/warehouse/product/:id", func(c *fiber.Ctx) error { return handlers.GetProductById(c, warehouseClient) })
	app.Get("/warehouse/product/:id/inventory", func(c *fiber.Ctx) error { return handlers.CheckInventory(c, warehouseClient) })
	app.Post("/warehouse/product", auth, func(c *fiber.Ctx) error { return handlers.CreateProduct(c, warehouseClient) })
	app.Patch("/warehouse/product", auth, func(c *fiber.Ctx) error { return handlers.UpdateProduct(c, warehouseClient) })
	app.Delete("/warehouse/product/:id", auth, func(c *fiber.Ctx) error { return handlers.DeleteProduct(c, warehouseClient) })
	app.Patch("/warehouse/product/inventory", auth, func(c *fiber.Ctx) error { return handlers.UpdateInventory(c, warehouseClient) })

	app.Get("/sales", func(c *fiber.Ctx) error { return handlers.GetAllSales(c, salesClient) })
	app.Get("/sales/:id", func(c *fiber.Ctx) error { return handlers.GetSalesById(c, salesClient) })
	app.Post("/sales", auth, func(c *fiber.Ctx) error { return handlers.CreateSales(c, salesClient) })
	app.Patch("/sales", auth, func(c *fiber.Ctx) error { return handlers.UpdateSales(c, salesClient) })
	app.Delete("/sales/:id", auth, func(c *fiber.Ctx) error { return handlers.DeleteSales(c, salesClient) })

	app.Get("/orders", auth, func(c *fiber.Ctx) error { return handlers.GetAllOrders(c, orderClient) })
	app.Get("/orders/:id", auth, func(c *fiber.Ctx) error { return handlers.GetOrderById(c, orderClient) })
	app.Post("/orders", auth, func(c *fiber.Ctx) error { return handlers.CreateOrder(c, orderClient) })
	app.Patch("/orders", auth, func(c *fiber.Ctx) error { return handlers.UpdateOrder(c, orderClient) })
	app.Delete("/orders/:id", auth, func(c *fiber.Ctx) error { return handlers.DeleteOrder(c, orderClient) })

	go func() {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
		<-stop
		app.Shutdown()
	}()

	if err := app.Listen(":" + getenv("PORT", "3000")); err != nil {
		log.Fatal(err)
	}
}
