// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/api/http.proto

package google_api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// `HttpRule` defines the mapping of an RPC method to one or more HTTP REST API
// methods. The mapping determines what portions of the request message are
// populated from the path, query parameters, or body of the HTTP request.  The
// mapping is typically specified as an `google.api.http` annotation, see
// "google/api/annotations.proto" for details.
//
// The mapping consists of a mandatory field specifying a path template and an
// optional `body` field specifying what data is represented in the HTTP request
// body. The field name for the path indicates the HTTP method. Example:
//
// ```
// package google.storage.v2;
//
// import "google/api/annotations.proto";
//
// service Storage {
//   rpc CreateObject(CreateObjectRequest) returns (Object) {
//     option (google.api.http) {
//       post: "/v2/{bucket_name=buckets/*}/objects"
//       body: "object"
//     };
//   };
// }
// ```
//
// Here `bucket_name` and `object` bind to fields of the request message
// `CreateObjectRequest`.
//
// The rules for mapping HTTP path, query parameters, and body fields
// to the request message are as follows:
//
// 1. The `body` field specifies either `*` or a field path, or is
//    omitted. If omitted, it assumes there is no HTTP body.
// 2. Leaf fields (recursive expansion of nested messages in the
//    request) can be classified into three types:
//     (a) Matched in the URL template.
//     (b) Covered by body (if body is `*`, everything except (a) fields;
//         else everything under the body field)
//     (c) All other fields.
// 3. URL query parameters found in the HTTP request are mapped to (c) fields.
// 4. Any body sent with an HTTP request can contain only (b) fields.
//
// The syntax of the path template is as follows:
//
//     Template = "/" Segments [ Verb ] ;
//     Segments = Segment { "/" Segment } ;
//     Segment  = "*" | "**" | LITERAL | Variable ;
//     Variable = "{" FieldPath [ "=" Segments ] "}" ;
//     FieldPath = IDENT { "." IDENT } ;
//     Verb     = ":" LITERAL ;
//
// `*` matches a single path component, `**` zero or more path components, and
// `LITERAL` a constant.  A `Variable` can match an entire path as specified
// again by a template; this nested template must not contain further variables.
// If no template is given with a variable, it matches a single path component.
// The notation `{var}` is henceforth equivalent to `{var=*}`.
//
// Use CustomHttpPattern to specify any HTTP method that is not included in the
// pattern field, such as HEAD, or "*" to leave the HTTP method unspecified for
// a given URL path rule. The wild-card rule is useful for services that provide
// content to Web (HTML) clients.
type HttpRule struct {
	// Determines the URL pattern is matched by this rules. This pattern can be
	// used with any of the {get|put|post|delete|patch} methods. A custom method
	// can be defined using the 'custom' field.
	//
	// Types that are valid to be assigned to Pattern:
	//	*HttpRule_Get
	//	*HttpRule_Put
	//	*HttpRule_Post
	//	*HttpRule_Delete
	//	*HttpRule_Patch
	//	*HttpRule_Custom
	Pattern isHttpRule_Pattern `protobuf_oneof:"pattern"`
	// The name of the request field whose value is mapped to the HTTP body, or
	// `*` for mapping all fields not captured by the path pattern to the HTTP
	// body.
	Body string `protobuf:"bytes,7,opt,name=body" json:"body,omitempty"`
	// Additional HTTP bindings for the selector. Nested bindings must not
	// specify a selector and must not contain additional bindings.
	AdditionalBindings   []*HttpRule `protobuf:"bytes,11,rep,name=additional_bindings,json=additionalBindings" json:"additional_bindings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *HttpRule) Reset()         { *m = HttpRule{} }
func (m *HttpRule) String() string { return proto.CompactTextString(m) }
func (*HttpRule) ProtoMessage()    {}
func (*HttpRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_http_995c0277c5ebfa94, []int{0}
}
func (m *HttpRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpRule.Unmarshal(m, b)
}
func (m *HttpRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpRule.Marshal(b, m, deterministic)
}
func (dst *HttpRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpRule.Merge(dst, src)
}
func (m *HttpRule) XXX_Size() int {
	return xxx_messageInfo_HttpRule.Size(m)
}
func (m *HttpRule) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpRule.DiscardUnknown(m)
}

var xxx_messageInfo_HttpRule proto.InternalMessageInfo

type isHttpRule_Pattern interface {
	isHttpRule_Pattern()
}

type HttpRule_Get struct {
	Get string `protobuf:"bytes,2,opt,name=get,oneof"`
}
type HttpRule_Put struct {
	Put string `protobuf:"bytes,3,opt,name=put,oneof"`
}
type HttpRule_Post struct {
	Post string `protobuf:"bytes,4,opt,name=post,oneof"`
}
type HttpRule_Delete struct {
	Delete string `protobuf:"bytes,5,opt,name=delete,oneof"`
}
type HttpRule_Patch struct {
	Patch string `protobuf:"bytes,6,opt,name=patch,oneof"`
}
type HttpRule_Custom struct {
	Custom *CustomHttpPattern `protobuf:"bytes,8,opt,name=custom,oneof"`
}

func (*HttpRule_Get) isHttpRule_Pattern()    {}
func (*HttpRule_Put) isHttpRule_Pattern()    {}
func (*HttpRule_Post) isHttpRule_Pattern()   {}
func (*HttpRule_Delete) isHttpRule_Pattern() {}
func (*HttpRule_Patch) isHttpRule_Pattern()  {}
func (*HttpRule_Custom) isHttpRule_Pattern() {}

func (m *HttpRule) GetPattern() isHttpRule_Pattern {
	if m != nil {
		return m.Pattern
	}
	return nil
}

func (m *HttpRule) GetGet() string {
	if x, ok := m.GetPattern().(*HttpRule_Get); ok {
		return x.Get
	}
	return ""
}

func (m *HttpRule) GetPut() string {
	if x, ok := m.GetPattern().(*HttpRule_Put); ok {
		return x.Put
	}
	return ""
}

func (m *HttpRule) GetPost() string {
	if x, ok := m.GetPattern().(*HttpRule_Post); ok {
		return x.Post
	}
	return ""
}

func (m *HttpRule) GetDelete() string {
	if x, ok := m.GetPattern().(*HttpRule_Delete); ok {
		return x.Delete
	}
	return ""
}

func (m *HttpRule) GetPatch() string {
	if x, ok := m.GetPattern().(*HttpRule_Patch); ok {
		return x.Patch
	}
	return ""
}

func (m *HttpRule) GetCustom() *CustomHttpPattern {
	if x, ok := m.GetPattern().(*HttpRule_Custom); ok {
		return x.Custom
	}
	return nil
}

func (m *HttpRule) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *HttpRule) GetAdditionalBindings() []*HttpRule {
	if m != nil {
		return m.AdditionalBindings
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*HttpRule) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _HttpRule_OneofMarshaler, _HttpRule_OneofUnmarshaler, _HttpRule_OneofSizer, []interface{}{
		(*HttpRule_Get)(nil),
		(*HttpRule_Put)(nil),
		(*HttpRule_Post)(nil),
		(*HttpRule_Delete)(nil),
		(*HttpRule_Patch)(nil),
		(*HttpRule_Custom)(nil),
	}
}

func _HttpRule_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*HttpRule)
	// pattern
	switch x := m.Pattern.(type) {
	case *HttpRule_Get:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Get)
	case *HttpRule_Put:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Put)
	case *HttpRule_Post:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Post)
	case *HttpRule_Delete:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Delete)
	case *HttpRule_Patch:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Patch)
	case *HttpRule_Custom:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Custom); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("HttpRule.Pattern has unexpected type %T", x)
	}
	return nil
}

func _HttpRule_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*HttpRule)
	switch tag {
	case 2: // pattern.get
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Pattern = &HttpRule_Get{x}
		return true, err
	case 3: // pattern.put
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Pattern = &HttpRule_Put{x}
		return true, err
	case 4: // pattern.post
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Pattern = &HttpRule_Post{x}
		return true, err
	case 5: // pattern.delete
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Pattern = &HttpRule_Delete{x}
		return true, err
	case 6: // pattern.patch
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Pattern = &HttpRule_Patch{x}
		return true, err
	case 8: // pattern.custom
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CustomHttpPattern)
		err := b.DecodeMessage(msg)
		m.Pattern = &HttpRule_Custom{msg}
		return true, err
	default:
		return false, nil
	}
}

func _HttpRule_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*HttpRule)
	// pattern
	switch x := m.Pattern.(type) {
	case *HttpRule_Get:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Get)))
		n += len(x.Get)
	case *HttpRule_Put:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Put)))
		n += len(x.Put)
	case *HttpRule_Post:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Post)))
		n += len(x.Post)
	case *HttpRule_Delete:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Delete)))
		n += len(x.Delete)
	case *HttpRule_Patch:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Patch)))
		n += len(x.Patch)
	case *HttpRule_Custom:
		s := proto.Size(x.Custom)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// A custom pattern is used for defining custom HTTP verb.
type CustomHttpPattern struct {
	// The name of this custom HTTP verb.
	Kind string `protobuf:"bytes,1,opt,name=kind" json:"kind,omitempty"`
	// The path matched by this custom verb.
	Path                 string   `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomHttpPattern) Reset()         { *m = CustomHttpPattern{} }
func (m *CustomHttpPattern) String() string { return proto.CompactTextString(m) }
func (*CustomHttpPattern) ProtoMessage()    {}
func (*CustomHttpPattern) Descriptor() ([]byte, []int) {
	return fileDescriptor_http_995c0277c5ebfa94, []int{1}
}
func (m *CustomHttpPattern) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomHttpPattern.Unmarshal(m, b)
}
func (m *CustomHttpPattern) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomHttpPattern.Marshal(b, m, deterministic)
}
func (dst *CustomHttpPattern) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomHttpPattern.Merge(dst, src)
}
func (m *CustomHttpPattern) XXX_Size() int {
	return xxx_messageInfo_CustomHttpPattern.Size(m)
}
func (m *CustomHttpPattern) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomHttpPattern.DiscardUnknown(m)
}

var xxx_messageInfo_CustomHttpPattern proto.InternalMessageInfo

func (m *CustomHttpPattern) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *CustomHttpPattern) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func init() {
	proto.RegisterType((*HttpRule)(nil), "google.api.HttpRule")
	proto.RegisterType((*CustomHttpPattern)(nil), "google.api.CustomHttpPattern")
}

func init() { proto.RegisterFile("google/api/http.proto", fileDescriptor_http_995c0277c5ebfa94) }

var fileDescriptor_http_995c0277c5ebfa94 = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xc1, 0x6a, 0xbb, 0x40,
	0x10, 0xc6, 0xff, 0x46, 0x63, 0xe2, 0x08, 0x7f, 0xe8, 0x36, 0x2d, 0x7b, 0x09, 0x48, 0x4e, 0x9e,
	0x0c, 0xa4, 0x87, 0x1e, 0x7a, 0xb3, 0x14, 0x72, 0x0c, 0xbe, 0x40, 0x59, 0xdd, 0x45, 0x97, 0x1a,
	0x77, 0x88, 0xe3, 0xa1, 0x8f, 0xd2, 0xb7, 0x2d, 0xbb, 0x6b, 0x48, 0xa0, 0xb7, 0xf9, 0x7e, 0x33,
	0x7e, 0xdf, 0xb8, 0x03, 0x4f, 0xad, 0x31, 0x6d, 0xaf, 0xf6, 0x02, 0xf5, 0xbe, 0x23, 0xc2, 0x02,
	0x2f, 0x86, 0x0c, 0x03, 0x8f, 0x0b, 0x81, 0x7a, 0xf7, 0xb3, 0x80, 0xf5, 0x91, 0x08, 0xab, 0xa9,
	0x57, 0x8c, 0x41, 0xd8, 0x2a, 0xe2, 0x8b, 0x2c, 0xc8, 0x93, 0xe3, 0xbf, 0xca, 0x0a, 0xcb, 0x70,
	0x22, 0x1e, 0x5e, 0x19, 0x4e, 0xc4, 0x36, 0x10, 0xa1, 0x19, 0x89, 0x47, 0x33, 0x74, 0x8a, 0x71,
	0x88, 0xa5, 0xea, 0x15, 0x29, 0xbe, 0x9c, 0xf9, 0xac, 0xd9, 0x33, 0x2c, 0x51, 0x50, 0xd3, 0xf1,
	0x78, 0x6e, 0x78, 0xc9, 0x5e, 0x21, 0x6e, 0xa6, 0x91, 0xcc, 0x99, 0xaf, 0xb3, 0x20, 0x4f, 0x0f,
	0xdb, 0xe2, 0xb6, 0x59, 0xf1, 0xee, 0x3a, 0x76, 0xb7, 0x93, 0x20, 0x52, 0x97, 0xc1, 0x1a, 0xfa,
	0x71, 0xc6, 0x20, 0xaa, 0x8d, 0xfc, 0xe6, 0x2b, 0xeb, 0x57, 0xb9, 0x9a, 0x7d, 0xc0, 0xa3, 0x90,
	0x52, 0x93, 0x36, 0x83, 0xe8, 0x3f, 0x6b, 0x3d, 0x48, 0x3d, 0xb4, 0x23, 0x4f, 0xb3, 0x30, 0x4f,
	0x0f, 0x9b, 0x7b, 0xe7, 0xeb, 0xff, 0x56, 0xec, 0xf6, 0x41, 0x39, 0xcf, 0x97, 0x09, 0xac, 0xd0,
	0xe7, 0xed, 0xde, 0xe0, 0xe1, 0xcf, 0x12, 0x36, 0xfa, 0x4b, 0x0f, 0x92, 0x07, 0x3e, 0xda, 0xd6,
	0x96, 0xa1, 0xa0, 0xce, 0x3f, 0x5c, 0xe5, 0xea, 0x72, 0x0b, 0xff, 0x1b, 0x73, 0xbe, 0x8b, 0x2d,
	0x13, 0x67, 0x63, 0x2f, 0x70, 0x0a, 0xea, 0xd8, 0x9d, 0xe2, 0xe5, 0x37, 0x00, 0x00, 0xff, 0xff,
	0x2f, 0x89, 0x57, 0x7f, 0xa3, 0x01, 0x00, 0x00,
}
