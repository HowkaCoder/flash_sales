package grpc

import (
	"context"
	"log"
	wb "sales_service/protoc/warehouse"
	"time"
)

func GetProductById(client wb.WarehouseClient, product_id uint32) *wb.GetProductByIdResponse {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &wb.GetProductByIdRequest{
		ProductId: product_id,
	}

	res, err := client.GetProductById(ctx, req)
	if err != nil {
		log.Fatalf("Error with Get product by ID request")
	}

	log.Printf("Msg : %v   Status Code : %v and Product : %v ", res.Msg, res.Status, res.Product)
	return res
}
