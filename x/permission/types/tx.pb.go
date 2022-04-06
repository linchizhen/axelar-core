// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: permission/v1beta1/tx.proto

package types

import (
	fmt "fmt"
	_ "github.com/axelarnetwork/axelar-core/x/permission/exported"
	multisig "github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type UpdateGovernanceKeyRequest struct {
	Sender        github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=sender,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"sender,omitempty"`
	GovernanceKey multisig.LegacyAminoPubKey                    `protobuf:"bytes,2,opt,name=governance_key,json=governanceKey,proto3" json:"governance_key"`
}

func (m *UpdateGovernanceKeyRequest) Reset()         { *m = UpdateGovernanceKeyRequest{} }
func (m *UpdateGovernanceKeyRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateGovernanceKeyRequest) ProtoMessage()    {}
func (*UpdateGovernanceKeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d3d2ab82ee63830, []int{0}
}
func (m *UpdateGovernanceKeyRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdateGovernanceKeyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdateGovernanceKeyRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdateGovernanceKeyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateGovernanceKeyRequest.Merge(m, src)
}
func (m *UpdateGovernanceKeyRequest) XXX_Size() int {
	return m.Size()
}
func (m *UpdateGovernanceKeyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateGovernanceKeyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateGovernanceKeyRequest proto.InternalMessageInfo

type UpdateGovernanceKeyResponse struct {
}

func (m *UpdateGovernanceKeyResponse) Reset()         { *m = UpdateGovernanceKeyResponse{} }
func (m *UpdateGovernanceKeyResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateGovernanceKeyResponse) ProtoMessage()    {}
func (*UpdateGovernanceKeyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d3d2ab82ee63830, []int{1}
}
func (m *UpdateGovernanceKeyResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdateGovernanceKeyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdateGovernanceKeyResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdateGovernanceKeyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateGovernanceKeyResponse.Merge(m, src)
}
func (m *UpdateGovernanceKeyResponse) XXX_Size() int {
	return m.Size()
}
func (m *UpdateGovernanceKeyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateGovernanceKeyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateGovernanceKeyResponse proto.InternalMessageInfo

// MsgRegisterController represents a message to register a controller account
type RegisterControllerRequest struct {
	Sender     github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=sender,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"sender,omitempty"`
	Controller github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=controller,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"controller,omitempty"`
}

func (m *RegisterControllerRequest) Reset()         { *m = RegisterControllerRequest{} }
func (m *RegisterControllerRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterControllerRequest) ProtoMessage()    {}
func (*RegisterControllerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d3d2ab82ee63830, []int{2}
}
func (m *RegisterControllerRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisterControllerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisterControllerRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisterControllerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterControllerRequest.Merge(m, src)
}
func (m *RegisterControllerRequest) XXX_Size() int {
	return m.Size()
}
func (m *RegisterControllerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterControllerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterControllerRequest proto.InternalMessageInfo

type RegisterControllerResponse struct {
}

func (m *RegisterControllerResponse) Reset()         { *m = RegisterControllerResponse{} }
func (m *RegisterControllerResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterControllerResponse) ProtoMessage()    {}
func (*RegisterControllerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d3d2ab82ee63830, []int{3}
}
func (m *RegisterControllerResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisterControllerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisterControllerResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisterControllerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterControllerResponse.Merge(m, src)
}
func (m *RegisterControllerResponse) XXX_Size() int {
	return m.Size()
}
func (m *RegisterControllerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterControllerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterControllerResponse proto.InternalMessageInfo

// DeregisterController represents a message to deregister a controller account
type DeregisterControllerRequest struct {
	Sender     github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=sender,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"sender,omitempty"`
	Controller github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=controller,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"controller,omitempty"`
}

func (m *DeregisterControllerRequest) Reset()         { *m = DeregisterControllerRequest{} }
func (m *DeregisterControllerRequest) String() string { return proto.CompactTextString(m) }
func (*DeregisterControllerRequest) ProtoMessage()    {}
func (*DeregisterControllerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d3d2ab82ee63830, []int{4}
}
func (m *DeregisterControllerRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DeregisterControllerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DeregisterControllerRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DeregisterControllerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeregisterControllerRequest.Merge(m, src)
}
func (m *DeregisterControllerRequest) XXX_Size() int {
	return m.Size()
}
func (m *DeregisterControllerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeregisterControllerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeregisterControllerRequest proto.InternalMessageInfo

type DeregisterControllerResponse struct {
}

func (m *DeregisterControllerResponse) Reset()         { *m = DeregisterControllerResponse{} }
func (m *DeregisterControllerResponse) String() string { return proto.CompactTextString(m) }
func (*DeregisterControllerResponse) ProtoMessage()    {}
func (*DeregisterControllerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d3d2ab82ee63830, []int{5}
}
func (m *DeregisterControllerResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DeregisterControllerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DeregisterControllerResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DeregisterControllerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeregisterControllerResponse.Merge(m, src)
}
func (m *DeregisterControllerResponse) XXX_Size() int {
	return m.Size()
}
func (m *DeregisterControllerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeregisterControllerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeregisterControllerResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*UpdateGovernanceKeyRequest)(nil), "permission.v1beta1.UpdateGovernanceKeyRequest")
	proto.RegisterType((*UpdateGovernanceKeyResponse)(nil), "permission.v1beta1.UpdateGovernanceKeyResponse")
	proto.RegisterType((*RegisterControllerRequest)(nil), "permission.v1beta1.RegisterControllerRequest")
	proto.RegisterType((*RegisterControllerResponse)(nil), "permission.v1beta1.RegisterControllerResponse")
	proto.RegisterType((*DeregisterControllerRequest)(nil), "permission.v1beta1.DeregisterControllerRequest")
	proto.RegisterType((*DeregisterControllerResponse)(nil), "permission.v1beta1.DeregisterControllerResponse")
}

func init() { proto.RegisterFile("permission/v1beta1/tx.proto", fileDescriptor_1d3d2ab82ee63830) }

var fileDescriptor_1d3d2ab82ee63830 = []byte{
	// 409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x93, 0x3f, 0x8f, 0xd3, 0x30,
	0x18, 0xc6, 0x63, 0x38, 0xdd, 0x60, 0xfe, 0x0c, 0x11, 0x43, 0x49, 0x0f, 0xdf, 0xd1, 0x85, 0x63,
	0x68, 0xac, 0xc2, 0x80, 0xc4, 0xd6, 0x82, 0x84, 0xd0, 0x31, 0x40, 0xa4, 0x63, 0x60, 0x41, 0x89,
	0xf3, 0xca, 0x44, 0x49, 0xfc, 0x06, 0xdb, 0x39, 0x92, 0x8d, 0x8f, 0xc0, 0x97, 0x61, 0x46, 0x0c,
	0x48, 0x1d, 0x6f, 0x64, 0x3a, 0x41, 0xfb, 0x2d, 0x98, 0x10, 0x89, 0xaf, 0xcd, 0xd0, 0x09, 0x89,
	0x81, 0xc9, 0xb6, 0xde, 0xc7, 0xcf, 0xf3, 0xfc, 0x86, 0x97, 0x8e, 0x2b, 0xd0, 0x65, 0x66, 0x4c,
	0x86, 0x8a, 0x9f, 0xcd, 0x12, 0xb0, 0xf1, 0x8c, 0xdb, 0x26, 0xac, 0x34, 0x5a, 0xf4, 0xfd, 0xed,
	0x30, 0x74, 0xc3, 0xe0, 0x96, 0x44, 0x89, 0xdd, 0x98, 0xff, 0xb9, 0xf5, 0xca, 0xe0, 0xae, 0x40,
	0x53, 0xa2, 0xe1, 0x42, 0xb7, 0x95, 0x45, 0x5e, 0xd6, 0x85, 0xcd, 0x4c, 0x26, 0x79, 0x0e, 0xad,
	0x71, 0x92, 0x7b, 0x83, 0x24, 0x68, 0x2a, 0xd4, 0x16, 0xd2, 0x6d, 0x64, 0x5b, 0x81, 0x13, 0x4e,
	0xbe, 0x11, 0x1a, 0x9c, 0x56, 0x69, 0x6c, 0xe1, 0x19, 0x9e, 0x81, 0x56, 0xb1, 0x12, 0x70, 0x02,
	0x6d, 0x04, 0xef, 0x6b, 0x30, 0xd6, 0x7f, 0x4e, 0xf7, 0x0d, 0xa8, 0x14, 0xf4, 0x88, 0x1c, 0x91,
	0xe3, 0xeb, 0x8b, 0xd9, 0xaf, 0x8b, 0xc3, 0xa9, 0xcc, 0xec, 0xbb, 0x3a, 0x09, 0x05, 0x96, 0xfc,
	0xb2, 0x49, 0x77, 0x4c, 0x4d, 0x9a, 0x3b, 0xf3, 0xb9, 0x10, 0xf3, 0x34, 0xd5, 0x60, 0x4c, 0xe4,
	0x0c, 0xfc, 0xd7, 0xf4, 0xa6, 0xdc, 0x44, 0xbc, 0xcd, 0xa1, 0x1d, 0x5d, 0x39, 0x22, 0xc7, 0xd7,
	0x1e, 0xdc, 0x0f, 0xfb, 0xdf, 0x61, 0x8f, 0x13, 0x5e, 0xe2, 0x84, 0x2f, 0x40, 0xc6, 0xa2, 0x9d,
	0x97, 0x99, 0xc2, 0x97, 0x75, 0x72, 0x02, 0xed, 0x62, 0x6f, 0x79, 0x71, 0xe8, 0x45, 0x37, 0xe4,
	0xb0, 0xe9, 0xe3, 0xbd, 0x8f, 0x9f, 0x47, 0x57, 0x27, 0x77, 0xe8, 0x78, 0x27, 0x86, 0xa9, 0x50,
	0x19, 0x98, 0x7c, 0x21, 0xf4, 0x76, 0x04, 0x32, 0x33, 0x16, 0xf4, 0x13, 0x54, 0x56, 0x63, 0x51,
	0x80, 0xfe, 0x07, 0x94, 0xaf, 0x28, 0x15, 0x1b, 0xff, 0x8e, 0xf0, 0xaf, 0xec, 0x06, 0x26, 0x0e,
	0xf0, 0x80, 0x06, 0xbb, 0x00, 0x1c, 0xdf, 0x57, 0x42, 0xc7, 0x4f, 0x41, 0xff, 0xd7, 0x84, 0x8c,
	0x1e, 0xec, 0x46, 0xe8, 0x19, 0x17, 0xa7, 0xcb, 0x9f, 0xcc, 0x5b, 0xae, 0x18, 0x39, 0x5f, 0x31,
	0xf2, 0x63, 0xc5, 0xc8, 0xa7, 0x35, 0xf3, 0xce, 0xd7, 0xcc, 0xfb, 0xbe, 0x66, 0xde, 0x9b, 0x47,
	0x83, 0xf8, 0xb8, 0x81, 0x22, 0xd6, 0x0a, 0xec, 0x07, 0xd4, 0xb9, 0x7b, 0x4d, 0x05, 0x6a, 0xe0,
	0x0d, 0x1f, 0x2c, 0x46, 0xd7, 0x29, 0xd9, 0xef, 0x16, 0xe1, 0xe1, 0xef, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x2c, 0xf0, 0x21, 0xb4, 0x9d, 0x03, 0x00, 0x00,
}

func (m *UpdateGovernanceKeyRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateGovernanceKeyRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpdateGovernanceKeyRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.GovernanceKey.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UpdateGovernanceKeyResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateGovernanceKeyResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpdateGovernanceKeyResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *RegisterControllerRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisterControllerRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegisterControllerRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Controller) > 0 {
		i -= len(m.Controller)
		copy(dAtA[i:], m.Controller)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Controller)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RegisterControllerResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisterControllerResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegisterControllerResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *DeregisterControllerRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeregisterControllerRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DeregisterControllerRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Controller) > 0 {
		i -= len(m.Controller)
		copy(dAtA[i:], m.Controller)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Controller)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DeregisterControllerResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeregisterControllerResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DeregisterControllerResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
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
func (m *UpdateGovernanceKeyRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.GovernanceKey.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *UpdateGovernanceKeyResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *RegisterControllerRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Controller)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *RegisterControllerResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *DeregisterControllerRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Controller)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *DeregisterControllerResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UpdateGovernanceKeyRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: UpdateGovernanceKeyRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdateGovernanceKeyRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = append(m.Sender[:0], dAtA[iNdEx:postIndex]...)
			if m.Sender == nil {
				m.Sender = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GovernanceKey", wireType)
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
			if err := m.GovernanceKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *UpdateGovernanceKeyResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: UpdateGovernanceKeyResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdateGovernanceKeyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *RegisterControllerRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: RegisterControllerRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisterControllerRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = append(m.Sender[:0], dAtA[iNdEx:postIndex]...)
			if m.Sender == nil {
				m.Sender = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Controller", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Controller = append(m.Controller[:0], dAtA[iNdEx:postIndex]...)
			if m.Controller == nil {
				m.Controller = []byte{}
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
func (m *RegisterControllerResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: RegisterControllerResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisterControllerResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *DeregisterControllerRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: DeregisterControllerRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeregisterControllerRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = append(m.Sender[:0], dAtA[iNdEx:postIndex]...)
			if m.Sender == nil {
				m.Sender = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Controller", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Controller = append(m.Controller[:0], dAtA[iNdEx:postIndex]...)
			if m.Controller == nil {
				m.Controller = []byte{}
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
func (m *DeregisterControllerResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: DeregisterControllerResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeregisterControllerResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
