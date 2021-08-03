// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fbinternal/cloud/go/services/testcontroller/storage/storage.proto

package storage

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// TestCase is an end-to-end test case.
type TestCase struct {
	// pk uniquely identifies a test case
	Pk int64 `protobuf:"varint,1,opt,name=pk,proto3" json:"pk,omitempty"`
	// type of the test case which identifies its state machine and how to
	// deserialize its configuration
	TestCaseType string `protobuf:"bytes,2,opt,name=testCaseType,proto3" json:"testCaseType,omitempty"`
	// serialized configuration for the test case
	TestConfig []byte `protobuf:"bytes,10,opt,name=testConfig,proto3" json:"testConfig,omitempty"`
	// flag indicating if the test case is currently being run by a worker
	IsCurrentlyExecuting bool `protobuf:"varint,20,opt,name=isCurrentlyExecuting,proto3" json:"isCurrentlyExecuting,omitempty"`
	// timestamp of the last time the test case was claimed for execution
	LastExecutionTime *timestamp.Timestamp `protobuf:"bytes,21,opt,name=lastExecutionTime,proto3" json:"lastExecutionTime,omitempty"`
	// current state machine state of the test case
	State string `protobuf:"bytes,30,opt,name=state,proto3" json:"state,omitempty"`
	// error message, if any, for the test case
	Error string `protobuf:"bytes,31,opt,name=error,proto3" json:"error,omitempty"`
	// next scheduled runtime for the test case if it is currently idle
	NextScheduledTime    *timestamp.Timestamp `protobuf:"bytes,32,opt,name=nextScheduledTime,proto3" json:"nextScheduledTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TestCase) Reset()         { *m = TestCase{} }
func (m *TestCase) String() string { return proto.CompactTextString(m) }
func (*TestCase) ProtoMessage()    {}
func (*TestCase) Descriptor() ([]byte, []int) {
	return fileDescriptor_b34a1d824da7f137, []int{0}
}

func (m *TestCase) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestCase.Unmarshal(m, b)
}
func (m *TestCase) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestCase.Marshal(b, m, deterministic)
}
func (m *TestCase) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestCase.Merge(m, src)
}
func (m *TestCase) XXX_Size() int {
	return xxx_messageInfo_TestCase.Size(m)
}
func (m *TestCase) XXX_DiscardUnknown() {
	xxx_messageInfo_TestCase.DiscardUnknown(m)
}

var xxx_messageInfo_TestCase proto.InternalMessageInfo

func (m *TestCase) GetPk() int64 {
	if m != nil {
		return m.Pk
	}
	return 0
}

func (m *TestCase) GetTestCaseType() string {
	if m != nil {
		return m.TestCaseType
	}
	return ""
}

func (m *TestCase) GetTestConfig() []byte {
	if m != nil {
		return m.TestConfig
	}
	return nil
}

func (m *TestCase) GetIsCurrentlyExecuting() bool {
	if m != nil {
		return m.IsCurrentlyExecuting
	}
	return false
}

func (m *TestCase) GetLastExecutionTime() *timestamp.Timestamp {
	if m != nil {
		return m.LastExecutionTime
	}
	return nil
}

func (m *TestCase) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *TestCase) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *TestCase) GetNextScheduledTime() *timestamp.Timestamp {
	if m != nil {
		return m.NextScheduledTime
	}
	return nil
}

// MutableTestCase encapsulates the set of fields available to clients for
// modification. See TestCase for documentation on fields.
type MutableTestCase struct {
	Pk                   int64    `protobuf:"varint,1,opt,name=pk,proto3" json:"pk,omitempty"`
	TestCaseType         string   `protobuf:"bytes,2,opt,name=testCaseType,proto3" json:"testCaseType,omitempty"`
	TestConfig           []byte   `protobuf:"bytes,10,opt,name=testConfig,proto3" json:"testConfig,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutableTestCase) Reset()         { *m = MutableTestCase{} }
func (m *MutableTestCase) String() string { return proto.CompactTextString(m) }
func (*MutableTestCase) ProtoMessage()    {}
func (*MutableTestCase) Descriptor() ([]byte, []int) {
	return fileDescriptor_b34a1d824da7f137, []int{1}
}

func (m *MutableTestCase) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutableTestCase.Unmarshal(m, b)
}
func (m *MutableTestCase) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutableTestCase.Marshal(b, m, deterministic)
}
func (m *MutableTestCase) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutableTestCase.Merge(m, src)
}
func (m *MutableTestCase) XXX_Size() int {
	return xxx_messageInfo_MutableTestCase.Size(m)
}
func (m *MutableTestCase) XXX_DiscardUnknown() {
	xxx_messageInfo_MutableTestCase.DiscardUnknown(m)
}

var xxx_messageInfo_MutableTestCase proto.InternalMessageInfo

func (m *MutableTestCase) GetPk() int64 {
	if m != nil {
		return m.Pk
	}
	return 0
}

func (m *MutableTestCase) GetTestCaseType() string {
	if m != nil {
		return m.TestCaseType
	}
	return ""
}

func (m *MutableTestCase) GetTestConfig() []byte {
	if m != nil {
		return m.TestConfig
	}
	return nil
}

// CINode is a baremetal CI workload executor
type CINode struct {
	// unique ID for the node (e.g. VPN client ID)
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// IP address for the node on the VPN
	VpnIp string `protobuf:"bytes,2,opt,name=vpnIp,proto3" json:"vpnIp,omitempty"`
	// Tag for the node
	Tag string `protobuf:"bytes,3,opt,name=tag,proto3" json:"tag,omitempty"`
	// is the node available or not
	Available bool `protobuf:"varint,10,opt,name=available,proto3" json:"available,omitempty"`
	// the last time this node was leased out
	LastLeaseTime        *timestamp.Timestamp `protobuf:"bytes,11,opt,name=lastLeaseTime,proto3" json:"lastLeaseTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CINode) Reset()         { *m = CINode{} }
func (m *CINode) String() string { return proto.CompactTextString(m) }
func (*CINode) ProtoMessage()    {}
func (*CINode) Descriptor() ([]byte, []int) {
	return fileDescriptor_b34a1d824da7f137, []int{2}
}

func (m *CINode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CINode.Unmarshal(m, b)
}
func (m *CINode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CINode.Marshal(b, m, deterministic)
}
func (m *CINode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CINode.Merge(m, src)
}
func (m *CINode) XXX_Size() int {
	return xxx_messageInfo_CINode.Size(m)
}
func (m *CINode) XXX_DiscardUnknown() {
	xxx_messageInfo_CINode.DiscardUnknown(m)
}

var xxx_messageInfo_CINode proto.InternalMessageInfo

func (m *CINode) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CINode) GetVpnIp() string {
	if m != nil {
		return m.VpnIp
	}
	return ""
}

func (m *CINode) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *CINode) GetAvailable() bool {
	if m != nil {
		return m.Available
	}
	return false
}

func (m *CINode) GetLastLeaseTime() *timestamp.Timestamp {
	if m != nil {
		return m.LastLeaseTime
	}
	return nil
}

// MutableCINode encapsulates the set iof fields available to clients for
// modification. See CINode for documentation on fields.
type MutableCINode struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	VpnIP                string   `protobuf:"bytes,2,opt,name=vpnIP,proto3" json:"vpnIP,omitempty"`
	Tag                  string   `protobuf:"bytes,3,opt,name=tag,proto3" json:"tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutableCINode) Reset()         { *m = MutableCINode{} }
func (m *MutableCINode) String() string { return proto.CompactTextString(m) }
func (*MutableCINode) ProtoMessage()    {}
func (*MutableCINode) Descriptor() ([]byte, []int) {
	return fileDescriptor_b34a1d824da7f137, []int{3}
}

func (m *MutableCINode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutableCINode.Unmarshal(m, b)
}
func (m *MutableCINode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutableCINode.Marshal(b, m, deterministic)
}
func (m *MutableCINode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutableCINode.Merge(m, src)
}
func (m *MutableCINode) XXX_Size() int {
	return xxx_messageInfo_MutableCINode.Size(m)
}
func (m *MutableCINode) XXX_DiscardUnknown() {
	xxx_messageInfo_MutableCINode.DiscardUnknown(m)
}

var xxx_messageInfo_MutableCINode proto.InternalMessageInfo

func (m *MutableCINode) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MutableCINode) GetVpnIP() string {
	if m != nil {
		return m.VpnIP
	}
	return ""
}

func (m *MutableCINode) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

// NodeLease encapsulates a successful node lease. To release the lease on the
// node, the same leaseID must be provided.
type NodeLease struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	LeaseID              string   `protobuf:"bytes,2,opt,name=leaseID,proto3" json:"leaseID,omitempty"`
	VpnIP                string   `protobuf:"bytes,10,opt,name=vpnIP,proto3" json:"vpnIP,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeLease) Reset()         { *m = NodeLease{} }
func (m *NodeLease) String() string { return proto.CompactTextString(m) }
func (*NodeLease) ProtoMessage()    {}
func (*NodeLease) Descriptor() ([]byte, []int) {
	return fileDescriptor_b34a1d824da7f137, []int{4}
}

func (m *NodeLease) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeLease.Unmarshal(m, b)
}
func (m *NodeLease) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeLease.Marshal(b, m, deterministic)
}
func (m *NodeLease) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeLease.Merge(m, src)
}
func (m *NodeLease) XXX_Size() int {
	return xxx_messageInfo_NodeLease.Size(m)
}
func (m *NodeLease) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeLease.DiscardUnknown(m)
}

var xxx_messageInfo_NodeLease proto.InternalMessageInfo

func (m *NodeLease) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *NodeLease) GetLeaseID() string {
	if m != nil {
		return m.LeaseID
	}
	return ""
}

func (m *NodeLease) GetVpnIP() string {
	if m != nil {
		return m.VpnIP
	}
	return ""
}

func init() {
	proto.RegisterType((*TestCase)(nil), "magma.fbinternal.testcontroller.storage.TestCase")
	proto.RegisterType((*MutableTestCase)(nil), "magma.fbinternal.testcontroller.storage.MutableTestCase")
	proto.RegisterType((*CINode)(nil), "magma.fbinternal.testcontroller.storage.CINode")
	proto.RegisterType((*MutableCINode)(nil), "magma.fbinternal.testcontroller.storage.MutableCINode")
	proto.RegisterType((*NodeLease)(nil), "magma.fbinternal.testcontroller.storage.NodeLease")
}

func init() {
	proto.RegisterFile("fbinternal/cloud/go/services/testcontroller/storage/storage.proto", fileDescriptor_b34a1d824da7f137)
}

var fileDescriptor_b34a1d824da7f137 = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0xcf, 0x6a, 0xdb, 0x40,
	0x10, 0xc6, 0x91, 0x4c, 0x53, 0x7b, 0x92, 0xf4, 0xcf, 0xe2, 0x82, 0x08, 0x25, 0x11, 0xba, 0x54,
	0x27, 0x09, 0xd2, 0x53, 0xe9, 0xa5, 0xad, 0x5b, 0x5a, 0xd3, 0x3f, 0x04, 0xd5, 0xa7, 0xde, 0xd6,
	0xd2, 0x58, 0x5d, 0xb2, 0xda, 0x15, 0xbb, 0x23, 0x93, 0xbc, 0x4d, 0x9f, 0xa9, 0x4f, 0x54, 0x76,
	0x25, 0x63, 0x1b, 0x97, 0x06, 0x0a, 0x39, 0xd9, 0xf3, 0x8d, 0xbe, 0xdf, 0x7c, 0x33, 0xb0, 0xf0,
	0x76, 0xb5, 0x14, 0x8a, 0xd0, 0x28, 0x2e, 0xf3, 0x52, 0xea, 0xae, 0xca, 0x6b, 0x9d, 0x5b, 0x34,
	0x6b, 0x51, 0xa2, 0xcd, 0x09, 0x2d, 0x95, 0x5a, 0x91, 0xd1, 0x52, 0xa2, 0xc9, 0x2d, 0x69, 0xc3,
	0x6b, 0xdc, 0xfc, 0x66, 0xad, 0xd1, 0xa4, 0xd9, 0x8b, 0x86, 0xd7, 0x0d, 0xcf, 0xb6, 0xa0, 0x6c,
	0xdf, 0x96, 0x0d, 0x9f, 0x9f, 0x5d, 0xd4, 0x5a, 0xd7, 0x12, 0x73, 0x6f, 0x5b, 0x76, 0xab, 0x9c,
	0x44, 0x83, 0x96, 0x78, 0xd3, 0xf6, 0xa4, 0xe4, 0x77, 0x08, 0xe3, 0x05, 0x5a, 0x9a, 0x71, 0x8b,
	0xec, 0x11, 0x84, 0xed, 0x75, 0x14, 0xc4, 0x41, 0x3a, 0x2a, 0xc2, 0xf6, 0x9a, 0x25, 0x70, 0x42,
	0x43, 0x6f, 0x71, 0xdb, 0x62, 0x14, 0xc6, 0x41, 0x3a, 0x29, 0xf6, 0x34, 0x76, 0x0e, 0xe0, 0x6b,
	0xad, 0x56, 0xa2, 0x8e, 0x20, 0x0e, 0xd2, 0x93, 0x62, 0x47, 0x61, 0x97, 0x30, 0x15, 0x76, 0xd6,
	0x19, 0x83, 0x8a, 0xe4, 0xed, 0x87, 0x1b, 0x2c, 0x3b, 0x12, 0xaa, 0x8e, 0xa6, 0x71, 0x90, 0x8e,
	0x8b, 0xbf, 0xf6, 0xd8, 0x27, 0x78, 0x2a, 0xb9, 0xa5, 0x41, 0xd0, 0x6a, 0x21, 0x1a, 0x8c, 0x9e,
	0xc5, 0x41, 0x7a, 0x7c, 0x79, 0x96, 0xf5, 0x1b, 0x65, 0x9b, 0x8d, 0xb2, 0xc5, 0x66, 0xa3, 0xe2,
	0xd0, 0xc4, 0xa6, 0xf0, 0xc0, 0x12, 0x27, 0x8c, 0xce, 0x7d, 0xf4, 0xbe, 0x70, 0x2a, 0x1a, 0xa3,
	0x4d, 0x74, 0xd1, 0xab, 0xbe, 0x70, 0x53, 0x15, 0xde, 0xd0, 0xf7, 0xf2, 0x27, 0x56, 0x9d, 0xc4,
	0xca, 0x4f, 0x8d, 0xef, 0x9e, 0x7a, 0x60, 0x4a, 0x10, 0x1e, 0x7f, 0xed, 0x88, 0x2f, 0x25, 0xde,
	0xe7, 0x69, 0x93, 0x5f, 0x01, 0x1c, 0xcd, 0xe6, 0xdf, 0x74, 0xe5, 0xf1, 0xa2, 0xf2, 0xf8, 0x49,
	0x11, 0x8a, 0xca, 0x6d, 0xb8, 0x6e, 0xd5, 0xbc, 0x1d, 0xb8, 0x7d, 0xc1, 0x9e, 0xc0, 0x88, 0x78,
	0x1d, 0x8d, 0xbc, 0xe6, 0xfe, 0xb2, 0xe7, 0x30, 0xe1, 0x6b, 0x2e, 0xa4, 0xcb, 0xea, 0x27, 0x8c,
	0x8b, 0xad, 0xc0, 0xde, 0xc0, 0xa9, 0x3b, 0xe9, 0x17, 0x74, 0x89, 0xdc, 0x35, 0x8e, 0xef, 0xbc,
	0xc6, 0xbe, 0x21, 0xf9, 0x08, 0xa7, 0xc3, 0x25, 0xfe, 0x1d, 0xf4, 0x6a, 0x37, 0xe8, 0xd5, 0x61,
	0xd0, 0xe4, 0x33, 0x4c, 0x9c, 0xdf, 0x93, 0x0f, 0x20, 0x11, 0x3c, 0x94, 0xae, 0x31, 0x7f, 0x3f,
	0x60, 0x36, 0xe5, 0x16, 0x0f, 0x3b, 0xf8, 0x77, 0xaf, 0x7f, 0xbc, 0xf2, 0x0f, 0x28, 0xff, 0x8f,
	0x97, 0xb8, 0x3c, 0xf2, 0x5b, 0xbf, 0xfc, 0x13, 0x00, 0x00, 0xff, 0xff, 0x02, 0x66, 0xe3, 0xc1,
	0xc7, 0x03, 0x00, 0x00,
}
