// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mixmessages.proto

/*
Package mixmessages is a generated protocol buffer package.

It is generated from these files:
	mixmessages.proto

It has these top-level messages:
	Ack
	Ping
	Pong
	PrecompDecryptSlot
	PrecompDecryptMessage
	PrecompEncryptSlot
	PrecompEncryptMessage
	PrecompPermuteSlot
	PrecompPermuteMessage
	ErrorMessage
	ErrorAck
*/
package mixmessages

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

// Generic empty Ack message
type Ack struct {
}

func (m *Ack) Reset()                    { *m = Ack{} }
func (m *Ack) String() string            { return proto.CompactTextString(m) }
func (*Ack) ProtoMessage()               {}
func (*Ack) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// The request message asking if server is online
type Ping struct {
}

func (m *Ping) Reset()                    { *m = Ping{} }
func (m *Ping) String() string            { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()               {}
func (*Ping) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// The response message containing the online confirmation
type Pong struct {
}

func (m *Pong) Reset()                    { *m = Pong{} }
func (m *Pong) String() string            { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()               {}
func (*Pong) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// Message for individual Precomp Decrypt Slot
type PrecompDecryptSlot struct {
	Slot                         uint64 `protobuf:"varint,1,opt,name=Slot" json:"Slot,omitempty"`
	EncryptedMessageKeys         []byte `protobuf:"bytes,2,opt,name=EncryptedMessageKeys,proto3" json:"EncryptedMessageKeys,omitempty"`
	EncryptedRecipientIDKeys     []byte `protobuf:"bytes,3,opt,name=EncryptedRecipientIDKeys,proto3" json:"EncryptedRecipientIDKeys,omitempty"`
	PartialMessageCypherText     []byte `protobuf:"bytes,4,opt,name=PartialMessageCypherText,proto3" json:"PartialMessageCypherText,omitempty"`
	PartialRecipientIDCypherText []byte `protobuf:"bytes,5,opt,name=PartialRecipientIDCypherText,proto3" json:"PartialRecipientIDCypherText,omitempty"`
}

func (m *PrecompDecryptSlot) Reset()                    { *m = PrecompDecryptSlot{} }
func (m *PrecompDecryptSlot) String() string            { return proto.CompactTextString(m) }
func (*PrecompDecryptSlot) ProtoMessage()               {}
func (*PrecompDecryptSlot) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PrecompDecryptSlot) GetSlot() uint64 {
	if m != nil {
		return m.Slot
	}
	return 0
}

func (m *PrecompDecryptSlot) GetEncryptedMessageKeys() []byte {
	if m != nil {
		return m.EncryptedMessageKeys
	}
	return nil
}

func (m *PrecompDecryptSlot) GetEncryptedRecipientIDKeys() []byte {
	if m != nil {
		return m.EncryptedRecipientIDKeys
	}
	return nil
}

func (m *PrecompDecryptSlot) GetPartialMessageCypherText() []byte {
	if m != nil {
		return m.PartialMessageCypherText
	}
	return nil
}

func (m *PrecompDecryptSlot) GetPartialRecipientIDCypherText() []byte {
	if m != nil {
		return m.PartialRecipientIDCypherText
	}
	return nil
}

// Message for batch of Precomp Decrypt Slots
type PrecompDecryptMessage struct {
	RoundID string                `protobuf:"bytes,1,opt,name=RoundID" json:"RoundID,omitempty"`
	Slots   []*PrecompDecryptSlot `protobuf:"bytes,2,rep,name=Slots" json:"Slots,omitempty"`
}

func (m *PrecompDecryptMessage) Reset()                    { *m = PrecompDecryptMessage{} }
func (m *PrecompDecryptMessage) String() string            { return proto.CompactTextString(m) }
func (*PrecompDecryptMessage) ProtoMessage()               {}
func (*PrecompDecryptMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *PrecompDecryptMessage) GetRoundID() string {
	if m != nil {
		return m.RoundID
	}
	return ""
}

func (m *PrecompDecryptMessage) GetSlots() []*PrecompDecryptSlot {
	if m != nil {
		return m.Slots
	}
	return nil
}

// Message for individual Precomp Encrypt Slot
type PrecompEncryptSlot struct {
	Slot                     uint64 `protobuf:"varint,1,opt,name=Slot" json:"Slot,omitempty"`
	EncryptedMessageKeys     []byte `protobuf:"bytes,2,opt,name=EncryptedMessageKeys,proto3" json:"EncryptedMessageKeys,omitempty"`
	PartialMessageCypherText []byte `protobuf:"bytes,3,opt,name=PartialMessageCypherText,proto3" json:"PartialMessageCypherText,omitempty"`
}

func (m *PrecompEncryptSlot) Reset()                    { *m = PrecompEncryptSlot{} }
func (m *PrecompEncryptSlot) String() string            { return proto.CompactTextString(m) }
func (*PrecompEncryptSlot) ProtoMessage()               {}
func (*PrecompEncryptSlot) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PrecompEncryptSlot) GetSlot() uint64 {
	if m != nil {
		return m.Slot
	}
	return 0
}

func (m *PrecompEncryptSlot) GetEncryptedMessageKeys() []byte {
	if m != nil {
		return m.EncryptedMessageKeys
	}
	return nil
}

func (m *PrecompEncryptSlot) GetPartialMessageCypherText() []byte {
	if m != nil {
		return m.PartialMessageCypherText
	}
	return nil
}

// Message for batch of Precomp Encrypt Slots
type PrecompEncryptMessage struct {
	RoundID string                `protobuf:"bytes,1,opt,name=RoundID" json:"RoundID,omitempty"`
	Slots   []*PrecompEncryptSlot `protobuf:"bytes,2,rep,name=Slots" json:"Slots,omitempty"`
}

func (m *PrecompEncryptMessage) Reset()                    { *m = PrecompEncryptMessage{} }
func (m *PrecompEncryptMessage) String() string            { return proto.CompactTextString(m) }
func (*PrecompEncryptMessage) ProtoMessage()               {}
func (*PrecompEncryptMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *PrecompEncryptMessage) GetRoundID() string {
	if m != nil {
		return m.RoundID
	}
	return ""
}

func (m *PrecompEncryptMessage) GetSlots() []*PrecompEncryptSlot {
	if m != nil {
		return m.Slots
	}
	return nil
}

// Message for individual Precomp Permute Slot
type PrecompPermuteSlot struct {
	Slot                         uint64 `protobuf:"varint,1,opt,name=Slot" json:"Slot,omitempty"`
	EncryptedMessageKeys         []byte `protobuf:"bytes,2,opt,name=EncryptedMessageKeys,proto3" json:"EncryptedMessageKeys,omitempty"`
	EncryptedRecipientIDKeys     []byte `protobuf:"bytes,3,opt,name=EncryptedRecipientIDKeys,proto3" json:"EncryptedRecipientIDKeys,omitempty"`
	PartialMessageCypherText     []byte `protobuf:"bytes,4,opt,name=PartialMessageCypherText,proto3" json:"PartialMessageCypherText,omitempty"`
	PartialRecipientIDCypherText []byte `protobuf:"bytes,5,opt,name=PartialRecipientIDCypherText,proto3" json:"PartialRecipientIDCypherText,omitempty"`
}

func (m *PrecompPermuteSlot) Reset()                    { *m = PrecompPermuteSlot{} }
func (m *PrecompPermuteSlot) String() string            { return proto.CompactTextString(m) }
func (*PrecompPermuteSlot) ProtoMessage()               {}
func (*PrecompPermuteSlot) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *PrecompPermuteSlot) GetSlot() uint64 {
	if m != nil {
		return m.Slot
	}
	return 0
}

func (m *PrecompPermuteSlot) GetEncryptedMessageKeys() []byte {
	if m != nil {
		return m.EncryptedMessageKeys
	}
	return nil
}

func (m *PrecompPermuteSlot) GetEncryptedRecipientIDKeys() []byte {
	if m != nil {
		return m.EncryptedRecipientIDKeys
	}
	return nil
}

func (m *PrecompPermuteSlot) GetPartialMessageCypherText() []byte {
	if m != nil {
		return m.PartialMessageCypherText
	}
	return nil
}

func (m *PrecompPermuteSlot) GetPartialRecipientIDCypherText() []byte {
	if m != nil {
		return m.PartialRecipientIDCypherText
	}
	return nil
}

// Message for batch of Precomp Permute Slots
type PrecompPermuteMessage struct {
	RoundID string                `protobuf:"bytes,1,opt,name=RoundID" json:"RoundID,omitempty"`
	Slots   []*PrecompPermuteSlot `protobuf:"bytes,2,rep,name=Slots" json:"Slots,omitempty"`
}

func (m *PrecompPermuteMessage) Reset()                    { *m = PrecompPermuteMessage{} }
func (m *PrecompPermuteMessage) String() string            { return proto.CompactTextString(m) }
func (*PrecompPermuteMessage) ProtoMessage()               {}
func (*PrecompPermuteMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *PrecompPermuteMessage) GetRoundID() string {
	if m != nil {
		return m.RoundID
	}
	return ""
}

func (m *PrecompPermuteMessage) GetSlots() []*PrecompPermuteSlot {
	if m != nil {
		return m.Slots
	}
	return nil
}

// ErrorMessage encodes an error message
type ErrorMessage struct {
	Message string `protobuf:"bytes,1,opt,name=Message" json:"Message,omitempty"`
}

func (m *ErrorMessage) Reset()                    { *m = ErrorMessage{} }
func (m *ErrorMessage) String() string            { return proto.CompactTextString(m) }
func (*ErrorMessage) ProtoMessage()               {}
func (*ErrorMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ErrorMessage) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// ErrorAck returns the length of the received messages
type ErrorAck struct {
	MsgLen int32 `protobuf:"varint,1,opt,name=MsgLen" json:"MsgLen,omitempty"`
}

func (m *ErrorAck) Reset()                    { *m = ErrorAck{} }
func (m *ErrorAck) String() string            { return proto.CompactTextString(m) }
func (*ErrorAck) ProtoMessage()               {}
func (*ErrorAck) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *ErrorAck) GetMsgLen() int32 {
	if m != nil {
		return m.MsgLen
	}
	return 0
}

func init() {
	proto.RegisterType((*Ack)(nil), "mixmessages.Ack")
	proto.RegisterType((*Ping)(nil), "mixmessages.Ping")
	proto.RegisterType((*Pong)(nil), "mixmessages.Pong")
	proto.RegisterType((*PrecompDecryptSlot)(nil), "mixmessages.PrecompDecryptSlot")
	proto.RegisterType((*PrecompDecryptMessage)(nil), "mixmessages.PrecompDecryptMessage")
	proto.RegisterType((*PrecompEncryptSlot)(nil), "mixmessages.PrecompEncryptSlot")
	proto.RegisterType((*PrecompEncryptMessage)(nil), "mixmessages.PrecompEncryptMessage")
	proto.RegisterType((*PrecompPermuteSlot)(nil), "mixmessages.PrecompPermuteSlot")
	proto.RegisterType((*PrecompPermuteMessage)(nil), "mixmessages.PrecompPermuteMessage")
	proto.RegisterType((*ErrorMessage)(nil), "mixmessages.ErrorMessage")
	proto.RegisterType((*ErrorAck)(nil), "mixmessages.ErrorAck")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MixMessageService service

type MixMessageServiceClient interface {
	// Handles an error message
	NetworkError(ctx context.Context, in *ErrorMessage, opts ...grpc.CallOption) (*ErrorAck, error)
	// Handles an AskOnline message
	AskOnline(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Pong, error)
	// Handles Precomp Decrypt
	PrecompDecrypt(ctx context.Context, in *PrecompDecryptMessage, opts ...grpc.CallOption) (*Ack, error)
	// Handles Precomp Encrypt
	PrecompEncrypt(ctx context.Context, in *PrecompEncryptMessage, opts ...grpc.CallOption) (*Ack, error)
	// Handles Precomp Permute
	PrecompPermute(ctx context.Context, in *PrecompPermuteMessage, opts ...grpc.CallOption) (*Ack, error)
}

type mixMessageServiceClient struct {
	cc *grpc.ClientConn
}

func NewMixMessageServiceClient(cc *grpc.ClientConn) MixMessageServiceClient {
	return &mixMessageServiceClient{cc}
}

func (c *mixMessageServiceClient) NetworkError(ctx context.Context, in *ErrorMessage, opts ...grpc.CallOption) (*ErrorAck, error) {
	out := new(ErrorAck)
	err := grpc.Invoke(ctx, "/mixmessages.MixMessageService/NetworkError", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mixMessageServiceClient) AskOnline(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Pong, error) {
	out := new(Pong)
	err := grpc.Invoke(ctx, "/mixmessages.MixMessageService/AskOnline", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mixMessageServiceClient) PrecompDecrypt(ctx context.Context, in *PrecompDecryptMessage, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := grpc.Invoke(ctx, "/mixmessages.MixMessageService/PrecompDecrypt", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mixMessageServiceClient) PrecompEncrypt(ctx context.Context, in *PrecompEncryptMessage, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := grpc.Invoke(ctx, "/mixmessages.MixMessageService/PrecompEncrypt", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mixMessageServiceClient) PrecompPermute(ctx context.Context, in *PrecompPermuteMessage, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := grpc.Invoke(ctx, "/mixmessages.MixMessageService/PrecompPermute", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MixMessageService service

type MixMessageServiceServer interface {
	// Handles an error message
	NetworkError(context.Context, *ErrorMessage) (*ErrorAck, error)
	// Handles an AskOnline message
	AskOnline(context.Context, *Ping) (*Pong, error)
	// Handles Precomp Decrypt
	PrecompDecrypt(context.Context, *PrecompDecryptMessage) (*Ack, error)
	// Handles Precomp Encrypt
	PrecompEncrypt(context.Context, *PrecompEncryptMessage) (*Ack, error)
	// Handles Precomp Permute
	PrecompPermute(context.Context, *PrecompPermuteMessage) (*Ack, error)
}

func RegisterMixMessageServiceServer(s *grpc.Server, srv MixMessageServiceServer) {
	s.RegisterService(&_MixMessageService_serviceDesc, srv)
}

func _MixMessageService_NetworkError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ErrorMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MixMessageServiceServer).NetworkError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mixmessages.MixMessageService/NetworkError",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MixMessageServiceServer).NetworkError(ctx, req.(*ErrorMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _MixMessageService_AskOnline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ping)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MixMessageServiceServer).AskOnline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mixmessages.MixMessageService/AskOnline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MixMessageServiceServer).AskOnline(ctx, req.(*Ping))
	}
	return interceptor(ctx, in, info, handler)
}

func _MixMessageService_PrecompDecrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrecompDecryptMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MixMessageServiceServer).PrecompDecrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mixmessages.MixMessageService/PrecompDecrypt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MixMessageServiceServer).PrecompDecrypt(ctx, req.(*PrecompDecryptMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _MixMessageService_PrecompEncrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrecompEncryptMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MixMessageServiceServer).PrecompEncrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mixmessages.MixMessageService/PrecompEncrypt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MixMessageServiceServer).PrecompEncrypt(ctx, req.(*PrecompEncryptMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _MixMessageService_PrecompPermute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrecompPermuteMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MixMessageServiceServer).PrecompPermute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mixmessages.MixMessageService/PrecompPermute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MixMessageServiceServer).PrecompPermute(ctx, req.(*PrecompPermuteMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _MixMessageService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mixmessages.MixMessageService",
	HandlerType: (*MixMessageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NetworkError",
			Handler:    _MixMessageService_NetworkError_Handler,
		},
		{
			MethodName: "AskOnline",
			Handler:    _MixMessageService_AskOnline_Handler,
		},
		{
			MethodName: "PrecompDecrypt",
			Handler:    _MixMessageService_PrecompDecrypt_Handler,
		},
		{
			MethodName: "PrecompEncrypt",
			Handler:    _MixMessageService_PrecompEncrypt_Handler,
		},
		{
			MethodName: "PrecompPermute",
			Handler:    _MixMessageService_PrecompPermute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mixmessages.proto",
}

func init() { proto.RegisterFile("mixmessages.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 440 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x6e, 0x12, 0x3b, 0xd0, 0x69, 0x84, 0xc8, 0x88, 0x22, 0x53, 0x21, 0x51, 0xed, 0xc9, 0xa7,
	0x1e, 0x52, 0x71, 0xe1, 0x96, 0x92, 0x4a, 0x54, 0x10, 0xb0, 0x5c, 0x5e, 0x20, 0xb8, 0x23, 0x77,
	0xe5, 0x78, 0xd7, 0x5a, 0x6f, 0x21, 0x79, 0x93, 0x3e, 0x23, 0x4f, 0x81, 0xbc, 0x5e, 0xa7, 0x6b,
	0xe2, 0x84, 0x4b, 0x8e, 0x9c, 0xf6, 0x67, 0xe6, 0xfb, 0x66, 0xf7, 0xfb, 0x76, 0x07, 0xc6, 0x39,
	0x5f, 0xe5, 0x54, 0x96, 0x8b, 0x94, 0xca, 0x8b, 0x42, 0x49, 0x2d, 0xf1, 0xc4, 0xd9, 0x62, 0x3e,
	0x0c, 0xa6, 0x49, 0xc6, 0x86, 0xe0, 0x45, 0x5c, 0xa4, 0x66, 0x94, 0x22, 0x65, 0x8f, 0x7d, 0xc0,
	0x48, 0x51, 0x22, 0xf3, 0x62, 0x46, 0x89, 0x5a, 0x17, 0xfa, 0x76, 0x29, 0x35, 0x22, 0x78, 0xd5,
	0x18, 0xf4, 0xce, 0x7b, 0xa1, 0x17, 0x9b, 0x39, 0x4e, 0xe0, 0xd5, 0xb5, 0x30, 0x29, 0x74, 0x37,
	0xaf, 0x69, 0x3f, 0xd3, 0xba, 0x0c, 0xfa, 0xe7, 0xbd, 0x70, 0x14, 0x77, 0xc6, 0xf0, 0x03, 0x04,
	0x9b, 0xfd, 0x98, 0x12, 0x5e, 0x70, 0x12, 0xfa, 0x66, 0x66, 0x70, 0x03, 0x83, 0xdb, 0x19, 0xaf,
	0xb0, 0xd1, 0x42, 0x69, 0xbe, 0x58, 0x5a, 0xc6, 0x8f, 0xeb, 0xe2, 0x9e, 0xd4, 0x77, 0x5a, 0xe9,
	0xc0, 0xab, 0xb1, 0xbb, 0xe2, 0x78, 0x05, 0x6f, 0x6d, 0xcc, 0x61, 0x75, 0xf0, 0xbe, 0xc1, 0xef,
	0xcd, 0x61, 0xf7, 0x70, 0xda, 0x56, 0xc6, 0x96, 0xc1, 0x00, 0x9e, 0xc5, 0xf2, 0x41, 0xdc, 0xdd,
	0xcc, 0x8c, 0x3e, 0xc7, 0x71, 0xb3, 0xc4, 0xf7, 0xe0, 0x57, 0x52, 0x55, 0x9a, 0x0c, 0xc2, 0x93,
	0xc9, 0xbb, 0x0b, 0xd7, 0x94, 0x6d, 0x99, 0xe3, 0x3a, 0x9b, 0x3d, 0xf6, 0x36, 0x26, 0x58, 0x35,
	0x0e, 0x6d, 0xc2, 0x4e, 0x21, 0x07, 0xfb, 0x85, 0x74, 0x44, 0xb0, 0xd4, 0x87, 0x11, 0xc1, 0xb9,
	0xe6, 0x46, 0x84, 0xa7, 0x97, 0x18, 0x91, 0xca, 0x1f, 0x34, 0xfd, 0x7f, 0x89, 0x5b, 0x26, 0x58,
	0x65, 0x0e, 0x63, 0x82, 0x23, 0x73, 0x63, 0x42, 0x08, 0xa3, 0x6b, 0xa5, 0xa4, 0x72, 0x0a, 0xd8,
	0x69, 0x53, 0xc0, 0x2e, 0x19, 0x83, 0xe7, 0x26, 0x73, 0x9a, 0x64, 0xf8, 0x1a, 0x86, 0xf3, 0x32,
	0xfd, 0x42, 0xc2, 0x24, 0xf9, 0xb1, 0x5d, 0x4d, 0x7e, 0xf7, 0x61, 0x3c, 0xe7, 0x2b, 0x0b, 0xb9,
	0x25, 0xf5, 0x93, 0x27, 0x84, 0x57, 0x30, 0xfa, 0x4a, 0xfa, 0x97, 0x54, 0x99, 0x21, 0xc0, 0x37,
	0xad, 0xb3, 0xb9, 0xe5, 0xcf, 0x4e, 0xb7, 0x43, 0x55, 0x13, 0x3b, 0xc2, 0x4b, 0x38, 0x9e, 0x96,
	0xd9, 0x37, 0xb1, 0xe4, 0x82, 0x70, 0xdc, 0xbe, 0x1c, 0x17, 0xe9, 0xd9, 0x5f, 0x5b, 0x55, 0xa7,
	0x3b, 0xc2, 0x4f, 0xf0, 0xa2, 0xfd, 0x07, 0x91, 0xed, 0xf9, 0xa0, 0xcd, 0x19, 0x5e, 0xb6, 0x72,
	0xea, 0xf2, 0x4f, 0x4c, 0xf6, 0xcd, 0x74, 0x33, 0xb5, 0xbf, 0xcc, 0x3f, 0x98, 0xac, 0x1b, 0xdd,
	0x4c, 0x6d, 0xdf, 0xbb, 0x98, 0x7e, 0x0c, 0x4d, 0xd3, 0xbf, 0xfc, 0x13, 0x00, 0x00, 0xff, 0xff,
	0xfd, 0x2e, 0x87, 0x4b, 0x09, 0x06, 0x00, 0x00,
}
