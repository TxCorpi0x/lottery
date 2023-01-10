package keeper

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/vjdmhd/lottery/x/lottery/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramstore.GetParamSet(ctx, &params)
	return params
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

func (k Keeper) GetLotteryFee(ctx sdk.Context) sdkmath.Int {
	return k.GetParams(ctx).LotteryFee
}

func (k Keeper) GetMinMaxBetAllowed(ctx sdk.Context) (sdkmath.Int, sdkmath.Int) {
	betSize := k.GetParams(ctx).BetSize
	return betSize.MinBet, betSize.MaxBet
}
