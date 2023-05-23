// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.4
// source: stock_service.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SubscribeStockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Symbol string `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Name   string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SubscribeStockRequest) Reset() {
	*x = SubscribeStockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeStockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeStockRequest) ProtoMessage() {}

func (x *SubscribeStockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stock_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeStockRequest.ProtoReflect.Descriptor instead.
func (*SubscribeStockRequest) Descriptor() ([]byte, []int) {
	return file_stock_service_proto_rawDescGZIP(), []int{0}
}

func (x *SubscribeStockRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *SubscribeStockRequest) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *SubscribeStockRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UnSubscribeStockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Symbol string `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
}

func (x *UnSubscribeStockRequest) Reset() {
	*x = UnSubscribeStockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnSubscribeStockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnSubscribeStockRequest) ProtoMessage() {}

func (x *UnSubscribeStockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stock_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnSubscribeStockRequest.ProtoReflect.Descriptor instead.
func (*UnSubscribeStockRequest) Descriptor() ([]byte, []int) {
	return file_stock_service_proto_rawDescGZIP(), []int{1}
}

func (x *UnSubscribeStockRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UnSubscribeStockRequest) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

type TradeData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date        string  `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Close       float64 `protobuf:"fixed64,2,opt,name=close,proto3" json:"close,omitempty"`
	Volume      int64   `protobuf:"varint,3,opt,name=volume,proto3" json:"volume,omitempty"`
	XueqiuCount int64   `protobuf:"varint,4,opt,name=xueqiu_count,json=xueqiuCount,proto3" json:"xueqiu_count,omitempty"`
}

func (x *TradeData) Reset() {
	*x = TradeData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TradeData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TradeData) ProtoMessage() {}

func (x *TradeData) ProtoReflect() protoreflect.Message {
	mi := &file_stock_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TradeData.ProtoReflect.Descriptor instead.
func (*TradeData) Descriptor() ([]byte, []int) {
	return file_stock_service_proto_rawDescGZIP(), []int{2}
}

func (x *TradeData) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *TradeData) GetClose() float64 {
	if x != nil {
		return x.Close
	}
	return 0
}

func (x *TradeData) GetVolume() int64 {
	if x != nil {
		return x.Volume
	}
	return 0
}

func (x *TradeData) GetXueqiuCount() int64 {
	if x != nil {
		return x.XueqiuCount
	}
	return 0
}

type StockInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Symbol string `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Name   string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Bull   int64  `protobuf:"varint,4,opt,name=bull,proto3" json:"bull,omitempty"`
}

func (x *StockInfo) Reset() {
	*x = StockInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockInfo) ProtoMessage() {}

func (x *StockInfo) ProtoReflect() protoreflect.Message {
	mi := &file_stock_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockInfo.ProtoReflect.Descriptor instead.
func (*StockInfo) Descriptor() ([]byte, []int) {
	return file_stock_service_proto_rawDescGZIP(), []int{3}
}

func (x *StockInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *StockInfo) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *StockInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StockInfo) GetBull() int64 {
	if x != nil {
		return x.Bull
	}
	return 0
}

type SubscribeStocksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *SubscribeStocksRequest) Reset() {
	*x = SubscribeStocksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeStocksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeStocksRequest) ProtoMessage() {}

func (x *SubscribeStocksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stock_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeStocksRequest.ProtoReflect.Descriptor instead.
func (*SubscribeStocksRequest) Descriptor() ([]byte, []int) {
	return file_stock_service_proto_rawDescGZIP(), []int{4}
}

func (x *SubscribeStocksRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type SubscribeStocksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*StockInfo `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *SubscribeStocksResponse) Reset() {
	*x = SubscribeStocksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeStocksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeStocksResponse) ProtoMessage() {}

func (x *SubscribeStocksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stock_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeStocksResponse.ProtoReflect.Descriptor instead.
func (*SubscribeStocksResponse) Descriptor() ([]byte, []int) {
	return file_stock_service_proto_rawDescGZIP(), []int{5}
}

func (x *SubscribeStocksResponse) GetData() []*StockInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type StockTradeDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *StockTradeDataRequest) Reset() {
	*x = StockTradeDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockTradeDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockTradeDataRequest) ProtoMessage() {}

func (x *StockTradeDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stock_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockTradeDataRequest.ProtoReflect.Descriptor instead.
func (*StockTradeDataRequest) Descriptor() ([]byte, []int) {
	return file_stock_service_proto_rawDescGZIP(), []int{6}
}

func (x *StockTradeDataRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type StockTradeDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info      *StockInfo   `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	TradeData []*TradeData `protobuf:"bytes,2,rep,name=trade_data,json=tradeData,proto3" json:"trade_data,omitempty"`
}

func (x *StockTradeDataResponse) Reset() {
	*x = StockTradeDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockTradeDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockTradeDataResponse) ProtoMessage() {}

func (x *StockTradeDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stock_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockTradeDataResponse.ProtoReflect.Descriptor instead.
func (*StockTradeDataResponse) Descriptor() ([]byte, []int) {
	return file_stock_service_proto_rawDescGZIP(), []int{7}
}

func (x *StockTradeDataResponse) GetInfo() *StockInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *StockTradeDataResponse) GetTradeData() []*TradeData {
	if x != nil {
		return x.TradeData
	}
	return nil
}

type SearchStockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbolorname string `protobuf:"bytes,1,opt,name=symbolorname,proto3" json:"symbolorname,omitempty"`
}

func (x *SearchStockRequest) Reset() {
	*x = SearchStockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchStockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchStockRequest) ProtoMessage() {}

func (x *SearchStockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stock_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchStockRequest.ProtoReflect.Descriptor instead.
func (*SearchStockRequest) Descriptor() ([]byte, []int) {
	return file_stock_service_proto_rawDescGZIP(), []int{8}
}

func (x *SearchStockRequest) GetSymbolorname() string {
	if x != nil {
		return x.Symbolorname
	}
	return ""
}

type SearchStockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stocks []*StockInfo `protobuf:"bytes,1,rep,name=stocks,proto3" json:"stocks,omitempty"`
}

func (x *SearchStockResponse) Reset() {
	*x = SearchStockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchStockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchStockResponse) ProtoMessage() {}

func (x *SearchStockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stock_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchStockResponse.ProtoReflect.Descriptor instead.
func (*SearchStockResponse) Descriptor() ([]byte, []int) {
	return file_stock_service_proto_rawDescGZIP(), []int{9}
}

func (x *SearchStockResponse) GetStocks() []*StockInfo {
	if x != nil {
		return x.Stocks
	}
	return nil
}

var File_stock_service_proto protoreflect.FileDescriptor

var file_stock_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5c, 0x0a, 0x15, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d,
	0x62, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4a, 0x0a, 0x17, 0x55, 0x6e, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d,
	0x62, 0x6f, 0x6c, 0x22, 0x70, 0x0a, 0x09, 0x54, 0x72, 0x61, 0x64, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x78, 0x75, 0x65, 0x71, 0x69, 0x75, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x78, 0x75, 0x65, 0x71, 0x69, 0x75,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x5b, 0x0a, 0x09, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x62, 0x75, 0x6c, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x62, 0x75,
	0x6c, 0x6c, 0x22, 0x31, 0x0a, 0x16, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53,
	0x74, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x3f, 0x0a, 0x17, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x24, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x27, 0x0a, 0x15, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x54,
	0x72, 0x61, 0x64, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x6f, 0x0a, 0x16, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x64, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e,
	0x53, 0x74, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12,
	0x2f, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x54, 0x72, 0x61, 0x64,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x09, 0x74, 0x72, 0x61, 0x64, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x22, 0x38, 0x0a, 0x12, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c,
	0x6f, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x6f, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3f, 0x0a, 0x13, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x28, 0x0a, 0x06, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x06, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x32, 0x94, 0x03, 0x0a, 0x0c,
	0x53, 0x74, 0x6f, 0x63, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x0e,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x1c,
	0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x10, 0x55, 0x6e, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x1e, 0x2e, 0x73, 0x74, 0x6f,
	0x63, 0x6b, 0x2e, 0x55, 0x6e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x74,
	0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x64, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x1d, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b,
	0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0b, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x19, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b,
	0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x4f, 0x0a, 0x0e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x64, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x1c, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x53, 0x74, 0x6f, 0x63,
	0x6b, 0x54, 0x72, 0x61, 0x64, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x54,
	0x72, 0x61, 0x64, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_stock_service_proto_rawDescOnce sync.Once
	file_stock_service_proto_rawDescData = file_stock_service_proto_rawDesc
)

func file_stock_service_proto_rawDescGZIP() []byte {
	file_stock_service_proto_rawDescOnce.Do(func() {
		file_stock_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_stock_service_proto_rawDescData)
	})
	return file_stock_service_proto_rawDescData
}

var file_stock_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_stock_service_proto_goTypes = []interface{}{
	(*SubscribeStockRequest)(nil),   // 0: stock.SubscribeStockRequest
	(*UnSubscribeStockRequest)(nil), // 1: stock.UnSubscribeStockRequest
	(*TradeData)(nil),               // 2: stock.TradeData
	(*StockInfo)(nil),               // 3: stock.StockInfo
	(*SubscribeStocksRequest)(nil),  // 4: stock.SubscribeStocksRequest
	(*SubscribeStocksResponse)(nil), // 5: stock.SubscribeStocksResponse
	(*StockTradeDataRequest)(nil),   // 6: stock.StockTradeDataRequest
	(*StockTradeDataResponse)(nil),  // 7: stock.StockTradeDataResponse
	(*SearchStockRequest)(nil),      // 8: stock.SearchStockRequest
	(*SearchStockResponse)(nil),     // 9: stock.SearchStockResponse
	(*emptypb.Empty)(nil),           // 10: google.protobuf.Empty
}
var file_stock_service_proto_depIdxs = []int32{
	3,  // 0: stock.SubscribeStocksResponse.data:type_name -> stock.StockInfo
	3,  // 1: stock.StockTradeDataResponse.info:type_name -> stock.StockInfo
	2,  // 2: stock.StockTradeDataResponse.trade_data:type_name -> stock.TradeData
	3,  // 3: stock.SearchStockResponse.stocks:type_name -> stock.StockInfo
	0,  // 4: stock.StockService.SubscribeStock:input_type -> stock.SubscribeStockRequest
	1,  // 5: stock.StockService.UnSubscribeStock:input_type -> stock.UnSubscribeStockRequest
	4,  // 6: stock.StockService.SubscribedStocks:input_type -> stock.SubscribeStocksRequest
	8,  // 7: stock.StockService.SearchStock:input_type -> stock.SearchStockRequest
	6,  // 8: stock.StockService.StockTradeData:input_type -> stock.StockTradeDataRequest
	10, // 9: stock.StockService.SubscribeStock:output_type -> google.protobuf.Empty
	10, // 10: stock.StockService.UnSubscribeStock:output_type -> google.protobuf.Empty
	5,  // 11: stock.StockService.SubscribedStocks:output_type -> stock.SubscribeStocksResponse
	9,  // 12: stock.StockService.SearchStock:output_type -> stock.SearchStockResponse
	7,  // 13: stock.StockService.StockTradeData:output_type -> stock.StockTradeDataResponse
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_stock_service_proto_init() }
func file_stock_service_proto_init() {
	if File_stock_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stock_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeStockRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_stock_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnSubscribeStockRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_stock_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TradeData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_stock_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_stock_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeStocksRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_stock_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeStocksResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_stock_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockTradeDataRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_stock_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockTradeDataResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_stock_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchStockRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_stock_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchStockResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_stock_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stock_service_proto_goTypes,
		DependencyIndexes: file_stock_service_proto_depIdxs,
		MessageInfos:      file_stock_service_proto_msgTypes,
	}.Build()
	File_stock_service_proto = out.File
	file_stock_service_proto_rawDesc = nil
	file_stock_service_proto_goTypes = nil
	file_stock_service_proto_depIdxs = nil
}
