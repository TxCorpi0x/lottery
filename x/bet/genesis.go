package bet

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/vjdmhd/lottery/x/bet/keeper"
	"github.com/vjdmhd/lottery/x/bet/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the active bet
	for _, elem := range genState.ActiveBetList {
		k.SetActiveBet(ctx, elem)
	}
	// Set all the settled bet
	for _, elem := range genState.SettledBetList {
		k.SetSettledBet(ctx, elem)
	}

	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.ActiveBetList = k.GetAllActiveBet(ctx)
	genesis.SettledBetList = k.GetAllSettledBet(ctx)

	return genesis
}
