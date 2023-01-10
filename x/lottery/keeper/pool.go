package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/vjdmhd/lottery/app/params"
	"github.com/vjdmhd/lottery/x/lottery/types"
)

func (k Keeper) GetPoolBalance(ctx sdk.Context) sdk.Coin {
	moduleAcc := k.AccountKeeper.GetModuleAccount(ctx, types.ModuleName)
	return k.BankKeeper.GetBalance(ctx, moduleAcc.GetAddress(), params.DefaultBondDenom)
}

func (k Keeper) TransferFeesAndAmount(ctx sdk.Context, senderAcc sdk.AccAddress, amount sdk.Coin) {
	totalCoins := sdk.NewCoins(amount)
	k.BankKeeper.SendCoinsFromAccountToModule(ctx, senderAcc, types.ModuleName, totalCoins)
}
