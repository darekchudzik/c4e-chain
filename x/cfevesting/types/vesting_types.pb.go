// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cfevesting/vesting_types.proto

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

type VestingTypes struct {
	VestingTypes []*VestingType `protobuf:"bytes,1,rep,name=vesting_types,json=vestingTypes,proto3" json:"vesting_types,omitempty"`
}

func (m *VestingTypes) Reset()         { *m = VestingTypes{} }
func (m *VestingTypes) String() string { return proto.CompactTextString(m) }
func (*VestingTypes) ProtoMessage()    {}
func (*VestingTypes) Descriptor() ([]byte, []int) {
	return fileDescriptor_67b318d31cacf480, []int{0}
}
func (m *VestingTypes) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VestingTypes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VestingTypes.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VestingTypes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VestingTypes.Merge(m, src)
}
func (m *VestingTypes) XXX_Size() int {
	return m.Size()
}
func (m *VestingTypes) XXX_DiscardUnknown() {
	xxx_messageInfo_VestingTypes.DiscardUnknown(m)
}

var xxx_messageInfo_VestingTypes proto.InternalMessageInfo

func (m *VestingTypes) GetVestingTypes() []*VestingType {
	if m != nil {
		return m.VestingTypes
	}
	return nil
}

type VestingType struct {
	// vesting type name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// period of locked coins (minutes) from vesting start
	LockupPeriod int64 `protobuf:"varint,2,opt,name=lockup_period,json=lockupPeriod,proto3" json:"lockup_period,omitempty"`
	// period of veesting coins (minutes) from lockup period end
	VestingPeriod int64 `protobuf:"varint,3,opt,name=vesting_period,json=vestingPeriod,proto3" json:"vesting_period,omitempty"`
	// vested coin periodical releasing (minutes)
	TokenReleasingPeriod int64 `protobuf:"varint,4,opt,name=token_releasing_period,json=tokenReleasingPeriod,proto3" json:"token_releasing_period,omitempty"`
	// defines if vesting type allows delegation
	DelegationsAllowed bool `protobuf:"varint,5,opt,name=delegations_allowed,json=delegationsAllowed,proto3" json:"delegations_allowed,omitempty"`
}

func (m *VestingType) Reset()         { *m = VestingType{} }
func (m *VestingType) String() string { return proto.CompactTextString(m) }
func (*VestingType) ProtoMessage()    {}
func (*VestingType) Descriptor() ([]byte, []int) {
	return fileDescriptor_67b318d31cacf480, []int{1}
}
func (m *VestingType) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VestingType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VestingType.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VestingType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VestingType.Merge(m, src)
}
func (m *VestingType) XXX_Size() int {
	return m.Size()
}
func (m *VestingType) XXX_DiscardUnknown() {
	xxx_messageInfo_VestingType.DiscardUnknown(m)
}

var xxx_messageInfo_VestingType proto.InternalMessageInfo

func (m *VestingType) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VestingType) GetLockupPeriod() int64 {
	if m != nil {
		return m.LockupPeriod
	}
	return 0
}

func (m *VestingType) GetVestingPeriod() int64 {
	if m != nil {
		return m.VestingPeriod
	}
	return 0
}

func (m *VestingType) GetTokenReleasingPeriod() int64 {
	if m != nil {
		return m.TokenReleasingPeriod
	}
	return 0
}

func (m *VestingType) GetDelegationsAllowed() bool {
	if m != nil {
		return m.DelegationsAllowed
	}
	return false
}

func init() {
	proto.RegisterType((*VestingTypes)(nil), "chain4energy.c4echain.cfevesting.VestingTypes")
	proto.RegisterType((*VestingType)(nil), "chain4energy.c4echain.cfevesting.VestingType")
}

func init() { proto.RegisterFile("cfevesting/vesting_types.proto", fileDescriptor_67b318d31cacf480) }

var fileDescriptor_67b318d31cacf480 = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xb1, 0x4e, 0xfb, 0x30,
	0x10, 0xc6, 0xeb, 0x7f, 0xfb, 0x47, 0xe0, 0xb6, 0x0c, 0x06, 0xa1, 0x4c, 0x56, 0x54, 0x84, 0x94,
	0xa5, 0xb6, 0x04, 0x79, 0x01, 0x78, 0x00, 0x84, 0x2c, 0xc4, 0xc0, 0x12, 0xa5, 0xce, 0x91, 0x46,
	0x4d, 0xed, 0x28, 0x71, 0x0b, 0x7d, 0x0b, 0x1e, 0x8b, 0xb1, 0x03, 0x03, 0x23, 0x4a, 0x5e, 0x04,
	0xe1, 0x24, 0xc2, 0x4c, 0x4c, 0x3e, 0xdf, 0xf7, 0xfb, 0x3e, 0x9d, 0xee, 0x30, 0x95, 0x4f, 0xb0,
	0x85, 0xca, 0x64, 0x2a, 0xe5, 0xdd, 0x1b, 0x99, 0x5d, 0x01, 0x15, 0x2b, 0x4a, 0x6d, 0x34, 0xf1,
	0xe5, 0x32, 0xce, 0x54, 0x08, 0x0a, 0xca, 0x74, 0xc7, 0x64, 0x08, 0xf6, 0xcf, 0x7e, 0x5c, 0xb3,
	0x05, 0x9e, 0x3c, 0xb4, 0xe5, 0xfd, 0xb7, 0x8f, 0x08, 0x3c, 0xfd, 0x15, 0xe4, 0x21, 0x7f, 0x18,
	0x8c, 0x2f, 0xe7, 0xec, 0xaf, 0x24, 0xe6, 0xc4, 0x88, 0xc9, 0xd6, 0xc9, 0x9c, 0xbd, 0x23, 0x3c,
	0x76, 0x54, 0x42, 0xf0, 0x48, 0xc5, 0x6b, 0xf0, 0x90, 0x8f, 0x82, 0x23, 0x61, 0x6b, 0x72, 0x8e,
	0xa7, 0xb9, 0x96, 0xab, 0x4d, 0x11, 0x15, 0x50, 0x66, 0x3a, 0xf1, 0xfe, 0xf9, 0x28, 0x18, 0x8a,
	0x49, 0xdb, 0xbc, 0xb3, 0x3d, 0x72, 0x81, 0x8f, 0xfb, 0xe1, 0x3a, 0x6a, 0x68, 0xa9, 0x7e, 0xe4,
	0x0e, 0x0b, 0xf1, 0x99, 0xd1, 0x2b, 0x50, 0x51, 0x09, 0x39, 0xc4, 0x95, 0x83, 0x8f, 0x2c, 0x7e,
	0x6a, 0x55, 0xd1, 0x8b, 0x9d, 0x8b, 0xe3, 0x93, 0x04, 0x72, 0x48, 0x63, 0x93, 0x69, 0x55, 0x45,
	0x71, 0x9e, 0xeb, 0x67, 0x48, 0xbc, 0xff, 0x3e, 0x0a, 0x0e, 0x05, 0x71, 0xa4, 0xeb, 0x56, 0xb9,
	0xb9, 0x7d, 0xab, 0x29, 0xda, 0xd7, 0x14, 0x7d, 0xd6, 0x14, 0xbd, 0x36, 0x74, 0xb0, 0x6f, 0xe8,
	0xe0, 0xa3, 0xa1, 0x83, 0xc7, 0x30, 0xcd, 0xcc, 0x72, 0xb3, 0x60, 0x52, 0xaf, 0xb9, 0xbb, 0x37,
	0x2e, 0x43, 0x98, 0xdb, 0x06, 0x7f, 0xe1, 0xce, 0xe9, 0xec, 0xa6, 0x17, 0x07, 0xf6, 0x66, 0x57,
	0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x27, 0x91, 0x42, 0x6c, 0xd5, 0x01, 0x00, 0x00,
}

func (m *VestingTypes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VestingTypes) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VestingTypes) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.VestingTypes) > 0 {
		for iNdEx := len(m.VestingTypes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VestingTypes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintVestingTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *VestingType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VestingType) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VestingType) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DelegationsAllowed {
		i--
		if m.DelegationsAllowed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if m.TokenReleasingPeriod != 0 {
		i = encodeVarintVestingTypes(dAtA, i, uint64(m.TokenReleasingPeriod))
		i--
		dAtA[i] = 0x20
	}
	if m.VestingPeriod != 0 {
		i = encodeVarintVestingTypes(dAtA, i, uint64(m.VestingPeriod))
		i--
		dAtA[i] = 0x18
	}
	if m.LockupPeriod != 0 {
		i = encodeVarintVestingTypes(dAtA, i, uint64(m.LockupPeriod))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintVestingTypes(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintVestingTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovVestingTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *VestingTypes) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.VestingTypes) > 0 {
		for _, e := range m.VestingTypes {
			l = e.Size()
			n += 1 + l + sovVestingTypes(uint64(l))
		}
	}
	return n
}

func (m *VestingType) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovVestingTypes(uint64(l))
	}
	if m.LockupPeriod != 0 {
		n += 1 + sovVestingTypes(uint64(m.LockupPeriod))
	}
	if m.VestingPeriod != 0 {
		n += 1 + sovVestingTypes(uint64(m.VestingPeriod))
	}
	if m.TokenReleasingPeriod != 0 {
		n += 1 + sovVestingTypes(uint64(m.TokenReleasingPeriod))
	}
	if m.DelegationsAllowed {
		n += 2
	}
	return n
}

func sovVestingTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVestingTypes(x uint64) (n int) {
	return sovVestingTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *VestingTypes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVestingTypes
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
			return fmt.Errorf("proto: VestingTypes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VestingTypes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VestingTypes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVestingTypes
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
				return ErrInvalidLengthVestingTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVestingTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VestingTypes = append(m.VestingTypes, &VestingType{})
			if err := m.VestingTypes[len(m.VestingTypes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVestingTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVestingTypes
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
func (m *VestingType) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVestingTypes
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
			return fmt.Errorf("proto: VestingType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VestingType: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVestingTypes
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
				return ErrInvalidLengthVestingTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVestingTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LockupPeriod", wireType)
			}
			m.LockupPeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVestingTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LockupPeriod |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VestingPeriod", wireType)
			}
			m.VestingPeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVestingTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VestingPeriod |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenReleasingPeriod", wireType)
			}
			m.TokenReleasingPeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVestingTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TokenReleasingPeriod |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegationsAllowed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVestingTypes
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
			m.DelegationsAllowed = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipVestingTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVestingTypes
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
func skipVestingTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVestingTypes
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
					return 0, ErrIntOverflowVestingTypes
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
					return 0, ErrIntOverflowVestingTypes
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
				return 0, ErrInvalidLengthVestingTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVestingTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVestingTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVestingTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVestingTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVestingTypes = fmt.Errorf("proto: unexpected end of group")
)
