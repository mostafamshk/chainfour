package keeper_test

import (
	"testing"

	testkeeper "chainfour/testutil/keeper"
	"chainfour/x/chainfour/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ChainfourKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
