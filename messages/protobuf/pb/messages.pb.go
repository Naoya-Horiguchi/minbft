// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages.proto

// Package pb defines MinBFT protocol messages in Protobuf format.

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// MessageType represents message type tag.
type MessageType int32

const (
	MessageType_UNKNOWN         MessageType = 0
	MessageType_REQUEST         MessageType = 1
	MessageType_REPLY           MessageType = 2
	MessageType_PREPARE         MessageType = 3
	MessageType_COMMIT          MessageType = 4
	MessageType_REQ_VIEW_CHANGE MessageType = 5
)

var MessageType_name = map[int32]string{
	0: "UNKNOWN",
	1: "REQUEST",
	2: "REPLY",
	3: "PREPARE",
	4: "COMMIT",
	5: "REQ_VIEW_CHANGE",
}

var MessageType_value = map[string]int32{
	"UNKNOWN":         0,
	"REQUEST":         1,
	"REPLY":           2,
	"PREPARE":         3,
	"COMMIT":          4,
	"REQ_VIEW_CHANGE": 5,
}

func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}

func (MessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{0}
}

// Message represents arbitrary protocol message.
type Message struct {
	// Types that are valid to be assigned to Typed:
	//	*Message_Request
	//	*Message_Reply
	//	*Message_Prepare
	//	*Message_Commit
	//	*Message_ReqViewChange
	Typed                isMessage_Typed `protobuf_oneof:"typed"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

type isMessage_Typed interface {
	isMessage_Typed()
}

type Message_Request struct {
	Request *Request `protobuf:"bytes,1,opt,name=request,proto3,oneof"`
}

type Message_Reply struct {
	Reply *Reply `protobuf:"bytes,2,opt,name=reply,proto3,oneof"`
}

type Message_Prepare struct {
	Prepare *Prepare `protobuf:"bytes,3,opt,name=prepare,proto3,oneof"`
}

type Message_Commit struct {
	Commit *Commit `protobuf:"bytes,4,opt,name=commit,proto3,oneof"`
}

type Message_ReqViewChange struct {
	ReqViewChange *ReqViewChange `protobuf:"bytes,5,opt,name=req_view_change,json=reqViewChange,proto3,oneof"`
}

func (*Message_Request) isMessage_Typed() {}

func (*Message_Reply) isMessage_Typed() {}

func (*Message_Prepare) isMessage_Typed() {}

func (*Message_Commit) isMessage_Typed() {}

func (*Message_ReqViewChange) isMessage_Typed() {}

func (m *Message) GetTyped() isMessage_Typed {
	if m != nil {
		return m.Typed
	}
	return nil
}

func (m *Message) GetRequest() *Request {
	if x, ok := m.GetTyped().(*Message_Request); ok {
		return x.Request
	}
	return nil
}

func (m *Message) GetReply() *Reply {
	if x, ok := m.GetTyped().(*Message_Reply); ok {
		return x.Reply
	}
	return nil
}

func (m *Message) GetPrepare() *Prepare {
	if x, ok := m.GetTyped().(*Message_Prepare); ok {
		return x.Prepare
	}
	return nil
}

func (m *Message) GetCommit() *Commit {
	if x, ok := m.GetTyped().(*Message_Commit); ok {
		return x.Commit
	}
	return nil
}

func (m *Message) GetReqViewChange() *ReqViewChange {
	if x, ok := m.GetTyped().(*Message_ReqViewChange); ok {
		return x.ReqViewChange
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Message) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Message_Request)(nil),
		(*Message_Reply)(nil),
		(*Message_Prepare)(nil),
		(*Message_Commit)(nil),
		(*Message_ReqViewChange)(nil),
	}
}

// Request represents REQUEST message.
type Request struct {
	// Client identifier
	ClientId uint32 `protobuf:"varint,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	// Request identifier (timestamp / sequence number)
	Seq uint64 `protobuf:"varint,2,opt,name=seq,proto3" json:"seq,omitempty"`
	// Operation to execute on replicated state machine
	Operation []byte `protobuf:"bytes,3,opt,name=operation,proto3" json:"operation,omitempty"`
	// Client's signature
	Signature            []byte   `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetClientId() uint32 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *Request) GetSeq() uint64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *Request) GetOperation() []byte {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *Request) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

// Reply represents REPLY message.
type Reply struct {
	// Replica identifier
	ReplicaId uint32 `protobuf:"varint,1,opt,name=replica_id,json=replicaId,proto3" json:"replica_id,omitempty"`
	// Client identifier
	ClientId uint32 `protobuf:"varint,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	// Request identifier
	Seq uint64 `protobuf:"varint,3,opt,name=seq,proto3" json:"seq,omitempty"`
	// Result of requested operation execution
	Result []byte `protobuf:"bytes,4,opt,name=result,proto3" json:"result,omitempty"`
	// Replica's signature
	Signature            []byte   `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Reply) Reset()         { *m = Reply{} }
func (m *Reply) String() string { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()    {}
func (*Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{2}
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

func (m *Reply) GetReplicaId() uint32 {
	if m != nil {
		return m.ReplicaId
	}
	return 0
}

func (m *Reply) GetClientId() uint32 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *Reply) GetSeq() uint64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *Reply) GetResult() []byte {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *Reply) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

// Prepare represents PREPARE message.
type Prepare struct {
	// Replica identifier
	ReplicaId uint32 `protobuf:"varint,1,opt,name=replica_id,json=replicaId,proto3" json:"replica_id,omitempty"`
	// View number
	View uint64 `protobuf:"varint,2,opt,name=view,proto3" json:"view,omitempty"`
	// Client's REQUEST
	Request *Request `protobuf:"bytes,3,opt,name=request,proto3" json:"request,omitempty"`
	// Replica's UI
	Ui                   []byte   `protobuf:"bytes,4,opt,name=ui,proto3" json:"ui,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Prepare) Reset()         { *m = Prepare{} }
func (m *Prepare) String() string { return proto.CompactTextString(m) }
func (*Prepare) ProtoMessage()    {}
func (*Prepare) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{3}
}

func (m *Prepare) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Prepare.Unmarshal(m, b)
}
func (m *Prepare) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Prepare.Marshal(b, m, deterministic)
}
func (m *Prepare) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Prepare.Merge(m, src)
}
func (m *Prepare) XXX_Size() int {
	return xxx_messageInfo_Prepare.Size(m)
}
func (m *Prepare) XXX_DiscardUnknown() {
	xxx_messageInfo_Prepare.DiscardUnknown(m)
}

var xxx_messageInfo_Prepare proto.InternalMessageInfo

func (m *Prepare) GetReplicaId() uint32 {
	if m != nil {
		return m.ReplicaId
	}
	return 0
}

func (m *Prepare) GetView() uint64 {
	if m != nil {
		return m.View
	}
	return 0
}

func (m *Prepare) GetRequest() *Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *Prepare) GetUi() []byte {
	if m != nil {
		return m.Ui
	}
	return nil
}

// Commit represents COMMIT message.
type Commit struct {
	// Replica identifier
	ReplicaId uint32 `protobuf:"varint,1,opt,name=replica_id,json=replicaId,proto3" json:"replica_id,omitempty"`
	// Primary's PREPARE
	Prepare *Prepare `protobuf:"bytes,2,opt,name=prepare,proto3" json:"prepare,omitempty"`
	// Replica's UI
	Ui                   []byte   `protobuf:"bytes,3,opt,name=ui,proto3" json:"ui,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Commit) Reset()         { *m = Commit{} }
func (m *Commit) String() string { return proto.CompactTextString(m) }
func (*Commit) ProtoMessage()    {}
func (*Commit) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{4}
}

func (m *Commit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Commit.Unmarshal(m, b)
}
func (m *Commit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Commit.Marshal(b, m, deterministic)
}
func (m *Commit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Commit.Merge(m, src)
}
func (m *Commit) XXX_Size() int {
	return xxx_messageInfo_Commit.Size(m)
}
func (m *Commit) XXX_DiscardUnknown() {
	xxx_messageInfo_Commit.DiscardUnknown(m)
}

var xxx_messageInfo_Commit proto.InternalMessageInfo

func (m *Commit) GetReplicaId() uint32 {
	if m != nil {
		return m.ReplicaId
	}
	return 0
}

func (m *Commit) GetPrepare() *Prepare {
	if m != nil {
		return m.Prepare
	}
	return nil
}

func (m *Commit) GetUi() []byte {
	if m != nil {
		return m.Ui
	}
	return nil
}

// ReqViewChange represents REQ-VIEW-CHANGE message.
type ReqViewChange struct {
	// Replica identifier
	ReplicaId uint32 `protobuf:"varint,1,opt,name=replica_id,json=replicaId,proto3" json:"replica_id,omitempty"`
	// New view number
	NewView uint64 `protobuf:"varint,2,opt,name=new_view,json=newView,proto3" json:"new_view,omitempty"`
	// Replica's signature
	Signature            []byte   `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqViewChange) Reset()         { *m = ReqViewChange{} }
func (m *ReqViewChange) String() string { return proto.CompactTextString(m) }
func (*ReqViewChange) ProtoMessage()    {}
func (*ReqViewChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{5}
}

func (m *ReqViewChange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqViewChange.Unmarshal(m, b)
}
func (m *ReqViewChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqViewChange.Marshal(b, m, deterministic)
}
func (m *ReqViewChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqViewChange.Merge(m, src)
}
func (m *ReqViewChange) XXX_Size() int {
	return xxx_messageInfo_ReqViewChange.Size(m)
}
func (m *ReqViewChange) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqViewChange.DiscardUnknown(m)
}

var xxx_messageInfo_ReqViewChange proto.InternalMessageInfo

func (m *ReqViewChange) GetReplicaId() uint32 {
	if m != nil {
		return m.ReplicaId
	}
	return 0
}

func (m *ReqViewChange) GetNewView() uint64 {
	if m != nil {
		return m.NewView
	}
	return 0
}

func (m *ReqViewChange) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterEnum("pb.MessageType", MessageType_name, MessageType_value)
	proto.RegisterType((*Message)(nil), "pb.Message")
	proto.RegisterType((*Request)(nil), "pb.Request")
	proto.RegisterType((*Reply)(nil), "pb.Reply")
	proto.RegisterType((*Prepare)(nil), "pb.Prepare")
	proto.RegisterType((*Commit)(nil), "pb.Commit")
	proto.RegisterType((*ReqViewChange)(nil), "pb.ReqViewChange")
}

func init() { proto.RegisterFile("messages.proto", fileDescriptor_4dc296cbfe5ffcd5) }

var fileDescriptor_4dc296cbfe5ffcd5 = []byte{
	// 480 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x6d, 0x92, 0x26, 0x59, 0x6e, 0xd7, 0x2d, 0x18, 0x09, 0x05, 0x01, 0x12, 0x54, 0x4c, 0x20,
	0x1e, 0xfa, 0x00, 0x8f, 0x3c, 0x8d, 0x2a, 0x22, 0x15, 0xb4, 0xeb, 0x4c, 0xb7, 0x89, 0x17, 0x42,
	0xda, 0x5e, 0x15, 0x4b, 0x6d, 0xe2, 0x3a, 0x09, 0x55, 0xbf, 0x81, 0x3f, 0xe5, 0x2b, 0x90, 0x1d,
	0x83, 0xd7, 0x0d, 0xa9, 0x6f, 0xf1, 0xb9, 0x37, 0xf7, 0x9e, 0x73, 0x7c, 0x0c, 0x27, 0x6b, 0x2c,
	0xcb, 0x6c, 0x89, 0x65, 0x9f, 0x8b, 0xa2, 0x2a, 0x88, 0xcd, 0x67, 0xbd, 0xdf, 0x16, 0xf8, 0xa3,
	0x06, 0x26, 0xaf, 0xc0, 0x17, 0xb8, 0xa9, 0xb1, 0xac, 0x22, 0xeb, 0xb9, 0xf5, 0xba, 0xf3, 0xb6,
	0xd3, 0xe7, 0xb3, 0x3e, 0x6d, 0xa0, 0xa4, 0x45, 0xff, 0x56, 0xc9, 0x0b, 0x70, 0x05, 0xf2, 0xd5,
	0x2e, 0xb2, 0x55, 0x5b, 0xd0, 0xb4, 0xf1, 0xd5, 0x2e, 0x69, 0xd1, 0xa6, 0x22, 0x67, 0x71, 0x81,
	0x3c, 0x13, 0x18, 0x39, 0x66, 0xd6, 0xa4, 0x81, 0xe4, 0x2c, 0x5d, 0x25, 0x2f, 0xc1, 0x9b, 0x17,
	0xeb, 0x35, 0xab, 0xa2, 0xb6, 0xea, 0x03, 0xd9, 0x37, 0x50, 0x48, 0xd2, 0xa2, 0xba, 0x46, 0xde,
	0xc3, 0xa9, 0xc0, 0x4d, 0xfa, 0x93, 0xe1, 0x36, 0x9d, 0xff, 0xc8, 0xf2, 0x25, 0x46, 0xae, 0x6a,
	0x7f, 0xa0, 0x29, 0x5e, 0x33, 0xdc, 0x0e, 0x54, 0x21, 0x69, 0xd1, 0xae, 0xb8, 0x0d, 0x7c, 0xf0,
	0xc1, 0xad, 0x76, 0x1c, 0x17, 0xbd, 0x0a, 0x7c, 0xad, 0x86, 0x3c, 0x81, 0x60, 0xbe, 0x62, 0x98,
	0x57, 0x29, 0x5b, 0x28, 0xb5, 0x5d, 0x7a, 0xd4, 0x00, 0xc3, 0x05, 0x09, 0xc1, 0x29, 0x71, 0xa3,
	0xd4, 0xb5, 0xa9, 0xfc, 0x24, 0x4f, 0x21, 0x28, 0x38, 0x8a, 0xac, 0x62, 0x45, 0xae, 0x04, 0x1d,
	0x53, 0x03, 0xc8, 0x6a, 0xc9, 0x96, 0x79, 0x56, 0xd5, 0x02, 0x95, 0x8c, 0x63, 0x6a, 0x80, 0xde,
	0x2f, 0x0b, 0x5c, 0xe5, 0x0e, 0x79, 0x06, 0x20, 0xdd, 0x61, 0xf3, 0xcc, 0x6c, 0x0d, 0x34, 0x32,
	0x5c, 0xec, 0x73, 0xb2, 0xff, 0xcf, 0xc9, 0x31, 0x9c, 0x1e, 0x81, 0x27, 0xb0, 0xac, 0x57, 0x95,
	0x5e, 0xa9, 0x4f, 0xfb, 0x6c, 0xdc, 0xbb, 0x6c, 0x4a, 0xf0, 0xf5, 0x2d, 0x1c, 0xa2, 0x43, 0xa0,
	0x2d, 0xfd, 0xd6, 0x36, 0xa8, 0x6f, 0x72, 0x66, 0x22, 0xe2, 0xdc, 0x8b, 0x88, 0x09, 0xc8, 0x09,
	0xd8, 0x35, 0xd3, 0xb4, 0xec, 0x9a, 0xf5, 0xbe, 0x81, 0xd7, 0x5c, 0xe9, 0xa1, 0x9d, 0x67, 0x26,
	0x36, 0xf6, 0xbd, 0xd8, 0x98, 0xd0, 0x34, 0xf3, 0x9d, 0x7f, 0xf3, 0x97, 0xd0, 0xdd, 0xcb, 0xc0,
	0xa1, 0x35, 0x8f, 0xe1, 0x28, 0xc7, 0x6d, 0x7a, 0x4b, 0x9e, 0x9f, 0xe3, 0x56, 0xfe, 0xbf, 0xef,
	0x9e, 0x73, 0xc7, 0xbd, 0x37, 0xdf, 0xa1, 0xa3, 0x5f, 0xcb, 0x74, 0xc7, 0x91, 0x74, 0xc0, 0xbf,
	0x1a, 0x7f, 0x1a, 0x5f, 0xdc, 0x8c, 0xc3, 0x96, 0x3c, 0xd0, 0xf8, 0xf2, 0x2a, 0xfe, 0x32, 0x0d,
	0x2d, 0x12, 0x80, 0x4b, 0xe3, 0xc9, 0xe7, 0xaf, 0xa1, 0x2d, 0xf1, 0x09, 0x8d, 0x27, 0xe7, 0x34,
	0x0e, 0x1d, 0x02, 0xe0, 0x0d, 0x2e, 0x46, 0xa3, 0xe1, 0x34, 0x6c, 0x93, 0x87, 0x70, 0x4a, 0xe3,
	0xcb, 0xf4, 0x7a, 0x18, 0xdf, 0xa4, 0x83, 0xe4, 0x7c, 0xfc, 0x31, 0x0e, 0xdd, 0x99, 0xa7, 0xde,
	0xe6, 0xbb, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x12, 0x67, 0x4b, 0x45, 0xad, 0x03, 0x00, 0x00,
}