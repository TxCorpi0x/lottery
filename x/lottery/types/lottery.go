package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewLottery creates new lottery with unique and incfremental id
// the best option is the timestamp of blocktime
// the winner id is set as null because it is not finished yet
func NewLottery(ctx sdk.Context) Lottery {
	return Lottery{
		Id:         uint64(ctx.BlockTime().UnixNano()),
		StartBlock: ctx.BlockHeight(),
		WinnerId:   UnknownWinnerID,
	}
}
