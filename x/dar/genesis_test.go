package dar_test

import (
	"testing"

	keepertest "github.com/jaecobd1/dar/testutil/keeper"
	"github.com/jaecobd1/dar/testutil/nullify"
	"github.com/jaecobd1/dar/x/dar"
	"github.com/jaecobd1/dar/x/dar/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DarKeeper(t)
	dar.InitGenesis(ctx, *k, genesisState)
	got := dar.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
