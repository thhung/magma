//
//Copyright 2020 The Magma Authors.
//
//This source code is licensed under the BSD-style license found in the
//LICENSE file in the root directory of this source tree.
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.10.0
// source: orc8r/cloud/go/obsidian/swagger/protos/swagger.proto

package protos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PartialSpecRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PartialSpecRequest) Reset() {
	*x = PartialSpecRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_cloud_go_obsidian_swagger_protos_swagger_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartialSpecRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartialSpecRequest) ProtoMessage() {}

func (x *PartialSpecRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_cloud_go_obsidian_swagger_protos_swagger_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartialSpecRequest.ProtoReflect.Descriptor instead.
func (*PartialSpecRequest) Descriptor() ([]byte, []int) {
	return file_orc8r_cloud_go_obsidian_swagger_protos_swagger_proto_rawDescGZIP(), []int{0}
}

type PartialSpecResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SwaggerSpec string `protobuf:"bytes,1,opt,name=swagger_spec,json=swaggerSpec,proto3" json:"swagger_spec,omitempty"`
}

func (x *PartialSpecResponse) Reset() {
	*x = PartialSpecResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_cloud_go_obsidian_swagger_protos_swagger_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartialSpecResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartialSpecResponse) ProtoMessage() {}

func (x *PartialSpecResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_cloud_go_obsidian_swagger_protos_swagger_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartialSpecResponse.ProtoReflect.Descriptor instead.
func (*PartialSpecResponse) Descriptor() ([]byte, []int) {
	return file_orc8r_cloud_go_obsidian_swagger_protos_swagger_proto_rawDescGZIP(), []int{1}
}

func (x *PartialSpecResponse) GetSwaggerSpec() string {
	if x != nil {
		return x.SwaggerSpec
	}
	return ""
}

type StandaloneSpecRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StandaloneSpecRequest) Reset() {
	*x = StandaloneSpecRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_cloud_go_obsidian_swagger_protos_swagger_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StandaloneSpecRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StandaloneSpecRequest) ProtoMessage() {}

func (x *StandaloneSpecRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_cloud_go_obsidian_swagger_protos_swagger_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StandaloneSpecRequest.ProtoReflect.Descriptor instead.
func (*StandaloneSpecRequest) Descriptor() ([]byte, []int) {
	return file_orc8r_cloud_go_obsidian_swagger_protos_swagger_proto_rawDescGZIP(), []int{2}
}

type StandaloneSpecResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SwaggerSpec string `protobuf:"bytes,1,opt,name=swagger_spec,json=swaggerSpec,proto3" json:"swagger_spec,omitempty"`
}

func (x *StandaloneSpecResponse) Reset() {
	*x = StandaloneSpecResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_cloud_go_obsidian_swagger_protos_swagger_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StandaloneSpecResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StandaloneSpecResponse) ProtoMessage() {}

func (x *StandaloneSpecResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_cloud_go_obsidian_swagger_protos_swagger_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StandaloneSpecResponse.ProtoReflect.Descriptor instead.
func (*StandaloneSpecResponse) Descriptor() ([]byte, []int) {
	return file_orc8r_cloud_go_obsidian_swagger_protos_swagger_proto_rawDescGZIP(), []int{3}
}

func (x *StandaloneSpecResponse) GetSwaggerSpec() string {
	if x != nil {
		return x.SwaggerSpec
	}
	return ""
}

var File_orc8r_cloud_go_obsidian_swagger_protos_swagger_proto protoreflect.FileDescriptor

var file_orc8r_cloud_go_obsidian_swagger_protos_swagger_proto_rawDesc = []byte{
	0x0a, 0x34, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f,
	0x2f, 0x6f, 0x62, 0x73, 0x69, 0x64, 0x69, 0x61, 0x6e, 0x2f, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65,
	0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72,
	0x63, 0x38, 0x72, 0x2e, 0x6f, 0x62, 0x73, 0x69, 0x64, 0x69, 0x61, 0x6e, 0x2e, 0x73, 0x77, 0x61,
	0x67, 0x67, 0x65, 0x72, 0x22, 0x14, 0x0a, 0x12, 0x50, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x53,
	0x70, 0x65, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x38, 0x0a, 0x13, 0x50, 0x61,
	0x72, 0x74, 0x69, 0x61, 0x6c, 0x53, 0x70, 0x65, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x73, 0x70, 0x65,
	0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72,
	0x53, 0x70, 0x65, 0x63, 0x22, 0x17, 0x0a, 0x15, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x6c, 0x6f,
	0x6e, 0x65, 0x53, 0x70, 0x65, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3b, 0x0a,
	0x16, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x6c, 0x6f, 0x6e, 0x65, 0x53, 0x70, 0x65, 0x63, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x77, 0x61, 0x67, 0x67,
	0x65, 0x72, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73,
	0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x32, 0x89, 0x02, 0x0a, 0x0b, 0x53,
	0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x12, 0x77, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x50, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x53, 0x70, 0x65, 0x63, 0x12, 0x30, 0x2e, 0x6d,
	0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x6f, 0x62, 0x73, 0x69, 0x64,
	0x69, 0x61, 0x6e, 0x2e, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x72, 0x74,
	0x69, 0x61, 0x6c, 0x53, 0x70, 0x65, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31,
	0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x6f, 0x62, 0x73,
	0x69, 0x64, 0x69, 0x61, 0x6e, 0x2e, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x50, 0x61,
	0x72, 0x74, 0x69, 0x61, 0x6c, 0x53, 0x70, 0x65, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x80, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x6e, 0x64,
	0x61, 0x6c, 0x6f, 0x6e, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x33, 0x2e, 0x6d, 0x61, 0x67, 0x6d,
	0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x6f, 0x62, 0x73, 0x69, 0x64, 0x69, 0x61, 0x6e,
	0x2e, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x6c,
	0x6f, 0x6e, 0x65, 0x53, 0x70, 0x65, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34,
	0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x6f, 0x62, 0x73,
	0x69, 0x64, 0x69, 0x61, 0x6e, 0x2e, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x53, 0x74,
	0x61, 0x6e, 0x64, 0x61, 0x6c, 0x6f, 0x6e, 0x65, 0x53, 0x70, 0x65, 0x63, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2e, 0x5a, 0x2c, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2f,
	0x6f, 0x72, 0x63, 0x38, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x6f,
	0x62, 0x73, 0x69, 0x64, 0x69, 0x61, 0x6e, 0x2f, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orc8r_cloud_go_obsidian_swagger_protos_swagger_proto_rawDescOnce sync.Once
	file_orc8r_cloud_go_obsidian_swagger_protos_swagger_proto_rawDescData = file_orc8r_cloud_go_obsidian_swagger_protos_swagger_proto_rawDesc
)

func file_orc8r_cloud_go_obsidian_swagger_protos_swagger_proto_rawDescGZIP() []byte {
	file_orc8r_cloud_go_obsidian_swagger_protos_swagger_proto_rawDescOnce.Do(func() {
		file_orc8r_cloud_go_obsidian_swagger_protos_swagger_proto_rawDescData = protoimpl.X.CompressGZIP(file_orc8r_cloud_go_obsidian_swagger_protos_swagger_proto_rawDescData)
	})
	return file_orc8r_cloud_go_obsidian_swagger_protos_swagger_proto_rawDescData
}

var file_orc8r_cloud_go_obsidian_swagger_protos_swagger_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_orc8r_cloud_go_obsidian_swagger_protos_swagger_proto_goTypes = []interface{}{
	(*PartialSpecRequest)(nil),     // 0: magma.orc8r.obsidian.swagger.PartialSpecRequest
	(*PartialSpecResponse)(nil),    // 1: magma.orc8r.obsidian.swagger.PartialSpecResponse
	(*StandaloneSpecRequest)(nil),  // 2: magma.orc8r.obsidian.swagger.StandaloneSpecRequest
	(*StandaloneSpecResponse)(nil), // 3: magma.orc8r.obsidian.swagger.StandaloneSpecResponse
}
var file_orc8r_cloud_go_obsidian_swagger_protos_swagger_proto_depIdxs = []int32{
	0, // 0: magma.orc8r.obsidian.swagger.SwaggerSpec.GetPartialSpec:input_type -> magma.orc8r.obsidian.swagger.PartialSpecRequest
	2, // 1: magma.orc8r.obsidian.swagger.SwaggerSpec.GetStandaloneSpec:input_type -> magma.orc8r.obsidian.swagger.StandaloneSpecRequest
	1, // 2: magma.orc8r.obsidian.swagger.SwaggerSpec.GetPartialSpec:output_type -> magma.orc8r.obsidian.swagger.PartialSpecResponse
	3, // 3: magma.orc8r.obsidian.swagger.SwaggerSpec.GetStandaloneSpec:output_type -> magma.orc8r.obsidian.swagger.StandaloneSpecResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_orc8r_cloud_go_obsidian_swagger_protos_swagger_proto_init() }
func file_orc8r_cloud_go_obsidian_swagger_protos_swagger_proto_init() {
	if File_orc8r_cloud_go_obsidian_swagger_protos_swagger_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orc8r_cloud_go_obsidian_swagger_protos_swagger_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartialSpecRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_orc8r_cloud_go_obsidian_swagger_protos_swagger_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartialSpecResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_orc8r_cloud_go_obsidian_swagger_protos_swagger_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StandaloneSpecRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_orc8r_cloud_go_obsidian_swagger_protos_swagger_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StandaloneSpecResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_orc8r_cloud_go_obsidian_swagger_protos_swagger_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_orc8r_cloud_go_obsidian_swagger_protos_swagger_proto_goTypes,
		DependencyIndexes: file_orc8r_cloud_go_obsidian_swagger_protos_swagger_proto_depIdxs,
		MessageInfos:      file_orc8r_cloud_go_obsidian_swagger_protos_swagger_proto_msgTypes,
	}.Build()
	File_orc8r_cloud_go_obsidian_swagger_protos_swagger_proto = out.File
	file_orc8r_cloud_go_obsidian_swagger_protos_swagger_proto_rawDesc = nil
	file_orc8r_cloud_go_obsidian_swagger_protos_swagger_proto_goTypes = nil
	file_orc8r_cloud_go_obsidian_swagger_protos_swagger_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SwaggerSpecClient is the client API for SwaggerSpec service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SwaggerSpecClient interface {
	// GetPartialSpec returns partial Swagger specification of the service's REST API.
	GetPartialSpec(ctx context.Context, in *PartialSpecRequest, opts ...grpc.CallOption) (*PartialSpecResponse, error)
	// GetStandaloneSpec returns a standalone Swagger specification of the
	// service's REST API.
	GetStandaloneSpec(ctx context.Context, in *StandaloneSpecRequest, opts ...grpc.CallOption) (*StandaloneSpecResponse, error)
}

type swaggerSpecClient struct {
	cc grpc.ClientConnInterface
}

func NewSwaggerSpecClient(cc grpc.ClientConnInterface) SwaggerSpecClient {
	return &swaggerSpecClient{cc}
}

func (c *swaggerSpecClient) GetPartialSpec(ctx context.Context, in *PartialSpecRequest, opts ...grpc.CallOption) (*PartialSpecResponse, error) {
	out := new(PartialSpecResponse)
	err := c.cc.Invoke(ctx, "/magma.orc8r.obsidian.swagger.SwaggerSpec/GetPartialSpec", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *swaggerSpecClient) GetStandaloneSpec(ctx context.Context, in *StandaloneSpecRequest, opts ...grpc.CallOption) (*StandaloneSpecResponse, error) {
	out := new(StandaloneSpecResponse)
	err := c.cc.Invoke(ctx, "/magma.orc8r.obsidian.swagger.SwaggerSpec/GetStandaloneSpec", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SwaggerSpecServer is the server API for SwaggerSpec service.
type SwaggerSpecServer interface {
	// GetPartialSpec returns partial Swagger specification of the service's REST API.
	GetPartialSpec(context.Context, *PartialSpecRequest) (*PartialSpecResponse, error)
	// GetStandaloneSpec returns a standalone Swagger specification of the
	// service's REST API.
	GetStandaloneSpec(context.Context, *StandaloneSpecRequest) (*StandaloneSpecResponse, error)
}

// UnimplementedSwaggerSpecServer can be embedded to have forward compatible implementations.
type UnimplementedSwaggerSpecServer struct {
}

func (*UnimplementedSwaggerSpecServer) GetPartialSpec(context.Context, *PartialSpecRequest) (*PartialSpecResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPartialSpec not implemented")
}
func (*UnimplementedSwaggerSpecServer) GetStandaloneSpec(context.Context, *StandaloneSpecRequest) (*StandaloneSpecResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStandaloneSpec not implemented")
}

func RegisterSwaggerSpecServer(s *grpc.Server, srv SwaggerSpecServer) {
	s.RegisterService(&_SwaggerSpec_serviceDesc, srv)
}

func _SwaggerSpec_GetPartialSpec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PartialSpecRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwaggerSpecServer).GetPartialSpec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.obsidian.swagger.SwaggerSpec/GetPartialSpec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwaggerSpecServer).GetPartialSpec(ctx, req.(*PartialSpecRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SwaggerSpec_GetStandaloneSpec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StandaloneSpecRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwaggerSpecServer).GetStandaloneSpec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.obsidian.swagger.SwaggerSpec/GetStandaloneSpec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwaggerSpecServer).GetStandaloneSpec(ctx, req.(*StandaloneSpecRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SwaggerSpec_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.orc8r.obsidian.swagger.SwaggerSpec",
	HandlerType: (*SwaggerSpecServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPartialSpec",
			Handler:    _SwaggerSpec_GetPartialSpec_Handler,
		},
		{
			MethodName: "GetStandaloneSpec",
			Handler:    _SwaggerSpec_GetStandaloneSpec_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orc8r/cloud/go/obsidian/swagger/protos/swagger.proto",
}
