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
	betmodulekeepr "github.com/vjdmhd/lottery/x/bet/keeper"
	betmoduletypes "github.com/vjdmhd/lottery/x/bet/types"
	"github.com/vjdmhd/lottery/x/lottery/keeper"
	"github.com/vjdmhd/lottery/x/lottery/types"
)

func LotteryKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	paramsSubspace := typesparams.NewSubspace(cdc,
		types.Amino,
		storeKey,
		memStoreKey,
		"LotteryParams",
	)

	maccPerms := map[string][]string{
		authtypes.FeeCollectorName: nil,
		minttypes.ModuleName:       {authtypes.Minter},
		types.ModuleName:           nil,
		// this line is used by starport scaffolding # stargate/app/maccPerms
	}

	accStoreKey := sdk.NewKVStoreKey(authtypes.StoreKey)
	accMemStoreKey := storetypes.NewMemoryStoreKey("mem_" + authtypes.StoreKey)
	accParamsSubspace := typesparams.NewSubspace(cdc,
		types.Amino,
		accStoreKey,
		accMemStoreKey,
		"AccParams",
	)
	accountKpr := authkeeper.NewAccountKeeper(
		cdc,
		accStoreKey,
		accParamsSubspace,
		authtypes.ProtoBaseAccount,
		maccPerms,
		sdk.Bech32PrefixAccAddr,
	)

	bankStoreKey := sdk.NewKVStoreKey(banktypes.StoreKey)
	bankMemStoreKey := storetypes.NewMemoryStoreKey("mem_" + banktypes.StoreKey)
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

	stakingStoreKey := sdk.NewKVStoreKey(stakingtypes.StoreKey)
	stakingMemStoreKey := storetypes.NewMemoryStoreKey("mem_" + stakingtypes.StoreKey)
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

	betStoreKey := sdk.NewKVStoreKey(betmoduletypes.StoreKey)
	betMemStoreKey := storetypes.NewMemoryStoreKey(betmoduletypes.MemStoreKey)
	betParamsSubspace := typesparams.NewSubspace(cdc,
		types.Amino,
		betStoreKey,
		betMemStoreKey,
		"BetParams",
	)
	betKpr := betmodulekeepr.NewKeeper(
		cdc,
		betStoreKey,
		betMemStoreKey,
		betParamsSubspace,
		bankKpr,
	)

	k := keeper.NewKeeper(
		cdc,
		storeKey,
		memStoreKey,
		paramsSubspace,
		accountKpr,
		bankKpr,
		betKpr,
		stakingkKpr,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())

	return k, ctx
}
