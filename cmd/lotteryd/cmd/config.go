package cmd

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/vjdmhd/lottery/app/params"
)

const (
	// AccountAddressPrefix prefix used for generating account address
	AccountAddressPrefix = "lot"
)

func initSDKConfig() {
	// Set prefixes
	accountPubKeyPrefix := AccountAddressPrefix + "pub"
	validatorAddressPrefix := AccountAddressPrefix + "valoper"
	validatorPubKeyPrefix := AccountAddressPrefix + "valoperpub"
	consNodeAddressPrefix := AccountAddressPrefix + "valcons"
	consNodePubKeyPrefix := AccountAddressPrefix + "valconspub"

	// Set and seal config
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(AccountAddressPrefix, accountPubKeyPrefix)
	config.SetBech32PrefixForValidator(validatorAddressPrefix, validatorPubKeyPrefix)
	config.SetBech32PrefixForConsensusNode(consNodeAddressPrefix, consNodePubKeyPrefix)

	err := sdk.RegisterDenom(params.HumanCoinUnit, sdk.OneDec())
	if err != nil {
		panic(err)
	}
	err = sdk.RegisterDenom(params.BaseCoinUnit, sdk.NewDecWithPrec(1, params.LOTExponent))
	if err != nil {
		panic(err)
	}

	config.Seal()
}
