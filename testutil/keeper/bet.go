package keeper

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"
	"github.com/vjdmhd/lottery/x/bet/keeper"
	"github.com/vjdmhd/lottery/x/bet/types"
	lotterymodulekeeper "github.com/vjdmhd/lottery/x/lottery/keeper"
	lotterymoduletypes "github.com/vjdmhd/lottery/x/lottery/types"
)

func BetKeeper(t testing.TB) (*keeper.Keeper, bankkeeper.Keeper, sdk.Context) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	accStoreKey := sdk.NewKVStoreKey(authtypes.StoreKey)
	accMemStoreKey := storetypes.NewMemoryStoreKey("mem_" + authtypes.StoreKey)

	bankStoreKey := sdk.NewKVStoreKey(banktypes.StoreKey)
	bankMemStoreKey := storetypes.NewMemoryStoreKey("mem_" + banktypes.StoreKey)

	stakingStoreKey := sdk.NewKVStoreKey(stakingtypes.StoreKey)
	stakingMemStoreKey := storetypes.NewMemoryStoreKey("mem_" + stakingtypes.StoreKey)

	lotteryStoreKey := sdk.NewKVStoreKey(lotterymoduletypes.StoreKey)
	lotteryMemStoreKey := storetypes.NewMemoryStoreKey(lotterymoduletypes.MemStoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)
	stateStore.MountStoreWithDB(accStoreKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(accMemStoreKey, storetypes.StoreTypeMemory, nil)
	stateStore.MountStoreWithDB(bankStoreKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(bankMemStoreKey, storetypes.StoreTypeMemory, nil)
	stateStore.MountStoreWithDB(stakingStoreKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(stakingMemStoreKey, storetypes.StoreTypeMemory, nil)
	stateStore.MountStoreWithDB(lotteryStoreKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(lotteryMemStoreKey, storetypes.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	accParamsSubspace := typesparams.NewSubspace(cdc,
		types.Amino,
		accStoreKey,
		accMemStoreKey,
		"AccParams",
	)

	maccPerms := map[string][]string{
		authtypes.FeeCollectorName:     nil,
		minttypes.ModuleName:           {authtypes.Minter},
		stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
		stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
		types.ModuleName:               nil,
	}

	accountKpr := authkeeper.NewAccountKeeper(
		cdc,
		accStoreKey,
		accParamsSubspace,
		authtypes.ProtoBaseAccount,
		maccPerms,
		sdk.Bech32PrefixAccAddr,
	)

	bankParamsSubspace := typesparams.NewSubspace(cdc,
		types.Amino,
		bankStoreKey,
		bankMemStoreKey,
		"BankParams",
	)
	bankKpr := bankkeeper.NewBaseKeeper(
		cdc,
		bankStoreKey,
		accountKpr,
		bankParamsSubspace,
		map[string]bool{},
	)

	stakingParamsSubspace := typesparams.NewSubspace(cdc,
		types.Amino,
		stakingStoreKey,
		stakingMemStoreKey,
		"StakingParams",
	)
	stakingkKpr := stakingkeeper.NewKeeper(
		cdc,
		stakingStoreKey,
		accountKpr,
		bankKpr,
		stakingParamsSubspace,
	)

	paramsSubspace := typesparams.NewSubspace(cdc,
		types.Amino,
		storeKey,
		memStoreKey,
		"BetParams",
	)
	k := keeper.NewKeeper(
		cdc,
		storeKey,
		memStoreKey,
		paramsSubspace,
		bankKpr,
	)

	lotteryParamsSubspace := typesparams.NewSubspace(cdc,
		types.Amino,
		lotteryStoreKey,
		lotteryMemStoreKey,
		"LotteryParams",
	)
	lotteryKpr := lotterymodulekeeper.NewKeeper(
		cdc,
		lotteryStoreKey,
		lotteryMemStoreKey,
		lotteryParamsSubspace,
		accountKpr,
		bankKpr,
		k,
		stakingkKpr,
	)
	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	lotteryKpr.SetParams(ctx, lotterymoduletypes.DefaultParams())
	k.SetLotteryKeeper(lotteryKpr)

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())

	return k, bankKpr, ctx
}
