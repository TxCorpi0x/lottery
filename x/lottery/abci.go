package lottery

import (
	"bytes"
	"encoding/gob"

	"github.com/cespare/xxhash/v2"
	sdk "github.com/cosmos/cosmos-sdk/types"
	betmoduletypes "github.com/vjdmhd/lottery/x/bet/types"
	"github.com/vjdmhd/lottery/x/lottery/keeper"
)

func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	currentLottery := k.GetOrCreateCurrentLottery(ctx)

	allActiveBets := k.BetKeeper.GetAllActiveBet(ctx)
	betCount := uint64(len(allActiveBets))

	// get consensus address of the current block proposer
	proposerConsAddr := sdk.ConsAddress(ctx.BlockHeader().ProposerAddress)

	// check all active bets for any existing bet from proposer operator account
	for _, v := range allActiveBets {

		// extract consensus address from bet creator bech32
		creatorConsAddress, err := sdk.ConsAddressFromBech32(v.Creator)
		if err != nil {
			// if this happens, it means that ther is big fatal problem in the chain
			// so we should halt it
			panic(err)
		}

		if creatorConsAddress.Equals(proposerConsAddr) {
			// block proposer has bet in active bets,
			// so will choose the winner in next block
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
}

// get hash of bets slice using the most efficient hash algoritm
// http://cyan4973.github.io/xxHash/
func calculateHash(bets []betmoduletypes.Bet) uint64 {
	var b bytes.Buffer
	gob.NewEncoder(&b).Encode(bets)

	return xxhash.Sum64(b.Bytes())
}
