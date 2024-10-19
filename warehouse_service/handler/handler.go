package handler

import (
	"context"
	"log"
	"warehouse_service/model"
	pb "warehouse_service/protoc/warehouse"

	"gorm.io/gorm"
)

type WarehouseService struct {
	DB *gorm.DB
	pb.UnimplementedWarehouseServer
}

func (s *WarehouseService) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	product := model.Product{
		Name:     req.Name,
		Desc:     req.Desc,
		Price:    uint(req.Price),
		Quantity: uint(req.Quantity),
	}
	if err := s.DB.Create(&product).Error; err != nil {
		log.Printf("Error creating product: %v", err)
		return &pb.CreateProductResponse{Msg: "Error creating product", Status: 500}, nil
	}
	return &pb.CreateProductResponse{Msg: "Product successfully created", Status: 200}, nil
}

func (s *WarehouseService) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.UpdateProductResponse, error) {
	var product model.Product
	if err := s.DB.First(&product, req.Id).Error; err != nil {
		log.Printf("Product not found: %v", err)
		return &pb.UpdateProductResponse{Msg: "Product not found", Status: 404}, nil
	}

	product.Name = req.Name
	product.Desc = req.Desc
	product.Price = uint(req.Price)
	product.Quantity = uint(req.Quantity)

	if err := s.DB.Save(&product).Error; err != nil {
		log.Printf("Error updating product: %v", err)
		return &pb.UpdateProductResponse{Msg: "Error updating product", Status: 500}, nil
	}
	return &pb.UpdateProductResponse{Msg: "Product successfully updated", Status: 200}, nil
}

func (s *WarehouseService) GetAllProducts(ctx context.Context, req *pb.GetAllProductsRequest) (*pb.GetAllProductsResponse, error) {
	var products []model.Product
	if err := s.DB.Find(&products).Error; err != nil {
		log.Printf("Error retrieving products: %v", err)
		return nil, err
	}

	var pbProducts []*pb.Product
	for _, product := range products {
		pbProducts = append(pbProducts, &pb.Product{
			Id:       uint32(product.ID),
			Name:     product.Name,
			Desc:     product.Desc,
			Price:    uint32(product.Price),
			Quantity: uint32(product.Quantity),
		})
	}

	return &pb.GetAllProductsResponse{Products: pbProducts}, nil
}

func (s *WarehouseService) GetProductById(ctx context.Context, req *pb.GetProductByIdRequest) (*pb.GetProductByIdResponse, error) {
	var product model.Product
	if err := s.DB.First(&product, req.ProductId).Error; err != nil {
		log.Printf("Product not found: %v", err)
		return &pb.GetProductByIdResponse{Msg: "Product not found", Status: 404}, nil
	}

	return &pb.GetProductByIdResponse{
		Msg:    "Product found",
		Status: 200,
		Product: &pb.Product{
			Id:       uint32(product.ID),
			Name:     product.Name,
			Desc:     product.Desc,
			Price:    uint32(product.Price),
			Quantity: uint32(product.Quantity),
		},
	}, nil
}

func (s *WarehouseService) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error) {
	if err := s.DB.Delete(&model.Product{}, req.ProductId).Error; err != nil {
		log.Printf("Error deleting product: %v", err)
		return &pb.DeleteProductResponse{Msg: "Error deleting product", Status: 500}, nil
	}
	return &pb.DeleteProductResponse{Msg: "Product successfully deleted", Status: 200}, nil
}

func (s *WarehouseService) CheckInventory(ctx context.Context, req *pb.CheckInventoryRequest) (*pb.CheckInventoryResponse, error) {
	var product model.Product
	if err := s.DB.First(&product, req.ProductId).Error; err != nil {
		log.Printf("Product not found: %v", err)
		return nil, err
	}

	return &pb.CheckInventoryResponse{
		Product: &pb.Product{
			Id:       uint32(product.ID),
			Name:     product.Name,
			Desc:     product.Desc,
			Price:    uint32(product.Price),
			Quantity: uint32(product.Quantity),
		},
	}, nil
}

func (s *WarehouseService) UpdateInventory(ctx context.Context, req *pb.UpdateInventoryRequest) (*pb.UpdateInventoryResponse, error) {
	var product model.Product
	if err := s.DB.First(&product, req.ProductId).Error; err != nil {
		log.Printf("Product not found: %v", err)
		return &pb.UpdateInventoryResponse{Msg: "Product not found", Status: 404}, nil
	}

	if req.Move {
		product.Quantity += uint(req.Quantity)
	} else {
		if product.Quantity < uint(req.Quantity) {
			return &pb.UpdateInventoryResponse{Msg: "Not enough quantity in stock", Status: 400}, nil
		}
		product.Quantity -= uint(req.Quantity)
	}

	if err := s.DB.Save(&product).Error; err != nil {
		log.Printf("Error updating inventory: %v", err)
		return &pb.UpdateInventoryResponse{Msg: "Error updating inventory", Status: 500}, nil
	}

	return &pb.UpdateInventoryResponse{Msg: "Inventory successfully updated", Status: 200}, nil
}
