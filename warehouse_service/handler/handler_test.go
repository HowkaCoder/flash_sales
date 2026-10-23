package handler

import (
	"context"
	"path/filepath"
	"sync"
	"testing"

	"warehouse_service/database"
	"warehouse_service/model"
	pb "warehouse_service/protoc/warehouse"
)

func newService(t *testing.T) *WarehouseService {
	t.Helper()
	db := database.SetupDB(filepath.Join(t.TempDir(), "test.db"))
	sqlDB, err := db.DB()
	if err != nil {
		t.Fatal(err)
	}
	sqlDB.SetMaxOpenConns(1)
	t.Cleanup(func() { sqlDB.Close() })
	return &WarehouseService{DB: db}
}

func TestNoOversellUnderLoad(t *testing.T) {
	s := newService(t)
	product := model.Product{Name: "Phone", Desc: "flash sale", Price: 100, Quantity: 10}
	if err := s.DB.Create(&product).Error; err != nil {
		t.Fatal(err)
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	sold := 0
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			res, err := s.UpdateInventory(context.Background(), &pb.UpdateInventoryRequest{ProductId: uint32(product.ID), Quantity: 1})
			if err != nil {
				t.Error(err)
				return
			}
			if res.Status == 200 {
				mu.Lock()
				sold++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()

	if sold != 10 {
		t.Fatalf("sold %d items, stock was 10", sold)
	}
	var got model.Product
	s.DB.First(&got, product.ID)
	if got.Quantity != 0 {
		t.Fatalf("quantity left: %d", got.Quantity)
	}
}

func TestUpdateInventory(t *testing.T) {
	s := newService(t)
	product := model.Product{Name: "Laptop", Desc: "x", Price: 500, Quantity: 2}
	s.DB.Create(&product)
	ctx := context.Background()

	res, _ := s.UpdateInventory(ctx, &pb.UpdateInventoryRequest{ProductId: uint32(product.ID), Quantity: 3})
	if res.Status != 409 {
		t.Fatalf("taking more than in stock: status %d", res.Status)
	}

	res, _ = s.UpdateInventory(ctx, &pb.UpdateInventoryRequest{ProductId: 999, Quantity: 1})
	if res.Status != 404 {
		t.Fatalf("unknown product: status %d", res.Status)
	}

	res, _ = s.UpdateInventory(ctx, &pb.UpdateInventoryRequest{ProductId: uint32(product.ID), Quantity: 5, Move: true})
	if res.Status != 200 {
		t.Fatalf("restock: status %d", res.Status)
	}

	inv, err := s.CheckInventory(ctx, &pb.CheckInventoryRequest{ProductId: uint32(product.ID)})
	if err != nil {
		t.Fatal(err)
	}
	if inv.Product.Quantity != 7 {
		t.Fatalf("quantity = %d, want 7", inv.Product.Quantity)
	}
}
