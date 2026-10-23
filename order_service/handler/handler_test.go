package handler

import (
	"context"
	"path/filepath"
	"sync"
	"testing"

	"order_service/database"
	"order_service/model"
	ab "order_service/protoc/auth"
	pb "order_service/protoc/order"
	wb "order_service/protoc/warehouse"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type fakeAuth struct {
	ab.AuthServiceClient
	users map[uint32]bool
}

func (f *fakeAuth) GetUserById(_ context.Context, in *ab.GetUserByIdRequest, _ ...grpc.CallOption) (*ab.GetUserByIdResponse, error) {
	if !f.users[in.UserId] {
		return nil, status.Error(codes.NotFound, "user not found")
	}
	return &ab.GetUserByIdResponse{Status: 200, User: &ab.User{Id: in.UserId}}, nil
}

type fakeWarehouse struct {
	wb.WarehouseClient
	mu    sync.Mutex
	stock map[uint32]uint32
}

func (f *fakeWarehouse) UpdateInventory(_ context.Context, in *wb.UpdateInventoryRequest, _ ...grpc.CallOption) (*wb.UpdateInventoryResponse, error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	qty, ok := f.stock[in.ProductId]
	if !ok {
		return &wb.UpdateInventoryResponse{Msg: "Product not found", Status: 404}, nil
	}
	if in.Mov {
		f.stock[in.ProductId] = qty + in.Quantity
		return &wb.UpdateInventoryResponse{Status: 200}, nil
	}
	if qty < in.Quantity {
		return &wb.UpdateInventoryResponse{Msg: "Not enough quantity in stock", Status: 409}, nil
	}
	f.stock[in.ProductId] = qty - in.Quantity
	return &wb.UpdateInventoryResponse{Status: 200}, nil
}

func setup(t *testing.T) (*OrderService, *fakeWarehouse) {
	t.Helper()
	db := database.SetupDB(filepath.Join(t.TempDir(), "orders.db"))
	sqlDB, _ := db.DB()
	t.Cleanup(func() { sqlDB.Close() })

	wh := &fakeWarehouse{stock: map[uint32]uint32{1: 5}}
	return &OrderService{
		DB:      db,
		Wclient: wh,
		Aclient: &fakeAuth{users: map[uint32]bool{7: true}},
	}, wh
}

func countOrders(t *testing.T, s *OrderService) int64 {
	var n int64
	s.DB.Model(&model.Order{}).Count(&n)
	return n
}

func TestCreateOrder(t *testing.T) {
	s, wh := setup(t)
	ctx := context.Background()

	res, err := s.CreateOrder(ctx, &pb.CreateOrderRequest{UserId: 7, ProductId: 1, Quantity: 2})
	if err != nil || res.Status != 200 {
		t.Fatalf("create: %v %+v", err, res)
	}
	if wh.stock[1] != 3 || countOrders(t, s) != 1 {
		t.Fatalf("stock %d, orders %d", wh.stock[1], countOrders(t, s))
	}
}

func TestCreateOrderRejected(t *testing.T) {
	s, wh := setup(t)
	ctx := context.Background()

	if res, _ := s.CreateOrder(ctx, &pb.CreateOrderRequest{UserId: 99, ProductId: 1, Quantity: 1}); res.Status != 404 {
		t.Fatalf("unknown user: status %d", res.Status)
	}
	if res, _ := s.CreateOrder(ctx, &pb.CreateOrderRequest{UserId: 7, ProductId: 1, Quantity: 10}); res.Status != 409 {
		t.Fatalf("out of stock: status %d", res.Status)
	}
	if res, _ := s.CreateOrder(ctx, &pb.CreateOrderRequest{UserId: 7, ProductId: 1, Quantity: 0}); res.Status != 400 {
		t.Fatalf("zero quantity: status %d", res.Status)
	}
	if wh.stock[1] != 5 || countOrders(t, s) != 0 {
		t.Fatalf("rejected orders changed state: stock %d, orders %d", wh.stock[1], countOrders(t, s))
	}
}

func TestUpdateAndDeleteOrderMoveStock(t *testing.T) {
	s, wh := setup(t)
	ctx := context.Background()

	s.CreateOrder(ctx, &pb.CreateOrderRequest{UserId: 7, ProductId: 1, Quantity: 2})
	var order model.Order
	s.DB.First(&order)

	if res, _ := s.UpdateOrder(ctx, &pb.UpdateOrderRequest{Id: uint32(order.ID), Quantity: 4}); res.Status != 200 || wh.stock[1] != 1 {
		t.Fatalf("increase: status %d, stock %d", res.Status, wh.stock[1])
	}
	if res, _ := s.UpdateOrder(ctx, &pb.UpdateOrderRequest{Id: uint32(order.ID), Quantity: 1}); res.Status != 200 || wh.stock[1] != 4 {
		t.Fatalf("decrease: status %d, stock %d", res.Status, wh.stock[1])
	}
	if res, _ := s.DeleteOrder(ctx, &pb.DeleteOrderRequest{OrderId: uint32(order.ID)}); res.Status != 200 || wh.stock[1] != 5 {
		t.Fatalf("delete: status %d, stock %d", res.Status, wh.stock[1])
	}
}
