// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: peggy/v1/genesis.proto

package types

import (
	fmt "fmt"
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

// Params represent the peggy genesis and store parameters
// PEGGYID:
// a random 32 byte value to prevent signature reuse
// CONTRACTHASH:
// the code hash of a known good version of the Peggy contract
// solidity code. It will be used to verify exactly which version of the
// bridge will be deployed.
// STARTTHRESHOLD:
// the percentage of total voting power that must be online
// and participating in Peggy operations before a bridge can start operating
// BRIDGECONTRACTADDRESS:
// is address of the bridge contract on the Ethereum side
// BRIDGECHAINID:
// the unique identifier of the Ethereum chain
type Params struct {
	PeggyId             string                                 `protobuf:"bytes,1,opt,name=peggy_id,json=peggyId,proto3" json:"peggy_id,omitempty"`
	ContractSourceHash  string                                 `protobuf:"bytes,2,opt,name=contract_source_hash,json=contractSourceHash,proto3" json:"contract_source_hash,omitempty"`
	StartThreshold      uint64                                 `protobuf:"varint,3,opt,name=start_threshold,json=startThreshold,proto3" json:"start_threshold,omitempty"`
	EthereumAddress     string                                 `protobuf:"bytes,4,opt,name=ethereum_address,json=ethereumAddress,proto3" json:"ethereum_address,omitempty"`
	BridgeChainId       uint64                                 `protobuf:"varint,5,opt,name=bridge_chain_id,json=bridgeChainId,proto3" json:"bridge_chain_id,omitempty"`
	SignedBlocksWindow  uint64                                 `protobuf:"varint,6,opt,name=signed_blocks_window,json=signedBlocksWindow,proto3" json:"signed_blocks_window,omitempty"`
	SlashFractionValset github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,7,opt,name=slash_fraction_valset,json=slashFractionValset,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"slash_fraction_valset"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_84231c3b3f050761, []int{0}
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

func (m *Params) GetPeggyId() string {
	if m != nil {
		return m.PeggyId
	}
	return ""
}

func (m *Params) GetContractSourceHash() string {
	if m != nil {
		return m.ContractSourceHash
	}
	return ""
}

func (m *Params) GetStartThreshold() uint64 {
	if m != nil {
		return m.StartThreshold
	}
	return 0
}

func (m *Params) GetEthereumAddress() string {
	if m != nil {
		return m.EthereumAddress
	}
	return ""
}

func (m *Params) GetBridgeChainId() uint64 {
	if m != nil {
		return m.BridgeChainId
	}
	return 0
}

func (m *Params) GetSignedBlocksWindow() uint64 {
	if m != nil {
		return m.SignedBlocksWindow
	}
	return 0
}

// GenesisState struct
type GenesisState struct {
	Params *Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_84231c3b3f050761, []int{1}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() *Params {
	if m != nil {
		return m.Params
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "peggy.v1.Params")
	proto.RegisterType((*GenesisState)(nil), "peggy.v1.GenesisState")
}

func init() { proto.RegisterFile("peggy/v1/genesis.proto", fileDescriptor_84231c3b3f050761) }

var fileDescriptor_84231c3b3f050761 = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x92, 0x4f, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0xed, 0x12, 0x52, 0x58, 0x0a, 0xa9, 0x96, 0x82, 0x0c, 0x07, 0x37, 0xea, 0xa1, 0x84,
	0x43, 0xed, 0x16, 0x2e, 0x88, 0x1b, 0x01, 0xf1, 0xef, 0x84, 0x52, 0x04, 0x12, 0x97, 0xd5, 0xda,
	0x3b, 0x78, 0x57, 0xb5, 0xbd, 0xd1, 0xce, 0x26, 0xa5, 0x37, 0x3e, 0x02, 0x1f, 0xab, 0xc7, 0x1e,
	0x11, 0x42, 0x11, 0x4a, 0xce, 0x7c, 0x07, 0x94, 0xd9, 0x58, 0xe2, 0x64, 0xcf, 0xef, 0x8d, 0x9e,
	0x9e, 0xe7, 0x99, 0xdd, 0x9f, 0x42, 0x55, 0x5d, 0xe4, 0xf3, 0x93, 0xbc, 0x82, 0x16, 0xd0, 0x60,
	0x36, 0x75, 0xd6, 0x5b, 0x7e, 0x83, 0x78, 0x36, 0x3f, 0x79, 0xb8, 0x57, 0xd9, 0xca, 0x12, 0xcc,
	0xd7, 0x6f, 0x41, 0x3f, 0xf8, 0xbb, 0xc5, 0xfa, 0x1f, 0xa4, 0x93, 0x0d, 0xf2, 0x07, 0x2c, 0x2c,
	0x0b, 0xa3, 0x92, 0x78, 0x18, 0x8f, 0x6e, 0x4e, 0xb6, 0x69, 0x7e, 0xa7, 0xf8, 0x31, 0xdb, 0x2b,
	0x6d, 0xeb, 0x9d, 0x2c, 0xbd, 0x40, 0x3b, 0x73, 0x25, 0x08, 0x2d, 0x51, 0x27, 0x5b, 0xb4, 0xc6,
	0x3b, 0xed, 0x94, 0xa4, 0xb7, 0x12, 0x35, 0x7f, 0xc4, 0x06, 0xe8, 0xa5, 0xf3, 0xc2, 0x6b, 0x07,
	0xa8, 0x6d, 0xad, 0x92, 0x6b, 0xc3, 0x78, 0xd4, 0x9b, 0xdc, 0x21, 0xfc, 0xb1, 0xa3, 0xfc, 0x31,
	0xdb, 0x05, 0xaf, 0xc1, 0xc1, 0xac, 0x11, 0x52, 0x29, 0x07, 0x88, 0x49, 0x8f, 0x6c, 0x07, 0x1d,
	0x7f, 0x11, 0x30, 0x3f, 0x64, 0x83, 0xc2, 0x19, 0x55, 0x81, 0x28, 0xb5, 0x34, 0xed, 0x3a, 0xe7,
	0x75, 0xf2, 0xbc, 0x1d, 0xf0, 0xcb, 0x35, 0x0d, 0x69, 0xd1, 0x54, 0x2d, 0x28, 0x51, 0xd4, 0xb6,
	0x3c, 0x43, 0x71, 0x6e, 0x5a, 0x65, 0xcf, 0x93, 0x3e, 0x2d, 0xf3, 0xa0, 0x8d, 0x49, 0xfa, 0x4c,
	0x0a, 0x2f, 0xd8, 0x3d, 0xac, 0x25, 0x6a, 0xf1, 0x75, 0xfd, 0x19, 0xc6, 0xb6, 0x62, 0x2e, 0x6b,
	0x04, 0x9f, 0x6c, 0x0f, 0xe3, 0xd1, 0xce, 0x38, 0xbb, 0x5c, 0xec, 0x47, 0xbf, 0x16, 0xfb, 0x87,
	0x95, 0xf1, 0x7a, 0x56, 0x64, 0xa5, 0x6d, 0xf2, 0xd2, 0x62, 0x63, 0x71, 0xf3, 0x38, 0x42, 0x75,
	0x96, 0xfb, 0x8b, 0x29, 0x60, 0xf6, 0x0a, 0xca, 0xc9, 0x5d, 0x32, 0x7b, 0xbd, 0xf1, 0xfa, 0x44,
	0x56, 0xcf, 0x7b, 0xdf, 0x7f, 0x0f, 0xa3, 0x83, 0x67, 0x6c, 0xe7, 0x4d, 0x28, 0xe8, 0xd4, 0x4b,
	0x0f, 0x7c, 0xc4, 0xfa, 0x53, 0x3a, 0x3f, 0x9d, 0xfc, 0xd6, 0x93, 0xdd, 0xac, 0x2b, 0x2c, 0x0b,
	0xb5, 0x4c, 0x36, 0xfa, 0xf8, 0xfd, 0xe5, 0x32, 0x8d, 0xaf, 0x96, 0x69, 0xfc, 0x67, 0x99, 0xc6,
	0x3f, 0x56, 0x69, 0x74, 0xb5, 0x4a, 0xa3, 0x9f, 0xab, 0x34, 0xfa, 0x72, 0xfc, 0x5f, 0x2c, 0x59,
	0x7b, 0x0d, 0xf2, 0xa8, 0x05, 0x9f, 0x87, 0x3f, 0xa2, 0xb1, 0x6a, 0x56, 0x43, 0xfe, 0x6d, 0x33,
	0x52, 0xc8, 0xa2, 0x4f, 0xe5, 0x3f, 0xfd, 0x17, 0x00, 0x00, 0xff, 0xff, 0x4f, 0x93, 0x94, 0x36,
	0x36, 0x02, 0x00, 0x00,
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
	{
		size := m.SlashFractionValset.Size()
		i -= size
		if _, err := m.SlashFractionValset.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if m.SignedBlocksWindow != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.SignedBlocksWindow))
		i--
		dAtA[i] = 0x30
	}
	if m.BridgeChainId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.BridgeChainId))
		i--
		dAtA[i] = 0x28
	}
	if len(m.EthereumAddress) > 0 {
		i -= len(m.EthereumAddress)
		copy(dAtA[i:], m.EthereumAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.EthereumAddress)))
		i--
		dAtA[i] = 0x22
	}
	if m.StartThreshold != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.StartThreshold))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ContractSourceHash) > 0 {
		i -= len(m.ContractSourceHash)
		copy(dAtA[i:], m.ContractSourceHash)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ContractSourceHash)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PeggyId) > 0 {
		i -= len(m.PeggyId)
		copy(dAtA[i:], m.PeggyId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.PeggyId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Params != nil {
		{
			size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
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
	l = len(m.PeggyId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.ContractSourceHash)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.StartThreshold != 0 {
		n += 1 + sovGenesis(uint64(m.StartThreshold))
	}
	l = len(m.EthereumAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.BridgeChainId != 0 {
		n += 1 + sovGenesis(uint64(m.BridgeChainId))
	}
	if m.SignedBlocksWindow != 0 {
		n += 1 + sovGenesis(uint64(m.SignedBlocksWindow))
	}
	l = m.SlashFractionValset.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Params != nil {
		l = m.Params.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeggyId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeggyId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractSourceHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractSourceHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartThreshold", wireType)
			}
			m.StartThreshold = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartThreshold |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthereumAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EthereumAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BridgeChainId", wireType)
			}
			m.BridgeChainId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BridgeChainId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignedBlocksWindow", wireType)
			}
			m.SignedBlocksWindow = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SignedBlocksWindow |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashFractionValset", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SlashFractionValset.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Params == nil {
				m.Params = &Params{}
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
