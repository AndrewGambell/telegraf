// Code generated by protoc-gen-go. DO NOT EDIT.
// source: plugins/parsers/jtinative/pfe_mpls_sr_sid_oc/pfe_mpls_sr_sid_oc.proto

package pfe_mpls_sr_sid_oc

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

type MplsPfeMplsSrSid struct {
	SignalingProtocols   *MplsPfeMplsSrSidSignalingProtocolsType `protobuf:"bytes,151,opt,name=signaling_protocols,json=signalingProtocols" json:"signaling_protocols,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *MplsPfeMplsSrSid) Reset()         { *m = MplsPfeMplsSrSid{} }
func (m *MplsPfeMplsSrSid) String() string { return proto.CompactTextString(m) }
func (*MplsPfeMplsSrSid) ProtoMessage()    {}
func (*MplsPfeMplsSrSid) Descriptor() ([]byte, []int) {
	return fileDescriptor_pfe_mpls_sr_sid_oc_3ef755138710cb14, []int{0}
}
func (m *MplsPfeMplsSrSid) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsPfeMplsSrSid.Unmarshal(m, b)
}
func (m *MplsPfeMplsSrSid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsPfeMplsSrSid.Marshal(b, m, deterministic)
}
func (dst *MplsPfeMplsSrSid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsPfeMplsSrSid.Merge(dst, src)
}
func (m *MplsPfeMplsSrSid) XXX_Size() int {
	return xxx_messageInfo_MplsPfeMplsSrSid.Size(m)
}
func (m *MplsPfeMplsSrSid) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsPfeMplsSrSid.DiscardUnknown(m)
}

var xxx_messageInfo_MplsPfeMplsSrSid proto.InternalMessageInfo

func (m *MplsPfeMplsSrSid) GetSignalingProtocols() *MplsPfeMplsSrSidSignalingProtocolsType {
	if m != nil {
		return m.SignalingProtocols
	}
	return nil
}

type MplsPfeMplsSrSidSignalingProtocolsType struct {
	SegmentRouting       *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingType `protobuf:"bytes,151,opt,name=segment_routing,json=segmentRouting" json:"segment_routing,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                  `json:"-"`
	XXX_unrecognized     []byte                                                    `json:"-"`
	XXX_sizecache        int32                                                     `json:"-"`
}

func (m *MplsPfeMplsSrSidSignalingProtocolsType) Reset() {
	*m = MplsPfeMplsSrSidSignalingProtocolsType{}
}
func (m *MplsPfeMplsSrSidSignalingProtocolsType) String() string { return proto.CompactTextString(m) }
func (*MplsPfeMplsSrSidSignalingProtocolsType) ProtoMessage()    {}
func (*MplsPfeMplsSrSidSignalingProtocolsType) Descriptor() ([]byte, []int) {
	return fileDescriptor_pfe_mpls_sr_sid_oc_3ef755138710cb14, []int{0, 0}
}
func (m *MplsPfeMplsSrSidSignalingProtocolsType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsType.Unmarshal(m, b)
}
func (m *MplsPfeMplsSrSidSignalingProtocolsType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsType.Marshal(b, m, deterministic)
}
func (dst *MplsPfeMplsSrSidSignalingProtocolsType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsType.Merge(dst, src)
}
func (m *MplsPfeMplsSrSidSignalingProtocolsType) XXX_Size() int {
	return xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsType.Size(m)
}
func (m *MplsPfeMplsSrSidSignalingProtocolsType) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsType.DiscardUnknown(m)
}

var xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsType proto.InternalMessageInfo

func (m *MplsPfeMplsSrSidSignalingProtocolsType) GetSegmentRouting() *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingType {
	if m != nil {
		return m.SegmentRouting
	}
	return nil
}

type MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingType struct {
	AggregateSidCounters *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersType `protobuf:"bytes,151,opt,name=aggregate_sid_counters,json=aggregateSidCounters" json:"aggregate_sid_counters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                                          `json:"-"`
	XXX_unrecognized     []byte                                                                            `json:"-"`
	XXX_sizecache        int32                                                                             `json:"-"`
}

func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingType) Reset() {
	*m = MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingType{}
}
func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingType) String() string {
	return proto.CompactTextString(m)
}
func (*MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingType) ProtoMessage() {}
func (*MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingType) Descriptor() ([]byte, []int) {
	return fileDescriptor_pfe_mpls_sr_sid_oc_3ef755138710cb14, []int{0, 0, 0}
}
func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingType.Unmarshal(m, b)
}
func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingType.Marshal(b, m, deterministic)
}
func (dst *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingType.Merge(dst, src)
}
func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingType) XXX_Size() int {
	return xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingType.Size(m)
}
func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingType) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingType.DiscardUnknown(m)
}

var xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingType proto.InternalMessageInfo

func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingType) GetAggregateSidCounters() *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersType {
	if m != nil {
		return m.AggregateSidCounters
	}
	return nil
}

type MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersType struct {
	AggregateSidCounter  []*MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterList `protobuf:"bytes,151,rep,name=aggregate_sid_counter,json=aggregateSidCounter" json:"aggregate_sid_counter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                                                                   `json:"-"`
	XXX_unrecognized     []byte                                                                                                     `json:"-"`
	XXX_sizecache        int32                                                                                                      `json:"-"`
}

func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersType) Reset() {
	*m = MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersType{}
}
func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersType) String() string {
	return proto.CompactTextString(m)
}
func (*MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersType) ProtoMessage() {
}
func (*MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersType) Descriptor() ([]byte, []int) {
	return fileDescriptor_pfe_mpls_sr_sid_oc_3ef755138710cb14, []int{0, 0, 0, 0}
}
func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersType.Unmarshal(m, b)
}
func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersType.Marshal(b, m, deterministic)
}
func (dst *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersType.Merge(dst, src)
}
func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersType) XXX_Size() int {
	return xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersType.Size(m)
}
func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersType) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersType.DiscardUnknown(m)
}

var xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersType proto.InternalMessageInfo

func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersType) GetAggregateSidCounter() []*MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterList {
	if m != nil {
		return m.AggregateSidCounter
	}
	return nil
}

type MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterList struct {
	State                *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterListStateType `protobuf:"bytes,151,opt,name=state" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                                                                          `json:"-"`
	XXX_unrecognized     []byte                                                                                                            `json:"-"`
	XXX_sizecache        int32                                                                                                             `json:"-"`
}

func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterList) Reset() {
	*m = MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterList{}
}
func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterList) String() string {
	return proto.CompactTextString(m)
}
func (*MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterList) ProtoMessage() {
}
func (*MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterList) Descriptor() ([]byte, []int) {
	return fileDescriptor_pfe_mpls_sr_sid_oc_3ef755138710cb14, []int{0, 0, 0, 0, 0}
}
func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterList.Unmarshal(m, b)
}
func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterList.Marshal(b, m, deterministic)
}
func (dst *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterList.Merge(dst, src)
}
func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterList) XXX_Size() int {
	return xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterList.Size(m)
}
func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterList) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterList.DiscardUnknown(m)
}

var xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterList proto.InternalMessageInfo

func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterList) GetState() *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterListStateType {
	if m != nil {
		return m.State
	}
	return nil
}

type MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterListStateType struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterListStateType) Reset() {
	*m = MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterListStateType{}
}
func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterListStateType) String() string {
	return proto.CompactTextString(m)
}
func (*MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterListStateType) ProtoMessage() {
}
func (*MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterListStateType) Descriptor() ([]byte, []int) {
	return fileDescriptor_pfe_mpls_sr_sid_oc_3ef755138710cb14, []int{0, 0, 0, 0, 0, 0}
}
func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterListStateType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterListStateType.Unmarshal(m, b)
}
func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterListStateType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterListStateType.Marshal(b, m, deterministic)
}
func (dst *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterListStateType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterListStateType.Merge(dst, src)
}
func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterListStateType) XXX_Size() int {
	return xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterListStateType.Size(m)
}
func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterListStateType) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterListStateType.DiscardUnknown(m)
}

var xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterListStateType proto.InternalMessageInfo

var E_JnprMplsPfeMplsSrSidExt = &proto.ExtensionDesc{
	ExtendedType:  (*telemetry_top.JuniperNetworksSensors)(nil),
	ExtensionType: (*MplsPfeMplsSrSid)(nil),
	Field:         75,
	Name:          "jnpr_mpls_pfe_mpls_sr_sid_ext",
	Tag:           "bytes,75,opt,name=jnpr_mpls_pfe_mpls_sr_sid_ext,json=jnprMplsPfeMplsSrSidExt",
	Filename:      "plugins/parsers/jtinative/pfe_mpls_sr_sid_oc/pfe_mpls_sr_sid_oc.proto",
}

func init() {
	proto.RegisterType((*MplsPfeMplsSrSid)(nil), "mpls_pfe_mpls_sr_sid")
	proto.RegisterType((*MplsPfeMplsSrSidSignalingProtocolsType)(nil), "mpls_pfe_mpls_sr_sid.signaling_protocols_type")
	proto.RegisterType((*MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingType)(nil), "mpls_pfe_mpls_sr_sid.signaling_protocols_type.segment_routing_type")
	proto.RegisterType((*MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersType)(nil), "mpls_pfe_mpls_sr_sid.signaling_protocols_type.segment_routing_type.aggregate_sid_counters_type")
	proto.RegisterType((*MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterList)(nil), "mpls_pfe_mpls_sr_sid.signaling_protocols_type.segment_routing_type.aggregate_sid_counters_type.aggregate_sid_counter_list")
	proto.RegisterType((*MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterListStateType)(nil), "mpls_pfe_mpls_sr_sid.signaling_protocols_type.segment_routing_type.aggregate_sid_counters_type.aggregate_sid_counter_list.state_type")
	proto.RegisterExtension(E_JnprMplsPfeMplsSrSidExt)
}

func init() {
	proto.RegisterFile("plugins/parsers/jtinative/pfe_mpls_sr_sid_oc/pfe_mpls_sr_sid_oc.proto", fileDescriptor_pfe_mpls_sr_sid_oc_3ef755138710cb14)
}

var fileDescriptor_pfe_mpls_sr_sid_oc_3ef755138710cb14 = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x91, 0xc1, 0x4a, 0x23, 0x41,
	0x10, 0x86, 0x19, 0xb2, 0xd9, 0x43, 0x67, 0xd9, 0x85, 0x49, 0xb2, 0x09, 0xb3, 0x2c, 0x2c, 0x7b,
	0xca, 0xa9, 0x03, 0x39, 0x7a, 0x0d, 0xb9, 0x28, 0x4a, 0x98, 0x79, 0x80, 0x76, 0x48, 0x2a, 0x4d,
	0xc7, 0x99, 0xee, 0xa6, 0xab, 0x46, 0x13, 0xef, 0xbe, 0x82, 0x5e, 0xf4, 0x19, 0x7c, 0x0f, 0x0f,
	0xde, 0xbc, 0xf8, 0x34, 0x92, 0x99, 0x24, 0x42, 0xe8, 0x08, 0x82, 0xe2, 0x69, 0xe0, 0xff, 0xeb,
	0xef, 0xfa, 0xfe, 0x1a, 0x36, 0xb2, 0x59, 0x21, 0x95, 0xc6, 0xbe, 0x4d, 0x1d, 0x82, 0xc3, 0xfe,
	0x9c, 0x94, 0x4e, 0x49, 0x9d, 0x43, 0xdf, 0xce, 0x40, 0xe4, 0x36, 0x43, 0x81, 0x4e, 0xa0, 0x9a,
	0x0a, 0x33, 0xf1, 0x48, 0xdc, 0x3a, 0x43, 0x26, 0x6a, 0x12, 0x64, 0x90, 0x03, 0xb9, 0xa5, 0x20,
	0x63, 0x2b, 0xf1, 0xff, 0xd3, 0x77, 0xd6, 0x2a, 0xa7, 0x77, 0x62, 0xe1, 0x29, 0x6b, 0xa2, 0x92,
	0x3a, 0xcd, 0x94, 0x96, 0xa2, 0x9c, 0x9d, 0x98, 0x0c, 0xbb, 0x37, 0xc1, 0xbf, 0xa0, 0xd7, 0x18,
	0x70, 0xee, 0x0b, 0x71, 0x4f, 0x42, 0xd0, 0xd2, 0x42, 0x1c, 0x6e, 0x9d, 0xf1, 0xc6, 0x88, 0x9e,
	0xeb, 0xac, 0xbb, 0x2f, 0x10, 0xe6, 0xec, 0x17, 0x82, 0xcc, 0x41, 0x93, 0x70, 0xa6, 0x20, 0xa5,
	0xe5, 0x66, 0xf5, 0xf0, 0x7d, 0xab, 0xf9, 0xce, 0x33, 0x15, 0xcf, 0xcf, 0xb5, 0x1a, 0x57, 0x62,
	0xf4, 0xf8, 0x8d, 0xb5, 0x7c, 0x83, 0xe1, 0x75, 0xc0, 0x7e, 0xa7, 0x52, 0x3a, 0x90, 0x29, 0x41,
	0x79, 0xcf, 0x89, 0x29, 0x34, 0x81, 0xdb, 0x9e, 0x42, 0x7c, 0x00, 0x0f, 0xf7, 0xaf, 0xa8, 0x58,
	0x5b, 0x5b, 0x33, 0x51, 0xd3, 0xe1, 0xda, 0x8a, 0x6e, 0x6b, 0xec, 0xcf, 0x1b, 0xa9, 0xf0, 0x3e,
	0x60, 0x6d, 0xaf, 0xbf, 0xe2, 0xae, 0xf5, 0x1a, 0x83, 0xcb, 0x4f, 0xe6, 0xf6, 0x7b, 0x22, 0x53,
	0x48, 0x71, 0xd3, 0x53, 0x29, 0x7a, 0x08, 0x58, 0xb4, 0x3f, 0x13, 0xde, 0x05, 0xac, 0x8e, 0x94,
	0x12, 0x6c, 0x0e, 0x7f, 0x15, 0x7c, 0x5d, 0x03, 0x5e, 0x92, 0x54, 0xff, 0xa7, 0xa2, 0x8a, 0x7e,
	0x30, 0xf6, 0x2a, 0x1e, 0x68, 0xf6, 0x77, 0xae, 0xad, 0x13, 0x3e, 0x42, 0x01, 0x0b, 0x0a, 0x3b,
	0xfc, 0xb0, 0xd0, 0xca, 0x82, 0x3b, 0x01, 0xba, 0x30, 0xee, 0x0c, 0x13, 0xd0, 0x68, 0x1c, 0x76,
	0x8f, 0xca, 0x72, 0x6d, 0x6f, 0xb7, 0xb8, 0xb3, 0x7a, 0xf4, 0xd8, 0x66, 0x38, 0x9e, 0xc1, 0xea,
	0x93, 0xb8, 0x44, 0x4d, 0x47, 0x0b, 0x7a, 0x09, 0x00, 0x00, 0xff, 0xff, 0x70, 0xa9, 0xe8, 0xea,
	0x24, 0x04, 0x00, 0x00,
}
