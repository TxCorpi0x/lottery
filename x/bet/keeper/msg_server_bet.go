package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/vjdmhd/lottery/app/params"
	"github.com/vjdmhd/lottery/x/bet/types"
)

func (k msgServer) CreateBet(goCtx context.Context, msg *types.MsgCreateBet) (*types.MsgCreateBetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	min, max := k.lotteryKeeper.GetMinMaxBetAllowed(ctx)

	accAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return &types.MsgCreateBetResponse{},
			sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address: %s", msg.Creator)

	}

	if msg.Amount.LT(min) || msg.Amount.GT(max) {
		return &types.MsgCreateBetResponse{},
			sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "bet amount should be between %s and %s", min, max)
	}

	totalAmount := k.lotteryKeeper.GetLotteryFee(ctx).Add(msg.Amount)
	totalCoin := sdk.NewCoin(params.DefaultBondDenom, totalAmount)

	if k.BankKeeper.SpendableCoins(ctx, accAddr).AmountOf(params.DefaultBondDenom).LT(totalAmount) {
		return &types.MsgCreateBetResponse{},
			sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, "account balance is less than sum of amount and fees %s", totalAmount)
	}

	k.lotteryKeeper.TransferFeesAndAmount(ctx, accAddr, totalCoin)

	// get current number of bets to create incremental id
	betStats := k.getBetStats(ctx)
	betStats.Count += 1

	// create active bet object to be set in store
	bet := types.NewBet(betStats.Count, msg.Creator, msg.Amount, ctx.BlockHeight())
	k.SetActiveBet(ctx, bet)

	// update the counts of the bets
	k.setBetStats(ctx, betStats)

	return &types.MsgCreateBetResponse{}, nil
}
