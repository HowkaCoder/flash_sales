package grpc

import (
	"context"
	wb "gateway/protoc/warehouse"
	"log"
	"time"
)

func CreateProduct(client wb.WarehouseClient, name, desc string, price, quantity uint32) (string, uint32) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &wb.CreateProductRequest{
		Name:     name,
		Desc:     desc,
		Price:    price,
		Quantity: quantity,
	}

	res, err := client.CreateProduct(ctx, req)
	if err != nil {
		log.Fatalf("Error with CreateProduct request: %v", err)
	}

	log.Printf("Response Message: %v and Response code : %v", res.Msg, res.Status)
	return res.Msg, res.Status
}

func UpdateProduct(client wb.WarehouseClient, id uint32, name, desc string, price, quantity uint32) (string, uint32) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &wb.UpdateProductRequest{
		Id:       id,
		Name:     name,
		Desc:     desc,
		Price:    price,
		Quantity: quantity,
	}

	res, err := client.UpdateProduct(ctx, req)
	if err != nil {
		log.Fatalf("Error with UpdateProduct request: %v", err)
	}

	log.Printf("Response Message: %v and Response code : %v", res.Msg, res.Status)
	return res.Msg, res.Status
}

func GetAllProducts(client wb.WarehouseClient) ([]*wb.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &wb.GetAllProductsRequest{}

	res, err := client.GetAllProducts(ctx, req)
	if err != nil {
		log.Fatalf("Error with GetAllProducts request: %v", err)
		return nil, err
	}

	log.Printf("Received %d products", len(res.Products))
	return res.Products, nil
}

func GetProductById(client wb.WarehouseClient, productId uint32) (*wb.Product, string, uint32) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &wb.GetProductByIdRequest{
		ProductId: productId,
	}

	res, err := client.GetProductById(ctx, req)
	if err != nil {
		log.Fatalf("Error with GetProductById request: %v", err)
	}

	log.Printf("Response Message: %v and Response code : %v", res.Msg, res.Status)
	return res.Product, res.Msg, res.Status
}

func DeleteProduct(client wb.WarehouseClient, productId uint32) (string, uint32) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &wb.DeleteProductRequest{
		ProductId: productId,
	}

	res, err := client.DeleteProduct(ctx, req)
	if err != nil {
		log.Fatalf("Error with DeleteProduct request: %v", err)
	}

	log.Printf("Response Message: %v and Response code : %v", res.Msg, res.Status)
	return res.Msg, res.Status
}

func CheckInventory(client wb.WarehouseClient, productId uint32) (*wb.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &wb.CheckInventoryRequest{
		ProductId: productId,
	}

	res, err := client.CheckInventory(ctx, req)
	if err != nil {
		log.Fatalf("Error with CheckInventory request: %v", err)
		return nil, err
	}

	log.Printf("Product ID: %v, Name: %v, Quantity: %v", res.Product.Id, res.Product.Name, res.Product.Quantity)
	return res.Product, nil
}

func UpdateInventory(client wb.WarehouseClient, productId, quantity uint32, move bool) (string, uint32) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &wb.UpdateInventoryRequest{
		ProductId: productId,
		Quantity:  quantity,
		Move:      move,
	}

	res, err := client.UpdateInventory(ctx, req)
	if err != nil {
		log.Fatalf("Error with UpdateInventory request: %v", err)
	}

	log.Printf("Response Message: %v and Response code : %v", res.Msg, res.Status)
	return res.Msg, res.Status
}
