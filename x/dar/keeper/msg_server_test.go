package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/jaecobd1/dar/testutil/keeper"
	"github.com/jaecobd1/dar/x/dar/keeper"
	"github.com/jaecobd1/dar/x/dar/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.DarKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
