package types

import sdkmath "cosmossdk.io/math"

func NewBet(count uint64, creator string, amount sdkmath.Int, blockHeigh int64) Bet {
	return Bet{
		Id:      uint64(count),
		Creator: creator,
		Amount:  amount,
		Height:  blockHeigh,
	}
}
