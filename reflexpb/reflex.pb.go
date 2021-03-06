// Code generated by protoc-gen-go. DO NOT EDIT.
// source: reflex.proto

package reflexpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import duration "github.com/golang/protobuf/ptypes/duration"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

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

type StreamRequest struct {
	AfterInt             int64          `protobuf:"varint,1,opt,name=after_int,json=afterInt,proto3" json:"after_int,omitempty"`
	Options              *StreamOptions `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"`
	After                string         `protobuf:"bytes,3,opt,name=after,proto3" json:"after,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *StreamRequest) Reset()         { *m = StreamRequest{} }
func (m *StreamRequest) String() string { return proto.CompactTextString(m) }
func (*StreamRequest) ProtoMessage()    {}
func (*StreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_reflex_9b85d283972ad706, []int{0}
}
func (m *StreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamRequest.Unmarshal(m, b)
}
func (m *StreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamRequest.Marshal(b, m, deterministic)
}
func (dst *StreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamRequest.Merge(dst, src)
}
func (m *StreamRequest) XXX_Size() int {
	return xxx_messageInfo_StreamRequest.Size(m)
}
func (m *StreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamRequest proto.InternalMessageInfo

func (m *StreamRequest) GetAfterInt() int64 {
	if m != nil {
		return m.AfterInt
	}
	return 0
}

func (m *StreamRequest) GetOptions() *StreamOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *StreamRequest) GetAfter() string {
	if m != nil {
		return m.After
	}
	return ""
}

type Event struct {
	IdInt                int64                `protobuf:"varint,1,opt,name=id_int,json=idInt,proto3" json:"id_int,omitempty"`
	Type                 int32                `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	ForeignId            string               `protobuf:"bytes,5,opt,name=foreign_id,json=foreignId,proto3" json:"foreign_id,omitempty"`
	Id                   string               `protobuf:"bytes,6,opt,name=id,proto3" json:"id,omitempty"`
	Metadata             []byte               `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_reflex_9b85d283972ad706, []int{1}
}
func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (dst *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(dst, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetIdInt() int64 {
	if m != nil {
		return m.IdInt
	}
	return 0
}

func (m *Event) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Event) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Event) GetForeignId() string {
	if m != nil {
		return m.ForeignId
	}
	return ""
}

func (m *Event) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Event) GetMetadata() []byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type StreamOptions struct {
	Lag                  *duration.Duration `protobuf:"bytes,1,opt,name=lag,proto3" json:"lag,omitempty"`
	FromHead             bool               `protobuf:"varint,2,opt,name=fromHead,proto3" json:"fromHead,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *StreamOptions) Reset()         { *m = StreamOptions{} }
func (m *StreamOptions) String() string { return proto.CompactTextString(m) }
func (*StreamOptions) ProtoMessage()    {}
func (*StreamOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_reflex_9b85d283972ad706, []int{2}
}
func (m *StreamOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamOptions.Unmarshal(m, b)
}
func (m *StreamOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamOptions.Marshal(b, m, deterministic)
}
func (dst *StreamOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamOptions.Merge(dst, src)
}
func (m *StreamOptions) XXX_Size() int {
	return xxx_messageInfo_StreamOptions.Size(m)
}
func (m *StreamOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamOptions.DiscardUnknown(m)
}

var xxx_messageInfo_StreamOptions proto.InternalMessageInfo

func (m *StreamOptions) GetLag() *duration.Duration {
	if m != nil {
		return m.Lag
	}
	return nil
}

func (m *StreamOptions) GetFromHead() bool {
	if m != nil {
		return m.FromHead
	}
	return false
}

func init() {
	proto.RegisterType((*StreamRequest)(nil), "reflexpb.StreamRequest")
	proto.RegisterType((*Event)(nil), "reflexpb.Event")
	proto.RegisterType((*StreamOptions)(nil), "reflexpb.StreamOptions")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ReflexClient is the client API for Reflex service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReflexClient interface {
	Stream(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (Reflex_StreamClient, error)
}

type reflexClient struct {
	cc *grpc.ClientConn
}

func NewReflexClient(cc *grpc.ClientConn) ReflexClient {
	return &reflexClient{cc}
}

func (c *reflexClient) Stream(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (Reflex_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Reflex_serviceDesc.Streams[0], "/reflexpb.Reflex/stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &reflexStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Reflex_StreamClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type reflexStreamClient struct {
	grpc.ClientStream
}

func (x *reflexStreamClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ReflexServer is the server API for Reflex service.
type ReflexServer interface {
	Stream(*StreamRequest, Reflex_StreamServer) error
}

func RegisterReflexServer(s *grpc.Server, srv ReflexServer) {
	s.RegisterService(&_Reflex_serviceDesc, srv)
}

func _Reflex_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ReflexServer).Stream(m, &reflexStreamServer{stream})
}

type Reflex_StreamServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type reflexStreamServer struct {
	grpc.ServerStream
}

func (x *reflexStreamServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

var _Reflex_serviceDesc = grpc.ServiceDesc{
	ServiceName: "reflexpb.Reflex",
	HandlerType: (*ReflexServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "stream",
			Handler:       _Reflex_Stream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "reflex.proto",
}

func init() { proto.RegisterFile("reflex.proto", fileDescriptor_reflex_9b85d283972ad706) }

var fileDescriptor_reflex_9b85d283972ad706 = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x51, 0x4d, 0x4b, 0xc3, 0x40,
	0x14, 0x34, 0x9f, 0x4d, 0x5e, 0xab, 0x96, 0x45, 0x31, 0x46, 0xd4, 0x90, 0x53, 0x40, 0x48, 0xb5,
	0x82, 0x78, 0xf4, 0xa0, 0x60, 0x7b, 0x11, 0x56, 0xcf, 0x96, 0x2d, 0xbb, 0x09, 0x0b, 0x4d, 0x36,
	0xa6, 0x1b, 0xd1, 0x1f, 0xe7, 0x7f, 0x93, 0xec, 0x26, 0xad, 0xd8, 0x5b, 0x66, 0xde, 0xbc, 0x9d,
	0x37, 0x13, 0x18, 0xd5, 0x2c, 0x5b, 0xb1, 0xaf, 0xb4, 0xaa, 0x85, 0x14, 0xc8, 0xd3, 0xa8, 0x5a,
	0x86, 0x97, 0xb9, 0x10, 0xf9, 0x8a, 0x4d, 0x14, 0xbf, 0x6c, 0xb2, 0x89, 0xe4, 0x05, 0x5b, 0x4b,
	0x52, 0x54, 0x5a, 0x1a, 0x5e, 0xfc, 0x17, 0xd0, 0xa6, 0x26, 0x92, 0x8b, 0x52, 0xcf, 0xe3, 0x06,
	0xf6, 0x5f, 0x65, 0xcd, 0x48, 0x81, 0xd9, 0x47, 0xc3, 0xd6, 0x12, 0x9d, 0x81, 0x4f, 0x32, 0xc9,
	0xea, 0x05, 0x2f, 0x65, 0x60, 0x44, 0x46, 0x62, 0x61, 0x4f, 0x11, 0xb3, 0x52, 0xa2, 0x1b, 0x18,
	0x88, 0xaa, 0xdd, 0x5e, 0x07, 0x66, 0x64, 0x24, 0xc3, 0xe9, 0x49, 0xda, 0x9f, 0x92, 0xea, 0x67,
	0x5e, 0xf4, 0x18, 0xf7, 0x3a, 0x74, 0x04, 0x8e, 0x5a, 0x0f, 0xac, 0xc8, 0x48, 0x7c, 0xac, 0x41,
	0xfc, 0x63, 0x80, 0xf3, 0xf4, 0xc9, 0x4a, 0x89, 0x8e, 0xc1, 0xe5, 0xf4, 0x8f, 0x99, 0xc3, 0x69,
	0xeb, 0x84, 0xc0, 0x96, 0xdf, 0x15, 0x53, 0x5b, 0x0e, 0x56, 0xdf, 0xe8, 0x1e, 0xfc, 0x4d, 0xbc,
	0xc0, 0x56, 0xfe, 0x61, 0xaa, 0xf3, 0xa5, 0x7d, 0xbe, 0xf4, 0xad, 0x57, 0xe0, 0xad, 0x18, 0x9d,
	0x03, 0x64, 0xa2, 0x66, 0x3c, 0x2f, 0x17, 0x9c, 0x06, 0x8e, 0xba, 0xc4, 0xef, 0x98, 0x19, 0x45,
	0x07, 0x60, 0x72, 0x1a, 0xb8, 0x8a, 0x36, 0x39, 0x45, 0x21, 0x78, 0x05, 0x93, 0x84, 0x12, 0x49,
	0x82, 0x41, 0x64, 0x24, 0x23, 0xbc, 0xc1, 0x73, 0xdb, 0x33, 0xc7, 0x56, 0xfc, 0xde, 0xd7, 0xd6,
	0xe5, 0x45, 0x57, 0x60, 0xad, 0x48, 0xae, 0x32, 0x0c, 0xa7, 0xa7, 0x3b, 0x57, 0x3d, 0x76, 0xad,
	0xe3, 0x56, 0xd5, 0xbe, 0x9f, 0xd5, 0xa2, 0x78, 0x66, 0x84, 0xaa, 0x1e, 0x3d, 0xbc, 0xc1, 0x73,
	0xdb, 0xb3, 0xc6, 0xf6, 0xf4, 0x01, 0x5c, 0xac, 0x8a, 0x45, 0x77, 0xe0, 0x6a, 0x27, 0xb4, 0xd3,
	0x75, 0xf7, 0xcb, 0xc2, 0xc3, 0xed, 0x40, 0x75, 0x1a, 0xef, 0x5d, 0x1b, 0x4b, 0x57, 0x79, 0xdf,
	0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x90, 0x66, 0x3b, 0xa2, 0x3a, 0x02, 0x00, 0x00,
}
