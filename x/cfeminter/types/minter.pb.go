// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: c4echain/cfeminter/minter.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Minter struct {
	SequenceId uint32     `protobuf:"varint,1,opt,name=sequence_id,json=sequenceId,proto3" json:"sequence_id,omitempty"`
	EndTime    *time.Time `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3,stdtime" json:"end_time,omitempty"`
	Config     *types.Any `protobuf:"bytes,3,opt,name=config,proto3" json:"config,omitempty"`
}

func (m *Minter) Reset()      { *m = Minter{} }
func (*Minter) ProtoMessage() {}
func (*Minter) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b192c5e4cd70467, []int{0}
}
func (m *Minter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Minter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Minter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Minter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Minter.Merge(m, src)
}
func (m *Minter) XXX_Size() int {
	return m.Size()
}
func (m *Minter) XXX_DiscardUnknown() {
	xxx_messageInfo_Minter.DiscardUnknown(m)
}

var xxx_messageInfo_Minter proto.InternalMessageInfo

func (m *Minter) GetSequenceId() uint32 {
	if m != nil {
		return m.SequenceId
	}
	return 0
}

func (m *Minter) GetEndTime() *time.Time {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *Minter) GetConfig() *types.Any {
	if m != nil {
		return m.Config
	}
	return nil
}

type NoMinting struct {
}

func (m *NoMinting) Reset()         { *m = NoMinting{} }
func (m *NoMinting) String() string { return proto.CompactTextString(m) }
func (*NoMinting) ProtoMessage()    {}
func (*NoMinting) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b192c5e4cd70467, []int{1}
}
func (m *NoMinting) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NoMinting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NoMinting.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NoMinting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NoMinting.Merge(m, src)
}
func (m *NoMinting) XXX_Size() int {
	return m.Size()
}
func (m *NoMinting) XXX_DiscardUnknown() {
	xxx_messageInfo_NoMinting.DiscardUnknown(m)
}

var xxx_messageInfo_NoMinting proto.InternalMessageInfo

type LinearMinting struct {
	Amount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount"`
}

func (m *LinearMinting) Reset()         { *m = LinearMinting{} }
func (m *LinearMinting) String() string { return proto.CompactTextString(m) }
func (*LinearMinting) ProtoMessage()    {}
func (*LinearMinting) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b192c5e4cd70467, []int{2}
}
func (m *LinearMinting) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LinearMinting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LinearMinting.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LinearMinting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LinearMinting.Merge(m, src)
}
func (m *LinearMinting) XXX_Size() int {
	return m.Size()
}
func (m *LinearMinting) XXX_DiscardUnknown() {
	xxx_messageInfo_LinearMinting.DiscardUnknown(m)
}

var xxx_messageInfo_LinearMinting proto.InternalMessageInfo

type ExponentialStepMinting struct {
	StepDuration     time.Duration                          `protobuf:"bytes,1,opt,name=step_duration,json=stepDuration,proto3,stdduration" json:"step_duration"`
	Amount           github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount"`
	AmountMultiplier github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=amount_multiplier,json=amountMultiplier,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"amount_multiplier"`
}

func (m *ExponentialStepMinting) Reset()         { *m = ExponentialStepMinting{} }
func (m *ExponentialStepMinting) String() string { return proto.CompactTextString(m) }
func (*ExponentialStepMinting) ProtoMessage()    {}
func (*ExponentialStepMinting) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b192c5e4cd70467, []int{3}
}
func (m *ExponentialStepMinting) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExponentialStepMinting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExponentialStepMinting.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExponentialStepMinting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExponentialStepMinting.Merge(m, src)
}
func (m *ExponentialStepMinting) XXX_Size() int {
	return m.Size()
}
func (m *ExponentialStepMinting) XXX_DiscardUnknown() {
	xxx_messageInfo_ExponentialStepMinting.DiscardUnknown(m)
}

var xxx_messageInfo_ExponentialStepMinting proto.InternalMessageInfo

func (m *ExponentialStepMinting) GetStepDuration() time.Duration {
	if m != nil {
		return m.StepDuration
	}
	return 0
}

type MinterState struct {
	SequenceId                  uint32                                 `protobuf:"varint,1,opt,name=sequence_id,json=sequenceId,proto3" json:"sequence_id,omitempty"`
	AmountMinted                github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=amount_minted,json=amountMinted,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount_minted"`
	RemainderToMint             github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=remainder_to_mint,json=remainderToMint,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"remainder_to_mint"`
	LastMintBlockTime           time.Time                              `protobuf:"bytes,4,opt,name=last_mint_block_time,json=lastMintBlockTime,proto3,stdtime" json:"last_mint_block_time"`
	RemainderFromPreviousMinter github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=remainder_from_previous_minter,json=remainderFromPreviousMinter,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"remainder_from_previous_minter"`
}

func (m *MinterState) Reset()         { *m = MinterState{} }
func (m *MinterState) String() string { return proto.CompactTextString(m) }
func (*MinterState) ProtoMessage()    {}
func (*MinterState) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b192c5e4cd70467, []int{4}
}
func (m *MinterState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MinterState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MinterState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MinterState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MinterState.Merge(m, src)
}
func (m *MinterState) XXX_Size() int {
	return m.Size()
}
func (m *MinterState) XXX_DiscardUnknown() {
	xxx_messageInfo_MinterState.DiscardUnknown(m)
}

var xxx_messageInfo_MinterState proto.InternalMessageInfo

func (m *MinterState) GetSequenceId() uint32 {
	if m != nil {
		return m.SequenceId
	}
	return 0
}

func (m *MinterState) GetLastMintBlockTime() time.Time {
	if m != nil {
		return m.LastMintBlockTime
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*Minter)(nil), "chain4energy.c4echain.cfeminter.Minter")
	proto.RegisterType((*NoMinting)(nil), "chain4energy.c4echain.cfeminter.NoMinting")
	proto.RegisterType((*LinearMinting)(nil), "chain4energy.c4echain.cfeminter.LinearMinting")
	proto.RegisterType((*ExponentialStepMinting)(nil), "chain4energy.c4echain.cfeminter.ExponentialStepMinting")
	proto.RegisterType((*MinterState)(nil), "chain4energy.c4echain.cfeminter.MinterState")
}

func init() { proto.RegisterFile("c4echain/cfeminter/minter.proto", fileDescriptor_7b192c5e4cd70467) }

var fileDescriptor_7b192c5e4cd70467 = []byte{
	// 587 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0x8e, 0xdb, 0x10, 0xda, 0x4b, 0x23, 0x88, 0x15, 0xa1, 0x24, 0x48, 0x76, 0x95, 0x01, 0x75,
	0x89, 0x2d, 0xb5, 0x99, 0x60, 0x40, 0x98, 0x52, 0x11, 0x89, 0x20, 0x94, 0x14, 0x21, 0x95, 0xc1,
	0x72, 0xec, 0x17, 0xf7, 0x54, 0xfb, 0xce, 0x9c, 0xcf, 0xa8, 0xf9, 0x05, 0xac, 0x1d, 0x3b, 0xb2,
	0x23, 0x36, 0x7e, 0x44, 0xc5, 0xd4, 0x11, 0x31, 0x14, 0x94, 0xfc, 0x11, 0xe4, 0xbb, 0x73, 0x88,
	0xda, 0xa1, 0x22, 0x4c, 0x3e, 0xbf, 0xfb, 0xbe, 0xef, 0x7d, 0xef, 0xdd, 0xd3, 0x43, 0xa6, 0xdf,
	0x03, 0xff, 0xd8, 0xc3, 0xc4, 0xf6, 0x27, 0x10, 0x63, 0xc2, 0x81, 0xd9, 0xf2, 0x63, 0x25, 0x8c,
	0x72, 0xaa, 0x9b, 0xe2, 0xb6, 0x07, 0x04, 0x58, 0x38, 0xb5, 0x0a, 0xb4, 0xb5, 0x40, 0xb7, 0x5b,
	0x21, 0xa5, 0x61, 0x04, 0xb6, 0x80, 0x8f, 0xb3, 0x89, 0xed, 0x91, 0xa9, 0xe4, 0xb6, 0x1b, 0x21,
	0x0d, 0xa9, 0x38, 0xda, 0xf9, 0x49, 0x45, 0xcd, 0xeb, 0x04, 0x8e, 0x63, 0x48, 0xb9, 0x17, 0x27,
	0x0a, 0x60, 0x5c, 0x07, 0x04, 0x19, 0xf3, 0x38, 0xa6, 0x44, 0xdd, 0xb7, 0x7c, 0x9a, 0xc6, 0x34,
	0x75, 0xa5, 0xb2, 0xfc, 0x91, 0x57, 0x9d, 0xaf, 0x1a, 0xaa, 0x0c, 0x84, 0x2f, 0xdd, 0x44, 0xd5,
	0x14, 0x3e, 0x64, 0x40, 0x7c, 0x70, 0x71, 0xd0, 0xd4, 0xb6, 0xb5, 0x9d, 0xda, 0x10, 0x15, 0xa1,
	0x7e, 0xa0, 0x3f, 0x41, 0x1b, 0x40, 0x02, 0x37, 0xcf, 0xde, 0x5c, 0xdb, 0xd6, 0x76, 0xaa, 0xbb,
	0x6d, 0x4b, 0x66, 0xb6, 0x8a, 0xcc, 0xd6, 0x61, 0x61, 0xcd, 0x29, 0x9f, 0xfd, 0x32, 0xb5, 0xe1,
	0x5d, 0x20, 0x41, 0x1e, 0xd3, 0x9f, 0xa2, 0x8a, 0x4f, 0xc9, 0x04, 0x87, 0xcd, 0x75, 0x41, 0x6d,
	0xdc, 0xa0, 0x3e, 0x23, 0x53, 0xa7, 0xfe, 0xfd, 0x5b, 0xb7, 0x26, 0xfd, 0x3c, 0x17, 0xe8, 0xfe,
	0x50, 0xd1, 0x1e, 0x97, 0xcf, 0x3f, 0x9b, 0xa5, 0x4e, 0x15, 0x6d, 0xbe, 0xa6, 0x39, 0x00, 0x93,
	0xb0, 0xf3, 0x0e, 0xd5, 0x5e, 0x61, 0x02, 0x1e, 0x53, 0x01, 0xfd, 0x00, 0x55, 0xbc, 0x98, 0x66,
	0x84, 0x0b, 0xf7, 0x9b, 0x8e, 0x75, 0x71, 0x65, 0x96, 0x7e, 0x5e, 0x99, 0x8f, 0x42, 0xcc, 0x8f,
	0xb3, 0xb1, 0xe5, 0xd3, 0x58, 0x95, 0xaf, 0x3e, 0xdd, 0x34, 0x38, 0xb1, 0xf9, 0x34, 0x81, 0xd4,
	0xea, 0x13, 0x3e, 0x54, 0xec, 0xce, 0xa7, 0x35, 0xf4, 0xe0, 0xc5, 0x69, 0x42, 0x09, 0x10, 0x8e,
	0xbd, 0x68, 0xc4, 0x21, 0x29, 0x52, 0xbc, 0x44, 0xb5, 0x94, 0x43, 0xe2, 0x16, 0x2d, 0x16, 0x99,
	0xaa, 0xbb, 0xad, 0x1b, 0xe5, 0xec, 0x2b, 0x80, 0xb3, 0x91, 0x9b, 0x38, 0xcf, 0x9b, 0xb1, 0x95,
	0x33, 0x8b, 0xf8, 0x92, 0xd9, 0xb5, 0xff, 0x31, 0xab, 0xbf, 0x47, 0x75, 0x79, 0x72, 0xe3, 0x2c,
	0xe2, 0x38, 0x89, 0x30, 0x30, 0xd1, 0xe4, 0x7f, 0x93, 0xdc, 0x07, 0x7f, 0x78, 0x5f, 0x0a, 0x0d,
	0x16, 0x3a, 0x9d, 0x2f, 0xeb, 0xa8, 0x2a, 0xdf, 0x63, 0xc4, 0x3d, 0x0e, 0xb7, 0x0f, 0xc9, 0x08,
	0xd5, 0x0a, 0x37, 0x39, 0x2d, 0x58, 0xb1, 0xb8, 0x2d, 0xe5, 0x44, 0x68, 0xe8, 0x47, 0xa8, 0xce,
	0x20, 0xf6, 0x30, 0x09, 0x80, 0xb9, 0x9c, 0x0a, 0xe9, 0x15, 0x4b, 0xbc, 0xb7, 0x10, 0x3a, 0x14,
	0x73, 0xa4, 0xbf, 0x45, 0x8d, 0xc8, 0x4b, 0xa5, 0x5d, 0x77, 0x1c, 0x51, 0xff, 0x44, 0x4e, 0x78,
	0xf9, 0xd6, 0x09, 0x17, 0x0f, 0x2b, 0xa6, 0xbc, 0x9e, 0x2b, 0xe4, 0x6a, 0x4e, 0xce, 0x17, 0xf3,
	0x9e, 0x22, 0xe3, 0xaf, 0xe5, 0x09, 0xa3, 0xb1, 0x9b, 0x30, 0xf8, 0x88, 0x69, 0x96, 0xca, 0xc6,
	0xb0, 0xe6, 0x9d, 0x95, 0xfc, 0x3f, 0x5c, 0xa8, 0x1e, 0x30, 0x1a, 0xbf, 0x51, 0x9a, 0xf2, 0x89,
	0x9c, 0xc1, 0xc5, 0xcc, 0xd0, 0x2e, 0x67, 0x86, 0xf6, 0x7b, 0x66, 0x68, 0x67, 0x73, 0xa3, 0x74,
	0x39, 0x37, 0x4a, 0x3f, 0xe6, 0x46, 0xe9, 0x68, 0x6f, 0x59, 0x7e, 0x69, 0x41, 0xd9, 0x7e, 0x0f,
	0xba, 0x72, 0x9f, 0x9d, 0x2e, 0x6d, 0x34, 0x91, 0x6f, 0x5c, 0x11, 0x45, 0xef, 0xfd, 0x09, 0x00,
	0x00, 0xff, 0xff, 0x11, 0x80, 0xeb, 0x55, 0xf4, 0x04, 0x00, 0x00,
}

func (m *Minter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Minter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Minter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Config != nil {
		{
			size, err := m.Config.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMinter(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.EndTime != nil {
		n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.EndTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.EndTime):])
		if err2 != nil {
			return 0, err2
		}
		i -= n2
		i = encodeVarintMinter(dAtA, i, uint64(n2))
		i--
		dAtA[i] = 0x12
	}
	if m.SequenceId != 0 {
		i = encodeVarintMinter(dAtA, i, uint64(m.SequenceId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *NoMinting) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NoMinting) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NoMinting) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *LinearMinting) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LinearMinting) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LinearMinting) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMinter(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ExponentialStepMinting) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExponentialStepMinting) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExponentialStepMinting) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.AmountMultiplier.Size()
		i -= size
		if _, err := m.AmountMultiplier.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMinter(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMinter(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	n3, err3 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.StepDuration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.StepDuration):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintMinter(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *MinterState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MinterState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MinterState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.RemainderFromPreviousMinter.Size()
		i -= size
		if _, err := m.RemainderFromPreviousMinter.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMinter(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	n4, err4 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.LastMintBlockTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.LastMintBlockTime):])
	if err4 != nil {
		return 0, err4
	}
	i -= n4
	i = encodeVarintMinter(dAtA, i, uint64(n4))
	i--
	dAtA[i] = 0x22
	{
		size := m.RemainderToMint.Size()
		i -= size
		if _, err := m.RemainderToMint.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMinter(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.AmountMinted.Size()
		i -= size
		if _, err := m.AmountMinted.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMinter(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.SequenceId != 0 {
		i = encodeVarintMinter(dAtA, i, uint64(m.SequenceId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMinter(dAtA []byte, offset int, v uint64) int {
	offset -= sovMinter(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Minter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SequenceId != 0 {
		n += 1 + sovMinter(uint64(m.SequenceId))
	}
	if m.EndTime != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.EndTime)
		n += 1 + l + sovMinter(uint64(l))
	}
	if m.Config != nil {
		l = m.Config.Size()
		n += 1 + l + sovMinter(uint64(l))
	}
	return n
}

func (m *NoMinting) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *LinearMinting) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Amount.Size()
	n += 1 + l + sovMinter(uint64(l))
	return n
}

func (m *ExponentialStepMinting) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.StepDuration)
	n += 1 + l + sovMinter(uint64(l))
	l = m.Amount.Size()
	n += 1 + l + sovMinter(uint64(l))
	l = m.AmountMultiplier.Size()
	n += 1 + l + sovMinter(uint64(l))
	return n
}

func (m *MinterState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SequenceId != 0 {
		n += 1 + sovMinter(uint64(m.SequenceId))
	}
	l = m.AmountMinted.Size()
	n += 1 + l + sovMinter(uint64(l))
	l = m.RemainderToMint.Size()
	n += 1 + l + sovMinter(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.LastMintBlockTime)
	n += 1 + l + sovMinter(uint64(l))
	l = m.RemainderFromPreviousMinter.Size()
	n += 1 + l + sovMinter(uint64(l))
	return n
}

func sovMinter(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMinter(x uint64) (n int) {
	return sovMinter(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Minter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMinter
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
			return fmt.Errorf("proto: Minter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Minter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SequenceId", wireType)
			}
			m.SequenceId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMinter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SequenceId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMinter
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
				return ErrInvalidLengthMinter
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMinter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.EndTime == nil {
				m.EndTime = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.EndTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Config", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMinter
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
				return ErrInvalidLengthMinter
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMinter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Config == nil {
				m.Config = &types.Any{}
			}
			if err := m.Config.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMinter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMinter
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
func (m *NoMinting) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMinter
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
			return fmt.Errorf("proto: NoMinting: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NoMinting: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMinter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMinter
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
func (m *LinearMinting) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMinter
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
			return fmt.Errorf("proto: LinearMinting: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LinearMinting: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMinter
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
				return ErrInvalidLengthMinter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMinter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMinter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMinter
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
func (m *ExponentialStepMinting) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMinter
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
			return fmt.Errorf("proto: ExponentialStepMinting: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExponentialStepMinting: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StepDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMinter
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
				return ErrInvalidLengthMinter
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMinter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.StepDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMinter
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
				return ErrInvalidLengthMinter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMinter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountMultiplier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMinter
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
				return ErrInvalidLengthMinter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMinter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AmountMultiplier.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMinter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMinter
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
func (m *MinterState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMinter
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
			return fmt.Errorf("proto: MinterState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MinterState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SequenceId", wireType)
			}
			m.SequenceId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMinter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SequenceId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountMinted", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMinter
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
				return ErrInvalidLengthMinter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMinter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AmountMinted.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemainderToMint", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMinter
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
				return ErrInvalidLengthMinter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMinter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RemainderToMint.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastMintBlockTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMinter
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
				return ErrInvalidLengthMinter
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMinter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.LastMintBlockTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemainderFromPreviousMinter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMinter
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
				return ErrInvalidLengthMinter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMinter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RemainderFromPreviousMinter.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMinter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMinter
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
func skipMinter(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMinter
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
					return 0, ErrIntOverflowMinter
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
					return 0, ErrIntOverflowMinter
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
				return 0, ErrInvalidLengthMinter
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMinter
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMinter
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMinter        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMinter          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMinter = fmt.Errorf("proto: unexpected end of group")
)
