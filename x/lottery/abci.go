package lottery

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/vjdmhd/lottery/x/lottery/keeper"
	"github.com/vjdmhd/lottery/x/lottery/types"
)

func EndBlocker(ctx sdk.Context, k keeper.Keeper) {

	currentLottery := k.GetOrCreateCurrentLottery(ctx)

	allActiveBets := k.BetKeeper.GetAllActiveBet(ctx)
	betCount := uint64(len(allActiveBets))

	fmt.Printf("Block height: %d, Lottery ID: %d, Bet Count: %d\n", ctx.BlockHeight(), currentLottery.Id, betCount)

	// get consensus address of the current block proposer
	proposerConsAddr := sdk.ConsAddress(ctx.BlockHeader().ProposerAddress)

	// fond proposer validator
	proposerValidator := k.StakingKeeper.ValidatorByConsAddr(ctx, proposerConsAddr)

	// find proposer operator
	operator := proposerValidator.GetOperator()

	// check all active bets for any existing bet from proposer operator account
	for _, v := range allActiveBets {

		accAddr := sdk.MustAccAddressFromBech32(v.Creator)
		if operator.Equals(accAddr) {
			fmt.Printf("Block Proposer operator has a bet in the active bet list, continue to the next block.\nPool Balance: %s\n",
				k.GetPoolBalance(ctx))
			return
		}

	}

	// if the bets count does not satisfy the min count
	// should return and continue in nex end blocker
	if betCount < types.MinBetCount {
		return
	}

	// determine winner index according to the hash and remainder
	winnerBet := keeper.DeceideWinnerByBetsHash(allActiveBets, betCount)
	// decide the winner bet according to the proposer cons addres
	// winnerBet := keeper.DeceideWinnerByProposerHash(allActiveBets, betCount, ctx.BlockHeader().ProposerAddress)

	// calculate payout and then transfer from pool
	payout := k.CalculateAndTransferPayout(ctx, winnerBet, allActiveBets)

	// mark lottery as finished because the winer is determined
	k.FinishLottery(ctx, currentLottery, betCount, payout, winnerBet.Id)

	k.BetKeeper.SettleAllActiveBets(ctx, currentLottery.Id)

	poolBalance := k.GetPoolBalance(ctx)
	fmt.Printf("Winner Bet => creator: %s, Amount: %s Payout: %s\nPool balance: %s\n", winnerBet.Creator, winnerBet.Amount, payout, poolBalance)

}
