// Code generated by protoc-gen-go. DO NOT EDIT.
// source: axxonsoft/bl/auth/Authentication.proto

package auth // import "bitbucket.org/Axxonsoft/bl/auth"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google/api"

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

type AuthenticateResponseEx_EAuthenticateCode int32

const (
	AuthenticateResponseEx_AUTHENTICATE_CODE_OK                       AuthenticateResponseEx_EAuthenticateCode = 0
	AuthenticateResponseEx_AUTHENTICATE_CODE_GENERAL_ERROR            AuthenticateResponseEx_EAuthenticateCode = 1
	AuthenticateResponseEx_AUTHENTICATE_CODE_WRONG_CREDENTIALS        AuthenticateResponseEx_EAuthenticateCode = 2
	AuthenticateResponseEx_AUTHENTICATE_CODE_PASSWORD_EXPIRED         AuthenticateResponseEx_EAuthenticateCode = 3
	AuthenticateResponseEx_AUTHENTICATE_CODE_USER_LOCKED              AuthenticateResponseEx_EAuthenticateCode = 4
	AuthenticateResponseEx_AUTHENTICATE_CODE_IP_BLOCKED               AuthenticateResponseEx_EAuthenticateCode = 5
	AuthenticateResponseEx_AUTHENTICATE_CODE_ALREADY_LOGGED           AuthenticateResponseEx_EAuthenticateCode = 6
	AuthenticateResponseEx_AUTHENTICATE_CODE_FORCE_CHANGE_CREDENTIALS AuthenticateResponseEx_EAuthenticateCode = 7
	AuthenticateResponseEx_AUTHENTICATE_CODE_PASSWORD_INVALID         AuthenticateResponseEx_EAuthenticateCode = 8
	AuthenticateResponseEx_AUTHENTICATE_CODE_PASSWORD_EXPIRE_SOON     AuthenticateResponseEx_EAuthenticateCode = 9
	AuthenticateResponseEx_AUTHENTICATE_CODE_WRONG_SUPERVISOR_ROLE    AuthenticateResponseEx_EAuthenticateCode = 10
	AuthenticateResponseEx_AUTHENTICATE_CODE_TIMEZONE_DENIED          AuthenticateResponseEx_EAuthenticateCode = 11
)

var AuthenticateResponseEx_EAuthenticateCode_name = map[int32]string{
	0:  "AUTHENTICATE_CODE_OK",
	1:  "AUTHENTICATE_CODE_GENERAL_ERROR",
	2:  "AUTHENTICATE_CODE_WRONG_CREDENTIALS",
	3:  "AUTHENTICATE_CODE_PASSWORD_EXPIRED",
	4:  "AUTHENTICATE_CODE_USER_LOCKED",
	5:  "AUTHENTICATE_CODE_IP_BLOCKED",
	6:  "AUTHENTICATE_CODE_ALREADY_LOGGED",
	7:  "AUTHENTICATE_CODE_FORCE_CHANGE_CREDENTIALS",
	8:  "AUTHENTICATE_CODE_PASSWORD_INVALID",
	9:  "AUTHENTICATE_CODE_PASSWORD_EXPIRE_SOON",
	10: "AUTHENTICATE_CODE_WRONG_SUPERVISOR_ROLE",
	11: "AUTHENTICATE_CODE_TIMEZONE_DENIED",
}
var AuthenticateResponseEx_EAuthenticateCode_value = map[string]int32{
	"AUTHENTICATE_CODE_OK":                       0,
	"AUTHENTICATE_CODE_GENERAL_ERROR":            1,
	"AUTHENTICATE_CODE_WRONG_CREDENTIALS":        2,
	"AUTHENTICATE_CODE_PASSWORD_EXPIRED":         3,
	"AUTHENTICATE_CODE_USER_LOCKED":              4,
	"AUTHENTICATE_CODE_IP_BLOCKED":               5,
	"AUTHENTICATE_CODE_ALREADY_LOGGED":           6,
	"AUTHENTICATE_CODE_FORCE_CHANGE_CREDENTIALS": 7,
	"AUTHENTICATE_CODE_PASSWORD_INVALID":         8,
	"AUTHENTICATE_CODE_PASSWORD_EXPIRE_SOON":     9,
	"AUTHENTICATE_CODE_WRONG_SUPERVISOR_ROLE":    10,
	"AUTHENTICATE_CODE_TIMEZONE_DENIED":          11,
}

func (x AuthenticateResponseEx_EAuthenticateCode) String() string {
	return proto.EnumName(AuthenticateResponseEx_EAuthenticateCode_name, int32(x))
}
func (AuthenticateResponseEx_EAuthenticateCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_Authentication_d9f3199dfb1c8a85, []int{2, 0}
}

type AuthenticateRequest struct {
	UserName             string   `protobuf:"bytes,1,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthenticateRequest) Reset()         { *m = AuthenticateRequest{} }
func (m *AuthenticateRequest) String() string { return proto.CompactTextString(m) }
func (*AuthenticateRequest) ProtoMessage()    {}
func (*AuthenticateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_Authentication_d9f3199dfb1c8a85, []int{0}
}
func (m *AuthenticateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticateRequest.Unmarshal(m, b)
}
func (m *AuthenticateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticateRequest.Marshal(b, m, deterministic)
}
func (dst *AuthenticateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticateRequest.Merge(dst, src)
}
func (m *AuthenticateRequest) XXX_Size() int {
	return xxx_messageInfo_AuthenticateRequest.Size(m)
}
func (m *AuthenticateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticateRequest proto.InternalMessageInfo

func (m *AuthenticateRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *AuthenticateRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type AuthenticateResponse struct {
	// token name to specify as request metadata authentication
	TokenName string `protobuf:"bytes,1,opt,name=token_name,json=tokenName" json:"token_name,omitempty"`
	// token value for request metadata auth
	TokenValue string `protobuf:"bytes,2,opt,name=token_value,json=tokenValue" json:"token_value,omitempty"`
	// ISO 8601 string with date toe
	// Token is good due this date.
	// Client have to obtain fresh token before token expiration.
	ExpiresAt string `protobuf:"bytes,3,opt,name=expires_at,json=expiresAt" json:"expires_at,omitempty"`
	// Indicates whether current user relates to admins or not
	IsUnrestricted bool `protobuf:"varint,10,opt,name=is_unrestricted,json=isUnrestricted" json:"is_unrestricted,omitempty"`
	// Current user id
	UserId string `protobuf:"bytes,16,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	// Active roles IDs
	RolesIds             []string `protobuf:"bytes,17,rep,name=roles_ids,json=rolesIds" json:"roles_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthenticateResponse) Reset()         { *m = AuthenticateResponse{} }
func (m *AuthenticateResponse) String() string { return proto.CompactTextString(m) }
func (*AuthenticateResponse) ProtoMessage()    {}
func (*AuthenticateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_Authentication_d9f3199dfb1c8a85, []int{1}
}
func (m *AuthenticateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticateResponse.Unmarshal(m, b)
}
func (m *AuthenticateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticateResponse.Marshal(b, m, deterministic)
}
func (dst *AuthenticateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticateResponse.Merge(dst, src)
}
func (m *AuthenticateResponse) XXX_Size() int {
	return xxx_messageInfo_AuthenticateResponse.Size(m)
}
func (m *AuthenticateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticateResponse proto.InternalMessageInfo

func (m *AuthenticateResponse) GetTokenName() string {
	if m != nil {
		return m.TokenName
	}
	return ""
}

func (m *AuthenticateResponse) GetTokenValue() string {
	if m != nil {
		return m.TokenValue
	}
	return ""
}

func (m *AuthenticateResponse) GetExpiresAt() string {
	if m != nil {
		return m.ExpiresAt
	}
	return ""
}

func (m *AuthenticateResponse) GetIsUnrestricted() bool {
	if m != nil {
		return m.IsUnrestricted
	}
	return false
}

func (m *AuthenticateResponse) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *AuthenticateResponse) GetRolesIds() []string {
	if m != nil {
		return m.RolesIds
	}
	return nil
}

type AuthenticateResponseEx struct {
	// token name to specify as request metadata authentication
	TokenName string `protobuf:"bytes,1,opt,name=token_name,json=tokenName" json:"token_name,omitempty"`
	// token value for request metadata auth
	TokenValue string `protobuf:"bytes,2,opt,name=token_value,json=tokenValue" json:"token_value,omitempty"`
	// ISO 8601 string with date toe
	// Token is good due this date.
	// Client have to obtain fresh token before token expiration.
	ExpiresAt string `protobuf:"bytes,3,opt,name=expires_at,json=expiresAt" json:"expires_at,omitempty"`
	// Indicates whether current user relates to admins or not
	IsUnrestricted bool `protobuf:"varint,10,opt,name=is_unrestricted,json=isUnrestricted" json:"is_unrestricted,omitempty"`
	// Current user id
	UserId string `protobuf:"bytes,16,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	// Active roles IDs
	RolesIds             []string                                 `protobuf:"bytes,17,rep,name=roles_ids,json=rolesIds" json:"roles_ids,omitempty"`
	ErrorCode            AuthenticateResponseEx_EAuthenticateCode `protobuf:"varint,20,opt,name=error_code,json=errorCode,enum=axxonsoft.bl.auth.AuthenticateResponseEx_EAuthenticateCode" json:"error_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                 `json:"-"`
	XXX_unrecognized     []byte                                   `json:"-"`
	XXX_sizecache        int32                                    `json:"-"`
}

func (m *AuthenticateResponseEx) Reset()         { *m = AuthenticateResponseEx{} }
func (m *AuthenticateResponseEx) String() string { return proto.CompactTextString(m) }
func (*AuthenticateResponseEx) ProtoMessage()    {}
func (*AuthenticateResponseEx) Descriptor() ([]byte, []int) {
	return fileDescriptor_Authentication_d9f3199dfb1c8a85, []int{2}
}
func (m *AuthenticateResponseEx) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticateResponseEx.Unmarshal(m, b)
}
func (m *AuthenticateResponseEx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticateResponseEx.Marshal(b, m, deterministic)
}
func (dst *AuthenticateResponseEx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticateResponseEx.Merge(dst, src)
}
func (m *AuthenticateResponseEx) XXX_Size() int {
	return xxx_messageInfo_AuthenticateResponseEx.Size(m)
}
func (m *AuthenticateResponseEx) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticateResponseEx.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticateResponseEx proto.InternalMessageInfo

func (m *AuthenticateResponseEx) GetTokenName() string {
	if m != nil {
		return m.TokenName
	}
	return ""
}

func (m *AuthenticateResponseEx) GetTokenValue() string {
	if m != nil {
		return m.TokenValue
	}
	return ""
}

func (m *AuthenticateResponseEx) GetExpiresAt() string {
	if m != nil {
		return m.ExpiresAt
	}
	return ""
}

func (m *AuthenticateResponseEx) GetIsUnrestricted() bool {
	if m != nil {
		return m.IsUnrestricted
	}
	return false
}

func (m *AuthenticateResponseEx) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *AuthenticateResponseEx) GetRolesIds() []string {
	if m != nil {
		return m.RolesIds
	}
	return nil
}

func (m *AuthenticateResponseEx) GetErrorCode() AuthenticateResponseEx_EAuthenticateCode {
	if m != nil {
		return m.ErrorCode
	}
	return AuthenticateResponseEx_AUTHENTICATE_CODE_OK
}

func init() {
	proto.RegisterType((*AuthenticateRequest)(nil), "axxonsoft.bl.auth.AuthenticateRequest")
	proto.RegisterType((*AuthenticateResponse)(nil), "axxonsoft.bl.auth.AuthenticateResponse")
	proto.RegisterType((*AuthenticateResponseEx)(nil), "axxonsoft.bl.auth.AuthenticateResponseEx")
	proto.RegisterEnum("axxonsoft.bl.auth.AuthenticateResponseEx_EAuthenticateCode", AuthenticateResponseEx_EAuthenticateCode_name, AuthenticateResponseEx_EAuthenticateCode_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for AuthenticationService service

type AuthenticationServiceClient interface {
	Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error)
	// noexcept. Instead of exception AuthenticateEx returns error code.
	AuthenticateEx(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponseEx, error)
}

type authenticationServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthenticationServiceClient(cc *grpc.ClientConn) AuthenticationServiceClient {
	return &authenticationServiceClient{cc}
}

func (c *authenticationServiceClient) Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error) {
	out := new(AuthenticateResponse)
	err := grpc.Invoke(ctx, "/axxonsoft.bl.auth.AuthenticationService/Authenticate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationServiceClient) AuthenticateEx(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponseEx, error) {
	out := new(AuthenticateResponseEx)
	err := grpc.Invoke(ctx, "/axxonsoft.bl.auth.AuthenticationService/AuthenticateEx", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AuthenticationService service

type AuthenticationServiceServer interface {
	Authenticate(context.Context, *AuthenticateRequest) (*AuthenticateResponse, error)
	// noexcept. Instead of exception AuthenticateEx returns error code.
	AuthenticateEx(context.Context, *AuthenticateRequest) (*AuthenticateResponseEx, error)
}

func RegisterAuthenticationServiceServer(s *grpc.Server, srv AuthenticationServiceServer) {
	s.RegisterService(&_AuthenticationService_serviceDesc, srv)
}

func _AuthenticationService_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axxonsoft.bl.auth.AuthenticationService/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).Authenticate(ctx, req.(*AuthenticateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationService_AuthenticateEx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).AuthenticateEx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axxonsoft.bl.auth.AuthenticationService/AuthenticateEx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).AuthenticateEx(ctx, req.(*AuthenticateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthenticationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "axxonsoft.bl.auth.AuthenticationService",
	HandlerType: (*AuthenticationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authenticate",
			Handler:    _AuthenticationService_Authenticate_Handler,
		},
		{
			MethodName: "AuthenticateEx",
			Handler:    _AuthenticationService_AuthenticateEx_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "axxonsoft/bl/auth/Authentication.proto",
}

func init() {
	proto.RegisterFile("axxonsoft/bl/auth/Authentication.proto", fileDescriptor_Authentication_d9f3199dfb1c8a85)
}

var fileDescriptor_Authentication_d9f3199dfb1c8a85 = []byte{
	// 651 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x94, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0x71, 0x02, 0x69, 0x32, 0x45, 0xc5, 0x5d, 0x0a, 0x58, 0xa5, 0x55, 0x12, 0xb7, 0x24,
	0xa5, 0x48, 0x8e, 0x28, 0x47, 0x4e, 0x6e, 0xbc, 0xa4, 0x56, 0x83, 0x1d, 0xad, 0x93, 0x16, 0x7a,
	0x59, 0x39, 0xf1, 0xd2, 0x5a, 0x4d, 0xbd, 0xc1, 0xbb, 0x29, 0x39, 0x73, 0xe2, 0xc6, 0x81, 0x23,
	0xaf, 0xc1, 0x9b, 0xc0, 0x23, 0x20, 0x9e, 0x03, 0xd9, 0x29, 0x90, 0xe2, 0xa2, 0xf4, 0xca, 0x2d,
	0xfb, 0xcf, 0x37, 0xbf, 0xe6, 0x8f, 0xc6, 0x03, 0x35, 0x7f, 0x32, 0xe1, 0x91, 0xe0, 0x6f, 0x64,
	0xa3, 0x3f, 0x6c, 0xf8, 0x63, 0x79, 0xd2, 0x30, 0xc7, 0xf2, 0x84, 0x45, 0x32, 0x1c, 0xf8, 0x32,
	0xe4, 0x91, 0x31, 0x8a, 0xb9, 0xe4, 0x68, 0xf9, 0x37, 0x67, 0xf4, 0x87, 0x46, 0xc2, 0xad, 0xae,
	0x1d, 0x73, 0x7e, 0x3c, 0x64, 0x0d, 0x7f, 0x14, 0x36, 0xfc, 0x28, 0xe2, 0x32, 0xe5, 0xc5, 0xb4,
	0x41, 0x77, 0xe0, 0xee, 0x8c, 0x11, 0x23, 0xec, 0xed, 0x98, 0x09, 0x89, 0x1e, 0x42, 0x69, 0x2c,
	0x58, 0x4c, 0x23, 0xff, 0x8c, 0x69, 0x4a, 0x45, 0xd9, 0x2a, 0x91, 0x62, 0x22, 0x38, 0xfe, 0x19,
	0x43, 0xab, 0x50, 0x1c, 0xf9, 0x42, 0xbc, 0xe3, 0x71, 0xa0, 0xe5, 0xa6, 0xb5, 0x5f, 0x6f, 0xfd,
	0x9b, 0x02, 0x2b, 0x97, 0x0d, 0xc5, 0x88, 0x47, 0x82, 0xa1, 0x75, 0x00, 0xc9, 0x4f, 0x59, 0x34,
	0x6b, 0x59, 0x4a, 0x95, 0xd4, 0xb3, 0x0c, 0x8b, 0xd3, 0xf2, 0xb9, 0x3f, 0x1c, 0xb3, 0x0b, 0xdb,
	0x69, 0xc7, 0x41, 0xa2, 0x24, 0xfd, 0x6c, 0x32, 0x0a, 0x63, 0x26, 0xa8, 0x2f, 0xb5, 0xfc, 0xb4,
	0xff, 0x42, 0x31, 0x25, 0xaa, 0xc3, 0x9d, 0x50, 0xd0, 0x71, 0x14, 0x33, 0x21, 0xe3, 0x70, 0x20,
	0x59, 0xa0, 0x41, 0x45, 0xd9, 0x2a, 0x92, 0xa5, 0x50, 0xf4, 0x66, 0x54, 0xf4, 0x00, 0x16, 0xd2,
	0x64, 0x61, 0xa0, 0xa9, 0xa9, 0x49, 0x21, 0x79, 0xda, 0x41, 0x12, 0x39, 0xe6, 0x43, 0x26, 0x68,
	0x18, 0x08, 0x6d, 0xb9, 0x92, 0x4f, 0x62, 0xa5, 0x82, 0x1d, 0x08, 0xfd, 0x73, 0x01, 0xee, 0x5f,
	0x15, 0x0b, 0x4f, 0xfe, 0xef, 0x60, 0xe8, 0x08, 0x80, 0xc5, 0x31, 0x8f, 0xe9, 0x80, 0x07, 0x4c,
	0x5b, 0xa9, 0x28, 0x5b, 0x4b, 0x3b, 0xcf, 0x8d, 0xcc, 0x16, 0x19, 0x57, 0x87, 0x37, 0xf0, 0xac,
	0xde, 0xe4, 0x01, 0x23, 0xa5, 0xd4, 0x2e, 0xf9, 0xa9, 0xff, 0xc8, 0xc3, 0x72, 0x06, 0x40, 0x1a,
	0xac, 0x98, 0xbd, 0xee, 0x1e, 0x76, 0xba, 0x76, 0xd3, 0xec, 0x62, 0xda, 0x74, 0x2d, 0x4c, 0xdd,
	0x7d, 0xf5, 0x06, 0xda, 0x80, 0x72, 0xb6, 0xd2, 0xc2, 0x0e, 0x26, 0x66, 0x9b, 0x62, 0x42, 0x5c,
	0xa2, 0x2a, 0xa8, 0x0e, 0x1b, 0x59, 0xe8, 0x90, 0xb8, 0x4e, 0x8b, 0x36, 0x09, 0xb6, 0x12, 0xdd,
	0x6c, 0x7b, 0x6a, 0x0e, 0xd5, 0x40, 0xcf, 0x82, 0x1d, 0xd3, 0xf3, 0x0e, 0x5d, 0x62, 0x51, 0xfc,
	0xaa, 0x63, 0x13, 0x6c, 0xa9, 0x79, 0x54, 0x85, 0xf5, 0x2c, 0xd7, 0xf3, 0x30, 0xa1, 0x6d, 0xb7,
	0xb9, 0x8f, 0x2d, 0xf5, 0x26, 0xaa, 0xc0, 0x5a, 0x16, 0xb1, 0x3b, 0x74, 0xf7, 0x82, 0xb8, 0x85,
	0x36, 0xa1, 0x92, 0x25, 0xcc, 0x36, 0xc1, 0xa6, 0xf5, 0x9a, 0xb6, 0xdd, 0x56, 0x0b, 0x5b, 0x6a,
	0x01, 0x19, 0xb0, 0x9d, 0xa5, 0x5e, 0xb8, 0xa4, 0x89, 0x69, 0x73, 0xcf, 0x74, 0x5a, 0xf8, 0x52,
	0x84, 0x85, 0x39, 0x11, 0x6c, 0xe7, 0xc0, 0x6c, 0xdb, 0x96, 0x5a, 0x44, 0xdb, 0x50, 0x9b, 0x1b,
	0x95, 0x7a, 0xae, 0xeb, 0xa8, 0x25, 0xf4, 0x04, 0xea, 0xff, 0xfa, 0xff, 0xbc, 0x5e, 0x07, 0x93,
	0x03, 0xdb, 0x73, 0x09, 0x25, 0x6e, 0x1b, 0xab, 0x80, 0x1e, 0x41, 0x35, 0x0b, 0x77, 0xed, 0x97,
	0xf8, 0xc8, 0x75, 0x30, 0xb5, 0xb0, 0x63, 0x63, 0x4b, 0x5d, 0xdc, 0xf9, 0x92, 0x83, 0x7b, 0x97,
	0xcf, 0x91, 0xc7, 0xe2, 0xf3, 0x70, 0xc0, 0xd0, 0x07, 0x05, 0x6e, 0xcf, 0x6e, 0x00, 0xaa, 0xcd,
	0xdd, 0xad, 0xf4, 0x00, 0xad, 0xd6, 0xaf, 0xb9, 0x83, 0x7a, 0xfd, 0xfd, 0xd7, 0xef, 0x9f, 0x72,
	0x55, 0x54, 0x6e, 0x9c, 0x3f, 0x4d, 0x0f, 0xe3, 0x9f, 0x21, 0x66, 0x9f, 0x0c, 0x7d, 0x54, 0x60,
	0x69, 0xd6, 0x01, 0x4f, 0xae, 0x3d, 0xcc, 0xe3, 0x6b, 0x7f, 0x10, 0xfa, 0x76, 0x3a, 0xce, 0x26,
	0xd2, 0xe7, 0x8c, 0x43, 0xd9, 0x64, 0xb7, 0x7a, 0x54, 0xee, 0x87, 0xb2, 0x3f, 0x1e, 0x9c, 0x32,
	0x69, 0xf0, 0xf8, 0xb8, 0x61, 0xfe, 0x7d, 0xe4, 0xfb, 0x85, 0xf4, 0x4a, 0x3f, 0xfb, 0x19, 0x00,
	0x00, 0xff, 0xff, 0xd3, 0x6f, 0xe3, 0xde, 0x00, 0x06, 0x00, 0x00,
}
