// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.1
// source: sales.proto

package sales

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateSalesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId     uint32 `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	DiscountPrice uint32 `protobuf:"varint,2,opt,name=discount_price,json=discountPrice,proto3" json:"discount_price,omitempty"`
	StartTime     string `protobuf:"bytes,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime       string `protobuf:"bytes,4,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
}

func (x *CreateSalesRequest) Reset() {
	*x = CreateSalesRequest{}
	mi := &file_sales_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateSalesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSalesRequest) ProtoMessage() {}

func (x *CreateSalesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sales_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSalesRequest.ProtoReflect.Descriptor instead.
func (*CreateSalesRequest) Descriptor() ([]byte, []int) {
	return file_sales_proto_rawDescGZIP(), []int{0}
}

func (x *CreateSalesRequest) GetProductId() uint32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *CreateSalesRequest) GetDiscountPrice() uint32 {
	if x != nil {
		return x.DiscountPrice
	}
	return 0
}

func (x *CreateSalesRequest) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *CreateSalesRequest) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

type CreateSalesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg    string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Status uint32 `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CreateSalesResponse) Reset() {
	*x = CreateSalesResponse{}
	mi := &file_sales_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateSalesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSalesResponse) ProtoMessage() {}

func (x *CreateSalesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sales_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSalesResponse.ProtoReflect.Descriptor instead.
func (*CreateSalesResponse) Descriptor() ([]byte, []int) {
	return file_sales_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSalesResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *CreateSalesResponse) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type GetAllSalesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllSalesRequest) Reset() {
	*x = GetAllSalesRequest{}
	mi := &file_sales_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllSalesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllSalesRequest) ProtoMessage() {}

func (x *GetAllSalesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sales_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllSalesRequest.ProtoReflect.Descriptor instead.
func (*GetAllSalesRequest) Descriptor() ([]byte, []int) {
	return file_sales_proto_rawDescGZIP(), []int{2}
}

type Sale struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ProductId     uint32 `protobuf:"varint,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	DiscountPrice uint32 `protobuf:"varint,3,opt,name=discount_price,json=discountPrice,proto3" json:"discount_price,omitempty"`
	StartTime     string `protobuf:"bytes,4,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime       string `protobuf:"bytes,5,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
}

func (x *Sale) Reset() {
	*x = Sale{}
	mi := &file_sales_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Sale) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sale) ProtoMessage() {}

func (x *Sale) ProtoReflect() protoreflect.Message {
	mi := &file_sales_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sale.ProtoReflect.Descriptor instead.
func (*Sale) Descriptor() ([]byte, []int) {
	return file_sales_proto_rawDescGZIP(), []int{3}
}

func (x *Sale) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Sale) GetProductId() uint32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *Sale) GetDiscountPrice() uint32 {
	if x != nil {
		return x.DiscountPrice
	}
	return 0
}

func (x *Sale) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *Sale) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

type GetAllSalesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sales []*Sale `protobuf:"bytes,1,rep,name=sales,proto3" json:"sales,omitempty"`
}

func (x *GetAllSalesResponse) Reset() {
	*x = GetAllSalesResponse{}
	mi := &file_sales_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllSalesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllSalesResponse) ProtoMessage() {}

func (x *GetAllSalesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sales_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllSalesResponse.ProtoReflect.Descriptor instead.
func (*GetAllSalesResponse) Descriptor() ([]byte, []int) {
	return file_sales_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllSalesResponse) GetSales() []*Sale {
	if x != nil {
		return x.Sales
	}
	return nil
}

type GetSalesByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SalesId uint32 `protobuf:"varint,1,opt,name=sales_id,json=salesId,proto3" json:"sales_id,omitempty"`
}

func (x *GetSalesByIdRequest) Reset() {
	*x = GetSalesByIdRequest{}
	mi := &file_sales_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSalesByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSalesByIdRequest) ProtoMessage() {}

func (x *GetSalesByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sales_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSalesByIdRequest.ProtoReflect.Descriptor instead.
func (*GetSalesByIdRequest) Descriptor() ([]byte, []int) {
	return file_sales_proto_rawDescGZIP(), []int{5}
}

func (x *GetSalesByIdRequest) GetSalesId() uint32 {
	if x != nil {
		return x.SalesId
	}
	return 0
}

type GetSalesByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg    string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Status uint32 `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Sale   *Sale  `protobuf:"bytes,3,opt,name=sale,proto3" json:"sale,omitempty"`
}

func (x *GetSalesByIdResponse) Reset() {
	*x = GetSalesByIdResponse{}
	mi := &file_sales_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSalesByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSalesByIdResponse) ProtoMessage() {}

func (x *GetSalesByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sales_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSalesByIdResponse.ProtoReflect.Descriptor instead.
func (*GetSalesByIdResponse) Descriptor() ([]byte, []int) {
	return file_sales_proto_rawDescGZIP(), []int{6}
}

func (x *GetSalesByIdResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetSalesByIdResponse) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GetSalesByIdResponse) GetSale() *Sale {
	if x != nil {
		return x.Sale
	}
	return nil
}

type UpdateSalesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ProductId     uint32 `protobuf:"varint,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	DiscountPrice uint32 `protobuf:"varint,3,opt,name=discount_price,json=discountPrice,proto3" json:"discount_price,omitempty"`
	StartTime     string `protobuf:"bytes,4,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime       string `protobuf:"bytes,5,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
}

func (x *UpdateSalesRequest) Reset() {
	*x = UpdateSalesRequest{}
	mi := &file_sales_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateSalesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSalesRequest) ProtoMessage() {}

func (x *UpdateSalesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sales_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSalesRequest.ProtoReflect.Descriptor instead.
func (*UpdateSalesRequest) Descriptor() ([]byte, []int) {
	return file_sales_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateSalesRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateSalesRequest) GetProductId() uint32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *UpdateSalesRequest) GetDiscountPrice() uint32 {
	if x != nil {
		return x.DiscountPrice
	}
	return 0
}

func (x *UpdateSalesRequest) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *UpdateSalesRequest) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

type UpdateSalesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg    string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Status uint32 `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UpdateSalesResponse) Reset() {
	*x = UpdateSalesResponse{}
	mi := &file_sales_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateSalesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSalesResponse) ProtoMessage() {}

func (x *UpdateSalesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sales_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSalesResponse.ProtoReflect.Descriptor instead.
func (*UpdateSalesResponse) Descriptor() ([]byte, []int) {
	return file_sales_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateSalesResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *UpdateSalesResponse) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type DeleteSalesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SalesId uint32 `protobuf:"varint,1,opt,name=sales_id,json=salesId,proto3" json:"sales_id,omitempty"`
}

func (x *DeleteSalesRequest) Reset() {
	*x = DeleteSalesRequest{}
	mi := &file_sales_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteSalesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSalesRequest) ProtoMessage() {}

func (x *DeleteSalesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sales_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSalesRequest.ProtoReflect.Descriptor instead.
func (*DeleteSalesRequest) Descriptor() ([]byte, []int) {
	return file_sales_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteSalesRequest) GetSalesId() uint32 {
	if x != nil {
		return x.SalesId
	}
	return 0
}

type DeleteSalesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg    string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Status uint32 `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *DeleteSalesResponse) Reset() {
	*x = DeleteSalesResponse{}
	mi := &file_sales_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteSalesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSalesResponse) ProtoMessage() {}

func (x *DeleteSalesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sales_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSalesResponse.ProtoReflect.Descriptor instead.
func (*DeleteSalesResponse) Descriptor() ([]byte, []int) {
	return file_sales_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteSalesResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *DeleteSalesResponse) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

var File_sales_proto protoreflect.FileDescriptor

var file_sales_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73,
	0x61, 0x6c, 0x65, 0x73, 0x22, 0x94, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x61, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0d, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x3f, 0x0a, 0x13, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x73, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x14, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x96, 0x01, 0x0a, 0x04, 0x53, 0x61, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0d, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x38, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x21, 0x0a, 0x05, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x2e, 0x53, 0x61, 0x6c, 0x65, 0x52, 0x05,
	0x73, 0x61, 0x6c, 0x65, 0x73, 0x22, 0x30, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x61, 0x6c, 0x65,
	0x73, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08,
	0x73, 0x61, 0x6c, 0x65, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07,
	0x73, 0x61, 0x6c, 0x65, 0x73, 0x49, 0x64, 0x22, 0x61, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x53, 0x61,
	0x6c, 0x65, 0x73, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x04, 0x73, 0x61, 0x6c,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x2e,
	0x53, 0x61, 0x6c, 0x65, 0x52, 0x04, 0x73, 0x61, 0x6c, 0x65, 0x22, 0xa4, 0x01, 0x0a, 0x12, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64,
	0x12, 0x25, 0x0a, 0x0e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x22, 0x3f, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x61, 0x6c, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x2f, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x61, 0x6c, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x61, 0x6c, 0x65,
	0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x61, 0x6c, 0x65,
	0x73, 0x49, 0x64, 0x22, 0x3f, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x61, 0x6c,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x32, 0xe8, 0x02, 0x0a, 0x05, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x12, 0x44,
	0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x12, 0x19, 0x2e,
	0x73, 0x61, 0x6c, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x61, 0x6c, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x61, 0x6c, 0x65, 0x73,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x61,
	0x6c, 0x65, 0x73, 0x12, 0x19, 0x2e, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x61, 0x6c,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0c, 0x47, 0x65,
	0x74, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1a, 0x2e, 0x73, 0x61, 0x6c,
	0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x42, 0x79, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x61, 0x6c,
	0x65, 0x73, 0x12, 0x19, 0x2e, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x73, 0x61, 0x6c, 0x65, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x61, 0x6c, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0b, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x12, 0x19, 0x2e, 0x73, 0x61, 0x6c, 0x65, 0x73,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x14, 0x5a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2f, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x3b,
	0x73, 0x61, 0x6c, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sales_proto_rawDescOnce sync.Once
	file_sales_proto_rawDescData = file_sales_proto_rawDesc
)

func file_sales_proto_rawDescGZIP() []byte {
	file_sales_proto_rawDescOnce.Do(func() {
		file_sales_proto_rawDescData = protoimpl.X.CompressGZIP(file_sales_proto_rawDescData)
	})
	return file_sales_proto_rawDescData
}

var file_sales_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_sales_proto_goTypes = []any{
	(*CreateSalesRequest)(nil),   // 0: sales.CreateSalesRequest
	(*CreateSalesResponse)(nil),  // 1: sales.CreateSalesResponse
	(*GetAllSalesRequest)(nil),   // 2: sales.GetAllSalesRequest
	(*Sale)(nil),                 // 3: sales.Sale
	(*GetAllSalesResponse)(nil),  // 4: sales.GetAllSalesResponse
	(*GetSalesByIdRequest)(nil),  // 5: sales.GetSalesByIdRequest
	(*GetSalesByIdResponse)(nil), // 6: sales.GetSalesByIdResponse
	(*UpdateSalesRequest)(nil),   // 7: sales.UpdateSalesRequest
	(*UpdateSalesResponse)(nil),  // 8: sales.UpdateSalesResponse
	(*DeleteSalesRequest)(nil),   // 9: sales.DeleteSalesRequest
	(*DeleteSalesResponse)(nil),  // 10: sales.DeleteSalesResponse
}
var file_sales_proto_depIdxs = []int32{
	3,  // 0: sales.GetAllSalesResponse.sales:type_name -> sales.Sale
	3,  // 1: sales.GetSalesByIdResponse.sale:type_name -> sales.Sale
	0,  // 2: sales.Sales.CreateSales:input_type -> sales.CreateSalesRequest
	2,  // 3: sales.Sales.GetAllSales:input_type -> sales.GetAllSalesRequest
	5,  // 4: sales.Sales.GetSalesById:input_type -> sales.GetSalesByIdRequest
	7,  // 5: sales.Sales.UpdateSales:input_type -> sales.UpdateSalesRequest
	9,  // 6: sales.Sales.DeleteSales:input_type -> sales.DeleteSalesRequest
	1,  // 7: sales.Sales.CreateSales:output_type -> sales.CreateSalesResponse
	4,  // 8: sales.Sales.GetAllSales:output_type -> sales.GetAllSalesResponse
	6,  // 9: sales.Sales.GetSalesById:output_type -> sales.GetSalesByIdResponse
	8,  // 10: sales.Sales.UpdateSales:output_type -> sales.UpdateSalesResponse
	10, // 11: sales.Sales.DeleteSales:output_type -> sales.DeleteSalesResponse
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_sales_proto_init() }
func file_sales_proto_init() {
	if File_sales_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sales_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sales_proto_goTypes,
		DependencyIndexes: file_sales_proto_depIdxs,
		MessageInfos:      file_sales_proto_msgTypes,
	}.Build()
	File_sales_proto = out.File
	file_sales_proto_rawDesc = nil
	file_sales_proto_goTypes = nil
	file_sales_proto_depIdxs = nil
}
