// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cfeairdrop/account.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	types1 "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
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

// Period defines a length of time and amount of coins that will vest.
type ContinuousVestingPeriod struct {
	StartTime int64                                    `protobuf:"varint,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime   int64                                    `protobuf:"varint,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Amount    github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount"`
}

func (m *ContinuousVestingPeriod) Reset()      { *m = ContinuousVestingPeriod{} }
func (*ContinuousVestingPeriod) ProtoMessage() {}
func (*ContinuousVestingPeriod) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdc2cd1d846a008b, []int{0}
}
func (m *ContinuousVestingPeriod) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContinuousVestingPeriod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContinuousVestingPeriod.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContinuousVestingPeriod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContinuousVestingPeriod.Merge(m, src)
}
func (m *ContinuousVestingPeriod) XXX_Size() int {
	return m.Size()
}
func (m *ContinuousVestingPeriod) XXX_DiscardUnknown() {
	xxx_messageInfo_ContinuousVestingPeriod.DiscardUnknown(m)
}

var xxx_messageInfo_ContinuousVestingPeriod proto.InternalMessageInfo

func (m *ContinuousVestingPeriod) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *ContinuousVestingPeriod) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *ContinuousVestingPeriod) GetAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Amount
	}
	return nil
}

// PeriodicVestingAccount implements the VestingAccount interface. It
// periodically vests by unlocking coins during each specified period.
type AirdropVestingAccount struct {
	*types1.BaseVestingAccount `protobuf:"bytes,1,opt,name=base_vesting_account,json=baseVestingAccount,proto3,embedded=base_vesting_account" json:"base_vesting_account,omitempty"`
	StartTime                  int64                     `protobuf:"varint,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	VestingPeriods             []ContinuousVestingPeriod `protobuf:"bytes,3,rep,name=vesting_periods,json=vestingPeriods,proto3" json:"vesting_periods"`
}

func (m *AirdropVestingAccount) Reset()      { *m = AirdropVestingAccount{} }
func (*AirdropVestingAccount) ProtoMessage() {}
func (*AirdropVestingAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdc2cd1d846a008b, []int{1}
}
func (m *AirdropVestingAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AirdropVestingAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AirdropVestingAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AirdropVestingAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AirdropVestingAccount.Merge(m, src)
}
func (m *AirdropVestingAccount) XXX_Size() int {
	return m.Size()
}
func (m *AirdropVestingAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_AirdropVestingAccount.DiscardUnknown(m)
}

var xxx_messageInfo_AirdropVestingAccount proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ContinuousVestingPeriod)(nil), "chain4energy.c4echain.cfeairdrop.ContinuousVestingPeriod")
	proto.RegisterType((*AirdropVestingAccount)(nil), "chain4energy.c4echain.cfeairdrop.AirdropVestingAccount")
}

func init() { proto.RegisterFile("cfeairdrop/account.proto", fileDescriptor_bdc2cd1d846a008b) }

var fileDescriptor_bdc2cd1d846a008b = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x31, 0x8f, 0xd3, 0x30,
	0x18, 0x86, 0x93, 0xb6, 0x3a, 0x0e, 0x9f, 0x04, 0x52, 0x74, 0x88, 0xde, 0x49, 0x24, 0x55, 0xc5,
	0x50, 0x21, 0xd5, 0xa6, 0x25, 0x0b, 0xdd, 0x9a, 0xee, 0x08, 0x45, 0x88, 0x81, 0xa5, 0x72, 0x1c,
	0x93, 0x5a, 0x28, 0x76, 0x14, 0x3b, 0x15, 0xfd, 0x07, 0x4c, 0x88, 0x91, 0xb1, 0x33, 0x3f, 0x82,
	0xb9, 0x63, 0x47, 0xa6, 0x82, 0xda, 0x3f, 0x82, 0x62, 0x3b, 0x10, 0x8a, 0x10, 0x53, 0xe2, 0xef,
	0x7d, 0xed, 0xef, 0xfd, 0x1e, 0x1b, 0xf4, 0xc9, 0x5b, 0x8a, 0x59, 0x99, 0x96, 0xa2, 0x40, 0x98,
	0x10, 0x51, 0x71, 0x05, 0x8b, 0x52, 0x28, 0xe1, 0x0d, 0xc8, 0x0a, 0x33, 0x1e, 0x52, 0x4e, 0xcb,
	0x6c, 0x03, 0x49, 0x48, 0xf5, 0x1a, 0xfe, 0xf6, 0xdf, 0x5e, 0x67, 0x22, 0x13, 0xda, 0x8c, 0xea,
	0x3f, 0xb3, 0xef, 0xd6, 0x27, 0x42, 0xe6, 0x42, 0xa2, 0x04, 0x4b, 0x8a, 0xd6, 0x93, 0x84, 0x2a,
	0x3c, 0x41, 0x44, 0x30, 0x6e, 0xf5, 0xc7, 0x56, 0x5f, 0x53, 0xa9, 0x18, 0xcf, 0x7e, 0x59, 0xec,
	0xda, 0xb8, 0x86, 0x5f, 0x5d, 0xf0, 0x70, 0x21, 0xb8, 0x62, 0xbc, 0x12, 0x95, 0x7c, 0x6d, 0xb4,
	0x97, 0xb4, 0x64, 0x22, 0xf5, 0x1e, 0x01, 0x20, 0x15, 0x2e, 0xd5, 0x52, 0xb1, 0x9c, 0xf6, 0xdd,
	0x81, 0x3b, 0xea, 0xc6, 0x77, 0x75, 0xe5, 0x15, 0xcb, 0xa9, 0x77, 0x03, 0x2e, 0x29, 0x4f, 0x8d,
	0xd8, 0xd1, 0xe2, 0x1d, 0xca, 0x53, 0x2d, 0x11, 0x70, 0x81, 0xf3, 0x7a, 0xc6, 0x7e, 0x77, 0xd0,
	0x1d, 0x5d, 0x4d, 0x6f, 0xa0, 0x09, 0x03, 0xeb, 0xb0, 0xd0, 0x26, 0x81, 0x0b, 0xc1, 0x78, 0xf4,
	0x74, 0x77, 0x08, 0x9c, 0x2f, 0xdf, 0x83, 0x51, 0xc6, 0xd4, 0xaa, 0x4a, 0x20, 0x11, 0x39, 0xb2,
	0xc9, 0xcd, 0x67, 0x2c, 0xd3, 0x77, 0x48, 0x6d, 0x0a, 0x2a, 0xf5, 0x06, 0x19, 0xdb, 0xa3, 0x67,
	0xbd, 0xcf, 0xdb, 0xc0, 0x19, 0x7e, 0xec, 0x80, 0x07, 0x73, 0x03, 0xca, 0xa6, 0x9f, 0x1b, 0xbc,
	0x5e, 0x02, 0xae, 0xeb, 0x76, 0x4b, 0x3b, 0xf0, 0xd2, 0x62, 0xd7, 0x83, 0x5c, 0x4d, 0x9f, 0x34,
	0x91, 0x1a, 0x1e, 0x4d, 0xaa, 0x08, 0x4b, 0xfa, 0xe7, 0x49, 0x51, 0x6f, 0x7f, 0x08, 0xdc, 0xd8,
	0x4b, 0xfe, 0x52, 0xce, 0x10, 0x75, 0xce, 0x11, 0xad, 0xc0, 0xfd, 0xa6, 0x7b, 0xa1, 0x99, 0x4a,
	0x0b, 0xe4, 0x39, 0xfc, 0xdf, 0xad, 0xc3, 0x7f, 0xdc, 0x4a, 0xd4, 0xab, 0x81, 0xc5, 0xf7, 0xd6,
	0xed, 0xa2, 0x9c, 0x5d, 0x7e, 0xd8, 0x06, 0x4e, 0x0d, 0x24, 0x7a, 0xb1, 0x3b, 0xfa, 0xee, 0xfe,
	0xe8, 0xbb, 0x3f, 0x8e, 0xbe, 0xfb, 0xe9, 0xe4, 0x3b, 0xfb, 0x93, 0xef, 0x7c, 0x3b, 0xf9, 0xce,
	0x9b, 0xb0, 0x8d, 0xb8, 0xd5, 0x1e, 0x91, 0x90, 0x8e, 0x75, 0x01, 0xbd, 0x47, 0xad, 0x77, 0xaa,
	0xa1, 0x27, 0x17, 0xfa, 0xa1, 0x3c, 0xfb, 0x19, 0x00, 0x00, 0xff, 0xff, 0x15, 0xa1, 0x64, 0x4d,
	0xc2, 0x02, 0x00, 0x00,
}

func (m *ContinuousVestingPeriod) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContinuousVestingPeriod) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContinuousVestingPeriod) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAccount(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.EndTime != 0 {
		i = encodeVarintAccount(dAtA, i, uint64(m.EndTime))
		i--
		dAtA[i] = 0x10
	}
	if m.StartTime != 0 {
		i = encodeVarintAccount(dAtA, i, uint64(m.StartTime))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *AirdropVestingAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AirdropVestingAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AirdropVestingAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.VestingPeriods) > 0 {
		for iNdEx := len(m.VestingPeriods) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VestingPeriods[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAccount(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.StartTime != 0 {
		i = encodeVarintAccount(dAtA, i, uint64(m.StartTime))
		i--
		dAtA[i] = 0x10
	}
	if m.BaseVestingAccount != nil {
		{
			size, err := m.BaseVestingAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAccount(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAccount(dAtA []byte, offset int, v uint64) int {
	offset -= sovAccount(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ContinuousVestingPeriod) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StartTime != 0 {
		n += 1 + sovAccount(uint64(m.StartTime))
	}
	if m.EndTime != 0 {
		n += 1 + sovAccount(uint64(m.EndTime))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovAccount(uint64(l))
		}
	}
	return n
}

func (m *AirdropVestingAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BaseVestingAccount != nil {
		l = m.BaseVestingAccount.Size()
		n += 1 + l + sovAccount(uint64(l))
	}
	if m.StartTime != 0 {
		n += 1 + sovAccount(uint64(m.StartTime))
	}
	if len(m.VestingPeriods) > 0 {
		for _, e := range m.VestingPeriods {
			l = e.Size()
			n += 1 + l + sovAccount(uint64(l))
		}
	}
	return n
}

func sovAccount(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAccount(x uint64) (n int) {
	return sovAccount(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ContinuousVestingPeriod) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccount
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
			return fmt.Errorf("proto: ContinuousVestingPeriod: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContinuousVestingPeriod: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			m.StartTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			m.EndTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
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
				return ErrInvalidLengthAccount
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAccount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccount
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
func (m *AirdropVestingAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccount
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
			return fmt.Errorf("proto: AirdropVestingAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AirdropVestingAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseVestingAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
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
				return ErrInvalidLengthAccount
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BaseVestingAccount == nil {
				m.BaseVestingAccount = &types1.BaseVestingAccount{}
			}
			if err := m.BaseVestingAccount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			m.StartTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VestingPeriods", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
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
				return ErrInvalidLengthAccount
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VestingPeriods = append(m.VestingPeriods, ContinuousVestingPeriod{})
			if err := m.VestingPeriods[len(m.VestingPeriods)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAccount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccount
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
func skipAccount(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAccount
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
					return 0, ErrIntOverflowAccount
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
					return 0, ErrIntOverflowAccount
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
				return 0, ErrInvalidLengthAccount
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAccount
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAccount
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAccount        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAccount          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAccount = fmt.Errorf("proto: unexpected end of group")
)