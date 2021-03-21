// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: calculator/calcpb/calc.proto

package calcpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type SumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Summand1 int32 `protobuf:"varint,1,opt,name=summand1,proto3" json:"summand1,omitempty"`
	Summand2 int32 `protobuf:"varint,2,opt,name=summand2,proto3" json:"summand2,omitempty"`
}

func (x *SumRequest) Reset() {
	*x = SumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calcpb_calc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumRequest) ProtoMessage() {}

func (x *SumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calcpb_calc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumRequest.ProtoReflect.Descriptor instead.
func (*SumRequest) Descriptor() ([]byte, []int) {
	return file_calculator_calcpb_calc_proto_rawDescGZIP(), []int{0}
}

func (x *SumRequest) GetSummand1() int32 {
	if x != nil {
		return x.Summand1
	}
	return 0
}

func (x *SumRequest) GetSummand2() int32 {
	if x != nil {
		return x.Summand2
	}
	return 0
}

type SumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SumResult int32 `protobuf:"varint,1,opt,name=sum_result,json=sumResult,proto3" json:"sum_result,omitempty"`
}

func (x *SumResponse) Reset() {
	*x = SumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calcpb_calc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumResponse) ProtoMessage() {}

func (x *SumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calcpb_calc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumResponse.ProtoReflect.Descriptor instead.
func (*SumResponse) Descriptor() ([]byte, []int) {
	return file_calculator_calcpb_calc_proto_rawDescGZIP(), []int{1}
}

func (x *SumResponse) GetSumResult() int32 {
	if x != nil {
		return x.SumResult
	}
	return 0
}

type PrimeNumberDecompositionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *PrimeNumberDecompositionRequest) Reset() {
	*x = PrimeNumberDecompositionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calcpb_calc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimeNumberDecompositionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeNumberDecompositionRequest) ProtoMessage() {}

func (x *PrimeNumberDecompositionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calcpb_calc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimeNumberDecompositionRequest.ProtoReflect.Descriptor instead.
func (*PrimeNumberDecompositionRequest) Descriptor() ([]byte, []int) {
	return file_calculator_calcpb_calc_proto_rawDescGZIP(), []int{2}
}

func (x *PrimeNumberDecompositionRequest) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type PrimeNumberDecompositionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *PrimeNumberDecompositionResponse) Reset() {
	*x = PrimeNumberDecompositionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calcpb_calc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimeNumberDecompositionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeNumberDecompositionResponse) ProtoMessage() {}

func (x *PrimeNumberDecompositionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calcpb_calc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimeNumberDecompositionResponse.ProtoReflect.Descriptor instead.
func (*PrimeNumberDecompositionResponse) Descriptor() ([]byte, []int) {
	return file_calculator_calcpb_calc_proto_rawDescGZIP(), []int{3}
}

func (x *PrimeNumberDecompositionResponse) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type ComputeAverageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *ComputeAverageRequest) Reset() {
	*x = ComputeAverageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calcpb_calc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputeAverageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputeAverageRequest) ProtoMessage() {}

func (x *ComputeAverageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calcpb_calc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputeAverageRequest.ProtoReflect.Descriptor instead.
func (*ComputeAverageRequest) Descriptor() ([]byte, []int) {
	return file_calculator_calcpb_calc_proto_rawDescGZIP(), []int{4}
}

func (x *ComputeAverageRequest) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type ComputeAverageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result float64 `protobuf:"fixed64,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *ComputeAverageResponse) Reset() {
	*x = ComputeAverageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calcpb_calc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputeAverageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputeAverageResponse) ProtoMessage() {}

func (x *ComputeAverageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calcpb_calc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputeAverageResponse.ProtoReflect.Descriptor instead.
func (*ComputeAverageResponse) Descriptor() ([]byte, []int) {
	return file_calculator_calcpb_calc_proto_rawDescGZIP(), []int{5}
}

func (x *ComputeAverageResponse) GetResult() float64 {
	if x != nil {
		return x.Result
	}
	return 0
}

type FindMaximumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *FindMaximumRequest) Reset() {
	*x = FindMaximumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calcpb_calc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindMaximumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindMaximumRequest) ProtoMessage() {}

func (x *FindMaximumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calcpb_calc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindMaximumRequest.ProtoReflect.Descriptor instead.
func (*FindMaximumRequest) Descriptor() ([]byte, []int) {
	return file_calculator_calcpb_calc_proto_rawDescGZIP(), []int{6}
}

func (x *FindMaximumRequest) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type FindMaximumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *FindMaximumResponse) Reset() {
	*x = FindMaximumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calcpb_calc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindMaximumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindMaximumResponse) ProtoMessage() {}

func (x *FindMaximumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calcpb_calc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindMaximumResponse.ProtoReflect.Descriptor instead.
func (*FindMaximumResponse) Descriptor() ([]byte, []int) {
	return file_calculator_calcpb_calc_proto_rawDescGZIP(), []int{7}
}

func (x *FindMaximumResponse) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_calculator_calcpb_calc_proto protoreflect.FileDescriptor

var file_calculator_calcpb_calc_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x63, 0x61, 0x6c,
	0x63, 0x70, 0x62, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x63, 0x61, 0x6c, 0x63, 0x70, 0x62, 0x22, 0x44, 0x0a, 0x0a, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x31,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x31,
	0x12, 0x1a, 0x0a, 0x08, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x32, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x32, 0x22, 0x2c, 0x0a, 0x0b,
	0x53, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x75, 0x6d, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x73, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x39, 0x0a, 0x1f, 0x50, 0x72,
	0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x44, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x3a, 0x0a, 0x20, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x44, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x22, 0x2f, 0x0a, 0x15, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x41, 0x76, 0x65, 0x72,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x22, 0x30, 0x0a, 0x16, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x41, 0x76, 0x65,
	0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x22, 0x2c, 0x0a, 0x12, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x61, 0x78, 0x69,
	0x6d, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x22, 0x2d, 0x0a, 0x13, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x32, 0xd5, 0x02, 0x0a, 0x0b, 0x43, 0x61, 0x6c, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x30, 0x0a, 0x03, 0x53, 0x75, 0x6d, 0x12, 0x12, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x70,
	0x62, 0x2e, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x63,
	0x61, 0x6c, 0x63, 0x70, 0x62, 0x2e, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x71, 0x0a, 0x18, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x44, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x27, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x44, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x70,
	0x62, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x44, 0x65, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x53, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74,
	0x65, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x70,
	0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x70, 0x62,
	0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x4c, 0x0a, 0x0b, 0x46,
	0x69, 0x6e, 0x64, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x12, 0x1a, 0x2e, 0x63, 0x61, 0x6c,
	0x63, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x70, 0x62, 0x2e,
	0x46, 0x69, 0x6e, 0x64, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x13, 0x5a, 0x11, 0x63, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_calculator_calcpb_calc_proto_rawDescOnce sync.Once
	file_calculator_calcpb_calc_proto_rawDescData = file_calculator_calcpb_calc_proto_rawDesc
)

func file_calculator_calcpb_calc_proto_rawDescGZIP() []byte {
	file_calculator_calcpb_calc_proto_rawDescOnce.Do(func() {
		file_calculator_calcpb_calc_proto_rawDescData = protoimpl.X.CompressGZIP(file_calculator_calcpb_calc_proto_rawDescData)
	})
	return file_calculator_calcpb_calc_proto_rawDescData
}

var file_calculator_calcpb_calc_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_calculator_calcpb_calc_proto_goTypes = []interface{}{
	(*SumRequest)(nil),                       // 0: calcpb.SumRequest
	(*SumResponse)(nil),                      // 1: calcpb.SumResponse
	(*PrimeNumberDecompositionRequest)(nil),  // 2: calcpb.PrimeNumberDecompositionRequest
	(*PrimeNumberDecompositionResponse)(nil), // 3: calcpb.PrimeNumberDecompositionResponse
	(*ComputeAverageRequest)(nil),            // 4: calcpb.ComputeAverageRequest
	(*ComputeAverageResponse)(nil),           // 5: calcpb.ComputeAverageResponse
	(*FindMaximumRequest)(nil),               // 6: calcpb.FindMaximumRequest
	(*FindMaximumResponse)(nil),              // 7: calcpb.FindMaximumResponse
}
var file_calculator_calcpb_calc_proto_depIdxs = []int32{
	0, // 0: calcpb.CalcService.Sum:input_type -> calcpb.SumRequest
	2, // 1: calcpb.CalcService.PrimeNumberDecomposition:input_type -> calcpb.PrimeNumberDecompositionRequest
	4, // 2: calcpb.CalcService.ComputeAverage:input_type -> calcpb.ComputeAverageRequest
	6, // 3: calcpb.CalcService.FindMaximum:input_type -> calcpb.FindMaximumRequest
	1, // 4: calcpb.CalcService.Sum:output_type -> calcpb.SumResponse
	3, // 5: calcpb.CalcService.PrimeNumberDecomposition:output_type -> calcpb.PrimeNumberDecompositionResponse
	5, // 6: calcpb.CalcService.ComputeAverage:output_type -> calcpb.ComputeAverageResponse
	7, // 7: calcpb.CalcService.FindMaximum:output_type -> calcpb.FindMaximumResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_calculator_calcpb_calc_proto_init() }
func file_calculator_calcpb_calc_proto_init() {
	if File_calculator_calcpb_calc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_calculator_calcpb_calc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumRequest); i {
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
		file_calculator_calcpb_calc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumResponse); i {
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
		file_calculator_calcpb_calc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimeNumberDecompositionRequest); i {
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
		file_calculator_calcpb_calc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimeNumberDecompositionResponse); i {
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
		file_calculator_calcpb_calc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputeAverageRequest); i {
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
		file_calculator_calcpb_calc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputeAverageResponse); i {
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
		file_calculator_calcpb_calc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindMaximumRequest); i {
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
		file_calculator_calcpb_calc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindMaximumResponse); i {
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
			RawDescriptor: file_calculator_calcpb_calc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calculator_calcpb_calc_proto_goTypes,
		DependencyIndexes: file_calculator_calcpb_calc_proto_depIdxs,
		MessageInfos:      file_calculator_calcpb_calc_proto_msgTypes,
	}.Build()
	File_calculator_calcpb_calc_proto = out.File
	file_calculator_calcpb_calc_proto_rawDesc = nil
	file_calculator_calcpb_calc_proto_goTypes = nil
	file_calculator_calcpb_calc_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CalcServiceClient is the client API for CalcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalcServiceClient interface {
	// unary
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
	// server streaming
	PrimeNumberDecomposition(ctx context.Context, in *PrimeNumberDecompositionRequest, opts ...grpc.CallOption) (CalcService_PrimeNumberDecompositionClient, error)
	// Client Streaming
	ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (CalcService_ComputeAverageClient, error)
	// BiDi Streaming
	FindMaximum(ctx context.Context, opts ...grpc.CallOption) (CalcService_FindMaximumClient, error)
}

type calcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalcServiceClient(cc grpc.ClientConnInterface) CalcServiceClient {
	return &calcServiceClient{cc}
}

func (c *calcServiceClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/calcpb.CalcService/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calcServiceClient) PrimeNumberDecomposition(ctx context.Context, in *PrimeNumberDecompositionRequest, opts ...grpc.CallOption) (CalcService_PrimeNumberDecompositionClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalcService_serviceDesc.Streams[0], "/calcpb.CalcService/PrimeNumberDecomposition", opts...)
	if err != nil {
		return nil, err
	}
	x := &calcServicePrimeNumberDecompositionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalcService_PrimeNumberDecompositionClient interface {
	Recv() (*PrimeNumberDecompositionResponse, error)
	grpc.ClientStream
}

type calcServicePrimeNumberDecompositionClient struct {
	grpc.ClientStream
}

func (x *calcServicePrimeNumberDecompositionClient) Recv() (*PrimeNumberDecompositionResponse, error) {
	m := new(PrimeNumberDecompositionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calcServiceClient) ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (CalcService_ComputeAverageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalcService_serviceDesc.Streams[1], "/calcpb.CalcService/ComputeAverage", opts...)
	if err != nil {
		return nil, err
	}
	x := &calcServiceComputeAverageClient{stream}
	return x, nil
}

type CalcService_ComputeAverageClient interface {
	Send(*ComputeAverageRequest) error
	CloseAndRecv() (*ComputeAverageResponse, error)
	grpc.ClientStream
}

type calcServiceComputeAverageClient struct {
	grpc.ClientStream
}

func (x *calcServiceComputeAverageClient) Send(m *ComputeAverageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calcServiceComputeAverageClient) CloseAndRecv() (*ComputeAverageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ComputeAverageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calcServiceClient) FindMaximum(ctx context.Context, opts ...grpc.CallOption) (CalcService_FindMaximumClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalcService_serviceDesc.Streams[2], "/calcpb.CalcService/FindMaximum", opts...)
	if err != nil {
		return nil, err
	}
	x := &calcServiceFindMaximumClient{stream}
	return x, nil
}

type CalcService_FindMaximumClient interface {
	Send(*FindMaximumRequest) error
	Recv() (*FindMaximumResponse, error)
	grpc.ClientStream
}

type calcServiceFindMaximumClient struct {
	grpc.ClientStream
}

func (x *calcServiceFindMaximumClient) Send(m *FindMaximumRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calcServiceFindMaximumClient) Recv() (*FindMaximumResponse, error) {
	m := new(FindMaximumResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalcServiceServer is the server API for CalcService service.
type CalcServiceServer interface {
	// unary
	Sum(context.Context, *SumRequest) (*SumResponse, error)
	// server streaming
	PrimeNumberDecomposition(*PrimeNumberDecompositionRequest, CalcService_PrimeNumberDecompositionServer) error
	// Client Streaming
	ComputeAverage(CalcService_ComputeAverageServer) error
	// BiDi Streaming
	FindMaximum(CalcService_FindMaximumServer) error
}

// UnimplementedCalcServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalcServiceServer struct {
}

func (*UnimplementedCalcServiceServer) Sum(context.Context, *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (*UnimplementedCalcServiceServer) PrimeNumberDecomposition(*PrimeNumberDecompositionRequest, CalcService_PrimeNumberDecompositionServer) error {
	return status.Errorf(codes.Unimplemented, "method PrimeNumberDecomposition not implemented")
}
func (*UnimplementedCalcServiceServer) ComputeAverage(CalcService_ComputeAverageServer) error {
	return status.Errorf(codes.Unimplemented, "method ComputeAverage not implemented")
}
func (*UnimplementedCalcServiceServer) FindMaximum(CalcService_FindMaximumServer) error {
	return status.Errorf(codes.Unimplemented, "method FindMaximum not implemented")
}

func RegisterCalcServiceServer(s *grpc.Server, srv CalcServiceServer) {
	s.RegisterService(&_CalcService_serviceDesc, srv)
}

func _CalcService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calcpb.CalcService/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServiceServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalcService_PrimeNumberDecomposition_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeNumberDecompositionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalcServiceServer).PrimeNumberDecomposition(m, &calcServicePrimeNumberDecompositionServer{stream})
}

type CalcService_PrimeNumberDecompositionServer interface {
	Send(*PrimeNumberDecompositionResponse) error
	grpc.ServerStream
}

type calcServicePrimeNumberDecompositionServer struct {
	grpc.ServerStream
}

func (x *calcServicePrimeNumberDecompositionServer) Send(m *PrimeNumberDecompositionResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CalcService_ComputeAverage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalcServiceServer).ComputeAverage(&calcServiceComputeAverageServer{stream})
}

type CalcService_ComputeAverageServer interface {
	SendAndClose(*ComputeAverageResponse) error
	Recv() (*ComputeAverageRequest, error)
	grpc.ServerStream
}

type calcServiceComputeAverageServer struct {
	grpc.ServerStream
}

func (x *calcServiceComputeAverageServer) SendAndClose(m *ComputeAverageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calcServiceComputeAverageServer) Recv() (*ComputeAverageRequest, error) {
	m := new(ComputeAverageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalcService_FindMaximum_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalcServiceServer).FindMaximum(&calcServiceFindMaximumServer{stream})
}

type CalcService_FindMaximumServer interface {
	Send(*FindMaximumResponse) error
	Recv() (*FindMaximumRequest, error)
	grpc.ServerStream
}

type calcServiceFindMaximumServer struct {
	grpc.ServerStream
}

func (x *calcServiceFindMaximumServer) Send(m *FindMaximumResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calcServiceFindMaximumServer) Recv() (*FindMaximumRequest, error) {
	m := new(FindMaximumRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _CalcService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calcpb.CalcService",
	HandlerType: (*CalcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _CalcService_Sum_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PrimeNumberDecomposition",
			Handler:       _CalcService_PrimeNumberDecomposition_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ComputeAverage",
			Handler:       _CalcService_ComputeAverage_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "FindMaximum",
			Handler:       _CalcService_FindMaximum_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "calculator/calcpb/calc.proto",
}
