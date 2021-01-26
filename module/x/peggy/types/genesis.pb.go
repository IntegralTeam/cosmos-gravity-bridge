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
	PeggyId                       string                                 `protobuf:"bytes,1,opt,name=peggy_id,json=peggyId,proto3" json:"peggy_id,omitempty"`
	ContractSourceHash            string                                 `protobuf:"bytes,2,opt,name=contract_source_hash,json=contractSourceHash,proto3" json:"contract_source_hash,omitempty"`
	StartThreshold                uint64                                 `protobuf:"varint,3,opt,name=start_threshold,json=startThreshold,proto3" json:"start_threshold,omitempty"`
	EthereumAddress               string                                 `protobuf:"bytes,4,opt,name=ethereum_address,json=ethereumAddress,proto3" json:"ethereum_address,omitempty"`
	BridgeChainId                 uint64                                 `protobuf:"varint,5,opt,name=bridge_chain_id,json=bridgeChainId,proto3" json:"bridge_chain_id,omitempty"`
	SignedValsetsWindow           uint64                                 `protobuf:"varint,6,opt,name=signed_valsets_window,json=signedValsetsWindow,proto3" json:"signed_valsets_window,omitempty"`
	SignedBatchesWindow           uint64                                 `protobuf:"varint,7,opt,name=signed_batches_window,json=signedBatchesWindow,proto3" json:"signed_batches_window,omitempty"`
	SignedClaimsWindow            uint64                                 `protobuf:"varint,8,opt,name=signed_claims_window,json=signedClaimsWindow,proto3" json:"signed_claims_window,omitempty"`
	TargetBatchTimeout            uint64                                 `protobuf:"varint,10,opt,name=target_batch_timeout,json=targetBatchTimeout,proto3" json:"target_batch_timeout,omitempty"`
	AverageBlockTime              uint64                                 `protobuf:"varint,11,opt,name=average_block_time,json=averageBlockTime,proto3" json:"average_block_time,omitempty"`
	AverageEthereumBlockTime      uint64                                 `protobuf:"varint,12,opt,name=average_ethereum_block_time,json=averageEthereumBlockTime,proto3" json:"average_ethereum_block_time,omitempty"`
	SlashFractionValset           github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,13,opt,name=slash_fraction_valset,json=slashFractionValset,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"slash_fraction_valset"`
	SlashFractionBatch            github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,14,opt,name=slash_fraction_batch,json=slashFractionBatch,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"slash_fraction_batch"`
	SlashFractionClaim            github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,15,opt,name=slash_fraction_claim,json=slashFractionClaim,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"slash_fraction_claim"`
	SlashFractionConflictingClaim github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,16,opt,name=slash_fraction_conflicting_claim,json=slashFractionConflictingClaim,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"slash_fraction_conflicting_claim"`
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

func (m *Params) GetSignedValsetsWindow() uint64 {
	if m != nil {
		return m.SignedValsetsWindow
	}
	return 0
}

func (m *Params) GetSignedBatchesWindow() uint64 {
	if m != nil {
		return m.SignedBatchesWindow
	}
	return 0
}

func (m *Params) GetSignedClaimsWindow() uint64 {
	if m != nil {
		return m.SignedClaimsWindow
	}
	return 0
}

func (m *Params) GetTargetBatchTimeout() uint64 {
	if m != nil {
		return m.TargetBatchTimeout
	}
	return 0
}

func (m *Params) GetAverageBlockTime() uint64 {
	if m != nil {
		return m.AverageBlockTime
	}
	return 0
}

func (m *Params) GetAverageEthereumBlockTime() uint64 {
	if m != nil {
		return m.AverageEthereumBlockTime
	}
	return 0
}

// GenesisState struct
type GenesisState struct {
	Params         *Params             `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
	Valsets        []*Valset           `protobuf:"bytes,2,rep,name=valsets,proto3" json:"valsets,omitempty"`
	ValsetConfirms []*MsgValsetConfirm `protobuf:"bytes,3,rep,name=valset_confirms,json=valsetConfirms,proto3" json:"valset_confirms,omitempty"`
	Batches        []*OutgoingTxBatch  `protobuf:"bytes,4,rep,name=batches,proto3" json:"batches,omitempty"`
	BatchConfirms  []MsgConfirmBatch   `protobuf:"bytes,5,rep,name=batch_confirms,json=batchConfirms,proto3" json:"batch_confirms"`
	Attestations   []Attestation       `protobuf:"bytes,6,rep,name=attestations,proto3" json:"attestations"`
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

func (m *GenesisState) GetValsets() []*Valset {
	if m != nil {
		return m.Valsets
	}
	return nil
}

func (m *GenesisState) GetValsetConfirms() []*MsgValsetConfirm {
	if m != nil {
		return m.ValsetConfirms
	}
	return nil
}

func (m *GenesisState) GetBatches() []*OutgoingTxBatch {
	if m != nil {
		return m.Batches
	}
	return nil
}

func (m *GenesisState) GetBatchConfirms() []MsgConfirmBatch {
	if m != nil {
		return m.BatchConfirms
	}
	return nil
}

func (m *GenesisState) GetAttestations() []Attestation {
	if m != nil {
		return m.Attestations
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "peggy.v1.Params")
	proto.RegisterType((*GenesisState)(nil), "peggy.v1.GenesisState")
}

func init() { proto.RegisterFile("peggy/v1/genesis.proto", fileDescriptor_84231c3b3f050761) }

var fileDescriptor_84231c3b3f050761 = []byte{
	// 723 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xcf, 0x4e, 0x1b, 0x49,
	0x10, 0xc6, 0x6d, 0x30, 0x36, 0xdb, 0x18, 0xdb, 0x6a, 0xcc, 0xaa, 0xf1, 0x6a, 0x8d, 0xc5, 0x81,
	0xf5, 0xae, 0x16, 0x9b, 0x3f, 0xb7, 0x48, 0x51, 0x84, 0x9d, 0x90, 0x10, 0x09, 0x25, 0x1a, 0x50,
	0x22, 0xe5, 0x32, 0x69, 0xcf, 0x34, 0x3d, 0x2d, 0x66, 0xa6, 0xad, 0xe9, 0xb6, 0x81, 0x5b, 0x1e,
	0x21, 0xaf, 0x92, 0xb7, 0xe0, 0xc8, 0x31, 0x8a, 0x22, 0x14, 0xc1, 0x21, 0xaf, 0x11, 0xb9, 0xba,
	0x3d, 0xb6, 0x49, 0x4e, 0x28, 0x27, 0x4f, 0xd7, 0x57, 0xbf, 0xfa, 0xc6, 0x55, 0x53, 0x8d, 0xfe,
	0xec, 0x33, 0xce, 0x2f, 0xdb, 0xc3, 0x9d, 0x36, 0x67, 0x31, 0x53, 0x42, 0xb5, 0xfa, 0x89, 0xd4,
	0x12, 0x2f, 0x42, 0xbc, 0x35, 0xdc, 0xa9, 0x55, 0xb9, 0xe4, 0x12, 0x82, 0xed, 0xd1, 0x93, 0xd1,
	0x6b, 0xd5, 0x94, 0xd3, 0x97, 0x7d, 0x66, 0xa9, 0xda, 0x4a, 0x1a, 0x8d, 0x14, 0x57, 0x3f, 0xa5,
	0xf6, 0xa8, 0xf6, 0x02, 0x1b, 0xad, 0xa5, 0x51, 0xaa, 0x35, 0x53, 0x9a, 0x6a, 0x21, 0x63, 0xa3,
	0x6d, 0x7c, 0x2a, 0xa0, 0xfc, 0x6b, 0x9a, 0xd0, 0x48, 0xe1, 0x35, 0x64, 0xde, 0xc4, 0x15, 0x3e,
	0xc9, 0x36, 0xb2, 0xcd, 0x3f, 0x9c, 0x02, 0x9c, 0x0f, 0x7d, 0xbc, 0x8d, 0xaa, 0x9e, 0x8c, 0x75,
	0x42, 0x3d, 0xed, 0x2a, 0x39, 0x48, 0x3c, 0xe6, 0x06, 0x54, 0x05, 0x64, 0x0e, 0xd2, 0xf0, 0x58,
	0x3b, 0x06, 0xe9, 0x05, 0x55, 0x01, 0xfe, 0x07, 0x95, 0x95, 0xa6, 0x89, 0x76, 0x75, 0x90, 0x30,
	0x15, 0xc8, 0xd0, 0x27, 0xf3, 0x8d, 0x6c, 0x33, 0xe7, 0x94, 0x20, 0x7c, 0x32, 0x8e, 0xe2, 0x7f,
	0x51, 0x85, 0xe9, 0x80, 0x25, 0x6c, 0x10, 0xb9, 0xd4, 0xf7, 0x13, 0xa6, 0x14, 0xc9, 0x41, 0xd9,
	0xf2, 0x38, 0xbe, 0x6f, 0xc2, 0x78, 0x13, 0x95, 0x7b, 0x89, 0xf0, 0x39, 0x73, 0xbd, 0x80, 0x8a,
	0x78, 0xf4, 0x9e, 0x0b, 0x50, 0x73, 0xd9, 0x84, 0xbb, 0xa3, 0xe8, 0xa1, 0x8f, 0x77, 0xd1, 0xaa,
	0x12, 0x3c, 0x66, 0xbe, 0x3b, 0xa4, 0xa1, 0x62, 0x5a, 0xb9, 0xe7, 0x22, 0xf6, 0xe5, 0x39, 0xc9,
	0x43, 0xf6, 0x8a, 0x11, 0xdf, 0x18, 0xed, 0x2d, 0x48, 0x53, 0x0c, 0x74, 0x8e, 0xa5, 0x4c, 0x61,
	0x9a, 0xe9, 0x18, 0xcd, 0x32, 0xdb, 0xa8, 0x6a, 0x19, 0x2f, 0xa4, 0x22, 0x4a, 0x91, 0x45, 0x40,
	0xb0, 0xd1, 0xba, 0x20, 0x4d, 0x08, 0x4d, 0x13, 0xce, 0xb4, 0x71, 0x71, 0xb5, 0x88, 0x98, 0x1c,
	0x68, 0x82, 0x0c, 0x61, 0x34, 0x30, 0x39, 0x31, 0x0a, 0xfe, 0x1f, 0x61, 0x3a, 0x64, 0x09, 0xe5,
	0xcc, 0xed, 0x85, 0xd2, 0x3b, 0x03, 0x84, 0x2c, 0x41, 0x7e, 0xc5, 0x2a, 0x9d, 0x91, 0x30, 0x02,
	0xf0, 0x63, 0xf4, 0xd7, 0x38, 0x3b, 0x6d, 0xea, 0x14, 0x56, 0x04, 0x8c, 0xd8, 0x94, 0x67, 0x36,
	0x63, 0x82, 0xf7, 0xd0, 0xaa, 0x0a, 0xa9, 0x0a, 0xdc, 0xd3, 0xd1, 0x34, 0x85, 0x8c, 0x6d, 0x03,
	0xc9, 0x72, 0x23, 0xdb, 0x2c, 0x76, 0x5a, 0x57, 0x37, 0xeb, 0x99, 0x2f, 0x37, 0xeb, 0x9b, 0x5c,
	0xe8, 0x60, 0xd0, 0x6b, 0x79, 0x32, 0x6a, 0x7b, 0x52, 0x45, 0x52, 0xd9, 0x9f, 0x2d, 0xe5, 0x9f,
	0xd9, 0x8f, 0xf4, 0x29, 0xf3, 0x9c, 0x15, 0x28, 0x76, 0x60, 0x6b, 0x99, 0x7e, 0xe3, 0xf7, 0xa8,
	0x7a, 0xcf, 0x03, 0x5a, 0x41, 0x4a, 0x0f, 0xb2, 0xc0, 0x33, 0x16, 0xd0, 0xb9, 0x5f, 0x38, 0xc0,
	0x78, 0x48, 0xf9, 0x37, 0x38, 0xc0, 0x34, 0xf1, 0x39, 0x6a, 0xdc, 0x77, 0x90, 0xf1, 0x69, 0x28,
	0x3c, 0x2d, 0x62, 0x6e, 0xdd, 0x2a, 0x0f, 0x72, 0xfb, 0x7b, 0xd6, 0x6d, 0x52, 0x15, 0x8c, 0x1f,
	0xe5, 0x3e, 0x7c, 0x6d, 0x64, 0x36, 0xbe, 0xcf, 0xa1, 0xe2, 0x73, 0x73, 0x85, 0x1c, 0x6b, 0xaa,
	0x19, 0x6e, 0xa2, 0x7c, 0x1f, 0x76, 0x18, 0xf6, 0x76, 0x69, 0xb7, 0xd2, 0x1a, 0x5f, 0x29, 0x2d,
	0xb3, 0xdb, 0x8e, 0xd5, 0xf1, 0x7f, 0xa8, 0x60, 0x77, 0x82, 0xcc, 0x35, 0xe6, 0x67, 0x53, 0xcd,
	0x80, 0x9c, 0x71, 0x02, 0xee, 0xa2, 0xb2, 0x79, 0x84, 0x7f, 0x27, 0x92, 0x48, 0x91, 0x79, 0x60,
	0x6a, 0x13, 0xe6, 0x48, 0x71, 0x83, 0x75, 0x4d, 0x8a, 0x53, 0x1a, 0x4e, 0x1f, 0x15, 0xde, 0x43,
	0x05, 0xbb, 0x50, 0x24, 0x07, 0xf0, 0xda, 0x04, 0x7e, 0x35, 0xd0, 0x5c, 0x8a, 0x98, 0x9f, 0x5c,
	0xc0, 0xe0, 0x9c, 0x71, 0x26, 0x3e, 0x40, 0x25, 0xb3, 0x1f, 0xa9, 0xf1, 0xc2, 0x7d, 0xf6, 0x48,
	0x71, 0xeb, 0x01, 0x6c, 0x27, 0x37, 0x6a, 0xb4, 0xb3, 0x0c, 0x58, 0x6a, 0xfe, 0x04, 0x15, 0xa7,
	0x6e, 0x3c, 0x45, 0xf2, 0x50, 0x65, 0x75, 0x52, 0x65, 0x7f, 0xa2, 0xda, 0x0a, 0x33, 0x40, 0xe7,
	0xe5, 0xd5, 0x6d, 0x3d, 0x7b, 0x7d, 0x5b, 0xcf, 0x7e, 0xbb, 0xad, 0x67, 0x3f, 0xde, 0xd5, 0x33,
	0xd7, 0x77, 0xf5, 0xcc, 0xe7, 0xbb, 0x7a, 0xe6, 0xdd, 0xf6, 0xd4, 0x40, 0x69, 0xa8, 0x03, 0x46,
	0xb7, 0x62, 0xa6, 0xdb, 0xe6, 0xa6, 0x8d, 0xa4, 0x3f, 0x08, 0x59, 0xfb, 0xc2, 0x1e, 0x61, 0xbc,
	0xbd, 0x3c, 0x5c, 0xb8, 0x7b, 0x3f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa6, 0x07, 0x94, 0x19, 0x07,
	0x06, 0x00, 0x00,
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
		size := m.SlashFractionConflictingClaim.Size()
		i -= size
		if _, err := m.SlashFractionConflictingClaim.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1
	i--
	dAtA[i] = 0x82
	{
		size := m.SlashFractionClaim.Size()
		i -= size
		if _, err := m.SlashFractionClaim.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x7a
	{
		size := m.SlashFractionBatch.Size()
		i -= size
		if _, err := m.SlashFractionBatch.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x72
	{
		size := m.SlashFractionValset.Size()
		i -= size
		if _, err := m.SlashFractionValset.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x6a
	if m.AverageEthereumBlockTime != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.AverageEthereumBlockTime))
		i--
		dAtA[i] = 0x60
	}
	if m.AverageBlockTime != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.AverageBlockTime))
		i--
		dAtA[i] = 0x58
	}
	if m.TargetBatchTimeout != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.TargetBatchTimeout))
		i--
		dAtA[i] = 0x50
	}
	if m.SignedClaimsWindow != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.SignedClaimsWindow))
		i--
		dAtA[i] = 0x40
	}
	if m.SignedBatchesWindow != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.SignedBatchesWindow))
		i--
		dAtA[i] = 0x38
	}
	if m.SignedValsetsWindow != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.SignedValsetsWindow))
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
	if len(m.Attestations) > 0 {
		for iNdEx := len(m.Attestations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Attestations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.BatchConfirms) > 0 {
		for iNdEx := len(m.BatchConfirms) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BatchConfirms[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Batches) > 0 {
		for iNdEx := len(m.Batches) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Batches[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.ValsetConfirms) > 0 {
		for iNdEx := len(m.ValsetConfirms) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ValsetConfirms[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Valsets) > 0 {
		for iNdEx := len(m.Valsets) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Valsets[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
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
	if m.SignedValsetsWindow != 0 {
		n += 1 + sovGenesis(uint64(m.SignedValsetsWindow))
	}
	if m.SignedBatchesWindow != 0 {
		n += 1 + sovGenesis(uint64(m.SignedBatchesWindow))
	}
	if m.SignedClaimsWindow != 0 {
		n += 1 + sovGenesis(uint64(m.SignedClaimsWindow))
	}
	if m.TargetBatchTimeout != 0 {
		n += 1 + sovGenesis(uint64(m.TargetBatchTimeout))
	}
	if m.AverageBlockTime != 0 {
		n += 1 + sovGenesis(uint64(m.AverageBlockTime))
	}
	if m.AverageEthereumBlockTime != 0 {
		n += 1 + sovGenesis(uint64(m.AverageEthereumBlockTime))
	}
	l = m.SlashFractionValset.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.SlashFractionBatch.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.SlashFractionClaim.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.SlashFractionConflictingClaim.Size()
	n += 2 + l + sovGenesis(uint64(l))
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
	if len(m.Valsets) > 0 {
		for _, e := range m.Valsets {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ValsetConfirms) > 0 {
		for _, e := range m.ValsetConfirms {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Batches) > 0 {
		for _, e := range m.Batches {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.BatchConfirms) > 0 {
		for _, e := range m.BatchConfirms {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Attestations) > 0 {
		for _, e := range m.Attestations {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
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
				return fmt.Errorf("proto: wrong wireType = %d for field SignedValsetsWindow", wireType)
			}
			m.SignedValsetsWindow = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SignedValsetsWindow |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignedBatchesWindow", wireType)
			}
			m.SignedBatchesWindow = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SignedBatchesWindow |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignedClaimsWindow", wireType)
			}
			m.SignedClaimsWindow = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SignedClaimsWindow |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TargetBatchTimeout", wireType)
			}
			m.TargetBatchTimeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TargetBatchTimeout |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AverageBlockTime", wireType)
			}
			m.AverageBlockTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AverageBlockTime |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AverageEthereumBlockTime", wireType)
			}
			m.AverageEthereumBlockTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AverageEthereumBlockTime |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
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
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashFractionBatch", wireType)
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
			if err := m.SlashFractionBatch.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashFractionClaim", wireType)
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
			if err := m.SlashFractionClaim.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashFractionConflictingClaim", wireType)
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
			if err := m.SlashFractionConflictingClaim.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Valsets", wireType)
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
			m.Valsets = append(m.Valsets, &Valset{})
			if err := m.Valsets[len(m.Valsets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValsetConfirms", wireType)
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
			m.ValsetConfirms = append(m.ValsetConfirms, &MsgValsetConfirm{})
			if err := m.ValsetConfirms[len(m.ValsetConfirms)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Batches", wireType)
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
			m.Batches = append(m.Batches, &OutgoingTxBatch{})
			if err := m.Batches[len(m.Batches)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BatchConfirms", wireType)
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
			m.BatchConfirms = append(m.BatchConfirms, MsgConfirmBatch{})
			if err := m.BatchConfirms[len(m.BatchConfirms)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Attestations", wireType)
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
			m.Attestations = append(m.Attestations, Attestation{})
			if err := m.Attestations[len(m.Attestations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
