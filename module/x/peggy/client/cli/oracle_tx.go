package cli

import (
	"bufio"
	"encoding/base64"
	"strconv"

	"github.com/althea-net/peggy/module/x/peggy/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/spf13/cobra"
)

func GetObservedCmd(cdc *codec.Codec) *cobra.Command {
	testingTxCmd := &cobra.Command{
		Use:                        "observed",
		Short:                      "submit observed ETH events",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	testingTxCmd.AddCommand(flags.PostCommands(
		CmdSendETHDepositRequest(cdc),
		CmdSendETHWithdrawalRequest(cdc),
		CmdSendETHMultiSigRequest(cdc),
	)...)

	return testingTxCmd
}

func CmdSendETHDepositRequest(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "deposit [eth chain id] [eth contract address] [nonce] [cosmos receiver] [amount] [eth erc20 symbol] [eth erc20 contract addr] [eth sender address]",
		Short: "Submit an eth event observed by an orchestrator",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cosmosAddr := cliCtx.GetFromAddress()

			ethChainID := args[0]
			ethContractAddress := args[1]
			nonce, err := parseNonce(args[2])
			if err != nil {
				return err
			}
			receiverAddr, err := sdk.AccAddressFromBech32(args[3])
			if err != nil {
				return sdkerrors.Wrap(err, "cosmos receiver")
			}
			amount, err := strconv.ParseInt(args[4], 10, 64)
			if err != nil {
				return sdkerrors.Wrap(err, "amount")
			}
			tokenSymbol := args[5]

			// Make the message
			tokenContractAddr := types.NewEthereumAddress(args[6])
			ethSenderAddr := types.NewEthereumAddress(args[7])
			msg := types.MsgCreateEthereumClaims{
				EthereumChainID:       ethChainID,
				BridgeContractAddress: types.NewEthereumAddress(ethContractAddress),
				Orchestrator:          cosmosAddr,
				Claims: []types.EthereumClaim{
					types.EthereumBridgeDepositClaim{
						Nonce:          nonce,
						ERC20Token:     types.NewERC20Token(uint64(amount), tokenSymbol, tokenContractAddr),
						EthereumSender: ethSenderAddr,
						CosmosReceiver: receiverAddr,
					},
				},
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func CmdSendETHWithdrawalRequest(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "withdrawal [eth chain id] [eth contract address] [nonce]",
		Short: "Submit an eth event observed by an orchestrator",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cosmosAddr := cliCtx.GetFromAddress()

			ethChainID := args[0]
			ethContractAddress := args[1]
			nonce, err := parseNonce(args[2])
			if err != nil {
				return err
			}
			msg := types.MsgCreateEthereumClaims{
				EthereumChainID:       ethChainID,
				BridgeContractAddress: types.NewEthereumAddress(ethContractAddress),
				Orchestrator:          cosmosAddr,
				Claims: []types.EthereumClaim{
					types.EthereumBridgeWithdrawalBatchClaim{
						Nonce: nonce,
					},
				},
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func CmdSendETHMultiSigRequest(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "multisig-update [eth chain id] [eth contract address] [nonce]",
		Short: "Submit an eth event observed by an orchestrator",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cosmosAddr := cliCtx.GetFromAddress()

			ethChainID := args[0]
			ethContractAddress := types.NewEthereumAddress(args[1])
			nonce, err := parseNonce(args[2])
			if err != nil {
				return err
			}
			msg := types.MsgCreateEthereumClaims{
				EthereumChainID:       ethChainID,
				BridgeContractAddress: ethContractAddress,
				Orchestrator:          cosmosAddr,
				Claims: []types.EthereumClaim{
					types.EthereumBridgeMultiSigUpdateClaim{
						Nonce: nonce,
					},
				},
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// todo: we mix nonces as int64 and base64 bytes at the moment
func parseNonce(nonceArg string) (types.Nonce, error) {
	if len(nonceArg) != base64.StdEncoding.EncodedLen(8) {
		// not a byte nonce byte representation
		v, err := strconv.ParseUint(nonceArg, 10, 64)
		if err != nil {
			return nil, sdkerrors.Wrap(err, "nonce")
		}
		return types.NonceFromUint64(v), nil
	}
	return base64.StdEncoding.DecodeString(nonceArg)
}
