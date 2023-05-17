package chainfour_test

import (
	"testing"

	keepertest "chainfour/testutil/keeper"
	"chainfour/testutil/nullify"
	"chainfour/x/chainfour"
	"chainfour/x/chainfour/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ChainfourKeeper(t)
	chainfour.InitGenesis(ctx, *k, genesisState)
	got := chainfour.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
