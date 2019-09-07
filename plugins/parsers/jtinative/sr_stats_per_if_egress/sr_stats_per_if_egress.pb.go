// Code generated by protoc-gen-go. DO NOT EDIT.
// source: plugins/parsers/jtinative/sr_stats_per_if_egress/sr_stats_per_if_egress.proto

package sr_stats_per_if_egress

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

//
// Top-level message
//
type SrStatsPerIfEgress struct {
	// List of SR stats per IF egress records
	PerIfRecords         []*SegmentRoutingInterfaceRecord `protobuf:"bytes,1,rep,name=per_if_records,json=perIfRecords" json:"per_if_records,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *SrStatsPerIfEgress) Reset()         { *m = SrStatsPerIfEgress{} }
func (m *SrStatsPerIfEgress) String() string { return proto.CompactTextString(m) }
func (*SrStatsPerIfEgress) ProtoMessage()    {}
func (*SrStatsPerIfEgress) Descriptor() ([]byte, []int) {
	return fileDescriptor_sr_stats_per_if_egress_9167b4ebc03e7ddd, []int{0}
}
func (m *SrStatsPerIfEgress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SrStatsPerIfEgress.Unmarshal(m, b)
}
func (m *SrStatsPerIfEgress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SrStatsPerIfEgress.Marshal(b, m, deterministic)
}
func (dst *SrStatsPerIfEgress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SrStatsPerIfEgress.Merge(dst, src)
}
func (m *SrStatsPerIfEgress) XXX_Size() int {
	return xxx_messageInfo_SrStatsPerIfEgress.Size(m)
}
func (m *SrStatsPerIfEgress) XXX_DiscardUnknown() {
	xxx_messageInfo_SrStatsPerIfEgress.DiscardUnknown(m)
}

var xxx_messageInfo_SrStatsPerIfEgress proto.InternalMessageInfo

func (m *SrStatsPerIfEgress) GetPerIfRecords() []*SegmentRoutingInterfaceRecord {
	if m != nil {
		return m.PerIfRecords
	}
	return nil
}

//
// SR statistics record
//
type SegmentRoutingInterfaceRecord struct {
	// Interface name, e.g., xe-0/0/0
	IfName *string `protobuf:"bytes,1,req,name=if_name,json=ifName" json:"if_name,omitempty"`
	// Name of parent for AE interface, if applicable
	ParentAeName *string `protobuf:"bytes,2,opt,name=parent_ae_name,json=parentAeName" json:"parent_ae_name,omitempty"`
	// Name of the counter. This is useful when an interface has multiple counters.
	// for some scenarios, it is possible that a new counter is
	// created in the hardware.
	CounterName *string `protobuf:"bytes,3,opt,name=counter_name,json=counterName" json:"counter_name,omitempty"`
	// Traffic statistics
	IngressStats         *SegmentRoutingIfStats `protobuf:"bytes,4,opt,name=ingress_stats,json=ingressStats" json:"ingress_stats,omitempty"`
	EgressStats          *SegmentRoutingIfStats `protobuf:"bytes,5,opt,name=egress_stats,json=egressStats" json:"egress_stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SegmentRoutingInterfaceRecord) Reset()         { *m = SegmentRoutingInterfaceRecord{} }
func (m *SegmentRoutingInterfaceRecord) String() string { return proto.CompactTextString(m) }
func (*SegmentRoutingInterfaceRecord) ProtoMessage()    {}
func (*SegmentRoutingInterfaceRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_sr_stats_per_if_egress_9167b4ebc03e7ddd, []int{1}
}
func (m *SegmentRoutingInterfaceRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SegmentRoutingInterfaceRecord.Unmarshal(m, b)
}
func (m *SegmentRoutingInterfaceRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SegmentRoutingInterfaceRecord.Marshal(b, m, deterministic)
}
func (dst *SegmentRoutingInterfaceRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SegmentRoutingInterfaceRecord.Merge(dst, src)
}
func (m *SegmentRoutingInterfaceRecord) XXX_Size() int {
	return xxx_messageInfo_SegmentRoutingInterfaceRecord.Size(m)
}
func (m *SegmentRoutingInterfaceRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_SegmentRoutingInterfaceRecord.DiscardUnknown(m)
}

var xxx_messageInfo_SegmentRoutingInterfaceRecord proto.InternalMessageInfo

func (m *SegmentRoutingInterfaceRecord) GetIfName() string {
	if m != nil && m.IfName != nil {
		return *m.IfName
	}
	return ""
}

func (m *SegmentRoutingInterfaceRecord) GetParentAeName() string {
	if m != nil && m.ParentAeName != nil {
		return *m.ParentAeName
	}
	return ""
}

func (m *SegmentRoutingInterfaceRecord) GetCounterName() string {
	if m != nil && m.CounterName != nil {
		return *m.CounterName
	}
	return ""
}

func (m *SegmentRoutingInterfaceRecord) GetIngressStats() *SegmentRoutingIfStats {
	if m != nil {
		return m.IngressStats
	}
	return nil
}

func (m *SegmentRoutingInterfaceRecord) GetEgressStats() *SegmentRoutingIfStats {
	if m != nil {
		return m.EgressStats
	}
	return nil
}

type SegmentRoutingIfStats struct {
	// Packet and Byte statistics
	Packets *uint64 `protobuf:"varint,1,opt,name=packets" json:"packets,omitempty"`
	Bytes   *uint64 `protobuf:"varint,2,opt,name=bytes" json:"bytes,omitempty"`
	// Rates of the above counters.
	PacketRate           *uint64  `protobuf:"varint,3,opt,name=packet_rate,json=packetRate" json:"packet_rate,omitempty"`
	ByteRate             *uint64  `protobuf:"varint,4,opt,name=byte_rate,json=byteRate" json:"byte_rate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SegmentRoutingIfStats) Reset()         { *m = SegmentRoutingIfStats{} }
func (m *SegmentRoutingIfStats) String() string { return proto.CompactTextString(m) }
func (*SegmentRoutingIfStats) ProtoMessage()    {}
func (*SegmentRoutingIfStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_sr_stats_per_if_egress_9167b4ebc03e7ddd, []int{2}
}
func (m *SegmentRoutingIfStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SegmentRoutingIfStats.Unmarshal(m, b)
}
func (m *SegmentRoutingIfStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SegmentRoutingIfStats.Marshal(b, m, deterministic)
}
func (dst *SegmentRoutingIfStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SegmentRoutingIfStats.Merge(dst, src)
}
func (m *SegmentRoutingIfStats) XXX_Size() int {
	return xxx_messageInfo_SegmentRoutingIfStats.Size(m)
}
func (m *SegmentRoutingIfStats) XXX_DiscardUnknown() {
	xxx_messageInfo_SegmentRoutingIfStats.DiscardUnknown(m)
}

var xxx_messageInfo_SegmentRoutingIfStats proto.InternalMessageInfo

func (m *SegmentRoutingIfStats) GetPackets() uint64 {
	if m != nil && m.Packets != nil {
		return *m.Packets
	}
	return 0
}

func (m *SegmentRoutingIfStats) GetBytes() uint64 {
	if m != nil && m.Bytes != nil {
		return *m.Bytes
	}
	return 0
}

func (m *SegmentRoutingIfStats) GetPacketRate() uint64 {
	if m != nil && m.PacketRate != nil {
		return *m.PacketRate
	}
	return 0
}

func (m *SegmentRoutingIfStats) GetByteRate() uint64 {
	if m != nil && m.ByteRate != nil {
		return *m.ByteRate
	}
	return 0
}

var E_JnprSrStatsPerIfEgressExt = &proto.ExtensionDesc{
	ExtendedType:  (*telemetry_top.JuniperNetworksSensors)(nil),
	ExtensionType: (*SrStatsPerIfEgress)(nil),
	Field:         17,
	Name:          "jnpr_sr_stats_per_if_egress_ext",
	Tag:           "bytes,17,opt,name=jnpr_sr_stats_per_if_egress_ext,json=jnprSrStatsPerIfEgressExt",
	Filename:      "plugins/parsers/jtinative/sr_stats_per_if_egress/sr_stats_per_if_egress.proto",
}

func init() {
	proto.RegisterType((*SrStatsPerIfEgress)(nil), "SrStatsPerIfEgress")
	proto.RegisterType((*SegmentRoutingInterfaceRecord)(nil), "SegmentRoutingInterfaceRecord")
	proto.RegisterType((*SegmentRoutingIfStats)(nil), "SegmentRoutingIfStats")
	proto.RegisterExtension(E_JnprSrStatsPerIfEgressExt)
}

func init() {
	proto.RegisterFile("plugins/parsers/jtinative/sr_stats_per_if_egress/sr_stats_per_if_egress.proto", fileDescriptor_sr_stats_per_if_egress_9167b4ebc03e7ddd)
}

var fileDescriptor_sr_stats_per_if_egress_9167b4ebc03e7ddd = []byte{
	// 406 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x41, 0x6f, 0x13, 0x31,
	0x10, 0x85, 0xe5, 0x34, 0xa1, 0xed, 0xec, 0x52, 0x09, 0x57, 0xc0, 0x02, 0xa2, 0x5d, 0xe5, 0x80,
	0x56, 0x42, 0x4a, 0xa4, 0xde, 0x80, 0x0b, 0x20, 0x7a, 0x28, 0x12, 0x15, 0x72, 0x6e, 0x5c, 0x2c,
	0x13, 0x66, 0x23, 0xb7, 0x8d, 0xd7, 0x1a, 0xcf, 0x42, 0x7b, 0xe5, 0x07, 0xf0, 0x1f, 0xf8, 0xa7,
	0xc8, 0x76, 0xa0, 0x11, 0x6d, 0x73, 0xf4, 0x7b, 0xdf, 0x9b, 0xdd, 0x79, 0x1a, 0xf8, 0xe4, 0x2f,
	0xfa, 0x85, 0x75, 0x61, 0xea, 0x0d, 0x05, 0xa4, 0x30, 0x3d, 0x63, 0xeb, 0x0c, 0xdb, 0xef, 0x38,
	0x0d, 0xa4, 0x03, 0x1b, 0x0e, 0xda, 0x23, 0x69, 0xdb, 0x6a, 0x5c, 0x10, 0x86, 0x70, 0x87, 0x3c,
	0xf1, 0xd4, 0x71, 0xf7, 0x74, 0x9f, 0xf1, 0x02, 0x97, 0xc8, 0x74, 0xa5, 0xb9, 0xf3, 0x59, 0x1c,
	0x7f, 0x01, 0x39, 0xa3, 0x59, 0xcc, 0x7c, 0x46, 0x3a, 0x69, 0x8f, 0x53, 0x40, 0x7e, 0x80, 0xbd,
	0xd5, 0x04, 0xc2, 0x79, 0x47, 0xdf, 0x42, 0x25, 0xea, 0xad, 0xa6, 0x38, 0x3a, 0x98, 0xcc, 0x70,
	0xb1, 0x44, 0xc7, 0xaa, 0xeb, 0xd9, 0xba, 0xc5, 0x89, 0x63, 0xa4, 0xd6, 0xcc, 0x51, 0x25, 0x4c,
	0x95, 0x3e, 0x0e, 0xc9, 0x8f, 0x30, 0xfe, 0x35, 0x80, 0xe7, 0x1b, 0x79, 0x79, 0x00, 0xdb, 0xb6,
	0xd5, 0xce, 0x2c, 0xb1, 0x12, 0xf5, 0xa0, 0xd9, 0x7d, 0x3f, 0xfa, 0xf9, 0x76, 0xb0, 0x23, 0xd4,
	0x3d, 0xdb, 0x9e, 0x9a, 0x25, 0xca, 0x97, 0xb0, 0xe7, 0x0d, 0xa1, 0x63, 0x6d, 0x30, 0x63, 0x83,
	0x5a, 0x5c, 0x63, 0x65, 0x36, 0xdf, 0x61, 0x82, 0x1b, 0x28, 0xe7, 0x5d, 0x1f, 0xbf, 0x90, 0xd1,
	0xad, 0x75, 0xb4, 0x58, 0x59, 0x89, 0x7c, 0x03, 0xf7, 0xad, 0x4b, 0x9b, 0xe6, 0xba, 0xaa, 0x61,
	0x2d, 0x9a, 0xe2, 0xe8, 0xd1, 0xff, 0xdb, 0xb5, 0xa9, 0x18, 0x55, 0xae, 0xe0, 0xf4, 0x92, 0xaf,
	0xa0, 0xc4, 0xf5, 0xec, 0x68, 0x63, 0xb6, 0xc0, 0xeb, 0xe8, 0xf8, 0xb7, 0x80, 0x87, 0xb7, 0x62,
	0xf2, 0x10, 0xb6, 0xbd, 0x99, 0x9f, 0x23, 0xc7, 0xa6, 0x45, 0x33, 0x4c, 0xbf, 0x5d, 0x09, 0xf5,
	0x57, 0x95, 0xcf, 0x60, 0xf4, 0xf5, 0x8a, 0x31, 0xa4, 0x02, 0xfe, 0xd9, 0x59, 0x93, 0x2f, 0xa0,
	0xc8, 0x9c, 0x26, 0xc3, 0x79, 0xf1, 0x8c, 0xd4, 0x42, 0x41, 0x76, 0x94, 0x61, 0x94, 0x63, 0xd8,
	0x8d, 0x81, 0x4c, 0x0d, 0xd7, 0xa9, 0x9d, 0xa8, 0x47, 0xe6, 0xb5, 0x87, 0xc3, 0x33, 0xe7, 0x49,
	0xdf, 0x7e, 0x4a, 0x1a, 0x2f, 0x59, 0x3e, 0x9e, 0x7c, 0xec, 0x9d, 0xf5, 0x48, 0xa7, 0xc8, 0x3f,
	0x3a, 0x3a, 0x0f, 0x33, 0x74, 0xa1, 0xa3, 0x50, 0x3d, 0x48, 0x55, 0xec, 0x4f, 0x6e, 0x5e, 0x94,
	0x7a, 0x12, 0x87, 0xde, 0xd4, 0x8f, 0x2f, 0xf9, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x42, 0xfe,
	0x07, 0x31, 0xe7, 0x02, 0x00, 0x00,
}