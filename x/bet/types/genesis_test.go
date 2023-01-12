package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vjdmhd/lottery/x/bet/types"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				ActiveBetList: []types.Bet{
					{
						Id:      0,
						Creator: "creator1",
					},
					{
						Id:      1,
						Creator: "creator2",
					},
				},
				SettledBetList: []types.Bet{
					{
						Id:      0,
						Creator: "creator1",
					},
					{
						Id:      1,
						Creator: "creator2",
					},
				},
			},
			valid: true,
		},
		{
			desc: "duplicated bet",
			genState: &types.GenesisState{
				ActiveBetList: []types.Bet{
					{
						Id:      0,
						Creator: "creator1",
					},
					{
						Id:      1,
						Creator: "creator1",
					},
				},
				SettledBetList: []types.Bet{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
			},
			valid: false,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
