package testutil

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	mintmoduletypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/vjdmhd/lottery/app/params"
)

// SetCoins sets the balance of accounts for testing
func SetCoins(ctx *sdk.Context, k bankkeeper.Keeper, addr sdk.AccAddress, amount int64) error {
	coin := sdk.NewCoins(sdk.NewCoin(params.DefaultBondDenom, sdk.NewInt(amount)))
	err := k.MintCoins(*ctx, mintmoduletypes.ModuleName, coin)
	if err != nil {
		return err
	}
	err = k.SendCoinsFromModuleToAccount(*ctx, mintmoduletypes.ModuleName, addr, coin)
	if err != nil {
		return err
	}
	return nil
}

// SetModuleCoins sets the balance of accounts for testing
func SetModuleCoins(ctx *sdk.Context, k bankkeeper.Keeper, moduleName string, amount int64) error {
	coin := sdk.NewCoins(sdk.NewCoin(params.DefaultBondDenom, sdk.NewInt(amount)))
	err := k.MintCoins(*ctx, mintmoduletypes.ModuleName, coin)
	if err != nil {
		return err
	}
	err = k.SendCoinsFromModuleToModule(*ctx, mintmoduletypes.ModuleName, moduleName, coin)
	if err != nil {
		return err
	}
	return nil
}
