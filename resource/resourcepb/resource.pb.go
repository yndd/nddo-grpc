//
//Copyright 2020 Ndd.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.17.3
// source: resource/resourcepb/resource.proto

package resourcepb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenant          string                 `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	NetworkInstance string                 `protobuf:"bytes,2,opt,name=networkInstance,proto3" json:"networkInstance,omitempty"`
	Kind            string                 `protobuf:"bytes,3,opt,name=kind,proto3" json:"kind,omitempty"`                                                                                                             // ipam, route-target, vlan
	Data            map[string]*TypedValue `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`                     // Map of tags (attribute) name to value.
	QueryTags       map[string]string      `protobuf:"bytes,5,rep,name=queryTags,proto3" json:"queryTags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`           // Map of tags (attribute) name to value.
	AdditionalTags  map[string]string      `protobuf:"bytes,6,rep,name=additionalTags,proto3" json:"additionalTags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // Map of tags (attribute) name to value.
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_resourcepb_resource_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_resource_resourcepb_resource_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_resource_resourcepb_resource_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *Request) GetNetworkInstance() string {
	if x != nil {
		return x.NetworkInstance
	}
	return ""
}

func (x *Request) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *Request) GetData() map[string]*TypedValue {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Request) GetQueryTags() map[string]string {
	if x != nil {
		return x.QueryTags
	}
	return nil
}

func (x *Request) GetAdditionalTags() map[string]string {
	if x != nil {
		return x.AdditionalTags
	}
	return nil
}

type Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp  int64                  `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`                                                                                        // Timestamp in nanoseconds since Epoch.
	ExpiryTime int64                  `protobuf:"varint,2,opt,name=expiryTime,proto3" json:"expiryTime,omitempty"`                                                                                      // ExpiryTime in nanoseconds since Epoch.
	QueryTags  map[string]*TypedValue `protobuf:"bytes,5,rep,name=queryTags,proto3" json:"queryTags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // Map of tags (attribute) name to value.
}

func (x *Reply) Reset() {
	*x = Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_resourcepb_resource_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply) ProtoMessage() {}

func (x *Reply) ProtoReflect() protoreflect.Message {
	mi := &file_resource_resourcepb_resource_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reply.ProtoReflect.Descriptor instead.
func (*Reply) Descriptor() ([]byte, []int) {
	return file_resource_resourcepb_resource_proto_rawDescGZIP(), []int{1}
}

func (x *Reply) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Reply) GetExpiryTime() int64 {
	if x != nil {
		return x.ExpiryTime
	}
	return 0
}

func (x *Reply) GetQueryTags() map[string]*TypedValue {
	if x != nil {
		return x.QueryTags
	}
	return nil
}

// TypedValue is used to encode a value being sent between the client and
// target (originated by either entity).
type TypedValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// One of the fields within the val oneof is populated with the value
	// of the update. The type of the value being included in the Update
	// determines which field should be populated. In the case that the
	// encoding is a particular form of the base protobuf type, a specific
	// field is used to store the value (e.g., json_val).
	//
	// Types that are assignable to Value:
	//	*TypedValue_StringVal
	//	*TypedValue_IntVal
	//	*TypedValue_UintVal
	//	*TypedValue_BoolVal
	//	*TypedValue_BytesVal
	//	*TypedValue_FloatVal
	//	*TypedValue_DecimalVal
	//	*TypedValue_LeaflistVal
	//	*TypedValue_AnyVal
	//	*TypedValue_JsonVal
	//	*TypedValue_JsonIetfVal
	//	*TypedValue_AsciiVal
	//	*TypedValue_ProtoBytes
	Value isTypedValue_Value `protobuf_oneof:"value"`
}

func (x *TypedValue) Reset() {
	*x = TypedValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_resourcepb_resource_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TypedValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TypedValue) ProtoMessage() {}

func (x *TypedValue) ProtoReflect() protoreflect.Message {
	mi := &file_resource_resourcepb_resource_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TypedValue.ProtoReflect.Descriptor instead.
func (*TypedValue) Descriptor() ([]byte, []int) {
	return file_resource_resourcepb_resource_proto_rawDescGZIP(), []int{2}
}

func (m *TypedValue) GetValue() isTypedValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *TypedValue) GetStringVal() string {
	if x, ok := x.GetValue().(*TypedValue_StringVal); ok {
		return x.StringVal
	}
	return ""
}

func (x *TypedValue) GetIntVal() int64 {
	if x, ok := x.GetValue().(*TypedValue_IntVal); ok {
		return x.IntVal
	}
	return 0
}

func (x *TypedValue) GetUintVal() uint64 {
	if x, ok := x.GetValue().(*TypedValue_UintVal); ok {
		return x.UintVal
	}
	return 0
}

func (x *TypedValue) GetBoolVal() bool {
	if x, ok := x.GetValue().(*TypedValue_BoolVal); ok {
		return x.BoolVal
	}
	return false
}

func (x *TypedValue) GetBytesVal() []byte {
	if x, ok := x.GetValue().(*TypedValue_BytesVal); ok {
		return x.BytesVal
	}
	return nil
}

func (x *TypedValue) GetFloatVal() float32 {
	if x, ok := x.GetValue().(*TypedValue_FloatVal); ok {
		return x.FloatVal
	}
	return 0
}

func (x *TypedValue) GetDecimalVal() *Decimal64 {
	if x, ok := x.GetValue().(*TypedValue_DecimalVal); ok {
		return x.DecimalVal
	}
	return nil
}

func (x *TypedValue) GetLeaflistVal() *ScalarArray {
	if x, ok := x.GetValue().(*TypedValue_LeaflistVal); ok {
		return x.LeaflistVal
	}
	return nil
}

func (x *TypedValue) GetAnyVal() *anypb.Any {
	if x, ok := x.GetValue().(*TypedValue_AnyVal); ok {
		return x.AnyVal
	}
	return nil
}

func (x *TypedValue) GetJsonVal() []byte {
	if x, ok := x.GetValue().(*TypedValue_JsonVal); ok {
		return x.JsonVal
	}
	return nil
}

func (x *TypedValue) GetJsonIetfVal() []byte {
	if x, ok := x.GetValue().(*TypedValue_JsonIetfVal); ok {
		return x.JsonIetfVal
	}
	return nil
}

func (x *TypedValue) GetAsciiVal() string {
	if x, ok := x.GetValue().(*TypedValue_AsciiVal); ok {
		return x.AsciiVal
	}
	return ""
}

func (x *TypedValue) GetProtoBytes() []byte {
	if x, ok := x.GetValue().(*TypedValue_ProtoBytes); ok {
		return x.ProtoBytes
	}
	return nil
}

type isTypedValue_Value interface {
	isTypedValue_Value()
}

type TypedValue_StringVal struct {
	StringVal string `protobuf:"bytes,1,opt,name=string_val,json=stringVal,proto3,oneof"` // String value.
}

type TypedValue_IntVal struct {
	IntVal int64 `protobuf:"varint,2,opt,name=int_val,json=intVal,proto3,oneof"` // Integer value.
}

type TypedValue_UintVal struct {
	UintVal uint64 `protobuf:"varint,3,opt,name=uint_val,json=uintVal,proto3,oneof"` // Unsigned integer value.
}

type TypedValue_BoolVal struct {
	BoolVal bool `protobuf:"varint,4,opt,name=bool_val,json=boolVal,proto3,oneof"` // Bool value.
}

type TypedValue_BytesVal struct {
	BytesVal []byte `protobuf:"bytes,5,opt,name=bytes_val,json=bytesVal,proto3,oneof"` // Arbitrary byte sequence value.
}

type TypedValue_FloatVal struct {
	FloatVal float32 `protobuf:"fixed32,6,opt,name=float_val,json=floatVal,proto3,oneof"` // Floating point value.
}

type TypedValue_DecimalVal struct {
	DecimalVal *Decimal64 `protobuf:"bytes,7,opt,name=decimal_val,json=decimalVal,proto3,oneof"` // Decimal64 encoded value.
}

type TypedValue_LeaflistVal struct {
	LeaflistVal *ScalarArray `protobuf:"bytes,8,opt,name=leaflist_val,json=leaflistVal,proto3,oneof"` // Mixed type scalar array value.
}

type TypedValue_AnyVal struct {
	AnyVal *anypb.Any `protobuf:"bytes,9,opt,name=any_val,json=anyVal,proto3,oneof"` // protobuf.Any encoded bytes.
}

type TypedValue_JsonVal struct {
	JsonVal []byte `protobuf:"bytes,10,opt,name=json_val,json=jsonVal,proto3,oneof"` // JSON-encoded text.
}

type TypedValue_JsonIetfVal struct {
	JsonIetfVal []byte `protobuf:"bytes,11,opt,name=json_ietf_val,json=jsonIetfVal,proto3,oneof"` // JSON-encoded text per RFC7951.
}

type TypedValue_AsciiVal struct {
	AsciiVal string `protobuf:"bytes,12,opt,name=ascii_val,json=asciiVal,proto3,oneof"` // Arbitrary ASCII text.
}

type TypedValue_ProtoBytes struct {
	// Protobuf binary encoded bytes. The message type is not included.
	// See the specification at
	// github.com/openconfig/reference/blob/master/rpc/gnmi/protobuf-vals.md
	// for a complete specification.
	ProtoBytes []byte `protobuf:"bytes,13,opt,name=proto_bytes,json=protoBytes,proto3,oneof"`
}

func (*TypedValue_StringVal) isTypedValue_Value() {}

func (*TypedValue_IntVal) isTypedValue_Value() {}

func (*TypedValue_UintVal) isTypedValue_Value() {}

func (*TypedValue_BoolVal) isTypedValue_Value() {}

func (*TypedValue_BytesVal) isTypedValue_Value() {}

func (*TypedValue_FloatVal) isTypedValue_Value() {}

func (*TypedValue_DecimalVal) isTypedValue_Value() {}

func (*TypedValue_LeaflistVal) isTypedValue_Value() {}

func (*TypedValue_AnyVal) isTypedValue_Value() {}

func (*TypedValue_JsonVal) isTypedValue_Value() {}

func (*TypedValue_JsonIetfVal) isTypedValue_Value() {}

func (*TypedValue_AsciiVal) isTypedValue_Value() {}

func (*TypedValue_ProtoBytes) isTypedValue_Value() {}

// Decimal64 is used to encode a fixed precision decimal number. The value
// is expressed as a set of digits with the precision specifying the
// number of digits following the decimal point in the digit set.
type Decimal64 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Digits    int64  `protobuf:"varint,1,opt,name=digits,proto3" json:"digits,omitempty"`       // Set of digits.
	Precision uint32 `protobuf:"varint,2,opt,name=precision,proto3" json:"precision,omitempty"` // Number of digits following the decimal point.
}

func (x *Decimal64) Reset() {
	*x = Decimal64{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_resourcepb_resource_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Decimal64) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Decimal64) ProtoMessage() {}

func (x *Decimal64) ProtoReflect() protoreflect.Message {
	mi := &file_resource_resourcepb_resource_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Decimal64.ProtoReflect.Descriptor instead.
func (*Decimal64) Descriptor() ([]byte, []int) {
	return file_resource_resourcepb_resource_proto_rawDescGZIP(), []int{3}
}

func (x *Decimal64) GetDigits() int64 {
	if x != nil {
		return x.Digits
	}
	return 0
}

func (x *Decimal64) GetPrecision() uint32 {
	if x != nil {
		return x.Precision
	}
	return 0
}

// ScalarArray is used to encode a mixed-type array of values.
type ScalarArray struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The set of elements within the array. Each TypedValue message should
	// specify only elements that have a field identifier of 1-7 (i.e., the
	// values are scalar values).
	Element []*TypedValue `protobuf:"bytes,1,rep,name=element,proto3" json:"element,omitempty"`
}

func (x *ScalarArray) Reset() {
	*x = ScalarArray{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_resourcepb_resource_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScalarArray) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScalarArray) ProtoMessage() {}

func (x *ScalarArray) ProtoReflect() protoreflect.Message {
	mi := &file_resource_resourcepb_resource_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScalarArray.ProtoReflect.Descriptor instead.
func (*ScalarArray) Descriptor() ([]byte, []int) {
	return file_resource_resourcepb_resource_proto_rawDescGZIP(), []int{4}
}

func (x *ScalarArray) GetElement() []*TypedValue {
	if x != nil {
		return x.Element
	}
	return nil
}

var File_resource_resourcepb_resource_proto protoreflect.FileDescriptor

var file_resource_resourcepb_resource_proto_rawDesc = []byte{
	0x0a, 0x22, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xef, 0x03, 0x0a, 0x07, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x28, 0x0a,
	0x0f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x2f, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3e, 0x0a, 0x09,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x54, 0x61, 0x67, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x09, 0x71, 0x75, 0x65, 0x72, 0x79, 0x54, 0x61, 0x67, 0x73, 0x12, 0x4d, 0x0a, 0x0e,
	0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x54, 0x61, 0x67, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e, 0x61, 0x64, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x54, 0x61, 0x67, 0x73, 0x1a, 0x4d, 0x0a, 0x09, 0x44,
	0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3c, 0x0a, 0x0e, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x41, 0x0a, 0x13, 0x41, 0x64, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xd7, 0x01, 0x0a, 0x05,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x09, 0x71, 0x75, 0x65, 0x72, 0x79, 0x54, 0x61, 0x67, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x61, 0x67,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x71, 0x75, 0x65, 0x72, 0x79, 0x54, 0x61, 0x67,
	0x73, 0x1a, 0x52, 0x0a, 0x0e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x54, 0x79, 0x70, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xf3, 0x03, 0x0a, 0x0a, 0x54, 0x79, 0x70, 0x65, 0x64, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f, 0x0a, 0x0a, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76,
	0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x73, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x12, 0x19, 0x0a, 0x07, 0x69, 0x6e, 0x74, 0x5f, 0x76, 0x61, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x06, 0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c,
	0x12, 0x1b, 0x0a, 0x08, 0x75, 0x69, 0x6e, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x48, 0x00, 0x52, 0x07, 0x75, 0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x12, 0x1b, 0x0a,
	0x08, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x48,
	0x00, 0x52, 0x07, 0x62, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x12, 0x1d, 0x0a, 0x09, 0x62, 0x79,
	0x74, 0x65, 0x73, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52,
	0x08, 0x62, 0x79, 0x74, 0x65, 0x73, 0x56, 0x61, 0x6c, 0x12, 0x1d, 0x0a, 0x09, 0x66, 0x6c, 0x6f,
	0x61, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x48, 0x00, 0x52, 0x08,
	0x66, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x12, 0x36, 0x0a, 0x0b, 0x64, 0x65, 0x63, 0x69,
	0x6d, 0x61, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c,
	0x36, 0x34, 0x48, 0x00, 0x52, 0x0a, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x56, 0x61, 0x6c,
	0x12, 0x3a, 0x0a, 0x0c, 0x6c, 0x65, 0x61, 0x66, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x76, 0x61, 0x6c,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x41, 0x72, 0x72, 0x61, 0x79, 0x48, 0x00, 0x52,
	0x0b, 0x6c, 0x65, 0x61, 0x66, 0x6c, 0x69, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x12, 0x2f, 0x0a, 0x07,
	0x61, 0x6e, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x41, 0x6e, 0x79, 0x48, 0x00, 0x52, 0x06, 0x61, 0x6e, 0x79, 0x56, 0x61, 0x6c, 0x12, 0x1b, 0x0a,
	0x08, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0c, 0x48,
	0x00, 0x52, 0x07, 0x6a, 0x73, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x12, 0x24, 0x0a, 0x0d, 0x6a, 0x73,
	0x6f, 0x6e, 0x5f, 0x69, 0x65, 0x74, 0x66, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0c, 0x48, 0x00, 0x52, 0x0b, 0x6a, 0x73, 0x6f, 0x6e, 0x49, 0x65, 0x74, 0x66, 0x56, 0x61, 0x6c,
	0x12, 0x1d, 0x0a, 0x09, 0x61, 0x73, 0x63, 0x69, 0x69, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x61, 0x73, 0x63, 0x69, 0x69, 0x56, 0x61, 0x6c, 0x12,
	0x21, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x79, 0x74,
	0x65, 0x73, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x41, 0x0a, 0x09, 0x44,
	0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x36, 0x34, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x69, 0x67, 0x69,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x64, 0x69, 0x67, 0x69, 0x74, 0x73,
	0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x72, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x3d,
	0x0a, 0x0b, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x41, 0x72, 0x72, 0x61, 0x79, 0x12, 0x2e, 0x0a,
	0x07, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x64, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x32, 0x39, 0x0a,
	0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x41, 0x6c, 0x6c,
	0x6f, 0x63, 0x12, 0x11, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x2d, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x2d, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2f, 0x6e, 0x64, 0x64, 0x2d, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resource_resourcepb_resource_proto_rawDescOnce sync.Once
	file_resource_resourcepb_resource_proto_rawDescData = file_resource_resourcepb_resource_proto_rawDesc
)

func file_resource_resourcepb_resource_proto_rawDescGZIP() []byte {
	file_resource_resourcepb_resource_proto_rawDescOnce.Do(func() {
		file_resource_resourcepb_resource_proto_rawDescData = protoimpl.X.CompressGZIP(file_resource_resourcepb_resource_proto_rawDescData)
	})
	return file_resource_resourcepb_resource_proto_rawDescData
}

var file_resource_resourcepb_resource_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_resource_resourcepb_resource_proto_goTypes = []interface{}{
	(*Request)(nil),     // 0: resource.Request
	(*Reply)(nil),       // 1: resource.Reply
	(*TypedValue)(nil),  // 2: resource.TypedValue
	(*Decimal64)(nil),   // 3: resource.Decimal64
	(*ScalarArray)(nil), // 4: resource.ScalarArray
	nil,                 // 5: resource.Request.DataEntry
	nil,                 // 6: resource.Request.QueryTagsEntry
	nil,                 // 7: resource.Request.AdditionalTagsEntry
	nil,                 // 8: resource.Reply.QueryTagsEntry
	(*anypb.Any)(nil),   // 9: google.protobuf.Any
}
var file_resource_resourcepb_resource_proto_depIdxs = []int32{
	5,  // 0: resource.Request.data:type_name -> resource.Request.DataEntry
	6,  // 1: resource.Request.queryTags:type_name -> resource.Request.QueryTagsEntry
	7,  // 2: resource.Request.additionalTags:type_name -> resource.Request.AdditionalTagsEntry
	8,  // 3: resource.Reply.queryTags:type_name -> resource.Reply.QueryTagsEntry
	3,  // 4: resource.TypedValue.decimal_val:type_name -> resource.Decimal64
	4,  // 5: resource.TypedValue.leaflist_val:type_name -> resource.ScalarArray
	9,  // 6: resource.TypedValue.any_val:type_name -> google.protobuf.Any
	2,  // 7: resource.ScalarArray.element:type_name -> resource.TypedValue
	2,  // 8: resource.Request.DataEntry.value:type_name -> resource.TypedValue
	2,  // 9: resource.Reply.QueryTagsEntry.value:type_name -> resource.TypedValue
	0,  // 10: resource.Resource.Alloc:input_type -> resource.Request
	1,  // 11: resource.Resource.Alloc:output_type -> resource.Reply
	11, // [11:12] is the sub-list for method output_type
	10, // [10:11] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_resource_resourcepb_resource_proto_init() }
func file_resource_resourcepb_resource_proto_init() {
	if File_resource_resourcepb_resource_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resource_resourcepb_resource_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_resource_resourcepb_resource_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reply); i {
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
		file_resource_resourcepb_resource_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TypedValue); i {
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
		file_resource_resourcepb_resource_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Decimal64); i {
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
		file_resource_resourcepb_resource_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScalarArray); i {
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
	file_resource_resourcepb_resource_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*TypedValue_StringVal)(nil),
		(*TypedValue_IntVal)(nil),
		(*TypedValue_UintVal)(nil),
		(*TypedValue_BoolVal)(nil),
		(*TypedValue_BytesVal)(nil),
		(*TypedValue_FloatVal)(nil),
		(*TypedValue_DecimalVal)(nil),
		(*TypedValue_LeaflistVal)(nil),
		(*TypedValue_AnyVal)(nil),
		(*TypedValue_JsonVal)(nil),
		(*TypedValue_JsonIetfVal)(nil),
		(*TypedValue_AsciiVal)(nil),
		(*TypedValue_ProtoBytes)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_resource_resourcepb_resource_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_resource_resourcepb_resource_proto_goTypes,
		DependencyIndexes: file_resource_resourcepb_resource_proto_depIdxs,
		MessageInfos:      file_resource_resourcepb_resource_proto_msgTypes,
	}.Build()
	File_resource_resourcepb_resource_proto = out.File
	file_resource_resourcepb_resource_proto_rawDesc = nil
	file_resource_resourcepb_resource_proto_goTypes = nil
	file_resource_resourcepb_resource_proto_depIdxs = nil
}
