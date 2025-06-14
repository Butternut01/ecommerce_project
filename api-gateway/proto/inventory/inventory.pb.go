// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: proto/inventory.proto

package inventory

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProductRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price         float64                `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
	Stock         int32                  `protobuf:"varint,5,opt,name=stock,proto3" json:"stock,omitempty"`
	Category      string                 `protobuf:"bytes,6,opt,name=category,proto3" json:"category,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProductRequest) Reset() {
	*x = ProductRequest{}
	mi := &file_proto_inventory_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductRequest) ProtoMessage() {}

func (x *ProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inventory_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductRequest.ProtoReflect.Descriptor instead.
func (*ProductRequest) Descriptor() ([]byte, []int) {
	return file_proto_inventory_proto_rawDescGZIP(), []int{0}
}

func (x *ProductRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ProductRequest) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *ProductRequest) GetStock() int32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *ProductRequest) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

type ProductResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price         float64                `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
	Stock         int32                  `protobuf:"varint,5,opt,name=stock,proto3" json:"stock,omitempty"`
	Category      string                 `protobuf:"bytes,6,opt,name=category,proto3" json:"category,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProductResponse) Reset() {
	*x = ProductResponse{}
	mi := &file_proto_inventory_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductResponse) ProtoMessage() {}

func (x *ProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inventory_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductResponse.ProtoReflect.Descriptor instead.
func (*ProductResponse) Descriptor() ([]byte, []int) {
	return file_proto_inventory_proto_rawDescGZIP(), []int{1}
}

func (x *ProductResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ProductResponse) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *ProductResponse) GetStock() int32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *ProductResponse) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

type GetProductRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProductRequest) Reset() {
	*x = GetProductRequest{}
	mi := &file_proto_inventory_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductRequest) ProtoMessage() {}

func (x *GetProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inventory_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductRequest.ProtoReflect.Descriptor instead.
func (*GetProductRequest) Descriptor() ([]byte, []int) {
	return file_proto_inventory_proto_rawDescGZIP(), []int{2}
}

func (x *GetProductRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteProductRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteProductRequest) Reset() {
	*x = DeleteProductRequest{}
	mi := &file_proto_inventory_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProductRequest) ProtoMessage() {}

func (x *DeleteProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inventory_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProductRequest.ProtoReflect.Descriptor instead.
func (*DeleteProductRequest) Descriptor() ([]byte, []int) {
	return file_proto_inventory_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteProductRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteProductResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteProductResponse) Reset() {
	*x = DeleteProductResponse{}
	mi := &file_proto_inventory_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProductResponse) ProtoMessage() {}

func (x *DeleteProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inventory_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProductResponse.ProtoReflect.Descriptor instead.
func (*DeleteProductResponse) Descriptor() ([]byte, []int) {
	return file_proto_inventory_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteProductResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ListProductsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Category      string                 `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
	MinPrice      float64                `protobuf:"fixed64,3,opt,name=min_price,json=minPrice,proto3" json:"min_price,omitempty"`
	MaxPrice      float64                `protobuf:"fixed64,4,opt,name=max_price,json=maxPrice,proto3" json:"max_price,omitempty"`
	Page          int32                  `protobuf:"varint,5,opt,name=page,proto3" json:"page,omitempty"`
	Limit         int32                  `protobuf:"varint,6,opt,name=limit,proto3" json:"limit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListProductsRequest) Reset() {
	*x = ListProductsRequest{}
	mi := &file_proto_inventory_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListProductsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProductsRequest) ProtoMessage() {}

func (x *ListProductsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inventory_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProductsRequest.ProtoReflect.Descriptor instead.
func (*ListProductsRequest) Descriptor() ([]byte, []int) {
	return file_proto_inventory_proto_rawDescGZIP(), []int{5}
}

func (x *ListProductsRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListProductsRequest) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *ListProductsRequest) GetMinPrice() float64 {
	if x != nil {
		return x.MinPrice
	}
	return 0
}

func (x *ListProductsRequest) GetMaxPrice() float64 {
	if x != nil {
		return x.MaxPrice
	}
	return 0
}

func (x *ListProductsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListProductsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ListProductsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Products      []*ProductResponse     `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListProductsResponse) Reset() {
	*x = ListProductsResponse{}
	mi := &file_proto_inventory_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListProductsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProductsResponse) ProtoMessage() {}

func (x *ListProductsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inventory_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProductsResponse.ProtoReflect.Descriptor instead.
func (*ListProductsResponse) Descriptor() ([]byte, []int) {
	return file_proto_inventory_proto_rawDescGZIP(), []int{6}
}

func (x *ListProductsResponse) GetProducts() []*ProductResponse {
	if x != nil {
		return x.Products
	}
	return nil
}

var File_proto_inventory_proto protoreflect.FileDescriptor

const file_proto_inventory_proto_rawDesc = "" +
	"\n" +
	"\x15proto/inventory.proto\x12\tinventory\"\x9e\x01\n" +
	"\x0eProductRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12\x14\n" +
	"\x05price\x18\x04 \x01(\x01R\x05price\x12\x14\n" +
	"\x05stock\x18\x05 \x01(\x05R\x05stock\x12\x1a\n" +
	"\bcategory\x18\x06 \x01(\tR\bcategory\"\x9f\x01\n" +
	"\x0fProductResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12\x14\n" +
	"\x05price\x18\x04 \x01(\x01R\x05price\x12\x14\n" +
	"\x05stock\x18\x05 \x01(\x05R\x05stock\x12\x1a\n" +
	"\bcategory\x18\x06 \x01(\tR\bcategory\"#\n" +
	"\x11GetProductRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"&\n" +
	"\x14DeleteProductRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"1\n" +
	"\x15DeleteProductResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\"\xa9\x01\n" +
	"\x13ListProductsRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x1a\n" +
	"\bcategory\x18\x02 \x01(\tR\bcategory\x12\x1b\n" +
	"\tmin_price\x18\x03 \x01(\x01R\bminPrice\x12\x1b\n" +
	"\tmax_price\x18\x04 \x01(\x01R\bmaxPrice\x12\x12\n" +
	"\x04page\x18\x05 \x01(\x05R\x04page\x12\x14\n" +
	"\x05limit\x18\x06 \x01(\x05R\x05limit\"N\n" +
	"\x14ListProductsResponse\x126\n" +
	"\bproducts\x18\x01 \x03(\v2\x1a.inventory.ProductResponseR\bproducts2\x8f\x03\n" +
	"\x10InventoryService\x12F\n" +
	"\rCreateProduct\x12\x19.inventory.ProductRequest\x1a\x1a.inventory.ProductResponse\x12F\n" +
	"\n" +
	"GetProduct\x12\x1c.inventory.GetProductRequest\x1a\x1a.inventory.ProductResponse\x12F\n" +
	"\rUpdateProduct\x12\x19.inventory.ProductRequest\x1a\x1a.inventory.ProductResponse\x12R\n" +
	"\rDeleteProduct\x12\x1f.inventory.DeleteProductRequest\x1a .inventory.DeleteProductResponse\x12O\n" +
	"\fListProducts\x12\x1e.inventory.ListProductsRequest\x1a\x1f.inventory.ListProductsResponseB\x1dZ\x1bapi-gateway/proto/inventoryb\x06proto3"

var (
	file_proto_inventory_proto_rawDescOnce sync.Once
	file_proto_inventory_proto_rawDescData []byte
)

func file_proto_inventory_proto_rawDescGZIP() []byte {
	file_proto_inventory_proto_rawDescOnce.Do(func() {
		file_proto_inventory_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_inventory_proto_rawDesc), len(file_proto_inventory_proto_rawDesc)))
	})
	return file_proto_inventory_proto_rawDescData
}

var file_proto_inventory_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_inventory_proto_goTypes = []any{
	(*ProductRequest)(nil),        // 0: inventory.ProductRequest
	(*ProductResponse)(nil),       // 1: inventory.ProductResponse
	(*GetProductRequest)(nil),     // 2: inventory.GetProductRequest
	(*DeleteProductRequest)(nil),  // 3: inventory.DeleteProductRequest
	(*DeleteProductResponse)(nil), // 4: inventory.DeleteProductResponse
	(*ListProductsRequest)(nil),   // 5: inventory.ListProductsRequest
	(*ListProductsResponse)(nil),  // 6: inventory.ListProductsResponse
}
var file_proto_inventory_proto_depIdxs = []int32{
	1, // 0: inventory.ListProductsResponse.products:type_name -> inventory.ProductResponse
	0, // 1: inventory.InventoryService.CreateProduct:input_type -> inventory.ProductRequest
	2, // 2: inventory.InventoryService.GetProduct:input_type -> inventory.GetProductRequest
	0, // 3: inventory.InventoryService.UpdateProduct:input_type -> inventory.ProductRequest
	3, // 4: inventory.InventoryService.DeleteProduct:input_type -> inventory.DeleteProductRequest
	5, // 5: inventory.InventoryService.ListProducts:input_type -> inventory.ListProductsRequest
	1, // 6: inventory.InventoryService.CreateProduct:output_type -> inventory.ProductResponse
	1, // 7: inventory.InventoryService.GetProduct:output_type -> inventory.ProductResponse
	1, // 8: inventory.InventoryService.UpdateProduct:output_type -> inventory.ProductResponse
	4, // 9: inventory.InventoryService.DeleteProduct:output_type -> inventory.DeleteProductResponse
	6, // 10: inventory.InventoryService.ListProducts:output_type -> inventory.ListProductsResponse
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_inventory_proto_init() }
func file_proto_inventory_proto_init() {
	if File_proto_inventory_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_inventory_proto_rawDesc), len(file_proto_inventory_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_inventory_proto_goTypes,
		DependencyIndexes: file_proto_inventory_proto_depIdxs,
		MessageInfos:      file_proto_inventory_proto_msgTypes,
	}.Build()
	File_proto_inventory_proto = out.File
	file_proto_inventory_proto_goTypes = nil
	file_proto_inventory_proto_depIdxs = nil
}
