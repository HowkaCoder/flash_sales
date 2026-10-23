package handler

import (
	"context"
	"errors"
	"log"
	"warehouse_service/model"
	pb "warehouse_service/protoc/warehouse"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	if err := s.DB.WithContext(ctx).First(&product, req.ProductId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, "product not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.CheckInventoryResponse{Product: toPB(product)}, nil
}

func (s *WarehouseService) UpdateInventory(ctx context.Context, req *pb.UpdateInventoryRequest) (*pb.UpdateInventoryResponse, error) {
	if req.Quantity == 0 {
		return &pb.UpdateInventoryResponse{Msg: "Quantity must be greater than zero", Status: 400}, nil
	}

	db := s.DB.WithContext(ctx).Model(&model.Product{}).Where("id = ?", req.ProductId)
	var res *gorm.DB
	if req.Move {
		res = db.UpdateColumn("quantity", gorm.Expr("quantity + ?", req.Quantity))
	} else {
		res = db.Where("quantity >= ?", req.Quantity).UpdateColumn("quantity", gorm.Expr("quantity - ?", req.Quantity))
	}
	if res.Error != nil {
		log.Printf("Error updating inventory: %v", res.Error)
		return &pb.UpdateInventoryResponse{Msg: "Error updating inventory", Status: 500}, nil
	}

	if res.RowsAffected == 0 {
		var count int64
		s.DB.WithContext(ctx).Model(&model.Product{}).Where("id = ?", req.ProductId).Count(&count)
		if count == 0 {
			return &pb.UpdateInventoryResponse{Msg: "Product not found", Status: 404}, nil
		}
		return &pb.UpdateInventoryResponse{Msg: "Not enough quantity in stock", Status: 409}, nil
	}

	return &pb.UpdateInventoryResponse{Msg: "Inventory successfully updated", Status: 200}, nil
}

func toPB(p model.Product) *pb.Product {
	return &pb.Product{
		Id:       uint32(p.ID),
		Name:     p.Name,
		Desc:     p.Desc,
		Price:    uint32(p.Price),
		Quantity: uint32(p.Quantity),
	}
}
