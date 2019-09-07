// Code generated by protoc-gen-go. DO NOT EDIT.
// source: plugins/parsers/jtinative/telemetry_top/telemetry_top.proto

package telemetry_top

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

type TelemetryFieldOptions struct {
	IsKey                *bool    `protobuf:"varint,1,opt,name=is_key,json=isKey" json:"is_key,omitempty"`
	IsTimestamp          *bool    `protobuf:"varint,2,opt,name=is_timestamp,json=isTimestamp" json:"is_timestamp,omitempty"`
	IsCounter            *bool    `protobuf:"varint,3,opt,name=is_counter,json=isCounter" json:"is_counter,omitempty"`
	IsGauge              *bool    `protobuf:"varint,4,opt,name=is_gauge,json=isGauge" json:"is_gauge,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TelemetryFieldOptions) Reset()         { *m = TelemetryFieldOptions{} }
func (m *TelemetryFieldOptions) String() string { return proto.CompactTextString(m) }
func (*TelemetryFieldOptions) ProtoMessage()    {}
func (*TelemetryFieldOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_telemetry_top_7141970063c30f9d, []int{0}
}
func (m *TelemetryFieldOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TelemetryFieldOptions.Unmarshal(m, b)
}
func (m *TelemetryFieldOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TelemetryFieldOptions.Marshal(b, m, deterministic)
}
func (dst *TelemetryFieldOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TelemetryFieldOptions.Merge(dst, src)
}
func (m *TelemetryFieldOptions) XXX_Size() int {
	return xxx_messageInfo_TelemetryFieldOptions.Size(m)
}
func (m *TelemetryFieldOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_TelemetryFieldOptions.DiscardUnknown(m)
}

var xxx_messageInfo_TelemetryFieldOptions proto.InternalMessageInfo

func (m *TelemetryFieldOptions) GetIsKey() bool {
	if m != nil && m.IsKey != nil {
		return *m.IsKey
	}
	return false
}

func (m *TelemetryFieldOptions) GetIsTimestamp() bool {
	if m != nil && m.IsTimestamp != nil {
		return *m.IsTimestamp
	}
	return false
}

func (m *TelemetryFieldOptions) GetIsCounter() bool {
	if m != nil && m.IsCounter != nil {
		return *m.IsCounter
	}
	return false
}

func (m *TelemetryFieldOptions) GetIsGauge() bool {
	if m != nil && m.IsGauge != nil {
		return *m.IsGauge
	}
	return false
}

type TelemetryStream struct {
	// router name or export IP address
	SystemId *string `protobuf:"bytes,1,req,name=system_id,json=systemId" json:"system_id,omitempty"`
	// line card / RE (slot number). For RE, it will be 65535
	ComponentId *uint32 `protobuf:"varint,2,opt,name=component_id,json=componentId" json:"component_id,omitempty"`
	// PFE (if applicable)
	SubComponentId *uint32 `protobuf:"varint,3,opt,name=sub_component_id,json=subComponentId" json:"sub_component_id,omitempty"`
	// Overload sensor name with "senor name, internal path, external path
	// and component" seperated by ":". For RE sensors, component will be
	// daemon-name and for PFE sensors it will be "PFE".
	SensorName *string `protobuf:"bytes,4,opt,name=sensor_name,json=sensorName" json:"sensor_name,omitempty"`
	// sequence number, monotonically increasesing for each
	// system_id, component_id, sub_component_id + sensor_name.
	SequenceNumber *uint32 `protobuf:"varint,5,opt,name=sequence_number,json=sequenceNumber" json:"sequence_number,omitempty"`
	// timestamp (milliseconds since 00:00:00 UTC 1/1/1970)
	Timestamp *uint64 `protobuf:"varint,6,opt,name=timestamp" json:"timestamp,omitempty"`
	// major version
	VersionMajor *uint32 `protobuf:"varint,7,opt,name=version_major,json=versionMajor" json:"version_major,omitempty"`
	// minor version
	VersionMinor         *uint32            `protobuf:"varint,8,opt,name=version_minor,json=versionMinor" json:"version_minor,omitempty"`
	Ietf                 *IETFSensors       `protobuf:"bytes,100,opt,name=ietf" json:"ietf,omitempty"`
	Enterprise           *EnterpriseSensors `protobuf:"bytes,101,opt,name=enterprise" json:"enterprise,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *TelemetryStream) Reset()         { *m = TelemetryStream{} }
func (m *TelemetryStream) String() string { return proto.CompactTextString(m) }
func (*TelemetryStream) ProtoMessage()    {}
func (*TelemetryStream) Descriptor() ([]byte, []int) {
	return fileDescriptor_telemetry_top_7141970063c30f9d, []int{1}
}
func (m *TelemetryStream) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TelemetryStream.Unmarshal(m, b)
}
func (m *TelemetryStream) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TelemetryStream.Marshal(b, m, deterministic)
}
func (dst *TelemetryStream) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TelemetryStream.Merge(dst, src)
}
func (m *TelemetryStream) XXX_Size() int {
	return xxx_messageInfo_TelemetryStream.Size(m)
}
func (m *TelemetryStream) XXX_DiscardUnknown() {
	xxx_messageInfo_TelemetryStream.DiscardUnknown(m)
}

var xxx_messageInfo_TelemetryStream proto.InternalMessageInfo

func (m *TelemetryStream) GetSystemId() string {
	if m != nil && m.SystemId != nil {
		return *m.SystemId
	}
	return ""
}

func (m *TelemetryStream) GetComponentId() uint32 {
	if m != nil && m.ComponentId != nil {
		return *m.ComponentId
	}
	return 0
}

func (m *TelemetryStream) GetSubComponentId() uint32 {
	if m != nil && m.SubComponentId != nil {
		return *m.SubComponentId
	}
	return 0
}

func (m *TelemetryStream) GetSensorName() string {
	if m != nil && m.SensorName != nil {
		return *m.SensorName
	}
	return ""
}

func (m *TelemetryStream) GetSequenceNumber() uint32 {
	if m != nil && m.SequenceNumber != nil {
		return *m.SequenceNumber
	}
	return 0
}

func (m *TelemetryStream) GetTimestamp() uint64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *TelemetryStream) GetVersionMajor() uint32 {
	if m != nil && m.VersionMajor != nil {
		return *m.VersionMajor
	}
	return 0
}

func (m *TelemetryStream) GetVersionMinor() uint32 {
	if m != nil && m.VersionMinor != nil {
		return *m.VersionMinor
	}
	return 0
}

func (m *TelemetryStream) GetIetf() *IETFSensors {
	if m != nil {
		return m.Ietf
	}
	return nil
}

func (m *TelemetryStream) GetEnterprise() *EnterpriseSensors {
	if m != nil {
		return m.Enterprise
	}
	return nil
}

type IETFSensors struct {
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
	XXX_sizecache                int32  `json:"-"`
}

func (m *IETFSensors) Reset()         { *m = IETFSensors{} }
func (m *IETFSensors) String() string { return proto.CompactTextString(m) }
func (*IETFSensors) ProtoMessage()    {}
func (*IETFSensors) Descriptor() ([]byte, []int) {
	return fileDescriptor_telemetry_top_7141970063c30f9d, []int{2}
}

var extRange_IETFSensors = []proto.ExtensionRange{
	{Start: 1, End: 536870911},
}

func (*IETFSensors) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_IETFSensors
}
func (m *IETFSensors) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IETFSensors.Unmarshal(m, b)
}
func (m *IETFSensors) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IETFSensors.Marshal(b, m, deterministic)
}
func (dst *IETFSensors) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IETFSensors.Merge(dst, src)
}
func (m *IETFSensors) XXX_Size() int {
	return xxx_messageInfo_IETFSensors.Size(m)
}
func (m *IETFSensors) XXX_DiscardUnknown() {
	xxx_messageInfo_IETFSensors.DiscardUnknown(m)
}

var xxx_messageInfo_IETFSensors proto.InternalMessageInfo

type EnterpriseSensors struct {
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
	XXX_sizecache                int32  `json:"-"`
}

func (m *EnterpriseSensors) Reset()         { *m = EnterpriseSensors{} }
func (m *EnterpriseSensors) String() string { return proto.CompactTextString(m) }
func (*EnterpriseSensors) ProtoMessage()    {}
func (*EnterpriseSensors) Descriptor() ([]byte, []int) {
	return fileDescriptor_telemetry_top_7141970063c30f9d, []int{3}
}

var extRange_EnterpriseSensors = []proto.ExtensionRange{
	{Start: 1, End: 536870911},
}

func (*EnterpriseSensors) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_EnterpriseSensors
}
func (m *EnterpriseSensors) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnterpriseSensors.Unmarshal(m, b)
}
func (m *EnterpriseSensors) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnterpriseSensors.Marshal(b, m, deterministic)
}
func (dst *EnterpriseSensors) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnterpriseSensors.Merge(dst, src)
}
func (m *EnterpriseSensors) XXX_Size() int {
	return xxx_messageInfo_EnterpriseSensors.Size(m)
}
func (m *EnterpriseSensors) XXX_DiscardUnknown() {
	xxx_messageInfo_EnterpriseSensors.DiscardUnknown(m)
}

var xxx_messageInfo_EnterpriseSensors proto.InternalMessageInfo

type JuniperNetworksSensors struct {
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
	XXX_sizecache                int32  `json:"-"`
}

func (m *JuniperNetworksSensors) Reset()         { *m = JuniperNetworksSensors{} }
func (m *JuniperNetworksSensors) String() string { return proto.CompactTextString(m) }
func (*JuniperNetworksSensors) ProtoMessage()    {}
func (*JuniperNetworksSensors) Descriptor() ([]byte, []int) {
	return fileDescriptor_telemetry_top_7141970063c30f9d, []int{4}
}

var extRange_JuniperNetworksSensors = []proto.ExtensionRange{
	{Start: 1, End: 536870911},
}

func (*JuniperNetworksSensors) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_JuniperNetworksSensors
}
func (m *JuniperNetworksSensors) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JuniperNetworksSensors.Unmarshal(m, b)
}
func (m *JuniperNetworksSensors) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JuniperNetworksSensors.Marshal(b, m, deterministic)
}
func (dst *JuniperNetworksSensors) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JuniperNetworksSensors.Merge(dst, src)
}
func (m *JuniperNetworksSensors) XXX_Size() int {
	return xxx_messageInfo_JuniperNetworksSensors.Size(m)
}
func (m *JuniperNetworksSensors) XXX_DiscardUnknown() {
	xxx_messageInfo_JuniperNetworksSensors.DiscardUnknown(m)
}

var xxx_messageInfo_JuniperNetworksSensors proto.InternalMessageInfo

var E_TelemetryOptions = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*TelemetryFieldOptions)(nil),
	Field:         1024,
	Name:          "telemetry_options",
	Tag:           "bytes,1024,opt,name=telemetry_options,json=telemetryOptions",
	Filename:      "plugins/parsers/jtinative/telemetry_top/telemetry_top.proto",
}

var E_JuniperNetworks = &proto.ExtensionDesc{
	ExtendedType:  (*EnterpriseSensors)(nil),
	ExtensionType: (*JuniperNetworksSensors)(nil),
	Field:         2636,
	Name:          "juniperNetworks",
	Tag:           "bytes,2636,opt,name=juniperNetworks",
	Filename:      "plugins/parsers/jtinative/telemetry_top/telemetry_top.proto",
}

func init() {
	proto.RegisterType((*TelemetryFieldOptions)(nil), "TelemetryFieldOptions")
	proto.RegisterType((*TelemetryStream)(nil), "TelemetryStream")
	proto.RegisterType((*IETFSensors)(nil), "IETFSensors")
	proto.RegisterType((*EnterpriseSensors)(nil), "EnterpriseSensors")
	proto.RegisterType((*JuniperNetworksSensors)(nil), "JuniperNetworksSensors")
	proto.RegisterExtension(E_TelemetryOptions)
	proto.RegisterExtension(E_JuniperNetworks)
}

func init() {
	proto.RegisterFile("plugins/parsers/jtinative/telemetry_top/telemetry_top.proto", fileDescriptor_telemetry_top_7141970063c30f9d)
}

var fileDescriptor_telemetry_top_7141970063c30f9d = []byte{
	// 515 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x4f, 0x6f, 0xd3, 0x40,
	0x14, 0xc4, 0xe5, 0xf4, 0x9f, 0xf3, 0x92, 0xb6, 0xe9, 0x8a, 0xb6, 0x0b, 0x52, 0x25, 0xe3, 0x4a,
	0x10, 0x71, 0xb0, 0xa5, 0x1e, 0xcb, 0x05, 0x51, 0xb5, 0xa8, 0x20, 0x8a, 0xe4, 0x46, 0x5c, 0x2d,
	0x27, 0x7e, 0x8d, 0x36, 0x8d, 0x77, 0xcd, 0xbe, 0x75, 0x51, 0x6e, 0x11, 0x57, 0xbe, 0x07, 0x9f,
	0x86, 0x0f, 0x85, 0xbc, 0x8e, 0x13, 0x87, 0xe6, 0xb8, 0x33, 0xbf, 0x37, 0x9a, 0xc4, 0x03, 0xef,
	0xf3, 0x69, 0x31, 0x16, 0x92, 0xc2, 0x3c, 0xd1, 0x84, 0x9a, 0xc2, 0x89, 0x11, 0x32, 0x31, 0xe2,
	0x09, 0x43, 0x83, 0x53, 0xcc, 0xd0, 0xe8, 0x59, 0x6c, 0x54, 0xbe, 0xfe, 0x0a, 0x72, 0xad, 0x8c,
	0x7a, 0xe5, 0x8d, 0x95, 0x1a, 0x4f, 0x31, 0xb4, 0xaf, 0x61, 0xf1, 0x10, 0xa6, 0x48, 0x23, 0x2d,
	0x72, 0xa3, 0x74, 0x45, 0xf8, 0xbf, 0x1d, 0x38, 0x1e, 0xd4, 0x97, 0x37, 0x02, 0xa7, 0xe9, 0xb7,
	0xdc, 0x08, 0x25, 0x89, 0x1d, 0xc3, 0xae, 0xa0, 0xf8, 0x11, 0x67, 0xdc, 0xf1, 0x9c, 0xbe, 0x1b,
	0xed, 0x08, 0xfa, 0x82, 0x33, 0xf6, 0x1a, 0xba, 0x82, 0x62, 0x23, 0x32, 0x24, 0x93, 0x64, 0x39,
	0x6f, 0x59, 0xb3, 0x23, 0x68, 0x50, 0x4b, 0xec, 0x0c, 0x40, 0x50, 0x3c, 0x52, 0x85, 0x34, 0xa8,
	0xf9, 0x96, 0x05, 0xda, 0x82, 0xae, 0x2a, 0x81, 0xbd, 0x04, 0x57, 0x50, 0x3c, 0x4e, 0x8a, 0x31,
	0xf2, 0x6d, 0x6b, 0xee, 0x09, 0xfa, 0x54, 0x3e, 0xfd, 0x3f, 0x5b, 0x70, 0xb8, 0x6c, 0x73, 0x6f,
	0x34, 0x26, 0x19, 0xf3, 0xa1, 0x4d, 0x33, 0x32, 0x98, 0xc5, 0x22, 0xe5, 0x8e, 0xd7, 0xea, 0xb7,
	0x3f, 0xee, 0xfc, 0xfa, 0xd0, 0x72, 0x9d, 0xc8, 0xad, 0xf4, 0xdb, 0x94, 0xf5, 0xa1, 0x3b, 0x52,
	0x59, 0xae, 0x24, 0x4a, 0x53, 0x62, 0x65, 0xa9, 0xfd, 0x1a, 0xeb, 0x2c, 0xad, 0xdb, 0x94, 0x85,
	0xd0, 0xa3, 0x62, 0x18, 0xaf, 0xd1, 0x5b, 0x4d, 0xfa, 0x80, 0x8a, 0xe1, 0x55, 0xe3, 0xe0, 0x0d,
	0x74, 0x08, 0x25, 0x29, 0x1d, 0xcb, 0x24, 0xab, 0x0a, 0x2f, 0x0b, 0x40, 0xe5, 0xdc, 0x25, 0x19,
	0xb2, 0xb7, 0x70, 0x48, 0xf8, 0xa3, 0x40, 0x39, 0xc2, 0x58, 0x16, 0xd9, 0x10, 0x35, 0xdf, 0x29,
	0x73, 0xa3, 0x83, 0x5a, 0xbe, 0xb3, 0x2a, 0x3b, 0x87, 0xf6, 0xea, 0xdf, 0xdb, 0xf5, 0x9c, 0xfe,
	0xb6, 0x8d, 0xeb, 0x39, 0xd1, 0x4a, 0x67, 0xe7, 0xb0, 0xff, 0x84, 0x9a, 0x84, 0x92, 0x71, 0x96,
	0x4c, 0x94, 0xe6, 0x7b, 0x36, 0xab, 0xbb, 0x10, 0xbf, 0x96, 0xda, 0x1a, 0x24, 0xa4, 0xd2, 0xdc,
	0x5d, 0x87, 0x4a, 0x8d, 0x79, 0xb0, 0x2d, 0xd0, 0x3c, 0xf0, 0xd4, 0x73, 0xfa, 0x9d, 0x8b, 0x6e,
	0x70, 0x7b, 0x3d, 0xb8, 0xb9, 0xb7, 0xb5, 0x29, 0xb2, 0x0e, 0xbb, 0x00, 0xc0, 0xf2, 0xc3, 0xe4,
	0x5a, 0x10, 0x72, 0xb4, 0x1c, 0x0b, 0xae, 0x97, 0x52, 0x4d, 0x37, 0x28, 0xff, 0x14, 0x3a, 0x8d,
	0xa0, 0x77, 0xae, 0xeb, 0xf4, 0xe6, 0xf3, 0xf9, 0xbc, 0xe5, 0x9f, 0xc1, 0xd1, 0xb3, 0xcb, 0x86,
	0xed, 0xc3, 0xc9, 0xe7, 0x42, 0x8a, 0x1c, 0xf5, 0x1d, 0x9a, 0x9f, 0x4a, 0x3f, 0xd2, 0x33, 0xe6,
	0x72, 0x04, 0x47, 0xab, 0x2d, 0xab, 0xc5, 0x1a, 0xcf, 0x82, 0x6a, 0xca, 0x41, 0x3d, 0xe5, 0xa0,
	0x39, 0x56, 0x3e, 0x77, 0x6d, 0xed, 0x93, 0x60, 0xe3, 0x96, 0xa3, 0xde, 0x32, 0x70, 0xa1, 0x5c,
	0x7e, 0x87, 0xc3, 0xc9, 0x7a, 0x11, 0xb6, 0xe1, 0x37, 0xf3, 0xbf, 0x2f, 0x6c, 0xee, 0x69, 0xb0,
	0xb9, 0x75, 0xf4, 0x7f, 0xc8, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa9, 0xdf, 0x8f, 0x2c, 0xaf,
	0x03, 0x00, 0x00,
}