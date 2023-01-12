package types

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
}

// LotteryKeeper defines the expected interface needed to retrieve lottery info.
type LotteryKeeper interface {
	GetMinMaxBetAllowed(ctx sdk.Context) (sdkmath.Int, sdkmath.Int)
	TransferFeesAndAmount(ctx sdk.Context, senderAcc sdk.AccAddress, amount sdk.Coin)
	GetLotteryFee(ctx sdk.Context) sdkmath.Int
}
