package keeper_test

import (
	"strconv"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/vjdmhd/lottery/app/params"
	keepertest "github.com/vjdmhd/lottery/testutil/keeper"
	"github.com/vjdmhd/lottery/testutil/nullify"
	"github.com/vjdmhd/lottery/x/lottery/keeper"
	"github.com/vjdmhd/lottery/x/lottery/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNLottery(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Lottery {
	items := make([]types.Lottery, n)
	for i := range items {
		items[i].Id = uint64(i)

		keeper.SetLottery(ctx, items[i])
	}
	return items
}

func TestLotteryGet(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	items := createNLottery(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetLottery(ctx,
			item.Id,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestLotteryRemove(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	items := createNLottery(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveLottery(ctx,
			item.Id,
		)
		_, found := keeper.GetLottery(ctx,
			item.Id,
		)
		require.False(t, found)
	}
}

func TestLotteryGetAll(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	items := createNLottery(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllLottery(ctx)),
	)
}

func TestGetOrCreateCurrentLottery(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)

	// check if the lottery has not finished in the current block
	// it should not create a new lottery
	currentLottery := keeper.GetOrCreateCurrentLottery(ctx)
	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(5 * time.Second))
	currentLottery2 := keeper.GetOrCreateCurrentLottery(ctx)
	require.Equal(t, currentLottery.Id, currentLottery2.Id)

	// set lottery as finished, then the new lottery should be generated
	keeper.FinishLottery(ctx, currentLottery, 0, sdk.NewCoin(params.DefaultBondDenom, sdk.NewInt(0)), 100)
	currentLottery3 := keeper.GetOrCreateCurrentLottery(ctx)
	require.NotEqual(t, currentLottery3.Id, currentLottery2.Id)

	lotteries := keeper.GetAllLottery(ctx)
	require.Equal(t, 2, len(lotteries))
}

func TestCurrentCurrentLottery(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	ctx = ctx.WithBlockTime(time.Now())

	// simulate 10 blocks of create new lottery and finish
	for i := 1; i <= 10; i++ {
		ctx = ctx.WithBlockTime(ctx.BlockTime().Add(5 * time.Second))
		currentLottery := keeper.GetOrCreateCurrentLottery(ctx)

		keeper.FinishLottery(ctx, currentLottery, 0, sdk.NewCoin(params.DefaultBondDenom, sdk.NewInt(0)), 100)
	}

	// check if the current lottery is returned currectly or not
	latestLottery := keeper.GetCurrentLottery(ctx)
	require.Equal(t, uint64(ctx.BlockTime().UnixNano()), latestLottery.Id)
}
