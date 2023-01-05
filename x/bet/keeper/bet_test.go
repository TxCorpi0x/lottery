package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/vjdmhd/lottery/testutil/keeper"
	"github.com/vjdmhd/lottery/testutil/nullify"
	"github.com/vjdmhd/lottery/x/bet/keeper"
	"github.com/vjdmhd/lottery/x/bet/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNActiveBet(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Bet {
	items := make([]types.Bet, n)
	for i := range items {
		items[i].Amount = uint64(i)

		keeper.SetActiveBet(ctx, items[i])
	}

	return items
}

func TestActiveBetGet(t *testing.T) {
	keeper, ctx := keepertest.BetKeeper(t)
	items := createNActiveBet(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetActiveBet(ctx,
			item.Creator,
		)
		// this means after modifying the active bet
		// only s single item should be updated
		// and value should be largest number in the loop
		item.Amount = 9
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestActiveBetGetAll(t *testing.T) {
	keeper, ctx := keepertest.BetKeeper(t)
	items := createNActiveBet(keeper, ctx, 10)
	// this means after modifying the active bet
	// only s single item should be updated
	// and value should be largest number in the loop
	items[0].Amount = 9
	require.ElementsMatch(t,
		nullify.Fill([]types.Bet{items[0]}),
		nullify.Fill(keeper.GetAllActiveBet(ctx)),
	)
}
