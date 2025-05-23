// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v6.30.0--rc1
// source: api/category/v1/category.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateCategoryRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateCategoryRequest) Reset() {
	*x = CreateCategoryRequest{}
	mi := &file_api_category_v1_category_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCategoryRequest) ProtoMessage() {}

func (x *CreateCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_category_v1_category_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCategoryRequest.ProtoReflect.Descriptor instead.
func (*CreateCategoryRequest) Descriptor() ([]byte, []int) {
	return file_api_category_v1_category_proto_rawDescGZIP(), []int{0}
}

type CreateCategoryReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateCategoryReply) Reset() {
	*x = CreateCategoryReply{}
	mi := &file_api_category_v1_category_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCategoryReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCategoryReply) ProtoMessage() {}

func (x *CreateCategoryReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_category_v1_category_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCategoryReply.ProtoReflect.Descriptor instead.
func (*CreateCategoryReply) Descriptor() ([]byte, []int) {
	return file_api_category_v1_category_proto_rawDescGZIP(), []int{1}
}

type UpdateCategoryRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateCategoryRequest) Reset() {
	*x = UpdateCategoryRequest{}
	mi := &file_api_category_v1_category_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCategoryRequest) ProtoMessage() {}

func (x *UpdateCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_category_v1_category_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCategoryRequest.ProtoReflect.Descriptor instead.
func (*UpdateCategoryRequest) Descriptor() ([]byte, []int) {
	return file_api_category_v1_category_proto_rawDescGZIP(), []int{2}
}

type UpdateCategoryReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateCategoryReply) Reset() {
	*x = UpdateCategoryReply{}
	mi := &file_api_category_v1_category_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCategoryReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCategoryReply) ProtoMessage() {}

func (x *UpdateCategoryReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_category_v1_category_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCategoryReply.ProtoReflect.Descriptor instead.
func (*UpdateCategoryReply) Descriptor() ([]byte, []int) {
	return file_api_category_v1_category_proto_rawDescGZIP(), []int{3}
}

type DeleteCategoryRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteCategoryRequest) Reset() {
	*x = DeleteCategoryRequest{}
	mi := &file_api_category_v1_category_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCategoryRequest) ProtoMessage() {}

func (x *DeleteCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_category_v1_category_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCategoryRequest.ProtoReflect.Descriptor instead.
func (*DeleteCategoryRequest) Descriptor() ([]byte, []int) {
	return file_api_category_v1_category_proto_rawDescGZIP(), []int{4}
}

type DeleteCategoryReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteCategoryReply) Reset() {
	*x = DeleteCategoryReply{}
	mi := &file_api_category_v1_category_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteCategoryReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCategoryReply) ProtoMessage() {}

func (x *DeleteCategoryReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_category_v1_category_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCategoryReply.ProtoReflect.Descriptor instead.
func (*DeleteCategoryReply) Descriptor() ([]byte, []int) {
	return file_api_category_v1_category_proto_rawDescGZIP(), []int{5}
}

type GetCategoryRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCategoryRequest) Reset() {
	*x = GetCategoryRequest{}
	mi := &file_api_category_v1_category_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCategoryRequest) ProtoMessage() {}

func (x *GetCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_category_v1_category_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCategoryRequest.ProtoReflect.Descriptor instead.
func (*GetCategoryRequest) Descriptor() ([]byte, []int) {
	return file_api_category_v1_category_proto_rawDescGZIP(), []int{6}
}

type GetCategoryReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCategoryReply) Reset() {
	*x = GetCategoryReply{}
	mi := &file_api_category_v1_category_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCategoryReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCategoryReply) ProtoMessage() {}

func (x *GetCategoryReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_category_v1_category_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCategoryReply.ProtoReflect.Descriptor instead.
func (*GetCategoryReply) Descriptor() ([]byte, []int) {
	return file_api_category_v1_category_proto_rawDescGZIP(), []int{7}
}

type CategoryWrapper struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CategoryWrapper) Reset() {
	*x = CategoryWrapper{}
	mi := &file_api_category_v1_category_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CategoryWrapper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryWrapper) ProtoMessage() {}

func (x *CategoryWrapper) ProtoReflect() protoreflect.Message {
	mi := &file_api_category_v1_category_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryWrapper.ProtoReflect.Descriptor instead.
func (*CategoryWrapper) Descriptor() ([]byte, []int) {
	return file_api_category_v1_category_proto_rawDescGZIP(), []int{8}
}

func (x *CategoryWrapper) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CategoryWrapper) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListCategoryRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListCategoryRequest) Reset() {
	*x = ListCategoryRequest{}
	mi := &file_api_category_v1_category_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCategoryRequest) ProtoMessage() {}

func (x *ListCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_category_v1_category_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCategoryRequest.ProtoReflect.Descriptor instead.
func (*ListCategoryRequest) Descriptor() ([]byte, []int) {
	return file_api_category_v1_category_proto_rawDescGZIP(), []int{9}
}

type ListCategoryReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Category      []*CategoryWrapper     `protobuf:"bytes,1,rep,name=category,proto3" json:"category,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListCategoryReply) Reset() {
	*x = ListCategoryReply{}
	mi := &file_api_category_v1_category_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCategoryReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCategoryReply) ProtoMessage() {}

func (x *ListCategoryReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_category_v1_category_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCategoryReply.ProtoReflect.Descriptor instead.
func (*ListCategoryReply) Descriptor() ([]byte, []int) {
	return file_api_category_v1_category_proto_rawDescGZIP(), []int{10}
}

func (x *ListCategoryReply) GetCategory() []*CategoryWrapper {
	if x != nil {
		return x.Category
	}
	return nil
}

var File_api_category_v1_category_proto protoreflect.FileDescriptor

var file_api_category_v1_category_proto_rawDesc = string([]byte{
	0x0a, 0x1e, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x76,
	0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x17, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x15, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x17, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x15, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x17, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x15, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x35, 0x0a, 0x0f, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x15, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x51, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3c, 0x0a, 0x08,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x32, 0xfc, 0x03, 0x0a, 0x08, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x5e, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x5e, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x5e, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x55, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x79,
	0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x24,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19,
	0x12, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x75, 0x66, 0x66, 0x2f, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x2d, 0x0a, 0x0f, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x18,
	0x73, 0x74, 0x75, 0x66, 0x66, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_api_category_v1_category_proto_rawDescOnce sync.Once
	file_api_category_v1_category_proto_rawDescData []byte
)

func file_api_category_v1_category_proto_rawDescGZIP() []byte {
	file_api_category_v1_category_proto_rawDescOnce.Do(func() {
		file_api_category_v1_category_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_category_v1_category_proto_rawDesc), len(file_api_category_v1_category_proto_rawDesc)))
	})
	return file_api_category_v1_category_proto_rawDescData
}

var file_api_category_v1_category_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_category_v1_category_proto_goTypes = []any{
	(*CreateCategoryRequest)(nil), // 0: api.category.v1.CreateCategoryRequest
	(*CreateCategoryReply)(nil),   // 1: api.category.v1.CreateCategoryReply
	(*UpdateCategoryRequest)(nil), // 2: api.category.v1.UpdateCategoryRequest
	(*UpdateCategoryReply)(nil),   // 3: api.category.v1.UpdateCategoryReply
	(*DeleteCategoryRequest)(nil), // 4: api.category.v1.DeleteCategoryRequest
	(*DeleteCategoryReply)(nil),   // 5: api.category.v1.DeleteCategoryReply
	(*GetCategoryRequest)(nil),    // 6: api.category.v1.GetCategoryRequest
	(*GetCategoryReply)(nil),      // 7: api.category.v1.GetCategoryReply
	(*CategoryWrapper)(nil),       // 8: api.category.v1.CategoryWrapper
	(*ListCategoryRequest)(nil),   // 9: api.category.v1.ListCategoryRequest
	(*ListCategoryReply)(nil),     // 10: api.category.v1.ListCategoryReply
}
var file_api_category_v1_category_proto_depIdxs = []int32{
	8,  // 0: api.category.v1.ListCategoryReply.category:type_name -> api.category.v1.CategoryWrapper
	0,  // 1: api.category.v1.Category.CreateCategory:input_type -> api.category.v1.CreateCategoryRequest
	2,  // 2: api.category.v1.Category.UpdateCategory:input_type -> api.category.v1.UpdateCategoryRequest
	4,  // 3: api.category.v1.Category.DeleteCategory:input_type -> api.category.v1.DeleteCategoryRequest
	6,  // 4: api.category.v1.Category.GetCategory:input_type -> api.category.v1.GetCategoryRequest
	9,  // 5: api.category.v1.Category.ListCategory:input_type -> api.category.v1.ListCategoryRequest
	1,  // 6: api.category.v1.Category.CreateCategory:output_type -> api.category.v1.CreateCategoryReply
	3,  // 7: api.category.v1.Category.UpdateCategory:output_type -> api.category.v1.UpdateCategoryReply
	5,  // 8: api.category.v1.Category.DeleteCategory:output_type -> api.category.v1.DeleteCategoryReply
	7,  // 9: api.category.v1.Category.GetCategory:output_type -> api.category.v1.GetCategoryReply
	10, // 10: api.category.v1.Category.ListCategory:output_type -> api.category.v1.ListCategoryReply
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_api_category_v1_category_proto_init() }
func file_api_category_v1_category_proto_init() {
	if File_api_category_v1_category_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_category_v1_category_proto_rawDesc), len(file_api_category_v1_category_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_category_v1_category_proto_goTypes,
		DependencyIndexes: file_api_category_v1_category_proto_depIdxs,
		MessageInfos:      file_api_category_v1_category_proto_msgTypes,
	}.Build()
	File_api_category_v1_category_proto = out.File
	file_api_category_v1_category_proto_goTypes = nil
	file_api_category_v1_category_proto_depIdxs = nil
}
