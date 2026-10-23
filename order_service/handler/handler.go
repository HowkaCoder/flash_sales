package handler

import (
	"context"
	"fmt"
	"log"
	"order_service/grpc"
	"order_service/kafka"
	"order_service/model"
	ab "order_service/protoc/auth"
	pb "order_service/protoc/order"
	wb "order_service/protoc/warehouse"

	"gorm.io/gorm"
)

type OrderService struct {
	DB        *gorm.DB
	Wclient   wb.WarehouseClient
	Aclient   ab.AuthServiceClient
	Publisher *kafka.Publisher
	pb.UnimplementedOrderServiceServer
}

func (s *OrderService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	if req.Quantity == 0 {
		return &pb.CreateOrderResponse{Msg: "Quantity must be greater than zero", Status: 400}, nil
	}

	ok, err := grpc.UserExists(s.Aclient, req.UserId)
	if err != nil {
		log.Printf("check user %d: %v", req.UserId, err)
		return &pb.CreateOrderResponse{Msg: "Auth service is unavailable", Status: 503}, nil
	}
	if !ok {
		return &pb.CreateOrderResponse{Msg: "User not found", Status: 404}, nil
	}

	code, msg, err := grpc.ChangeInventory(s.Wclient, req.ProductId, req.Quantity, false)
	if err != nil {
		log.Printf("reserve product %d: %v", req.ProductId, err)
		return &pb.CreateOrderResponse{Msg: "Warehouse service is unavailable", Status: 503}, nil
	}
	if code != 200 {
		return &pb.CreateOrderResponse{Msg: msg, Status: code}, nil
	}

	order := model.Order{
		UserID:    uint(req.UserId),
		ProductID: uint(req.ProductId),
		Quantity:  uint(req.Quantity),
		Status:    req.Status,
	}
	if order.Status == "" {
		order.Status = "created"
	}

	if err := s.DB.WithContext(ctx).Create(&order).Error; err != nil {
		log.Printf("create order: %v", err)
		if _, _, rerr := grpc.ChangeInventory(s.Wclient, req.ProductId, req.Quantity, true); rerr != nil {
			log.Printf("return %d of product %d to stock: %v", req.Quantity, req.ProductId, rerr)
		}
		return &pb.CreateOrderResponse{Msg: "Error creating order", Status: 500}, nil
	}

	s.Publisher.Notify("order_created", fmt.Sprintf("order %d: user %d, product %d x %d", order.ID, order.UserID, order.ProductID, order.Quantity))
	return &pb.CreateOrderResponse{Msg: "Order successfully created", Status: 200}, nil
}

func (s *OrderService) UpdateOrder(ctx context.Context, req *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error) {
	var order model.Order
	if err := s.DB.WithContext(ctx).First(&order, req.Id).Error; err != nil {
		return &pb.UpdateOrderResponse{Msg: "Order not found", Status: 404}, nil
	}
	if req.ProductId != 0 && uint(req.ProductId) != order.ProductID {
		return &pb.UpdateOrderResponse{Msg: "Product of an order can't be changed", Status: 400}, nil
	}

	if req.Quantity != 0 && uint(req.Quantity) != order.Quantity {
		add := uint(req.Quantity) < order.Quantity
		var delta uint32
		if add {
			delta = uint32(order.Quantity) - req.Quantity
		} else {
			delta = req.Quantity - uint32(order.Quantity)
		}

		code, msg, err := grpc.ChangeInventory(s.Wclient, uint32(order.ProductID), delta, add)
		if err != nil {
			log.Printf("change inventory: %v", err)
			return &pb.UpdateOrderResponse{Msg: "Warehouse service is unavailable", Status: 503}, nil
		}
		if code != 200 {
			return &pb.UpdateOrderResponse{Msg: msg, Status: code}, nil
		}
		order.Quantity = uint(req.Quantity)
	}
	if req.Status != "" {
		order.Status = req.Status
	}

	if err := s.DB.WithContext(ctx).Save(&order).Error; err != nil {
		log.Printf("update order: %v", err)
		return &pb.UpdateOrderResponse{Msg: "Error updating order", Status: 500}, nil
	}
	return &pb.UpdateOrderResponse{Msg: "Order successfully updated", Status: 200}, nil
}

func (s *OrderService) GetAllOrders(ctx context.Context, req *pb.GetAllOrdersRequest) (*pb.GetAllOrdersResponse, error) {
	var orders []model.Order
	if err := s.DB.Find(&orders).Error; err != nil {
		log.Printf("Error retrieving orders: %v", err)
		return nil, err
	}

	var pbOrders []*pb.Order
	for _, order := range orders {
		pbOrders = append(pbOrders, &pb.Order{
			Id:        uint32(order.ID),
			UserId:    uint32(order.UserID),
			ProductId: uint32(order.ProductID),
			Quantity:  uint32(order.Quantity),
			Status:    order.Status,
		})
	}

	return &pb.GetAllOrdersResponse{Orders: pbOrders}, nil
}

func (s *OrderService) GetOrderById(ctx context.Context, req *pb.GetOrderByIdRequest) (*pb.GetOrderByIdResponse, error) {
	var order model.Order
	if err := s.DB.First(&order, req.OrderId).Error; err != nil {
		log.Printf("Order not found: %v", err)
		return &pb.GetOrderByIdResponse{Msg: "Order not found", Status: 404}, nil
	}

	return &pb.GetOrderByIdResponse{
		Msg:    "Order found",
		Status: 200,
		Order: &pb.Order{
			Id:        uint32(order.ID),
			UserId:    uint32(order.UserID),
			ProductId: uint32(order.ProductID),
			Quantity:  uint32(order.Quantity),
			Status:    order.Status,
		},
	}, nil
}

func (s *OrderService) DeleteOrder(ctx context.Context, req *pb.DeleteOrderRequest) (*pb.DeleteOrderResponse, error) {
	var order model.Order
	if err := s.DB.WithContext(ctx).First(&order, req.OrderId).Error; err != nil {
		return &pb.DeleteOrderResponse{Msg: "Order not found", Status: 404}, nil
	}
	if err := s.DB.WithContext(ctx).Delete(&order).Error; err != nil {
		log.Printf("Error deleting order: %v", err)
		return &pb.DeleteOrderResponse{Msg: "Error deleting order", Status: 500}, nil
	}

	if _, _, err := grpc.ChangeInventory(s.Wclient, uint32(order.ProductID), uint32(order.Quantity), true); err != nil {
		log.Printf("return product %d to stock: %v", order.ProductID, err)
	}
	return &pb.DeleteOrderResponse{Msg: "Order successfully deleted", Status: 200}, nil
}
