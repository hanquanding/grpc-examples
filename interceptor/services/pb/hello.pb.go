// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello.proto

package services

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloReply struct {
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReply) Reset()         { *m = HelloReply{} }
func (m *HelloReply) String() string { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()    {}
func (*HelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{1}
}

func (m *HelloReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReply.Unmarshal(m, b)
}
func (m *HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReply.Marshal(b, m, deterministic)
}
func (m *HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReply.Merge(m, src)
}
func (m *HelloReply) XXX_Size() int {
	return xxx_messageInfo_HelloReply.Size(m)
}
func (m *HelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReply proto.InternalMessageInfo

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "services.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "services.HelloReply")
}

func init() {
	proto.RegisterFile("hello.proto", fileDescriptor_61ef911816e0a8ce)
}

var fileDescriptor_61ef911816e0a8ce = []byte{
	// 159 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x48, 0xcd, 0xc9,
	0xc9, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x28, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e,
	0x2d, 0x56, 0x52, 0xe2, 0xe2, 0xf1, 0x00, 0x49, 0x04, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08,
	0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9,
	0x4a, 0x6a, 0x5c, 0x5c, 0x50, 0x35, 0x05, 0x39, 0x95, 0x42, 0x12, 0x5c, 0xec, 0xb9, 0xa9, 0xc5,
	0xc5, 0x89, 0xe9, 0xa9, 0x12, 0x4c, 0x60, 0x45, 0x30, 0xae, 0x51, 0x3b, 0x23, 0xd4, 0xb0, 0x60,
	0x88, 0xe9, 0x42, 0x16, 0x5c, 0x1c, 0xc1, 0x89, 0x95, 0x60, 0x21, 0x21, 0x31, 0x3d, 0x98, 0x9d,
	0x7a, 0xc8, 0x16, 0x4a, 0x89, 0x60, 0x88, 0x83, 0x2c, 0xb1, 0xe5, 0xe2, 0x0c, 0x2e, 0x29, 0x4a,
	0x4d, 0xcc, 0xcd, 0xcc, 0x4b, 0x27, 0x4d, 0xab, 0x06, 0xa3, 0x01, 0x63, 0x12, 0x1b, 0xd8, 0x9b,
	0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x57, 0x0b, 0xe5, 0xf5, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HelloServiceClient is the client API for HelloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloServiceClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	Streaming(ctx context.Context, opts ...grpc.CallOption) (HelloService_StreamingClient, error)
}

type helloServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloServiceClient(cc grpc.ClientConnInterface) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/services.HelloService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloServiceClient) Streaming(ctx context.Context, opts ...grpc.CallOption) (HelloService_StreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &_HelloService_serviceDesc.Streams[0], "/services.HelloService/Streaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloServiceStreamingClient{stream}
	return x, nil
}

type HelloService_StreamingClient interface {
	Send(*HelloRequest) error
	Recv() (*HelloReply, error)
	grpc.ClientStream
}

type helloServiceStreamingClient struct {
	grpc.ClientStream
}

func (x *helloServiceStreamingClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *helloServiceStreamingClient) Recv() (*HelloReply, error) {
	m := new(HelloReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HelloServiceServer is the server API for HelloService service.
type HelloServiceServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	Streaming(HelloService_StreamingServer) error
}

// UnimplementedHelloServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHelloServiceServer struct {
}

func (*UnimplementedHelloServiceServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (*UnimplementedHelloServiceServer) Streaming(srv HelloService_StreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method Streaming not implemented")
}

func RegisterHelloServiceServer(s *grpc.Server, srv HelloServiceServer) {
	s.RegisterService(&_HelloService_serviceDesc, srv)
}

func _HelloService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.HelloService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloService_Streaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloServiceServer).Streaming(&helloServiceStreamingServer{stream})
}

type HelloService_StreamingServer interface {
	Send(*HelloReply) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type helloServiceStreamingServer struct {
	grpc.ServerStream
}

func (x *helloServiceStreamingServer) Send(m *HelloReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *helloServiceStreamingServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _HelloService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "services.HelloService",
	HandlerType: (*HelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _HelloService_SayHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Streaming",
			Handler:       _HelloService_Streaming_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "hello.proto",
}