package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cast"
	"github.com/vjdmhd/lottery/x/bet/types"
)

func (k msgServer) CreateBet(goCtx context.Context, msg *types.MsgCreateBet) (*types.MsgCreateBetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: validate minimal bet amount and fees
	// TODO: subtract fees

	betStats := k.GetBetStats(ctx)
	betStats.Count += 1

	var bet = types.Bet{
		Id:      cast.ToString(betStats.Count),
		Creator: msg.Creator,
		Amount:  msg.Amount,
		Height:  ctx.BlockHeight(),
	}

	k.SetActiveBet(
		ctx,
		bet,
	)

	k.SetBetStats(
		ctx,
		betStats,
	)

	return &types.MsgCreateBetResponse{}, nil
}
