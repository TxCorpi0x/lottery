package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/vjdmhd/lottery/x/lottery/types"
)

// SetLottery set a specific lottery in the store from its index
func (k Keeper) SetLottery(ctx sdk.Context, lottery types.Lottery) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.LotteryKeyPrefix)
	b := k.cdc.MustMarshal(&lottery)
	store.Set(types.LotteryKey(
		lottery.Id,
	), b)
}

// GetLottery returns a lottery from its index
func (k Keeper) GetLottery(
	ctx sdk.Context,
	id uint64,

) (val types.Lottery, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.LotteryKeyPrefix)

	b := store.Get(types.LotteryKey(
		id,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveLottery removes a lottery from the store
func (k Keeper) RemoveLottery(
	ctx sdk.Context,
	id uint64,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.LotteryKeyPrefix)
	store.Delete(types.LotteryKey(
		id,
	))
}

// GetAllLottery returns all lottery
func (k Keeper) GetAllLottery(ctx sdk.Context) (list []types.Lottery) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.LotteryKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Lottery
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetCurrentLottery returns current lottery
// reversely sorted lottery kvstore returns the latest lottery
// because we have used number as key, it is always sorted ascending
func (k Keeper) GetCurrentLottery(ctx sdk.Context) (current types.Lottery) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.LotteryKeyPrefix)
	// to get the latest lottery item KVStoreReversePrefixIterator
	// is being chosen to get latest item
	iterator := sdk.KVStoreReversePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Lottery
		k.cdc.MustUnmarshal(iterator.Value(), &val)

		// set the return value
		current = val

		// do not continue the loop because we set the current lottery value
		break
	}

	// if it is the first block, this returns no lottery
	// in the begin blocker the GetOrCreateCurrentLottery
	// will create the first one
	return
}

// GetOrCreateCurrentLottery returns current unfinished lottery if available
// creates and return if the unfinished lottery is not available
func (k Keeper) GetOrCreateCurrentLottery(ctx sdk.Context) (current types.Lottery) {
	currentLottery := k.GetCurrentLottery(ctx)

	// if the winner is determined ths means that the lottery has not finished
	if currentLottery.WinnerId == types.UnknownWinnerID {
		return currentLottery
	}

	newLottery := types.NewLottery(ctx)
	// sets new lottery
	k.SetLottery(ctx, newLottery)

	// return the latest unfinished lottery
	return k.GetCurrentLottery(ctx)
}

// FinishLottery set the finished lottery attributes
func (k Keeper) FinishLottery(
	ctx sdk.Context,
	lottery types.Lottery,
	betCount uint64,
	payout sdk.Coin,
	winnerID uint64,
) {
	lottery.BetCount = betCount
	lottery.EndBlock = ctx.BlockHeight()
	lottery.Payout = payout
	lottery.WinnerId = winnerID
	k.SetLottery(ctx, lottery)
}
