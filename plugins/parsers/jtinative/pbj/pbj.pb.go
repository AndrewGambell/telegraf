// Code generated by protoc-gen-go. DO NOT EDIT.
// source: plugins/parsers/jtinative/pbj/pbj.proto

package pbj

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FieldType int32

const (
	FieldType_FT_DEFAULT  FieldType = 0
	FieldType_FT_CALLBACK FieldType = 1
	FieldType_FT_POINTER  FieldType = 4
	FieldType_FT_STATIC   FieldType = 2
	FieldType_FT_IGNORE   FieldType = 3
)

var FieldType_name = map[int32]string{
	0: "FT_DEFAULT",
	1: "FT_CALLBACK",
	4: "FT_POINTER",
	2: "FT_STATIC",
	3: "FT_IGNORE",
}
var FieldType_value = map[string]int32{
	"FT_DEFAULT":  0,
	"FT_CALLBACK": 1,
	"FT_POINTER":  4,
	"FT_STATIC":   2,
	"FT_IGNORE":   3,
}

func (x FieldType) Enum() *FieldType {
	p := new(FieldType)
	*p = x
	return p
}
func (x FieldType) String() string {
	return proto.EnumName(FieldType_name, int32(x))
}
func (x *FieldType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(FieldType_value, data, "FieldType")
	if err != nil {
		return err
	}
	*x = FieldType(value)
	return nil
}
func (FieldType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_pbj_a5c5af291a1a3c81, []int{0}
}

// This is the inner options message, which basically defines options for
// a field. When it is used in message or file scope, it applies to all
// fields.
type PBJOptions struct {
	// Allocated size for 'bytes' and 'string' fields.
	MaxSize *int32 `protobuf:"varint,1,opt,name=max_size,json=maxSize" json:"max_size,omitempty"`
	// Allocated number of entries in arrays ('repeated' fields)
	MaxCount *int32 `protobuf:"varint,2,opt,name=max_count,json=maxCount" json:"max_count,omitempty"`
	// Force type of field (callback or static allocation)
	Type *FieldType `protobuf:"varint,3,opt,name=type,enum=FieldType,def=0" json:"type,omitempty"`
	// Use long names for enums, i.e. EnumName_EnumValue.
	LongNames *bool `protobuf:"varint,4,opt,name=long_names,json=longNames,def=1" json:"long_names,omitempty"`
	// Add 'packed' attribute to generated structs.
	// Note: this cannot be used on CPUs that break on unaligned
	// accesses to variables.
	PackedStruct *bool `protobuf:"varint,5,opt,name=packed_struct,json=packedStruct,def=0" json:"packed_struct,omitempty"`
	// Skip this message
	SkipMessage          *bool    `protobuf:"varint,6,opt,name=skip_message,json=skipMessage,def=0" json:"skip_message,omitempty"`
	CacheSize            *bool    `protobuf:"varint,7,opt,name=cache_size,json=cacheSize,def=1" json:"cache_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PBJOptions) Reset()         { *m = PBJOptions{} }
func (m *PBJOptions) String() string { return proto.CompactTextString(m) }
func (*PBJOptions) ProtoMessage()    {}
func (*PBJOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_pbj_a5c5af291a1a3c81, []int{0}
}
func (m *PBJOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PBJOptions.Unmarshal(m, b)
}
func (m *PBJOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PBJOptions.Marshal(b, m, deterministic)
}
func (dst *PBJOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PBJOptions.Merge(dst, src)
}
func (m *PBJOptions) XXX_Size() int {
	return xxx_messageInfo_PBJOptions.Size(m)
}
func (m *PBJOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_PBJOptions.DiscardUnknown(m)
}

var xxx_messageInfo_PBJOptions proto.InternalMessageInfo

const Default_PBJOptions_Type FieldType = FieldType_FT_DEFAULT
const Default_PBJOptions_LongNames bool = true
const Default_PBJOptions_PackedStruct bool = false
const Default_PBJOptions_SkipMessage bool = false
const Default_PBJOptions_CacheSize bool = true

func (m *PBJOptions) GetMaxSize() int32 {
	if m != nil && m.MaxSize != nil {
		return *m.MaxSize
	}
	return 0
}

func (m *PBJOptions) GetMaxCount() int32 {
	if m != nil && m.MaxCount != nil {
		return *m.MaxCount
	}
	return 0
}

func (m *PBJOptions) GetType() FieldType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Default_PBJOptions_Type
}

func (m *PBJOptions) GetLongNames() bool {
	if m != nil && m.LongNames != nil {
		return *m.LongNames
	}
	return Default_PBJOptions_LongNames
}

func (m *PBJOptions) GetPackedStruct() bool {
	if m != nil && m.PackedStruct != nil {
		return *m.PackedStruct
	}
	return Default_PBJOptions_PackedStruct
}

func (m *PBJOptions) GetSkipMessage() bool {
	if m != nil && m.SkipMessage != nil {
		return *m.SkipMessage
	}
	return Default_PBJOptions_SkipMessage
}

func (m *PBJOptions) GetCacheSize() bool {
	if m != nil && m.CacheSize != nil {
		return *m.CacheSize
	}
	return Default_PBJOptions_CacheSize
}

var E_PbjFileOption = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*PBJOptions)(nil),
	Field:         1020,
	Name:          "pbj_file_option",
	Tag:           "bytes,1020,opt,name=pbj_file_option,json=pbjFileOption",
	Filename:      "plugins/parsers/jtinative/pbj/pbj.proto",
}

var E_PbjMessageOption = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*PBJOptions)(nil),
	Field:         1020,
	Name:          "pbj_message_option",
	Tag:           "bytes,1020,opt,name=pbj_message_option,json=pbjMessageOption",
	Filename:      "plugins/parsers/jtinative/pbj/pbj.proto",
}

var E_PbjEnumOption = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumOptions)(nil),
	ExtensionType: (*PBJOptions)(nil),
	Field:         1020,
	Name:          "pbj_enum_option",
	Tag:           "bytes,1020,opt,name=pbj_enum_option,json=pbjEnumOption",
	Filename:      "plugins/parsers/jtinative/pbj/pbj.proto",
}

var E_PbjFieldOption = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*PBJOptions)(nil),
	Field:         1020,
	Name:          "pbj_field_option",
	Tag:           "bytes,1020,opt,name=pbj_field_option,json=pbjFieldOption",
	Filename:      "plugins/parsers/jtinative/pbj/pbj.proto",
}

func init() {
	proto.RegisterType((*PBJOptions)(nil), "PBJOptions")
	proto.RegisterEnum("FieldType", FieldType_name, FieldType_value)
	proto.RegisterExtension(E_PbjFileOption)
	proto.RegisterExtension(E_PbjMessageOption)
	proto.RegisterExtension(E_PbjEnumOption)
	proto.RegisterExtension(E_PbjFieldOption)
}

func init() {
	proto.RegisterFile("plugins/parsers/jtinative/pbj/pbj.proto", fileDescriptor_pbj_a5c5af291a1a3c81)
}

var fileDescriptor_pbj_a5c5af291a1a3c81 = []byte{
	// 467 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x4f, 0x6f, 0xda, 0x40,
	0x10, 0xc5, 0x0b, 0x81, 0x26, 0x0c, 0x01, 0xdc, 0x3d, 0xb9, 0xff, 0x54, 0xd4, 0x1e, 0x8a, 0x38,
	0x18, 0x29, 0x47, 0xdf, 0x80, 0xe2, 0x8a, 0x96, 0x42, 0x64, 0x9c, 0x4b, 0x7b, 0xb0, 0xfc, 0x67,
	0x70, 0xd7, 0xb1, 0xd7, 0x2b, 0xef, 0xba, 0x4a, 0xf2, 0x29, 0xfa, 0x81, 0x7b, 0xa8, 0xd6, 0x36,
	0x81, 0xaa, 0x6d, 0x72, 0xd8, 0xc3, 0xcc, 0xfb, 0xe9, 0xed, 0xcc, 0x1b, 0x78, 0xcf, 0x93, 0x22,
	0xa2, 0x4c, 0x4c, 0xb8, 0x97, 0x0b, 0xcc, 0xc5, 0x24, 0x96, 0x94, 0x79, 0x92, 0xfe, 0xc0, 0x09,
	0xf7, 0x63, 0xf5, 0x0c, 0x9e, 0x67, 0x32, 0x7b, 0x31, 0x8c, 0xb2, 0x2c, 0x4a, 0x70, 0x52, 0x56,
	0x7e, 0xb1, 0x9b, 0x84, 0x28, 0x82, 0x9c, 0x72, 0x99, 0xe5, 0x15, 0xf1, 0xf6, 0x67, 0x13, 0xe0,
	0x72, 0xf6, 0x69, 0xc3, 0x25, 0xcd, 0x98, 0x20, 0xcf, 0xe1, 0x2c, 0xf5, 0x6e, 0x5c, 0x41, 0xef,
	0x50, 0x6f, 0x0c, 0x1b, 0xa3, 0xb6, 0x7d, 0x9a, 0x7a, 0x37, 0x5b, 0x7a, 0x87, 0xe4, 0x25, 0x74,
	0x94, 0x14, 0x64, 0x05, 0x93, 0x7a, 0xb3, 0xd4, 0x14, 0x3b, 0x57, 0x35, 0x19, 0x43, 0x4b, 0xde,
	0x72, 0xd4, 0x4f, 0x86, 0x8d, 0x51, 0xff, 0x02, 0x0c, 0x8b, 0x62, 0x12, 0x3a, 0xb7, 0x1c, 0x4d,
	0xb0, 0x1c, 0xf7, 0xc3, 0xc2, 0x9a, 0x5e, 0xad, 0x1c, 0xbb, 0x64, 0xc8, 0x3b, 0x80, 0x24, 0x63,
	0x91, 0xcb, 0xbc, 0x14, 0x85, 0xde, 0x1a, 0x36, 0x46, 0x67, 0x66, 0x4b, 0xe6, 0x05, 0xda, 0x1d,
	0xd5, 0x5f, 0xab, 0x36, 0x19, 0x43, 0x8f, 0x7b, 0xc1, 0x35, 0x86, 0xae, 0x90, 0x79, 0x11, 0x48,
	0xbd, 0x5d, 0x72, 0xed, 0x9d, 0x97, 0x08, 0xb4, 0xcf, 0x2b, 0x6d, 0x5b, 0x4a, 0x64, 0x04, 0xe7,
	0xe2, 0x9a, 0x72, 0x37, 0x45, 0x21, 0xbc, 0x08, 0xf5, 0xa7, 0xc7, 0x68, 0x57, 0x49, 0x5f, 0x2a,
	0x45, 0x7d, 0x1d, 0x78, 0xc1, 0x77, 0xac, 0x16, 0x3c, 0x3d, 0xfe, 0xba, 0xec, 0xab, 0x45, 0xc7,
	0x5f, 0xa1, 0x73, 0x3f, 0x3e, 0xe9, 0xc3, 0xd1, 0x02, 0xda, 0x13, 0x32, 0x80, 0xae, 0xe5, 0xb8,
	0xf3, 0xe9, 0x6a, 0x35, 0x9b, 0xce, 0x3f, 0x6b, 0x8d, 0x1a, 0xb8, 0xdc, 0x2c, 0xd7, 0xce, 0xc2,
	0xd6, 0x5a, 0xa4, 0x07, 0x1d, 0xcb, 0x71, 0xb7, 0xce, 0xd4, 0x59, 0xce, 0xb5, 0x66, 0x5d, 0x2e,
	0x3f, 0xae, 0x37, 0xf6, 0x42, 0x3b, 0x31, 0x6d, 0x18, 0x70, 0x3f, 0x76, 0x77, 0x34, 0x41, 0x37,
	0x2b, 0x33, 0x27, 0xaf, 0x8c, 0xea, 0x48, 0xc6, 0xfe, 0x48, 0x86, 0x45, 0x13, 0xac, 0x0f, 0xa2,
	0xff, 0x52, 0x43, 0x76, 0x2f, 0xba, 0xc6, 0xe1, 0x48, 0x76, 0x8f, 0xfb, 0xf1, 0x81, 0x31, 0xbf,
	0x01, 0x51, 0x9e, 0xf5, 0xf6, 0x7b, 0xdb, 0x37, 0x7f, 0xd9, 0xd6, 0x21, 0x3c, 0xe4, 0xac, 0x71,
	0x3f, 0xfe, 0x03, 0xdb, 0x0f, 0x8c, 0xac, 0x48, 0xff, 0x3f, 0xf0, 0x82, 0x15, 0xe9, 0x63, 0x03,
	0x1f, 0x18, 0xf3, 0x0a, 0xb4, 0x2a, 0x04, 0x4c, 0xc2, 0xbd, 0xe9, 0xeb, 0x7f, 0xa4, 0x80, 0x49,
	0xf8, 0x90, 0x6b, 0xbf, 0x8c, 0xe1, 0x1e, 0x9a, 0x3d, 0x83, 0x01, 0x43, 0x69, 0xc4, 0x05, 0xa3,
	0x1c, 0x73, 0x83, 0xfb, 0xf1, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5d, 0xfd, 0x45, 0x85, 0x29,
	0x03, 0x00, 0x00,
}
