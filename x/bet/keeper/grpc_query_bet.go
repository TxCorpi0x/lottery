package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/vjdmhd/lottery/x/bet/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ActiveBetAll(c context.Context, req *types.QueryAllActiveBetRequest) (*types.QueryAllBetResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var bets []types.Bet
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	betStore := prefix.NewStore(store, types.ActiveBetKeyPrefix)

	pageRes, err := query.Paginate(betStore, req.Pagination, func(key []byte, value []byte) error {
		var bet types.Bet
		if err := k.cdc.Unmarshal(value, &bet); err != nil {
			return err
		}

		bets = append(bets, bet)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllBetResponse{Bet: bets, Pagination: pageRes}, nil
}

func (k Keeper) ActiveBet(c context.Context, req *types.QueryGetBetRequest) (*types.QueryGetBetResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetActiveBet(
		ctx,
		req.Creator,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetBetResponse{Bet: val}, nil
}

func (k Keeper) SettledBetAll(c context.Context, req *types.QueryAllSettledBetRequest) (*types.QueryAllBetResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var bets []types.Bet
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	betStore := prefix.NewStore(store, types.SettledBetOfLotteryPrefix(req.LotteryId))

	pageRes, err := query.Paginate(betStore, req.Pagination, func(key []byte, value []byte) error {
		var bet types.Bet
		if err := k.cdc.Unmarshal(value, &bet); err != nil {
			return err
		}

		bets = append(bets, bet)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllBetResponse{Bet: bets, Pagination: pageRes}, nil
}
