package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/vjdmhd/lottery/testutil/keeper"
	"github.com/vjdmhd/lottery/x/bet/keeper"
	"github.com/vjdmhd/lottery/x/bet/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestBetMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.BetKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateBet{
			Creator: creator,
			Amount:  uint64(i),
		}
		_, err := srv.CreateBet(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetActiveBet(ctx,
			expected.Creator,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}
