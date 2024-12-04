package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/merlingo/intravision/testutil/keeper"
	"github.com/merlingo/intravision/x/labour/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.LabourKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
