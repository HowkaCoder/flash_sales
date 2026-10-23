package grpc

import (
	"context"
	"fmt"
	wb "sales_service/protoc/warehouse"
	"time"
)

func ProductExists(client wb.WarehouseClient, productID uint32) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	res, err := client.GetProductById(ctx, &wb.GetProductByIdRequest{ProductId: productID})
	if err != nil {
		return false, fmt.Errorf("warehouse service: %w", err)
	}
	return res.Status == 200 && res.Product != nil && res.Product.Id != 0, nil
}
