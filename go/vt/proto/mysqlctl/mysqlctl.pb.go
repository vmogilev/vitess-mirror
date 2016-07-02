// Code generated by protoc-gen-go.
// source: mysqlctl.proto
// DO NOT EDIT!

/*
Package mysqlctl is a generated protocol buffer package.

It is generated from these files:
	mysqlctl.proto

It has these top-level messages:
	StartRequest
	StartResponse
	ShutdownRequest
	ShutdownResponse
	RunMysqlUpgradeRequest
	RunMysqlUpgradeResponse
	ReinitConfigRequest
	ReinitConfigResponse
*/
package mysqlctl

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StartRequest struct {
}

func (m *StartRequest) Reset()                    { *m = StartRequest{} }
func (m *StartRequest) String() string            { return proto.CompactTextString(m) }
func (*StartRequest) ProtoMessage()               {}
func (*StartRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type StartResponse struct {
}

func (m *StartResponse) Reset()                    { *m = StartResponse{} }
func (m *StartResponse) String() string            { return proto.CompactTextString(m) }
func (*StartResponse) ProtoMessage()               {}
func (*StartResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ShutdownRequest struct {
	WaitForMysqld bool `protobuf:"varint,1,opt,name=wait_for_mysqld,json=waitForMysqld" json:"wait_for_mysqld,omitempty"`
}

func (m *ShutdownRequest) Reset()                    { *m = ShutdownRequest{} }
func (m *ShutdownRequest) String() string            { return proto.CompactTextString(m) }
func (*ShutdownRequest) ProtoMessage()               {}
func (*ShutdownRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type ShutdownResponse struct {
}

func (m *ShutdownResponse) Reset()                    { *m = ShutdownResponse{} }
func (m *ShutdownResponse) String() string            { return proto.CompactTextString(m) }
func (*ShutdownResponse) ProtoMessage()               {}
func (*ShutdownResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type RunMysqlUpgradeRequest struct {
}

func (m *RunMysqlUpgradeRequest) Reset()                    { *m = RunMysqlUpgradeRequest{} }
func (m *RunMysqlUpgradeRequest) String() string            { return proto.CompactTextString(m) }
func (*RunMysqlUpgradeRequest) ProtoMessage()               {}
func (*RunMysqlUpgradeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type RunMysqlUpgradeResponse struct {
}

func (m *RunMysqlUpgradeResponse) Reset()                    { *m = RunMysqlUpgradeResponse{} }
func (m *RunMysqlUpgradeResponse) String() string            { return proto.CompactTextString(m) }
func (*RunMysqlUpgradeResponse) ProtoMessage()               {}
func (*RunMysqlUpgradeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type ReinitConfigRequest struct {
}

func (m *ReinitConfigRequest) Reset()                    { *m = ReinitConfigRequest{} }
func (m *ReinitConfigRequest) String() string            { return proto.CompactTextString(m) }
func (*ReinitConfigRequest) ProtoMessage()               {}
func (*ReinitConfigRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type ReinitConfigResponse struct {
}

func (m *ReinitConfigResponse) Reset()                    { *m = ReinitConfigResponse{} }
func (m *ReinitConfigResponse) String() string            { return proto.CompactTextString(m) }
func (*ReinitConfigResponse) ProtoMessage()               {}
func (*ReinitConfigResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func init() {
	proto.RegisterType((*StartRequest)(nil), "mysqlctl.StartRequest")
	proto.RegisterType((*StartResponse)(nil), "mysqlctl.StartResponse")
	proto.RegisterType((*ShutdownRequest)(nil), "mysqlctl.ShutdownRequest")
	proto.RegisterType((*ShutdownResponse)(nil), "mysqlctl.ShutdownResponse")
	proto.RegisterType((*RunMysqlUpgradeRequest)(nil), "mysqlctl.RunMysqlUpgradeRequest")
	proto.RegisterType((*RunMysqlUpgradeResponse)(nil), "mysqlctl.RunMysqlUpgradeResponse")
	proto.RegisterType((*ReinitConfigRequest)(nil), "mysqlctl.ReinitConfigRequest")
	proto.RegisterType((*ReinitConfigResponse)(nil), "mysqlctl.ReinitConfigResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for MysqlCtl service

type MysqlCtlClient interface {
	Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error)
	Shutdown(ctx context.Context, in *ShutdownRequest, opts ...grpc.CallOption) (*ShutdownResponse, error)
	RunMysqlUpgrade(ctx context.Context, in *RunMysqlUpgradeRequest, opts ...grpc.CallOption) (*RunMysqlUpgradeResponse, error)
	ReinitConfig(ctx context.Context, in *ReinitConfigRequest, opts ...grpc.CallOption) (*ReinitConfigResponse, error)
}

type mysqlCtlClient struct {
	cc *grpc.ClientConn
}

func NewMysqlCtlClient(cc *grpc.ClientConn) MysqlCtlClient {
	return &mysqlCtlClient{cc}
}

func (c *mysqlCtlClient) Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error) {
	out := new(StartResponse)
	err := grpc.Invoke(ctx, "/mysqlctl.MysqlCtl/Start", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mysqlCtlClient) Shutdown(ctx context.Context, in *ShutdownRequest, opts ...grpc.CallOption) (*ShutdownResponse, error) {
	out := new(ShutdownResponse)
	err := grpc.Invoke(ctx, "/mysqlctl.MysqlCtl/Shutdown", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mysqlCtlClient) RunMysqlUpgrade(ctx context.Context, in *RunMysqlUpgradeRequest, opts ...grpc.CallOption) (*RunMysqlUpgradeResponse, error) {
	out := new(RunMysqlUpgradeResponse)
	err := grpc.Invoke(ctx, "/mysqlctl.MysqlCtl/RunMysqlUpgrade", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mysqlCtlClient) ReinitConfig(ctx context.Context, in *ReinitConfigRequest, opts ...grpc.CallOption) (*ReinitConfigResponse, error) {
	out := new(ReinitConfigResponse)
	err := grpc.Invoke(ctx, "/mysqlctl.MysqlCtl/ReinitConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MysqlCtl service

type MysqlCtlServer interface {
	Start(context.Context, *StartRequest) (*StartResponse, error)
	Shutdown(context.Context, *ShutdownRequest) (*ShutdownResponse, error)
	RunMysqlUpgrade(context.Context, *RunMysqlUpgradeRequest) (*RunMysqlUpgradeResponse, error)
	ReinitConfig(context.Context, *ReinitConfigRequest) (*ReinitConfigResponse, error)
}

func RegisterMysqlCtlServer(s *grpc.Server, srv MysqlCtlServer) {
	s.RegisterService(&_MysqlCtl_serviceDesc, srv)
}

func _MysqlCtl_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MysqlCtlServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mysqlctl.MysqlCtl/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MysqlCtlServer).Start(ctx, req.(*StartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MysqlCtl_Shutdown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShutdownRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MysqlCtlServer).Shutdown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mysqlctl.MysqlCtl/Shutdown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MysqlCtlServer).Shutdown(ctx, req.(*ShutdownRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MysqlCtl_RunMysqlUpgrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunMysqlUpgradeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MysqlCtlServer).RunMysqlUpgrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mysqlctl.MysqlCtl/RunMysqlUpgrade",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MysqlCtlServer).RunMysqlUpgrade(ctx, req.(*RunMysqlUpgradeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MysqlCtl_ReinitConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReinitConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MysqlCtlServer).ReinitConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mysqlctl.MysqlCtl/ReinitConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MysqlCtlServer).ReinitConfig(ctx, req.(*ReinitConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MysqlCtl_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mysqlctl.MysqlCtl",
	HandlerType: (*MysqlCtlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _MysqlCtl_Start_Handler,
		},
		{
			MethodName: "Shutdown",
			Handler:    _MysqlCtl_Shutdown_Handler,
		},
		{
			MethodName: "RunMysqlUpgrade",
			Handler:    _MysqlCtl_RunMysqlUpgrade_Handler,
		},
		{
			MethodName: "ReinitConfig",
			Handler:    _MysqlCtl_ReinitConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("mysqlctl.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 265 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x52, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0x76, 0x82, 0x52, 0x1e, 0xdb, 0x2a, 0x4f, 0xed, 0xb6, 0x80, 0x32, 0x7b, 0x10, 0x4f, 0x3b,
	0xe8, 0x49, 0xaf, 0x05, 0x6f, 0x22, 0x74, 0x08, 0xde, 0x46, 0xb5, 0xd9, 0x2c, 0xd4, 0xa4, 0x4b,
	0x5e, 0x19, 0xfe, 0x63, 0xfe, 0x7d, 0x62, 0x4c, 0xba, 0xcc, 0x56, 0x8f, 0xf9, 0x7e, 0x3d, 0xbe,
	0x8f, 0xc0, 0xf0, 0xfd, 0x43, 0xaf, 0xcb, 0x57, 0x2a, 0x67, 0x95, 0x92, 0x24, 0x31, 0x70, 0xef,
	0x78, 0x08, 0xfd, 0x39, 0x65, 0x8a, 0x52, 0xbe, 0xae, 0xb9, 0xa6, 0x38, 0x84, 0x81, 0x7d, 0xeb,
	0x4a, 0x0a, 0xcd, 0xe3, 0x5b, 0x08, 0xe7, 0x6f, 0x35, 0xe5, 0x72, 0x23, 0xac, 0x06, 0x2f, 0x21,
	0xdc, 0x64, 0x05, 0x2d, 0x96, 0x52, 0x2d, 0x4c, 0x50, 0x3e, 0xee, 0x4d, 0x7b, 0x57, 0x41, 0x3a,
	0xf8, 0x86, 0xef, 0xa5, 0x7a, 0x30, 0x60, 0x8c, 0x70, 0xb4, 0xb5, 0xda, 0xb8, 0x31, 0x44, 0x69,
	0x2d, 0x8c, 0xe0, 0xa9, 0x5a, 0xa9, 0x2c, 0xe7, 0xee, 0xf2, 0x04, 0x46, 0x2d, 0xc6, 0x9a, 0x4e,
	0xe1, 0x38, 0xe5, 0x85, 0x28, 0x28, 0x91, 0x62, 0x59, 0xac, 0x9c, 0x23, 0x82, 0x93, 0x5d, 0xf8,
	0x47, 0x7e, 0xfd, 0xb9, 0x0f, 0x81, 0xc9, 0x49, 0xa8, 0xc4, 0x3b, 0x38, 0x30, 0x85, 0x30, 0x9a,
	0x35, 0x23, 0xf8, 0x8d, 0xd9, 0xa8, 0x85, 0xdb, 0xab, 0x7b, 0x98, 0x40, 0xe0, 0x0a, 0xe0, 0xc4,
	0x93, 0xed, 0xee, 0xc1, 0x58, 0x17, 0xd5, 0x84, 0x3c, 0x43, 0xf8, 0xab, 0x17, 0x4e, 0xb7, 0x86,
	0xee, 0x31, 0xd8, 0xc5, 0x3f, 0x8a, 0x26, 0xf9, 0x11, 0xfa, 0x7e, 0x7f, 0x3c, 0xf3, 0x4c, 0xed,
	0xb9, 0xd8, 0xf9, 0x5f, 0xb4, 0x0b, 0x7c, 0x39, 0x34, 0xbf, 0xe3, 0xe6, 0x2b, 0x00, 0x00, 0xff,
	0xff, 0x1f, 0x63, 0x65, 0xd4, 0x2f, 0x02, 0x00, 0x00,
}
