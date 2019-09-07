// Code generated by protoc-gen-go. DO NOT EDIT.
// source: plugins/parsers/jtinative/eventd/eventd.proto

package eventd

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

type JunosEvents struct {
	Events               *JunosEventsEventsType `protobuf:"bytes,151,opt,name=events" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *JunosEvents) Reset()         { *m = JunosEvents{} }
func (m *JunosEvents) String() string { return proto.CompactTextString(m) }
func (*JunosEvents) ProtoMessage()    {}
func (*JunosEvents) Descriptor() ([]byte, []int) {
	return fileDescriptor_eventd_fc33bd6a082269fc, []int{0}
}
func (m *JunosEvents) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JunosEvents.Unmarshal(m, b)
}
func (m *JunosEvents) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JunosEvents.Marshal(b, m, deterministic)
}
func (dst *JunosEvents) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JunosEvents.Merge(dst, src)
}
func (m *JunosEvents) XXX_Size() int {
	return xxx_messageInfo_JunosEvents.Size(m)
}
func (m *JunosEvents) XXX_DiscardUnknown() {
	xxx_messageInfo_JunosEvents.DiscardUnknown(m)
}

var xxx_messageInfo_JunosEvents proto.InternalMessageInfo

func (m *JunosEvents) GetEvents() *JunosEventsEventsType {
	if m != nil {
		return m.Events
	}
	return nil
}

type JunosEventsEventsType struct {
	Event                []*JunosEventsEventsTypeEventList `protobuf:"bytes,151,rep,name=event" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *JunosEventsEventsType) Reset()         { *m = JunosEventsEventsType{} }
func (m *JunosEventsEventsType) String() string { return proto.CompactTextString(m) }
func (*JunosEventsEventsType) ProtoMessage()    {}
func (*JunosEventsEventsType) Descriptor() ([]byte, []int) {
	return fileDescriptor_eventd_fc33bd6a082269fc, []int{0, 0}
}
func (m *JunosEventsEventsType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JunosEventsEventsType.Unmarshal(m, b)
}
func (m *JunosEventsEventsType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JunosEventsEventsType.Marshal(b, m, deterministic)
}
func (dst *JunosEventsEventsType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JunosEventsEventsType.Merge(dst, src)
}
func (m *JunosEventsEventsType) XXX_Size() int {
	return xxx_messageInfo_JunosEventsEventsType.Size(m)
}
func (m *JunosEventsEventsType) XXX_DiscardUnknown() {
	xxx_messageInfo_JunosEventsEventsType.DiscardUnknown(m)
}

var xxx_messageInfo_JunosEventsEventsType proto.InternalMessageInfo

func (m *JunosEventsEventsType) GetEvent() []*JunosEventsEventsTypeEventList {
	if m != nil {
		return m.Event
	}
	return nil
}

type JunosEventsEventsTypeEventList struct {
	Id                   *string                                         `protobuf:"bytes,51,opt,name=id" json:"id,omitempty"`
	Type                 *string                                         `protobuf:"bytes,52,opt,name=type" json:"type,omitempty"`
	Timestamp            *JunosEventsEventsTypeEventListTimestampType    `protobuf:"bytes,151,opt,name=timestamp" json:"timestamp,omitempty"`
	Priority             *string                                         `protobuf:"bytes,53,opt,name=priority" json:"priority,omitempty"`
	Facility             *string                                         `protobuf:"bytes,54,opt,name=facility" json:"facility,omitempty"`
	Pid                  *uint32                                         `protobuf:"varint,55,opt,name=pid" json:"pid,omitempty"`
	Message              *string                                         `protobuf:"bytes,56,opt,name=message" json:"message,omitempty"`
	Daemon               *string                                         `protobuf:"bytes,57,opt,name=daemon" json:"daemon,omitempty"`
	Hostname             *string                                         `protobuf:"bytes,58,opt,name=hostname" json:"hostname,omitempty"`
	Lsname               *string                                         `protobuf:"bytes,59,opt,name=lsname" json:"lsname,omitempty"`
	Attributes           []*JunosEventsEventsTypeEventListAttributesList `protobuf:"bytes,152,rep,name=attributes" json:"attributes,omitempty"`
	Logoptions           *int32                                          `protobuf:"varint,60,opt,name=logoptions" json:"logoptions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                        `json:"-"`
	XXX_unrecognized     []byte                                          `json:"-"`
	XXX_sizecache        int32                                           `json:"-"`
}

func (m *JunosEventsEventsTypeEventList) Reset()         { *m = JunosEventsEventsTypeEventList{} }
func (m *JunosEventsEventsTypeEventList) String() string { return proto.CompactTextString(m) }
func (*JunosEventsEventsTypeEventList) ProtoMessage()    {}
func (*JunosEventsEventsTypeEventList) Descriptor() ([]byte, []int) {
	return fileDescriptor_eventd_fc33bd6a082269fc, []int{0, 0, 0}
}
func (m *JunosEventsEventsTypeEventList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JunosEventsEventsTypeEventList.Unmarshal(m, b)
}
func (m *JunosEventsEventsTypeEventList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JunosEventsEventsTypeEventList.Marshal(b, m, deterministic)
}
func (dst *JunosEventsEventsTypeEventList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JunosEventsEventsTypeEventList.Merge(dst, src)
}
func (m *JunosEventsEventsTypeEventList) XXX_Size() int {
	return xxx_messageInfo_JunosEventsEventsTypeEventList.Size(m)
}
func (m *JunosEventsEventsTypeEventList) XXX_DiscardUnknown() {
	xxx_messageInfo_JunosEventsEventsTypeEventList.DiscardUnknown(m)
}

var xxx_messageInfo_JunosEventsEventsTypeEventList proto.InternalMessageInfo

func (m *JunosEventsEventsTypeEventList) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *JunosEventsEventsTypeEventList) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *JunosEventsEventsTypeEventList) GetTimestamp() *JunosEventsEventsTypeEventListTimestampType {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *JunosEventsEventsTypeEventList) GetPriority() string {
	if m != nil && m.Priority != nil {
		return *m.Priority
	}
	return ""
}

func (m *JunosEventsEventsTypeEventList) GetFacility() string {
	if m != nil && m.Facility != nil {
		return *m.Facility
	}
	return ""
}

func (m *JunosEventsEventsTypeEventList) GetPid() uint32 {
	if m != nil && m.Pid != nil {
		return *m.Pid
	}
	return 0
}

func (m *JunosEventsEventsTypeEventList) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *JunosEventsEventsTypeEventList) GetDaemon() string {
	if m != nil && m.Daemon != nil {
		return *m.Daemon
	}
	return ""
}

func (m *JunosEventsEventsTypeEventList) GetHostname() string {
	if m != nil && m.Hostname != nil {
		return *m.Hostname
	}
	return ""
}

func (m *JunosEventsEventsTypeEventList) GetLsname() string {
	if m != nil && m.Lsname != nil {
		return *m.Lsname
	}
	return ""
}

func (m *JunosEventsEventsTypeEventList) GetAttributes() []*JunosEventsEventsTypeEventListAttributesList {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *JunosEventsEventsTypeEventList) GetLogoptions() int32 {
	if m != nil && m.Logoptions != nil {
		return *m.Logoptions
	}
	return 0
}

type JunosEventsEventsTypeEventListTimestampType struct {
	Seconds              *uint32  `protobuf:"varint,51,opt,name=seconds" json:"seconds,omitempty"`
	Microseconds         *uint32  `protobuf:"varint,52,opt,name=microseconds" json:"microseconds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JunosEventsEventsTypeEventListTimestampType) Reset() {
	*m = JunosEventsEventsTypeEventListTimestampType{}
}
func (m *JunosEventsEventsTypeEventListTimestampType) String() string {
	return proto.CompactTextString(m)
}
func (*JunosEventsEventsTypeEventListTimestampType) ProtoMessage() {}
func (*JunosEventsEventsTypeEventListTimestampType) Descriptor() ([]byte, []int) {
	return fileDescriptor_eventd_fc33bd6a082269fc, []int{0, 0, 0, 0}
}
func (m *JunosEventsEventsTypeEventListTimestampType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JunosEventsEventsTypeEventListTimestampType.Unmarshal(m, b)
}
func (m *JunosEventsEventsTypeEventListTimestampType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JunosEventsEventsTypeEventListTimestampType.Marshal(b, m, deterministic)
}
func (dst *JunosEventsEventsTypeEventListTimestampType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JunosEventsEventsTypeEventListTimestampType.Merge(dst, src)
}
func (m *JunosEventsEventsTypeEventListTimestampType) XXX_Size() int {
	return xxx_messageInfo_JunosEventsEventsTypeEventListTimestampType.Size(m)
}
func (m *JunosEventsEventsTypeEventListTimestampType) XXX_DiscardUnknown() {
	xxx_messageInfo_JunosEventsEventsTypeEventListTimestampType.DiscardUnknown(m)
}

var xxx_messageInfo_JunosEventsEventsTypeEventListTimestampType proto.InternalMessageInfo

func (m *JunosEventsEventsTypeEventListTimestampType) GetSeconds() uint32 {
	if m != nil && m.Seconds != nil {
		return *m.Seconds
	}
	return 0
}

func (m *JunosEventsEventsTypeEventListTimestampType) GetMicroseconds() uint32 {
	if m != nil && m.Microseconds != nil {
		return *m.Microseconds
	}
	return 0
}

type JunosEventsEventsTypeEventListAttributesList struct {
	Key                  *string  `protobuf:"bytes,51,opt,name=key" json:"key,omitempty"`
	Value                *string  `protobuf:"bytes,52,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JunosEventsEventsTypeEventListAttributesList) Reset() {
	*m = JunosEventsEventsTypeEventListAttributesList{}
}
func (m *JunosEventsEventsTypeEventListAttributesList) String() string {
	return proto.CompactTextString(m)
}
func (*JunosEventsEventsTypeEventListAttributesList) ProtoMessage() {}
func (*JunosEventsEventsTypeEventListAttributesList) Descriptor() ([]byte, []int) {
	return fileDescriptor_eventd_fc33bd6a082269fc, []int{0, 0, 0, 1}
}
func (m *JunosEventsEventsTypeEventListAttributesList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JunosEventsEventsTypeEventListAttributesList.Unmarshal(m, b)
}
func (m *JunosEventsEventsTypeEventListAttributesList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JunosEventsEventsTypeEventListAttributesList.Marshal(b, m, deterministic)
}
func (dst *JunosEventsEventsTypeEventListAttributesList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JunosEventsEventsTypeEventListAttributesList.Merge(dst, src)
}
func (m *JunosEventsEventsTypeEventListAttributesList) XXX_Size() int {
	return xxx_messageInfo_JunosEventsEventsTypeEventListAttributesList.Size(m)
}
func (m *JunosEventsEventsTypeEventListAttributesList) XXX_DiscardUnknown() {
	xxx_messageInfo_JunosEventsEventsTypeEventListAttributesList.DiscardUnknown(m)
}

var xxx_messageInfo_JunosEventsEventsTypeEventListAttributesList proto.InternalMessageInfo

func (m *JunosEventsEventsTypeEventListAttributesList) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *JunosEventsEventsTypeEventListAttributesList) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

var E_JnprJunosEventsExt = &proto.ExtensionDesc{
	ExtendedType:  (*telemetry_top.JuniperNetworksSensors)(nil),
	ExtensionType: (*JunosEvents)(nil),
	Field:         42,
	Name:          "jnpr_junos_events_ext",
	Tag:           "bytes,42,opt,name=jnpr_junos_events_ext,json=jnprJunosEventsExt",
	Filename:      "plugins/parsers/jtinative/eventd/eventd.proto",
}

func init() {
	proto.RegisterType((*JunosEvents)(nil), "junos_events")
	proto.RegisterType((*JunosEventsEventsType)(nil), "junos_events.events_type")
	proto.RegisterType((*JunosEventsEventsTypeEventList)(nil), "junos_events.events_type.event_list")
	proto.RegisterType((*JunosEventsEventsTypeEventListTimestampType)(nil), "junos_events.events_type.event_list.timestamp_type")
	proto.RegisterType((*JunosEventsEventsTypeEventListAttributesList)(nil), "junos_events.events_type.event_list.attributes_list")
	proto.RegisterExtension(E_JnprJunosEventsExt)
}

func init() {
	proto.RegisterFile("plugins/parsers/jtinative/eventd/eventd.proto", fileDescriptor_eventd_fc33bd6a082269fc)
}

var fileDescriptor_eventd_fc33bd6a082269fc = []byte{
	// 473 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcd, 0x6e, 0x13, 0x3d,
	0x14, 0xd5, 0x24, 0x9d, 0x7e, 0xed, 0x4d, 0xd3, 0x0f, 0x99, 0x9f, 0x9a, 0x59, 0xa0, 0xa8, 0x62,
	0x11, 0x21, 0x91, 0x48, 0x69, 0xf8, 0x6b, 0x59, 0x54, 0x48, 0xdd, 0x74, 0xd1, 0xc5, 0x20, 0x16,
	0xac, 0x46, 0x43, 0xe6, 0x12, 0x9c, 0xce, 0xd8, 0x96, 0x7d, 0x27, 0x34, 0x5b, 0x96, 0xbc, 0x00,
	0xbc, 0x0c, 0xcf, 0xc0, 0x2b, 0x21, 0xdb, 0x99, 0x66, 0x82, 0x84, 0xd4, 0x95, 0xef, 0x39, 0xc7,
	0xe7, 0xd8, 0xbe, 0xbe, 0xf0, 0x5c, 0x97, 0xf5, 0x5c, 0x48, 0x3b, 0xd6, 0xb9, 0xb1, 0x68, 0xec,
	0x78, 0x41, 0x42, 0xe6, 0x24, 0x96, 0x38, 0xc6, 0x25, 0x4a, 0x2a, 0xd6, 0xcb, 0x48, 0x1b, 0x45,
	0x2a, 0xb9, 0x4f, 0x58, 0x62, 0x85, 0x64, 0x56, 0x19, 0x29, 0x1d, 0xc8, 0xe3, 0xdf, 0x31, 0x1c,
	0x2c, 0x6a, 0xa9, 0x6c, 0xe6, 0xf7, 0x5a, 0x36, 0x81, 0xdd, 0x50, 0xf1, 0x1f, 0xd1, 0x20, 0x1a,
	0xf6, 0x26, 0x8f, 0x47, 0x6d, 0x7d, 0x14, 0x96, 0x8c, 0x56, 0x1a, 0xd3, 0xf5, 0xce, 0xe4, 0x7b,
	0x0c, 0xbd, 0x16, 0xcf, 0xce, 0x20, 0xf6, 0xd0, 0x45, 0x74, 0x87, 0xbd, 0xc9, 0xd3, 0x7f, 0x46,
	0x84, 0x3a, 0x2b, 0x85, 0xa5, 0x34, 0x78, 0x92, 0x5f, 0x3b, 0x00, 0x1b, 0x96, 0x1d, 0x42, 0x47,
	0x14, 0xfc, 0x64, 0x10, 0x0d, 0xf7, 0xd3, 0x8e, 0x28, 0x18, 0x83, 0x1d, 0x67, 0xe4, 0x53, 0xcf,
	0xf8, 0x9a, 0xa5, 0xb0, 0x4f, 0xa2, 0x42, 0x4b, 0x79, 0xa5, 0x9b, 0x6b, 0x9f, 0xdc, 0xe5, 0xcc,
	0xd1, 0xad, 0x2d, 0x3c, 0x68, 0x13, 0xc3, 0x12, 0xd8, 0xd3, 0x46, 0x28, 0x23, 0x68, 0xc5, 0x5f,
	0xf8, 0xb3, 0x6e, 0xb1, 0xd3, 0x3e, 0xe7, 0x33, 0x51, 0x3a, 0xed, 0x65, 0xd0, 0x1a, 0xcc, 0xee,
	0x41, 0x57, 0x8b, 0x82, 0xbf, 0x1a, 0x44, 0xc3, 0x7e, 0xea, 0x4a, 0xc6, 0xe1, 0xbf, 0x0a, 0xad,
	0xcd, 0xe7, 0xc8, 0x5f, 0xfb, 0xcd, 0x0d, 0x64, 0x8f, 0x60, 0xb7, 0xc8, 0xb1, 0x52, 0x92, 0xbf,
	0xf1, 0xc2, 0x1a, 0xb9, 0xfc, 0x2f, 0xca, 0x92, 0xcc, 0x2b, 0xe4, 0xa7, 0x21, 0xbf, 0xc1, 0xce,
	0x53, 0x5a, 0xaf, 0x9c, 0x05, 0x4f, 0x40, 0xec, 0x03, 0x40, 0x4e, 0x64, 0xc4, 0xa7, 0x9a, 0xd0,
	0xf2, 0x9f, 0xa1, 0xf1, 0xd3, 0x3b, 0x35, 0x61, 0xe3, 0x0b, 0x1f, 0xd1, 0x0a, 0x62, 0x4f, 0x00,
	0x4a, 0x35, 0x57, 0x9a, 0x84, 0x92, 0x96, 0xbf, 0x1d, 0x44, 0xc3, 0x38, 0x6d, 0x31, 0xc9, 0x15,
	0x1c, 0x6e, 0xf7, 0xd0, 0x3d, 0xd7, 0xe2, 0x4c, 0xc9, 0xc2, 0xfa, 0x5f, 0xeb, 0xa7, 0x0d, 0x64,
	0xc7, 0x70, 0x50, 0x89, 0x99, 0x51, 0x8d, 0x3c, 0xf5, 0xf2, 0x16, 0x97, 0x9c, 0xc3, 0xff, 0x7f,
	0x5d, 0x87, 0x1d, 0x41, 0xf7, 0x1a, 0x57, 0x61, 0x04, 0xde, 0xc5, 0xdf, 0xce, 0x3b, 0x7b, 0x51,
	0xea, 0x18, 0xf6, 0x00, 0xe2, 0x65, 0x5e, 0xd6, 0xcd, 0x2c, 0x04, 0x70, 0xfa, 0x11, 0x1e, 0x2e,
	0xa4, 0x36, 0x59, 0xfb, 0xe5, 0x19, 0xde, 0x10, 0x3b, 0x1a, 0x5d, 0xd6, 0x52, 0x68, 0x34, 0x57,
	0x48, 0x5f, 0x95, 0xb9, 0xb6, 0xef, 0x51, 0x5a, 0x65, 0x2c, 0x7f, 0xe6, 0x07, 0xa6, 0xbf, 0xd5,
	0xab, 0x94, 0xb9, 0x90, 0x4b, 0xc7, 0x5c, 0x78, 0xe2, 0xe2, 0x86, 0xfe, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x06, 0xc6, 0xf7, 0x7e, 0x71, 0x03, 0x00, 0x00,
}