// Code generated by protoc-gen-go. DO NOT EDIT.
// source: plugins/parsers/jtinative/mib2d_nd6_oc/mib2d_nd6_oc.proto

package mib2d_nd6_oc

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

type Nd6InformationMibNd6 struct {
	Ipv6                 *Nd6InformationMibNd6Ipv6Type `protobuf:"bytes,151,opt,name=ipv6" json:"ipv6,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *Nd6InformationMibNd6) Reset()         { *m = Nd6InformationMibNd6{} }
func (m *Nd6InformationMibNd6) String() string { return proto.CompactTextString(m) }
func (*Nd6InformationMibNd6) ProtoMessage()    {}
func (*Nd6InformationMibNd6) Descriptor() ([]byte, []int) {
	return fileDescriptor_mib2d_nd6_oc_151b2b27e2fb1104, []int{0}
}
func (m *Nd6InformationMibNd6) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nd6InformationMibNd6.Unmarshal(m, b)
}
func (m *Nd6InformationMibNd6) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nd6InformationMibNd6.Marshal(b, m, deterministic)
}
func (dst *Nd6InformationMibNd6) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nd6InformationMibNd6.Merge(dst, src)
}
func (m *Nd6InformationMibNd6) XXX_Size() int {
	return xxx_messageInfo_Nd6InformationMibNd6.Size(m)
}
func (m *Nd6InformationMibNd6) XXX_DiscardUnknown() {
	xxx_messageInfo_Nd6InformationMibNd6.DiscardUnknown(m)
}

var xxx_messageInfo_Nd6InformationMibNd6 proto.InternalMessageInfo

func (m *Nd6InformationMibNd6) GetIpv6() *Nd6InformationMibNd6Ipv6Type {
	if m != nil {
		return m.Ipv6
	}
	return nil
}

type Nd6InformationMibNd6Ipv6Type struct {
	Neighbors            *Nd6InformationMibNd6Ipv6TypeNeighborsType `protobuf:"bytes,151,opt,name=neighbors" json:"neighbors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                   `json:"-"`
	XXX_unrecognized     []byte                                     `json:"-"`
	XXX_sizecache        int32                                      `json:"-"`
}

func (m *Nd6InformationMibNd6Ipv6Type) Reset()         { *m = Nd6InformationMibNd6Ipv6Type{} }
func (m *Nd6InformationMibNd6Ipv6Type) String() string { return proto.CompactTextString(m) }
func (*Nd6InformationMibNd6Ipv6Type) ProtoMessage()    {}
func (*Nd6InformationMibNd6Ipv6Type) Descriptor() ([]byte, []int) {
	return fileDescriptor_mib2d_nd6_oc_151b2b27e2fb1104, []int{0, 0}
}
func (m *Nd6InformationMibNd6Ipv6Type) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nd6InformationMibNd6Ipv6Type.Unmarshal(m, b)
}
func (m *Nd6InformationMibNd6Ipv6Type) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nd6InformationMibNd6Ipv6Type.Marshal(b, m, deterministic)
}
func (dst *Nd6InformationMibNd6Ipv6Type) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nd6InformationMibNd6Ipv6Type.Merge(dst, src)
}
func (m *Nd6InformationMibNd6Ipv6Type) XXX_Size() int {
	return xxx_messageInfo_Nd6InformationMibNd6Ipv6Type.Size(m)
}
func (m *Nd6InformationMibNd6Ipv6Type) XXX_DiscardUnknown() {
	xxx_messageInfo_Nd6InformationMibNd6Ipv6Type.DiscardUnknown(m)
}

var xxx_messageInfo_Nd6InformationMibNd6Ipv6Type proto.InternalMessageInfo

func (m *Nd6InformationMibNd6Ipv6Type) GetNeighbors() *Nd6InformationMibNd6Ipv6TypeNeighborsType {
	if m != nil {
		return m.Neighbors
	}
	return nil
}

type Nd6InformationMibNd6Ipv6TypeNeighborsType struct {
	Neighbor             []*Nd6InformationMibNd6Ipv6TypeNeighborsTypeNeighborList `protobuf:"bytes,151,rep,name=neighbor" json:"neighbor,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                 `json:"-"`
	XXX_unrecognized     []byte                                                   `json:"-"`
	XXX_sizecache        int32                                                    `json:"-"`
}

func (m *Nd6InformationMibNd6Ipv6TypeNeighborsType) Reset() {
	*m = Nd6InformationMibNd6Ipv6TypeNeighborsType{}
}
func (m *Nd6InformationMibNd6Ipv6TypeNeighborsType) String() string { return proto.CompactTextString(m) }
func (*Nd6InformationMibNd6Ipv6TypeNeighborsType) ProtoMessage()    {}
func (*Nd6InformationMibNd6Ipv6TypeNeighborsType) Descriptor() ([]byte, []int) {
	return fileDescriptor_mib2d_nd6_oc_151b2b27e2fb1104, []int{0, 0, 0}
}
func (m *Nd6InformationMibNd6Ipv6TypeNeighborsType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nd6InformationMibNd6Ipv6TypeNeighborsType.Unmarshal(m, b)
}
func (m *Nd6InformationMibNd6Ipv6TypeNeighborsType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nd6InformationMibNd6Ipv6TypeNeighborsType.Marshal(b, m, deterministic)
}
func (dst *Nd6InformationMibNd6Ipv6TypeNeighborsType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nd6InformationMibNd6Ipv6TypeNeighborsType.Merge(dst, src)
}
func (m *Nd6InformationMibNd6Ipv6TypeNeighborsType) XXX_Size() int {
	return xxx_messageInfo_Nd6InformationMibNd6Ipv6TypeNeighborsType.Size(m)
}
func (m *Nd6InformationMibNd6Ipv6TypeNeighborsType) XXX_DiscardUnknown() {
	xxx_messageInfo_Nd6InformationMibNd6Ipv6TypeNeighborsType.DiscardUnknown(m)
}

var xxx_messageInfo_Nd6InformationMibNd6Ipv6TypeNeighborsType proto.InternalMessageInfo

func (m *Nd6InformationMibNd6Ipv6TypeNeighborsType) GetNeighbor() []*Nd6InformationMibNd6Ipv6TypeNeighborsTypeNeighborList {
	if m != nil {
		return m.Neighbor
	}
	return nil
}

type Nd6InformationMibNd6Ipv6TypeNeighborsTypeNeighborList struct {
	Ip                   *string                                                         `protobuf:"bytes,51,opt,name=ip" json:"ip,omitempty"`
	State                *Nd6InformationMibNd6Ipv6TypeNeighborsTypeNeighborListStateType `protobuf:"bytes,151,opt,name=state" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                        `json:"-"`
	XXX_unrecognized     []byte                                                          `json:"-"`
	XXX_sizecache        int32                                                           `json:"-"`
}

func (m *Nd6InformationMibNd6Ipv6TypeNeighborsTypeNeighborList) Reset() {
	*m = Nd6InformationMibNd6Ipv6TypeNeighborsTypeNeighborList{}
}
func (m *Nd6InformationMibNd6Ipv6TypeNeighborsTypeNeighborList) String() string {
	return proto.CompactTextString(m)
}
func (*Nd6InformationMibNd6Ipv6TypeNeighborsTypeNeighborList) ProtoMessage() {}
func (*Nd6InformationMibNd6Ipv6TypeNeighborsTypeNeighborList) Descriptor() ([]byte, []int) {
	return fileDescriptor_mib2d_nd6_oc_151b2b27e2fb1104, []int{0, 0, 0, 0}
}
func (m *Nd6InformationMibNd6Ipv6TypeNeighborsTypeNeighborList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nd6InformationMibNd6Ipv6TypeNeighborsTypeNeighborList.Unmarshal(m, b)
}
func (m *Nd6InformationMibNd6Ipv6TypeNeighborsTypeNeighborList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nd6InformationMibNd6Ipv6TypeNeighborsTypeNeighborList.Marshal(b, m, deterministic)
}
func (dst *Nd6InformationMibNd6Ipv6TypeNeighborsTypeNeighborList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nd6InformationMibNd6Ipv6TypeNeighborsTypeNeighborList.Merge(dst, src)
}
func (m *Nd6InformationMibNd6Ipv6TypeNeighborsTypeNeighborList) XXX_Size() int {
	return xxx_messageInfo_Nd6InformationMibNd6Ipv6TypeNeighborsTypeNeighborList.Size(m)
}
func (m *Nd6InformationMibNd6Ipv6TypeNeighborsTypeNeighborList) XXX_DiscardUnknown() {
	xxx_messageInfo_Nd6InformationMibNd6Ipv6TypeNeighborsTypeNeighborList.DiscardUnknown(m)
}

var xxx_messageInfo_Nd6InformationMibNd6Ipv6TypeNeighborsTypeNeighborList proto.InternalMessageInfo

func (m *Nd6InformationMibNd6Ipv6TypeNeighborsTypeNeighborList) GetIp() string {
	if m != nil && m.Ip != nil {
		return *m.Ip
	}
	return ""
}

func (m *Nd6InformationMibNd6Ipv6TypeNeighborsTypeNeighborList) GetState() *Nd6InformationMibNd6Ipv6TypeNeighborsTypeNeighborListStateType {
	if m != nil {
		return m.State
	}
	return nil
}

type Nd6InformationMibNd6Ipv6TypeNeighborsTypeNeighborListStateType struct {
	Ip                   *string  `protobuf:"bytes,51,opt,name=ip" json:"ip,omitempty"`
	LinkLayerAddress     *string  `protobuf:"bytes,52,opt,name=link_layer_address,json=linkLayerAddress" json:"link_layer_address,omitempty"`
	Origin               *string  `protobuf:"bytes,53,opt,name=origin" json:"origin,omitempty"`
	IsRouter             *string  `protobuf:"bytes,54,opt,name=is_router,json=isRouter" json:"is_router,omitempty"`
	NeighborState        *string  `protobuf:"bytes,55,opt,name=neighbor_state,json=neighborState" json:"neighbor_state,omitempty"`
	TableId              *uint32  `protobuf:"varint,61,opt,name=table_id,json=tableId" json:"table_id,omitempty"`
	IsSecure             *bool    `protobuf:"varint,62,opt,name=is_secure,json=isSecure" json:"is_secure,omitempty"`
	Expiry               *uint32  `protobuf:"varint,64,opt,name=expiry" json:"expiry,omitempty"`
	IsPublish            *bool    `protobuf:"varint,63,opt,name=is_publish,json=isPublish" json:"is_publish,omitempty"`
	InterfaceName        *string  `protobuf:"bytes,65,opt,name=interface_name,json=interfaceName" json:"interface_name,omitempty"`
	LogicalRouterId      *uint32  `protobuf:"varint,66,opt,name=logical_router_id,json=logicalRouterId" json:"logical_router_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Nd6InformationMibNd6Ipv6TypeNeighborsTypeNeighborListStateType) Reset() {
	*m = Nd6InformationMibNd6Ipv6TypeNeighborsTypeNeighborListStateType{}
}
func (m *Nd6InformationMibNd6Ipv6TypeNeighborsTypeNeighborListStateType) String() string {
	return proto.CompactTextString(m)
}
func (*Nd6InformationMibNd6Ipv6TypeNeighborsTypeNeighborListStateType) ProtoMessage() {}
func (*Nd6InformationMibNd6Ipv6TypeNeighborsTypeNeighborListStateType) Descriptor() ([]byte, []int) {
	return fileDescriptor_mib2d_nd6_oc_151b2b27e2fb1104, []int{0, 0, 0, 0, 0}
}
func (m *Nd6InformationMibNd6Ipv6TypeNeighborsTypeNeighborListStateType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nd6InformationMibNd6Ipv6TypeNeighborsTypeNeighborListStateType.Unmarshal(m, b)
}
func (m *Nd6InformationMibNd6Ipv6TypeNeighborsTypeNeighborListStateType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nd6InformationMibNd6Ipv6TypeNeighborsTypeNeighborListStateType.Marshal(b, m, deterministic)
}
func (dst *Nd6InformationMibNd6Ipv6TypeNeighborsTypeNeighborListStateType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nd6InformationMibNd6Ipv6TypeNeighborsTypeNeighborListStateType.Merge(dst, src)
}
func (m *Nd6InformationMibNd6Ipv6TypeNeighborsTypeNeighborListStateType) XXX_Size() int {
	return xxx_messageInfo_Nd6InformationMibNd6Ipv6TypeNeighborsTypeNeighborListStateType.Size(m)
}
func (m *Nd6InformationMibNd6Ipv6TypeNeighborsTypeNeighborListStateType) XXX_DiscardUnknown() {
	xxx_messageInfo_Nd6InformationMibNd6Ipv6TypeNeighborsTypeNeighborListStateType.DiscardUnknown(m)
}

var xxx_messageInfo_Nd6InformationMibNd6Ipv6TypeNeighborsTypeNeighborListStateType proto.InternalMessageInfo

func (m *Nd6InformationMibNd6Ipv6TypeNeighborsTypeNeighborListStateType) GetIp() string {
	if m != nil && m.Ip != nil {
		return *m.Ip
	}
	return ""
}

func (m *Nd6InformationMibNd6Ipv6TypeNeighborsTypeNeighborListStateType) GetLinkLayerAddress() string {
	if m != nil && m.LinkLayerAddress != nil {
		return *m.LinkLayerAddress
	}
	return ""
}

func (m *Nd6InformationMibNd6Ipv6TypeNeighborsTypeNeighborListStateType) GetOrigin() string {
	if m != nil && m.Origin != nil {
		return *m.Origin
	}
	return ""
}

func (m *Nd6InformationMibNd6Ipv6TypeNeighborsTypeNeighborListStateType) GetIsRouter() string {
	if m != nil && m.IsRouter != nil {
		return *m.IsRouter
	}
	return ""
}

func (m *Nd6InformationMibNd6Ipv6TypeNeighborsTypeNeighborListStateType) GetNeighborState() string {
	if m != nil && m.NeighborState != nil {
		return *m.NeighborState
	}
	return ""
}

func (m *Nd6InformationMibNd6Ipv6TypeNeighborsTypeNeighborListStateType) GetTableId() uint32 {
	if m != nil && m.TableId != nil {
		return *m.TableId
	}
	return 0
}

func (m *Nd6InformationMibNd6Ipv6TypeNeighborsTypeNeighborListStateType) GetIsSecure() bool {
	if m != nil && m.IsSecure != nil {
		return *m.IsSecure
	}
	return false
}

func (m *Nd6InformationMibNd6Ipv6TypeNeighborsTypeNeighborListStateType) GetExpiry() uint32 {
	if m != nil && m.Expiry != nil {
		return *m.Expiry
	}
	return 0
}

func (m *Nd6InformationMibNd6Ipv6TypeNeighborsTypeNeighborListStateType) GetIsPublish() bool {
	if m != nil && m.IsPublish != nil {
		return *m.IsPublish
	}
	return false
}

func (m *Nd6InformationMibNd6Ipv6TypeNeighborsTypeNeighborListStateType) GetInterfaceName() string {
	if m != nil && m.InterfaceName != nil {
		return *m.InterfaceName
	}
	return ""
}

func (m *Nd6InformationMibNd6Ipv6TypeNeighborsTypeNeighborListStateType) GetLogicalRouterId() uint32 {
	if m != nil && m.LogicalRouterId != nil {
		return *m.LogicalRouterId
	}
	return 0
}

var E_JnprNd6InformationMibNd6Ext = &proto.ExtensionDesc{
	ExtendedType:  (*telemetry_top.JuniperNetworksSensors)(nil),
	ExtensionType: (*Nd6InformationMibNd6)(nil),
	Field:         54,
	Name:          "jnpr_nd6_information_mib_nd6_ext",
	Tag:           "bytes,54,opt,name=jnpr_nd6_information_mib_nd6_ext,json=jnprNd6InformationMibNd6Ext",
	Filename:      "plugins/parsers/jtinative/mib2d_nd6_oc/mib2d_nd6_oc.proto",
}

func init() {
	proto.RegisterType((*Nd6InformationMibNd6)(nil), "nd6_information_mib_nd6")
	proto.RegisterType((*Nd6InformationMibNd6Ipv6Type)(nil), "nd6_information_mib_nd6.ipv6_type")
	proto.RegisterType((*Nd6InformationMibNd6Ipv6TypeNeighborsType)(nil), "nd6_information_mib_nd6.ipv6_type.neighbors_type")
	proto.RegisterType((*Nd6InformationMibNd6Ipv6TypeNeighborsTypeNeighborList)(nil), "nd6_information_mib_nd6.ipv6_type.neighbors_type.neighbor_list")
	proto.RegisterType((*Nd6InformationMibNd6Ipv6TypeNeighborsTypeNeighborListStateType)(nil), "nd6_information_mib_nd6.ipv6_type.neighbors_type.neighbor_list.state_type")
	proto.RegisterExtension(E_JnprNd6InformationMibNd6Ext)
}

func init() {
	proto.RegisterFile("plugins/parsers/jtinative/mib2d_nd6_oc/mib2d_nd6_oc.proto", fileDescriptor_mib2d_nd6_oc_151b2b27e2fb1104)
}

var fileDescriptor_mib2d_nd6_oc_151b2b27e2fb1104 = []byte{
	// 515 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0xdd, 0x6a, 0x13, 0x41,
	0x14, 0x66, 0xd3, 0x1f, 0x93, 0x29, 0xad, 0x3a, 0xa2, 0x1d, 0x53, 0x84, 0x50, 0x10, 0x82, 0xc8,
	0x16, 0xab, 0x6e, 0x51, 0xd0, 0xfe, 0x80, 0x17, 0x29, 0x1a, 0xca, 0xe6, 0x56, 0x18, 0x66, 0xb3,
	0xa7, 0xe9, 0x69, 0x77, 0x67, 0x86, 0x99, 0x49, 0x4d, 0x6e, 0x7d, 0x09, 0xc1, 0xf7, 0xf1, 0x51,
	0x04, 0x1f, 0xc0, 0x07, 0x90, 0x99, 0x4d, 0x36, 0xe4, 0x22, 0x88, 0x78, 0xf9, 0x7d, 0xdf, 0x39,
	0xdf, 0x7c, 0xe7, 0xcc, 0x21, 0x6f, 0x74, 0x31, 0x1e, 0xa1, 0xb4, 0x07, 0x5a, 0x18, 0x0b, 0xc6,
	0x1e, 0x5c, 0x3b, 0x94, 0xc2, 0xe1, 0x2d, 0x1c, 0x94, 0x98, 0x1d, 0xe6, 0x5c, 0xe6, 0x09, 0x57,
	0xc3, 0x25, 0x10, 0x6b, 0xa3, 0x9c, 0x6a, 0x3f, 0x70, 0x50, 0x40, 0x09, 0xce, 0x4c, 0xb9, 0x53,
	0xba, 0x22, 0xf7, 0xbf, 0x6f, 0x92, 0x5d, 0x5f, 0x85, 0xf2, 0x52, 0x99, 0x52, 0x38, 0x54, 0x92,
	0x97, 0x98, 0xf9, 0x4e, 0x7a, 0x44, 0xd6, 0x51, 0xdf, 0x26, 0xec, 0x5b, 0xd4, 0x89, 0xba, 0x5b,
	0x87, 0xfb, 0xf1, 0x8a, 0xc2, 0xd8, 0x57, 0x71, 0x37, 0xd5, 0x90, 0x86, 0x86, 0xf6, 0x8f, 0x0d,
	0xd2, 0xaa, 0x39, 0x7a, 0x41, 0x5a, 0x12, 0x70, 0x74, 0x95, 0x29, 0x63, 0xe7, 0x5e, 0x2f, 0xfe,
	0xee, 0x15, 0xd7, 0x4d, 0x95, 0xf5, 0xc2, 0xa4, 0xfd, 0x73, 0x9d, 0xec, 0x2c, 0xab, 0xf4, 0x33,
	0x69, 0xce, 0x19, 0xff, 0xc6, 0x5a, 0x77, 0xeb, 0xf0, 0xf8, 0x9f, 0xdf, 0xa8, 0x21, 0x2f, 0xd0,
	0xba, 0xb4, 0x76, 0x6c, 0xff, 0x5e, 0x23, 0xdb, 0x4b, 0x1a, 0x7d, 0x48, 0x1a, 0xa8, 0xd9, 0xcb,
	0x4e, 0xd4, 0x6d, 0x9d, 0x6d, 0x7c, 0x3d, 0x69, 0x34, 0xa3, 0xb4, 0x81, 0x9a, 0x0a, 0xb2, 0x61,
	0x9d, 0x70, 0x30, 0x9f, 0xf3, 0xfc, 0x3f, 0x33, 0xc4, 0xc1, 0xad, 0x5a, 0x40, 0xe5, 0xdc, 0xfe,
	0xd5, 0x20, 0x64, 0xc1, 0xd2, 0x9d, 0x45, 0x90, 0x90, 0xe0, 0x39, 0xa1, 0x05, 0xca, 0x1b, 0x5e,
	0x88, 0x29, 0x18, 0x2e, 0xf2, 0xdc, 0x80, 0xb5, 0xec, 0x55, 0xd0, 0xef, 0x79, 0xe5, 0xa3, 0x17,
	0x4e, 0x2b, 0x9e, 0x3e, 0x22, 0x9b, 0xca, 0xe0, 0x08, 0x25, 0x7b, 0x1d, 0x2a, 0x66, 0x88, 0xee,
	0x91, 0x16, 0x5a, 0x6e, 0xd4, 0xd8, 0x81, 0x61, 0x49, 0x90, 0x9a, 0x68, 0xd3, 0x80, 0xe9, 0xd3,
	0xc5, 0xf6, 0x79, 0x35, 0xed, 0x51, 0xa8, 0xa8, 0x57, 0x34, 0xf0, 0x24, 0x7d, 0x4c, 0x9a, 0x4e,
	0x64, 0x05, 0x70, 0xcc, 0xd9, 0xbb, 0x4e, 0xd4, 0xdd, 0x4e, 0xef, 0x04, 0xdc, 0xcb, 0x67, 0xf6,
	0x16, 0x86, 0x63, 0x03, 0xec, 0x7d, 0x27, 0xea, 0x36, 0xbd, 0xfd, 0x20, 0x60, 0x9f, 0x09, 0x26,
	0x1a, 0xcd, 0x94, 0x9d, 0x84, 0xae, 0x19, 0xa2, 0x4f, 0x08, 0x41, 0xcb, 0xf5, 0x38, 0x2b, 0xd0,
	0x5e, 0xb1, 0xe3, 0xd0, 0xd5, 0x42, 0x7b, 0x51, 0x11, 0x3e, 0x15, 0x4a, 0x07, 0xe6, 0x52, 0x0c,
	0x81, 0x4b, 0x51, 0x02, 0x3b, 0xad, 0x52, 0xd5, 0x6c, 0x5f, 0x94, 0x40, 0x9f, 0x91, 0xfb, 0x85,
	0x1a, 0xe1, 0x50, 0x14, 0xb3, 0xf1, 0x7c, 0xbc, 0xb3, 0xf0, 0xd0, 0xdd, 0x99, 0x50, 0x8d, 0xd9,
	0xcb, 0xdf, 0x4e, 0x48, 0xe7, 0x5a, 0x6a, 0xc3, 0x57, 0xfc, 0x21, 0x87, 0x89, 0xa3, 0xbb, 0xf1,
	0xf9, 0x58, 0xa2, 0x06, 0xd3, 0x07, 0xf7, 0x45, 0x99, 0x1b, 0x3b, 0x00, 0x69, 0xfd, 0xa5, 0x27,
	0xe1, 0x00, 0xd8, 0xaa, 0x03, 0x48, 0xf7, 0xbc, 0x75, 0x3f, 0x4f, 0x7a, 0x0b, 0xed, 0x13, 0x66,
	0xfd, 0x3c, 0xf9, 0x30, 0x71, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x11, 0xe6, 0xc6, 0x4d, 0xe7,
	0x03, 0x00, 0x00,
}
