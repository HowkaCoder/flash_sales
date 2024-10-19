package grpc

import (
	"context"
	salespb "gateway/protoc/sales"
	"log"
	"time"
)

func CreateSales(client salespb.SalesClient, productID, discountPrice uint32, startTime, endTime string) (string, uint32) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &salespb.CreateSalesRequest{
		ProductId:     productID,
		DiscountPrice: discountPrice,
		StartTime:     startTime,
		EndTime:       endTime,
	}

	res, err := client.CreateSales(ctx, req)
	if err != nil {
		log.Fatalf("Error with CreateSales request: %v", err)
	}

	log.Printf("Response Message: %v and Response code : %v", res.Msg, res.Status)
	return res.Msg, res.Status
}

func GetAllSales(client salespb.SalesClient) ([]*salespb.Sale, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &salespb.GetAllSalesRequest{}

	res, err := client.GetAllSales(ctx, req)
	if err != nil {
		log.Fatalf("Error with GetAllSales request: %v", err)
		return nil, err
	}

	log.Printf("Received %d sales", len(res.Sales))
	return res.Sales, nil
}

func GetSalesById(client salespb.SalesClient, salesID uint32) (*salespb.Sale, string, uint32) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &salespb.GetSalesByIdRequest{
		SalesId: salesID,
	}

	res, err := client.GetSalesById(ctx, req)
	if err != nil {
		log.Fatalf("Error with GetSalesById request: %v", err)
	}

	log.Printf("Response Message: %v and Response code : %v", res.Msg, res.Status)
	return res.Sale, res.Msg, res.Status
}

func UpdateSales(client salespb.SalesClient, id, productID, discountPrice uint32, startTime, endTime string) (string, uint32) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &salespb.UpdateSalesRequest{
		Id:            id,
		ProductId:     productID,
		DiscountPrice: discountPrice,
		StartTime:     startTime,
		EndTime:       endTime,
	}

	res, err := client.UpdateSales(ctx, req)
	if err != nil {
		log.Fatalf("Error with UpdateSales request: %v", err)
	}

	log.Printf("Response Message: %v and Response code : %v", res.Msg, res.Status)
	return res.Msg, res.Status
}

func DeleteSales(client salespb.SalesClient, salesID uint32) (string, uint32) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &salespb.DeleteSalesRequest{
		SalesId: salesID,
	}

	res, err := client.DeleteSales(ctx, req)
	if err != nil {
		log.Fatalf("Error with DeleteSales request: %v", err)
	}

	log.Printf("Response Message: %v and Response code : %v", res.Msg, res.Status)
	return res.Msg, res.Status
}
