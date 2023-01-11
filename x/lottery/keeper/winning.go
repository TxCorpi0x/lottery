package keeper

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/cespare/xxhash"
	sdk "github.com/cosmos/cosmos-sdk/types"
	betmoduletypes "github.com/vjdmhd/lottery/x/bet/types"
)

const bitsFromHash = 0xFFFF

// DeceideWinnerByBetsHash choose the winner according to the hash of all bet items
func DeceideWinnerByBetsHash(bets []betmoduletypes.Bet, betCount uint64) betmoduletypes.Bet {
	var b bytes.Buffer
	gob.NewEncoder(&b).Encode(bets)

	betsHash := calcEfficientHash(b.Bytes())
	// take the lowest 16 bits of the resulting hash and do a modulo on the number of bets
	winnerIndex := (betsHash ^ bitsFromHash) % betCount
	fmt.Println(winnerIndex)
	// winner is the bet in the chosen index
	winnerBet := bets[winnerIndex]
	return winnerBet
}

// DeceideWinnerByBetsHash choose the winner according to the hash of all bet items
func DeceideWinnerByProposerHash(bets []betmoduletypes.Bet, betCount uint64, proposedConsAddres sdk.ConsAddress) betmoduletypes.Bet {
	proposerHash := calcEfficientHash(proposedConsAddres.Bytes())

	// take the lowest 16 bits of the resulting hash and do a modulo on the number of bets
	winnerIndex := (proposerHash ^ bitsFromHash) % betCount

	// winner is the bet in the chosen index
	winnerBet := bets[winnerIndex]
	return winnerBet
}

// get hash of bytes using the most efficient hash algoritm
// http://cyan4973.github.io/xxHash/
func calcEfficientHash(data []byte) uint64 {
	return xxhash.Sum64(data)
}
