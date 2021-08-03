// Code generated by protoc-gen-go. DO NOT EDIT.
// source: feg/protos/envoy_controller.proto

package protos

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protos "magma/lte/cloud/go/protos"
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

type AddUEHeaderEnrichmentResult_Result int32

const (
	AddUEHeaderEnrichmentResult_SUCCESS          AddUEHeaderEnrichmentResult_Result = 0
	AddUEHeaderEnrichmentResult_WEBSITE_CONFLICT AddUEHeaderEnrichmentResult_Result = 1
	AddUEHeaderEnrichmentResult_RULE_ID_CONFLICT AddUEHeaderEnrichmentResult_Result = 2
	AddUEHeaderEnrichmentResult_FAILURE          AddUEHeaderEnrichmentResult_Result = 3
)

var AddUEHeaderEnrichmentResult_Result_name = map[int32]string{
	0: "SUCCESS",
	1: "WEBSITE_CONFLICT",
	2: "RULE_ID_CONFLICT",
	3: "FAILURE",
}

var AddUEHeaderEnrichmentResult_Result_value = map[string]int32{
	"SUCCESS":          0,
	"WEBSITE_CONFLICT": 1,
	"RULE_ID_CONFLICT": 2,
	"FAILURE":          3,
}

func (x AddUEHeaderEnrichmentResult_Result) String() string {
	return proto.EnumName(AddUEHeaderEnrichmentResult_Result_name, int32(x))
}

func (AddUEHeaderEnrichmentResult_Result) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c190610d29559b01, []int{2, 0}
}

type DeactivateUEHeaderEnrichmentResult_Result int32

const (
	DeactivateUEHeaderEnrichmentResult_SUCCESS        DeactivateUEHeaderEnrichmentResult_Result = 0
	DeactivateUEHeaderEnrichmentResult_UE_NOT_FOUND   DeactivateUEHeaderEnrichmentResult_Result = 1
	DeactivateUEHeaderEnrichmentResult_RULE_NOT_FOUND DeactivateUEHeaderEnrichmentResult_Result = 2
	DeactivateUEHeaderEnrichmentResult_FAILURE        DeactivateUEHeaderEnrichmentResult_Result = 3
)

var DeactivateUEHeaderEnrichmentResult_Result_name = map[int32]string{
	0: "SUCCESS",
	1: "UE_NOT_FOUND",
	2: "RULE_NOT_FOUND",
	3: "FAILURE",
}

var DeactivateUEHeaderEnrichmentResult_Result_value = map[string]int32{
	"SUCCESS":        0,
	"UE_NOT_FOUND":   1,
	"RULE_NOT_FOUND": 2,
	"FAILURE":        3,
}

func (x DeactivateUEHeaderEnrichmentResult_Result) String() string {
	return proto.EnumName(DeactivateUEHeaderEnrichmentResult_Result_name, int32(x))
}

func (DeactivateUEHeaderEnrichmentResult_Result) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c190610d29559b01, []int{4, 0}
}

type Header struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}
func (*Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_c190610d29559b01, []int{0}
}

func (m *Header) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Header.Unmarshal(m, b)
}
func (m *Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Header.Marshal(b, m, deterministic)
}
func (m *Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Header.Merge(m, src)
}
func (m *Header) XXX_Size() int {
	return xxx_messageInfo_Header.Size(m)
}
func (m *Header) XXX_DiscardUnknown() {
	xxx_messageInfo_Header.DiscardUnknown(m)
}

var xxx_messageInfo_Header proto.InternalMessageInfo

func (m *Header) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Header) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type AddUEHeaderEnrichmentRequest struct {
	UeIp                 *protos.IPAddress `protobuf:"bytes,1,opt,name=ue_ip,json=ueIp,proto3" json:"ue_ip,omitempty"`
	Websites             []string          `protobuf:"bytes,2,rep,name=websites,proto3" json:"websites,omitempty"`
	Headers              []*Header         `protobuf:"bytes,3,rep,name=headers,proto3" json:"headers,omitempty"`
	RuleId               string            `protobuf:"bytes,4,opt,name=rule_id,json=ruleId,proto3" json:"rule_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *AddUEHeaderEnrichmentRequest) Reset()         { *m = AddUEHeaderEnrichmentRequest{} }
func (m *AddUEHeaderEnrichmentRequest) String() string { return proto.CompactTextString(m) }
func (*AddUEHeaderEnrichmentRequest) ProtoMessage()    {}
func (*AddUEHeaderEnrichmentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c190610d29559b01, []int{1}
}

func (m *AddUEHeaderEnrichmentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddUEHeaderEnrichmentRequest.Unmarshal(m, b)
}
func (m *AddUEHeaderEnrichmentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddUEHeaderEnrichmentRequest.Marshal(b, m, deterministic)
}
func (m *AddUEHeaderEnrichmentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddUEHeaderEnrichmentRequest.Merge(m, src)
}
func (m *AddUEHeaderEnrichmentRequest) XXX_Size() int {
	return xxx_messageInfo_AddUEHeaderEnrichmentRequest.Size(m)
}
func (m *AddUEHeaderEnrichmentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddUEHeaderEnrichmentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddUEHeaderEnrichmentRequest proto.InternalMessageInfo

func (m *AddUEHeaderEnrichmentRequest) GetUeIp() *protos.IPAddress {
	if m != nil {
		return m.UeIp
	}
	return nil
}

func (m *AddUEHeaderEnrichmentRequest) GetWebsites() []string {
	if m != nil {
		return m.Websites
	}
	return nil
}

func (m *AddUEHeaderEnrichmentRequest) GetHeaders() []*Header {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *AddUEHeaderEnrichmentRequest) GetRuleId() string {
	if m != nil {
		return m.RuleId
	}
	return ""
}

type AddUEHeaderEnrichmentResult struct {
	Result               AddUEHeaderEnrichmentResult_Result `protobuf:"varint,1,opt,name=result,proto3,enum=magma.feg.AddUEHeaderEnrichmentResult_Result" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *AddUEHeaderEnrichmentResult) Reset()         { *m = AddUEHeaderEnrichmentResult{} }
func (m *AddUEHeaderEnrichmentResult) String() string { return proto.CompactTextString(m) }
func (*AddUEHeaderEnrichmentResult) ProtoMessage()    {}
func (*AddUEHeaderEnrichmentResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_c190610d29559b01, []int{2}
}

func (m *AddUEHeaderEnrichmentResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddUEHeaderEnrichmentResult.Unmarshal(m, b)
}
func (m *AddUEHeaderEnrichmentResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddUEHeaderEnrichmentResult.Marshal(b, m, deterministic)
}
func (m *AddUEHeaderEnrichmentResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddUEHeaderEnrichmentResult.Merge(m, src)
}
func (m *AddUEHeaderEnrichmentResult) XXX_Size() int {
	return xxx_messageInfo_AddUEHeaderEnrichmentResult.Size(m)
}
func (m *AddUEHeaderEnrichmentResult) XXX_DiscardUnknown() {
	xxx_messageInfo_AddUEHeaderEnrichmentResult.DiscardUnknown(m)
}

var xxx_messageInfo_AddUEHeaderEnrichmentResult proto.InternalMessageInfo

func (m *AddUEHeaderEnrichmentResult) GetResult() AddUEHeaderEnrichmentResult_Result {
	if m != nil {
		return m.Result
	}
	return AddUEHeaderEnrichmentResult_SUCCESS
}

type DeactivateUEHeaderEnrichmentRequest struct {
	UeIp                 *protos.IPAddress `protobuf:"bytes,1,opt,name=ue_ip,json=ueIp,proto3" json:"ue_ip,omitempty"`
	RuleId               string            `protobuf:"bytes,2,opt,name=rule_id,json=ruleId,proto3" json:"rule_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *DeactivateUEHeaderEnrichmentRequest) Reset()         { *m = DeactivateUEHeaderEnrichmentRequest{} }
func (m *DeactivateUEHeaderEnrichmentRequest) String() string { return proto.CompactTextString(m) }
func (*DeactivateUEHeaderEnrichmentRequest) ProtoMessage()    {}
func (*DeactivateUEHeaderEnrichmentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c190610d29559b01, []int{3}
}

func (m *DeactivateUEHeaderEnrichmentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeactivateUEHeaderEnrichmentRequest.Unmarshal(m, b)
}
func (m *DeactivateUEHeaderEnrichmentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeactivateUEHeaderEnrichmentRequest.Marshal(b, m, deterministic)
}
func (m *DeactivateUEHeaderEnrichmentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeactivateUEHeaderEnrichmentRequest.Merge(m, src)
}
func (m *DeactivateUEHeaderEnrichmentRequest) XXX_Size() int {
	return xxx_messageInfo_DeactivateUEHeaderEnrichmentRequest.Size(m)
}
func (m *DeactivateUEHeaderEnrichmentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeactivateUEHeaderEnrichmentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeactivateUEHeaderEnrichmentRequest proto.InternalMessageInfo

func (m *DeactivateUEHeaderEnrichmentRequest) GetUeIp() *protos.IPAddress {
	if m != nil {
		return m.UeIp
	}
	return nil
}

func (m *DeactivateUEHeaderEnrichmentRequest) GetRuleId() string {
	if m != nil {
		return m.RuleId
	}
	return ""
}

type DeactivateUEHeaderEnrichmentResult struct {
	Result               DeactivateUEHeaderEnrichmentResult_Result `protobuf:"varint,1,opt,name=result,proto3,enum=magma.feg.DeactivateUEHeaderEnrichmentResult_Result" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                  `json:"-"`
	XXX_unrecognized     []byte                                    `json:"-"`
	XXX_sizecache        int32                                     `json:"-"`
}

func (m *DeactivateUEHeaderEnrichmentResult) Reset()         { *m = DeactivateUEHeaderEnrichmentResult{} }
func (m *DeactivateUEHeaderEnrichmentResult) String() string { return proto.CompactTextString(m) }
func (*DeactivateUEHeaderEnrichmentResult) ProtoMessage()    {}
func (*DeactivateUEHeaderEnrichmentResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_c190610d29559b01, []int{4}
}

func (m *DeactivateUEHeaderEnrichmentResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeactivateUEHeaderEnrichmentResult.Unmarshal(m, b)
}
func (m *DeactivateUEHeaderEnrichmentResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeactivateUEHeaderEnrichmentResult.Marshal(b, m, deterministic)
}
func (m *DeactivateUEHeaderEnrichmentResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeactivateUEHeaderEnrichmentResult.Merge(m, src)
}
func (m *DeactivateUEHeaderEnrichmentResult) XXX_Size() int {
	return xxx_messageInfo_DeactivateUEHeaderEnrichmentResult.Size(m)
}
func (m *DeactivateUEHeaderEnrichmentResult) XXX_DiscardUnknown() {
	xxx_messageInfo_DeactivateUEHeaderEnrichmentResult.DiscardUnknown(m)
}

var xxx_messageInfo_DeactivateUEHeaderEnrichmentResult proto.InternalMessageInfo

func (m *DeactivateUEHeaderEnrichmentResult) GetResult() DeactivateUEHeaderEnrichmentResult_Result {
	if m != nil {
		return m.Result
	}
	return DeactivateUEHeaderEnrichmentResult_SUCCESS
}

func init() {
	proto.RegisterEnum("magma.feg.AddUEHeaderEnrichmentResult_Result", AddUEHeaderEnrichmentResult_Result_name, AddUEHeaderEnrichmentResult_Result_value)
	proto.RegisterEnum("magma.feg.DeactivateUEHeaderEnrichmentResult_Result", DeactivateUEHeaderEnrichmentResult_Result_name, DeactivateUEHeaderEnrichmentResult_Result_value)
	proto.RegisterType((*Header)(nil), "magma.feg.Header")
	proto.RegisterType((*AddUEHeaderEnrichmentRequest)(nil), "magma.feg.AddUEHeaderEnrichmentRequest")
	proto.RegisterType((*AddUEHeaderEnrichmentResult)(nil), "magma.feg.AddUEHeaderEnrichmentResult")
	proto.RegisterType((*DeactivateUEHeaderEnrichmentRequest)(nil), "magma.feg.DeactivateUEHeaderEnrichmentRequest")
	proto.RegisterType((*DeactivateUEHeaderEnrichmentResult)(nil), "magma.feg.DeactivateUEHeaderEnrichmentResult")
}

func init() { proto.RegisterFile("feg/protos/envoy_controller.proto", fileDescriptor_c190610d29559b01) }

var fileDescriptor_c190610d29559b01 = []byte{
	// 485 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0x51, 0x8f, 0xd2, 0x40,
	0x10, 0xc7, 0x29, 0x70, 0x20, 0x83, 0x39, 0xeb, 0x06, 0x63, 0xe5, 0xee, 0x01, 0x6b, 0xa2, 0x18,
	0x73, 0x25, 0xa9, 0x7e, 0x01, 0xae, 0x2c, 0xb9, 0x26, 0x04, 0x4c, 0xa1, 0x31, 0xf1, 0xa5, 0x29,
	0xec, 0x1c, 0x57, 0xb3, 0x6d, 0xb1, 0xdd, 0x62, 0xee, 0xc9, 0x0f, 0xe4, 0xab, 0x8f, 0x7e, 0x36,
	0x63, 0xba, 0xe5, 0x38, 0xee, 0x72, 0x72, 0x67, 0xe2, 0xd3, 0xee, 0x4e, 0xe6, 0xbf, 0xff, 0x99,
	0xdf, 0x76, 0x0a, 0x2f, 0xcf, 0x71, 0xd9, 0x5b, 0x25, 0xb1, 0x88, 0xd3, 0x1e, 0x46, 0xeb, 0xf8,
	0xd2, 0x5b, 0xc4, 0x91, 0x48, 0x62, 0xce, 0x31, 0x31, 0x64, 0x9c, 0x34, 0x42, 0x7f, 0x19, 0xfa,
	0xc6, 0x39, 0x2e, 0xdb, 0x6d, 0x2e, 0xf0, 0x2a, 0x3b, 0x8c, 0xe7, 0x01, 0x0f, 0xc4, 0x25, 0x2b,
	0xd2, 0x74, 0x13, 0x6a, 0x67, 0xe8, 0x33, 0x4c, 0x08, 0x81, 0x6a, 0xe4, 0x87, 0xa8, 0x29, 0x1d,
	0xa5, 0xdb, 0x70, 0xe4, 0x9e, 0xb4, 0xe0, 0x60, 0xed, 0xf3, 0x0c, 0xb5, 0xb2, 0x0c, 0x16, 0x07,
	0xfd, 0x87, 0x02, 0xc7, 0x7d, 0xc6, 0x5c, 0x5a, 0x28, 0x69, 0x94, 0x04, 0x8b, 0x8b, 0x10, 0x23,
	0xe1, 0xe0, 0xd7, 0x0c, 0x53, 0x41, 0xde, 0xc2, 0x41, 0x86, 0x5e, 0xb0, 0x92, 0x77, 0x35, 0xcd,
	0x96, 0x51, 0xd4, 0xc2, 0x05, 0x1a, 0xf6, 0xc7, 0x3e, 0x63, 0x09, 0xa6, 0xa9, 0x53, 0xcd, 0xd0,
	0x5e, 0x91, 0x36, 0x3c, 0xfa, 0x86, 0xf3, 0x34, 0x10, 0x98, 0x6a, 0xe5, 0x4e, 0xa5, 0xdb, 0x70,
	0xb6, 0x67, 0xf2, 0x0e, 0xea, 0x17, 0xd2, 0x21, 0xd5, 0x2a, 0x9d, 0x4a, 0xb7, 0x69, 0x3e, 0x35,
	0xb6, 0x4d, 0x19, 0x85, 0xb7, 0x73, 0x95, 0x41, 0x9e, 0x43, 0x3d, 0xc9, 0x38, 0x7a, 0x01, 0xd3,
	0xaa, 0xb2, 0xd8, 0x5a, 0x7e, 0xb4, 0x99, 0xfe, 0x53, 0x81, 0xa3, 0xbf, 0x54, 0x9b, 0x66, 0x5c,
	0x10, 0x0a, 0xb5, 0x44, 0xee, 0x64, 0xb5, 0x87, 0xe6, 0xc9, 0x8e, 0xc9, 0x1e, 0x9d, 0x51, 0x2c,
	0xce, 0x46, 0xac, 0x8f, 0xa1, 0xb6, 0xb9, 0xb0, 0x09, 0xf5, 0xa9, 0x6b, 0x59, 0x74, 0x3a, 0x55,
	0x4b, 0xa4, 0x05, 0xea, 0x27, 0x7a, 0x3a, 0xb5, 0x67, 0xd4, 0xb3, 0x26, 0xe3, 0xe1, 0xc8, 0xb6,
	0x66, 0xaa, 0x92, 0x47, 0x1d, 0x77, 0x44, 0x3d, 0x7b, 0x70, 0x1d, 0x2d, 0xe7, 0xc2, 0x61, 0xdf,
	0x1e, 0xb9, 0x0e, 0x55, 0x2b, 0x7a, 0x00, 0xaf, 0x06, 0xe8, 0x2f, 0x44, 0xb0, 0xf6, 0x05, 0xfe,
	0x17, 0xd4, 0x3b, 0x84, 0xca, 0x37, 0x08, 0xfd, 0x52, 0x40, 0xdf, 0xef, 0x25, 0xfb, 0x1a, 0xdd,
	0x02, 0xf5, 0x61, 0x07, 0xd4, 0xfd, 0xf2, 0xdb, 0xbc, 0xce, 0xee, 0xe6, 0xa5, 0xc2, 0x63, 0x97,
	0x7a, 0xe3, 0xc9, 0xcc, 0x1b, 0x4e, 0xdc, 0xf1, 0x40, 0x55, 0x08, 0x81, 0x43, 0xc9, 0xea, 0x3a,
	0x76, 0x93, 0x94, 0xf9, 0x5b, 0x81, 0x27, 0x34, 0x1f, 0x02, 0x6b, 0x3b, 0x03, 0xe4, 0x0b, 0x3c,
	0xbb, 0xf3, 0xed, 0xc8, 0x9b, 0xfb, 0x5f, 0x57, 0x82, 0x6d, 0xbf, 0x7e, 0xd8, 0x67, 0xa0, 0x97,
	0xc8, 0x77, 0x38, 0xde, 0xd7, 0x3e, 0x31, 0x1e, 0xcc, 0xa9, 0x70, 0x3e, 0xf9, 0x27, 0xae, 0x7a,
	0xe9, 0xf4, 0xe8, 0xf3, 0x0b, 0xa9, 0xe8, 0xe5, 0x7f, 0x85, 0x05, 0x8f, 0x33, 0xd6, 0x5b, 0xc6,
	0x9b, 0x81, 0x9f, 0xd7, 0xe4, 0xfa, 0xfe, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x53, 0x5e, 0x6c,
	0xca, 0x33, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EnvoyControllerClient is the client API for EnvoyController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EnvoyControllerClient interface {
	// Add UE to the header enrichment configuration
	AddUEHeaderEnrichment(ctx context.Context, in *AddUEHeaderEnrichmentRequest, opts ...grpc.CallOption) (*AddUEHeaderEnrichmentResult, error)
	// Remove UE from the header enrichment configuration
	DeactivateUEHeaderEnrichment(ctx context.Context, in *DeactivateUEHeaderEnrichmentRequest, opts ...grpc.CallOption) (*DeactivateUEHeaderEnrichmentResult, error)
}

type envoyControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewEnvoyControllerClient(cc grpc.ClientConnInterface) EnvoyControllerClient {
	return &envoyControllerClient{cc}
}

func (c *envoyControllerClient) AddUEHeaderEnrichment(ctx context.Context, in *AddUEHeaderEnrichmentRequest, opts ...grpc.CallOption) (*AddUEHeaderEnrichmentResult, error) {
	out := new(AddUEHeaderEnrichmentResult)
	err := c.cc.Invoke(ctx, "/magma.feg.EnvoyController/AddUEHeaderEnrichment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *envoyControllerClient) DeactivateUEHeaderEnrichment(ctx context.Context, in *DeactivateUEHeaderEnrichmentRequest, opts ...grpc.CallOption) (*DeactivateUEHeaderEnrichmentResult, error) {
	out := new(DeactivateUEHeaderEnrichmentResult)
	err := c.cc.Invoke(ctx, "/magma.feg.EnvoyController/DeactivateUEHeaderEnrichment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EnvoyControllerServer is the server API for EnvoyController service.
type EnvoyControllerServer interface {
	// Add UE to the header enrichment configuration
	AddUEHeaderEnrichment(context.Context, *AddUEHeaderEnrichmentRequest) (*AddUEHeaderEnrichmentResult, error)
	// Remove UE from the header enrichment configuration
	DeactivateUEHeaderEnrichment(context.Context, *DeactivateUEHeaderEnrichmentRequest) (*DeactivateUEHeaderEnrichmentResult, error)
}

// UnimplementedEnvoyControllerServer can be embedded to have forward compatible implementations.
type UnimplementedEnvoyControllerServer struct {
}

func (*UnimplementedEnvoyControllerServer) AddUEHeaderEnrichment(ctx context.Context, req *AddUEHeaderEnrichmentRequest) (*AddUEHeaderEnrichmentResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUEHeaderEnrichment not implemented")
}
func (*UnimplementedEnvoyControllerServer) DeactivateUEHeaderEnrichment(ctx context.Context, req *DeactivateUEHeaderEnrichmentRequest) (*DeactivateUEHeaderEnrichmentResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeactivateUEHeaderEnrichment not implemented")
}

func RegisterEnvoyControllerServer(s *grpc.Server, srv EnvoyControllerServer) {
	s.RegisterService(&_EnvoyController_serviceDesc, srv)
}

func _EnvoyController_AddUEHeaderEnrichment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUEHeaderEnrichmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvoyControllerServer).AddUEHeaderEnrichment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.feg.EnvoyController/AddUEHeaderEnrichment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvoyControllerServer).AddUEHeaderEnrichment(ctx, req.(*AddUEHeaderEnrichmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvoyController_DeactivateUEHeaderEnrichment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeactivateUEHeaderEnrichmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvoyControllerServer).DeactivateUEHeaderEnrichment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.feg.EnvoyController/DeactivateUEHeaderEnrichment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvoyControllerServer).DeactivateUEHeaderEnrichment(ctx, req.(*DeactivateUEHeaderEnrichmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EnvoyController_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.feg.EnvoyController",
	HandlerType: (*EnvoyControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddUEHeaderEnrichment",
			Handler:    _EnvoyController_AddUEHeaderEnrichment_Handler,
		},
		{
			MethodName: "DeactivateUEHeaderEnrichment",
			Handler:    _EnvoyController_DeactivateUEHeaderEnrichment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "feg/protos/envoy_controller.proto",
}
