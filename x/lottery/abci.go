package lottery

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/cespare/xxhash/v2"
	sdk "github.com/cosmos/cosmos-sdk/types"
	betmoduletypes "github.com/vjdmhd/lottery/x/bet/types"
	"github.com/vjdmhd/lottery/x/lottery/keeper"
)

func EndBlocker(ctx sdk.Context, k keeper.Keeper) {

	currentLottery := k.GetOrCreateCurrentLottery(ctx)

	allActiveBets := k.BetKeeper.GetAllActiveBet(ctx)
	betCount := uint64(len(allActiveBets))

	fmt.Printf("Block height: %d, Lottery ID: %d, Bet Count: %d \n", ctx.BlockHeight(), currentLottery.Id, betCount)

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
			fmt.Printf("Block Proposer operator has bet in the active bet list, continue to the next block.\n")
			return
		}

	}

	// if the bets count does not satisfy the min count
	// should return and continue in nex end blocker
	if betCount < 10 {
		return
	}

	// determine winner index according to the hash and remainder
	betsHash := calculateHash(allActiveBets)
	winnerIndex := (betsHash ^ 0xFFFF) % betCount
	winnerBet := allActiveBets[winnerIndex]

	// calculate payout and then transfer from pool
	payout := k.CalculateAndTransferPayout(ctx, winnerBet, allActiveBets)

	// mark lottery as finished because the winer is determined
	k.FinishLottery(ctx, currentLottery, betCount, payout, winnerBet.Id)

	k.BetKeeper.SettleAllActiveBets(ctx, currentLottery.Id)

	poolBalance := k.GetPoolBalance(ctx)
	fmt.Printf("Winner Bet => creator: %s, Amount: %s Payout: %s pool balance: %s\n", winnerBet.Creator, winnerBet.Amount, payout, poolBalance)

}

// get hash of bets slice using the most efficient hash algoritm
// http://cyan4973.github.io/xxHash/
func calculateHash(bets []betmoduletypes.Bet) uint64 {
	var b bytes.Buffer
	gob.NewEncoder(&b).Encode(bets)

	return xxhash.Sum64(b.Bytes())
}
