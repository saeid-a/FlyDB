// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: lib/proto/slave.proto

package proto

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

type SlaveGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *SlaveGetRequest) Reset() {
	*x = SlaveGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_proto_slave_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlaveGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlaveGetRequest) ProtoMessage() {}

func (x *SlaveGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lib_proto_slave_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlaveGetRequest.ProtoReflect.Descriptor instead.
func (*SlaveGetRequest) Descriptor() ([]byte, []int) {
	return file_lib_proto_slave_proto_rawDescGZIP(), []int{0}
}

func (x *SlaveGetRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type SlaveGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SlaveGetResponse) Reset() {
	*x = SlaveGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_proto_slave_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlaveGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlaveGetResponse) ProtoMessage() {}

func (x *SlaveGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lib_proto_slave_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlaveGetResponse.ProtoReflect.Descriptor instead.
func (*SlaveGetResponse) Descriptor() ([]byte, []int) {
	return file_lib_proto_slave_proto_rawDescGZIP(), []int{1}
}

func (x *SlaveGetResponse) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type SlaveSetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SlaveSetRequest) Reset() {
	*x = SlaveSetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_proto_slave_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlaveSetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlaveSetRequest) ProtoMessage() {}

func (x *SlaveSetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lib_proto_slave_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlaveSetRequest.ProtoReflect.Descriptor instead.
func (*SlaveSetRequest) Descriptor() ([]byte, []int) {
	return file_lib_proto_slave_proto_rawDescGZIP(), []int{2}
}

func (x *SlaveSetRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SlaveSetRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type SlaveSetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *SlaveSetResponse) Reset() {
	*x = SlaveSetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_proto_slave_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlaveSetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlaveSetResponse) ProtoMessage() {}

func (x *SlaveSetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lib_proto_slave_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlaveSetResponse.ProtoReflect.Descriptor instead.
func (*SlaveSetResponse) Descriptor() ([]byte, []int) {
	return file_lib_proto_slave_proto_rawDescGZIP(), []int{3}
}

func (x *SlaveSetResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type SlaveDelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *SlaveDelRequest) Reset() {
	*x = SlaveDelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_proto_slave_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlaveDelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlaveDelRequest) ProtoMessage() {}

func (x *SlaveDelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lib_proto_slave_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlaveDelRequest.ProtoReflect.Descriptor instead.
func (*SlaveDelRequest) Descriptor() ([]byte, []int) {
	return file_lib_proto_slave_proto_rawDescGZIP(), []int{4}
}

func (x *SlaveDelRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type SlaveDelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *SlaveDelResponse) Reset() {
	*x = SlaveDelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_proto_slave_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlaveDelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlaveDelResponse) ProtoMessage() {}

func (x *SlaveDelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lib_proto_slave_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlaveDelResponse.ProtoReflect.Descriptor instead.
func (*SlaveDelResponse) Descriptor() ([]byte, []int) {
	return file_lib_proto_slave_proto_rawDescGZIP(), []int{5}
}

func (x *SlaveDelResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type SlaveKeysRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pattern string `protobuf:"bytes,1,opt,name=pattern,proto3" json:"pattern,omitempty"`
}

func (x *SlaveKeysRequest) Reset() {
	*x = SlaveKeysRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_proto_slave_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlaveKeysRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlaveKeysRequest) ProtoMessage() {}

func (x *SlaveKeysRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lib_proto_slave_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlaveKeysRequest.ProtoReflect.Descriptor instead.
func (*SlaveKeysRequest) Descriptor() ([]byte, []int) {
	return file_lib_proto_slave_proto_rawDescGZIP(), []int{6}
}

func (x *SlaveKeysRequest) GetPattern() string {
	if x != nil {
		return x.Pattern
	}
	return ""
}

type SlaveKeysResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys []string `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *SlaveKeysResponse) Reset() {
	*x = SlaveKeysResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_proto_slave_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlaveKeysResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlaveKeysResponse) ProtoMessage() {}

func (x *SlaveKeysResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lib_proto_slave_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlaveKeysResponse.ProtoReflect.Descriptor instead.
func (*SlaveKeysResponse) Descriptor() ([]byte, []int) {
	return file_lib_proto_slave_proto_rawDescGZIP(), []int{7}
}

func (x *SlaveKeysResponse) GetKeys() []string {
	if x != nil {
		return x.Keys
	}
	return nil
}

type SlaveExistsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *SlaveExistsRequest) Reset() {
	*x = SlaveExistsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_proto_slave_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlaveExistsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlaveExistsRequest) ProtoMessage() {}

func (x *SlaveExistsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lib_proto_slave_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlaveExistsRequest.ProtoReflect.Descriptor instead.
func (*SlaveExistsRequest) Descriptor() ([]byte, []int) {
	return file_lib_proto_slave_proto_rawDescGZIP(), []int{8}
}

func (x *SlaveExistsRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type SlaveExistsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exists bool `protobuf:"varint,1,opt,name=exists,proto3" json:"exists,omitempty"`
}

func (x *SlaveExistsResponse) Reset() {
	*x = SlaveExistsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_proto_slave_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlaveExistsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlaveExistsResponse) ProtoMessage() {}

func (x *SlaveExistsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lib_proto_slave_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlaveExistsResponse.ProtoReflect.Descriptor instead.
func (*SlaveExistsResponse) Descriptor() ([]byte, []int) {
	return file_lib_proto_slave_proto_rawDescGZIP(), []int{9}
}

func (x *SlaveExistsResponse) GetExists() bool {
	if x != nil {
		return x.Exists
	}
	return false
}

type SlaveExpireRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Ttl int64  `protobuf:"varint,2,opt,name=ttl,proto3" json:"ttl,omitempty"`
}

func (x *SlaveExpireRequest) Reset() {
	*x = SlaveExpireRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_proto_slave_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlaveExpireRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlaveExpireRequest) ProtoMessage() {}

func (x *SlaveExpireRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lib_proto_slave_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlaveExpireRequest.ProtoReflect.Descriptor instead.
func (*SlaveExpireRequest) Descriptor() ([]byte, []int) {
	return file_lib_proto_slave_proto_rawDescGZIP(), []int{10}
}

func (x *SlaveExpireRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SlaveExpireRequest) GetTtl() int64 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

type SlaveExpireResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *SlaveExpireResponse) Reset() {
	*x = SlaveExpireResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_proto_slave_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlaveExpireResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlaveExpireResponse) ProtoMessage() {}

func (x *SlaveExpireResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lib_proto_slave_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlaveExpireResponse.ProtoReflect.Descriptor instead.
func (*SlaveExpireResponse) Descriptor() ([]byte, []int) {
	return file_lib_proto_slave_proto_rawDescGZIP(), []int{11}
}

func (x *SlaveExpireResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type SlaveTTLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *SlaveTTLRequest) Reset() {
	*x = SlaveTTLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_proto_slave_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlaveTTLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlaveTTLRequest) ProtoMessage() {}

func (x *SlaveTTLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lib_proto_slave_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlaveTTLRequest.ProtoReflect.Descriptor instead.
func (*SlaveTTLRequest) Descriptor() ([]byte, []int) {
	return file_lib_proto_slave_proto_rawDescGZIP(), []int{12}
}

func (x *SlaveTTLRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type SlaveTTLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ttl int64 `protobuf:"varint,1,opt,name=ttl,proto3" json:"ttl,omitempty"`
}

func (x *SlaveTTLResponse) Reset() {
	*x = SlaveTTLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_proto_slave_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlaveTTLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlaveTTLResponse) ProtoMessage() {}

func (x *SlaveTTLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lib_proto_slave_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlaveTTLResponse.ProtoReflect.Descriptor instead.
func (*SlaveTTLResponse) Descriptor() ([]byte, []int) {
	return file_lib_proto_slave_proto_rawDescGZIP(), []int{13}
}

func (x *SlaveTTLResponse) GetTtl() int64 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

type SlaveHeartbeatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SlaveHeartbeatRequest) Reset() {
	*x = SlaveHeartbeatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_proto_slave_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlaveHeartbeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlaveHeartbeatRequest) ProtoMessage() {}

func (x *SlaveHeartbeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lib_proto_slave_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlaveHeartbeatRequest.ProtoReflect.Descriptor instead.
func (*SlaveHeartbeatRequest) Descriptor() ([]byte, []int) {
	return file_lib_proto_slave_proto_rawDescGZIP(), []int{14}
}

func (x *SlaveHeartbeatRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SlaveHeartbeatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *SlaveHeartbeatResponse) Reset() {
	*x = SlaveHeartbeatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_proto_slave_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlaveHeartbeatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlaveHeartbeatResponse) ProtoMessage() {}

func (x *SlaveHeartbeatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lib_proto_slave_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlaveHeartbeatResponse.ProtoReflect.Descriptor instead.
func (*SlaveHeartbeatResponse) Descriptor() ([]byte, []int) {
	return file_lib_proto_slave_proto_rawDescGZIP(), []int{15}
}

func (x *SlaveHeartbeatResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

var File_lib_proto_slave_proto protoreflect.FileDescriptor

var file_lib_proto_slave_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6c, 0x69, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x6c, 0x61, 0x76,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x22, 0x23, 0x0a, 0x0f, 0x53, 0x6c, 0x61, 0x76, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x28, 0x0a, 0x10, 0x53, 0x6c, 0x61, 0x76, 0x65, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x39, 0x0a, 0x0f, 0x53, 0x6c, 0x61, 0x76, 0x65, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x22, 0x0a, 0x10, 0x53, 0x6c,
	0x61, 0x76, 0x65, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x23,
	0x0a, 0x0f, 0x53, 0x6c, 0x61, 0x76, 0x65, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x22, 0x22, 0x0a, 0x10, 0x53, 0x6c, 0x61, 0x76, 0x65, 0x44, 0x65, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x2c, 0x0a, 0x10, 0x53, 0x6c, 0x61, 0x76, 0x65,
	0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61,
	0x74, 0x74, 0x65, 0x72, 0x6e, 0x22, 0x27, 0x0a, 0x11, 0x53, 0x6c, 0x61, 0x76, 0x65, 0x4b, 0x65,
	0x79, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x65,
	0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x22, 0x26,
	0x0a, 0x12, 0x53, 0x6c, 0x61, 0x76, 0x65, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x2d, 0x0a, 0x13, 0x53, 0x6c, 0x61, 0x76, 0x65, 0x45,
	0x78, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65,
	0x78, 0x69, 0x73, 0x74, 0x73, 0x22, 0x38, 0x0a, 0x12, 0x53, 0x6c, 0x61, 0x76, 0x65, 0x45, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x74, 0x74, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x22,
	0x25, 0x0a, 0x13, 0x53, 0x6c, 0x61, 0x76, 0x65, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x23, 0x0a, 0x0f, 0x53, 0x6c, 0x61, 0x76, 0x65, 0x54,
	0x54, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x24, 0x0a, 0x10, 0x53,
	0x6c, 0x61, 0x76, 0x65, 0x54, 0x54, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x74, 0x74,
	0x6c, 0x22, 0x27, 0x0a, 0x15, 0x53, 0x6c, 0x61, 0x76, 0x65, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62,
	0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x28, 0x0a, 0x16, 0x53, 0x6c,
	0x61, 0x76, 0x65, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x02, 0x6f, 0x6b, 0x32, 0xa9, 0x04, 0x0a, 0x10, 0x53, 0x6c, 0x61, 0x76, 0x65, 0x47, 0x72,
	0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x03, 0x47, 0x65, 0x74,
	0x12, 0x18, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x6c, 0x61, 0x76, 0x65,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x6c, 0x61, 0x76, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x03, 0x53, 0x65, 0x74, 0x12, 0x18,
	0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x6c, 0x61, 0x76, 0x65, 0x53, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x2e, 0x53, 0x6c, 0x61, 0x76, 0x65, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x03, 0x44, 0x65, 0x6c, 0x12, 0x18, 0x2e, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x6c, 0x61, 0x76, 0x65, 0x44, 0x65, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x53, 0x6c, 0x61, 0x76, 0x65, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x04, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x19, 0x2e, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x6c, 0x61, 0x76, 0x65, 0x4b, 0x65, 0x79, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x53, 0x6c, 0x61, 0x76, 0x65, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x06, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x12, 0x1b,
	0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x6c, 0x61, 0x76, 0x65, 0x45, 0x78,
	0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x6c, 0x61, 0x76, 0x65, 0x45, 0x78, 0x69, 0x73, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x06, 0x45,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x12, 0x1b, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e,
	0x53, 0x6c, 0x61, 0x76, 0x65, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x6c, 0x61,
	0x76, 0x65, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x3c, 0x0a, 0x03, 0x54, 0x54, 0x4c, 0x12, 0x18, 0x2e, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x2e, 0x53, 0x6c, 0x61, 0x76, 0x65, 0x54, 0x54, 0x4c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x6c,
	0x61, 0x76, 0x65, 0x54, 0x54, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x4e, 0x0a, 0x09, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x1e, 0x2e,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x6c, 0x61, 0x76, 0x65, 0x48, 0x65, 0x61,
	0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x6c, 0x61, 0x76, 0x65, 0x48, 0x65, 0x61,
	0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x11, 0x5a, 0x0f, 0x66, 0x6c, 0x79, 0x64, 0x62, 0x2f, 0x6c, 0x69, 0x62, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lib_proto_slave_proto_rawDescOnce sync.Once
	file_lib_proto_slave_proto_rawDescData = file_lib_proto_slave_proto_rawDesc
)

func file_lib_proto_slave_proto_rawDescGZIP() []byte {
	file_lib_proto_slave_proto_rawDescOnce.Do(func() {
		file_lib_proto_slave_proto_rawDescData = protoimpl.X.CompressGZIP(file_lib_proto_slave_proto_rawDescData)
	})
	return file_lib_proto_slave_proto_rawDescData
}

var file_lib_proto_slave_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_lib_proto_slave_proto_goTypes = []interface{}{
	(*SlaveGetRequest)(nil),        // 0: cluster.SlaveGetRequest
	(*SlaveGetResponse)(nil),       // 1: cluster.SlaveGetResponse
	(*SlaveSetRequest)(nil),        // 2: cluster.SlaveSetRequest
	(*SlaveSetResponse)(nil),       // 3: cluster.SlaveSetResponse
	(*SlaveDelRequest)(nil),        // 4: cluster.SlaveDelRequest
	(*SlaveDelResponse)(nil),       // 5: cluster.SlaveDelResponse
	(*SlaveKeysRequest)(nil),       // 6: cluster.SlaveKeysRequest
	(*SlaveKeysResponse)(nil),      // 7: cluster.SlaveKeysResponse
	(*SlaveExistsRequest)(nil),     // 8: cluster.SlaveExistsRequest
	(*SlaveExistsResponse)(nil),    // 9: cluster.SlaveExistsResponse
	(*SlaveExpireRequest)(nil),     // 10: cluster.SlaveExpireRequest
	(*SlaveExpireResponse)(nil),    // 11: cluster.SlaveExpireResponse
	(*SlaveTTLRequest)(nil),        // 12: cluster.SlaveTTLRequest
	(*SlaveTTLResponse)(nil),       // 13: cluster.SlaveTTLResponse
	(*SlaveHeartbeatRequest)(nil),  // 14: cluster.SlaveHeartbeatRequest
	(*SlaveHeartbeatResponse)(nil), // 15: cluster.SlaveHeartbeatResponse
}
var file_lib_proto_slave_proto_depIdxs = []int32{
	0,  // 0: cluster.SlaveGrpcService.Get:input_type -> cluster.SlaveGetRequest
	2,  // 1: cluster.SlaveGrpcService.Set:input_type -> cluster.SlaveSetRequest
	4,  // 2: cluster.SlaveGrpcService.Del:input_type -> cluster.SlaveDelRequest
	6,  // 3: cluster.SlaveGrpcService.Keys:input_type -> cluster.SlaveKeysRequest
	8,  // 4: cluster.SlaveGrpcService.Exists:input_type -> cluster.SlaveExistsRequest
	10, // 5: cluster.SlaveGrpcService.Expire:input_type -> cluster.SlaveExpireRequest
	12, // 6: cluster.SlaveGrpcService.TTL:input_type -> cluster.SlaveTTLRequest
	14, // 7: cluster.SlaveGrpcService.Heartbeat:input_type -> cluster.SlaveHeartbeatRequest
	1,  // 8: cluster.SlaveGrpcService.Get:output_type -> cluster.SlaveGetResponse
	3,  // 9: cluster.SlaveGrpcService.Set:output_type -> cluster.SlaveSetResponse
	5,  // 10: cluster.SlaveGrpcService.Del:output_type -> cluster.SlaveDelResponse
	7,  // 11: cluster.SlaveGrpcService.Keys:output_type -> cluster.SlaveKeysResponse
	9,  // 12: cluster.SlaveGrpcService.Exists:output_type -> cluster.SlaveExistsResponse
	11, // 13: cluster.SlaveGrpcService.Expire:output_type -> cluster.SlaveExpireResponse
	13, // 14: cluster.SlaveGrpcService.TTL:output_type -> cluster.SlaveTTLResponse
	15, // 15: cluster.SlaveGrpcService.Heartbeat:output_type -> cluster.SlaveHeartbeatResponse
	8,  // [8:16] is the sub-list for method output_type
	0,  // [0:8] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_lib_proto_slave_proto_init() }
func file_lib_proto_slave_proto_init() {
	if File_lib_proto_slave_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lib_proto_slave_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlaveGetRequest); i {
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
		file_lib_proto_slave_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlaveGetResponse); i {
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
		file_lib_proto_slave_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlaveSetRequest); i {
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
		file_lib_proto_slave_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlaveSetResponse); i {
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
		file_lib_proto_slave_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlaveDelRequest); i {
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
		file_lib_proto_slave_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlaveDelResponse); i {
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
		file_lib_proto_slave_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlaveKeysRequest); i {
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
		file_lib_proto_slave_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlaveKeysResponse); i {
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
		file_lib_proto_slave_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlaveExistsRequest); i {
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
		file_lib_proto_slave_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlaveExistsResponse); i {
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
		file_lib_proto_slave_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlaveExpireRequest); i {
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
		file_lib_proto_slave_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlaveExpireResponse); i {
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
		file_lib_proto_slave_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlaveTTLRequest); i {
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
		file_lib_proto_slave_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlaveTTLResponse); i {
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
		file_lib_proto_slave_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlaveHeartbeatRequest); i {
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
		file_lib_proto_slave_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlaveHeartbeatResponse); i {
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
			RawDescriptor: file_lib_proto_slave_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_lib_proto_slave_proto_goTypes,
		DependencyIndexes: file_lib_proto_slave_proto_depIdxs,
		MessageInfos:      file_lib_proto_slave_proto_msgTypes,
	}.Build()
	File_lib_proto_slave_proto = out.File
	file_lib_proto_slave_proto_rawDesc = nil
	file_lib_proto_slave_proto_goTypes = nil
	file_lib_proto_slave_proto_depIdxs = nil
}
