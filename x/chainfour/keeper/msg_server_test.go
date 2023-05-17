package keeper_test

import (
	"context"
	"testing"

	keepertest "chainfour/testutil/keeper"
	"chainfour/x/chainfour/keeper"
	"chainfour/x/chainfour/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.ChainfourKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
