// Code generated by protoc-gen-go. DO NOT EDIT.
// source: plugins/parsers/jtinative/agentd/agentd.proto

package agentd

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

type TelemetrySystem struct {
	Subscriptions        *TelemetrySystemSubscriptionsType `protobuf:"bytes,151,opt,name=subscriptions" json:"subscriptions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *TelemetrySystem) Reset()         { *m = TelemetrySystem{} }
func (m *TelemetrySystem) String() string { return proto.CompactTextString(m) }
func (*TelemetrySystem) ProtoMessage()    {}
func (*TelemetrySystem) Descriptor() ([]byte, []int) {
	return fileDescriptor_agentd_df0ea598a328e1aa, []int{0}
}
func (m *TelemetrySystem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TelemetrySystem.Unmarshal(m, b)
}
func (m *TelemetrySystem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TelemetrySystem.Marshal(b, m, deterministic)
}
func (dst *TelemetrySystem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TelemetrySystem.Merge(dst, src)
}
func (m *TelemetrySystem) XXX_Size() int {
	return xxx_messageInfo_TelemetrySystem.Size(m)
}
func (m *TelemetrySystem) XXX_DiscardUnknown() {
	xxx_messageInfo_TelemetrySystem.DiscardUnknown(m)
}

var xxx_messageInfo_TelemetrySystem proto.InternalMessageInfo

func (m *TelemetrySystem) GetSubscriptions() *TelemetrySystemSubscriptionsType {
	if m != nil {
		return m.Subscriptions
	}
	return nil
}

type TelemetrySystemSubscriptionsType struct {
	Dynamic              *TelemetrySystemSubscriptionsTypeDynamicType `protobuf:"bytes,151,opt,name=dynamic" json:"dynamic,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                     `json:"-"`
	XXX_unrecognized     []byte                                       `json:"-"`
	XXX_sizecache        int32                                        `json:"-"`
}

func (m *TelemetrySystemSubscriptionsType) Reset()         { *m = TelemetrySystemSubscriptionsType{} }
func (m *TelemetrySystemSubscriptionsType) String() string { return proto.CompactTextString(m) }
func (*TelemetrySystemSubscriptionsType) ProtoMessage()    {}
func (*TelemetrySystemSubscriptionsType) Descriptor() ([]byte, []int) {
	return fileDescriptor_agentd_df0ea598a328e1aa, []int{0, 0}
}
func (m *TelemetrySystemSubscriptionsType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TelemetrySystemSubscriptionsType.Unmarshal(m, b)
}
func (m *TelemetrySystemSubscriptionsType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TelemetrySystemSubscriptionsType.Marshal(b, m, deterministic)
}
func (dst *TelemetrySystemSubscriptionsType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TelemetrySystemSubscriptionsType.Merge(dst, src)
}
func (m *TelemetrySystemSubscriptionsType) XXX_Size() int {
	return xxx_messageInfo_TelemetrySystemSubscriptionsType.Size(m)
}
func (m *TelemetrySystemSubscriptionsType) XXX_DiscardUnknown() {
	xxx_messageInfo_TelemetrySystemSubscriptionsType.DiscardUnknown(m)
}

var xxx_messageInfo_TelemetrySystemSubscriptionsType proto.InternalMessageInfo

func (m *TelemetrySystemSubscriptionsType) GetDynamic() *TelemetrySystemSubscriptionsTypeDynamicType {
	if m != nil {
		return m.Dynamic
	}
	return nil
}

type TelemetrySystemSubscriptionsTypeDynamicType struct {
	Subscription         []*TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionList `protobuf:"bytes,151,rep,name=subscription" json:"subscription,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                       `json:"-"`
	XXX_unrecognized     []byte                                                         `json:"-"`
	XXX_sizecache        int32                                                          `json:"-"`
}

func (m *TelemetrySystemSubscriptionsTypeDynamicType) Reset() {
	*m = TelemetrySystemSubscriptionsTypeDynamicType{}
}
func (m *TelemetrySystemSubscriptionsTypeDynamicType) String() string {
	return proto.CompactTextString(m)
}
func (*TelemetrySystemSubscriptionsTypeDynamicType) ProtoMessage() {}
func (*TelemetrySystemSubscriptionsTypeDynamicType) Descriptor() ([]byte, []int) {
	return fileDescriptor_agentd_df0ea598a328e1aa, []int{0, 0, 0}
}
func (m *TelemetrySystemSubscriptionsTypeDynamicType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicType.Unmarshal(m, b)
}
func (m *TelemetrySystemSubscriptionsTypeDynamicType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicType.Marshal(b, m, deterministic)
}
func (dst *TelemetrySystemSubscriptionsTypeDynamicType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicType.Merge(dst, src)
}
func (m *TelemetrySystemSubscriptionsTypeDynamicType) XXX_Size() int {
	return xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicType.Size(m)
}
func (m *TelemetrySystemSubscriptionsTypeDynamicType) XXX_DiscardUnknown() {
	xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicType.DiscardUnknown(m)
}

var xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicType proto.InternalMessageInfo

func (m *TelemetrySystemSubscriptionsTypeDynamicType) GetSubscription() []*TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionList {
	if m != nil {
		return m.Subscription
	}
	return nil
}

type TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionList struct {
	SubscriptionId       *uint64                                                                     `protobuf:"varint,51,opt,name=subscription_id,json=subscriptionId" json:"subscription_id,omitempty"`
	State                *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListStateType       `protobuf:"bytes,151,opt,name=state" json:"state,omitempty"`
	SensorPaths          *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsType `protobuf:"bytes,152,opt,name=sensor_paths,json=sensorPaths" json:"sensor_paths,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                                    `json:"-"`
	XXX_unrecognized     []byte                                                                      `json:"-"`
	XXX_sizecache        int32                                                                       `json:"-"`
}

func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionList) Reset() {
	*m = TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionList{}
}
func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionList) String() string {
	return proto.CompactTextString(m)
}
func (*TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionList) ProtoMessage() {}
func (*TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionList) Descriptor() ([]byte, []int) {
	return fileDescriptor_agentd_df0ea598a328e1aa, []int{0, 0, 0, 0}
}
func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionList.Unmarshal(m, b)
}
func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionList.Marshal(b, m, deterministic)
}
func (dst *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionList.Merge(dst, src)
}
func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionList) XXX_Size() int {
	return xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionList.Size(m)
}
func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionList) XXX_DiscardUnknown() {
	xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionList.DiscardUnknown(m)
}

var xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionList proto.InternalMessageInfo

func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionList) GetSubscriptionId() uint64 {
	if m != nil && m.SubscriptionId != nil {
		return *m.SubscriptionId
	}
	return 0
}

func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionList) GetState() *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListStateType {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionList) GetSensorPaths() *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsType {
	if m != nil {
		return m.SensorPaths
	}
	return nil
}

type TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListStateType struct {
	SubscriptionId       *uint64  `protobuf:"varint,51,opt,name=subscription_id,json=subscriptionId" json:"subscription_id,omitempty"`
	SampleInterval       *uint64  `protobuf:"varint,52,opt,name=sample_interval,json=sampleInterval" json:"sample_interval,omitempty"`
	HeartbeatInterval    *uint64  `protobuf:"varint,53,opt,name=heartbeat_interval,json=heartbeatInterval" json:"heartbeat_interval,omitempty"`
	SuppressRedundant    *bool    `protobuf:"varint,54,opt,name=suppress_redundant,json=suppressRedundant" json:"suppress_redundant,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListStateType) Reset() {
	*m = TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListStateType{}
}
func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListStateType) String() string {
	return proto.CompactTextString(m)
}
func (*TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListStateType) ProtoMessage() {}
func (*TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListStateType) Descriptor() ([]byte, []int) {
	return fileDescriptor_agentd_df0ea598a328e1aa, []int{0, 0, 0, 0, 0}
}
func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListStateType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListStateType.Unmarshal(m, b)
}
func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListStateType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListStateType.Marshal(b, m, deterministic)
}
func (dst *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListStateType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListStateType.Merge(dst, src)
}
func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListStateType) XXX_Size() int {
	return xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListStateType.Size(m)
}
func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListStateType) XXX_DiscardUnknown() {
	xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListStateType.DiscardUnknown(m)
}

var xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListStateType proto.InternalMessageInfo

func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListStateType) GetSubscriptionId() uint64 {
	if m != nil && m.SubscriptionId != nil {
		return *m.SubscriptionId
	}
	return 0
}

func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListStateType) GetSampleInterval() uint64 {
	if m != nil && m.SampleInterval != nil {
		return *m.SampleInterval
	}
	return 0
}

func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListStateType) GetHeartbeatInterval() uint64 {
	if m != nil && m.HeartbeatInterval != nil {
		return *m.HeartbeatInterval
	}
	return 0
}

func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListStateType) GetSuppressRedundant() bool {
	if m != nil && m.SuppressRedundant != nil {
		return *m.SuppressRedundant
	}
	return false
}

type TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsType struct {
	SensorPath           []*TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathList `protobuf:"bytes,151,rep,name=sensor_path,json=sensorPath" json:"sensor_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                                                    `json:"-"`
	XXX_unrecognized     []byte                                                                                      `json:"-"`
	XXX_sizecache        int32                                                                                       `json:"-"`
}

func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsType) Reset() {
	*m = TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsType{}
}
func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsType) String() string {
	return proto.CompactTextString(m)
}
func (*TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsType) ProtoMessage() {}
func (*TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsType) Descriptor() ([]byte, []int) {
	return fileDescriptor_agentd_df0ea598a328e1aa, []int{0, 0, 0, 0, 1}
}
func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsType.Unmarshal(m, b)
}
func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsType.Marshal(b, m, deterministic)
}
func (dst *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsType.Merge(dst, src)
}
func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsType) XXX_Size() int {
	return xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsType.Size(m)
}
func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsType) XXX_DiscardUnknown() {
	xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsType.DiscardUnknown(m)
}

var xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsType proto.InternalMessageInfo

func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsType) GetSensorPath() []*TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathList {
	if m != nil {
		return m.SensorPath
	}
	return nil
}

type TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathList struct {
	Path                 *string                                                                                            `protobuf:"bytes,51,opt,name=path" json:"path,omitempty"`
	State                *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathListStateType `protobuf:"bytes,151,opt,name=state" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                                                           `json:"-"`
	XXX_unrecognized     []byte                                                                                             `json:"-"`
	XXX_sizecache        int32                                                                                              `json:"-"`
}

func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathList) Reset() {
	*m = TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathList{}
}
func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathList) String() string {
	return proto.CompactTextString(m)
}
func (*TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathList) ProtoMessage() {
}
func (*TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathList) Descriptor() ([]byte, []int) {
	return fileDescriptor_agentd_df0ea598a328e1aa, []int{0, 0, 0, 0, 1, 0}
}
func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathList.Unmarshal(m, b)
}
func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathList.Marshal(b, m, deterministic)
}
func (dst *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathList.Merge(dst, src)
}
func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathList) XXX_Size() int {
	return xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathList.Size(m)
}
func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathList) XXX_DiscardUnknown() {
	xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathList.DiscardUnknown(m)
}

var xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathList proto.InternalMessageInfo

func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathList) GetPath() string {
	if m != nil && m.Path != nil {
		return *m.Path
	}
	return ""
}

func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathList) GetState() *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathListStateType {
	if m != nil {
		return m.State
	}
	return nil
}

type TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathListStateType struct {
	Path                 *string  `protobuf:"bytes,51,opt,name=path" json:"path,omitempty"`
	ExcludeFilter        *string  `protobuf:"bytes,52,opt,name=exclude_filter,json=excludeFilter" json:"exclude_filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathListStateType) Reset() {
	*m = TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathListStateType{}
}
func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathListStateType) String() string {
	return proto.CompactTextString(m)
}
func (*TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathListStateType) ProtoMessage() {
}
func (*TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathListStateType) Descriptor() ([]byte, []int) {
	return fileDescriptor_agentd_df0ea598a328e1aa, []int{0, 0, 0, 0, 1, 0, 0}
}
func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathListStateType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathListStateType.Unmarshal(m, b)
}
func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathListStateType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathListStateType.Marshal(b, m, deterministic)
}
func (dst *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathListStateType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathListStateType.Merge(dst, src)
}
func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathListStateType) XXX_Size() int {
	return xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathListStateType.Size(m)
}
func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathListStateType) XXX_DiscardUnknown() {
	xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathListStateType.DiscardUnknown(m)
}

var xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathListStateType proto.InternalMessageInfo

func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathListStateType) GetPath() string {
	if m != nil && m.Path != nil {
		return *m.Path
	}
	return ""
}

func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathListStateType) GetExcludeFilter() string {
	if m != nil && m.ExcludeFilter != nil {
		return *m.ExcludeFilter
	}
	return ""
}

var E_JnprTelemetrySystemExt = &proto.ExtensionDesc{
	ExtendedType:  (*telemetry_top.JuniperNetworksSensors)(nil),
	ExtensionType: (*TelemetrySystem)(nil),
	Field:         31,
	Name:          "jnpr_telemetry_system_ext",
	Tag:           "bytes,31,opt,name=jnpr_telemetry_system_ext,json=jnprTelemetrySystemExt",
	Filename:      "plugins/parsers/jtinative/agentd/agentd.proto",
}

func init() {
	proto.RegisterType((*TelemetrySystem)(nil), "telemetry_system")
	proto.RegisterType((*TelemetrySystemSubscriptionsType)(nil), "telemetry_system.subscriptions_type")
	proto.RegisterType((*TelemetrySystemSubscriptionsTypeDynamicType)(nil), "telemetry_system.subscriptions_type.dynamic_type")
	proto.RegisterType((*TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionList)(nil), "telemetry_system.subscriptions_type.dynamic_type.subscription_list")
	proto.RegisterType((*TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListStateType)(nil), "telemetry_system.subscriptions_type.dynamic_type.subscription_list.state_type")
	proto.RegisterType((*TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsType)(nil), "telemetry_system.subscriptions_type.dynamic_type.subscription_list.sensor_paths_type")
	proto.RegisterType((*TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathList)(nil), "telemetry_system.subscriptions_type.dynamic_type.subscription_list.sensor_paths_type.sensor_path_list")
	proto.RegisterType((*TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathListStateType)(nil), "telemetry_system.subscriptions_type.dynamic_type.subscription_list.sensor_paths_type.sensor_path_list.state_type")
	proto.RegisterExtension(E_JnprTelemetrySystemExt)
}

func init() {
	proto.RegisterFile("plugins/parsers/jtinative/agentd/agentd.proto", fileDescriptor_agentd_df0ea598a328e1aa)
}

var fileDescriptor_agentd_df0ea598a328e1aa = []byte{
	// 508 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xe5, 0xfe, 0x51, 0x93, 0x49, 0x5a, 0x9a, 0x45, 0x02, 0xd7, 0x17, 0x22, 0x04, 0x22,
	0x97, 0x3a, 0xa2, 0xfc, 0x39, 0x70, 0x42, 0x20, 0x40, 0x05, 0x54, 0x21, 0x17, 0xce, 0xd6, 0x36,
	0x9e, 0x36, 0x5b, 0x9c, 0xf5, 0xb2, 0x3b, 0x2e, 0xf1, 0x11, 0xc4, 0x01, 0x9e, 0x00, 0xce, 0x3c,
	0x07, 0x6f, 0xc0, 0x93, 0xf0, 0x16, 0xc8, 0x6b, 0x9b, 0xd8, 0xb1, 0x90, 0x40, 0x2a, 0x9c, 0x56,
	0xfa, 0xcd, 0xb7, 0xdf, 0x7c, 0x33, 0x1a, 0xd8, 0x55, 0x71, 0x7a, 0x22, 0xa4, 0x19, 0x2b, 0xae,
	0x0d, 0x6a, 0x33, 0x3e, 0x25, 0x21, 0x39, 0x89, 0x33, 0x1c, 0xf3, 0x13, 0x94, 0x14, 0x95, 0x8f,
	0xaf, 0x74, 0x42, 0x89, 0x77, 0x91, 0x30, 0xc6, 0x19, 0x92, 0xce, 0x42, 0x4a, 0x54, 0x01, 0xaf,
	0x7e, 0xef, 0xc2, 0xf6, 0x82, 0x9b, 0xcc, 0x10, 0xce, 0xd8, 0x33, 0xd8, 0x34, 0xe9, 0x91, 0x99,
	0x68, 0xa1, 0x48, 0x24, 0xd2, 0xb8, 0x9f, 0x9d, 0xa1, 0x33, 0xea, 0xed, 0x5d, 0xf3, 0x97, 0xa5,
	0x7e, 0x43, 0x17, 0x52, 0xa6, 0x30, 0x68, 0xfe, 0xf5, 0x7e, 0x74, 0x80, 0xb5, 0x55, 0xec, 0x39,
	0x6c, 0x44, 0x99, 0xe4, 0x33, 0x31, 0xa9, 0xdc, 0x6f, 0xfe, 0x89, 0xbb, 0x5f, 0x7e, 0x2a, 0x5a,
	0x55, 0x16, 0xde, 0x87, 0x0e, 0xf4, 0xeb, 0x15, 0x36, 0x85, 0x7e, 0xfd, 0x73, 0xde, 0x63, 0x75,
	0xd4, 0xdb, 0x7b, 0xf8, 0xd7, 0x3d, 0x1a, 0xf5, 0x30, 0x16, 0x86, 0x82, 0x86, 0xb3, 0xf7, 0x69,
	0x03, 0x06, 0x2d, 0x0d, 0xf3, 0xe1, 0x42, 0x03, 0x8a, 0xc8, 0xbd, 0x35, 0x74, 0x46, 0x6b, 0x0f,
	0xd6, 0xdf, 0xdf, 0x5f, 0xe9, 0x38, 0xc1, 0x56, 0xbd, 0xba, 0x1f, 0x31, 0x84, 0x75, 0x43, 0x9c,
	0xb0, 0x5a, 0xc6, 0xc1, 0x39, 0x04, 0xf5, 0xad, 0x63, 0xb1, 0xa9, 0xc2, 0x9d, 0x65, 0xd0, 0x37,
	0x28, 0x4d, 0xa2, 0x43, 0xc5, 0x69, 0x6a, 0xdc, 0x2f, 0x45, 0xb7, 0x57, 0xe7, 0xd2, 0xad, 0x66,
	0x5c, 0x34, 0xed, 0x15, 0xe8, 0x45, 0x4e, 0xbc, 0x6f, 0x0e, 0xc0, 0x22, 0x10, 0xbb, 0xf1, 0x9b,
	0x05, 0xb5, 0x36, 0x93, 0x0b, 0xf9, 0x4c, 0xc5, 0x18, 0x0a, 0x49, 0xa8, 0xcf, 0x78, 0xec, 0xde,
	0x2e, 0x85, 0x16, 0xef, 0x97, 0x94, 0xed, 0x02, 0x9b, 0x22, 0xd7, 0x74, 0x84, 0x9c, 0x16, 0xda,
	0x3b, 0x56, 0x3b, 0xf8, 0x55, 0xa9, 0xcb, 0x4d, 0xaa, 0x94, 0x46, 0x63, 0x42, 0x8d, 0x51, 0x2a,
	0x23, 0x2e, 0xc9, 0xbd, 0x3b, 0x74, 0x46, 0x9d, 0x60, 0x50, 0x55, 0x82, 0xaa, 0xe0, 0x7d, 0x5d,
	0x85, 0x41, 0x6b, 0x42, 0xf6, 0xd1, 0x81, 0x5e, 0x8d, 0x56, 0x67, 0x76, 0xfc, 0x4f, 0xf6, 0x59,
	0x27, 0xc5, 0x25, 0xc2, 0x62, 0xc1, 0xde, 0xbb, 0x15, 0xd8, 0x5e, 0x16, 0xb0, 0x1d, 0x58, 0xb3,
	0xb9, 0xf2, 0xd5, 0x76, 0xab, 0xdb, 0xb3, 0x28, 0x8f, 0xde, 0x3c, 0xb9, 0x37, 0xff, 0x27, 0x74,
	0xfb, 0x2a, 0xbd, 0x27, 0x8d, 0xcb, 0x60, 0xf5, 0xcc, 0x65, 0xd8, 0xeb, 0xb0, 0x85, 0xf3, 0x49,
	0x9c, 0x46, 0x18, 0x1e, 0x8b, 0x98, 0x50, 0xdb, 0x1b, 0xe8, 0x06, 0x9b, 0x25, 0x7d, 0x6c, 0xe1,
	0x3d, 0x84, 0x9d, 0x53, 0xa9, 0x74, 0xb8, 0x3c, 0x48, 0x88, 0x73, 0x62, 0x97, 0xfd, 0xa7, 0xa9,
	0x14, 0x0a, 0xf5, 0x01, 0xd2, 0xdb, 0x44, 0xbf, 0x36, 0x87, 0x36, 0x99, 0x71, 0xaf, 0xd8, 0xf1,
	0x07, 0xad, 0xf1, 0x83, 0x4b, 0xb9, 0xd9, 0xcb, 0x8a, 0x1e, 0x5a, 0xf8, 0x68, 0x4e, 0x3f, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x7c, 0xc8, 0xb9, 0x4e, 0x79, 0x05, 0x00, 0x00,
}
