// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: elys/oracle/price_feeder.proto

package types

import (
	fmt "fmt"
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

type PriceFeeder struct {
	Feeder   string `protobuf:"bytes,1,opt,name=feeder,proto3" json:"feeder,omitempty"`
	IsActive bool   `protobuf:"varint,2,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
}

func (m *PriceFeeder) Reset()         { *m = PriceFeeder{} }
func (m *PriceFeeder) String() string { return proto.CompactTextString(m) }
func (*PriceFeeder) ProtoMessage()    {}
func (*PriceFeeder) Descriptor() ([]byte, []int) {
	return fileDescriptor_f310b751e06e8fd3, []int{0}
}
func (m *PriceFeeder) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PriceFeeder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PriceFeeder.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PriceFeeder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PriceFeeder.Merge(m, src)
}
func (m *PriceFeeder) XXX_Size() int {
	return m.Size()
}
func (m *PriceFeeder) XXX_DiscardUnknown() {
	xxx_messageInfo_PriceFeeder.DiscardUnknown(m)
}

var xxx_messageInfo_PriceFeeder proto.InternalMessageInfo

func (m *PriceFeeder) GetFeeder() string {
	if m != nil {
		return m.Feeder
	}
	return ""
}

func (m *PriceFeeder) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
}

func init() {
	proto.RegisterType((*PriceFeeder)(nil), "elysnetwork.elys.oracle.PriceFeeder")
}

func init() { proto.RegisterFile("elys/oracle/price_feeder.proto", fileDescriptor_f310b751e06e8fd3) }

var fileDescriptor_f310b751e06e8fd3 = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0xcd, 0xa9, 0x2c,
	0xd6, 0xcf, 0x2f, 0x4a, 0x4c, 0xce, 0x49, 0xd5, 0x2f, 0x28, 0xca, 0x4c, 0x4e, 0x8d, 0x4f, 0x4b,
	0x4d, 0x4d, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x07, 0xc9, 0xe7, 0xa5,
	0x96, 0x94, 0xe7, 0x17, 0x65, 0xeb, 0x81, 0xd8, 0x7a, 0x10, 0xb5, 0x4a, 0x4e, 0x5c, 0xdc, 0x01,
	0x20, 0xe5, 0x6e, 0x60, 0xd5, 0x42, 0x62, 0x5c, 0x6c, 0x10, 0x7d, 0x12, 0x8c, 0x0a, 0x8c, 0x1a,
	0x9c, 0x41, 0x50, 0x9e, 0x90, 0x34, 0x17, 0x67, 0x66, 0x71, 0x7c, 0x62, 0x72, 0x49, 0x66, 0x59,
	0xaa, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x47, 0x10, 0x47, 0x66, 0xb1, 0x23, 0x98, 0xef, 0xe4, 0x7a,
	0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7,
	0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0xda, 0xe9, 0x99, 0x25, 0x19, 0xa5,
	0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x20, 0x5b, 0x75, 0xa1, 0x4e, 0x00, 0x73, 0xf4, 0x2b, 0x60,
	0x0e, 0x2e, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0x3b, 0xd5, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0x6e, 0x50, 0x16, 0x70, 0xcc, 0x00, 0x00, 0x00,
}

func (m *PriceFeeder) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PriceFeeder) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PriceFeeder) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IsActive {
		i--
		if m.IsActive {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.Feeder) > 0 {
		i -= len(m.Feeder)
		copy(dAtA[i:], m.Feeder)
		i = encodeVarintPriceFeeder(dAtA, i, uint64(len(m.Feeder)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPriceFeeder(dAtA []byte, offset int, v uint64) int {
	offset -= sovPriceFeeder(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PriceFeeder) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Feeder)
	if l > 0 {
		n += 1 + l + sovPriceFeeder(uint64(l))
	}
	if m.IsActive {
		n += 2
	}
	return n
}

func sovPriceFeeder(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPriceFeeder(x uint64) (n int) {
	return sovPriceFeeder(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PriceFeeder) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPriceFeeder
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
			return fmt.Errorf("proto: PriceFeeder: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PriceFeeder: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Feeder", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPriceFeeder
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
				return ErrInvalidLengthPriceFeeder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPriceFeeder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Feeder = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsActive", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPriceFeeder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsActive = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipPriceFeeder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPriceFeeder
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
func skipPriceFeeder(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPriceFeeder
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
					return 0, ErrIntOverflowPriceFeeder
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
					return 0, ErrIntOverflowPriceFeeder
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
				return 0, ErrInvalidLengthPriceFeeder
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPriceFeeder
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPriceFeeder
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPriceFeeder        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPriceFeeder          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPriceFeeder = fmt.Errorf("proto: unexpected end of group")
)
