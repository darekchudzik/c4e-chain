// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: energybank/token_params.proto

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

type TokenParams struct {
	Index          string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Name           string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	TradingCompany string `protobuf:"bytes,3,opt,name=tradingCompany,proto3" json:"tradingCompany,omitempty"`
	BurningTime    uint64 `protobuf:"varint,4,opt,name=burningTime,proto3" json:"burningTime,omitempty"`
	BurningType    string `protobuf:"bytes,5,opt,name=burningType,proto3" json:"burningType,omitempty"`
	SendPrice      uint64 `protobuf:"varint,6,opt,name=sendPrice,proto3" json:"sendPrice,omitempty"`
	MintAccount    string `protobuf:"bytes,7,opt,name=mintAccount,proto3" json:"mintAccount,omitempty"`
}

func (m *TokenParams) Reset()         { *m = TokenParams{} }
func (m *TokenParams) String() string { return proto.CompactTextString(m) }
func (*TokenParams) ProtoMessage()    {}
func (*TokenParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_70195a2a923b1281, []int{0}
}
func (m *TokenParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TokenParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TokenParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TokenParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenParams.Merge(m, src)
}
func (m *TokenParams) XXX_Size() int {
	return m.Size()
}
func (m *TokenParams) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenParams.DiscardUnknown(m)
}

var xxx_messageInfo_TokenParams proto.InternalMessageInfo

func (m *TokenParams) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *TokenParams) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TokenParams) GetTradingCompany() string {
	if m != nil {
		return m.TradingCompany
	}
	return ""
}

func (m *TokenParams) GetBurningTime() uint64 {
	if m != nil {
		return m.BurningTime
	}
	return 0
}

func (m *TokenParams) GetBurningType() string {
	if m != nil {
		return m.BurningType
	}
	return ""
}

func (m *TokenParams) GetSendPrice() uint64 {
	if m != nil {
		return m.SendPrice
	}
	return 0
}

func (m *TokenParams) GetMintAccount() string {
	if m != nil {
		return m.MintAccount
	}
	return ""
}

func init() {
	proto.RegisterType((*TokenParams)(nil), "chain4energy.c4echain.energybank.TokenParams")
}

func init() { proto.RegisterFile("energybank/token_params.proto", fileDescriptor_70195a2a923b1281) }

var fileDescriptor_70195a2a923b1281 = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xb1, 0x4e, 0xf3, 0x30,
	0x14, 0x85, 0xeb, 0xff, 0x4f, 0x8b, 0xea, 0x4a, 0x0c, 0x16, 0x83, 0x07, 0xb0, 0x22, 0x06, 0xd4,
	0x85, 0x64, 0x20, 0x2f, 0x00, 0xec, 0xa8, 0xaa, 0x3a, 0xb1, 0x20, 0xc7, 0xb9, 0x4a, 0xad, 0xca,
	0xd7, 0x91, 0xe3, 0x48, 0xcd, 0x5b, 0xf0, 0x58, 0x8c, 0x1d, 0x19, 0x51, 0xf2, 0x22, 0xa8, 0xce,
	0x50, 0x8b, 0xcd, 0xf7, 0xf3, 0xf9, 0xce, 0x70, 0xe8, 0x1d, 0x20, 0xb8, 0xba, 0x2f, 0x25, 0x1e,
	0x72, 0x6f, 0x0f, 0x80, 0x1f, 0x8d, 0x74, 0xd2, 0xb4, 0x59, 0xe3, 0xac, 0xb7, 0x2c, 0x55, 0x7b,
	0xa9, 0xb1, 0x98, 0x42, 0x99, 0x2a, 0x20, 0xdc, 0xd9, 0x45, 0xba, 0x1f, 0x09, 0x5d, 0xed, 0xce,
	0xe2, 0x26, 0x78, 0xec, 0x86, 0xce, 0x35, 0x56, 0x70, 0xe4, 0x24, 0x25, 0xeb, 0xe5, 0x76, 0x3a,
	0x18, 0xa3, 0x09, 0x4a, 0x03, 0xfc, 0x5f, 0x80, 0xe1, 0xcd, 0x1e, 0xe8, 0xb5, 0x77, 0xb2, 0xd2,
	0x58, 0xbf, 0x5a, 0xd3, 0x48, 0xec, 0xf9, 0xff, 0xf0, 0xfb, 0x87, 0xb2, 0x94, 0xae, 0xca, 0xce,
	0xa1, 0xc6, 0x7a, 0xa7, 0x0d, 0xf0, 0x24, 0x25, 0xeb, 0x64, 0x1b, 0xa3, 0x38, 0xd1, 0x37, 0xc0,
	0xe7, 0xa1, 0x26, 0x46, 0xec, 0x96, 0x2e, 0x5b, 0xc0, 0x6a, 0xe3, 0xb4, 0x02, 0xbe, 0x08, 0x0d,
	0x17, 0x70, 0xf6, 0x8d, 0x46, 0xff, 0xac, 0x94, 0xed, 0xd0, 0xf3, 0xab, 0xc9, 0x8f, 0xd0, 0xcb,
	0xdb, 0xd7, 0x20, 0xc8, 0x69, 0x10, 0xe4, 0x67, 0x10, 0xe4, 0x73, 0x14, 0xb3, 0xd3, 0x28, 0x66,
	0xdf, 0xa3, 0x98, 0xbd, 0x17, 0xb5, 0xf6, 0xfb, 0xae, 0xcc, 0x94, 0x35, 0x79, 0x3c, 0x56, 0xae,
	0x0a, 0x78, 0x0c, 0x20, 0x3f, 0xe6, 0xf1, 0xc8, 0x7d, 0x03, 0x6d, 0xb9, 0x08, 0xf3, 0x3e, 0xfd,
	0x06, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x64, 0x57, 0x1d, 0x7f, 0x01, 0x00, 0x00,
}

func (m *TokenParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TokenParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MintAccount) > 0 {
		i -= len(m.MintAccount)
		copy(dAtA[i:], m.MintAccount)
		i = encodeVarintTokenParams(dAtA, i, uint64(len(m.MintAccount)))
		i--
		dAtA[i] = 0x3a
	}
	if m.SendPrice != 0 {
		i = encodeVarintTokenParams(dAtA, i, uint64(m.SendPrice))
		i--
		dAtA[i] = 0x30
	}
	if len(m.BurningType) > 0 {
		i -= len(m.BurningType)
		copy(dAtA[i:], m.BurningType)
		i = encodeVarintTokenParams(dAtA, i, uint64(len(m.BurningType)))
		i--
		dAtA[i] = 0x2a
	}
	if m.BurningTime != 0 {
		i = encodeVarintTokenParams(dAtA, i, uint64(m.BurningTime))
		i--
		dAtA[i] = 0x20
	}
	if len(m.TradingCompany) > 0 {
		i -= len(m.TradingCompany)
		copy(dAtA[i:], m.TradingCompany)
		i = encodeVarintTokenParams(dAtA, i, uint64(len(m.TradingCompany)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTokenParams(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintTokenParams(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTokenParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovTokenParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TokenParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovTokenParams(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTokenParams(uint64(l))
	}
	l = len(m.TradingCompany)
	if l > 0 {
		n += 1 + l + sovTokenParams(uint64(l))
	}
	if m.BurningTime != 0 {
		n += 1 + sovTokenParams(uint64(m.BurningTime))
	}
	l = len(m.BurningType)
	if l > 0 {
		n += 1 + l + sovTokenParams(uint64(l))
	}
	if m.SendPrice != 0 {
		n += 1 + sovTokenParams(uint64(m.SendPrice))
	}
	l = len(m.MintAccount)
	if l > 0 {
		n += 1 + l + sovTokenParams(uint64(l))
	}
	return n
}

func sovTokenParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTokenParams(x uint64) (n int) {
	return sovTokenParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TokenParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTokenParams
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
			return fmt.Errorf("proto: TokenParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenParams
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
				return ErrInvalidLengthTokenParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTokenParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenParams
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
				return ErrInvalidLengthTokenParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTokenParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TradingCompany", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenParams
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
				return ErrInvalidLengthTokenParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTokenParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TradingCompany = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BurningTime", wireType)
			}
			m.BurningTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BurningTime |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BurningType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenParams
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
				return ErrInvalidLengthTokenParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTokenParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BurningType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SendPrice", wireType)
			}
			m.SendPrice = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SendPrice |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintAccount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenParams
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
				return ErrInvalidLengthTokenParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTokenParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MintAccount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTokenParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTokenParams
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
func skipTokenParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTokenParams
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
					return 0, ErrIntOverflowTokenParams
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
					return 0, ErrIntOverflowTokenParams
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
				return 0, ErrInvalidLengthTokenParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTokenParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTokenParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTokenParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTokenParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTokenParams = fmt.Errorf("proto: unexpected end of group")
)
