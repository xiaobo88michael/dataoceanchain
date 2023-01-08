package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/xiaobo88michael/dataoceanchain/x/dataocean/types"
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

				VideoList: []types.Video{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				VideoCount: 2,
				VideoLinkList: []types.VideoLink{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated video",
			genState: &types.GenesisState{
				VideoList: []types.Video{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid video count",
			genState: &types.GenesisState{
				VideoList: []types.Video{
					{
						Id: 1,
					},
				},
				VideoCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated videoLink",
			genState: &types.GenesisState{
				VideoLinkList: []types.VideoLink{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
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
