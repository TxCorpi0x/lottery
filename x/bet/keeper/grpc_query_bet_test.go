package keeper_test

import (
	"strconv"
	"testing"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/vjdmhd/lottery/testutil/keeper"
	"github.com/vjdmhd/lottery/testutil/nullify"
	"github.com/vjdmhd/lottery/x/bet/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestActiveBetQuerySingle(t *testing.T) {
	keeper, _, ctx := keepertest.BetKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNActiveBet(keeper, ctx, 2)
	msgs[0].Amount = math.NewInt(1)
	msgs[1].Amount = math.NewInt(1)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetBetRequest
		response *types.QueryGetBetResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetBetRequest{
				Creator: msgs[0].Creator,
			},
			response: &types.QueryGetBetResponse{Bet: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetBetRequest{
				Creator: msgs[1].Creator,
			},
			response: &types.QueryGetBetResponse{Bet: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetBetRequest{
				Creator: "notvalid",
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.ActiveBet(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestActiveBetQueryPaginated(t *testing.T) {
	keeper, _, ctx := keepertest.BetKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNActiveBet(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllActiveBetRequest {
		return &types.QueryAllActiveBetRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ActiveBetAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Bet), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Bet),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ActiveBetAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Bet), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Bet),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ActiveBetAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, 1, int(resp.Pagination.Total))
		msgs[0].Amount = math.NewInt(4)
		require.ElementsMatch(t,
			nullify.Fill([]types.Bet{msgs[0]}),
			nullify.Fill(resp.Bet),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ActiveBetAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
