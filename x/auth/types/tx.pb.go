// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/auth/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	any "github.com/cosmos/gogoproto/types/any"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgUpdateParams is the Msg/UpdateParams request type.
type MsgUpdateParams struct {
	// authority is the address that controls the module (defaults to x/gov unless overwritten).
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// params defines the x/auth parameters to update.
	//
	// NOTE: All parameters must be supplied.
	Params Params `protobuf:"bytes,2,opt,name=params,proto3" json:"params"`
}

func (m *MsgUpdateParams) Reset()         { *m = MsgUpdateParams{} }
func (m *MsgUpdateParams) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateParams) ProtoMessage()    {}
func (*MsgUpdateParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2d62bd9c4c212e5, []int{0}
}
func (m *MsgUpdateParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateParams.Merge(m, src)
}
func (m *MsgUpdateParams) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateParams) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateParams.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateParams proto.InternalMessageInfo

func (m *MsgUpdateParams) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *MsgUpdateParams) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// MsgUpdateParamsResponse defines the response structure for executing a
// MsgUpdateParams message.
type MsgUpdateParamsResponse struct {
}

func (m *MsgUpdateParamsResponse) Reset()         { *m = MsgUpdateParamsResponse{} }
func (m *MsgUpdateParamsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateParamsResponse) ProtoMessage()    {}
func (*MsgUpdateParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2d62bd9c4c212e5, []int{1}
}
func (m *MsgUpdateParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateParamsResponse.Merge(m, src)
}
func (m *MsgUpdateParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateParamsResponse proto.InternalMessageInfo

// MsgNonAtomicExec defines the Msg/NonAtomicExec request type.
type MsgNonAtomicExec struct {
	Signer string     `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty"`
	Msgs   []*any.Any `protobuf:"bytes,2,rep,name=msgs,proto3" json:"msgs,omitempty"`
}

func (m *MsgNonAtomicExec) Reset()         { *m = MsgNonAtomicExec{} }
func (m *MsgNonAtomicExec) String() string { return proto.CompactTextString(m) }
func (*MsgNonAtomicExec) ProtoMessage()    {}
func (*MsgNonAtomicExec) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2d62bd9c4c212e5, []int{2}
}
func (m *MsgNonAtomicExec) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgNonAtomicExec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgNonAtomicExec.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgNonAtomicExec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgNonAtomicExec.Merge(m, src)
}
func (m *MsgNonAtomicExec) XXX_Size() int {
	return m.Size()
}
func (m *MsgNonAtomicExec) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgNonAtomicExec.DiscardUnknown(m)
}

var xxx_messageInfo_MsgNonAtomicExec proto.InternalMessageInfo

func (m *MsgNonAtomicExec) GetSigner() string {
	if m != nil {
		return m.Signer
	}
	return ""
}

func (m *MsgNonAtomicExec) GetMsgs() []*any.Any {
	if m != nil {
		return m.Msgs
	}
	return nil
}

// NonAtomicExecResult defines the response structure for executing a
// MsgNonAtomicExec.
type NonAtomicExecResult struct {
	Error string   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Resp  *any.Any `protobuf:"bytes,2,opt,name=resp,proto3" json:"resp,omitempty"`
}

func (m *NonAtomicExecResult) Reset()         { *m = NonAtomicExecResult{} }
func (m *NonAtomicExecResult) String() string { return proto.CompactTextString(m) }
func (*NonAtomicExecResult) ProtoMessage()    {}
func (*NonAtomicExecResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2d62bd9c4c212e5, []int{3}
}
func (m *NonAtomicExecResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NonAtomicExecResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NonAtomicExecResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NonAtomicExecResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NonAtomicExecResult.Merge(m, src)
}
func (m *NonAtomicExecResult) XXX_Size() int {
	return m.Size()
}
func (m *NonAtomicExecResult) XXX_DiscardUnknown() {
	xxx_messageInfo_NonAtomicExecResult.DiscardUnknown(m)
}

var xxx_messageInfo_NonAtomicExecResult proto.InternalMessageInfo

func (m *NonAtomicExecResult) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *NonAtomicExecResult) GetResp() *any.Any {
	if m != nil {
		return m.Resp
	}
	return nil
}

// MsgNonAtomicExecResponse defines the response of MsgNonAtomicExec.
type MsgNonAtomicExecResponse struct {
	Results []*NonAtomicExecResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (m *MsgNonAtomicExecResponse) Reset()         { *m = MsgNonAtomicExecResponse{} }
func (m *MsgNonAtomicExecResponse) String() string { return proto.CompactTextString(m) }
func (*MsgNonAtomicExecResponse) ProtoMessage()    {}
func (*MsgNonAtomicExecResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2d62bd9c4c212e5, []int{4}
}
func (m *MsgNonAtomicExecResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgNonAtomicExecResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgNonAtomicExecResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgNonAtomicExecResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgNonAtomicExecResponse.Merge(m, src)
}
func (m *MsgNonAtomicExecResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgNonAtomicExecResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgNonAtomicExecResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgNonAtomicExecResponse proto.InternalMessageInfo

func (m *MsgNonAtomicExecResponse) GetResults() []*NonAtomicExecResult {
	if m != nil {
		return m.Results
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgUpdateParams)(nil), "cosmos.auth.v1beta1.MsgUpdateParams")
	proto.RegisterType((*MsgUpdateParamsResponse)(nil), "cosmos.auth.v1beta1.MsgUpdateParamsResponse")
	proto.RegisterType((*MsgNonAtomicExec)(nil), "cosmos.auth.v1beta1.MsgNonAtomicExec")
	proto.RegisterType((*NonAtomicExecResult)(nil), "cosmos.auth.v1beta1.NonAtomicExecResult")
	proto.RegisterType((*MsgNonAtomicExecResponse)(nil), "cosmos.auth.v1beta1.MsgNonAtomicExecResponse")
}

func init() { proto.RegisterFile("cosmos/auth/v1beta1/tx.proto", fileDescriptor_c2d62bd9c4c212e5) }

var fileDescriptor_c2d62bd9c4c212e5 = []byte{
	// 542 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x41, 0x8f, 0xd2, 0x40,
	0x14, 0x66, 0xdc, 0x5d, 0x0c, 0x83, 0x66, 0xb5, 0x90, 0x2c, 0xb2, 0xa6, 0x62, 0xa3, 0x09, 0x21,
	0x32, 0x5d, 0x58, 0xa3, 0x09, 0x07, 0x13, 0x48, 0x36, 0x9e, 0x30, 0xa6, 0x66, 0x2f, 0x1e, 0x34,
	0x05, 0xc6, 0xb1, 0xd9, 0x6d, 0xa7, 0xe9, 0x1b, 0x56, 0xb8, 0x19, 0x8f, 0x9e, 0x3c, 0xfa, 0x13,
	0x3c, 0x72, 0xe0, 0x47, 0x6c, 0xf6, 0xb4, 0xe1, 0xe4, 0xc9, 0x18, 0x38, 0x70, 0xf1, 0x47, 0x98,
	0xce, 0x4c, 0x25, 0x60, 0x8d, 0x7b, 0x69, 0x3b, 0xf3, 0x7d, 0xef, 0xbd, 0xef, 0xbd, 0xef, 0x15,
	0xdf, 0xed, 0x73, 0xf0, 0x39, 0xd8, 0xee, 0x50, 0xbc, 0xb7, 0xcf, 0x1a, 0x3d, 0x2a, 0xdc, 0x86,
	0x2d, 0x46, 0x24, 0x8c, 0xb8, 0xe0, 0x46, 0x41, 0xa1, 0x24, 0x46, 0x89, 0x46, 0xcb, 0x45, 0xc6,
	0x19, 0x97, 0xb8, 0x1d, 0x7f, 0x29, 0x6a, 0xf9, 0x0e, 0xe3, 0x9c, 0x9d, 0x52, 0x5b, 0x9e, 0x7a,
	0xc3, 0x77, 0xb6, 0x1b, 0x8c, 0x13, 0x48, 0x65, 0x79, 0xab, 0x62, 0x74, 0x4a, 0x05, 0xed, 0xe9,
	0xf2, 0x3e, 0x30, 0xfb, 0xac, 0x11, 0xbf, 0x34, 0x70, 0xdb, 0xf5, 0xbd, 0x80, 0xdb, 0xf2, 0xa9,
	0xaf, 0xcc, 0x34, 0xa9, 0x52, 0x99, 0xc4, 0xad, 0x19, 0xc2, 0xbb, 0x5d, 0x60, 0xc7, 0xe1, 0xc0,
	0x15, 0xf4, 0xa5, 0x1b, 0xb9, 0x3e, 0x18, 0x4f, 0x70, 0x2e, 0x66, 0xf0, 0xc8, 0x13, 0xe3, 0x12,
	0xaa, 0xa0, 0x6a, 0xae, 0x53, 0x9a, 0x4d, 0xeb, 0x45, 0x2d, 0xa2, 0x3d, 0x18, 0x44, 0x14, 0xe0,
	0x95, 0x88, 0xbc, 0x80, 0x39, 0x2b, 0xaa, 0xf1, 0x0c, 0x67, 0x43, 0x99, 0xa1, 0x74, 0xad, 0x82,
	0xaa, 0xf9, 0xe6, 0x3e, 0x49, 0x99, 0x04, 0x51, 0x45, 0x3a, 0xb9, 0xf3, 0x1f, 0xf7, 0x32, 0xdf,
	0x96, 0x93, 0x1a, 0x72, 0x74, 0x54, 0xeb, 0xf9, 0x6c, 0x5a, 0xdf, 0x55, 0x21, 0x75, 0x18, 0x9c,
	0x54, 0x0e, 0xc8, 0xe3, 0xa7, 0x9f, 0x96, 0x93, 0xda, 0xaa, 0xc4, 0xe7, 0xe5, 0xa4, 0x76, 0x7f,
	0xc5, 0xb0, 0x47, 0xaa, 0xaf, 0x8d, 0x06, 0x2c, 0x82, 0xf7, 0x36, 0xae, 0x1c, 0x0a, 0x21, 0x0f,
	0x80, 0xb6, 0x0a, 0x29, 0x35, 0xac, 0xaf, 0x08, 0xdf, 0xea, 0x02, 0x7b, 0xc1, 0x83, 0xb6, 0xe0,
	0xbe, 0xd7, 0x3f, 0x1a, 0xd1, 0xbe, 0x71, 0x80, 0xb3, 0xe0, 0xb1, 0x80, 0x46, 0xff, 0x1d, 0x81,
	0xe6, 0x19, 0x47, 0x78, 0xdb, 0x07, 0x16, 0x77, 0xbf, 0x55, 0xcd, 0x37, 0x8b, 0x44, 0x99, 0x4b,
	0x12, 0x73, 0x49, 0x3b, 0x18, 0x77, 0xf6, 0x2f, 0xa6, 0x75, 0xed, 0x1f, 0xe9, 0xb9, 0x40, 0xff,
	0x8c, 0xa5, 0x0b, 0xcc, 0x91, 0xe1, 0xad, 0x7c, 0xdc, 0xb3, 0xce, 0x69, 0x1d, 0xe3, 0xc2, 0x9a,
	0x2c, 0x87, 0xc2, 0xf0, 0x54, 0x18, 0x45, 0xbc, 0x43, 0xa3, 0x88, 0x6b, 0x6d, 0x8e, 0x3a, 0x18,
	0x55, 0xbc, 0x1d, 0x51, 0x08, 0xf5, 0xf8, 0x53, 0x05, 0x38, 0x92, 0x61, 0xbd, 0xc1, 0xa5, 0xcd,
	0x86, 0x93, 0x11, 0x19, 0x1d, 0x7c, 0x3d, 0x92, 0x55, 0xa0, 0x84, 0x64, 0x27, 0xd5, 0x54, 0x1f,
	0x53, 0x64, 0x39, 0x49, 0x60, 0xf3, 0x17, 0xc2, 0x5b, 0x5d, 0x60, 0xc6, 0x07, 0x7c, 0x63, 0x6d,
	0xb5, 0x1e, 0xa4, 0xa6, 0xda, 0x30, 0xab, 0xfc, 0xe8, 0x2a, 0xac, 0x44, 0xaf, 0x55, 0xb8, 0xf8,
	0xdb, 0x52, 0x83, 0xe2, 0x9b, 0xeb, 0x76, 0x3e, 0xfc, 0x57, 0xce, 0x35, 0x5a, 0xb9, 0x7e, 0x25,
	0x5a, 0x52, 0xbb, 0xbc, 0xf3, 0x31, 0xde, 0xe0, 0xce, 0xe1, 0xf9, 0xdc, 0x44, 0x97, 0x73, 0x13,
	0xfd, 0x9c, 0x9b, 0xe8, 0xcb, 0xc2, 0xcc, 0x5c, 0x2e, 0xcc, 0xcc, 0xf7, 0x85, 0x99, 0x79, 0xad,
	0x7f, 0x63, 0x18, 0x9c, 0x10, 0x8f, 0x27, 0xfb, 0x2a, 0xc6, 0x21, 0x85, 0x5e, 0x56, 0xfa, 0x72,
	0xf8, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xfe, 0x89, 0xaa, 0x66, 0x4e, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// UpdateParams defines a (governance) operation for updating the x/auth module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	// NonAtomicExec allows users to submit multiple messages for non-atomic execution.
	NonAtomicExec(ctx context.Context, in *MsgNonAtomicExec, opts ...grpc.CallOption) (*MsgNonAtomicExecResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.auth.v1beta1.Msg/UpdateParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) NonAtomicExec(ctx context.Context, in *MsgNonAtomicExec, opts ...grpc.CallOption) (*MsgNonAtomicExecResponse, error) {
	out := new(MsgNonAtomicExecResponse)
	err := c.cc.Invoke(ctx, "/cosmos.auth.v1beta1.Msg/NonAtomicExec", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// UpdateParams defines a (governance) operation for updating the x/auth module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	// NonAtomicExec allows users to submit multiple messages for non-atomic execution.
	NonAtomicExec(context.Context, *MsgNonAtomicExec) (*MsgNonAtomicExecResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) UpdateParams(ctx context.Context, req *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (*UnimplementedMsgServer) NonAtomicExec(ctx context.Context, req *MsgNonAtomicExec) (*MsgNonAtomicExecResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NonAtomicExec not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.auth.v1beta1.Msg/UpdateParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_NonAtomicExec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgNonAtomicExec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).NonAtomicExec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.auth.v1beta1.Msg/NonAtomicExec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).NonAtomicExec(ctx, req.(*MsgNonAtomicExec))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.auth.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "NonAtomicExec",
			Handler:    _Msg_NonAtomicExec_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/auth/v1beta1/tx.proto",
}

func (m *MsgUpdateParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgNonAtomicExec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgNonAtomicExec) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgNonAtomicExec) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Msgs) > 0 {
		for iNdEx := len(m.Msgs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Msgs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *NonAtomicExecResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NonAtomicExecResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NonAtomicExecResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Resp != nil {
		{
			size, err := m.Resp.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Error) > 0 {
		i -= len(m.Error)
		copy(dAtA[i:], m.Error)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Error)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgNonAtomicExecResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgNonAtomicExecResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgNonAtomicExecResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Results) > 0 {
		for iNdEx := len(m.Results) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Results[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgUpdateParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Params.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgUpdateParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgNonAtomicExec) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Msgs) > 0 {
		for _, e := range m.Msgs {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *NonAtomicExecResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Resp != nil {
		l = m.Resp.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgNonAtomicExecResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Results) > 0 {
		for _, e := range m.Results {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgUpdateParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUpdateParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgNonAtomicExec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgNonAtomicExec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgNonAtomicExec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msgs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msgs = append(m.Msgs, &any.Any{})
			if err := m.Msgs[len(m.Msgs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *NonAtomicExecResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NonAtomicExecResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NonAtomicExecResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Resp == nil {
				m.Resp = &any.Any{}
			}
			if err := m.Resp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgNonAtomicExecResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgNonAtomicExecResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgNonAtomicExecResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Results", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Results = append(m.Results, &NonAtomicExecResult{})
			if err := m.Results[len(m.Results)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
