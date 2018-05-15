// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: connect.proto

/*
Package connect is a generated protocol buffer package.

It is generated from these files:
	connect.proto

It has these top-level messages:
*/
package connect

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import kafka_scheduler_v1 "github.com/confluentinc/cc-structs/kafka/scheduler/v1"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Connect service

type ConnectClient interface {
	List(ctx context.Context, in *kafka_scheduler_v1.GetConnectClustersRequest, opts ...grpc.CallOption) (*kafka_scheduler_v1.GetConnectClustersReply, error)
	Describe(ctx context.Context, in *kafka_scheduler_v1.GetConnectClusterRequest, opts ...grpc.CallOption) (*kafka_scheduler_v1.GetConnectClusterReply, error)
	CreateS3Sink(ctx context.Context, in *kafka_scheduler_v1.CreateConnectS3SinkClusterRequest, opts ...grpc.CallOption) (*kafka_scheduler_v1.CreateConnectS3SinkClusterReply, error)
	UpdateS3Sink(ctx context.Context, in *kafka_scheduler_v1.UpdateConnectS3SinkClusterRequest, opts ...grpc.CallOption) (*kafka_scheduler_v1.UpdateConnectS3SinkClusterReply, error)
	Delete(ctx context.Context, in *kafka_scheduler_v1.DeleteConnectClusterRequest, opts ...grpc.CallOption) (*kafka_scheduler_v1.DeleteConnectClusterReply, error)
}

type connectClient struct {
	cc *grpc.ClientConn
}

func NewConnectClient(cc *grpc.ClientConn) ConnectClient {
	return &connectClient{cc}
}

func (c *connectClient) List(ctx context.Context, in *kafka_scheduler_v1.GetConnectClustersRequest, opts ...grpc.CallOption) (*kafka_scheduler_v1.GetConnectClustersReply, error) {
	out := new(kafka_scheduler_v1.GetConnectClustersReply)
	err := grpc.Invoke(ctx, "/connect.Connect/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectClient) Describe(ctx context.Context, in *kafka_scheduler_v1.GetConnectClusterRequest, opts ...grpc.CallOption) (*kafka_scheduler_v1.GetConnectClusterReply, error) {
	out := new(kafka_scheduler_v1.GetConnectClusterReply)
	err := grpc.Invoke(ctx, "/connect.Connect/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectClient) CreateS3Sink(ctx context.Context, in *kafka_scheduler_v1.CreateConnectS3SinkClusterRequest, opts ...grpc.CallOption) (*kafka_scheduler_v1.CreateConnectS3SinkClusterReply, error) {
	out := new(kafka_scheduler_v1.CreateConnectS3SinkClusterReply)
	err := grpc.Invoke(ctx, "/connect.Connect/CreateS3Sink", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectClient) UpdateS3Sink(ctx context.Context, in *kafka_scheduler_v1.UpdateConnectS3SinkClusterRequest, opts ...grpc.CallOption) (*kafka_scheduler_v1.UpdateConnectS3SinkClusterReply, error) {
	out := new(kafka_scheduler_v1.UpdateConnectS3SinkClusterReply)
	err := grpc.Invoke(ctx, "/connect.Connect/UpdateS3Sink", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectClient) Delete(ctx context.Context, in *kafka_scheduler_v1.DeleteConnectClusterRequest, opts ...grpc.CallOption) (*kafka_scheduler_v1.DeleteConnectClusterReply, error) {
	out := new(kafka_scheduler_v1.DeleteConnectClusterReply)
	err := grpc.Invoke(ctx, "/connect.Connect/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Connect service

type ConnectServer interface {
	List(context.Context, *kafka_scheduler_v1.GetConnectClustersRequest) (*kafka_scheduler_v1.GetConnectClustersReply, error)
	Describe(context.Context, *kafka_scheduler_v1.GetConnectClusterRequest) (*kafka_scheduler_v1.GetConnectClusterReply, error)
	CreateS3Sink(context.Context, *kafka_scheduler_v1.CreateConnectS3SinkClusterRequest) (*kafka_scheduler_v1.CreateConnectS3SinkClusterReply, error)
	UpdateS3Sink(context.Context, *kafka_scheduler_v1.UpdateConnectS3SinkClusterRequest) (*kafka_scheduler_v1.UpdateConnectS3SinkClusterReply, error)
	Delete(context.Context, *kafka_scheduler_v1.DeleteConnectClusterRequest) (*kafka_scheduler_v1.DeleteConnectClusterReply, error)
}

func RegisterConnectServer(s *grpc.Server, srv ConnectServer) {
	s.RegisterService(&_Connect_serviceDesc, srv)
}

func _Connect_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kafka_scheduler_v1.GetConnectClustersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connect.Connect/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectServer).List(ctx, req.(*kafka_scheduler_v1.GetConnectClustersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Connect_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kafka_scheduler_v1.GetConnectClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connect.Connect/Describe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectServer).Describe(ctx, req.(*kafka_scheduler_v1.GetConnectClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Connect_CreateS3Sink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kafka_scheduler_v1.CreateConnectS3SinkClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectServer).CreateS3Sink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connect.Connect/CreateS3Sink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectServer).CreateS3Sink(ctx, req.(*kafka_scheduler_v1.CreateConnectS3SinkClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Connect_UpdateS3Sink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kafka_scheduler_v1.UpdateConnectS3SinkClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectServer).UpdateS3Sink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connect.Connect/UpdateS3Sink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectServer).UpdateS3Sink(ctx, req.(*kafka_scheduler_v1.UpdateConnectS3SinkClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Connect_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kafka_scheduler_v1.DeleteConnectClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connect.Connect/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectServer).Delete(ctx, req.(*kafka_scheduler_v1.DeleteConnectClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Connect_serviceDesc = grpc.ServiceDesc{
	ServiceName: "connect.Connect",
	HandlerType: (*ConnectServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Connect_List_Handler,
		},
		{
			MethodName: "Describe",
			Handler:    _Connect_Describe_Handler,
		},
		{
			MethodName: "CreateS3Sink",
			Handler:    _Connect_CreateS3Sink_Handler,
		},
		{
			MethodName: "UpdateS3Sink",
			Handler:    _Connect_UpdateS3Sink_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Connect_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "connect.proto",
}

func init() { proto.RegisterFile("connect.proto", fileDescriptorConnect) }

var fileDescriptorConnect = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x6a, 0x84, 0x30,
	0x10, 0x40, 0x7b, 0x10, 0x2d, 0xa1, 0xbd, 0xe4, 0xe8, 0xd1, 0x63, 0x5b, 0x23, 0x56, 0xfa, 0x05,
	0x0a, 0xbd, 0xf4, 0x54, 0xe9, 0x07, 0x68, 0x1c, 0xa9, 0x35, 0xa8, 0x4d, 0x46, 0xa1, 0xd0, 0x7f,
	0xdb, 0x5f, 0x5b, 0x70, 0xdc, 0x15, 0x16, 0xd9, 0x8d, 0xc7, 0x21, 0xef, 0xcd, 0x83, 0x30, 0xec,
	0x51, 0xf6, 0x5d, 0x07, 0x12, 0xc5, 0xa0, 0x7b, 0xec, 0xb9, 0xb7, 0x8c, 0x7e, 0xd0, 0x16, 0x75,
	0x5b, 0x44, 0x46, 0x7e, 0x43, 0x35, 0x2a, 0xd0, 0xd1, 0x14, 0xaf, 0x03, 0xc1, 0xaf, 0x07, 0x87,
	0x79, 0x29, 0xf1, 0xbc, 0x62, 0xce, 0x47, 0x63, 0x90, 0x87, 0x62, 0x16, 0xc5, 0xca, 0x4e, 0xb1,
	0x78, 0x07, 0x5c, 0xb8, 0x54, 0x8d, 0x06, 0x41, 0x9b, 0x4f, 0xf8, 0x1d, 0xc1, 0xa0, 0xff, 0x6c,
	0x8b, 0x0f, 0xea, 0x2f, 0xb8, 0xe3, 0x35, 0xbb, 0xcf, 0xc0, 0x48, 0xdd, 0x94, 0xc0, 0x5f, 0xac,
	0xd4, 0x53, 0xe8, 0xc9, 0x92, 0xa6, 0xce, 0x3f, 0x7b, 0x48, 0x35, 0x14, 0x08, 0x79, 0x92, 0x37,
	0x5d, 0xcb, 0xdf, 0xb6, 0x6c, 0x22, 0x96, 0x05, 0x04, 0x5e, 0x44, 0x93, 0xbd, 0xda, 0xb9, 0xfe,
	0x35, 0x54, 0x37, 0xea, 0x44, 0xec, 0xae, 0x5f, 0xd3, 0xa8, 0xfe, 0xc3, 0xdc, 0x0c, 0x14, 0x20,
	0xf0, 0x68, 0x6b, 0x01, 0xbd, 0x6d, 0x7f, 0x72, 0x68, 0x2f, 0xcc, 0xad, 0xd2, 0x9d, 0x0f, 0x29,
	0x39, 0x06, 0x00, 0x00, 0xff, 0xff, 0x0a, 0xc7, 0x9f, 0x1e, 0x86, 0x02, 0x00, 0x00,
}
