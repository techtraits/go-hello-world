// Code generated by protoc-gen-go.
// source: proto/helloworld.proto
// DO NOT EDIT!

/*
Package hello is a generated protocol buffer package.

It is generated from these files:
	proto/helloworld.proto

It has these top-level messages:
	User
	Greeting
*/
package hello

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

type User struct {
	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=last_name,json=lastName" json:"last_name,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Greeting struct {
	Text string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
}

func (m *Greeting) Reset()                    { *m = Greeting{} }
func (m *Greeting) String() string            { return proto.CompactTextString(m) }
func (*Greeting) ProtoMessage()               {}
func (*Greeting) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*User)(nil), "hello.User")
	proto.RegisterType((*Greeting)(nil), "hello.Greeting")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Greeter service

type GreeterClient interface {
	SayHello(ctx context.Context, in *User, opts ...grpc.CallOption) (*Greeting, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *User, opts ...grpc.CallOption) (*Greeting, error) {
	out := new(Greeting)
	err := grpc.Invoke(ctx, "/hello.Greeter/SayHello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Greeter service

type GreeterServer interface {
	SayHello(context.Context, *User) (*Greeting, error)
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hello.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("proto/helloworld.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0x48, 0xcd, 0xc9, 0xc9, 0x2f, 0xcf, 0x2f, 0xca, 0x49, 0xd1, 0x03, 0x0b, 0x08,
	0xb1, 0x82, 0x45, 0x94, 0x9c, 0xb8, 0x58, 0x42, 0x8b, 0x53, 0x8b, 0x84, 0x64, 0xb9, 0xb8, 0xd2,
	0x32, 0x8b, 0x8a, 0x4b, 0xe2, 0xf3, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83,
	0x38, 0xc1, 0x22, 0x7e, 0x40, 0x01, 0x21, 0x69, 0x2e, 0xce, 0x9c, 0x44, 0x98, 0x2c, 0x13, 0x58,
	0x96, 0x03, 0x24, 0x00, 0x92, 0x54, 0x92, 0xe3, 0xe2, 0x70, 0x2f, 0x4a, 0x4d, 0x2d, 0xc9, 0xcc,
	0x4b, 0x17, 0x12, 0xe2, 0x62, 0x29, 0x49, 0xad, 0x28, 0x81, 0x9a, 0x00, 0x66, 0x1b, 0x99, 0x72,
	0xb1, 0x83, 0xe5, 0x81, 0xd6, 0x68, 0x71, 0x71, 0x04, 0x27, 0x56, 0x7a, 0x80, 0xac, 0x16, 0xe2,
	0xd6, 0x03, 0x3b, 0x41, 0x0f, 0x64, 0xbf, 0x14, 0x3f, 0x94, 0x03, 0x33, 0x48, 0x89, 0x21, 0x89,
	0x0d, 0xec, 0x50, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb1, 0x01, 0x05, 0x90, 0xc2, 0x00,
	0x00, 0x00,
}