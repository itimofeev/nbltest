// Code generated by protoc-gen-go. DO NOT EDIT.
// source: axxonsoft/bl/events/Notification.Internal.proto

package events // import "bitbucket.org/Axxonsoft/bl/events"

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

type PingRequest struct {
	TimeoutMs            int32    `protobuf:"varint,1,opt,name=timeoutMs" json:"timeoutMs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_Notification_Internal_e1da3af6fb7c2dbc, []int{0}
}
func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRequest.Unmarshal(m, b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
}
func (dst *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(dst, src)
}
func (m *PingRequest) XXX_Size() int {
	return xxx_messageInfo_PingRequest.Size(m)
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

func (m *PingRequest) GetTimeoutMs() int32 {
	if m != nil {
		return m.TimeoutMs
	}
	return 0
}

type PingResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingResponse) Reset()         { *m = PingResponse{} }
func (m *PingResponse) String() string { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()    {}
func (*PingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_Notification_Internal_e1da3af6fb7c2dbc, []int{1}
}
func (m *PingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingResponse.Unmarshal(m, b)
}
func (m *PingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingResponse.Marshal(b, m, deterministic)
}
func (dst *PingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingResponse.Merge(dst, src)
}
func (m *PingResponse) XXX_Size() int {
	return xxx_messageInfo_PingResponse.Size(m)
}
func (m *PingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PingResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PingRequest)(nil), "axxonsoft.bl.events.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "axxonsoft.bl.events.PingResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for NodeNotifier service

type NodeNotifierClient interface {
	PullEvents(ctx context.Context, in *PullEventsRequest, opts ...grpc.CallOption) (NodeNotifier_PullEventsClient, error)
	UpdateSubscription(ctx context.Context, in *UpdateSubscriptionRequest, opts ...grpc.CallOption) (*UpdateSubscriptionResponse, error)
	DisconnectEventChannel(ctx context.Context, in *DisconnectEventChannelRequest, opts ...grpc.CallOption) (*DisconnectEventChannelResponse, error)
	// Internal method for node status checking
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (NodeNotifier_PingClient, error)
}

type nodeNotifierClient struct {
	cc *grpc.ClientConn
}

func NewNodeNotifierClient(cc *grpc.ClientConn) NodeNotifierClient {
	return &nodeNotifierClient{cc}
}

func (c *nodeNotifierClient) PullEvents(ctx context.Context, in *PullEventsRequest, opts ...grpc.CallOption) (NodeNotifier_PullEventsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_NodeNotifier_serviceDesc.Streams[0], c.cc, "/axxonsoft.bl.events.NodeNotifier/PullEvents", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeNotifierPullEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodeNotifier_PullEventsClient interface {
	Recv() (*Events, error)
	grpc.ClientStream
}

type nodeNotifierPullEventsClient struct {
	grpc.ClientStream
}

func (x *nodeNotifierPullEventsClient) Recv() (*Events, error) {
	m := new(Events)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeNotifierClient) UpdateSubscription(ctx context.Context, in *UpdateSubscriptionRequest, opts ...grpc.CallOption) (*UpdateSubscriptionResponse, error) {
	out := new(UpdateSubscriptionResponse)
	err := grpc.Invoke(ctx, "/axxonsoft.bl.events.NodeNotifier/UpdateSubscription", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeNotifierClient) DisconnectEventChannel(ctx context.Context, in *DisconnectEventChannelRequest, opts ...grpc.CallOption) (*DisconnectEventChannelResponse, error) {
	out := new(DisconnectEventChannelResponse)
	err := grpc.Invoke(ctx, "/axxonsoft.bl.events.NodeNotifier/DisconnectEventChannel", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeNotifierClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (NodeNotifier_PingClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_NodeNotifier_serviceDesc.Streams[1], c.cc, "/axxonsoft.bl.events.NodeNotifier/Ping", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeNotifierPingClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodeNotifier_PingClient interface {
	Recv() (*PingResponse, error)
	grpc.ClientStream
}

type nodeNotifierPingClient struct {
	grpc.ClientStream
}

func (x *nodeNotifierPingClient) Recv() (*PingResponse, error) {
	m := new(PingResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for NodeNotifier service

type NodeNotifierServer interface {
	PullEvents(*PullEventsRequest, NodeNotifier_PullEventsServer) error
	UpdateSubscription(context.Context, *UpdateSubscriptionRequest) (*UpdateSubscriptionResponse, error)
	DisconnectEventChannel(context.Context, *DisconnectEventChannelRequest) (*DisconnectEventChannelResponse, error)
	// Internal method for node status checking
	Ping(*PingRequest, NodeNotifier_PingServer) error
}

func RegisterNodeNotifierServer(s *grpc.Server, srv NodeNotifierServer) {
	s.RegisterService(&_NodeNotifier_serviceDesc, srv)
}

func _NodeNotifier_PullEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PullEventsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeNotifierServer).PullEvents(m, &nodeNotifierPullEventsServer{stream})
}

type NodeNotifier_PullEventsServer interface {
	Send(*Events) error
	grpc.ServerStream
}

type nodeNotifierPullEventsServer struct {
	grpc.ServerStream
}

func (x *nodeNotifierPullEventsServer) Send(m *Events) error {
	return x.ServerStream.SendMsg(m)
}

func _NodeNotifier_UpdateSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSubscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeNotifierServer).UpdateSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axxonsoft.bl.events.NodeNotifier/UpdateSubscription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeNotifierServer).UpdateSubscription(ctx, req.(*UpdateSubscriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeNotifier_DisconnectEventChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisconnectEventChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeNotifierServer).DisconnectEventChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axxonsoft.bl.events.NodeNotifier/DisconnectEventChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeNotifierServer).DisconnectEventChannel(ctx, req.(*DisconnectEventChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeNotifier_Ping_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PingRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeNotifierServer).Ping(m, &nodeNotifierPingServer{stream})
}

type NodeNotifier_PingServer interface {
	Send(*PingResponse) error
	grpc.ServerStream
}

type nodeNotifierPingServer struct {
	grpc.ServerStream
}

func (x *nodeNotifierPingServer) Send(m *PingResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _NodeNotifier_serviceDesc = grpc.ServiceDesc{
	ServiceName: "axxonsoft.bl.events.NodeNotifier",
	HandlerType: (*NodeNotifierServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateSubscription",
			Handler:    _NodeNotifier_UpdateSubscription_Handler,
		},
		{
			MethodName: "DisconnectEventChannel",
			Handler:    _NodeNotifier_DisconnectEventChannel_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PullEvents",
			Handler:       _NodeNotifier_PullEvents_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Ping",
			Handler:       _NodeNotifier_Ping_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "axxonsoft/bl/events/Notification.Internal.proto",
}

func init() {
	proto.RegisterFile("axxonsoft/bl/events/Notification.Internal.proto", fileDescriptor_Notification_Internal_e1da3af6fb7c2dbc)
}

var fileDescriptor_Notification_Internal_e1da3af6fb7c2dbc = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4b, 0x4b, 0x03, 0x31,
	0x10, 0xc7, 0x59, 0x7c, 0x80, 0xb1, 0x78, 0x88, 0x20, 0x52, 0x3d, 0xb4, 0x15, 0x8a, 0x20, 0x64,
	0x4b, 0xfb, 0x09, 0x7c, 0x1d, 0x3c, 0xb4, 0x94, 0x16, 0x2f, 0xde, 0x36, 0xdb, 0x69, 0x0d, 0xc6,
	0x99, 0x75, 0x33, 0x91, 0x5e, 0xfd, 0x40, 0x7e, 0x47, 0xe9, 0xa6, 0x75, 0x0f, 0x0d, 0x3e, 0xae,
	0xe1, 0xf7, 0x7f, 0x64, 0x66, 0x44, 0x9a, 0x2d, 0x97, 0x84, 0x8e, 0xe6, 0x9c, 0x6a, 0x9b, 0xc2,
	0x3b, 0x20, 0xbb, 0x74, 0x44, 0x6c, 0xe6, 0x26, 0xcf, 0xd8, 0x10, 0xaa, 0x07, 0x64, 0x28, 0x31,
	0xb3, 0xaa, 0x28, 0x89, 0x49, 0x1e, 0x7f, 0x0b, 0x94, 0xb6, 0x2a, 0x08, 0x9a, 0xdd, 0x5f, 0x5d,
	0x2a, 0x71, 0xe7, 0x4a, 0x1c, 0x8e, 0x0d, 0x2e, 0x26, 0xf0, 0xe6, 0xc1, 0xb1, 0x3c, 0x17, 0x07,
	0x6c, 0x5e, 0x81, 0x3c, 0x0f, 0xdd, 0x69, 0xd2, 0x4a, 0x2e, 0xf7, 0x26, 0xf5, 0x43, 0xe7, 0x48,
	0x34, 0x02, 0xec, 0x0a, 0x42, 0x07, 0xfd, 0xcf, 0x1d, 0xd1, 0x18, 0xd1, 0x0c, 0x82, 0x2f, 0x94,
	0x72, 0x2a, 0xc4, 0xd8, 0x5b, 0x7b, 0x5f, 0xc5, 0xc9, 0xae, 0x8a, 0x34, 0x53, 0x35, 0xb0, 0x0e,
	0x6d, 0x9e, 0x45, 0xb9, 0xc0, 0xf4, 0x12, 0xe9, 0x85, 0x7c, 0x2c, 0x66, 0x19, 0xc3, 0xd4, 0x6b,
	0x97, 0x97, 0xa6, 0x58, 0xd5, 0x97, 0x2a, 0x2a, 0xda, 0x06, 0x37, 0x21, 0xe9, 0x9f, 0xf9, 0xf0,
	0x39, 0xf9, 0x91, 0x88, 0x93, 0x3b, 0xe3, 0x72, 0x42, 0x84, 0x9c, 0xab, 0x36, 0xb7, 0xcf, 0x19,
	0x22, 0x58, 0xd9, 0x8f, 0x7a, 0xc5, 0xe1, 0x4d, 0xfe, 0xe0, 0x5f, 0x9a, 0x75, 0x87, 0xa1, 0xd8,
	0x5d, 0x0d, 0x5c, 0xb6, 0xe2, 0x93, 0xac, 0x17, 0xd7, 0x6c, 0xff, 0x40, 0x04, 0xb3, 0x5e, 0x72,
	0x73, 0xf1, 0xd4, 0xd6, 0x86, 0xb5, 0xcf, 0x5f, 0x80, 0x15, 0x95, 0x8b, 0xf4, 0x7a, 0xfb, 0x48,
	0xf4, 0x7e, 0x75, 0x18, 0x83, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6e, 0x69, 0x6b, 0x8a, 0x88,
	0x02, 0x00, 0x00,
}
