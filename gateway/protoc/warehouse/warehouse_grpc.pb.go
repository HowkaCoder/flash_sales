// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.1
// source: warehouse.proto

package warehouse

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Warehouse_CreateProduct_FullMethodName   = "/warehouse.Warehouse/CreateProduct"
	Warehouse_UpdateProduct_FullMethodName   = "/warehouse.Warehouse/UpdateProduct"
	Warehouse_GetAllProducts_FullMethodName  = "/warehouse.Warehouse/GetAllProducts"
	Warehouse_GetProductById_FullMethodName  = "/warehouse.Warehouse/GetProductById"
	Warehouse_DeleteProduct_FullMethodName   = "/warehouse.Warehouse/DeleteProduct"
	Warehouse_CheckInventory_FullMethodName  = "/warehouse.Warehouse/CheckInventory"
	Warehouse_UpdateInventory_FullMethodName = "/warehouse.Warehouse/UpdateInventory"
)

// WarehouseClient is the client API for Warehouse service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WarehouseClient interface {
	CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*CreateProductResponse, error)
	UpdateProduct(ctx context.Context, in *UpdateProductRequest, opts ...grpc.CallOption) (*UpdateProductResponse, error)
	GetAllProducts(ctx context.Context, in *GetAllProductsRequest, opts ...grpc.CallOption) (*GetAllProductsResponse, error)
	GetProductById(ctx context.Context, in *GetProductByIdRequest, opts ...grpc.CallOption) (*GetProductByIdResponse, error)
	DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*DeleteProductResponse, error)
	CheckInventory(ctx context.Context, in *CheckInventoryRequest, opts ...grpc.CallOption) (*CheckInventoryResponse, error)
	UpdateInventory(ctx context.Context, in *UpdateInventoryRequest, opts ...grpc.CallOption) (*UpdateInventoryResponse, error)
}

type warehouseClient struct {
	cc grpc.ClientConnInterface
}

func NewWarehouseClient(cc grpc.ClientConnInterface) WarehouseClient {
	return &warehouseClient{cc}
}

func (c *warehouseClient) CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*CreateProductResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateProductResponse)
	err := c.cc.Invoke(ctx, Warehouse_CreateProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warehouseClient) UpdateProduct(ctx context.Context, in *UpdateProductRequest, opts ...grpc.CallOption) (*UpdateProductResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateProductResponse)
	err := c.cc.Invoke(ctx, Warehouse_UpdateProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warehouseClient) GetAllProducts(ctx context.Context, in *GetAllProductsRequest, opts ...grpc.CallOption) (*GetAllProductsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllProductsResponse)
	err := c.cc.Invoke(ctx, Warehouse_GetAllProducts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warehouseClient) GetProductById(ctx context.Context, in *GetProductByIdRequest, opts ...grpc.CallOption) (*GetProductByIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetProductByIdResponse)
	err := c.cc.Invoke(ctx, Warehouse_GetProductById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warehouseClient) DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*DeleteProductResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteProductResponse)
	err := c.cc.Invoke(ctx, Warehouse_DeleteProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warehouseClient) CheckInventory(ctx context.Context, in *CheckInventoryRequest, opts ...grpc.CallOption) (*CheckInventoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckInventoryResponse)
	err := c.cc.Invoke(ctx, Warehouse_CheckInventory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warehouseClient) UpdateInventory(ctx context.Context, in *UpdateInventoryRequest, opts ...grpc.CallOption) (*UpdateInventoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateInventoryResponse)
	err := c.cc.Invoke(ctx, Warehouse_UpdateInventory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WarehouseServer is the server API for Warehouse service.
// All implementations must embed UnimplementedWarehouseServer
// for forward compatibility.
type WarehouseServer interface {
	CreateProduct(context.Context, *CreateProductRequest) (*CreateProductResponse, error)
	UpdateProduct(context.Context, *UpdateProductRequest) (*UpdateProductResponse, error)
	GetAllProducts(context.Context, *GetAllProductsRequest) (*GetAllProductsResponse, error)
	GetProductById(context.Context, *GetProductByIdRequest) (*GetProductByIdResponse, error)
	DeleteProduct(context.Context, *DeleteProductRequest) (*DeleteProductResponse, error)
	CheckInventory(context.Context, *CheckInventoryRequest) (*CheckInventoryResponse, error)
	UpdateInventory(context.Context, *UpdateInventoryRequest) (*UpdateInventoryResponse, error)
	mustEmbedUnimplementedWarehouseServer()
}

// UnimplementedWarehouseServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedWarehouseServer struct{}

func (UnimplementedWarehouseServer) CreateProduct(context.Context, *CreateProductRequest) (*CreateProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (UnimplementedWarehouseServer) UpdateProduct(context.Context, *UpdateProductRequest) (*UpdateProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProduct not implemented")
}
func (UnimplementedWarehouseServer) GetAllProducts(context.Context, *GetAllProductsRequest) (*GetAllProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllProducts not implemented")
}
func (UnimplementedWarehouseServer) GetProductById(context.Context, *GetProductByIdRequest) (*GetProductByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductById not implemented")
}
func (UnimplementedWarehouseServer) DeleteProduct(context.Context, *DeleteProductRequest) (*DeleteProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProduct not implemented")
}
func (UnimplementedWarehouseServer) CheckInventory(context.Context, *CheckInventoryRequest) (*CheckInventoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckInventory not implemented")
}
func (UnimplementedWarehouseServer) UpdateInventory(context.Context, *UpdateInventoryRequest) (*UpdateInventoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateInventory not implemented")
}
func (UnimplementedWarehouseServer) mustEmbedUnimplementedWarehouseServer() {}
func (UnimplementedWarehouseServer) testEmbeddedByValue()                   {}

// UnsafeWarehouseServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WarehouseServer will
// result in compilation errors.
type UnsafeWarehouseServer interface {
	mustEmbedUnimplementedWarehouseServer()
}

func RegisterWarehouseServer(s grpc.ServiceRegistrar, srv WarehouseServer) {
	// If the following call pancis, it indicates UnimplementedWarehouseServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Warehouse_ServiceDesc, srv)
}

func _Warehouse_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarehouseServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Warehouse_CreateProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarehouseServer).CreateProduct(ctx, req.(*CreateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Warehouse_UpdateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarehouseServer).UpdateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Warehouse_UpdateProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarehouseServer).UpdateProduct(ctx, req.(*UpdateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Warehouse_GetAllProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarehouseServer).GetAllProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Warehouse_GetAllProducts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarehouseServer).GetAllProducts(ctx, req.(*GetAllProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Warehouse_GetProductById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarehouseServer).GetProductById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Warehouse_GetProductById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarehouseServer).GetProductById(ctx, req.(*GetProductByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Warehouse_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarehouseServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Warehouse_DeleteProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarehouseServer).DeleteProduct(ctx, req.(*DeleteProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Warehouse_CheckInventory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckInventoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarehouseServer).CheckInventory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Warehouse_CheckInventory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarehouseServer).CheckInventory(ctx, req.(*CheckInventoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Warehouse_UpdateInventory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateInventoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarehouseServer).UpdateInventory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Warehouse_UpdateInventory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarehouseServer).UpdateInventory(ctx, req.(*UpdateInventoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Warehouse_ServiceDesc is the grpc.ServiceDesc for Warehouse service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Warehouse_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "warehouse.Warehouse",
	HandlerType: (*WarehouseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProduct",
			Handler:    _Warehouse_CreateProduct_Handler,
		},
		{
			MethodName: "UpdateProduct",
			Handler:    _Warehouse_UpdateProduct_Handler,
		},
		{
			MethodName: "GetAllProducts",
			Handler:    _Warehouse_GetAllProducts_Handler,
		},
		{
			MethodName: "GetProductById",
			Handler:    _Warehouse_GetProductById_Handler,
		},
		{
			MethodName: "DeleteProduct",
			Handler:    _Warehouse_DeleteProduct_Handler,
		},
		{
			MethodName: "CheckInventory",
			Handler:    _Warehouse_CheckInventory_Handler,
		},
		{
			MethodName: "UpdateInventory",
			Handler:    _Warehouse_UpdateInventory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "warehouse.proto",
}
