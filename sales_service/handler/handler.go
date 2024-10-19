package handler

import (
	"context"
	"fmt"
	"log"
	"sales_service/grpc"
	"sales_service/kafka"
	"sales_service/model"
	pb "sales_service/protoc/sales"
	wb "sales_service/protoc/warehouse"
	"time"

	"gorm.io/gorm"
)

type SaleService struct {
	DB *gorm.DB

	Wclient wb.WarehouseClient
	pb.UnimplementedSalesServer
}

func (s *SaleService) DeleteSales(ctx context.Context, req *pb.DeleteSalesRequest) (*pb.DeleteSalesResponse, error) {
	if err := s.DB.Delete(&model.Sale{}, uint(req.SalesId)).Error; err != nil {
		return &pb.DeleteSalesResponse{
			Msg:    "Error with removing sale",
			Status: 500,
		}, nil
	}
	return &pb.DeleteSalesResponse{
		Msg:    "Succesfully Deleted",
		Status: 500,
	}, nil
}

func (s *SaleService) GetSalesById(ctx context.Context, req *pb.GetSalesByIdRequest) (*pb.GetSalesByIdResponse, error) {
	var sale model.Sale
	if err := s.DB.First(&sale, uint(req.SalesId)).Error; err != nil {
		fmt.Sprintf("Error with getting sale by id: %v", err)
		return &pb.GetSalesByIdResponse{
			Msg:    "Sales not found",
			Status: 404,
		}, nil
	}

	return &pb.GetSalesByIdResponse{
		Msg:    "Sale Found",
		Status: 200,
		Sale: &pb.Sale{
			Id:            uint32(sale.ID),
			ProductId:     uint32(sale.ProductID),
			DiscountPrice: uint32(sale.DiscountPrice),
			StartTime:     sale.StartTime.Format(time.RFC3339),
			EndTime:       sale.EndTime.Format(time.RFC3339),
		},
	}, nil
}

func (s *SaleService) GetAllSales(ctx context.Context, req *pb.GetAllSalesRequest) (*pb.GetAllSalesResponse, error) {

	var sales []model.Sale
	if err := s.DB.Find(&sales).Error; err != nil {
		log.Fatalf("Error with getting all sales: %v", err)
		return nil, nil
	}

	var pbSales []*pb.Sale
	for _, sale := range sales {
		pbSales = append(pbSales, &pb.Sale{
			Id:            uint32(sale.ID),
			ProductId:     uint32(sale.ProductID),
			DiscountPrice: uint32(sale.DiscountPrice),
			StartTime:     sale.StartTime.Format(time.RFC3339),
			EndTime:       sale.EndTime.Format(time.RFC3339),
		})
	}
	return &pb.GetAllSalesResponse{Sales: pbSales}, nil

}

func (s *SaleService) UpdateSales(ctx context.Context, req *pb.UpdateSalesRequest) (*pb.UpdateSalesResponse, error) {

	var eSale model.Sale
	if err := s.DB.First(&eSale, uint(req.Id)).Error; err != nil {
		return &pb.UpdateSalesResponse{
			Msg:    "There is no sales with this ID",
			Status: 404,
		}, nil
	}
	if req.DiscountPrice != 0 {
		eSale.DiscountPrice = uint(req.DiscountPrice)
	}
	if req.ProductId != 0 {
		res := grpc.GetProductById(s.Wclient, req.ProductId)

		if res.Product.Id == 0 {
			return &pb.UpdateSalesResponse{
				Msg:    "Product Not Found",
				Status: 404,
			}, nil
		}

		eSale.ProductID = uint(req.ProductId)
	}
	if req.StartTime != "" {
		start_time, err := time.Parse(time.RFC3339, req.StartTime)
		if err != nil {
			return &pb.UpdateSalesResponse{
				Msg:    "Error with start time parsing",
				Status: 500,
			}, nil
		}
		eSale.StartTime = start_time
	}

	if req.EndTime != "" {
		end_time, err := time.Parse(time.RFC3339, req.EndTime)
		if err != nil {
			return &pb.UpdateSalesResponse{
				Msg:    "Error with End time Parsing",
				Status: 500,
			}, nil
		}
		eSale.EndTime = end_time
	}

	if err := s.DB.Save(&eSale).Error; err != nil {
		return &pb.UpdateSalesResponse{
			Msg:    "Error with new sales saving",
			Status: 500,
		}, nil
	}

	return &pb.UpdateSalesResponse{
		Msg:    "Succesfully updated",
		Status: 200,
	}, nil

}

func (s *SaleService) CreateSales(ctx context.Context, req *pb.CreateSalesRequest) (*pb.CreateSalesResponse, error) {

	start_time, err := time.Parse(time.RFC3339, req.StartTime)
	if err != nil {
		return &pb.CreateSalesResponse{Msg: "Error with start_time parsing", Status: 500}, nil
	}

	end_time, err := time.Parse(time.RFC3339, req.EndTime)
	if err != nil {
		return &pb.CreateSalesResponse{Msg: "Error with end_time parsing", Status: 500}, nil
	}

	if !end_time.After(start_time) {
		return &pb.CreateSalesResponse{Msg: "end time must be after the start time", Status: 500}, nil
	}
	res := grpc.GetProductById(s.Wclient, req.ProductId)

	if res.Product.Id == 0 {
		return &pb.CreateSalesResponse{
			Msg:    "Product not found!",
			Status: 404,
		}, nil
	}

	sale := model.Sale{
		ProductID:     uint(req.ProductId),
		DiscountPrice: uint(req.DiscountPrice),
		StartTime:     start_time,
		EndTime:       end_time,
	}

	if err := s.DB.Create(&sale).Error; err != nil {
		return &pb.CreateSalesResponse{Msg: "Error with sales creating in database", Status: 500}, err
	}

	kafka.NotifyKafka("sales", "CreateSale - Request", "New Sales Guys!!!!!")
	return &pb.CreateSalesResponse{
		Msg:    "Succesfully created sales",
		Status: 200,
	}, nil

}
