package keeper

import (
	"chainfour/x/chainfour/types"
)

var _ types.QueryServer = Keeper{}
