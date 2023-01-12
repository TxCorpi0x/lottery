package keeper

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/vjdmhd/lottery/app/params"
	betmoduletypes "github.com/vjdmhd/lottery/x/bet/types"
	"github.com/vjdmhd/lottery/x/lottery/types"
)

func (k Keeper) CalculateAndTransferPayout(ctx sdk.Context, winnerBet betmoduletypes.Bet, bets []betmoduletypes.Bet) sdk.Coin {
	highestBetAmount := math.NewInt(0)
	totalAmount := math.NewInt(0)

	// set the first amount as the lowest, this is beting updated in the next loop
	lowestBetAmount := bets[0].Amount

	// loop through all active bets
	for _, v := range bets {

		if v.Amount.GT(highestBetAmount) {
			// update the highest bet amount
			highestBetAmount = v.Amount
		}

		if v.Amount.LT(lowestBetAmount) {
			// update the lowest bet amount
			lowestBetAmount = v.Amount
		}

		// update total amount paid in this lottery
		totalAmount = totalAmount.Add(v.Amount)
	}

	if winnerBet.Amount == highestBetAmount {
		// winner has paid the highest amount
		poolBalance := k.GetPoolBalance(ctx)
		coins := sdk.NewCoins(poolBalance)
		// transfer all of pool balance including fees to the winner account
		// payout the winner
		k.BankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.MustAccAddressFromBech32(winnerBet.Creator), coins)
		return poolBalance
	} else if winnerBet.Amount == lowestBetAmount {
		// bet has paid the lowest amount
		// pool balance remains the same because
		// winner will not get any reward
		return sdk.NewCoin(params.DefaultBondDenom, sdk.NewInt(0))
	} else {
		// transfer total amount of this lottery to the winner balance
		coin := sdk.NewCoin(params.DefaultBondDenom, totalAmount)

		// payuot the winner
		k.BankKeeper.SendCoinsFromModuleToAccount(
			ctx,
			types.ModuleName,
			sdk.MustAccAddressFromBech32(winnerBet.Creator),
			sdk.NewCoins(coin))
		return coin
	}
}
