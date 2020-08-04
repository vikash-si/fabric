// Code generated by protoc-gen-go. DO NOT EDIT.
// source: peer/snapshot.proto

package peer

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

// SnapshotRequest contains information for a generate/cancel snapshot request
type SnapshotRequest struct {
	ChannelId            string   `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	Height               uint64   `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SnapshotRequest) Reset()         { *m = SnapshotRequest{} }
func (m *SnapshotRequest) String() string { return proto.CompactTextString(m) }
func (*SnapshotRequest) ProtoMessage()    {}
func (*SnapshotRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d05a247df97d1516, []int{0}
}

func (m *SnapshotRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnapshotRequest.Unmarshal(m, b)
}
func (m *SnapshotRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnapshotRequest.Marshal(b, m, deterministic)
}
func (m *SnapshotRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnapshotRequest.Merge(m, src)
}
func (m *SnapshotRequest) XXX_Size() int {
	return xxx_messageInfo_SnapshotRequest.Size(m)
}
func (m *SnapshotRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SnapshotRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SnapshotRequest proto.InternalMessageInfo

func (m *SnapshotRequest) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *SnapshotRequest) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

// SnapshotQuery contains information for a query snapshot request
type SnapshotQuery struct {
	ChannelId            string   `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SnapshotQuery) Reset()         { *m = SnapshotQuery{} }
func (m *SnapshotQuery) String() string { return proto.CompactTextString(m) }
func (*SnapshotQuery) ProtoMessage()    {}
func (*SnapshotQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_d05a247df97d1516, []int{1}
}

func (m *SnapshotQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnapshotQuery.Unmarshal(m, b)
}
func (m *SnapshotQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnapshotQuery.Marshal(b, m, deterministic)
}
func (m *SnapshotQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnapshotQuery.Merge(m, src)
}
func (m *SnapshotQuery) XXX_Size() int {
	return xxx_messageInfo_SnapshotQuery.Size(m)
}
func (m *SnapshotQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_SnapshotQuery.DiscardUnknown(m)
}

var xxx_messageInfo_SnapshotQuery proto.InternalMessageInfo

func (m *SnapshotQuery) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

// SignedSnapshotRequest contains marshalled request bytes and signature
type SignedSnapshotRequest struct {
	// The bytes of SnapshotRequest or SnapshotQuery
	Request []byte `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	// Signaure over request bytes; this signature is to be verified against the creator identity
	Signature            []byte   `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignedSnapshotRequest) Reset()         { *m = SignedSnapshotRequest{} }
func (m *SignedSnapshotRequest) String() string { return proto.CompactTextString(m) }
func (*SignedSnapshotRequest) ProtoMessage()    {}
func (*SignedSnapshotRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d05a247df97d1516, []int{2}
}

func (m *SignedSnapshotRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignedSnapshotRequest.Unmarshal(m, b)
}
func (m *SignedSnapshotRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignedSnapshotRequest.Marshal(b, m, deterministic)
}
func (m *SignedSnapshotRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedSnapshotRequest.Merge(m, src)
}
func (m *SignedSnapshotRequest) XXX_Size() int {
	return xxx_messageInfo_SignedSnapshotRequest.Size(m)
}
func (m *SignedSnapshotRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedSnapshotRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignedSnapshotRequest proto.InternalMessageInfo

func (m *SignedSnapshotRequest) GetRequest() []byte {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *SignedSnapshotRequest) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

// QueryPendingSnapshotsResponse specifies the response payload of a query pending snapshots request
type QueryPendingSnapshotsResponse struct {
	Heights              []uint64 `protobuf:"varint,1,rep,packed,name=heights,proto3" json:"heights,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryPendingSnapshotsResponse) Reset()         { *m = QueryPendingSnapshotsResponse{} }
func (m *QueryPendingSnapshotsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryPendingSnapshotsResponse) ProtoMessage()    {}
func (*QueryPendingSnapshotsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d05a247df97d1516, []int{3}
}

func (m *QueryPendingSnapshotsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryPendingSnapshotsResponse.Unmarshal(m, b)
}
func (m *QueryPendingSnapshotsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryPendingSnapshotsResponse.Marshal(b, m, deterministic)
}
func (m *QueryPendingSnapshotsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryPendingSnapshotsResponse.Merge(m, src)
}
func (m *QueryPendingSnapshotsResponse) XXX_Size() int {
	return xxx_messageInfo_QueryPendingSnapshotsResponse.Size(m)
}
func (m *QueryPendingSnapshotsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryPendingSnapshotsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryPendingSnapshotsResponse proto.InternalMessageInfo

func (m *QueryPendingSnapshotsResponse) GetHeights() []uint64 {
	if m != nil {
		return m.Heights
	}
	return nil
}

func init() {
	proto.RegisterType((*SnapshotRequest)(nil), "protos.SnapshotRequest")
	proto.RegisterType((*SnapshotQuery)(nil), "protos.SnapshotQuery")
	proto.RegisterType((*SignedSnapshotRequest)(nil), "protos.SignedSnapshotRequest")
	proto.RegisterType((*QueryPendingSnapshotsResponse)(nil), "protos.QueryPendingSnapshotsResponse")
}

func init() { proto.RegisterFile("peer/snapshot.proto", fileDescriptor_d05a247df97d1516) }

var fileDescriptor_d05a247df97d1516 = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x4f, 0x4b, 0xeb, 0x40,
	0x14, 0xc5, 0x9b, 0xf7, 0x4a, 0x5f, 0x7b, 0x69, 0x79, 0x90, 0xc7, 0x2b, 0xa1, 0x5a, 0x28, 0x03,
	0x42, 0x17, 0x76, 0x02, 0xba, 0x72, 0xa9, 0x45, 0xd4, 0x95, 0x3a, 0x5d, 0x08, 0x6e, 0x24, 0x7f,
	0x6e, 0x27, 0x03, 0xe9, 0x4c, 0x9c, 0x99, 0x2c, 0xfa, 0x99, 0xfd, 0x12, 0x92, 0x4c, 0x07, 0x8b,
	0x88, 0x05, 0x57, 0xc9, 0xb9, 0x39, 0xf7, 0x77, 0x6f, 0x0e, 0x17, 0xfe, 0x55, 0x88, 0x3a, 0x36,
	0x32, 0xa9, 0x4c, 0xa1, 0x2c, 0xad, 0xb4, 0xb2, 0x2a, 0xec, 0xb5, 0x0f, 0x33, 0x39, 0xe2, 0x4a,
	0xf1, 0x12, 0xe3, 0x56, 0xa6, 0xf5, 0x3a, 0xc6, 0x4d, 0x65, 0xb7, 0xce, 0x44, 0x6e, 0xe1, 0xef,
	0x6a, 0xd7, 0xc6, 0xf0, 0xb5, 0x46, 0x63, 0xc3, 0x29, 0x40, 0x56, 0x24, 0x52, 0x62, 0xf9, 0x22,
	0xf2, 0x28, 0x98, 0x05, 0xf3, 0x01, 0x1b, 0xec, 0x2a, 0x77, 0x79, 0x38, 0x86, 0x5e, 0x81, 0x82,
	0x17, 0x36, 0xfa, 0x35, 0x0b, 0xe6, 0x5d, 0xb6, 0x53, 0x84, 0xc2, 0xc8, 0x93, 0x1e, 0x6b, 0xd4,
	0xdb, 0x03, 0x1c, 0x72, 0x0f, 0xff, 0x57, 0x82, 0x4b, 0xcc, 0x3f, 0xcf, 0x8f, 0xe0, 0x8f, 0x76,
	0xaf, 0x6d, 0xd3, 0x90, 0x79, 0x19, 0x1e, 0xc3, 0xc0, 0x08, 0x2e, 0x13, 0x5b, 0x6b, 0x6c, 0xa7,
	0x0f, 0xd9, 0x47, 0x81, 0x5c, 0xc0, 0xb4, 0x1d, 0xfc, 0x80, 0x32, 0x17, 0x92, 0x7b, 0xac, 0x61,
	0x68, 0x2a, 0x25, 0x0d, 0x36, 0x60, 0xb7, 0xab, 0x89, 0x82, 0xd9, 0xef, 0x79, 0x97, 0x79, 0x79,
	0xf6, 0x16, 0x40, 0xdf, 0xfb, 0xc3, 0x25, 0xf4, 0x6f, 0x50, 0xa2, 0x4e, 0x2c, 0x86, 0x53, 0x17,
	0x93, 0xa1, 0x5f, 0xae, 0x3a, 0x19, 0x53, 0x97, 0x2d, 0xf5, 0xd9, 0xd2, 0xeb, 0x26, 0x5b, 0xd2,
	0x09, 0x2f, 0xa1, 0xb7, 0x4c, 0x64, 0x86, 0xe5, 0xcf, 0x11, 0x4f, 0x30, 0xda, 0xff, 0x1f, 0x73,
	0x88, 0x74, 0xe2, 0x3f, 0x7f, 0x9b, 0x02, 0xe9, 0x5c, 0x31, 0x20, 0x4a, 0x73, 0x5a, 0x6c, 0x2b,
	0xd4, 0x25, 0xe6, 0x1c, 0x35, 0x5d, 0x27, 0xa9, 0x16, 0x99, 0x07, 0x34, 0xd7, 0xf4, 0x7c, 0xca,
	0x85, 0x2d, 0xea, 0x94, 0x66, 0x6a, 0x13, 0xef, 0x59, 0x63, 0x67, 0x5d, 0x38, 0xeb, 0x82, 0xab,
	0xb8, 0x71, 0xa7, 0xee, 0xd8, 0xce, 0xdf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x69, 0x92, 0xf8, 0x2c,
	0x8a, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SnapshotClient is the client API for Snapshot service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SnapshotClient interface {
	// Generate a snapshot reqeust. SignedSnapshotRequest contains marshalled bytes for SnaphostRequest
	Generate(ctx context.Context, in *SignedSnapshotRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Cancel a snapshot reqeust. SignedSnapshotRequest contains marshalled bytes for SnaphostRequest
	Cancel(ctx context.Context, in *SignedSnapshotRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Query pending snapshots query. SignedSnapshotRequest contains marshalled bytes for SnaphostQuery
	QueryPendings(ctx context.Context, in *SignedSnapshotRequest, opts ...grpc.CallOption) (*QueryPendingSnapshotsResponse, error)
}

type snapshotClient struct {
	cc *grpc.ClientConn
}

func NewSnapshotClient(cc *grpc.ClientConn) SnapshotClient {
	return &snapshotClient{cc}
}

func (c *snapshotClient) Generate(ctx context.Context, in *SignedSnapshotRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/protos.Snapshot/Generate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snapshotClient) Cancel(ctx context.Context, in *SignedSnapshotRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/protos.Snapshot/Cancel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snapshotClient) QueryPendings(ctx context.Context, in *SignedSnapshotRequest, opts ...grpc.CallOption) (*QueryPendingSnapshotsResponse, error) {
	out := new(QueryPendingSnapshotsResponse)
	err := c.cc.Invoke(ctx, "/protos.Snapshot/QueryPendings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SnapshotServer is the server API for Snapshot service.
type SnapshotServer interface {
	// Generate a snapshot reqeust. SignedSnapshotRequest contains marshalled bytes for SnaphostRequest
	Generate(context.Context, *SignedSnapshotRequest) (*empty.Empty, error)
	// Cancel a snapshot reqeust. SignedSnapshotRequest contains marshalled bytes for SnaphostRequest
	Cancel(context.Context, *SignedSnapshotRequest) (*empty.Empty, error)
	// Query pending snapshots query. SignedSnapshotRequest contains marshalled bytes for SnaphostQuery
	QueryPendings(context.Context, *SignedSnapshotRequest) (*QueryPendingSnapshotsResponse, error)
}

// UnimplementedSnapshotServer can be embedded to have forward compatible implementations.
type UnimplementedSnapshotServer struct {
}

func (*UnimplementedSnapshotServer) Generate(ctx context.Context, req *SignedSnapshotRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Generate not implemented")
}
func (*UnimplementedSnapshotServer) Cancel(ctx context.Context, req *SignedSnapshotRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cancel not implemented")
}
func (*UnimplementedSnapshotServer) QueryPendings(ctx context.Context, req *SignedSnapshotRequest) (*QueryPendingSnapshotsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryPendings not implemented")
}

func RegisterSnapshotServer(s *grpc.Server, srv SnapshotServer) {
	s.RegisterService(&_Snapshot_serviceDesc, srv)
}

func _Snapshot_Generate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignedSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotServer).Generate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Snapshot/Generate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotServer).Generate(ctx, req.(*SignedSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Snapshot_Cancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignedSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotServer).Cancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Snapshot/Cancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotServer).Cancel(ctx, req.(*SignedSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Snapshot_QueryPendings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignedSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotServer).QueryPendings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Snapshot/QueryPendings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotServer).QueryPendings(ctx, req.(*SignedSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Snapshot_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Snapshot",
	HandlerType: (*SnapshotServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Generate",
			Handler:    _Snapshot_Generate_Handler,
		},
		{
			MethodName: "Cancel",
			Handler:    _Snapshot_Cancel_Handler,
		},
		{
			MethodName: "QueryPendings",
			Handler:    _Snapshot_QueryPendings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "peer/snapshot.proto",
}
