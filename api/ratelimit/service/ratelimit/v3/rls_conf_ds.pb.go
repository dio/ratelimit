// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: ratelimit/service/ratelimit/v3/rls_conf_ds.proto

package ratelimitv3

import (
	context "context"
	v3 "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_ratelimit_service_ratelimit_v3_rls_conf_ds_proto protoreflect.FileDescriptor

var file_ratelimit_service_ratelimit_v3_rls_conf_ds_proto_rawDesc = []byte{
	0x0a, 0x30, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2f, 0x76, 0x33,
	0x2f, 0x72, 0x6c, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x5f, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e,
	0x76, 0x33, 0x1a, 0x2a, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x76, 0x33, 0x2f, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x8a,
	0x02, 0x0a, 0x1f, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x75, 0x0a, 0x10, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x6c, 0x73, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x2c, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x2e, 0x76, 0x33, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76,
	0x33, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x70, 0x0a, 0x0f, 0x46, 0x65, 0x74,
	0x63, 0x68, 0x52, 0x6c, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x2c, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x53, 0x5a, 0x51, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d,
	0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x2f, 0x76, 0x33, 0x3b, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x76, 0x33,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_ratelimit_service_ratelimit_v3_rls_conf_ds_proto_goTypes = []interface{}{
	(*v3.DiscoveryRequest)(nil),  // 0: envoy.service.discovery.v3.DiscoveryRequest
	(*v3.DiscoveryResponse)(nil), // 1: envoy.service.discovery.v3.DiscoveryResponse
}
var file_ratelimit_service_ratelimit_v3_rls_conf_ds_proto_depIdxs = []int32{
	0, // 0: ratelimit.service.ratelimit.v3.RateLimitConfigDiscoveryService.StreamRlsConfigs:input_type -> envoy.service.discovery.v3.DiscoveryRequest
	0, // 1: ratelimit.service.ratelimit.v3.RateLimitConfigDiscoveryService.FetchRlsConfigs:input_type -> envoy.service.discovery.v3.DiscoveryRequest
	1, // 2: ratelimit.service.ratelimit.v3.RateLimitConfigDiscoveryService.StreamRlsConfigs:output_type -> envoy.service.discovery.v3.DiscoveryResponse
	1, // 3: ratelimit.service.ratelimit.v3.RateLimitConfigDiscoveryService.FetchRlsConfigs:output_type -> envoy.service.discovery.v3.DiscoveryResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ratelimit_service_ratelimit_v3_rls_conf_ds_proto_init() }
func file_ratelimit_service_ratelimit_v3_rls_conf_ds_proto_init() {
	if File_ratelimit_service_ratelimit_v3_rls_conf_ds_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ratelimit_service_ratelimit_v3_rls_conf_ds_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ratelimit_service_ratelimit_v3_rls_conf_ds_proto_goTypes,
		DependencyIndexes: file_ratelimit_service_ratelimit_v3_rls_conf_ds_proto_depIdxs,
	}.Build()
	File_ratelimit_service_ratelimit_v3_rls_conf_ds_proto = out.File
	file_ratelimit_service_ratelimit_v3_rls_conf_ds_proto_rawDesc = nil
	file_ratelimit_service_ratelimit_v3_rls_conf_ds_proto_goTypes = nil
	file_ratelimit_service_ratelimit_v3_rls_conf_ds_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RateLimitConfigDiscoveryServiceClient is the client API for RateLimitConfigDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RateLimitConfigDiscoveryServiceClient interface {
	StreamRlsConfigs(ctx context.Context, opts ...grpc.CallOption) (RateLimitConfigDiscoveryService_StreamRlsConfigsClient, error)
	FetchRlsConfigs(ctx context.Context, in *v3.DiscoveryRequest, opts ...grpc.CallOption) (*v3.DiscoveryResponse, error)
}

type rateLimitConfigDiscoveryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRateLimitConfigDiscoveryServiceClient(cc grpc.ClientConnInterface) RateLimitConfigDiscoveryServiceClient {
	return &rateLimitConfigDiscoveryServiceClient{cc}
}

func (c *rateLimitConfigDiscoveryServiceClient) StreamRlsConfigs(ctx context.Context, opts ...grpc.CallOption) (RateLimitConfigDiscoveryService_StreamRlsConfigsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RateLimitConfigDiscoveryService_serviceDesc.Streams[0], "/ratelimit.service.ratelimit.v3.RateLimitConfigDiscoveryService/StreamRlsConfigs", opts...)
	if err != nil {
		return nil, err
	}
	x := &rateLimitConfigDiscoveryServiceStreamRlsConfigsClient{stream}
	return x, nil
}

type RateLimitConfigDiscoveryService_StreamRlsConfigsClient interface {
	Send(*v3.DiscoveryRequest) error
	Recv() (*v3.DiscoveryResponse, error)
	grpc.ClientStream
}

type rateLimitConfigDiscoveryServiceStreamRlsConfigsClient struct {
	grpc.ClientStream
}

func (x *rateLimitConfigDiscoveryServiceStreamRlsConfigsClient) Send(m *v3.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *rateLimitConfigDiscoveryServiceStreamRlsConfigsClient) Recv() (*v3.DiscoveryResponse, error) {
	m := new(v3.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rateLimitConfigDiscoveryServiceClient) FetchRlsConfigs(ctx context.Context, in *v3.DiscoveryRequest, opts ...grpc.CallOption) (*v3.DiscoveryResponse, error) {
	out := new(v3.DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/ratelimit.service.ratelimit.v3.RateLimitConfigDiscoveryService/FetchRlsConfigs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RateLimitConfigDiscoveryServiceServer is the server API for RateLimitConfigDiscoveryService service.
type RateLimitConfigDiscoveryServiceServer interface {
	StreamRlsConfigs(RateLimitConfigDiscoveryService_StreamRlsConfigsServer) error
	FetchRlsConfigs(context.Context, *v3.DiscoveryRequest) (*v3.DiscoveryResponse, error)
}

// UnimplementedRateLimitConfigDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRateLimitConfigDiscoveryServiceServer struct {
}

func (*UnimplementedRateLimitConfigDiscoveryServiceServer) StreamRlsConfigs(RateLimitConfigDiscoveryService_StreamRlsConfigsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamRlsConfigs not implemented")
}
func (*UnimplementedRateLimitConfigDiscoveryServiceServer) FetchRlsConfigs(context.Context, *v3.DiscoveryRequest) (*v3.DiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchRlsConfigs not implemented")
}

func RegisterRateLimitConfigDiscoveryServiceServer(s *grpc.Server, srv RateLimitConfigDiscoveryServiceServer) {
	s.RegisterService(&_RateLimitConfigDiscoveryService_serviceDesc, srv)
}

func _RateLimitConfigDiscoveryService_StreamRlsConfigs_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RateLimitConfigDiscoveryServiceServer).StreamRlsConfigs(&rateLimitConfigDiscoveryServiceStreamRlsConfigsServer{stream})
}

type RateLimitConfigDiscoveryService_StreamRlsConfigsServer interface {
	Send(*v3.DiscoveryResponse) error
	Recv() (*v3.DiscoveryRequest, error)
	grpc.ServerStream
}

type rateLimitConfigDiscoveryServiceStreamRlsConfigsServer struct {
	grpc.ServerStream
}

func (x *rateLimitConfigDiscoveryServiceStreamRlsConfigsServer) Send(m *v3.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *rateLimitConfigDiscoveryServiceStreamRlsConfigsServer) Recv() (*v3.DiscoveryRequest, error) {
	m := new(v3.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RateLimitConfigDiscoveryService_FetchRlsConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateLimitConfigDiscoveryServiceServer).FetchRlsConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ratelimit.service.ratelimit.v3.RateLimitConfigDiscoveryService/FetchRlsConfigs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateLimitConfigDiscoveryServiceServer).FetchRlsConfigs(ctx, req.(*v3.DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RateLimitConfigDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ratelimit.service.ratelimit.v3.RateLimitConfigDiscoveryService",
	HandlerType: (*RateLimitConfigDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchRlsConfigs",
			Handler:    _RateLimitConfigDiscoveryService_FetchRlsConfigs_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamRlsConfigs",
			Handler:       _RateLimitConfigDiscoveryService_StreamRlsConfigs_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "ratelimit/service/ratelimit/v3/rls_conf_ds.proto",
}
