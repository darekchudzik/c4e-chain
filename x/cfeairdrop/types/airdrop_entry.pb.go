// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cfeairdrop/airdrop_entry.proto

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

type AirdropEntry struct {
	Id      uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Amount  uint64 `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *AirdropEntry) Reset()         { *m = AirdropEntry{} }
func (m *AirdropEntry) String() string { return proto.CompactTextString(m) }
func (*AirdropEntry) ProtoMessage()    {}
func (*AirdropEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2f4f037284c5342, []int{0}
}
func (m *AirdropEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AirdropEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AirdropEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AirdropEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AirdropEntry.Merge(m, src)
}
func (m *AirdropEntry) XXX_Size() int {
	return m.Size()
}
func (m *AirdropEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_AirdropEntry.DiscardUnknown(m)
}

var xxx_messageInfo_AirdropEntry proto.InternalMessageInfo

func (m *AirdropEntry) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *AirdropEntry) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *AirdropEntry) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func init() {
	proto.RegisterType((*AirdropEntry)(nil), "chain4energy.c4echain.cfeairdrop.AirdropEntry")
}

func init() { proto.RegisterFile("cfeairdrop/airdrop_entry.proto", fileDescriptor_a2f4f037284c5342) }

var fileDescriptor_a2f4f037284c5342 = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0x4e, 0x4b, 0x4d,
	0xcc, 0x2c, 0x4a, 0x29, 0xca, 0x2f, 0xd0, 0x87, 0xd2, 0xf1, 0xa9, 0x79, 0x25, 0x45, 0x95, 0x7a,
	0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x0a, 0xc9, 0x19, 0x89, 0x99, 0x79, 0x26, 0xa9, 0x79, 0xa9,
	0x45, 0xe9, 0x95, 0x7a, 0xc9, 0x26, 0xa9, 0x60, 0xbe, 0x1e, 0x42, 0x97, 0x52, 0x00, 0x17, 0x8f,
	0x23, 0x84, 0xe9, 0x0a, 0xd2, 0x27, 0xc4, 0xc7, 0xc5, 0x94, 0x99, 0x22, 0xc1, 0xa8, 0xc0, 0xa8,
	0xc1, 0x12, 0xc4, 0x94, 0x99, 0x22, 0x24, 0xc1, 0xc5, 0x9e, 0x98, 0x92, 0x52, 0x94, 0x5a, 0x5c,
	0x2c, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe3, 0x0a, 0x89, 0x71, 0xb1, 0x25, 0xe6, 0xe6,
	0x97, 0xe6, 0x95, 0x48, 0x30, 0x83, 0x55, 0x43, 0x79, 0x4e, 0x7e, 0x27, 0x1e, 0xc9, 0x31, 0x5e,
	0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31,
	0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x65, 0x92, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f,
	0xab, 0x8f, 0xec, 0x30, 0xfd, 0x64, 0x93, 0x54, 0x5d, 0xb0, 0x80, 0x7e, 0x85, 0x3e, 0x92, 0x8f,
	0x4a, 0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0x5e, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0xa3, 0xed, 0xf9, 0xed, 0xec, 0x00, 0x00, 0x00,
}

func (m *AirdropEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AirdropEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AirdropEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Amount != 0 {
		i = encodeVarintAirdropEntry(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintAirdropEntry(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintAirdropEntry(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintAirdropEntry(dAtA []byte, offset int, v uint64) int {
	offset -= sovAirdropEntry(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AirdropEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovAirdropEntry(uint64(m.Id))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovAirdropEntry(uint64(l))
	}
	if m.Amount != 0 {
		n += 1 + sovAirdropEntry(uint64(m.Amount))
	}
	return n
}

func sovAirdropEntry(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAirdropEntry(x uint64) (n int) {
	return sovAirdropEntry(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AirdropEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAirdropEntry
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
			return fmt.Errorf("proto: AirdropEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AirdropEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAirdropEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAirdropEntry
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
				return ErrInvalidLengthAirdropEntry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAirdropEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAirdropEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAirdropEntry(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAirdropEntry
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
func skipAirdropEntry(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAirdropEntry
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
					return 0, ErrIntOverflowAirdropEntry
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
					return 0, ErrIntOverflowAirdropEntry
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
				return 0, ErrInvalidLengthAirdropEntry
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAirdropEntry
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAirdropEntry
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAirdropEntry        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAirdropEntry          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAirdropEntry = fmt.Errorf("proto: unexpected end of group")
)
