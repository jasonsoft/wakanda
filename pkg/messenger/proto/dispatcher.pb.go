// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/messenger/proto/dispatcher.proto

package proto

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

// The request message containing the user's name.
type HandleCommandRequest struct {
	OP                   string   `protobuf:"bytes,1,opt,name=OP,proto3" json:"OP,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HandleCommandRequest) Reset()         { *m = HandleCommandRequest{} }
func (m *HandleCommandRequest) String() string { return proto.CompactTextString(m) }
func (*HandleCommandRequest) ProtoMessage()    {}
func (*HandleCommandRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dispatcher_f606299cb4e37985, []int{0}
}
func (m *HandleCommandRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HandleCommandRequest.Unmarshal(m, b)
}
func (m *HandleCommandRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HandleCommandRequest.Marshal(b, m, deterministic)
}
func (dst *HandleCommandRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HandleCommandRequest.Merge(dst, src)
}
func (m *HandleCommandRequest) XXX_Size() int {
	return xxx_messageInfo_HandleCommandRequest.Size(m)
}
func (m *HandleCommandRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HandleCommandRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HandleCommandRequest proto.InternalMessageInfo

func (m *HandleCommandRequest) GetOP() string {
	if m != nil {
		return m.OP
	}
	return ""
}

func (m *HandleCommandRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// The response message containing the greetings
type HandleCommandReply struct {
	OP                   string   `protobuf:"bytes,1,opt,name=OP,proto3" json:"OP,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HandleCommandReply) Reset()         { *m = HandleCommandReply{} }
func (m *HandleCommandReply) String() string { return proto.CompactTextString(m) }
func (*HandleCommandReply) ProtoMessage()    {}
func (*HandleCommandReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_dispatcher_f606299cb4e37985, []int{1}
}
func (m *HandleCommandReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HandleCommandReply.Unmarshal(m, b)
}
func (m *HandleCommandReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HandleCommandReply.Marshal(b, m, deterministic)
}
func (dst *HandleCommandReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HandleCommandReply.Merge(dst, src)
}
func (m *HandleCommandReply) XXX_Size() int {
	return xxx_messageInfo_HandleCommandReply.Size(m)
}
func (m *HandleCommandReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HandleCommandReply.DiscardUnknown(m)
}

var xxx_messageInfo_HandleCommandReply proto.InternalMessageInfo

func (m *HandleCommandReply) GetOP() string {
	if m != nil {
		return m.OP
	}
	return ""
}

func (m *HandleCommandReply) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*HandleCommandRequest)(nil), "proto.HandleCommandRequest")
	proto.RegisterType((*HandleCommandReply)(nil), "proto.HandleCommandReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DispatcherClient is the client API for Dispatcher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DispatcherClient interface {
	// handle a command
	HandleCommand(ctx context.Context, in *HandleCommandRequest, opts ...grpc.CallOption) (*HandleCommandReply, error)
}

type dispatcherClient struct {
	cc *grpc.ClientConn
}

func NewDispatcherClient(cc *grpc.ClientConn) DispatcherClient {
	return &dispatcherClient{cc}
}

func (c *dispatcherClient) HandleCommand(ctx context.Context, in *HandleCommandRequest, opts ...grpc.CallOption) (*HandleCommandReply, error) {
	out := new(HandleCommandReply)
	err := c.cc.Invoke(ctx, "/proto.Dispatcher/HandleCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DispatcherServer is the server API for Dispatcher service.
type DispatcherServer interface {
	// handle a command
	HandleCommand(context.Context, *HandleCommandRequest) (*HandleCommandReply, error)
}

func RegisterDispatcherServer(s *grpc.Server, srv DispatcherServer) {
	s.RegisterService(&_Dispatcher_serviceDesc, srv)
}

func _Dispatcher_HandleCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandleCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherServer).HandleCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Dispatcher/HandleCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherServer).HandleCommand(ctx, req.(*HandleCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Dispatcher_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Dispatcher",
	HandlerType: (*DispatcherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandleCommand",
			Handler:    _Dispatcher_HandleCommand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/messenger/proto/dispatcher.proto",
}

func init() {
	proto.RegisterFile("pkg/messenger/proto/dispatcher.proto", fileDescriptor_dispatcher_f606299cb4e37985)
}

var fileDescriptor_dispatcher_f606299cb4e37985 = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x29, 0xc8, 0x4e, 0xd7,
	0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0xcd, 0x4b, 0x4f, 0x2d, 0xd2, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7,
	0x4f, 0xc9, 0x2c, 0x2e, 0x48, 0x2c, 0x49, 0xce, 0x48, 0x2d, 0xd2, 0x03, 0x0b, 0x08, 0xb1, 0x82,
	0x29, 0x25, 0x2b, 0x2e, 0x11, 0x8f, 0xc4, 0xbc, 0x94, 0x9c, 0x54, 0xe7, 0xfc, 0xdc, 0xdc, 0xc4,
	0xbc, 0x94, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x3e, 0x2e, 0x26, 0xff, 0x00, 0x09,
	0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x26, 0xff, 0x00, 0x21, 0x21, 0x2e, 0x16, 0x97, 0xc4, 0x92,
	0x44, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x9e, 0x20, 0x30, 0x5b, 0xc9, 0x82, 0x4b, 0x08, 0x4d, 0x6f,
	0x41, 0x4e, 0x25, 0x31, 0x3a, 0x8d, 0xc2, 0xb9, 0xb8, 0x5c, 0xe0, 0x0e, 0x12, 0xf2, 0xe4, 0xe2,
	0x45, 0x31, 0x47, 0x48, 0x1a, 0xe2, 0x46, 0x3d, 0x6c, 0x2e, 0x93, 0x92, 0xc4, 0x2e, 0x59, 0x90,
	0x53, 0xa9, 0xc4, 0x90, 0xc4, 0x06, 0x96, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x96, 0x73,
	0x4f, 0x64, 0x04, 0x01, 0x00, 0x00,
}
