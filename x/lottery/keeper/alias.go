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
