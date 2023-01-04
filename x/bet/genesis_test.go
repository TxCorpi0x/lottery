package bet_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/vjdmhd/lottery/testutil/keeper"
	"github.com/vjdmhd/lottery/testutil/nullify"
	"github.com/vjdmhd/lottery/x/bet"
	"github.com/vjdmhd/lottery/x/bet/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		BetList: []types.Bet{
			{
				Id: "0",
			},
			{
				Id: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BetKeeper(t)
	bet.InitGenesis(ctx, *k, genesisState)
	got := bet.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.BetList, got.BetList)
	// this line is used by starport scaffolding # genesis/test/assert
}
