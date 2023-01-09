package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

// params keys
var (
	KeyLotteryFee = []byte("LotteryFee")
	KeyBetSize    = []byte("BetSize")

	DefaultLotteryFee = uint64(5)
	DefaultBetSize    = BetSize{
		MinBet: uint64(1),
		MaxBet: uint64(100),
	}
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams() Params {
	return Params{
		LotteryFee: DefaultLotteryFee,
		BetSize:    DefaultBetSize,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams()
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyLotteryFee, &p.LotteryFee, validateLotteryFee),
		paramtypes.NewParamSetPair(KeyBetSize, &p.BetSize, validateBetSize),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateLotteryFee(p.LotteryFee); err != nil {
		return err
	}
	if err := validateBetSize(p.BetSize); err != nil {
		return err
	}
	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

func validateLotteryFee(i interface{}) error {
	_, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateBetSize(i interface{}) error {
	p, ok := i.(BetSize)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if p.MinBet >= p.MaxBet {
		return fmt.Errorf("minimum value should be less than maximum value")
	}

	return nil
}
