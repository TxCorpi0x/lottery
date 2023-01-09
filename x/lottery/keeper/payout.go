package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/vjdmhd/lottery/app/params"
	betmoduletypes "github.com/vjdmhd/lottery/x/bet/types"
	"github.com/vjdmhd/lottery/x/lottery/types"
)

func (k Keeper) CalculateAndTransferPayout(ctx sdk.Context, winnerBet betmoduletypes.Bet, bets []betmoduletypes.Bet) sdk.Coin {
	var highestBetAmount, totalAmount uint64
	lowestBetAmount := bets[0].Amount
	for _, v := range bets {
		if v.Amount > highestBetAmount {
			highestBetAmount = v.Amount
		}

		if v.Amount < lowestBetAmount {
			lowestBetAmount = v.Amount
		}

		totalAmount += v.Amount
	}

	if winnerBet.Amount == highestBetAmount {
		poolBalance := k.GetPoolBalance(ctx)
		coins := sdk.NewCoins(poolBalance)
		// transfer all of pool balance including fees to the winner account
		k.BankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.MustAccAddressFromBech32(winnerBet.Creator), coins)
		return poolBalance
	} else if winnerBet.Amount == lowestBetAmount {
		// pool balance remains the same because
		// winner will not get any reward
		return sdk.NewCoin(params.DefaultBondDenom, sdk.NewInt(0))
	} else {
		// transfer total amount of this lottery to the winner balance
		coin := sdk.NewCoin(params.DefaultBondDenom, sdk.NewIntFromUint64(totalAmount))
		k.BankKeeper.SendCoinsFromModuleToAccount(
			ctx,
			types.ModuleName,
			sdk.MustAccAddressFromBech32(winnerBet.Creator),
			sdk.NewCoins(coin))
		return coin
	}
}
