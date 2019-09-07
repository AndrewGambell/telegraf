// Code generated by protoc-gen-go. DO NOT EDIT.
// source: plugins/parsers/jtinative/pfe_port_oc/pfe_port_oc.proto

package pfe_port_oc

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

type InterfacesPfePort struct {
	Interface            []*InterfacesPfePortInterfaceList `protobuf:"bytes,151,rep,name=interface" json:"interface,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *InterfacesPfePort) Reset()         { *m = InterfacesPfePort{} }
func (m *InterfacesPfePort) String() string { return proto.CompactTextString(m) }
func (*InterfacesPfePort) ProtoMessage()    {}
func (*InterfacesPfePort) Descriptor() ([]byte, []int) {
	return fileDescriptor_pfe_port_oc_84acb85e80652239, []int{0}
}
func (m *InterfacesPfePort) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InterfacesPfePort.Unmarshal(m, b)
}
func (m *InterfacesPfePort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InterfacesPfePort.Marshal(b, m, deterministic)
}
func (dst *InterfacesPfePort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterfacesPfePort.Merge(dst, src)
}
func (m *InterfacesPfePort) XXX_Size() int {
	return xxx_messageInfo_InterfacesPfePort.Size(m)
}
func (m *InterfacesPfePort) XXX_DiscardUnknown() {
	xxx_messageInfo_InterfacesPfePort.DiscardUnknown(m)
}

var xxx_messageInfo_InterfacesPfePort proto.InternalMessageInfo

func (m *InterfacesPfePort) GetInterface() []*InterfacesPfePortInterfaceList {
	if m != nil {
		return m.Interface
	}
	return nil
}

type InterfacesPfePortInterfaceList struct {
	State                *InterfacesPfePortInterfaceListStateType `protobuf:"bytes,151,opt,name=state" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                 `json:"-"`
	XXX_unrecognized     []byte                                   `json:"-"`
	XXX_sizecache        int32                                    `json:"-"`
}

func (m *InterfacesPfePortInterfaceList) Reset()         { *m = InterfacesPfePortInterfaceList{} }
func (m *InterfacesPfePortInterfaceList) String() string { return proto.CompactTextString(m) }
func (*InterfacesPfePortInterfaceList) ProtoMessage()    {}
func (*InterfacesPfePortInterfaceList) Descriptor() ([]byte, []int) {
	return fileDescriptor_pfe_port_oc_84acb85e80652239, []int{0, 0}
}
func (m *InterfacesPfePortInterfaceList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InterfacesPfePortInterfaceList.Unmarshal(m, b)
}
func (m *InterfacesPfePortInterfaceList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InterfacesPfePortInterfaceList.Marshal(b, m, deterministic)
}
func (dst *InterfacesPfePortInterfaceList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterfacesPfePortInterfaceList.Merge(dst, src)
}
func (m *InterfacesPfePortInterfaceList) XXX_Size() int {
	return xxx_messageInfo_InterfacesPfePortInterfaceList.Size(m)
}
func (m *InterfacesPfePortInterfaceList) XXX_DiscardUnknown() {
	xxx_messageInfo_InterfacesPfePortInterfaceList.DiscardUnknown(m)
}

var xxx_messageInfo_InterfacesPfePortInterfaceList proto.InternalMessageInfo

func (m *InterfacesPfePortInterfaceList) GetState() *InterfacesPfePortInterfaceListStateType {
	if m != nil {
		return m.State
	}
	return nil
}

type InterfacesPfePortInterfaceListStateType struct {
	Counters             *InterfacesPfePortInterfaceListStateTypeCountersType `protobuf:"bytes,151,opt,name=counters" json:"counters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                             `json:"-"`
	XXX_unrecognized     []byte                                               `json:"-"`
	XXX_sizecache        int32                                                `json:"-"`
}

func (m *InterfacesPfePortInterfaceListStateType) Reset() {
	*m = InterfacesPfePortInterfaceListStateType{}
}
func (m *InterfacesPfePortInterfaceListStateType) String() string { return proto.CompactTextString(m) }
func (*InterfacesPfePortInterfaceListStateType) ProtoMessage()    {}
func (*InterfacesPfePortInterfaceListStateType) Descriptor() ([]byte, []int) {
	return fileDescriptor_pfe_port_oc_84acb85e80652239, []int{0, 0, 0}
}
func (m *InterfacesPfePortInterfaceListStateType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InterfacesPfePortInterfaceListStateType.Unmarshal(m, b)
}
func (m *InterfacesPfePortInterfaceListStateType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InterfacesPfePortInterfaceListStateType.Marshal(b, m, deterministic)
}
func (dst *InterfacesPfePortInterfaceListStateType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterfacesPfePortInterfaceListStateType.Merge(dst, src)
}
func (m *InterfacesPfePortInterfaceListStateType) XXX_Size() int {
	return xxx_messageInfo_InterfacesPfePortInterfaceListStateType.Size(m)
}
func (m *InterfacesPfePortInterfaceListStateType) XXX_DiscardUnknown() {
	xxx_messageInfo_InterfacesPfePortInterfaceListStateType.DiscardUnknown(m)
}

var xxx_messageInfo_InterfacesPfePortInterfaceListStateType proto.InternalMessageInfo

func (m *InterfacesPfePortInterfaceListStateType) GetCounters() *InterfacesPfePortInterfaceListStateTypeCountersType {
	if m != nil {
		return m.Counters
	}
	return nil
}

type InterfacesPfePortInterfaceListStateTypeCountersType struct {
	InOctets             *uint64  `protobuf:"varint,51,opt,name=in_octets,json=inOctets" json:"in_octets,omitempty"`
	InUnicastPkts        *uint64  `protobuf:"varint,52,opt,name=in_unicast_pkts,json=inUnicastPkts" json:"in_unicast_pkts,omitempty"`
	InBroadcastPkts      *uint64  `protobuf:"varint,53,opt,name=in_broadcast_pkts,json=inBroadcastPkts" json:"in_broadcast_pkts,omitempty"`
	InMulticastPkts      *uint64  `protobuf:"varint,54,opt,name=in_multicast_pkts,json=inMulticastPkts" json:"in_multicast_pkts,omitempty"`
	InDiscards           *uint64  `protobuf:"varint,55,opt,name=in_discards,json=inDiscards" json:"in_discards,omitempty"`
	InErrors             *uint64  `protobuf:"varint,56,opt,name=in_errors,json=inErrors" json:"in_errors,omitempty"`
	InUnknownProtos      *uint32  `protobuf:"varint,57,opt,name=in_unknown_protos,json=inUnknownProtos" json:"in_unknown_protos,omitempty"`
	OutOctets            *uint64  `protobuf:"varint,58,opt,name=out_octets,json=outOctets" json:"out_octets,omitempty"`
	OutUnicastPkts       *uint64  `protobuf:"varint,59,opt,name=out_unicast_pkts,json=outUnicastPkts" json:"out_unicast_pkts,omitempty"`
	OutBroadcastPkts     *uint64  `protobuf:"varint,60,opt,name=out_broadcast_pkts,json=outBroadcastPkts" json:"out_broadcast_pkts,omitempty"`
	OutMulticastPkts     *uint64  `protobuf:"varint,61,opt,name=out_multicast_pkts,json=outMulticastPkts" json:"out_multicast_pkts,omitempty"`
	OutDiscards          *uint64  `protobuf:"varint,62,opt,name=out_discards,json=outDiscards" json:"out_discards,omitempty"`
	OutErrors            *uint64  `protobuf:"varint,63,opt,name=out_errors,json=outErrors" json:"out_errors,omitempty"`
	LastClear            *string  `protobuf:"bytes,64,opt,name=last_clear,json=lastClear" json:"last_clear,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InterfacesPfePortInterfaceListStateTypeCountersType) Reset() {
	*m = InterfacesPfePortInterfaceListStateTypeCountersType{}
}
func (m *InterfacesPfePortInterfaceListStateTypeCountersType) String() string {
	return proto.CompactTextString(m)
}
func (*InterfacesPfePortInterfaceListStateTypeCountersType) ProtoMessage() {}
func (*InterfacesPfePortInterfaceListStateTypeCountersType) Descriptor() ([]byte, []int) {
	return fileDescriptor_pfe_port_oc_84acb85e80652239, []int{0, 0, 0, 0}
}
func (m *InterfacesPfePortInterfaceListStateTypeCountersType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InterfacesPfePortInterfaceListStateTypeCountersType.Unmarshal(m, b)
}
func (m *InterfacesPfePortInterfaceListStateTypeCountersType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InterfacesPfePortInterfaceListStateTypeCountersType.Marshal(b, m, deterministic)
}
func (dst *InterfacesPfePortInterfaceListStateTypeCountersType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterfacesPfePortInterfaceListStateTypeCountersType.Merge(dst, src)
}
func (m *InterfacesPfePortInterfaceListStateTypeCountersType) XXX_Size() int {
	return xxx_messageInfo_InterfacesPfePortInterfaceListStateTypeCountersType.Size(m)
}
func (m *InterfacesPfePortInterfaceListStateTypeCountersType) XXX_DiscardUnknown() {
	xxx_messageInfo_InterfacesPfePortInterfaceListStateTypeCountersType.DiscardUnknown(m)
}

var xxx_messageInfo_InterfacesPfePortInterfaceListStateTypeCountersType proto.InternalMessageInfo

func (m *InterfacesPfePortInterfaceListStateTypeCountersType) GetInOctets() uint64 {
	if m != nil && m.InOctets != nil {
		return *m.InOctets
	}
	return 0
}

func (m *InterfacesPfePortInterfaceListStateTypeCountersType) GetInUnicastPkts() uint64 {
	if m != nil && m.InUnicastPkts != nil {
		return *m.InUnicastPkts
	}
	return 0
}

func (m *InterfacesPfePortInterfaceListStateTypeCountersType) GetInBroadcastPkts() uint64 {
	if m != nil && m.InBroadcastPkts != nil {
		return *m.InBroadcastPkts
	}
	return 0
}

func (m *InterfacesPfePortInterfaceListStateTypeCountersType) GetInMulticastPkts() uint64 {
	if m != nil && m.InMulticastPkts != nil {
		return *m.InMulticastPkts
	}
	return 0
}

func (m *InterfacesPfePortInterfaceListStateTypeCountersType) GetInDiscards() uint64 {
	if m != nil && m.InDiscards != nil {
		return *m.InDiscards
	}
	return 0
}

func (m *InterfacesPfePortInterfaceListStateTypeCountersType) GetInErrors() uint64 {
	if m != nil && m.InErrors != nil {
		return *m.InErrors
	}
	return 0
}

func (m *InterfacesPfePortInterfaceListStateTypeCountersType) GetInUnknownProtos() uint32 {
	if m != nil && m.InUnknownProtos != nil {
		return *m.InUnknownProtos
	}
	return 0
}

func (m *InterfacesPfePortInterfaceListStateTypeCountersType) GetOutOctets() uint64 {
	if m != nil && m.OutOctets != nil {
		return *m.OutOctets
	}
	return 0
}

func (m *InterfacesPfePortInterfaceListStateTypeCountersType) GetOutUnicastPkts() uint64 {
	if m != nil && m.OutUnicastPkts != nil {
		return *m.OutUnicastPkts
	}
	return 0
}

func (m *InterfacesPfePortInterfaceListStateTypeCountersType) GetOutBroadcastPkts() uint64 {
	if m != nil && m.OutBroadcastPkts != nil {
		return *m.OutBroadcastPkts
	}
	return 0
}

func (m *InterfacesPfePortInterfaceListStateTypeCountersType) GetOutMulticastPkts() uint64 {
	if m != nil && m.OutMulticastPkts != nil {
		return *m.OutMulticastPkts
	}
	return 0
}

func (m *InterfacesPfePortInterfaceListStateTypeCountersType) GetOutDiscards() uint64 {
	if m != nil && m.OutDiscards != nil {
		return *m.OutDiscards
	}
	return 0
}

func (m *InterfacesPfePortInterfaceListStateTypeCountersType) GetOutErrors() uint64 {
	if m != nil && m.OutErrors != nil {
		return *m.OutErrors
	}
	return 0
}

func (m *InterfacesPfePortInterfaceListStateTypeCountersType) GetLastClear() string {
	if m != nil && m.LastClear != nil {
		return *m.LastClear
	}
	return ""
}

var E_JnprInterfacesPfePortExt = &proto.ExtensionDesc{
	ExtendedType:  (*telemetry_top.JuniperNetworksSensors)(nil),
	ExtensionType: (*InterfacesPfePort)(nil),
	Field:         60,
	Name:          "jnpr_interfaces_pfe_port_ext",
	Tag:           "bytes,60,opt,name=jnpr_interfaces_pfe_port_ext,json=jnprInterfacesPfePortExt",
	Filename:      "plugins/parsers/jtinative/pfe_port_oc/pfe_port_oc.proto",
}

func init() {
	proto.RegisterType((*InterfacesPfePort)(nil), "interfaces_pfe_port")
	proto.RegisterType((*InterfacesPfePortInterfaceList)(nil), "interfaces_pfe_port.interface_list")
	proto.RegisterType((*InterfacesPfePortInterfaceListStateType)(nil), "interfaces_pfe_port.interface_list.state_type")
	proto.RegisterType((*InterfacesPfePortInterfaceListStateTypeCountersType)(nil), "interfaces_pfe_port.interface_list.state_type.counters_type")
	proto.RegisterExtension(E_JnprInterfacesPfePortExt)
}

func init() {
	proto.RegisterFile("plugins/parsers/jtinative/pfe_port_oc/pfe_port_oc.proto", fileDescriptor_pfe_port_oc_84acb85e80652239)
}

var fileDescriptor_pfe_port_oc_84acb85e80652239 = []byte{
	// 502 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcb, 0x6e, 0x13, 0x31,
	0x14, 0x86, 0x35, 0xa2, 0x41, 0xcd, 0x09, 0x69, 0xc1, 0x45, 0x62, 0x14, 0xa8, 0x08, 0x20, 0xa1,
	0x11, 0x42, 0x13, 0xa9, 0x5c, 0x0a, 0x25, 0x5c, 0x54, 0xc8, 0x02, 0x24, 0x20, 0x1a, 0xd4, 0x05,
	0x2b, 0x6b, 0x98, 0x38, 0xc8, 0xcd, 0xc4, 0xb6, 0xec, 0x63, 0xda, 0xbe, 0x00, 0xaf, 0xc0, 0x86,
	0x47, 0xe1, 0x61, 0x78, 0x14, 0x64, 0xcf, 0x25, 0x33, 0x55, 0x16, 0x74, 0x67, 0xff, 0xe7, 0x3b,
	0xff, 0xb1, 0x7f, 0x1d, 0xd8, 0x57, 0xb9, 0xfd, 0xce, 0x85, 0x19, 0xa9, 0x54, 0x1b, 0xa6, 0xcd,
	0xe8, 0x18, 0xb9, 0x48, 0x91, 0xff, 0x60, 0x23, 0x35, 0x67, 0x54, 0x49, 0x8d, 0x54, 0x66, 0xcd,
	0x73, 0xac, 0xb4, 0x44, 0x39, 0xd8, 0x41, 0x96, 0xb3, 0x25, 0x43, 0x7d, 0x46, 0x51, 0xaa, 0x42,
	0xbc, 0xfb, 0xe7, 0x32, 0xec, 0x70, 0x81, 0x4c, 0xcf, 0xd3, 0x8c, 0x19, 0x5a, 0x75, 0x91, 0x43,
	0xe8, 0xd6, 0x72, 0xf8, 0x2b, 0x18, 0x5e, 0x8a, 0x7a, 0x7b, 0xf7, 0xe2, 0x35, 0xe4, 0x4a, 0xa3,
	0x39, 0x37, 0x98, 0xac, 0xda, 0x06, 0x7f, 0x3b, 0xb0, 0xd5, 0xae, 0x92, 0x09, 0x74, 0x0c, 0xa6,
	0xe8, 0x2d, 0x83, 0xa8, 0xb7, 0x17, 0xff, 0x87, 0x65, 0xec, 0x3b, 0x28, 0x9e, 0x29, 0x96, 0x14,
	0xdd, 0x83, 0x9f, 0x1d, 0x80, 0x95, 0x4a, 0xbe, 0xc2, 0x66, 0x26, 0xad, 0x6b, 0x32, 0x95, 0xf1,
	0xf8, 0x62, 0xc6, 0x71, 0xd5, 0x5f, 0x8c, 0xa9, 0xed, 0x06, 0xbf, 0x37, 0xa0, 0xdf, 0xaa, 0x91,
	0x9b, 0x2e, 0x19, 0x2a, 0x33, 0x64, 0x68, 0xc2, 0x47, 0xc3, 0x20, 0xda, 0x48, 0x36, 0xb9, 0xf8,
	0xec, 0xef, 0xe4, 0x3e, 0x6c, 0x73, 0x41, 0xad, 0xe0, 0x59, 0x6a, 0x90, 0xaa, 0x05, 0x9a, 0xf0,
	0xb1, 0x47, 0xfa, 0x5c, 0x1c, 0x15, 0xea, 0x74, 0x81, 0x86, 0x3c, 0x80, 0x6b, 0x5c, 0xd0, 0x6f,
	0x5a, 0xa6, 0xb3, 0x15, 0xf9, 0xc4, 0x93, 0xdb, 0x5c, 0x1c, 0x56, 0x7a, 0x83, 0x5d, 0xda, 0x1c,
	0x1b, 0xae, 0x4f, 0x2b, 0xf6, 0x63, 0xa5, 0x7b, 0xf6, 0x36, 0xf4, 0xb8, 0xa0, 0x33, 0x6e, 0xb2,
	0x54, 0xcf, 0x4c, 0xb8, 0xef, 0x29, 0xe0, 0xe2, 0x5d, 0xa9, 0x94, 0xaf, 0x67, 0x5a, 0x4b, 0x6d,
	0xc2, 0x67, 0xd5, 0xeb, 0x27, 0xfe, 0x5e, 0x4e, 0xb2, 0x62, 0x21, 0xe4, 0x89, 0xa0, 0x7e, 0x41,
	0x4c, 0xf8, 0x7c, 0x18, 0x44, 0x7d, 0x37, 0xe9, 0xa8, 0xd0, 0xa7, 0x5e, 0x26, 0xbb, 0x00, 0xd2,
	0x62, 0x95, 0xc3, 0x81, 0x77, 0xea, 0x4a, 0x8b, 0x65, 0x10, 0x11, 0x5c, 0x75, 0xe5, 0x56, 0x12,
	0x2f, 0x3c, 0xb4, 0x25, 0x2d, 0x36, 0xa3, 0x78, 0x08, 0xc4, 0x91, 0xe7, 0xb2, 0x18, 0x7b, 0xd6,
	0x79, 0xb4, 0xc3, 0x28, 0xe9, 0x73, 0x69, 0xbc, 0xac, 0xe9, 0x76, 0x1c, 0x77, 0xe0, 0x8a, 0xa3,
	0xeb, 0x3c, 0x5e, 0x79, 0xae, 0x27, 0x2d, 0xd6, 0x81, 0x94, 0xff, 0x28, 0x13, 0x79, 0x5d, 0xff,
	0xa3, 0x8c, 0x64, 0x17, 0x20, 0x77, 0x63, 0xb2, 0x9c, 0xa5, 0x3a, 0x7c, 0x33, 0x0c, 0xa2, 0x6e,
	0xd2, 0x75, 0xca, 0x5b, 0x27, 0x1c, 0x2c, 0xe1, 0xd6, 0xb1, 0x50, 0x9a, 0xae, 0x59, 0x36, 0xca,
	0x4e, 0x91, 0xdc, 0x88, 0x3f, 0x58, 0xc1, 0x15, 0xd3, 0x9f, 0x18, 0x9e, 0x48, 0xbd, 0x30, 0x5f,
	0x98, 0x30, 0x6e, 0xd0, 0xd8, 0x6f, 0xe9, 0xf5, 0x75, 0x5b, 0x9a, 0x84, 0xce, 0xf2, 0x7d, 0x5d,
	0x98, 0xce, 0xd9, 0x54, 0x6a, 0x9c, 0x9c, 0xe2, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x40,
	0xfe, 0xb2, 0xfc, 0x03, 0x00, 0x00,
}