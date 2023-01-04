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

func createNBet(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Bet {
	items := make([]types.Bet, n)
	for i := range items {
		items[i].Id = strconv.Itoa(i)

		keeper.SetBet(ctx, items[i])
	}
	return items
}

func TestBetGet(t *testing.T) {
	keeper, ctx := keepertest.BetKeeper(t)
	items := createNBet(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetBet(ctx,
			"TODO", // TODO: get current lottery id
			item.Id,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestBetGetAll(t *testing.T) {
	keeper, ctx := keepertest.BetKeeper(t)
	items := createNBet(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllBet(ctx)),
	)
}
