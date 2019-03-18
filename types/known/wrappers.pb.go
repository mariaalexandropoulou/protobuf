// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/protobuf/wrappers.proto

package known_proto

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "github.com/golang/protobuf/v2/reflect/protoreflect"
	protoimpl "github.com/golang/protobuf/v2/runtime/protoimpl"
	reflect "reflect"
)

// Wrapper message for `double`.
//
// The JSON representation for `DoubleValue` is JSON number.
type DoubleValue struct {
	// The double value.
	Value                float64  `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoubleValue) ProtoReflect() protoreflect.Message {
	return xxx_File_google_protobuf_wrappers_proto_messageTypes[0].MessageOf(m)
}
func (m *DoubleValue) Reset()         { *m = DoubleValue{} }
func (m *DoubleValue) String() string { return proto.CompactTextString(m) }
func (*DoubleValue) ProtoMessage()    {}

// Deprecated: Use DoubleValue.ProtoReflect.Type instead.
func (*DoubleValue) Descriptor() ([]byte, []int) {
	return xxx_File_google_protobuf_wrappers_proto_rawdesc_gzipped, []int{0}
}

func (*DoubleValue) XXX_WellKnownType() string { return "DoubleValue" }

func (m *DoubleValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoubleValue.Unmarshal(m, b)
}
func (m *DoubleValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoubleValue.Marshal(b, m, deterministic)
}
func (m *DoubleValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoubleValue.Merge(m, src)
}
func (m *DoubleValue) XXX_Size() int {
	return xxx_messageInfo_DoubleValue.Size(m)
}
func (m *DoubleValue) XXX_DiscardUnknown() {
	xxx_messageInfo_DoubleValue.DiscardUnknown(m)
}

var xxx_messageInfo_DoubleValue proto.InternalMessageInfo

func (m *DoubleValue) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

// Wrapper message for `float`.
//
// The JSON representation for `FloatValue` is JSON number.
type FloatValue struct {
	// The float value.
	Value                float32  `protobuf:"fixed32,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FloatValue) ProtoReflect() protoreflect.Message {
	return xxx_File_google_protobuf_wrappers_proto_messageTypes[1].MessageOf(m)
}
func (m *FloatValue) Reset()         { *m = FloatValue{} }
func (m *FloatValue) String() string { return proto.CompactTextString(m) }
func (*FloatValue) ProtoMessage()    {}

// Deprecated: Use FloatValue.ProtoReflect.Type instead.
func (*FloatValue) Descriptor() ([]byte, []int) {
	return xxx_File_google_protobuf_wrappers_proto_rawdesc_gzipped, []int{1}
}

func (*FloatValue) XXX_WellKnownType() string { return "FloatValue" }

func (m *FloatValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FloatValue.Unmarshal(m, b)
}
func (m *FloatValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FloatValue.Marshal(b, m, deterministic)
}
func (m *FloatValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FloatValue.Merge(m, src)
}
func (m *FloatValue) XXX_Size() int {
	return xxx_messageInfo_FloatValue.Size(m)
}
func (m *FloatValue) XXX_DiscardUnknown() {
	xxx_messageInfo_FloatValue.DiscardUnknown(m)
}

var xxx_messageInfo_FloatValue proto.InternalMessageInfo

func (m *FloatValue) GetValue() float32 {
	if m != nil {
		return m.Value
	}
	return 0
}

// Wrapper message for `int64`.
//
// The JSON representation for `Int64Value` is JSON string.
type Int64Value struct {
	// The int64 value.
	Value                int64    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Int64Value) ProtoReflect() protoreflect.Message {
	return xxx_File_google_protobuf_wrappers_proto_messageTypes[2].MessageOf(m)
}
func (m *Int64Value) Reset()         { *m = Int64Value{} }
func (m *Int64Value) String() string { return proto.CompactTextString(m) }
func (*Int64Value) ProtoMessage()    {}

// Deprecated: Use Int64Value.ProtoReflect.Type instead.
func (*Int64Value) Descriptor() ([]byte, []int) {
	return xxx_File_google_protobuf_wrappers_proto_rawdesc_gzipped, []int{2}
}

func (*Int64Value) XXX_WellKnownType() string { return "Int64Value" }

func (m *Int64Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Int64Value.Unmarshal(m, b)
}
func (m *Int64Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Int64Value.Marshal(b, m, deterministic)
}
func (m *Int64Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Int64Value.Merge(m, src)
}
func (m *Int64Value) XXX_Size() int {
	return xxx_messageInfo_Int64Value.Size(m)
}
func (m *Int64Value) XXX_DiscardUnknown() {
	xxx_messageInfo_Int64Value.DiscardUnknown(m)
}

var xxx_messageInfo_Int64Value proto.InternalMessageInfo

func (m *Int64Value) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

// Wrapper message for `uint64`.
//
// The JSON representation for `UInt64Value` is JSON string.
type UInt64Value struct {
	// The uint64 value.
	Value                uint64   `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UInt64Value) ProtoReflect() protoreflect.Message {
	return xxx_File_google_protobuf_wrappers_proto_messageTypes[3].MessageOf(m)
}
func (m *UInt64Value) Reset()         { *m = UInt64Value{} }
func (m *UInt64Value) String() string { return proto.CompactTextString(m) }
func (*UInt64Value) ProtoMessage()    {}

// Deprecated: Use UInt64Value.ProtoReflect.Type instead.
func (*UInt64Value) Descriptor() ([]byte, []int) {
	return xxx_File_google_protobuf_wrappers_proto_rawdesc_gzipped, []int{3}
}

func (*UInt64Value) XXX_WellKnownType() string { return "UInt64Value" }

func (m *UInt64Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UInt64Value.Unmarshal(m, b)
}
func (m *UInt64Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UInt64Value.Marshal(b, m, deterministic)
}
func (m *UInt64Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UInt64Value.Merge(m, src)
}
func (m *UInt64Value) XXX_Size() int {
	return xxx_messageInfo_UInt64Value.Size(m)
}
func (m *UInt64Value) XXX_DiscardUnknown() {
	xxx_messageInfo_UInt64Value.DiscardUnknown(m)
}

var xxx_messageInfo_UInt64Value proto.InternalMessageInfo

func (m *UInt64Value) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

// Wrapper message for `int32`.
//
// The JSON representation for `Int32Value` is JSON number.
type Int32Value struct {
	// The int32 value.
	Value                int32    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Int32Value) ProtoReflect() protoreflect.Message {
	return xxx_File_google_protobuf_wrappers_proto_messageTypes[4].MessageOf(m)
}
func (m *Int32Value) Reset()         { *m = Int32Value{} }
func (m *Int32Value) String() string { return proto.CompactTextString(m) }
func (*Int32Value) ProtoMessage()    {}

// Deprecated: Use Int32Value.ProtoReflect.Type instead.
func (*Int32Value) Descriptor() ([]byte, []int) {
	return xxx_File_google_protobuf_wrappers_proto_rawdesc_gzipped, []int{4}
}

func (*Int32Value) XXX_WellKnownType() string { return "Int32Value" }

func (m *Int32Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Int32Value.Unmarshal(m, b)
}
func (m *Int32Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Int32Value.Marshal(b, m, deterministic)
}
func (m *Int32Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Int32Value.Merge(m, src)
}
func (m *Int32Value) XXX_Size() int {
	return xxx_messageInfo_Int32Value.Size(m)
}
func (m *Int32Value) XXX_DiscardUnknown() {
	xxx_messageInfo_Int32Value.DiscardUnknown(m)
}

var xxx_messageInfo_Int32Value proto.InternalMessageInfo

func (m *Int32Value) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

// Wrapper message for `uint32`.
//
// The JSON representation for `UInt32Value` is JSON number.
type UInt32Value struct {
	// The uint32 value.
	Value                uint32   `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UInt32Value) ProtoReflect() protoreflect.Message {
	return xxx_File_google_protobuf_wrappers_proto_messageTypes[5].MessageOf(m)
}
func (m *UInt32Value) Reset()         { *m = UInt32Value{} }
func (m *UInt32Value) String() string { return proto.CompactTextString(m) }
func (*UInt32Value) ProtoMessage()    {}

// Deprecated: Use UInt32Value.ProtoReflect.Type instead.
func (*UInt32Value) Descriptor() ([]byte, []int) {
	return xxx_File_google_protobuf_wrappers_proto_rawdesc_gzipped, []int{5}
}

func (*UInt32Value) XXX_WellKnownType() string { return "UInt32Value" }

func (m *UInt32Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UInt32Value.Unmarshal(m, b)
}
func (m *UInt32Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UInt32Value.Marshal(b, m, deterministic)
}
func (m *UInt32Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UInt32Value.Merge(m, src)
}
func (m *UInt32Value) XXX_Size() int {
	return xxx_messageInfo_UInt32Value.Size(m)
}
func (m *UInt32Value) XXX_DiscardUnknown() {
	xxx_messageInfo_UInt32Value.DiscardUnknown(m)
}

var xxx_messageInfo_UInt32Value proto.InternalMessageInfo

func (m *UInt32Value) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

// Wrapper message for `bool`.
//
// The JSON representation for `BoolValue` is JSON `true` and `false`.
type BoolValue struct {
	// The bool value.
	Value                bool     `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BoolValue) ProtoReflect() protoreflect.Message {
	return xxx_File_google_protobuf_wrappers_proto_messageTypes[6].MessageOf(m)
}
func (m *BoolValue) Reset()         { *m = BoolValue{} }
func (m *BoolValue) String() string { return proto.CompactTextString(m) }
func (*BoolValue) ProtoMessage()    {}

// Deprecated: Use BoolValue.ProtoReflect.Type instead.
func (*BoolValue) Descriptor() ([]byte, []int) {
	return xxx_File_google_protobuf_wrappers_proto_rawdesc_gzipped, []int{6}
}

func (*BoolValue) XXX_WellKnownType() string { return "BoolValue" }

func (m *BoolValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoolValue.Unmarshal(m, b)
}
func (m *BoolValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoolValue.Marshal(b, m, deterministic)
}
func (m *BoolValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoolValue.Merge(m, src)
}
func (m *BoolValue) XXX_Size() int {
	return xxx_messageInfo_BoolValue.Size(m)
}
func (m *BoolValue) XXX_DiscardUnknown() {
	xxx_messageInfo_BoolValue.DiscardUnknown(m)
}

var xxx_messageInfo_BoolValue proto.InternalMessageInfo

func (m *BoolValue) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

// Wrapper message for `string`.
//
// The JSON representation for `StringValue` is JSON string.
type StringValue struct {
	// The string value.
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringValue) ProtoReflect() protoreflect.Message {
	return xxx_File_google_protobuf_wrappers_proto_messageTypes[7].MessageOf(m)
}
func (m *StringValue) Reset()         { *m = StringValue{} }
func (m *StringValue) String() string { return proto.CompactTextString(m) }
func (*StringValue) ProtoMessage()    {}

// Deprecated: Use StringValue.ProtoReflect.Type instead.
func (*StringValue) Descriptor() ([]byte, []int) {
	return xxx_File_google_protobuf_wrappers_proto_rawdesc_gzipped, []int{7}
}

func (*StringValue) XXX_WellKnownType() string { return "StringValue" }

func (m *StringValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringValue.Unmarshal(m, b)
}
func (m *StringValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringValue.Marshal(b, m, deterministic)
}
func (m *StringValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringValue.Merge(m, src)
}
func (m *StringValue) XXX_Size() int {
	return xxx_messageInfo_StringValue.Size(m)
}
func (m *StringValue) XXX_DiscardUnknown() {
	xxx_messageInfo_StringValue.DiscardUnknown(m)
}

var xxx_messageInfo_StringValue proto.InternalMessageInfo

func (m *StringValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// Wrapper message for `bytes`.
//
// The JSON representation for `BytesValue` is JSON string.
type BytesValue struct {
	// The bytes value.
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BytesValue) ProtoReflect() protoreflect.Message {
	return xxx_File_google_protobuf_wrappers_proto_messageTypes[8].MessageOf(m)
}
func (m *BytesValue) Reset()         { *m = BytesValue{} }
func (m *BytesValue) String() string { return proto.CompactTextString(m) }
func (*BytesValue) ProtoMessage()    {}

// Deprecated: Use BytesValue.ProtoReflect.Type instead.
func (*BytesValue) Descriptor() ([]byte, []int) {
	return xxx_File_google_protobuf_wrappers_proto_rawdesc_gzipped, []int{8}
}

func (*BytesValue) XXX_WellKnownType() string { return "BytesValue" }

func (m *BytesValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BytesValue.Unmarshal(m, b)
}
func (m *BytesValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BytesValue.Marshal(b, m, deterministic)
}
func (m *BytesValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BytesValue.Merge(m, src)
}
func (m *BytesValue) XXX_Size() int {
	return xxx_messageInfo_BytesValue.Size(m)
}
func (m *BytesValue) XXX_DiscardUnknown() {
	xxx_messageInfo_BytesValue.DiscardUnknown(m)
}

var xxx_messageInfo_BytesValue proto.InternalMessageInfo

func (m *BytesValue) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterFile("google/protobuf/wrappers.proto", xxx_File_google_protobuf_wrappers_proto_rawdesc_gzipped)
	proto.RegisterType((*DoubleValue)(nil), "google.protobuf.DoubleValue")
	proto.RegisterType((*FloatValue)(nil), "google.protobuf.FloatValue")
	proto.RegisterType((*Int64Value)(nil), "google.protobuf.Int64Value")
	proto.RegisterType((*UInt64Value)(nil), "google.protobuf.UInt64Value")
	proto.RegisterType((*Int32Value)(nil), "google.protobuf.Int32Value")
	proto.RegisterType((*UInt32Value)(nil), "google.protobuf.UInt32Value")
	proto.RegisterType((*BoolValue)(nil), "google.protobuf.BoolValue")
	proto.RegisterType((*StringValue)(nil), "google.protobuf.StringValue")
	proto.RegisterType((*BytesValue)(nil), "google.protobuf.BytesValue")
}

var xxx_File_google_protobuf_wrappers_proto_rawdesc = []byte{
	// 522 bytes of the wire-encoded FileDescriptorProto
	0x0a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x22, 0x23, 0x0a, 0x0b, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x22, 0x0a, 0x0a, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x22, 0x0a, 0x0a, 0x49, 0x6e,
	0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x23,
	0x0a, 0x0b, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x22, 0x0a, 0x0a, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x23, 0x0a, 0x0b, 0x55, 0x49, 0x6e, 0x74, 0x33,
	0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x21, 0x0a, 0x09,
	0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x23, 0x0a, 0x0b, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x22, 0x0a, 0x0a, 0x42, 0x79, 0x74, 0x65, 0x73, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x87, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x42, 0x0d, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x32,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x3b, 0x6b, 0x6e, 0x6f,
	0x77, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x03, 0x47, 0x50,
	0x42, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x57, 0x65, 0x6c, 0x6c, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var xxx_File_google_protobuf_wrappers_proto_rawdesc_gzipped = protoimpl.X.CompressGZIP(xxx_File_google_protobuf_wrappers_proto_rawdesc)

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

var File_google_protobuf_wrappers_proto protoreflect.FileDescriptor

var xxx_File_google_protobuf_wrappers_proto_messageTypes = make([]protoimpl.MessageType, 9)
var xxx_File_google_protobuf_wrappers_proto_goTypes = []interface{}{
	(*DoubleValue)(nil), // 0: google.protobuf.DoubleValue
	(*FloatValue)(nil),  // 1: google.protobuf.FloatValue
	(*Int64Value)(nil),  // 2: google.protobuf.Int64Value
	(*UInt64Value)(nil), // 3: google.protobuf.UInt64Value
	(*Int32Value)(nil),  // 4: google.protobuf.Int32Value
	(*UInt32Value)(nil), // 5: google.protobuf.UInt32Value
	(*BoolValue)(nil),   // 6: google.protobuf.BoolValue
	(*StringValue)(nil), // 7: google.protobuf.StringValue
	(*BytesValue)(nil),  // 8: google.protobuf.BytesValue
}
var xxx_File_google_protobuf_wrappers_proto_depIdxs = []int32{}

func init() { xxx_File_google_protobuf_wrappers_proto_init() }
func xxx_File_google_protobuf_wrappers_proto_init() {
	if File_google_protobuf_wrappers_proto != nil {
		return
	}
	messageTypes := make([]protoreflect.MessageType, 9)
	File_google_protobuf_wrappers_proto = protoimpl.FileBuilder{
		RawDescriptor:      xxx_File_google_protobuf_wrappers_proto_rawdesc,
		GoTypes:            xxx_File_google_protobuf_wrappers_proto_goTypes,
		DependencyIndexes:  xxx_File_google_protobuf_wrappers_proto_depIdxs,
		MessageOutputTypes: messageTypes,
	}.Init()
	messageGoTypes := xxx_File_google_protobuf_wrappers_proto_goTypes[0:][:9]
	for i, mt := range messageTypes {
		xxx_File_google_protobuf_wrappers_proto_messageTypes[i].GoType = reflect.TypeOf(messageGoTypes[i])
		xxx_File_google_protobuf_wrappers_proto_messageTypes[i].PBType = mt
	}
	xxx_File_google_protobuf_wrappers_proto_goTypes = nil
	xxx_File_google_protobuf_wrappers_proto_depIdxs = nil
}