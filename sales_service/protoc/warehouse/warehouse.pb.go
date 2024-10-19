// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.1
// source: warehouse.proto

package warehouse

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

type GetProductByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId uint32 `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
}

func (x *GetProductByIdRequest) Reset() {
	*x = GetProductByIdRequest{}
	mi := &file_warehouse_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductByIdRequest) ProtoMessage() {}

func (x *GetProductByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_warehouse_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductByIdRequest.ProtoReflect.Descriptor instead.
func (*GetProductByIdRequest) Descriptor() ([]byte, []int) {
	return file_warehouse_proto_rawDescGZIP(), []int{0}
}

func (x *GetProductByIdRequest) GetProductId() uint32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Desc     string `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
	Price    uint32 `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
	Quantity uint32 `protobuf:"varint,5,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	mi := &file_warehouse_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_warehouse_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_warehouse_proto_rawDescGZIP(), []int{1}
}

func (x *Product) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *Product) GetPrice() uint32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Product) GetQuantity() uint32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type GetProductByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg     string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Status  uint32   `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Product *Product `protobuf:"bytes,3,opt,name=product,proto3" json:"product,omitempty"`
}

func (x *GetProductByIdResponse) Reset() {
	*x = GetProductByIdResponse{}
	mi := &file_warehouse_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductByIdResponse) ProtoMessage() {}

func (x *GetProductByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_warehouse_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductByIdResponse.ProtoReflect.Descriptor instead.
func (*GetProductByIdResponse) Descriptor() ([]byte, []int) {
	return file_warehouse_proto_rawDescGZIP(), []int{2}
}

func (x *GetProductByIdResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetProductByIdResponse) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GetProductByIdResponse) GetProduct() *Product {
	if x != nil {
		return x.Product
	}
	return nil
}

var File_warehouse_proto protoreflect.FileDescriptor

var file_warehouse_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x22, 0x36, 0x0a, 0x15,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x49, 0x64, 0x22, 0x73, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x70, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2c, 0x0a,
	0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x32, 0x62, 0x0a, 0x09, 0x57,
	0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x20, 0x2e, 0x77, 0x61, 0x72,
	0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x77,
	0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x1c, 0x5a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2f, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f,
	0x75, 0x73, 0x65, 0x3b, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_warehouse_proto_rawDescOnce sync.Once
	file_warehouse_proto_rawDescData = file_warehouse_proto_rawDesc
)

func file_warehouse_proto_rawDescGZIP() []byte {
	file_warehouse_proto_rawDescOnce.Do(func() {
		file_warehouse_proto_rawDescData = protoimpl.X.CompressGZIP(file_warehouse_proto_rawDescData)
	})
	return file_warehouse_proto_rawDescData
}

var file_warehouse_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_warehouse_proto_goTypes = []any{
	(*GetProductByIdRequest)(nil),  // 0: warehouse.GetProductByIdRequest
	(*Product)(nil),                // 1: warehouse.Product
	(*GetProductByIdResponse)(nil), // 2: warehouse.GetProductByIdResponse
}
var file_warehouse_proto_depIdxs = []int32{
	1, // 0: warehouse.GetProductByIdResponse.product:type_name -> warehouse.Product
	0, // 1: warehouse.Warehouse.GetProductById:input_type -> warehouse.GetProductByIdRequest
	2, // 2: warehouse.Warehouse.GetProductById:output_type -> warehouse.GetProductByIdResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_warehouse_proto_init() }
func file_warehouse_proto_init() {
	if File_warehouse_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_warehouse_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_warehouse_proto_goTypes,
		DependencyIndexes: file_warehouse_proto_depIdxs,
		MessageInfos:      file_warehouse_proto_msgTypes,
	}.Build()
	File_warehouse_proto = out.File
	file_warehouse_proto_rawDesc = nil
	file_warehouse_proto_goTypes = nil
	file_warehouse_proto_depIdxs = nil
}
