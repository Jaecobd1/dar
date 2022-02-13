package keeper_test

import (
	"testing"

	testkeeper "github.com/jaecobd1/dar/testutil/keeper"
	"github.com/jaecobd1/dar/x/dar/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.DarKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
