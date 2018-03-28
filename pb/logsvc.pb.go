// Code generated by protoc-gen-go. DO NOT EDIT.
// source: logsvc.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	logsvc.proto

It has these top-level messages:
	Tags
	Log
	WriteResponse
	GetRequest
	GetResponse
*/
package pb

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

type Tags struct {
	Tags map[string]string `protobuf:"bytes,1,rep,name=tags" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Tags) Reset()                    { *m = Tags{} }
func (m *Tags) String() string            { return proto.CompactTextString(m) }
func (*Tags) ProtoMessage()               {}
func (*Tags) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Tags) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type Log struct {
	ClientIp string `protobuf:"bytes,1,opt,name=client_ip,json=clientIp" json:"client_ip,omitempty"`
	ServerIp string `protobuf:"bytes,2,opt,name=server_ip,json=serverIp" json:"server_ip,omitempty"`
	Tags     *Tags  `protobuf:"bytes,3,opt,name=tags" json:"tags,omitempty"`
	Msg      string `protobuf:"bytes,4,opt,name=msg" json:"msg,omitempty"`
}

func (m *Log) Reset()                    { *m = Log{} }
func (m *Log) String() string            { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()               {}
func (*Log) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Log) GetClientIp() string {
	if m != nil {
		return m.ClientIp
	}
	return ""
}

func (m *Log) GetServerIp() string {
	if m != nil {
		return m.ServerIp
	}
	return ""
}

func (m *Log) GetTags() *Tags {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Log) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type WriteResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
}

func (m *WriteResponse) Reset()                    { *m = WriteResponse{} }
func (m *WriteResponse) String() string            { return proto.CompactTextString(m) }
func (*WriteResponse) ProtoMessage()               {}
func (*WriteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *WriteResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type GetRequest struct {
	ClientIp string `protobuf:"bytes,1,opt,name=client_ip,json=clientIp" json:"client_ip,omitempty"`
	ServerIp string `protobuf:"bytes,2,opt,name=server_ip,json=serverIp" json:"server_ip,omitempty"`
	Tags     *Tags  `protobuf:"bytes,3,opt,name=tags" json:"tags,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetRequest) GetClientIp() string {
	if m != nil {
		return m.ClientIp
	}
	return ""
}

func (m *GetRequest) GetServerIp() string {
	if m != nil {
		return m.ServerIp
	}
	return ""
}

func (m *GetRequest) GetTags() *Tags {
	if m != nil {
		return m.Tags
	}
	return nil
}

type GetResponse struct {
	Logs []*Log `protobuf:"bytes,1,rep,name=logs" json:"logs,omitempty"`
}

func (m *GetResponse) Reset()                    { *m = GetResponse{} }
func (m *GetResponse) String() string            { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()               {}
func (*GetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetResponse) GetLogs() []*Log {
	if m != nil {
		return m.Logs
	}
	return nil
}

func init() {
	proto.RegisterType((*Tags)(nil), "pb.Tags")
	proto.RegisterType((*Log)(nil), "pb.Log")
	proto.RegisterType((*WriteResponse)(nil), "pb.WriteResponse")
	proto.RegisterType((*GetRequest)(nil), "pb.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "pb.GetResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for LogService service

type LogServiceClient interface {
	Write(ctx context.Context, in *Log, opts ...grpc.CallOption) (*WriteResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type logServiceClient struct {
	cc *grpc.ClientConn
}

func NewLogServiceClient(cc *grpc.ClientConn) LogServiceClient {
	return &logServiceClient{cc}
}

func (c *logServiceClient) Write(ctx context.Context, in *Log, opts ...grpc.CallOption) (*WriteResponse, error) {
	out := new(WriteResponse)
	err := grpc.Invoke(ctx, "/pb.LogService/Write", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := grpc.Invoke(ctx, "/pb.LogService/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LogService service

type LogServiceServer interface {
	Write(context.Context, *Log) (*WriteResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
}

func RegisterLogServiceServer(s *grpc.Server, srv LogServiceServer) {
	s.RegisterService(&_LogService_serviceDesc, srv)
}

func _LogService_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Log)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServiceServer).Write(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LogService/Write",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServiceServer).Write(ctx, req.(*Log))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LogService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LogService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.LogService",
	HandlerType: (*LogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Write",
			Handler:    _LogService_Write_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _LogService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "logsvc.proto",
}

func init() { proto.RegisterFile("logsvc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0x4f, 0x4b, 0x33, 0x31,
	0x10, 0xc6, 0xdf, 0xed, 0x6e, 0xdf, 0x6e, 0xa7, 0xfe, 0x1d, 0x3c, 0x2c, 0xad, 0x87, 0xb2, 0xa0,
	0xac, 0x1e, 0x7a, 0xa8, 0x07, 0xc5, 0xbb, 0x94, 0x42, 0x4f, 0x51, 0xf0, 0x24, 0xd2, 0xae, 0x43,
	0x58, 0x6c, 0x9b, 0x98, 0x49, 0x17, 0xfa, 0xed, 0x25, 0xc9, 0x6e, 0xc5, 0x0f, 0xe0, 0x25, 0xcc,
	0x3c, 0xcf, 0x24, 0xf3, 0x9b, 0x24, 0x70, 0xb4, 0x56, 0x92, 0xeb, 0x72, 0xa2, 0x8d, 0xb2, 0x0a,
	0x3b, 0x7a, 0x95, 0x4b, 0x48, 0x5e, 0x96, 0x92, 0xf1, 0x1a, 0x12, 0xbb, 0x94, 0x9c, 0x45, 0xe3,
	0xb8, 0x18, 0x4c, 0x71, 0xa2, 0x57, 0x13, 0xa7, 0xfb, 0xe5, 0x69, 0x6b, 0xcd, 0x5e, 0x78, 0x7f,
	0x78, 0x0f, 0xfd, 0x83, 0x84, 0x67, 0x10, 0x7f, 0xd2, 0x3e, 0x8b, 0xc6, 0x51, 0xd1, 0x17, 0x2e,
	0xc4, 0x0b, 0xe8, 0xd6, 0xcb, 0xf5, 0x8e, 0xb2, 0x8e, 0xd7, 0x42, 0xf2, 0xd8, 0x79, 0x88, 0x72,
	0x05, 0xf1, 0x42, 0x49, 0x1c, 0x41, 0xbf, 0x5c, 0x57, 0xb4, 0xb5, 0xef, 0x95, 0x6e, 0x36, 0xa6,
	0x41, 0x98, 0x6b, 0x67, 0x32, 0x99, 0x9a, 0x8c, 0x33, 0xc3, 0x09, 0x69, 0x10, 0xe6, 0x1a, 0x2f,
	0x1b, 0xc2, 0x78, 0x1c, 0x15, 0x83, 0x69, 0xda, 0x12, 0x06, 0x2e, 0x87, 0xb2, 0x61, 0x99, 0x25,
	0x01, 0x65, 0xc3, 0x32, 0xbf, 0x81, 0xe3, 0x57, 0x53, 0x59, 0x12, 0xc4, 0x5a, 0x6d, 0x99, 0x30,
	0x83, 0x1e, 0xef, 0xca, 0x92, 0x98, 0x7d, 0xe3, 0x54, 0xb4, 0x69, 0xfe, 0x01, 0x30, 0x23, 0x2b,
	0xe8, 0x6b, 0x47, 0x6c, 0xff, 0x0a, 0x31, 0xbf, 0x85, 0x81, 0xef, 0xd2, 0xe0, 0x8c, 0x20, 0x71,
	0xaf, 0xd1, 0xdc, 0x78, 0xcf, 0x15, 0x2f, 0x94, 0x14, 0x5e, 0x9c, 0xbe, 0x01, 0x2c, 0x94, 0x7c,
	0x26, 0x53, 0x57, 0x25, 0xe1, 0x15, 0x74, 0xfd, 0x28, 0xd8, 0x56, 0x0d, 0xcf, 0x5d, 0xf0, 0x6b,
	0xbc, 0xfc, 0x1f, 0x16, 0x10, 0xcf, 0xc8, 0xe2, 0x89, 0xf3, 0x7e, 0xe6, 0x19, 0x9e, 0x1e, 0xf2,
	0xb6, 0x72, 0xf5, 0xdf, 0x7f, 0x80, 0xbb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x18, 0x97, 0xc7,
	0x42, 0x10, 0x02, 0x00, 0x00,
}