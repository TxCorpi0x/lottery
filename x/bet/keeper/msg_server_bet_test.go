package keeper_test

import (
	"strconv"
	"testing"

	"cosmossdk.io/math"
	"github.com/stretchr/testify/require"

	lotsim "github.com/vjdmhd/lottery/testutil/simapp"
	"github.com/vjdmhd/lottery/x/bet/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestBetMsgServerCreate(t *testing.T) {
	_, k, srv, ctx, wctx := setupMsgServerAndApp(t)

	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateBet{
			Creator: lotsim.TestParamUsers["client1"].Address.String(),
			Amount:  math.NewInt(int64(50)),
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
