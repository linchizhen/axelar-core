// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: axelar/axelarnet/v1beta1/service.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("axelar/axelarnet/v1beta1/service.proto", fileDescriptor_4bc8aaff10777fba)
}
func init() {
	golang_proto.RegisterFile("axelar/axelarnet/v1beta1/service.proto", fileDescriptor_4bc8aaff10777fba)
}

var fileDescriptor_4bc8aaff10777fba = []byte{
	// 704 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x95, 0xcf, 0x6b, 0xd4, 0x4c,
	0x18, 0xc7, 0x3b, 0x2f, 0x2f, 0xef, 0x61, 0x68, 0x5f, 0x34, 0x16, 0x2b, 0x55, 0xa2, 0x0d, 0x6d,
	0xd1, 0x6d, 0x9b, 0xb4, 0xdb, 0x1f, 0x68, 0x6f, 0x6d, 0x54, 0x10, 0x2c, 0xe8, 0xda, 0x93, 0x97,
	0x30, 0x3b, 0x79, 0x9a, 0x86, 0x66, 0x67, 0xb6, 0x33, 0xb3, 0x75, 0xf7, 0xea, 0x5f, 0x20, 0x78,
	0xf2, 0x20, 0xe8, 0xd5, 0x83, 0x57, 0xc1, 0x93, 0x07, 0x11, 0x8f, 0x05, 0x11, 0x3c, 0x4a, 0xd7,
	0x3f, 0x44, 0x26, 0x99, 0xb4, 0xbb, 0x6e, 0xb3, 0x3f, 0x3c, 0xed, 0x86, 0x7c, 0x3f, 0x79, 0x3e,
	0xcf, 0xcc, 0xc3, 0x0c, 0x9e, 0x27, 0x4d, 0x48, 0x88, 0xf0, 0xb2, 0x1f, 0x06, 0xca, 0x3b, 0x5a,
	0xa9, 0x82, 0x22, 0x2b, 0x9e, 0x04, 0x71, 0x14, 0x53, 0x70, 0xeb, 0x82, 0x2b, 0x6e, 0x5d, 0xc9,
	0x02, 0xee, 0x69, 0xce, 0x35, 0xb9, 0xe9, 0xc9, 0x88, 0x47, 0x3c, 0x0d, 0x79, 0xfa, 0x5f, 0x96,
	0x9f, 0xbe, 0x16, 0x71, 0x1e, 0x25, 0xe0, 0x91, 0x7a, 0xec, 0x11, 0xc6, 0xb8, 0x22, 0x2a, 0xe6,
	0x4c, 0x9a, 0xb7, 0x33, 0x85, 0x55, 0x55, 0xd3, 0x44, 0x66, 0x0b, 0x23, 0x87, 0x0d, 0x10, 0xad,
	0x2c, 0x55, 0x7e, 0x33, 0x81, 0xf1, 0x8e, 0x8c, 0x9e, 0x64, 0xae, 0x56, 0x13, 0xff, 0xfb, 0x30,
	0x66, 0x07, 0xd6, 0x9c, 0x5b, 0xa4, 0xeb, 0xea, 0xf7, 0x15, 0x38, 0x6c, 0x80, 0x54, 0xd3, 0xf3,
	0x83, 0x62, 0xb2, 0xce, 0x99, 0x04, 0x67, 0xe6, 0xf9, 0xb7, 0x5f, 0x2f, 0xff, 0xb9, 0xea, 0x5c,
	0xf6, 0x7a, 0xa4, 0x92, 0x98, 0x1d, 0x6c, 0xa2, 0x92, 0xf5, 0x16, 0xe1, 0xff, 0x7d, 0xce, 0xf6,
	0x62, 0x51, 0xbb, 0x0b, 0x75, 0x2e, 0x63, 0x65, 0x79, 0xc5, 0x5f, 0xef, 0x4e, 0xe6, 0x3a, 0xcb,
	0xc3, 0x03, 0x46, 0x6c, 0x31, 0x15, 0x9b, 0x77, 0x66, 0x7a, 0xc5, 0x68, 0x46, 0x04, 0x61, 0x86,
	0x68, 0xc7, 0x2f, 0x08, 0x4f, 0xdd, 0x6b, 0x02, 0x6d, 0x28, 0x78, 0x04, 0x2c, 0x8c, 0x59, 0xb4,
	0x2b, 0x08, 0x93, 0x7b, 0x20, 0xa4, 0x75, 0xbb, 0xb8, 0x76, 0x01, 0x92, 0x5b, 0xdf, 0xf9, 0x0b,
	0xd2, 0xe8, 0x6f, 0xa4, 0xfa, 0xcb, 0xce, 0x42, 0xaf, 0x3e, 0x64, 0x68, 0x50, 0xcf, 0xd8, 0x40,
	0xe5, 0xb0, 0x6e, 0xe4, 0x03, 0xc2, 0x97, 0xb6, 0xc2, 0xd0, 0xe7, 0xb2, 0xc6, 0xe5, 0x36, 0x91,
	0x10, 0xfa, 0xfb, 0x24, 0x66, 0xd6, 0x5a, 0xb1, 0xca, 0x39, 0xf1, 0xbc, 0x81, 0xf5, 0x11, 0x29,
	0x23, 0xbf, 0x9a, 0xca, 0x2f, 0x39, 0x37, 0x7b, 0xe5, 0x49, 0x18, 0x06, 0x34, 0xe5, 0x82, 0xaa,
	0x06, 0x03, 0xaa, 0x49, 0x6d, 0xfe, 0x1a, 0xe1, 0x89, 0x0a, 0x44, 0xb1, 0x54, 0x20, 0xb6, 0xa4,
	0x04, 0x65, 0xb9, 0xc5, 0xd5, 0xbb, 0x82, 0xb9, 0xad, 0x37, 0x74, 0xde, 0x78, 0x2e, 0xa4, 0x9e,
	0x73, 0xce, 0x8d, 0x5e, 0x4f, 0x61, 0x80, 0x80, 0x68, 0x42, 0xfb, 0xbd, 0x47, 0xf8, 0x62, 0x85,
	0x37, 0x14, 0x3c, 0xd8, 0xf6, 0xcf, 0x86, 0xa3, 0xdc, 0xa7, 0xe6, 0x9f, 0xe1, 0xdc, 0x73, 0x75,
	0x24, 0xc6, 0xb8, 0x2e, 0xa7, 0xae, 0x25, 0x67, 0xee, 0x1c, 0x57, 0x0d, 0x05, 0x71, 0x95, 0x76,
	0x8f, 0xc2, 0x47, 0x84, 0x27, 0xf3, 0xbe, 0xef, 0x03, 0xf8, 0x3c, 0x49, 0x80, 0x2a, 0x2e, 0xac,
	0xf5, 0xc1, 0xeb, 0xd4, 0x99, 0xcf, 0xb5, 0x37, 0x46, 0xc5, 0x06, 0x4f, 0xc3, 0xe9, 0x2a, 0xef,
	0x01, 0x04, 0x34, 0x27, 0xb5, 0xfc, 0x3b, 0x84, 0x2f, 0x54, 0x40, 0x89, 0x56, 0xc7, 0x62, 0x58,
	0x2b, 0xfd, 0x0c, 0xba, 0xb3, 0xb9, 0x74, 0x79, 0x14, 0xc4, 0x08, 0x7b, 0xa9, 0xf0, 0x2d, 0x67,
	0xf6, 0x3c, 0x61, 0x25, 0x5a, 0x5d, 0x4b, 0xad, 0x65, 0x5f, 0x21, 0x3c, 0x9e, 0xee, 0xdc, 0x0e,
	0x48, 0x49, 0x22, 0xb0, 0x96, 0x06, 0xec, 0xb0, 0xc9, 0xe5, 0x92, 0xee, 0xb0, 0x71, 0x23, 0x58,
	0x4a, 0x05, 0x67, 0x9d, 0xeb, 0x45, 0xb3, 0x50, 0xcb, 0x80, 0xdc, 0xcd, 0x27, 0x49, 0xe2, 0x73,
	0xa6, 0x04, 0xa1, 0xaa, 0x9f, 0x5b, 0x67, 0x6e, 0x08, 0xb7, 0xee, 0xf8, 0x60, 0x37, 0x4a, 0x92,
	0x24, 0xa0, 0x06, 0xd8, 0x44, 0xa5, 0xf2, 0x77, 0x84, 0xc7, 0x1f, 0xeb, 0x2b, 0x2b, 0xbf, 0xa4,
	0x3e, 0x23, 0x3c, 0x65, 0x8e, 0xc4, 0x8e, 0x8d, 0xf1, 0x79, 0x83, 0xa9, 0x7e, 0xc7, 0x70, 0x01,
	0x32, 0xc4, 0x31, 0x5c, 0x48, 0x9a, 0x6e, 0xd6, 0xd2, 0x6e, 0x5c, 0x6b, 0xd1, 0x2b, 0xbc, 0x73,
	0x3b, 0x87, 0x21, 0xa0, 0x9a, 0xde, 0xde, 0xfd, 0x7a, 0x62, 0xa3, 0xe3, 0x13, 0x1b, 0xfd, 0x3c,
	0xb1, 0xd1, 0x8b, 0xb6, 0x3d, 0xf6, 0xa9, 0x6d, 0xa3, 0xe3, 0xb6, 0x3d, 0xf6, 0xa3, 0x6d, 0x8f,
	0x3d, 0xdd, 0x88, 0x62, 0xb5, 0xdf, 0xa8, 0xba, 0x94, 0xd7, 0xce, 0x3e, 0xf7, 0x8c, 0x8b, 0x03,
	0xf3, 0xb4, 0x44, 0xb9, 0x00, 0xaf, 0xd9, 0x51, 0x4a, 0xb5, 0xea, 0x20, 0xab, 0xff, 0xa5, 0xf7,
	0xfa, 0xea, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb5, 0x35, 0x66, 0x75, 0x98, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgServiceClient is the client API for MsgService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgServiceClient interface {
	Link(ctx context.Context, in *LinkRequest, opts ...grpc.CallOption) (*LinkResponse, error)
	ConfirmDeposit(ctx context.Context, in *ConfirmDepositRequest, opts ...grpc.CallOption) (*ConfirmDepositResponse, error)
	ExecutePendingTransfers(ctx context.Context, in *ExecutePendingTransfersRequest, opts ...grpc.CallOption) (*ExecutePendingTransfersResponse, error)
	AddCosmosBasedChain(ctx context.Context, in *AddCosmosBasedChainRequest, opts ...grpc.CallOption) (*AddCosmosBasedChainResponse, error)
	RegisterAsset(ctx context.Context, in *RegisterAssetRequest, opts ...grpc.CallOption) (*RegisterAssetResponse, error)
	RouteIBCTransfers(ctx context.Context, in *RouteIBCTransfersRequest, opts ...grpc.CallOption) (*RouteIBCTransfersResponse, error)
	RegisterFeeCollector(ctx context.Context, in *RegisterFeeCollectorRequest, opts ...grpc.CallOption) (*RegisterFeeCollectorResponse, error)
	RetryIBCTransfer(ctx context.Context, in *RetryIBCTransferRequest, opts ...grpc.CallOption) (*RetryIBCTransferResponse, error)
	RouteMessage(ctx context.Context, in *RouteMessageRequest, opts ...grpc.CallOption) (*RouteMessageResponse, error)
	CallContract(ctx context.Context, in *CallContractRequest, opts ...grpc.CallOption) (*CallContractResponse, error)
}

type msgServiceClient struct {
	cc grpc1.ClientConn
}

func NewMsgServiceClient(cc grpc1.ClientConn) MsgServiceClient {
	return &msgServiceClient{cc}
}

func (c *msgServiceClient) Link(ctx context.Context, in *LinkRequest, opts ...grpc.CallOption) (*LinkResponse, error) {
	out := new(LinkResponse)
	err := c.cc.Invoke(ctx, "/axelar.axelarnet.v1beta1.MsgService/Link", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) ConfirmDeposit(ctx context.Context, in *ConfirmDepositRequest, opts ...grpc.CallOption) (*ConfirmDepositResponse, error) {
	out := new(ConfirmDepositResponse)
	err := c.cc.Invoke(ctx, "/axelar.axelarnet.v1beta1.MsgService/ConfirmDeposit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) ExecutePendingTransfers(ctx context.Context, in *ExecutePendingTransfersRequest, opts ...grpc.CallOption) (*ExecutePendingTransfersResponse, error) {
	out := new(ExecutePendingTransfersResponse)
	err := c.cc.Invoke(ctx, "/axelar.axelarnet.v1beta1.MsgService/ExecutePendingTransfers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) AddCosmosBasedChain(ctx context.Context, in *AddCosmosBasedChainRequest, opts ...grpc.CallOption) (*AddCosmosBasedChainResponse, error) {
	out := new(AddCosmosBasedChainResponse)
	err := c.cc.Invoke(ctx, "/axelar.axelarnet.v1beta1.MsgService/AddCosmosBasedChain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) RegisterAsset(ctx context.Context, in *RegisterAssetRequest, opts ...grpc.CallOption) (*RegisterAssetResponse, error) {
	out := new(RegisterAssetResponse)
	err := c.cc.Invoke(ctx, "/axelar.axelarnet.v1beta1.MsgService/RegisterAsset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) RouteIBCTransfers(ctx context.Context, in *RouteIBCTransfersRequest, opts ...grpc.CallOption) (*RouteIBCTransfersResponse, error) {
	out := new(RouteIBCTransfersResponse)
	err := c.cc.Invoke(ctx, "/axelar.axelarnet.v1beta1.MsgService/RouteIBCTransfers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) RegisterFeeCollector(ctx context.Context, in *RegisterFeeCollectorRequest, opts ...grpc.CallOption) (*RegisterFeeCollectorResponse, error) {
	out := new(RegisterFeeCollectorResponse)
	err := c.cc.Invoke(ctx, "/axelar.axelarnet.v1beta1.MsgService/RegisterFeeCollector", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) RetryIBCTransfer(ctx context.Context, in *RetryIBCTransferRequest, opts ...grpc.CallOption) (*RetryIBCTransferResponse, error) {
	out := new(RetryIBCTransferResponse)
	err := c.cc.Invoke(ctx, "/axelar.axelarnet.v1beta1.MsgService/RetryIBCTransfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) RouteMessage(ctx context.Context, in *RouteMessageRequest, opts ...grpc.CallOption) (*RouteMessageResponse, error) {
	out := new(RouteMessageResponse)
	err := c.cc.Invoke(ctx, "/axelar.axelarnet.v1beta1.MsgService/RouteMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) CallContract(ctx context.Context, in *CallContractRequest, opts ...grpc.CallOption) (*CallContractResponse, error) {
	out := new(CallContractResponse)
	err := c.cc.Invoke(ctx, "/axelar.axelarnet.v1beta1.MsgService/CallContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServiceServer is the server API for MsgService service.
type MsgServiceServer interface {
	Link(context.Context, *LinkRequest) (*LinkResponse, error)
	ConfirmDeposit(context.Context, *ConfirmDepositRequest) (*ConfirmDepositResponse, error)
	ExecutePendingTransfers(context.Context, *ExecutePendingTransfersRequest) (*ExecutePendingTransfersResponse, error)
	AddCosmosBasedChain(context.Context, *AddCosmosBasedChainRequest) (*AddCosmosBasedChainResponse, error)
	RegisterAsset(context.Context, *RegisterAssetRequest) (*RegisterAssetResponse, error)
	RouteIBCTransfers(context.Context, *RouteIBCTransfersRequest) (*RouteIBCTransfersResponse, error)
	RegisterFeeCollector(context.Context, *RegisterFeeCollectorRequest) (*RegisterFeeCollectorResponse, error)
	RetryIBCTransfer(context.Context, *RetryIBCTransferRequest) (*RetryIBCTransferResponse, error)
	RouteMessage(context.Context, *RouteMessageRequest) (*RouteMessageResponse, error)
	CallContract(context.Context, *CallContractRequest) (*CallContractResponse, error)
}

// UnimplementedMsgServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServiceServer struct {
}

func (*UnimplementedMsgServiceServer) Link(ctx context.Context, req *LinkRequest) (*LinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Link not implemented")
}
func (*UnimplementedMsgServiceServer) ConfirmDeposit(ctx context.Context, req *ConfirmDepositRequest) (*ConfirmDepositResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmDeposit not implemented")
}
func (*UnimplementedMsgServiceServer) ExecutePendingTransfers(ctx context.Context, req *ExecutePendingTransfersRequest) (*ExecutePendingTransfersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecutePendingTransfers not implemented")
}
func (*UnimplementedMsgServiceServer) AddCosmosBasedChain(ctx context.Context, req *AddCosmosBasedChainRequest) (*AddCosmosBasedChainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCosmosBasedChain not implemented")
}
func (*UnimplementedMsgServiceServer) RegisterAsset(ctx context.Context, req *RegisterAssetRequest) (*RegisterAssetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterAsset not implemented")
}
func (*UnimplementedMsgServiceServer) RouteIBCTransfers(ctx context.Context, req *RouteIBCTransfersRequest) (*RouteIBCTransfersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RouteIBCTransfers not implemented")
}
func (*UnimplementedMsgServiceServer) RegisterFeeCollector(ctx context.Context, req *RegisterFeeCollectorRequest) (*RegisterFeeCollectorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterFeeCollector not implemented")
}
func (*UnimplementedMsgServiceServer) RetryIBCTransfer(ctx context.Context, req *RetryIBCTransferRequest) (*RetryIBCTransferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetryIBCTransfer not implemented")
}
func (*UnimplementedMsgServiceServer) RouteMessage(ctx context.Context, req *RouteMessageRequest) (*RouteMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RouteMessage not implemented")
}
func (*UnimplementedMsgServiceServer) CallContract(ctx context.Context, req *CallContractRequest) (*CallContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CallContract not implemented")
}

func RegisterMsgServiceServer(s grpc1.Server, srv MsgServiceServer) {
	s.RegisterService(&_MsgService_serviceDesc, srv)
}

func _MsgService_Link_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).Link(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelar.axelarnet.v1beta1.MsgService/Link",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).Link(ctx, req.(*LinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_ConfirmDeposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmDepositRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).ConfirmDeposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelar.axelarnet.v1beta1.MsgService/ConfirmDeposit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).ConfirmDeposit(ctx, req.(*ConfirmDepositRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_ExecutePendingTransfers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecutePendingTransfersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).ExecutePendingTransfers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelar.axelarnet.v1beta1.MsgService/ExecutePendingTransfers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).ExecutePendingTransfers(ctx, req.(*ExecutePendingTransfersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_AddCosmosBasedChain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCosmosBasedChainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).AddCosmosBasedChain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelar.axelarnet.v1beta1.MsgService/AddCosmosBasedChain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).AddCosmosBasedChain(ctx, req.(*AddCosmosBasedChainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_RegisterAsset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterAssetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).RegisterAsset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelar.axelarnet.v1beta1.MsgService/RegisterAsset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).RegisterAsset(ctx, req.(*RegisterAssetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_RouteIBCTransfers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RouteIBCTransfersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).RouteIBCTransfers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelar.axelarnet.v1beta1.MsgService/RouteIBCTransfers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).RouteIBCTransfers(ctx, req.(*RouteIBCTransfersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_RegisterFeeCollector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterFeeCollectorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).RegisterFeeCollector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelar.axelarnet.v1beta1.MsgService/RegisterFeeCollector",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).RegisterFeeCollector(ctx, req.(*RegisterFeeCollectorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_RetryIBCTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetryIBCTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).RetryIBCTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelar.axelarnet.v1beta1.MsgService/RetryIBCTransfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).RetryIBCTransfer(ctx, req.(*RetryIBCTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_RouteMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RouteMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).RouteMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelar.axelarnet.v1beta1.MsgService/RouteMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).RouteMessage(ctx, req.(*RouteMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_CallContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CallContractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).CallContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelar.axelarnet.v1beta1.MsgService/CallContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).CallContract(ctx, req.(*CallContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MsgService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "axelar.axelarnet.v1beta1.MsgService",
	HandlerType: (*MsgServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Link",
			Handler:    _MsgService_Link_Handler,
		},
		{
			MethodName: "ConfirmDeposit",
			Handler:    _MsgService_ConfirmDeposit_Handler,
		},
		{
			MethodName: "ExecutePendingTransfers",
			Handler:    _MsgService_ExecutePendingTransfers_Handler,
		},
		{
			MethodName: "AddCosmosBasedChain",
			Handler:    _MsgService_AddCosmosBasedChain_Handler,
		},
		{
			MethodName: "RegisterAsset",
			Handler:    _MsgService_RegisterAsset_Handler,
		},
		{
			MethodName: "RouteIBCTransfers",
			Handler:    _MsgService_RouteIBCTransfers_Handler,
		},
		{
			MethodName: "RegisterFeeCollector",
			Handler:    _MsgService_RegisterFeeCollector_Handler,
		},
		{
			MethodName: "RetryIBCTransfer",
			Handler:    _MsgService_RetryIBCTransfer_Handler,
		},
		{
			MethodName: "RouteMessage",
			Handler:    _MsgService_RouteMessage_Handler,
		},
		{
			MethodName: "CallContract",
			Handler:    _MsgService_CallContract_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "axelar/axelarnet/v1beta1/service.proto",
}

// QueryServiceClient is the client API for QueryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryServiceClient interface {
	// PendingIBCTransferCount queries the pending ibc transfers for all chains
	PendingIBCTransferCount(ctx context.Context, in *PendingIBCTransferCountRequest, opts ...grpc.CallOption) (*PendingIBCTransferCountResponse, error)
}

type queryServiceClient struct {
	cc grpc1.ClientConn
}

func NewQueryServiceClient(cc grpc1.ClientConn) QueryServiceClient {
	return &queryServiceClient{cc}
}

func (c *queryServiceClient) PendingIBCTransferCount(ctx context.Context, in *PendingIBCTransferCountRequest, opts ...grpc.CallOption) (*PendingIBCTransferCountResponse, error) {
	out := new(PendingIBCTransferCountResponse)
	err := c.cc.Invoke(ctx, "/axelar.axelarnet.v1beta1.QueryService/PendingIBCTransferCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServiceServer is the server API for QueryService service.
type QueryServiceServer interface {
	// PendingIBCTransferCount queries the pending ibc transfers for all chains
	PendingIBCTransferCount(context.Context, *PendingIBCTransferCountRequest) (*PendingIBCTransferCountResponse, error)
}

// UnimplementedQueryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServiceServer struct {
}

func (*UnimplementedQueryServiceServer) PendingIBCTransferCount(ctx context.Context, req *PendingIBCTransferCountRequest) (*PendingIBCTransferCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PendingIBCTransferCount not implemented")
}

func RegisterQueryServiceServer(s grpc1.Server, srv QueryServiceServer) {
	s.RegisterService(&_QueryService_serviceDesc, srv)
}

func _QueryService_PendingIBCTransferCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PendingIBCTransferCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).PendingIBCTransferCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelar.axelarnet.v1beta1.QueryService/PendingIBCTransferCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).PendingIBCTransferCount(ctx, req.(*PendingIBCTransferCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _QueryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "axelar.axelarnet.v1beta1.QueryService",
	HandlerType: (*QueryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PendingIBCTransferCount",
			Handler:    _QueryService_PendingIBCTransferCount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "axelar/axelarnet/v1beta1/service.proto",
}
