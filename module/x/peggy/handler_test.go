package peggy

import (
	"bytes"
	"testing"
	"time"

	"github.com/althea-net/peggy/module/x/peggy/keeper"
	"github.com/althea-net/peggy/module/x/peggy/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandleCreateEthereumClaims(t *testing.T) {
	var (
		myOrchestratorAddr sdk.AccAddress = make([]byte, sdk.AddrLen)
		myCosmosAddr       sdk.AccAddress = bytes.Repeat([]byte{1}, 12)
		myValAddr                         = sdk.ValAddress(myOrchestratorAddr) // revisit when proper mapping is impl in keeper
		myNonce                           = types.NewUInt64Nonce(1)
		anyETHAddr                        = types.NewEthereumAddress("any-address")
		tokenETHAddr                      = types.NewEthereumAddress("any-erc20-token-addr")
		myBlockTime                       = time.Date(2020, 9, 14, 15, 20, 10, 0, time.UTC)
	)
	k, ctx, keepers := keeper.CreateTestEnv(t)
	k.StakingKeeper = keeper.NewStakingKeeperMock(myValAddr)
	h := NewHandler(k)

	msg := MsgCreateEthereumClaims{
		EthereumChainID:       "0",
		BridgeContractAddress: types.NewEthereumAddress(""),
		Orchestrator:          myOrchestratorAddr,
		Claims: []EthereumClaim{
			EthereumBridgeDepositClaim{
				Nonce: myNonce,
				ERC20Token: types.ERC20Token{
					Amount:               12,
					Symbol:               "ALX",
					TokenContractAddress: tokenETHAddr,
				},
				EthereumSender: anyETHAddr,
				CosmosReceiver: myCosmosAddr,
			},
		},
	}
	// when
	ctx = ctx.WithBlockTime(myBlockTime)
	_, err := h(ctx, msg)
	// then
	require.NoError(t, err)
	// and claim persisted
	claimFound := k.HasClaim(ctx, types.ClaimTypeEthereumBridgeDeposit, myNonce, myValAddr, msg.Claims[0].Details())
	assert.True(t, claimFound)
	// and attestation persisted
	a := k.GetAttestation(ctx, types.ClaimTypeEthereumBridgeDeposit, myNonce)
	require.NotNil(t, a)
	// and vouchers added to the account
	balance := keepers.BankKeeper.GetCoins(ctx, myCosmosAddr)
	assert.Equal(t, sdk.Coins{sdk.NewInt64Coin("peggy96dde7db38", 12)}, balance)

}
