package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/vjdmhd/lottery/x/bet/types"
)

func (k msgServer) CreateBet(goCtx context.Context, msg *types.MsgCreateBet) (*types.MsgCreateBetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var bet = types.Bet{
		Creator: msg.Creator,
		Amount:  msg.Amount,
		Height:  ctx.BlockHeight(),
	}

	k.SetActiveBet(
		ctx,
		bet,
	)
	return &types.MsgCreateBetResponse{}, nil
}
