package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/vjdmhd/lottery/testutil/keeper"
	"github.com/vjdmhd/lottery/x/bet/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.BetKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
