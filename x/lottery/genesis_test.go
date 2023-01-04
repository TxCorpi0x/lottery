package lottery_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/vjdmhd/lottery/testutil/keeper"
	"github.com/vjdmhd/lottery/testutil/nullify"
	"github.com/vjdmhd/lottery/x/lottery"
	"github.com/vjdmhd/lottery/x/lottery/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		LotteryList: []types.Lottery{
			{
				ID: "0",
			},
			{
				ID: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LotteryKeeper(t)
	lottery.InitGenesis(ctx, *k, genesisState)
	got := lottery.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.LotteryList, got.LotteryList)
	// this line is used by starport scaffolding # genesis/test/assert
}
