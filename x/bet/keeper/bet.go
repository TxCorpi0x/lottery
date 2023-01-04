package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/vjdmhd/lottery/x/bet/types"
)

// SetBet set a specific bet in the store from its index
func (k Keeper) SetBet(ctx sdk.Context, bet types.Bet) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BetKeyPrefix))
	b := k.cdc.MustMarshal(&bet)
	store.Set(types.BetKey(
		bet.LotteryId,
		bet.Creator,
	), b)
}

// GetBet returns a bet from its index
func (k Keeper) GetBet(
	ctx sdk.Context,
	lotteryID,
	creator string,

) (val types.Bet, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BetKeyPrefix))

	b := store.Get(types.BetKey(lotteryID, creator))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetAllBet returns all bet
func (k Keeper) GetAllBet(ctx sdk.Context) (list []types.Bet) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BetKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Bet
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
