package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/vjdmhd/lottery/x/bet/types"
)

func CmdListBet() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-bet",
		Short: "list all bet",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllBetRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.BetAll(context.Background(), params)
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

func CmdShowBet() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-bet [lotteryID] [creator]",
		Short: "shows a bet",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argLotteryID := args[0]
			argCreator := args[1]

			params := &types.QueryGetBetRequest{
				LotteryId: argLotteryID,
				Creator:   argCreator,
			}

			res, err := queryClient.Bet(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
