package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/merlingo/intravision/testutil/keeper"
	"github.com/merlingo/intravision/x/intravision/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.IntravisionKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
