package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/vjdmhd/lottery/x/bet/types"
)

// SetBetCount sets total bets statistics
func (k Keeper) SetBetStats(ctx sdk.Context, betstats types.BetStats) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BetStatusKey))
	b := k.cdc.MustMarshal(&betstats)
	store.Set(types.KeyPrefix("0"), b)
}

// GetBetStats gets total bets statistics
func (k Keeper) GetBetStats(ctx sdk.Context) (val types.BetStats) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BetStatusKey))

	b := store.Get(types.KeyPrefix("0"))
	if b == nil {
		return val
	}

	k.cdc.MustUnmarshal(b, &val)
	return val
}

// SetActiveBet set a specific bet in the store for its creator
func (k Keeper) SetActiveBet(ctx sdk.Context, bet types.Bet) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActiveBetKeyPrefix))
	b := k.cdc.MustMarshal(&bet)
	store.Set(types.ActiveBetKey(bet.Creator), b)
}

// GetActiveBet returns an active bet from its creator
func (k Keeper) GetActiveBet(
	ctx sdk.Context,
	creator string,

) (val types.Bet, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActiveBetKeyPrefix))

	b := store.Get(types.ActiveBetKey(creator))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetAllActiveBet returns all active bet
func (k Keeper) GetAllActiveBet(ctx sdk.Context) (list []types.Bet) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActiveBetKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Bet
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// RemoveAllActiveBet removes all active bet items from the store
func (k Keeper) RemoveAllActiveBet(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActiveBetKeyPrefix))
	allBets := k.GetAllActiveBet(ctx)
	for _, v := range allBets {
		store.Delete(types.ActiveBetKey(v.Creator))
	}
}

// SetSettledBet set a specific settled bet in the store for its creator
func (k Keeper) SetSettledBet(ctx sdk.Context, bet types.Bet) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SettledBetKeyPrefix))
	b := k.cdc.MustMarshal(&bet)
	store.Set(types.SettledBetKey(bet.LotteryId, bet.Id), b)
}

// GetSettledBet returns an active bet from its creator
func (k Keeper) GetSettledBet(
	ctx sdk.Context,
	lotteryID string,
	betID string,
) (val types.Bet, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SettledBetKeyPrefix))

	b := store.Get(types.SettledBetKey(lotteryID, betID))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetAllSettledBet returns all settled bet
func (k Keeper) GetAllSettledBet(ctx sdk.Context) (list []types.Bet) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SettledBetKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Bet
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
