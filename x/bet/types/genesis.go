package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ActiveBetList:  []Bet{},
		SettledBetList: []Bet{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in bet
	activeBetIndexMap := make(map[string]struct{})

	for _, elem := range gs.ActiveBetList {
		index := string(ActiveBetKey(elem.Creator))
		if _, ok := activeBetIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for bet")
		}
		activeBetIndexMap[index] = struct{}{}
	}

	settledBetIndexMap := make(map[string]struct{})

	for _, elem := range gs.SettledBetList {
		index := string(SettledBetKey(elem.LotteryId, elem.Id))
		if _, ok := settledBetIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for bet")
		}
		settledBetIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
