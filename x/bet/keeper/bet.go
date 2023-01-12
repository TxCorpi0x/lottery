package keeper

import (
	"sort"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/vjdmhd/lottery/x/bet/types"
)

// SetBetCount sets total bets statistics
func (k Keeper) setBetStats(ctx sdk.Context, betstats types.BetStats) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.BetStatsKey)
	b := k.cdc.MustMarshal(&betstats)
	store.Set(types.KeyPrefix("0"), b)
}

// getBetStats gets total bets statistics
func (k Keeper) getBetStats(ctx sdk.Context) (val types.BetStats) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.BetStatsKey)

	b := store.Get(types.KeyPrefix("0"))
	if b == nil {
		return val
	}

	k.cdc.MustUnmarshal(b, &val)
	return val
}

// SetActiveBet set a specific bet in the store for its creator
func (k Keeper) SetActiveBet(ctx sdk.Context, bet types.Bet) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.ActiveBetKeyPrefix)
	b := k.cdc.MustMarshal(&bet)
	store.Set(types.ActiveBetKey(bet.Creator), b)
}

// GetActiveBet returns an active bet from its creator
func (k Keeper) GetActiveBet(
	ctx sdk.Context,
	creator string,

) (val types.Bet, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.ActiveBetKeyPrefix)

	b := store.Get(types.ActiveBetKey(creator))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetAllActiveBet returns all active bet
func (k Keeper) GetAllActiveBet(ctx sdk.Context) (list []types.Bet) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.ActiveBetKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Bet
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	// ascending sort of bets according to the time
	sort.Slice(list, func(i, j int) bool {
		return list[i].Id < list[j].Id
	})

	return
}

// SetSettledBet set a specific settled bet in the store for its creator
func (k Keeper) SetSettledBet(ctx sdk.Context, bet types.Bet) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.SettledBetKeyPrefix)
	b := k.cdc.MustMarshal(&bet)
	store.Set(types.SettledBetKey(bet.LotteryId, bet.Id), b)
}

// GetAllSettledBet returns all settled bet
func (k Keeper) GetAllSettledBet(ctx sdk.Context) (list []types.Bet) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.SettledBetKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Bet
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) SettleAllActiveBets(ctx sdk.Context, lotteryId uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.ActiveBetKeyPrefix)
	allBets := k.GetAllActiveBet(ctx)
	for _, v := range allBets {
		v.LotteryId = lotteryId
		k.SetSettledBet(ctx, v)
		store.Delete(types.ActiveBetKey(v.Creator))
	}

}
