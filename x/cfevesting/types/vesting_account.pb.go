// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: c4e-chain/cfevesting/vesting_account.proto

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

type VestingAccountTrace struct {
	Id                     uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Address                string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Genesis                bool   `protobuf:"varint,3,opt,name=genesis,proto3" json:"genesis,omitempty"`
	SourceVestingPoolOwner string `protobuf:"bytes,4,opt,name=source_vesting_pool_owner,json=sourceVestingPoolOwner,proto3" json:"source_vesting_pool_owner,omitempty"`
	SourceVestingPool      string `protobuf:"bytes,5,opt,name=source_vesting_pool,json=sourceVestingPool,proto3" json:"source_vesting_pool,omitempty"`
	SourceAccount          string `protobuf:"bytes,6,opt,name=source_account,json=sourceAccount,proto3" json:"source_account,omitempty"`
}

func (m *VestingAccountTrace) Reset()         { *m = VestingAccountTrace{} }
func (m *VestingAccountTrace) String() string { return proto.CompactTextString(m) }
func (*VestingAccountTrace) ProtoMessage()    {}
func (*VestingAccountTrace) Descriptor() ([]byte, []int) {
	return fileDescriptor_2177771eed4a6edc, []int{0}
}
func (m *VestingAccountTrace) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VestingAccountTrace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VestingAccountTrace.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VestingAccountTrace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VestingAccountTrace.Merge(m, src)
}
func (m *VestingAccountTrace) XXX_Size() int {
	return m.Size()
}
func (m *VestingAccountTrace) XXX_DiscardUnknown() {
	xxx_messageInfo_VestingAccountTrace.DiscardUnknown(m)
}

var xxx_messageInfo_VestingAccountTrace proto.InternalMessageInfo

func (m *VestingAccountTrace) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *VestingAccountTrace) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *VestingAccountTrace) GetGenesis() bool {
	if m != nil {
		return m.Genesis
	}
	return false
}

func (m *VestingAccountTrace) GetSourceVestingPoolOwner() string {
	if m != nil {
		return m.SourceVestingPoolOwner
	}
	return ""
}

func (m *VestingAccountTrace) GetSourceVestingPool() string {
	if m != nil {
		return m.SourceVestingPool
	}
	return ""
}

func (m *VestingAccountTrace) GetSourceAccount() string {
	if m != nil {
		return m.SourceAccount
	}
	return ""
}

func init() {
	proto.RegisterType((*VestingAccountTrace)(nil), "chain4energy.c4echain.cfevesting.VestingAccountTrace")
}

func init() {
	proto.RegisterFile("c4e-chain/cfevesting/vesting_account.proto", fileDescriptor_2177771eed4a6edc)
}

var fileDescriptor_2177771eed4a6edc = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xc1, 0x4a, 0xc3, 0x30,
	0x18, 0xc7, 0x9b, 0x3a, 0xa7, 0x06, 0x1c, 0x98, 0x81, 0xc4, 0x4b, 0x28, 0x82, 0x50, 0x04, 0xdb,
	0x83, 0xbd, 0x78, 0xd4, 0x07, 0x50, 0x29, 0xe2, 0xc1, 0x4b, 0xe9, 0xd2, 0xcf, 0x2e, 0x30, 0x93,
	0x92, 0xb4, 0xea, 0xde, 0xc2, 0xc7, 0xf2, 0xb8, 0xa3, 0x47, 0x69, 0x8f, 0xbe, 0x84, 0x2c, 0x4d,
	0xd9, 0x40, 0x4f, 0xed, 0x3f, 0xdf, 0xff, 0x17, 0xbe, 0xfc, 0xf0, 0x39, 0x4f, 0xe0, 0x82, 0xcf,
	0x73, 0x21, 0x63, 0xfe, 0x0c, 0xaf, 0x60, 0x6a, 0x21, 0xcb, 0xd8, 0x7d, 0xb3, 0x9c, 0x73, 0xd5,
	0xc8, 0x3a, 0xaa, 0xb4, 0xaa, 0x15, 0x09, 0x6c, 0x2f, 0x01, 0x09, 0xba, 0x5c, 0x46, 0x3c, 0x01,
	0x9b, 0xa3, 0x0d, 0x77, 0xfa, 0x83, 0xf0, 0xf4, 0xb1, 0xff, 0xbf, 0xee, 0xd1, 0x07, 0x9d, 0x73,
	0x20, 0x13, 0xec, 0x8b, 0x82, 0xa2, 0x00, 0x85, 0xa3, 0xd4, 0x17, 0x05, 0xa1, 0x78, 0x2f, 0x2f,
	0x0a, 0x0d, 0xc6, 0x50, 0x3f, 0x40, 0xe1, 0x41, 0x3a, 0xc4, 0xf5, 0xa4, 0x04, 0x09, 0x46, 0x18,
	0xba, 0x13, 0xa0, 0x70, 0x3f, 0x1d, 0x22, 0xb9, 0xc2, 0x27, 0x46, 0x35, 0x9a, 0x43, 0x36, 0x6c,
	0x57, 0x29, 0xb5, 0xc8, 0xd4, 0x9b, 0x04, 0x4d, 0x47, 0xf6, 0x96, 0xe3, 0xbe, 0xe0, 0x36, 0xb8,
	0x57, 0x6a, 0x71, 0xb7, 0x9e, 0x92, 0x08, 0x4f, 0xff, 0x41, 0xe9, 0xae, 0x85, 0x8e, 0xfe, 0x40,
	0xe4, 0x0c, 0x4f, 0x5c, 0xdf, 0x09, 0xa0, 0x63, 0x5b, 0x3d, 0xec, 0x4f, 0xdd, 0xd3, 0x6e, 0x6e,
	0x3f, 0x5b, 0x86, 0x56, 0x2d, 0x43, 0xdf, 0x2d, 0x43, 0x1f, 0x1d, 0xf3, 0x56, 0x1d, 0xf3, 0xbe,
	0x3a, 0xe6, 0x3d, 0x25, 0xa5, 0xa8, 0xe7, 0xcd, 0x2c, 0xe2, 0xea, 0x25, 0xde, 0x96, 0x16, 0x6f,
	0x6c, 0xbf, 0x6f, 0xfb, 0xae, 0x97, 0x15, 0x98, 0xd9, 0xd8, 0x6a, 0xbe, 0xfc, 0x0d, 0x00, 0x00,
	0xff, 0xff, 0x58, 0x94, 0xcb, 0xbb, 0x94, 0x01, 0x00, 0x00,
}

func (m *VestingAccountTrace) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VestingAccountTrace) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VestingAccountTrace) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SourceAccount) > 0 {
		i -= len(m.SourceAccount)
		copy(dAtA[i:], m.SourceAccount)
		i = encodeVarintVestingAccount(dAtA, i, uint64(len(m.SourceAccount)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.SourceVestingPool) > 0 {
		i -= len(m.SourceVestingPool)
		copy(dAtA[i:], m.SourceVestingPool)
		i = encodeVarintVestingAccount(dAtA, i, uint64(len(m.SourceVestingPool)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.SourceVestingPoolOwner) > 0 {
		i -= len(m.SourceVestingPoolOwner)
		copy(dAtA[i:], m.SourceVestingPoolOwner)
		i = encodeVarintVestingAccount(dAtA, i, uint64(len(m.SourceVestingPoolOwner)))
		i--
		dAtA[i] = 0x22
	}
	if m.Genesis {
		i--
		if m.Genesis {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintVestingAccount(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintVestingAccount(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintVestingAccount(dAtA []byte, offset int, v uint64) int {
	offset -= sovVestingAccount(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *VestingAccountTrace) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovVestingAccount(uint64(m.Id))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovVestingAccount(uint64(l))
	}
	if m.Genesis {
		n += 2
	}
	l = len(m.SourceVestingPoolOwner)
	if l > 0 {
		n += 1 + l + sovVestingAccount(uint64(l))
	}
	l = len(m.SourceVestingPool)
	if l > 0 {
		n += 1 + l + sovVestingAccount(uint64(l))
	}
	l = len(m.SourceAccount)
	if l > 0 {
		n += 1 + l + sovVestingAccount(uint64(l))
	}
	return n
}

func sovVestingAccount(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVestingAccount(x uint64) (n int) {
	return sovVestingAccount(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *VestingAccountTrace) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVestingAccount
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
			return fmt.Errorf("proto: VestingAccountTrace: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VestingAccountTrace: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVestingAccount
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
					return ErrIntOverflowVestingAccount
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
				return ErrInvalidLengthVestingAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVestingAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Genesis", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVestingAccount
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
			m.Genesis = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceVestingPoolOwner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVestingAccount
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
				return ErrInvalidLengthVestingAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVestingAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceVestingPoolOwner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceVestingPool", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVestingAccount
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
				return ErrInvalidLengthVestingAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVestingAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceVestingPool = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceAccount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVestingAccount
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
				return ErrInvalidLengthVestingAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVestingAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceAccount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVestingAccount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVestingAccount
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
func skipVestingAccount(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVestingAccount
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
					return 0, ErrIntOverflowVestingAccount
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
					return 0, ErrIntOverflowVestingAccount
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
				return 0, ErrInvalidLengthVestingAccount
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVestingAccount
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVestingAccount
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVestingAccount        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVestingAccount          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVestingAccount = fmt.Errorf("proto: unexpected end of group")
)
