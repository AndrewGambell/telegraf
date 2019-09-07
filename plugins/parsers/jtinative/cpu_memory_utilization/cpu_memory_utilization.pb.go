// Code generated by protoc-gen-go. DO NOT EDIT.
// source: plugins/parsers/jtinative/cpu_memory_utilization/cpu_memory_utilization.proto

package cpu_memory_utilization

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
// The top level message is CpuMemoryUtilization
//
type CpuMemoryUtilization struct {
	Utilization          []*CpuMemoryUtilizationSummary `protobuf:"bytes,1,rep,name=utilization" json:"utilization,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *CpuMemoryUtilization) Reset()         { *m = CpuMemoryUtilization{} }
func (m *CpuMemoryUtilization) String() string { return proto.CompactTextString(m) }
func (*CpuMemoryUtilization) ProtoMessage()    {}
func (*CpuMemoryUtilization) Descriptor() ([]byte, []int) {
	return fileDescriptor_cpu_memory_utilization_cdb8d37bc3bcd61e, []int{0}
}
func (m *CpuMemoryUtilization) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CpuMemoryUtilization.Unmarshal(m, b)
}
func (m *CpuMemoryUtilization) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CpuMemoryUtilization.Marshal(b, m, deterministic)
}
func (dst *CpuMemoryUtilization) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CpuMemoryUtilization.Merge(dst, src)
}
func (m *CpuMemoryUtilization) XXX_Size() int {
	return xxx_messageInfo_CpuMemoryUtilization.Size(m)
}
func (m *CpuMemoryUtilization) XXX_DiscardUnknown() {
	xxx_messageInfo_CpuMemoryUtilization.DiscardUnknown(m)
}

var xxx_messageInfo_CpuMemoryUtilization proto.InternalMessageInfo

func (m *CpuMemoryUtilization) GetUtilization() []*CpuMemoryUtilizationSummary {
	if m != nil {
		return m.Utilization
	}
	return nil
}

// This array gives the CPU  memory utilization on per partition basis
// and also the per application memory utilization for each partition
type CpuMemoryUtilizationSummary struct {
	// Name of the partition.
	Name *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The total size of the partition in bytes
	Size *uint64 `protobuf:"varint,2,opt,name=size" json:"size,omitempty"`
	// The amount of memory currently allocated from the partition in bytes
	BytesAllocated *uint64 `protobuf:"varint,3,opt,name=bytes_allocated,json=bytesAllocated" json:"bytes_allocated,omitempty"`
	// The amount of memory that is currently allocated, expressed
	// as percentage of the total (0--100).
	Utilization *int32 `protobuf:"varint,4,opt,name=utilization" json:"utilization,omitempty"`
	// Per application based memory utilization for this memory partition
	ApplicationUtilization []*CpuMemoryUtilizationPerApplication `protobuf:"bytes,5,rep,name=application_utilization,json=applicationUtilization" json:"application_utilization,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                              `json:"-"`
	XXX_unrecognized       []byte                                `json:"-"`
	XXX_sizecache          int32                                 `json:"-"`
}

func (m *CpuMemoryUtilizationSummary) Reset()         { *m = CpuMemoryUtilizationSummary{} }
func (m *CpuMemoryUtilizationSummary) String() string { return proto.CompactTextString(m) }
func (*CpuMemoryUtilizationSummary) ProtoMessage()    {}
func (*CpuMemoryUtilizationSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_cpu_memory_utilization_cdb8d37bc3bcd61e, []int{1}
}
func (m *CpuMemoryUtilizationSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CpuMemoryUtilizationSummary.Unmarshal(m, b)
}
func (m *CpuMemoryUtilizationSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CpuMemoryUtilizationSummary.Marshal(b, m, deterministic)
}
func (dst *CpuMemoryUtilizationSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CpuMemoryUtilizationSummary.Merge(dst, src)
}
func (m *CpuMemoryUtilizationSummary) XXX_Size() int {
	return xxx_messageInfo_CpuMemoryUtilizationSummary.Size(m)
}
func (m *CpuMemoryUtilizationSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_CpuMemoryUtilizationSummary.DiscardUnknown(m)
}

var xxx_messageInfo_CpuMemoryUtilizationSummary proto.InternalMessageInfo

func (m *CpuMemoryUtilizationSummary) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *CpuMemoryUtilizationSummary) GetSize() uint64 {
	if m != nil && m.Size != nil {
		return *m.Size
	}
	return 0
}

func (m *CpuMemoryUtilizationSummary) GetBytesAllocated() uint64 {
	if m != nil && m.BytesAllocated != nil {
		return *m.BytesAllocated
	}
	return 0
}

func (m *CpuMemoryUtilizationSummary) GetUtilization() int32 {
	if m != nil && m.Utilization != nil {
		return *m.Utilization
	}
	return 0
}

func (m *CpuMemoryUtilizationSummary) GetApplicationUtilization() []*CpuMemoryUtilizationPerApplication {
	if m != nil {
		return m.ApplicationUtilization
	}
	return nil
}

// This describes per Application specific CPU memory utilization
type CpuMemoryUtilizationPerApplication struct {
	// Application name
	Name *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Number of bytes allocated
	BytesAllocated *uint64 `protobuf:"varint,2,opt,name=bytes_allocated,json=bytesAllocated" json:"bytes_allocated,omitempty"`
	//  Number of allocations
	Allocations *uint64 `protobuf:"varint,3,opt,name=allocations" json:"allocations,omitempty"`
	//  Number of frees
	Frees *uint64 `protobuf:"varint,4,opt,name=frees" json:"frees,omitempty"`
	// Number of allocations failed
	AllocationsFailed    *uint64  `protobuf:"varint,5,opt,name=allocations_failed,json=allocationsFailed" json:"allocations_failed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CpuMemoryUtilizationPerApplication) Reset()         { *m = CpuMemoryUtilizationPerApplication{} }
func (m *CpuMemoryUtilizationPerApplication) String() string { return proto.CompactTextString(m) }
func (*CpuMemoryUtilizationPerApplication) ProtoMessage()    {}
func (*CpuMemoryUtilizationPerApplication) Descriptor() ([]byte, []int) {
	return fileDescriptor_cpu_memory_utilization_cdb8d37bc3bcd61e, []int{2}
}
func (m *CpuMemoryUtilizationPerApplication) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CpuMemoryUtilizationPerApplication.Unmarshal(m, b)
}
func (m *CpuMemoryUtilizationPerApplication) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CpuMemoryUtilizationPerApplication.Marshal(b, m, deterministic)
}
func (dst *CpuMemoryUtilizationPerApplication) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CpuMemoryUtilizationPerApplication.Merge(dst, src)
}
func (m *CpuMemoryUtilizationPerApplication) XXX_Size() int {
	return xxx_messageInfo_CpuMemoryUtilizationPerApplication.Size(m)
}
func (m *CpuMemoryUtilizationPerApplication) XXX_DiscardUnknown() {
	xxx_messageInfo_CpuMemoryUtilizationPerApplication.DiscardUnknown(m)
}

var xxx_messageInfo_CpuMemoryUtilizationPerApplication proto.InternalMessageInfo

func (m *CpuMemoryUtilizationPerApplication) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *CpuMemoryUtilizationPerApplication) GetBytesAllocated() uint64 {
	if m != nil && m.BytesAllocated != nil {
		return *m.BytesAllocated
	}
	return 0
}

func (m *CpuMemoryUtilizationPerApplication) GetAllocations() uint64 {
	if m != nil && m.Allocations != nil {
		return *m.Allocations
	}
	return 0
}

func (m *CpuMemoryUtilizationPerApplication) GetFrees() uint64 {
	if m != nil && m.Frees != nil {
		return *m.Frees
	}
	return 0
}

func (m *CpuMemoryUtilizationPerApplication) GetAllocationsFailed() uint64 {
	if m != nil && m.AllocationsFailed != nil {
		return *m.AllocationsFailed
	}
	return 0
}

var E_CpuMemoryUtilExt = &proto.ExtensionDesc{
	ExtendedType:  (*telemetry_top.JuniperNetworksSensors)(nil),
	ExtensionType: (*CpuMemoryUtilization)(nil),
	Field:         1,
	Name:          "cpu_memory_util_ext",
	Tag:           "bytes,1,opt,name=cpu_memory_util_ext,json=cpuMemoryUtilExt",
	Filename:      "plugins/parsers/jtinative/cpu_memory_utilization/cpu_memory_utilization.proto",
}

func init() {
	proto.RegisterType((*CpuMemoryUtilization)(nil), "CpuMemoryUtilization")
	proto.RegisterType((*CpuMemoryUtilizationSummary)(nil), "CpuMemoryUtilizationSummary")
	proto.RegisterType((*CpuMemoryUtilizationPerApplication)(nil), "CpuMemoryUtilizationPerApplication")
	proto.RegisterExtension(E_CpuMemoryUtilExt)
}

func init() {
	proto.RegisterFile("plugins/parsers/jtinative/cpu_memory_utilization/cpu_memory_utilization.proto", fileDescriptor_cpu_memory_utilization_cdb8d37bc3bcd61e)
}

var fileDescriptor_cpu_memory_utilization_cdb8d37bc3bcd61e = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xc1, 0x4a, 0x3b, 0x31,
	0x10, 0xc6, 0xd9, 0x76, 0xf7, 0xf0, 0x4f, 0xe1, 0xaf, 0xa6, 0xd5, 0x2e, 0xea, 0x61, 0x59, 0x0f,
	0xee, 0xc5, 0x16, 0x7a, 0xf4, 0x20, 0x14, 0xd1, 0x83, 0x50, 0x91, 0x2d, 0x7a, 0x52, 0x96, 0xb8,
	0x9d, 0x4a, 0x34, 0x9b, 0x84, 0x24, 0xab, 0xdd, 0x3e, 0xa1, 0x4f, 0xe3, 0x33, 0x88, 0x69, 0xa5,
	0xa9, 0x6c, 0xf5, 0x96, 0xf9, 0xe6, 0xfb, 0x86, 0xf9, 0x0d, 0x41, 0x23, 0xc9, 0xca, 0x27, 0xca,
	0x75, 0x5f, 0x12, 0xa5, 0x41, 0xe9, 0xfe, 0xb3, 0xa1, 0x9c, 0x18, 0xfa, 0x0a, 0xfd, 0x5c, 0x96,
	0x59, 0x01, 0x85, 0x50, 0x55, 0x56, 0x1a, 0xca, 0xe8, 0x9c, 0x18, 0x2a, 0xf8, 0x06, 0xb9, 0x27,
	0x95, 0x30, 0x62, 0xbf, 0x6d, 0x80, 0x41, 0x01, 0x46, 0x55, 0x99, 0x11, 0x72, 0x21, 0xc6, 0x77,
	0xa8, 0x73, 0x2e, 0xcb, 0x91, 0xcd, 0xdc, 0xae, 0x22, 0xf8, 0x0c, 0xb5, 0x9c, 0x09, 0xa1, 0x17,
	0x35, 0x93, 0xd6, 0xe0, 0xb0, 0x57, 0xe7, 0x1d, 0x97, 0x45, 0x41, 0x54, 0x95, 0xba, 0x81, 0xf8,
	0xc3, 0x43, 0x07, 0xbf, 0x98, 0x31, 0x46, 0x3e, 0x27, 0x05, 0x84, 0x5e, 0xe4, 0x25, 0xff, 0x52,
	0xfb, 0xfe, 0xd2, 0x34, 0x9d, 0x43, 0xd8, 0x88, 0xbc, 0xc4, 0x4f, 0xed, 0x1b, 0x1f, 0xa3, 0xad,
	0xc7, 0xca, 0x80, 0xce, 0x08, 0x63, 0x22, 0x27, 0x06, 0x26, 0x61, 0xd3, 0xb6, 0xff, 0x5b, 0x79,
	0xf8, 0xad, 0xe2, 0x68, 0x7d, 0x61, 0x3f, 0xf2, 0x92, 0x60, 0x6d, 0x25, 0x7c, 0x8f, 0xba, 0x44,
	0x4a, 0x46, 0x73, 0x5b, 0xba, 0x07, 0x0a, 0x03, 0x8b, 0x77, 0x54, 0x8b, 0x77, 0x03, 0x6a, 0xb8,
	0x8a, 0xa5, 0x7b, 0xce, 0x0c, 0xc7, 0x15, 0xbf, 0x7b, 0x28, 0xfe, 0x3b, 0x5e, 0xcb, 0x5d, 0xc3,
	0xd8, 0xd8, 0xc4, 0xb8, 0xb4, 0x50, 0xc1, 0xf5, 0xf2, 0x10, 0xae, 0x84, 0x3b, 0x28, 0x98, 0x2a,
	0x00, 0x6d, 0xf9, 0xfd, 0x74, 0x51, 0xe0, 0x13, 0x84, 0x1d, 0x53, 0x36, 0x25, 0x94, 0xc1, 0x24,
	0x0c, 0xac, 0x65, 0xc7, 0xe9, 0x5c, 0xda, 0xc6, 0xe9, 0x03, 0x6a, 0xff, 0xf8, 0x48, 0x19, 0xcc,
	0x0c, 0xee, 0xf6, 0xae, 0x4a, 0x4e, 0x25, 0xa8, 0x6b, 0x30, 0x6f, 0x42, 0xbd, 0xe8, 0x31, 0x70,
	0x2d, 0x94, 0xb6, 0x2c, 0xad, 0xc1, 0x6e, 0xed, 0xf5, 0xd2, 0xed, 0xdc, 0x55, 0x2f, 0x66, 0xe6,
	0x33, 0x00, 0x00, 0xff, 0xff, 0xa6, 0x1e, 0x51, 0x04, 0xd7, 0x02, 0x00, 0x00,
}