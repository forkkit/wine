// Code generated by protoc-gen-go. DO NOT EDIT.
// source: packet.proto

package websocket

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ef1a6541f9f9e7, []int{0}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Data struct {
	// Types that are valid to be assigned to V:
	//	*Data_Raw
	//	*Data_Json
	//	*Data_Protobuf
	V                    isData_V `protobuf_oneof:"v"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ef1a6541f9f9e7, []int{1}
}

func (m *Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data.Unmarshal(m, b)
}
func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data.Marshal(b, m, deterministic)
}
func (m *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(m, src)
}
func (m *Data) XXX_Size() int {
	return xxx_messageInfo_Data.Size(m)
}
func (m *Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Data proto.InternalMessageInfo

type isData_V interface {
	isData_V()
}

type Data_Raw struct {
	Raw []byte `protobuf:"bytes,1,opt,name=raw,proto3,oneof"`
}

type Data_Json struct {
	Json []byte `protobuf:"bytes,2,opt,name=json,proto3,oneof"`
}

type Data_Protobuf struct {
	Protobuf []byte `protobuf:"bytes,3,opt,name=protobuf,proto3,oneof"`
}

func (*Data_Raw) isData_V() {}

func (*Data_Json) isData_V() {}

func (*Data_Protobuf) isData_V() {}

func (m *Data) GetV() isData_V {
	if m != nil {
		return m.V
	}
	return nil
}

func (m *Data) GetRaw() []byte {
	if x, ok := m.GetV().(*Data_Raw); ok {
		return x.Raw
	}
	return nil
}

func (m *Data) GetJson() []byte {
	if x, ok := m.GetV().(*Data_Json); ok {
		return x.Json
	}
	return nil
}

func (m *Data) GetProtobuf() []byte {
	if x, ok := m.GetV().(*Data_Protobuf); ok {
		return x.Protobuf
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Data) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Data_Raw)(nil),
		(*Data_Json)(nil),
		(*Data_Protobuf)(nil),
	}
}

type Push struct {
	Type                 int32    `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Data                 *Data    `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Push) Reset()         { *m = Push{} }
func (m *Push) String() string { return proto.CompactTextString(m) }
func (*Push) ProtoMessage()    {}
func (*Push) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ef1a6541f9f9e7, []int{2}
}

func (m *Push) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Push.Unmarshal(m, b)
}
func (m *Push) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Push.Marshal(b, m, deterministic)
}
func (m *Push) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Push.Merge(m, src)
}
func (m *Push) XXX_Size() int {
	return xxx_messageInfo_Push.Size(m)
}
func (m *Push) XXX_DiscardUnknown() {
	xxx_messageInfo_Push.DiscardUnknown(m)
}

var xxx_messageInfo_Push proto.InternalMessageInfo

func (m *Push) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Push) GetData() *Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type Call struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Data                 *Data    `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Call) Reset()         { *m = Call{} }
func (m *Call) String() string { return proto.CompactTextString(m) }
func (*Call) ProtoMessage()    {}
func (*Call) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ef1a6541f9f9e7, []int{3}
}

func (m *Call) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Call.Unmarshal(m, b)
}
func (m *Call) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Call.Marshal(b, m, deterministic)
}
func (m *Call) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Call.Merge(m, src)
}
func (m *Call) XXX_Size() int {
	return xxx_messageInfo_Call.Size(m)
}
func (m *Call) XXX_DiscardUnknown() {
	xxx_messageInfo_Call.DiscardUnknown(m)
}

var xxx_messageInfo_Call proto.InternalMessageInfo

func (m *Call) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Call) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Call) GetData() *Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type Reply struct {
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are valid to be assigned to Result:
	//	*Reply_Data
	//	*Reply_Error
	Result               isReply_Result `protobuf_oneof:"result"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Reply) Reset()         { *m = Reply{} }
func (m *Reply) String() string { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()    {}
func (*Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ef1a6541f9f9e7, []int{4}
}

func (m *Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Reply.Unmarshal(m, b)
}
func (m *Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Reply.Marshal(b, m, deterministic)
}
func (m *Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reply.Merge(m, src)
}
func (m *Reply) XXX_Size() int {
	return xxx_messageInfo_Reply.Size(m)
}
func (m *Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_Reply.DiscardUnknown(m)
}

var xxx_messageInfo_Reply proto.InternalMessageInfo

func (m *Reply) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type isReply_Result interface {
	isReply_Result()
}

type Reply_Data struct {
	Data *Data `protobuf:"bytes,2,opt,name=data,proto3,oneof"`
}

type Reply_Error struct {
	Error *Error `protobuf:"bytes,3,opt,name=error,proto3,oneof"`
}

func (*Reply_Data) isReply_Result() {}

func (*Reply_Error) isReply_Result() {}

func (m *Reply) GetResult() isReply_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *Reply) GetData() *Data {
	if x, ok := m.GetResult().(*Reply_Data); ok {
		return x.Data
	}
	return nil
}

func (m *Reply) GetError() *Error {
	if x, ok := m.GetResult().(*Reply_Error); ok {
		return x.Error
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Reply) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Reply_Data)(nil),
		(*Reply_Error)(nil),
	}
}

type Header struct {
	Entries              map[string]string `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}
func (*Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ef1a6541f9f9e7, []int{5}
}

func (m *Header) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Header.Unmarshal(m, b)
}
func (m *Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Header.Marshal(b, m, deterministic)
}
func (m *Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Header.Merge(m, src)
}
func (m *Header) XXX_Size() int {
	return xxx_messageInfo_Header.Size(m)
}
func (m *Header) XXX_DiscardUnknown() {
	xxx_messageInfo_Header.DiscardUnknown(m)
}

var xxx_messageInfo_Header proto.InternalMessageInfo

func (m *Header) GetEntries() map[string]string {
	if m != nil {
		return m.Entries
	}
	return nil
}

type Hello struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Hello) Reset()         { *m = Hello{} }
func (m *Hello) String() string { return proto.CompactTextString(m) }
func (*Hello) ProtoMessage()    {}
func (*Hello) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ef1a6541f9f9e7, []int{6}
}

func (m *Hello) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Hello.Unmarshal(m, b)
}
func (m *Hello) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Hello.Marshal(b, m, deterministic)
}
func (m *Hello) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hello.Merge(m, src)
}
func (m *Hello) XXX_Size() int {
	return xxx_messageInfo_Hello.Size(m)
}
func (m *Hello) XXX_DiscardUnknown() {
	xxx_messageInfo_Hello.DiscardUnknown(m)
}

var xxx_messageInfo_Hello proto.InternalMessageInfo

type Packet struct {
	// Types that are valid to be assigned to V:
	//	*Packet_Call
	//	*Packet_Data
	//	*Packet_Header
	//	*Packet_Hello
	//	*Packet_Push
	//	*Packet_Reply
	V                    isPacket_V `protobuf_oneof:"v"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Packet) Reset()         { *m = Packet{} }
func (m *Packet) String() string { return proto.CompactTextString(m) }
func (*Packet) ProtoMessage()    {}
func (*Packet) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ef1a6541f9f9e7, []int{7}
}

func (m *Packet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Packet.Unmarshal(m, b)
}
func (m *Packet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Packet.Marshal(b, m, deterministic)
}
func (m *Packet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Packet.Merge(m, src)
}
func (m *Packet) XXX_Size() int {
	return xxx_messageInfo_Packet.Size(m)
}
func (m *Packet) XXX_DiscardUnknown() {
	xxx_messageInfo_Packet.DiscardUnknown(m)
}

var xxx_messageInfo_Packet proto.InternalMessageInfo

type isPacket_V interface {
	isPacket_V()
}

type Packet_Call struct {
	Call *Call `protobuf:"bytes,1,opt,name=call,proto3,oneof"`
}

type Packet_Data struct {
	Data *Data `protobuf:"bytes,2,opt,name=data,proto3,oneof"`
}

type Packet_Header struct {
	Header *Header `protobuf:"bytes,3,opt,name=header,proto3,oneof"`
}

type Packet_Hello struct {
	Hello *Hello `protobuf:"bytes,4,opt,name=hello,proto3,oneof"`
}

type Packet_Push struct {
	Push *Push `protobuf:"bytes,5,opt,name=push,proto3,oneof"`
}

type Packet_Reply struct {
	Reply *Reply `protobuf:"bytes,6,opt,name=reply,proto3,oneof"`
}

func (*Packet_Call) isPacket_V() {}

func (*Packet_Data) isPacket_V() {}

func (*Packet_Header) isPacket_V() {}

func (*Packet_Hello) isPacket_V() {}

func (*Packet_Push) isPacket_V() {}

func (*Packet_Reply) isPacket_V() {}

func (m *Packet) GetV() isPacket_V {
	if m != nil {
		return m.V
	}
	return nil
}

func (m *Packet) GetCall() *Call {
	if x, ok := m.GetV().(*Packet_Call); ok {
		return x.Call
	}
	return nil
}

func (m *Packet) GetData() *Data {
	if x, ok := m.GetV().(*Packet_Data); ok {
		return x.Data
	}
	return nil
}

func (m *Packet) GetHeader() *Header {
	if x, ok := m.GetV().(*Packet_Header); ok {
		return x.Header
	}
	return nil
}

func (m *Packet) GetHello() *Hello {
	if x, ok := m.GetV().(*Packet_Hello); ok {
		return x.Hello
	}
	return nil
}

func (m *Packet) GetPush() *Push {
	if x, ok := m.GetV().(*Packet_Push); ok {
		return x.Push
	}
	return nil
}

func (m *Packet) GetReply() *Reply {
	if x, ok := m.GetV().(*Packet_Reply); ok {
		return x.Reply
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Packet) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Packet_Call)(nil),
		(*Packet_Data)(nil),
		(*Packet_Header)(nil),
		(*Packet_Hello)(nil),
		(*Packet_Push)(nil),
		(*Packet_Reply)(nil),
	}
}

func init() {
	proto.RegisterType((*Error)(nil), "ws.Error")
	proto.RegisterType((*Data)(nil), "ws.Data")
	proto.RegisterType((*Push)(nil), "ws.Push")
	proto.RegisterType((*Call)(nil), "ws.Call")
	proto.RegisterType((*Reply)(nil), "ws.Reply")
	proto.RegisterType((*Header)(nil), "ws.Header")
	proto.RegisterMapType((map[string]string)(nil), "ws.Header.EntriesEntry")
	proto.RegisterType((*Hello)(nil), "ws.Hello")
	proto.RegisterType((*Packet)(nil), "ws.Packet")
}

func init() { proto.RegisterFile("packet.proto", fileDescriptor_e9ef1a6541f9f9e7) }

var fileDescriptor_e9ef1a6541f9f9e7 = []byte{
	// 438 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x4d, 0x6b, 0xdb, 0x40,
	0x10, 0xb5, 0x3e, 0x6d, 0x4f, 0x4c, 0x29, 0x4b, 0xa0, 0x4b, 0x09, 0xad, 0x23, 0x7a, 0xf0, 0x49,
	0xa6, 0x29, 0x85, 0x90, 0xa3, 0xdb, 0x80, 0x8e, 0x61, 0x7b, 0xeb, 0x6d, 0x65, 0x6d, 0x2d, 0x35,
	0x6b, 0xaf, 0xd8, 0x5d, 0x45, 0xf8, 0xa7, 0xf6, 0xdf, 0x94, 0x19, 0xcb, 0x76, 0xa1, 0x86, 0x9c,
	0x34, 0x33, 0x4f, 0xf3, 0xe6, 0xbd, 0xd9, 0x81, 0x59, 0x2b, 0xd7, 0xcf, 0xca, 0xe7, 0xad, 0x35,
	0xde, 0xb0, 0xb0, 0x77, 0xd9, 0x57, 0x48, 0x1e, 0xad, 0x35, 0x96, 0x31, 0x88, 0xd7, 0xa6, 0x52,
	0x3c, 0x98, 0x07, 0x8b, 0x44, 0x50, 0xcc, 0x38, 0x8c, 0xb7, 0xca, 0x39, 0xb9, 0x51, 0x3c, 0x9c,
	0x07, 0x8b, 0xa9, 0x38, 0xa6, 0xd9, 0x0f, 0x88, 0xbf, 0x4b, 0x2f, 0x19, 0x83, 0xc8, 0xca, 0x9e,
	0x9a, 0x66, 0xc5, 0x48, 0x60, 0xc2, 0xae, 0x21, 0xfe, 0xed, 0xcc, 0x8e, 0x5a, 0xb0, 0x48, 0x19,
	0xbb, 0x81, 0x09, 0x4d, 0x2d, 0xbb, 0x5f, 0x3c, 0x1a, 0x90, 0x53, 0x65, 0x15, 0x41, 0xf0, 0x92,
	0xdd, 0x43, 0xfc, 0xd4, 0xb9, 0x1a, 0xa5, 0xf8, 0x7d, 0x7b, 0x92, 0x82, 0x31, 0xbb, 0x81, 0xb8,
	0x92, 0x5e, 0x12, 0xe9, 0xd5, 0xdd, 0x24, 0xef, 0x5d, 0x8e, 0x02, 0x04, 0x55, 0xb3, 0x02, 0xe2,
	0x6f, 0x52, 0x6b, 0xf6, 0x06, 0xc2, 0xa6, 0x1a, 0xfa, 0xc2, 0xa6, 0x42, 0xa6, 0x9d, 0xdc, 0x1e,
	0xd5, 0x53, 0x7c, 0x62, 0x8a, 0x2e, 0x32, 0x55, 0x90, 0x08, 0xd5, 0xea, 0xfd, 0x7f, 0x54, 0x1f,
	0x2e, 0x0b, 0x40, 0x7f, 0x58, 0x67, 0xb7, 0x90, 0x28, 0x5c, 0xe4, 0xc0, 0x3b, 0xc5, 0x1f, 0x68,
	0xb3, 0xc5, 0x48, 0x1c, 0x90, 0xd5, 0x04, 0x52, 0xab, 0x5c, 0xa7, 0x7d, 0xd6, 0x43, 0x5a, 0x28,
	0x59, 0x29, 0xcb, 0x3e, 0xc3, 0x58, 0xed, 0xbc, 0x6d, 0x94, 0xe3, 0xc1, 0x3c, 0x5a, 0x5c, 0xdd,
	0xbd, 0xc3, 0xc6, 0x03, 0x98, 0x3f, 0x1e, 0x10, 0xfc, 0xec, 0xc5, 0xf1, 0xbf, 0xf7, 0x0f, 0x30,
	0xfb, 0x17, 0x60, 0x6f, 0x21, 0x7a, 0x56, 0x7b, 0x92, 0x3a, 0x15, 0x18, 0xb2, 0x6b, 0x48, 0x5e,
	0xa4, 0xee, 0x8e, 0xbe, 0x0f, 0xc9, 0x43, 0x78, 0x1f, 0x64, 0x63, 0x48, 0x0a, 0xa5, 0xb5, 0xc9,
	0xfe, 0x04, 0x90, 0x3e, 0xd1, 0x31, 0xa0, 0xb3, 0xb5, 0xd4, 0x9a, 0x08, 0x06, 0x67, 0xb8, 0x4c,
	0x74, 0x86, 0xf5, 0x57, 0x9d, 0x7f, 0x82, 0xb4, 0x26, 0xbd, 0x83, 0x75, 0x38, 0x3b, 0x28, 0x46,
	0x62, 0xc0, 0x70, 0x3f, 0x35, 0x4e, 0xe6, 0xf1, 0x79, 0x3f, 0x24, 0x05, 0xf7, 0x43, 0x08, 0x0e,
	0x6a, 0x3b, 0x57, 0xf3, 0xe4, 0x3c, 0x08, 0xef, 0x01, 0x07, 0x61, 0x1d, 0x29, 0x2c, 0xbe, 0x0d,
	0x4f, 0xcf, 0x14, 0xf4, 0x58, 0x48, 0x41, 0x08, 0xdd, 0xd1, 0xea, 0xf6, 0xe7, 0xc7, 0x4d, 0xe3,
	0xeb, 0xae, 0xcc, 0xd7, 0x66, 0xbb, 0xdc, 0x98, 0xb6, 0x2b, 0x97, 0x7d, 0xb3, 0x53, 0xcb, 0x5e,
	0x95, 0xce, 0xa0, 0xe7, 0x32, 0xa5, 0xcb, 0xfb, 0xf2, 0x37, 0x00, 0x00, 0xff, 0xff, 0x25, 0x2a,
	0x97, 0xab, 0x11, 0x03, 0x00, 0x00,
}
