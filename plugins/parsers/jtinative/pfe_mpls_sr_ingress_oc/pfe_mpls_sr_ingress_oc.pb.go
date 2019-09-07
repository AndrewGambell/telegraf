// Code generated by protoc-gen-go. DO NOT EDIT.
// source: plugins/parsers/jtinative/pfe_mpls_sr_ingress_oc/pfe_mpls_sr_ingress_oc.proto

package pfe_mpls_sr_ingress_oc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import telemetry_top "github.com/influxdata/telegraf/plugins/parsers/jtinative/telemetry_top"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MplsPfeMplsSeIngress struct {
	SignalingProtocols   *MplsPfeMplsSeIngressSignalingProtocolsType `protobuf:"bytes,151,opt,name=signaling_protocols,json=signalingProtocols" json:"signaling_protocols,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                    `json:"-"`
	XXX_unrecognized     []byte                                      `json:"-"`
	XXX_sizecache        int32                                       `json:"-"`
}

func (m *MplsPfeMplsSeIngress) Reset()         { *m = MplsPfeMplsSeIngress{} }
func (m *MplsPfeMplsSeIngress) String() string { return proto.CompactTextString(m) }
func (*MplsPfeMplsSeIngress) ProtoMessage()    {}
func (*MplsPfeMplsSeIngress) Descriptor() ([]byte, []int) {
	return fileDescriptor_pfe_mpls_sr_ingress_oc_4ed69a7e8587ed10, []int{0}
}
func (m *MplsPfeMplsSeIngress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsPfeMplsSeIngress.Unmarshal(m, b)
}
func (m *MplsPfeMplsSeIngress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsPfeMplsSeIngress.Marshal(b, m, deterministic)
}
func (dst *MplsPfeMplsSeIngress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsPfeMplsSeIngress.Merge(dst, src)
}
func (m *MplsPfeMplsSeIngress) XXX_Size() int {
	return xxx_messageInfo_MplsPfeMplsSeIngress.Size(m)
}
func (m *MplsPfeMplsSeIngress) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsPfeMplsSeIngress.DiscardUnknown(m)
}

var xxx_messageInfo_MplsPfeMplsSeIngress proto.InternalMessageInfo

func (m *MplsPfeMplsSeIngress) GetSignalingProtocols() *MplsPfeMplsSeIngressSignalingProtocolsType {
	if m != nil {
		return m.SignalingProtocols
	}
	return nil
}

type MplsPfeMplsSeIngressSignalingProtocolsType struct {
	SegmentRouting       *MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingType `protobuf:"bytes,151,opt,name=segment_routing,json=segmentRouting" json:"segment_routing,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                      `json:"-"`
	XXX_unrecognized     []byte                                                        `json:"-"`
	XXX_sizecache        int32                                                         `json:"-"`
}

func (m *MplsPfeMplsSeIngressSignalingProtocolsType) Reset() {
	*m = MplsPfeMplsSeIngressSignalingProtocolsType{}
}
func (m *MplsPfeMplsSeIngressSignalingProtocolsType) String() string {
	return proto.CompactTextString(m)
}
func (*MplsPfeMplsSeIngressSignalingProtocolsType) ProtoMessage() {}
func (*MplsPfeMplsSeIngressSignalingProtocolsType) Descriptor() ([]byte, []int) {
	return fileDescriptor_pfe_mpls_sr_ingress_oc_4ed69a7e8587ed10, []int{0, 0}
}
func (m *MplsPfeMplsSeIngressSignalingProtocolsType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsPfeMplsSeIngressSignalingProtocolsType.Unmarshal(m, b)
}
func (m *MplsPfeMplsSeIngressSignalingProtocolsType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsPfeMplsSeIngressSignalingProtocolsType.Marshal(b, m, deterministic)
}
func (dst *MplsPfeMplsSeIngressSignalingProtocolsType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsPfeMplsSeIngressSignalingProtocolsType.Merge(dst, src)
}
func (m *MplsPfeMplsSeIngressSignalingProtocolsType) XXX_Size() int {
	return xxx_messageInfo_MplsPfeMplsSeIngressSignalingProtocolsType.Size(m)
}
func (m *MplsPfeMplsSeIngressSignalingProtocolsType) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsPfeMplsSeIngressSignalingProtocolsType.DiscardUnknown(m)
}

var xxx_messageInfo_MplsPfeMplsSeIngressSignalingProtocolsType proto.InternalMessageInfo

func (m *MplsPfeMplsSeIngressSignalingProtocolsType) GetSegmentRouting() *MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingType {
	if m != nil {
		return m.SegmentRouting
	}
	return nil
}

type MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingType struct {
	Interfaces           *MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingTypeInterfacesType `protobuf:"bytes,151,opt,name=interfaces" json:"interfaces,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                                    `json:"-"`
	XXX_unrecognized     []byte                                                                      `json:"-"`
	XXX_sizecache        int32                                                                       `json:"-"`
}

func (m *MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingType) Reset() {
	*m = MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingType{}
}
func (m *MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingType) String() string {
	return proto.CompactTextString(m)
}
func (*MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingType) ProtoMessage() {}
func (*MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingType) Descriptor() ([]byte, []int) {
	return fileDescriptor_pfe_mpls_sr_ingress_oc_4ed69a7e8587ed10, []int{0, 0, 0}
}
func (m *MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingType.Unmarshal(m, b)
}
func (m *MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingType.Marshal(b, m, deterministic)
}
func (dst *MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingType.Merge(dst, src)
}
func (m *MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingType) XXX_Size() int {
	return xxx_messageInfo_MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingType.Size(m)
}
func (m *MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingType) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingType.DiscardUnknown(m)
}

var xxx_messageInfo_MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingType proto.InternalMessageInfo

func (m *MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingType) GetInterfaces() *MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingTypeInterfacesType {
	if m != nil {
		return m.Interfaces
	}
	return nil
}

type MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingTypeInterfacesType struct {
	Interface            []*MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingTypeInterfacesTypeInterfaceList `protobuf:"bytes,151,rep,name=interface" json:"interface,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                                                   `json:"-"`
	XXX_unrecognized     []byte                                                                                     `json:"-"`
	XXX_sizecache        int32                                                                                      `json:"-"`
}

func (m *MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingTypeInterfacesType) Reset() {
	*m = MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingTypeInterfacesType{}
}
func (m *MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingTypeInterfacesType) String() string {
	return proto.CompactTextString(m)
}
func (*MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingTypeInterfacesType) ProtoMessage() {}
func (*MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingTypeInterfacesType) Descriptor() ([]byte, []int) {
	return fileDescriptor_pfe_mpls_sr_ingress_oc_4ed69a7e8587ed10, []int{0, 0, 0, 0}
}
func (m *MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingTypeInterfacesType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingTypeInterfacesType.Unmarshal(m, b)
}
func (m *MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingTypeInterfacesType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingTypeInterfacesType.Marshal(b, m, deterministic)
}
func (dst *MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingTypeInterfacesType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingTypeInterfacesType.Merge(dst, src)
}
func (m *MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingTypeInterfacesType) XXX_Size() int {
	return xxx_messageInfo_MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingTypeInterfacesType.Size(m)
}
func (m *MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingTypeInterfacesType) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingTypeInterfacesType.DiscardUnknown(m)
}

var xxx_messageInfo_MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingTypeInterfacesType proto.InternalMessageInfo

func (m *MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingTypeInterfacesType) GetInterface() []*MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingTypeInterfacesTypeInterfaceList {
	if m != nil {
		return m.Interface
	}
	return nil
}

type MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingTypeInterfacesTypeInterfaceList struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingTypeInterfacesTypeInterfaceList) Reset() {
	*m = MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingTypeInterfacesTypeInterfaceList{}
}
func (m *MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingTypeInterfacesTypeInterfaceList) String() string {
	return proto.CompactTextString(m)
}
func (*MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingTypeInterfacesTypeInterfaceList) ProtoMessage() {
}
func (*MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingTypeInterfacesTypeInterfaceList) Descriptor() ([]byte, []int) {
	return fileDescriptor_pfe_mpls_sr_ingress_oc_4ed69a7e8587ed10, []int{0, 0, 0, 0, 0}
}
func (m *MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingTypeInterfacesTypeInterfaceList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingTypeInterfacesTypeInterfaceList.Unmarshal(m, b)
}
func (m *MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingTypeInterfacesTypeInterfaceList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingTypeInterfacesTypeInterfaceList.Marshal(b, m, deterministic)
}
func (dst *MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingTypeInterfacesTypeInterfaceList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingTypeInterfacesTypeInterfaceList.Merge(dst, src)
}
func (m *MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingTypeInterfacesTypeInterfaceList) XXX_Size() int {
	return xxx_messageInfo_MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingTypeInterfacesTypeInterfaceList.Size(m)
}
func (m *MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingTypeInterfacesTypeInterfaceList) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingTypeInterfacesTypeInterfaceList.DiscardUnknown(m)
}

var xxx_messageInfo_MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingTypeInterfacesTypeInterfaceList proto.InternalMessageInfo

var E_JnprMplsPfeMplsSeIngressExt = &proto.ExtensionDesc{
	ExtendedType:  (*telemetry_top.JuniperNetworksSensors)(nil),
	ExtensionType: (*MplsPfeMplsSeIngress)(nil),
	Field:         76,
	Name:          "jnpr_mpls_pfe_mpls_se_ingress_ext",
	Tag:           "bytes,76,opt,name=jnpr_mpls_pfe_mpls_se_ingress_ext,json=jnprMplsPfeMplsSeIngressExt",
	Filename:      "plugins/parsers/jtinative/pfe_mpls_sr_ingress_oc/pfe_mpls_sr_ingress_oc.proto",
}

func init() {
	proto.RegisterType((*MplsPfeMplsSeIngress)(nil), "mpls_pfe_mpls_se_ingress")
	proto.RegisterType((*MplsPfeMplsSeIngressSignalingProtocolsType)(nil), "mpls_pfe_mpls_se_ingress.signaling_protocols_type")
	proto.RegisterType((*MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingType)(nil), "mpls_pfe_mpls_se_ingress.signaling_protocols_type.segment_routing_type")
	proto.RegisterType((*MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingTypeInterfacesType)(nil), "mpls_pfe_mpls_se_ingress.signaling_protocols_type.segment_routing_type.interfaces_type")
	proto.RegisterType((*MplsPfeMplsSeIngressSignalingProtocolsTypeSegmentRoutingTypeInterfacesTypeInterfaceList)(nil), "mpls_pfe_mpls_se_ingress.signaling_protocols_type.segment_routing_type.interfaces_type.interface_list")
	proto.RegisterExtension(E_JnprMplsPfeMplsSeIngressExt)
}

func init() {
	proto.RegisterFile("plugins/parsers/jtinative/pfe_mpls_sr_ingress_oc/pfe_mpls_sr_ingress_oc.proto", fileDescriptor_pfe_mpls_sr_ingress_oc_4ed69a7e8587ed10)
}

var fileDescriptor_pfe_mpls_sr_ingress_oc_4ed69a7e8587ed10 = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x90, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0x86, 0x53, 0xf1, 0xe2, 0x92, 0x80, 0x29, 0x26, 0xd6, 0x7a, 0x51, 0x4f, 0x9c, 0x4a, 0xc2,
	0xd1, 0xbb, 0x31, 0x1a, 0x31, 0xa4, 0x1c, 0x3c, 0x6e, 0x1a, 0x9c, 0x6e, 0x16, 0x97, 0xdd, 0x75,
	0x66, 0x50, 0xf0, 0xee, 0x33, 0xf8, 0x10, 0xbe, 0x83, 0x8f, 0x60, 0xe2, 0x1b, 0x19, 0xa8, 0x50,
	0x25, 0xf4, 0x60, 0xa2, 0xa7, 0xa6, 0xff, 0x3f, 0xff, 0xfc, 0xdf, 0x8e, 0xe8, 0x79, 0x33, 0x51,
	0xda, 0x52, 0xc7, 0x67, 0x48, 0x80, 0xd4, 0x19, 0xb1, 0xb6, 0x19, 0xeb, 0x07, 0xe8, 0xf8, 0x1c,
	0xe4, 0xd8, 0x1b, 0x92, 0x84, 0x52, 0x5b, 0x85, 0x40, 0x24, 0xdd, 0xb0, 0x42, 0x4e, 0x3c, 0x3a,
	0x76, 0x71, 0x8b, 0xc1, 0xc0, 0x18, 0x18, 0x67, 0x92, 0x9d, 0x2f, 0xc4, 0x93, 0x8f, 0x6d, 0x11,
	0x2d, 0x12, 0x65, 0x14, 0x96, 0xd1, 0xf0, 0x56, 0xb4, 0x48, 0x2b, 0x9b, 0x19, 0x6d, 0x95, 0x5c,
	0xcc, 0x0f, 0x9d, 0xa1, 0xe8, 0x25, 0x38, 0x0a, 0xda, 0xf5, 0x6e, 0x37, 0xa9, 0x0a, 0x26, 0x1b,
	0x52, 0x92, 0x67, 0x1e, 0xd2, 0x70, 0xe5, 0xf4, 0x97, 0x46, 0xfc, 0x5e, 0x13, 0x51, 0x55, 0x20,
	0xbc, 0x17, 0x4d, 0x02, 0x35, 0x06, 0xcb, 0x12, 0xdd, 0x84, 0xb5, 0x55, 0xcb, 0xfa, 0xf3, 0xdf,
	0xd7, 0x27, 0x6b, 0xab, 0x0a, 0xa6, 0xc6, 0x97, 0x9a, 0x16, 0x62, 0xfc, 0xb6, 0x25, 0xf6, 0x36,
	0x0d, 0x86, 0x53, 0x21, 0xb4, 0x65, 0xc0, 0x3c, 0x1b, 0xc2, 0xea, 0x0a, 0x37, 0x7f, 0x84, 0x91,
	0x94, 0xab, 0x0b, 0xac, 0x6f, 0x5d, 0xf1, 0x6b, 0x20, 0x9a, 0x6b, 0x7e, 0xf8, 0x1c, 0x88, 0x9d,
	0x95, 0x36, 0xa7, 0xa9, 0xb5, 0xeb, 0xdd, 0xfc, 0x9f, 0x68, 0xca, 0x7f, 0x69, 0x34, 0x71, 0x5a,
	0x36, 0xc7, 0xbb, 0xa2, 0xf1, 0xd3, 0x3c, 0x7d, 0x12, 0xc7, 0x23, 0xeb, 0x51, 0x56, 0xa1, 0x48,
	0x98, 0x72, 0xb8, 0x9f, 0x5c, 0x4e, 0xac, 0xf6, 0x80, 0xd7, 0xc0, 0x8f, 0x0e, 0xef, 0x68, 0x00,
	0x96, 0x1c, 0x52, 0x74, 0xb5, 0x38, 0xeb, 0x41, 0xe5, 0x43, 0xd2, 0xc3, 0xf9, 0xf2, 0x9e, 0x37,
	0xd4, 0xcf, 0x61, 0xfe, 0x19, 0xc0, 0x45, 0xe1, 0x9c, 0x4d, 0xf9, 0x33, 0x00, 0x00, 0xff, 0xff,
	0xb1, 0x0f, 0x1f, 0x40, 0x34, 0x03, 0x00, 0x00,
}
