// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v6.30.0--rc1
// source: api/stuff/v1/stuff.proto

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

type CreateStuffRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Category      int64                  `protobuf:"varint,2,opt,name=category,proto3" json:"category,omitempty"`
	Price         float32                `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	Photos        string                 `protobuf:"bytes,4,opt,name=photos,proto3" json:"photos,omitempty"`
	Publisher     int64                  `protobuf:"varint,5,opt,name=publisher,proto3" json:"publisher,omitempty"`
	Condition     string                 `protobuf:"bytes,6,opt,name=condition,proto3" json:"condition,omitempty"`
	Description   string                 `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateStuffRequest) Reset() {
	*x = CreateStuffRequest{}
	mi := &file_api_stuff_v1_stuff_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateStuffRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStuffRequest) ProtoMessage() {}

func (x *CreateStuffRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_stuff_v1_stuff_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStuffRequest.ProtoReflect.Descriptor instead.
func (*CreateStuffRequest) Descriptor() ([]byte, []int) {
	return file_api_stuff_v1_stuff_proto_rawDescGZIP(), []int{0}
}

func (x *CreateStuffRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateStuffRequest) GetCategory() int64 {
	if x != nil {
		return x.Category
	}
	return 0
}

func (x *CreateStuffRequest) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateStuffRequest) GetPhotos() string {
	if x != nil {
		return x.Photos
	}
	return ""
}

func (x *CreateStuffRequest) GetPublisher() int64 {
	if x != nil {
		return x.Publisher
	}
	return 0
}

func (x *CreateStuffRequest) GetCondition() string {
	if x != nil {
		return x.Condition
	}
	return ""
}

func (x *CreateStuffRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type CreateStuffReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateStuffReply) Reset() {
	*x = CreateStuffReply{}
	mi := &file_api_stuff_v1_stuff_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateStuffReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStuffReply) ProtoMessage() {}

func (x *CreateStuffReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_stuff_v1_stuff_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStuffReply.ProtoReflect.Descriptor instead.
func (*CreateStuffReply) Descriptor() ([]byte, []int) {
	return file_api_stuff_v1_stuff_proto_rawDescGZIP(), []int{1}
}

type UpdateStuffRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateStuffRequest) Reset() {
	*x = UpdateStuffRequest{}
	mi := &file_api_stuff_v1_stuff_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateStuffRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStuffRequest) ProtoMessage() {}

func (x *UpdateStuffRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_stuff_v1_stuff_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStuffRequest.ProtoReflect.Descriptor instead.
func (*UpdateStuffRequest) Descriptor() ([]byte, []int) {
	return file_api_stuff_v1_stuff_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateStuffRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateStuffRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UpdateStuffReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateStuffReply) Reset() {
	*x = UpdateStuffReply{}
	mi := &file_api_stuff_v1_stuff_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateStuffReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStuffReply) ProtoMessage() {}

func (x *UpdateStuffReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_stuff_v1_stuff_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStuffReply.ProtoReflect.Descriptor instead.
func (*UpdateStuffReply) Descriptor() ([]byte, []int) {
	return file_api_stuff_v1_stuff_proto_rawDescGZIP(), []int{3}
}

type DeleteStuffRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteStuffRequest) Reset() {
	*x = DeleteStuffRequest{}
	mi := &file_api_stuff_v1_stuff_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteStuffRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteStuffRequest) ProtoMessage() {}

func (x *DeleteStuffRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_stuff_v1_stuff_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteStuffRequest.ProtoReflect.Descriptor instead.
func (*DeleteStuffRequest) Descriptor() ([]byte, []int) {
	return file_api_stuff_v1_stuff_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteStuffRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteStuffReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteStuffReply) Reset() {
	*x = DeleteStuffReply{}
	mi := &file_api_stuff_v1_stuff_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteStuffReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteStuffReply) ProtoMessage() {}

func (x *DeleteStuffReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_stuff_v1_stuff_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteStuffReply.ProtoReflect.Descriptor instead.
func (*DeleteStuffReply) Descriptor() ([]byte, []int) {
	return file_api_stuff_v1_stuff_proto_rawDescGZIP(), []int{5}
}

type GetStuffRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetStuffRequest) Reset() {
	*x = GetStuffRequest{}
	mi := &file_api_stuff_v1_stuff_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetStuffRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStuffRequest) ProtoMessage() {}

func (x *GetStuffRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_stuff_v1_stuff_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStuffRequest.ProtoReflect.Descriptor instead.
func (*GetStuffRequest) Descriptor() ([]byte, []int) {
	return file_api_stuff_v1_stuff_proto_rawDescGZIP(), []int{6}
}

func (x *GetStuffRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetStuffReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetStuffReply) Reset() {
	*x = GetStuffReply{}
	mi := &file_api_stuff_v1_stuff_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetStuffReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStuffReply) ProtoMessage() {}

func (x *GetStuffReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_stuff_v1_stuff_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStuffReply.ProtoReflect.Descriptor instead.
func (*GetStuffReply) Descriptor() ([]byte, []int) {
	return file_api_stuff_v1_stuff_proto_rawDescGZIP(), []int{7}
}

func (x *GetStuffReply) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetStuffReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListStuffRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListStuffRequest) Reset() {
	*x = ListStuffRequest{}
	mi := &file_api_stuff_v1_stuff_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListStuffRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStuffRequest) ProtoMessage() {}

func (x *ListStuffRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_stuff_v1_stuff_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStuffRequest.ProtoReflect.Descriptor instead.
func (*ListStuffRequest) Descriptor() ([]byte, []int) {
	return file_api_stuff_v1_stuff_proto_rawDescGZIP(), []int{8}
}

type ListStuffReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListStuffReply) Reset() {
	*x = ListStuffReply{}
	mi := &file_api_stuff_v1_stuff_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListStuffReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStuffReply) ProtoMessage() {}

func (x *ListStuffReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_stuff_v1_stuff_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStuffReply.ProtoReflect.Descriptor instead.
func (*ListStuffReply) Descriptor() ([]byte, []int) {
	return file_api_stuff_v1_stuff_proto_rawDescGZIP(), []int{9}
}

var File_api_stuff_v1_stuff_proto protoreflect.FileDescriptor

var file_api_stuff_v1_stuff_proto_rawDesc = string([]byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x75, 0x66, 0x66, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x74, 0x75, 0x66, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x74, 0x75, 0x66, 0x66, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd0, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x74, 0x75, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x12, 0x0a, 0x10, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x74, 0x75, 0x66, 0x66, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x38, 0x0a,
	0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x75, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x74, 0x75, 0x66, 0x66, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x24, 0x0a, 0x12, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x75, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x12, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x75, 0x66, 0x66,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x21, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x66,
	0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x33, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53,
	0x74, 0x75, 0x66, 0x66, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x12, 0x0a,
	0x10, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x75, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x10, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x75, 0x66, 0x66, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x32, 0x85, 0x04, 0x0a, 0x05, 0x53, 0x74, 0x75, 0x66, 0x66, 0x12, 0x6c, 0x0a,
	0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x75, 0x66, 0x66, 0x12, 0x20, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x74, 0x75, 0x66, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x74, 0x75, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x75, 0x66, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x75, 0x66, 0x66, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1b,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x74, 0x75, 0x66, 0x66, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x67, 0x0a, 0x0b, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x75, 0x66, 0x66, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x74, 0x75, 0x66, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x74, 0x75, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x74, 0x75, 0x66, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x74, 0x75, 0x66, 0x66, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x16, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x10, 0x1a, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x75, 0x66, 0x66, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0x12, 0x67, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74,
	0x75, 0x66, 0x66, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x75, 0x66, 0x66, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x75, 0x66, 0x66, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x75, 0x66,
	0x66, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x75, 0x66, 0x66,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x2a, 0x0e, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x74, 0x75, 0x66, 0x66, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x5e, 0x0a,
	0x08, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x66, 0x66, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x74, 0x75, 0x66, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x66,
	0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x74, 0x75, 0x66, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x66, 0x66,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x74, 0x75, 0x66, 0x66, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x5c, 0x0a,
	0x09, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x75, 0x66, 0x66, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x74, 0x75, 0x66, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74,
	0x75, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x74, 0x75, 0x66, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74,
	0x75, 0x66, 0x66, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b,
	0x12, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x75, 0x66, 0x66, 0x42, 0x27, 0x0a, 0x0c, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x74, 0x75, 0x66, 0x66, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x15, 0x73,
	0x74, 0x75, 0x66, 0x66, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x75, 0x66, 0x66, 0x2f, 0x76,
	0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_api_stuff_v1_stuff_proto_rawDescOnce sync.Once
	file_api_stuff_v1_stuff_proto_rawDescData []byte
)

func file_api_stuff_v1_stuff_proto_rawDescGZIP() []byte {
	file_api_stuff_v1_stuff_proto_rawDescOnce.Do(func() {
		file_api_stuff_v1_stuff_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_stuff_v1_stuff_proto_rawDesc), len(file_api_stuff_v1_stuff_proto_rawDesc)))
	})
	return file_api_stuff_v1_stuff_proto_rawDescData
}

var file_api_stuff_v1_stuff_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_stuff_v1_stuff_proto_goTypes = []any{
	(*CreateStuffRequest)(nil), // 0: api.stuff.v1.CreateStuffRequest
	(*CreateStuffReply)(nil),   // 1: api.stuff.v1.CreateStuffReply
	(*UpdateStuffRequest)(nil), // 2: api.stuff.v1.UpdateStuffRequest
	(*UpdateStuffReply)(nil),   // 3: api.stuff.v1.UpdateStuffReply
	(*DeleteStuffRequest)(nil), // 4: api.stuff.v1.DeleteStuffRequest
	(*DeleteStuffReply)(nil),   // 5: api.stuff.v1.DeleteStuffReply
	(*GetStuffRequest)(nil),    // 6: api.stuff.v1.GetStuffRequest
	(*GetStuffReply)(nil),      // 7: api.stuff.v1.GetStuffReply
	(*ListStuffRequest)(nil),   // 8: api.stuff.v1.ListStuffRequest
	(*ListStuffReply)(nil),     // 9: api.stuff.v1.ListStuffReply
}
var file_api_stuff_v1_stuff_proto_depIdxs = []int32{
	0, // 0: api.stuff.v1.Stuff.CreateStuff:input_type -> api.stuff.v1.CreateStuffRequest
	2, // 1: api.stuff.v1.Stuff.UpdateStuff:input_type -> api.stuff.v1.UpdateStuffRequest
	4, // 2: api.stuff.v1.Stuff.DeleteStuff:input_type -> api.stuff.v1.DeleteStuffRequest
	6, // 3: api.stuff.v1.Stuff.GetStuff:input_type -> api.stuff.v1.GetStuffRequest
	8, // 4: api.stuff.v1.Stuff.ListStuff:input_type -> api.stuff.v1.ListStuffRequest
	1, // 5: api.stuff.v1.Stuff.CreateStuff:output_type -> api.stuff.v1.CreateStuffReply
	3, // 6: api.stuff.v1.Stuff.UpdateStuff:output_type -> api.stuff.v1.UpdateStuffReply
	5, // 7: api.stuff.v1.Stuff.DeleteStuff:output_type -> api.stuff.v1.DeleteStuffReply
	7, // 8: api.stuff.v1.Stuff.GetStuff:output_type -> api.stuff.v1.GetStuffReply
	9, // 9: api.stuff.v1.Stuff.ListStuff:output_type -> api.stuff.v1.ListStuffReply
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_stuff_v1_stuff_proto_init() }
func file_api_stuff_v1_stuff_proto_init() {
	if File_api_stuff_v1_stuff_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_stuff_v1_stuff_proto_rawDesc), len(file_api_stuff_v1_stuff_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_stuff_v1_stuff_proto_goTypes,
		DependencyIndexes: file_api_stuff_v1_stuff_proto_depIdxs,
		MessageInfos:      file_api_stuff_v1_stuff_proto_msgTypes,
	}.Build()
	File_api_stuff_v1_stuff_proto = out.File
	file_api_stuff_v1_stuff_proto_goTypes = nil
	file_api_stuff_v1_stuff_proto_depIdxs = nil
}
