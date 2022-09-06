// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: axelar/axelarnet/v1beta1/params.proto

package types

import (
	fmt "fmt"
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

// Params represent the genesis parameters for the module
type Params struct {
	// IBC packet route timeout window
	RouteTimeoutWindow uint64 `protobuf:"varint,1,opt,name=route_timeout_window,json=routeTimeoutWindow,proto3" json:"route_timeout_window,omitempty"`
	TransferLimit      uint64 `protobuf:"varint,3,opt,name=transfer_limit,json=transferLimit,proto3" json:"transfer_limit,omitempty"`
	EndBlockerLimit    uint64 `protobuf:"varint,4,opt,name=end_blocker_limit,json=endBlockerLimit,proto3" json:"end_blocker_limit,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_241d9e88637795bf, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Params)(nil), "axelar.axelarnet.v1beta1.Params")
}

func init() {
	proto.RegisterFile("axelar/axelarnet/v1beta1/params.proto", fileDescriptor_241d9e88637795bf)
}

var fileDescriptor_241d9e88637795bf = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4d, 0xac, 0x48, 0xcd,
	0x49, 0x2c, 0xd2, 0x87, 0x50, 0x79, 0xa9, 0x25, 0xfa, 0x65, 0x86, 0x49, 0xa9, 0x25, 0x89, 0x86,
	0xfa, 0x05, 0x89, 0x45, 0x89, 0xb9, 0xc5, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x12, 0x10,
	0x79, 0x3d, 0xb8, 0x32, 0x3d, 0xa8, 0x32, 0x29, 0x91, 0xf4, 0xfc, 0xf4, 0x7c, 0xb0, 0x22, 0x7d,
	0x10, 0x0b, 0xa2, 0x5e, 0x69, 0x32, 0x23, 0x17, 0x5b, 0x00, 0xd8, 0x00, 0x21, 0x03, 0x2e, 0x91,
	0xa2, 0xfc, 0xd2, 0x92, 0xd4, 0xf8, 0x92, 0xcc, 0xdc, 0xd4, 0xfc, 0xd2, 0x92, 0xf8, 0xf2, 0xcc,
	0xbc, 0x94, 0xfc, 0x72, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x96, 0x20, 0x21, 0xb0, 0x5c, 0x08, 0x44,
	0x2a, 0x1c, 0x2c, 0x23, 0xa4, 0xca, 0xc5, 0x57, 0x52, 0x94, 0x98, 0x57, 0x9c, 0x96, 0x5a, 0x14,
	0x9f, 0x93, 0x99, 0x9b, 0x59, 0x22, 0xc1, 0x0c, 0x56, 0xcb, 0x0b, 0x13, 0xf5, 0x01, 0x09, 0x0a,
	0x69, 0x71, 0x09, 0xa6, 0xe6, 0xa5, 0xc4, 0x27, 0xe5, 0xe4, 0x27, 0x67, 0xc3, 0x55, 0xb2, 0x80,
	0x55, 0xf2, 0xa7, 0xe6, 0xa5, 0x38, 0x41, 0xc4, 0xc1, 0x6a, 0xbd, 0x58, 0x38, 0x98, 0x04, 0x98,
	0x9d, 0x42, 0x4e, 0x3c, 0x94, 0x63, 0x38, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07,
	0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86,
	0x28, 0xb3, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x44, 0x70, 0x94, 0xe7,
	0x17, 0x65, 0x43, 0x79, 0xba, 0xc9, 0xf9, 0x45, 0xa9, 0xfa, 0x15, 0x48, 0x41, 0x55, 0x52, 0x59,
	0x90, 0x5a, 0x9c, 0xc4, 0x06, 0xf6, 0xb2, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xa9, 0x28, 0x58,
	0x99, 0x4b, 0x01, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EndBlockerLimit != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.EndBlockerLimit))
		i--
		dAtA[i] = 0x20
	}
	if m.TransferLimit != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.TransferLimit))
		i--
		dAtA[i] = 0x18
	}
	if m.RouteTimeoutWindow != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.RouteTimeoutWindow))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RouteTimeoutWindow != 0 {
		n += 1 + sovParams(uint64(m.RouteTimeoutWindow))
	}
	if m.TransferLimit != 0 {
		n += 1 + sovParams(uint64(m.TransferLimit))
	}
	if m.EndBlockerLimit != 0 {
		n += 1 + sovParams(uint64(m.EndBlockerLimit))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RouteTimeoutWindow", wireType)
			}
			m.RouteTimeoutWindow = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RouteTimeoutWindow |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TransferLimit", wireType)
			}
			m.TransferLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TransferLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndBlockerLimit", wireType)
			}
			m.EndBlockerLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndBlockerLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
