// Code generated by protoc-gen-go. DO NOT EDIT.
// source: axxonsoft/bl/events/SystemMessages.Internal.proto

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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for NodeSystemMessagesService service

type NodeSystemMessagesServiceClient interface {
	GetJsonEvents(ctx context.Context, in *GetJsonEventsRequest, opts ...grpc.CallOption) (*EventsResponse, error)
	GetTextEvents(ctx context.Context, in *GetTextEventsRequest, opts ...grpc.CallOption) (*EventsResponse, error)
	GetAlerts(ctx context.Context, in *GetAlertsRequest, opts ...grpc.CallOption) (*GetAlertsResponse, error)
}

type nodeSystemMessagesServiceClient struct {
	cc *grpc.ClientConn
}

func NewNodeSystemMessagesServiceClient(cc *grpc.ClientConn) NodeSystemMessagesServiceClient {
	return &nodeSystemMessagesServiceClient{cc}
}

func (c *nodeSystemMessagesServiceClient) GetJsonEvents(ctx context.Context, in *GetJsonEventsRequest, opts ...grpc.CallOption) (*EventsResponse, error) {
	out := new(EventsResponse)
	err := grpc.Invoke(ctx, "/axxonsoft.bl.events.NodeSystemMessagesService/GetJsonEvents", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeSystemMessagesServiceClient) GetTextEvents(ctx context.Context, in *GetTextEventsRequest, opts ...grpc.CallOption) (*EventsResponse, error) {
	out := new(EventsResponse)
	err := grpc.Invoke(ctx, "/axxonsoft.bl.events.NodeSystemMessagesService/GetTextEvents", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeSystemMessagesServiceClient) GetAlerts(ctx context.Context, in *GetAlertsRequest, opts ...grpc.CallOption) (*GetAlertsResponse, error) {
	out := new(GetAlertsResponse)
	err := grpc.Invoke(ctx, "/axxonsoft.bl.events.NodeSystemMessagesService/GetAlerts", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NodeSystemMessagesService service

type NodeSystemMessagesServiceServer interface {
	GetJsonEvents(context.Context, *GetJsonEventsRequest) (*EventsResponse, error)
	GetTextEvents(context.Context, *GetTextEventsRequest) (*EventsResponse, error)
	GetAlerts(context.Context, *GetAlertsRequest) (*GetAlertsResponse, error)
}

func RegisterNodeSystemMessagesServiceServer(s *grpc.Server, srv NodeSystemMessagesServiceServer) {
	s.RegisterService(&_NodeSystemMessagesService_serviceDesc, srv)
}

func _NodeSystemMessagesService_GetJsonEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetJsonEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeSystemMessagesServiceServer).GetJsonEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axxonsoft.bl.events.NodeSystemMessagesService/GetJsonEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeSystemMessagesServiceServer).GetJsonEvents(ctx, req.(*GetJsonEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeSystemMessagesService_GetTextEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTextEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeSystemMessagesServiceServer).GetTextEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axxonsoft.bl.events.NodeSystemMessagesService/GetTextEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeSystemMessagesServiceServer).GetTextEvents(ctx, req.(*GetTextEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeSystemMessagesService_GetAlerts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAlertsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeSystemMessagesServiceServer).GetAlerts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axxonsoft.bl.events.NodeSystemMessagesService/GetAlerts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeSystemMessagesServiceServer).GetAlerts(ctx, req.(*GetAlertsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NodeSystemMessagesService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "axxonsoft.bl.events.NodeSystemMessagesService",
	HandlerType: (*NodeSystemMessagesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetJsonEvents",
			Handler:    _NodeSystemMessagesService_GetJsonEvents_Handler,
		},
		{
			MethodName: "GetTextEvents",
			Handler:    _NodeSystemMessagesService_GetTextEvents_Handler,
		},
		{
			MethodName: "GetAlerts",
			Handler:    _NodeSystemMessagesService_GetAlerts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "axxonsoft/bl/events/SystemMessages.Internal.proto",
}

func init() {
	proto.RegisterFile("axxonsoft/bl/events/SystemMessages.Internal.proto", fileDescriptor_SystemMessages_Internal_bf4f590e13080acf)
}

var fileDescriptor_SystemMessages_Internal_bf4f590e13080acf = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0xb1, 0x4a, 0xc5, 0x40,
	0x10, 0x45, 0xc1, 0x42, 0x70, 0xc1, 0x66, 0xad, 0x7c, 0x9d, 0x3c, 0x14, 0x6d, 0x76, 0x51, 0xbf,
	0xe0, 0x09, 0x12, 0x14, 0xb4, 0x30, 0x56, 0x69, 0x24, 0x1b, 0xaf, 0x21, 0xb8, 0xee, 0xc4, 0x9d,
	0x49, 0x88, 0x9f, 0xe6, 0xdf, 0x09, 0x09, 0x51, 0x82, 0x91, 0xf8, 0xea, 0x39, 0xf7, 0x1c, 0x18,
	0x75, 0x9e, 0x77, 0x1d, 0x05, 0xa6, 0x17, 0xb1, 0xce, 0x5b, 0xb4, 0x08, 0xc2, 0x36, 0xfd, 0x60,
	0xc1, 0xdb, 0x1d, 0x98, 0xf3, 0x12, 0x6c, 0x6e, 0x82, 0x20, 0x86, 0xdc, 0x9b, 0x3a, 0x92, 0x90,
	0x3e, 0xf8, 0x9e, 0x18, 0xe7, 0xcd, 0x30, 0x59, 0x9d, 0xfe, 0xc3, 0xd3, 0xcf, 0x2f, 0x3e, 0x77,
	0xd4, 0xe1, 0x3d, 0x3d, 0x63, 0x7a, 0x4c, 0x11, 0xdb, 0xaa, 0x80, 0x7e, 0x52, 0xfb, 0x09, 0xe4,
	0x96, 0x29, 0x5c, 0xf7, 0x0e, 0x7d, 0x66, 0x66, 0x72, 0x66, 0xc2, 0x3c, 0xe0, 0xbd, 0x01, 0xcb,
	0x6a, 0x3d, 0x8b, 0x8e, 0x0c, 0xd7, 0x14, 0x78, 0x0c, 0x3c, 0xa2, 0x93, 0xa5, 0xc0, 0x0f, 0xb3,
	0x55, 0x20, 0x53, 0x7b, 0x09, 0x64, 0xe3, 0x11, 0x85, 0xf5, 0xf1, 0x5f, 0xf2, 0xe1, 0x3e, 0x8a,
	0x4f, 0x96, 0xb0, 0xc1, 0x7d, 0xb5, 0xce, 0x8e, 0x5c, 0x25, 0xae, 0x29, 0x5e, 0x21, 0x86, 0x62,
	0x69, 0x37, 0xbf, 0xbf, 0xee, 0x76, 0xfb, 0x3f, 0x5f, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0xfd,
	0x73, 0x25, 0xb7, 0xdb, 0x01, 0x00, 0x00,
}
