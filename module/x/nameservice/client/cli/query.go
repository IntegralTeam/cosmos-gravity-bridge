package cli

import (
	"fmt"

	"github.com/althea-net/peggy/module/x/nameservice/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	nameserviceQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the nameservice module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	nameserviceQueryCmd.AddCommand(flags.GetCommands(
		// GetCmdResolveName(storeKey, cdc),
		// GetCmdWhois(storeKey, cdc),
		// GetCmdNames(storeKey, cdc),
		CmdGetCurrentValset(storeKey, cdc),
		CmdGetValsetByNonce(storeKey, cdc),
		CmdGetConfirmationsByNonce(storeKey, cdc)
	)...)

	return nameserviceQueryCmd
}

// // GetCmdResolveName queries information about a name
// func GetCmdResolveName(storeKey string, cdc *codec.Codec) *cobra.Command {
// 	return &cobra.Command{
// 		Use:   "resolve [name]",
// 		Short: "resolve name",
// 		Args:  cobra.ExactArgs(1),
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			cliCtx := context.NewCLIContext().WithCodec(cdc)
// 			name := args[0]

// 			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/resolve/%s", storeKey, name), nil)
// 			if err != nil {
// 				fmt.Printf("could not resolve name - %s \n", name)
// 				return nil
// 			}

// 			var out types.QueryResResolve
// 			cdc.MustUnmarshalJSON(res, &out)
// 			return cliCtx.PrintOutput(out)
// 		},
// 	}
// }

// // GetCmdWhois queries information about a domain
// func GetCmdWhois(storeKey string, cdc *codec.Codec) *cobra.Command {
// 	return &cobra.Command{
// 		Use:   "whois [name]",
// 		Short: "Query whois info of name",
// 		Args:  cobra.ExactArgs(1),
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			cliCtx := context.NewCLIContext().WithCodec(cdc)
// 			name := args[0]

// 			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/whois/%s", storeKey, name), nil)
// 			if err != nil {
// 				fmt.Printf("could not resolve whois - %s \n", name)
// 				return nil
// 			}

// 			var out types.Whois
// 			cdc.MustUnmarshalJSON(res, &out)
// 			return cliCtx.PrintOutput(out)
// 		},
// 	}
// }

func CmdGetCurrentValset(storeKey string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "current-valset",
		Short: "Query current valset",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/valset", storeKey), nil)
			if err != nil {
				fmt.Printf("could not get valset")
				return nil
			}

			var out types.Valset
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func CmdGetValsetByNonce(storeKey string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "valset-by-nonce [nonce]",
		Short: "Get valset with a particular nonce",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			nonce := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/valsetByNonce/%s", storeKey, nonce), nil)
			if err != nil {
				fmt.Printf("could not get valset")
				return nil
			}

			var out types.Valset
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func CmdGetConfirmationsByNonce(storeKey string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "confirmations-by-nonce [nonce]",
		Short: "Get valset confirmations with a particular nonce",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			nonce := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/confirmationsByNonce", storeKey, nonce), nil)
			if err != nil {
				fmt.Printf("could not get valset")
				return nil
			}

			var out []types.MsgValsetConfirm
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

// // GetCmdNames queries a list of all names
// func GetCmdNames(storeKey string, cdc *codec.Codec) *cobra.Command {
// 	return &cobra.Command{
// 		Use:   "names",
// 		Short: "names",
// 		// Args:  cobra.ExactArgs(1),
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			cliCtx := context.NewCLIContext().WithCodec(cdc)

// 			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/names", storeKey), nil)
// 			if err != nil {
// 				fmt.Printf("could not get query names\n")
// 				return nil
// 			}

// 			var out types.QueryResNames
// 			cdc.MustUnmarshalJSON(res, &out)
// 			return cliCtx.PrintOutput(out)
// 		},
// 	}
// }
