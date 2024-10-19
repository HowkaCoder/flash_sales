package handler

import (
	"context"
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
	DB      *gorm.DB
	Wclient wb.WarehouseClient
	Aclient ab.AuthServiceClient
	pb.UnimplementedOrderServiceServer
}

func (s *OrderService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	order := model.Order{
		UserID:    uint(req.UserId),
		ProductID: uint(req.ProductId),
		Quantity:  uint(req.Quantity),
		Status:    req.Status,
	}

	res := grpc.CheckInventory(s.Wclient, req.ProductId)

	a_res := grpc.GetUserById(s.Aclient, req.UserId)

	if a_res.User.Id == 0 {
		return &pb.CreateOrderResponse{Msg: "User not found!", Status: 404}, nil
	}

	if res.Product.Quantity != 0 {
		if res.Product.Quantity > req.Quantity || res.Product.Quantity == req.Quantity {
			if err := s.DB.Create(&order).Error; err != nil {
				log.Printf("Error creating order: %v", err)
				return &pb.CreateOrderResponse{Msg: "Error creating order", Status: 500}, nil
			}
			msg, status := grpc.UpdateInventory(s.Wclient, req.ProductId, req.Quantity, false)
			log.Printf("Message from WareHouse : %v and status code too : %v", msg, status)

		} else {
			return &pb.CreateOrderResponse{Msg: "There are no products for you bro", Status: 404}, nil
		}
	} else {
		return &pb.CreateOrderResponse{Msg: "There are no products for you bro", Status: 404}, nil
	}
	kafka.NotifyKafka("products", "CreateOrder - Reuqest", "New Product in Warehouse")
	return &pb.CreateOrderResponse{Msg: "Order successfully created", Status: 200}, nil
}

func (s *OrderService) UpdateOrder(ctx context.Context, req *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error) {
	var order model.Order
	if err := s.DB.First(&order, req.Id).Error; err != nil {
		log.Printf("Order not found: %v", err)
		return &pb.UpdateOrderResponse{Msg: "Order not found", Status: 404}, nil
	}
	var lop bool
	if req.Quantity > uint32(order.Quantity) {
		lop = true
	} else if req.Quantity < uint32(order.Quantity) {
		lop = false
	}

	order.UserID = uint(req.UserId)
	order.ProductID = uint(req.ProductId)
	order.Quantity = uint(req.Quantity)
	order.Status = req.Status

	a_res := grpc.GetUserById(s.Aclient, req.UserId)
	if a_res.User.Id == 0 {
		return &pb.UpdateOrderResponse{Msg: "Error , User not found", Status: 500}, nil
	}
	if err := s.DB.Save(&order).Error; err != nil {
		log.Printf("Error updating order: %v", err)
		return &pb.UpdateOrderResponse{Msg: "Error updating order", Status: 500}, nil
	}

	res := grpc.CheckInventory(s.Wclient, req.ProductId)

	if res.Product.Quantity != 0 {
		if res.Product.Quantity > req.Quantity || res.Product.Quantity == req.Quantity {
			if err := s.DB.Save(&order).Error; err != nil {
				log.Printf("Error creating order: %v", err)
				return &pb.UpdateOrderResponse{Msg: "Error creating order", Status: 500}, nil
			}
			if lop {
				msg, status := grpc.UpdateInventory(s.Wclient, req.ProductId, req.Quantity-uint32(order.Quantity), true)
				log.Printf("Message from WareHouse : %v and status code too : %v", msg, status)
			} else if !lop {
				msg, status := grpc.UpdateInventory(s.Wclient, req.ProductId, uint32(order.Quantity)-req.Quantity, false)
				log.Printf("Message from WareHouse : %v and Status code too : %v ", msg, status)
			}

		} else {
			return &pb.UpdateOrderResponse{Msg: "There are no products for you bro", Status: 404}, nil
		}
	} else {
		return &pb.UpdateOrderResponse{Msg: "There are no products for you bro", Status: 404}, nil
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
	if err := s.DB.Delete(&model.Order{}, req.OrderId).Error; err != nil {
		log.Printf("Error deleting order: %v", err)
		return &pb.DeleteOrderResponse{Msg: "Error deleting order", Status: 500}, nil
	}
	return &pb.DeleteOrderResponse{Msg: "Order successfully deleted", Status: 200}, nil
}
