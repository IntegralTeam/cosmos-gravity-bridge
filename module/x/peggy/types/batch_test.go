package types

import (
	"encoding/hex"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOutgoingTxBatchCheckpointGold1(t *testing.T) {
	senderAddr, err := sdk.AccAddressFromHex("527FBEE652609AB150F0AEE9D61A2F76CFC4A73E")
	require.NoError(t, err)
	var (
		receiverAddr     = EthereumAddress{0xc7, 0x83, 0xdf, 0x8a, 0x85, 0xf, 0x42, 0xe7, 0xf7, 0xe5, 0x70, 0x13, 0x75, 0x9c, 0x28, 0x5c, 0xaa, 0x70, 0x1e, 0xb6}
		erc20Addr        = EthereumAddress{0x7c, 0x2c, 0x19, 0x5c, 0xd6, 0xd3, 0x4b, 0x8f, 0x84, 0x59, 0x92, 0xd3, 0x80, 0xaa, 0xdb, 0x27, 0x30, 0xbb, 0x9c, 0x6f}
		orchestratorAddr = EthereumAddress{0xb4, 0x62, 0x86, 0x4e, 0x39, 0x5d, 0x88, 0xd6, 0xbc, 0x7c, 0x5d, 0xd5, 0xf3, 0xf5, 0xeb, 0x4c, 0xc2, 0x59, 0x92, 0x55}
	)
	src := OutgoingTxBatch{
		Nonce:              1,
		Elements:           []OutgoingTransferTx{{ID: 0x1, Sender: senderAddr, DestAddress: receiverAddr, Amount: ERC20Token{Amount: 0x1, Symbol: "MAX", TokenContractAddress: EthereumAddress{0x7c, 0x2c, 0x19, 0x5c, 0xd6, 0xd3, 0x4b, 0x8f, 0x84, 0x59, 0x92, 0xd3, 0x80, 0xaa, 0xdb, 0x27, 0x30, 0xbb, 0x9c, 0x6f}}, BridgeFee: ERC20Token{Amount: 0x0, Symbol: "MAX", TokenContractAddress: EthereumAddress{0x7c, 0x2c, 0x19, 0x5c, 0xd6, 0xd3, 0x4b, 0x8f, 0x84, 0x59, 0x92, 0xd3, 0x80, 0xaa, 0xdb, 0x27, 0x30, 0xbb, 0x9c, 0x6f}}}},
		TotalFee:           ERC20Token{Amount: 0x0, Symbol: "MAX", TokenContractAddress: erc20Addr},
		BridgedDenominator: BridgedDenominator{TokenContractAddress: erc20Addr, Symbol: "MAX", CosmosVoucherDenom: "peggy39b512461b"},
		BatchStatus:        1,
	}
	v := Valset{
		Nonce:        10,
		Powers:       []uint64{4294967295},
		EthAddresses: []EthereumAddress{orchestratorAddr},
	}
	ourHash, err := src.GetCheckpoint(&v)
	require.NoError(t, err)

	// hash from bridge contract
	goldHash := "0x0764db7c6ba5bdbd3bc902049e67530747a61bfad7c80530268eb47b3514b5ac"[2:]
	assert.Equal(t, goldHash, hex.EncodeToString(ourHash))
}
