package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/vjdmhd/lottery/x/bet/types"
)

func (k msgServer) CreateBet(goCtx context.Context, msg *types.MsgCreateBet) (*types.MsgCreateBetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: get current lottery id from lottery module
	// Check if the value already exists
	_, isFound := k.GetBet(
		ctx,
		"TODO",
		msg.Creator,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var bet = types.Bet{
		Creator: msg.Creator,
		Amount:  msg.Amount,
	}

	k.SetBet(
		ctx,
		bet,
	)
	return &types.MsgCreateBetResponse{}, nil
}
