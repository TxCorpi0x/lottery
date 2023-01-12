package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/vjdmhd/lottery/x/bet/types"
)

func CmdListActiveBet() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-active-bet",
		Short: "list all active bet",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllActiveBetRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ActiveBetAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowActiveBet() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-active-bet [creator]",
		Short: "shows an active bet",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argCreator := args[0]

			params := &types.QueryGetBetRequest{
				Creator: argCreator,
			}

			res, err := queryClient.ActiveBet(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdListSettledBet() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-settled-bet",
		Short: "list all settled bet",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			argLotteryID := cast.ToUint64(args[0])

			params := &types.QueryAllSettledBetRequest{
				Pagination: pageReq,
				LotteryId:  argLotteryID,
			}

			res, err := queryClient.SettledBetAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
