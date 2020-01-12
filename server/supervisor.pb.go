// Code generated by protoc-gen-go. DO NOT EDIT.
// source: supervisor.proto

package supervisor

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

type Node_NodeType int32

const (
	Node_Leader Node_NodeType = 0
	Node_Member Node_NodeType = 1
)

var Node_NodeType_name = map[int32]string{
	0: "Leader",
	1: "Member",
}

var Node_NodeType_value = map[string]int32{
	"Leader": 0,
	"Member": 1,
}

func (x Node_NodeType) String() string {
	return proto.EnumName(Node_NodeType_name, int32(x))
}

func (Node_NodeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b8b9452d77b1c7d2, []int{0, 0}
}

type Node struct {
	Id                   string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type                 Node_NodeType `protobuf:"varint,2,opt,name=type,proto3,enum=supervisor.Node_NodeType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8b9452d77b1c7d2, []int{0}
}

func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Node) GetType() Node_NodeType {
	if m != nil {
		return m.Type
	}
	return Node_Leader
}

type RegisterNodeRequest struct {
	Node                 *Node    `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterNodeRequest) Reset()         { *m = RegisterNodeRequest{} }
func (m *RegisterNodeRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterNodeRequest) ProtoMessage()    {}
func (*RegisterNodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8b9452d77b1c7d2, []int{1}
}

func (m *RegisterNodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterNodeRequest.Unmarshal(m, b)
}
func (m *RegisterNodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterNodeRequest.Marshal(b, m, deterministic)
}
func (m *RegisterNodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterNodeRequest.Merge(m, src)
}
func (m *RegisterNodeRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterNodeRequest.Size(m)
}
func (m *RegisterNodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterNodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterNodeRequest proto.InternalMessageInfo

func (m *RegisterNodeRequest) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

type RegisterNodeResponse struct {
	Result               bool     `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterNodeResponse) Reset()         { *m = RegisterNodeResponse{} }
func (m *RegisterNodeResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterNodeResponse) ProtoMessage()    {}
func (*RegisterNodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8b9452d77b1c7d2, []int{2}
}

func (m *RegisterNodeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterNodeResponse.Unmarshal(m, b)
}
func (m *RegisterNodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterNodeResponse.Marshal(b, m, deterministic)
}
func (m *RegisterNodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterNodeResponse.Merge(m, src)
}
func (m *RegisterNodeResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterNodeResponse.Size(m)
}
func (m *RegisterNodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterNodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterNodeResponse proto.InternalMessageInfo

func (m *RegisterNodeResponse) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

type LeaderStatusRequest struct {
	// Node id
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeaderStatusRequest) Reset()         { *m = LeaderStatusRequest{} }
func (m *LeaderStatusRequest) String() string { return proto.CompactTextString(m) }
func (*LeaderStatusRequest) ProtoMessage()    {}
func (*LeaderStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8b9452d77b1c7d2, []int{3}
}

func (m *LeaderStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LeaderStatusRequest.Unmarshal(m, b)
}
func (m *LeaderStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LeaderStatusRequest.Marshal(b, m, deterministic)
}
func (m *LeaderStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaderStatusRequest.Merge(m, src)
}
func (m *LeaderStatusRequest) XXX_Size() int {
	return xxx_messageInfo_LeaderStatusRequest.Size(m)
}
func (m *LeaderStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaderStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LeaderStatusRequest proto.InternalMessageInfo

func (m *LeaderStatusRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type LeaderStatusResponse struct {
	// Node id
	DependentID          string   `protobuf:"bytes,1,opt,name=dependentID,proto3" json:"dependentID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeaderStatusResponse) Reset()         { *m = LeaderStatusResponse{} }
func (m *LeaderStatusResponse) String() string { return proto.CompactTextString(m) }
func (*LeaderStatusResponse) ProtoMessage()    {}
func (*LeaderStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8b9452d77b1c7d2, []int{4}
}

func (m *LeaderStatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LeaderStatusResponse.Unmarshal(m, b)
}
func (m *LeaderStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LeaderStatusResponse.Marshal(b, m, deterministic)
}
func (m *LeaderStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaderStatusResponse.Merge(m, src)
}
func (m *LeaderStatusResponse) XXX_Size() int {
	return xxx_messageInfo_LeaderStatusResponse.Size(m)
}
func (m *LeaderStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaderStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LeaderStatusResponse proto.InternalMessageInfo

func (m *LeaderStatusResponse) GetDependentID() string {
	if m != nil {
		return m.DependentID
	}
	return ""
}

type MemberStatusRequest struct {
	// Node id
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MemberStatusRequest) Reset()         { *m = MemberStatusRequest{} }
func (m *MemberStatusRequest) String() string { return proto.CompactTextString(m) }
func (*MemberStatusRequest) ProtoMessage()    {}
func (*MemberStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8b9452d77b1c7d2, []int{5}
}

func (m *MemberStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemberStatusRequest.Unmarshal(m, b)
}
func (m *MemberStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemberStatusRequest.Marshal(b, m, deterministic)
}
func (m *MemberStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberStatusRequest.Merge(m, src)
}
func (m *MemberStatusRequest) XXX_Size() int {
	return xxx_messageInfo_MemberStatusRequest.Size(m)
}
func (m *MemberStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MemberStatusRequest proto.InternalMessageInfo

func (m *MemberStatusRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type MemberStatusResponse struct {
	// Node id
	DependentID          string   `protobuf:"bytes,1,opt,name=dependentID,proto3" json:"dependentID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MemberStatusResponse) Reset()         { *m = MemberStatusResponse{} }
func (m *MemberStatusResponse) String() string { return proto.CompactTextString(m) }
func (*MemberStatusResponse) ProtoMessage()    {}
func (*MemberStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8b9452d77b1c7d2, []int{6}
}

func (m *MemberStatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemberStatusResponse.Unmarshal(m, b)
}
func (m *MemberStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemberStatusResponse.Marshal(b, m, deterministic)
}
func (m *MemberStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberStatusResponse.Merge(m, src)
}
func (m *MemberStatusResponse) XXX_Size() int {
	return xxx_messageInfo_MemberStatusResponse.Size(m)
}
func (m *MemberStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MemberStatusResponse proto.InternalMessageInfo

func (m *MemberStatusResponse) GetDependentID() string {
	if m != nil {
		return m.DependentID
	}
	return ""
}

func init() {
	proto.RegisterEnum("supervisor.Node_NodeType", Node_NodeType_name, Node_NodeType_value)
	proto.RegisterType((*Node)(nil), "supervisor.Node")
	proto.RegisterType((*RegisterNodeRequest)(nil), "supervisor.RegisterNodeRequest")
	proto.RegisterType((*RegisterNodeResponse)(nil), "supervisor.RegisterNodeResponse")
	proto.RegisterType((*LeaderStatusRequest)(nil), "supervisor.LeaderStatusRequest")
	proto.RegisterType((*LeaderStatusResponse)(nil), "supervisor.LeaderStatusResponse")
	proto.RegisterType((*MemberStatusRequest)(nil), "supervisor.MemberStatusRequest")
	proto.RegisterType((*MemberStatusResponse)(nil), "supervisor.MemberStatusResponse")
}

func init() { proto.RegisterFile("supervisor.proto", fileDescriptor_b8b9452d77b1c7d2) }

var fileDescriptor_b8b9452d77b1c7d2 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0x6d, 0x19, 0xa5, 0x7b, 0x85, 0x51, 0xb2, 0x21, 0xd3, 0x8b, 0x25, 0x28, 0xec, 0x62,
	0x91, 0x7a, 0x11, 0xbc, 0x7a, 0x11, 0x9c, 0x48, 0x26, 0x78, 0xee, 0xcc, 0x43, 0x0b, 0xda, 0xc4,
	0x24, 0x15, 0x76, 0xf3, 0x4f, 0x97, 0xa6, 0x29, 0x36, 0xa3, 0xa2, 0x97, 0xd2, 0xbe, 0x7e, 0xef,
	0xfb, 0xf2, 0xfd, 0x08, 0xa4, 0xba, 0x91, 0xa8, 0x3e, 0x2b, 0x2d, 0x54, 0x2e, 0x95, 0x30, 0x82,
	0xc0, 0xcf, 0x84, 0x56, 0x30, 0xb9, 0x17, 0x1c, 0xc9, 0x0c, 0xc2, 0x8a, 0x2f, 0x83, 0x2c, 0x58,
	0x4d, 0x59, 0x58, 0x71, 0x72, 0x0e, 0x13, 0xb3, 0x93, 0xb8, 0x0c, 0xb3, 0x60, 0x35, 0x2b, 0x8e,
	0xf2, 0x81, 0x49, 0xab, 0xb7, 0x8f, 0xc7, 0x9d, 0x44, 0x66, 0x65, 0x94, 0x42, 0xdc, 0x4f, 0x08,
	0x40, 0x74, 0x87, 0x25, 0x47, 0x95, 0x1e, 0xb4, 0xef, 0x6b, 0x7c, 0xdf, 0xa2, 0x4a, 0x03, 0x7a,
	0x0d, 0x73, 0x86, 0x2f, 0x95, 0x36, 0xa8, 0x5a, 0x2d, 0xc3, 0x8f, 0x06, 0xb5, 0x21, 0xa7, 0x30,
	0xa9, 0x05, 0x47, 0x9b, 0x9d, 0x14, 0xe9, 0x7e, 0x12, 0xb3, 0x7f, 0x69, 0x0e, 0x0b, 0x7f, 0x59,
	0x4b, 0x51, 0x6b, 0x24, 0x87, 0x10, 0x29, 0xd4, 0xcd, 0x9b, 0xb1, 0xfb, 0x31, 0x73, 0x5f, 0xf4,
	0x0c, 0xe6, 0xdd, 0x21, 0x36, 0xa6, 0x34, 0x8d, 0xee, 0xc3, 0xf6, 0x6a, 0xd2, 0x2b, 0x58, 0xf8,
	0x32, 0x67, 0x9b, 0x41, 0xc2, 0x51, 0x62, 0xcd, 0xb1, 0x36, 0xb7, 0x37, 0x6e, 0x61, 0x38, 0x6a,
	0x03, 0xba, 0x66, 0x7f, 0x06, 0xf8, 0xb2, 0xff, 0x06, 0x14, 0x5f, 0x21, 0x4c, 0x37, 0x8e, 0x05,
	0x92, 0x35, 0xc4, 0x7d, 0x7f, 0x72, 0x32, 0x64, 0x34, 0x82, 0xf4, 0x38, 0xfb, 0x5d, 0xe0, 0xe2,
	0x1f, 0x20, 0x79, 0x2a, 0xcd, 0xf3, 0x6b, 0x57, 0xde, 0x77, 0x1c, 0xe1, 0xe6, 0x3b, 0x8e, 0x12,
	0x63, 0xce, 0xb1, 0x6b, 0xeb, 0x3b, 0x8e, 0x80, 0xf2, 0x1d, 0xc7, 0x10, 0x5d, 0x04, 0xdb, 0xc8,
	0xde, 0xd7, 0xcb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x00, 0xc7, 0x2a, 0xc3, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SuperviseClient is the client API for Supervise service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SuperviseClient interface {
	Register(ctx context.Context, in *RegisterNodeRequest, opts ...grpc.CallOption) (*RegisterNodeResponse, error)
	WatchLeader(ctx context.Context, in *LeaderStatusRequest, opts ...grpc.CallOption) (*LeaderStatusResponse, error)
	WatchMember(ctx context.Context, in *MemberStatusRequest, opts ...grpc.CallOption) (Supervise_WatchMemberClient, error)
}

type superviseClient struct {
	cc *grpc.ClientConn
}

func NewSuperviseClient(cc *grpc.ClientConn) SuperviseClient {
	return &superviseClient{cc}
}

func (c *superviseClient) Register(ctx context.Context, in *RegisterNodeRequest, opts ...grpc.CallOption) (*RegisterNodeResponse, error) {
	out := new(RegisterNodeResponse)
	err := c.cc.Invoke(ctx, "/supervisor.Supervise/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *superviseClient) WatchLeader(ctx context.Context, in *LeaderStatusRequest, opts ...grpc.CallOption) (*LeaderStatusResponse, error) {
	out := new(LeaderStatusResponse)
	err := c.cc.Invoke(ctx, "/supervisor.Supervise/WatchLeader", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *superviseClient) WatchMember(ctx context.Context, in *MemberStatusRequest, opts ...grpc.CallOption) (Supervise_WatchMemberClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Supervise_serviceDesc.Streams[0], "/supervisor.Supervise/WatchMember", opts...)
	if err != nil {
		return nil, err
	}
	x := &superviseWatchMemberClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Supervise_WatchMemberClient interface {
	Recv() (*MemberStatusResponse, error)
	grpc.ClientStream
}

type superviseWatchMemberClient struct {
	grpc.ClientStream
}

func (x *superviseWatchMemberClient) Recv() (*MemberStatusResponse, error) {
	m := new(MemberStatusResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SuperviseServer is the server API for Supervise service.
type SuperviseServer interface {
	Register(context.Context, *RegisterNodeRequest) (*RegisterNodeResponse, error)
	WatchLeader(context.Context, *LeaderStatusRequest) (*LeaderStatusResponse, error)
	WatchMember(*MemberStatusRequest, Supervise_WatchMemberServer) error
}

// UnimplementedSuperviseServer can be embedded to have forward compatible implementations.
type UnimplementedSuperviseServer struct {
}

func (*UnimplementedSuperviseServer) Register(ctx context.Context, req *RegisterNodeRequest) (*RegisterNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedSuperviseServer) WatchLeader(ctx context.Context, req *LeaderStatusRequest) (*LeaderStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WatchLeader not implemented")
}
func (*UnimplementedSuperviseServer) WatchMember(req *MemberStatusRequest, srv Supervise_WatchMemberServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchMember not implemented")
}

func RegisterSuperviseServer(s *grpc.Server, srv SuperviseServer) {
	s.RegisterService(&_Supervise_serviceDesc, srv)
}

func _Supervise_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SuperviseServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supervisor.Supervise/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SuperviseServer).Register(ctx, req.(*RegisterNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Supervise_WatchLeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaderStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SuperviseServer).WatchLeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supervisor.Supervise/WatchLeader",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SuperviseServer).WatchLeader(ctx, req.(*LeaderStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Supervise_WatchMember_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MemberStatusRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SuperviseServer).WatchMember(m, &superviseWatchMemberServer{stream})
}

type Supervise_WatchMemberServer interface {
	Send(*MemberStatusResponse) error
	grpc.ServerStream
}

type superviseWatchMemberServer struct {
	grpc.ServerStream
}

func (x *superviseWatchMemberServer) Send(m *MemberStatusResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Supervise_serviceDesc = grpc.ServiceDesc{
	ServiceName: "supervisor.Supervise",
	HandlerType: (*SuperviseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Supervise_Register_Handler,
		},
		{
			MethodName: "WatchLeader",
			Handler:    _Supervise_WatchLeader_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchMember",
			Handler:       _Supervise_WatchMember_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "supervisor.proto",
}
