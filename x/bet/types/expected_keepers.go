package types

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
}

// LotteryKeeper defines the expected interface needed to retrieve lottery info.
type LotteryKeeper interface {
	GetMinMaxBetAllowed(ctx sdk.Context) (sdkmath.Int, sdkmath.Int)
	TransferFeesAndAmount(ctx sdk.Context, senderAcc sdk.AccAddress, amount sdk.Coin)
	GetLotteryFee(ctx sdk.Context) sdkmath.Int

	// Methods imported from bank should be defined here
}
