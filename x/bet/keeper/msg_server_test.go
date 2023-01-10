package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/vjdmhd/lottery/testutil/keeper"
	lotsim "github.com/vjdmhd/lottery/testutil/simapp"
	"github.com/vjdmhd/lottery/x/bet/keeper"
	"github.com/vjdmhd/lottery/x/bet/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, _, ctx := keepertest.BetKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

func setupMsgServerAndApp(t testing.TB) (*lotsim.TestApp, *keeper.KeeperTest, types.MsgServer, sdk.Context, context.Context) {
	tApp, ctx, err := lotsim.GetTestObjects()
	if err != nil {
		panic(err)
	}
	return tApp, &tApp.BetKeeper, keeper.NewMsgServerImpl(tApp.BetKeeper), ctx, sdk.WrapSDKContext(ctx)
}
