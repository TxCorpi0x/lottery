package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		LotteryList: []Lottery{},
		Params:      DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in lottery
	lotteryIndexMap := make(map[string]struct{})

	for _, elem := range gs.LotteryList {
		index := string(LotteryKey(elem.Id))
		if _, ok := lotteryIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for lottery")
		}
		lotteryIndexMap[index] = struct{}{}
	}

	return gs.Params.Validate()
}
