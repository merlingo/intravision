package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/merlingo/intravision/testutil/keeper"
	"github.com/merlingo/intravision/x/labour/keeper"
	labour "github.com/merlingo/intravision/x/labour/module"
	"github.com/merlingo/intravision/x/labour/types"
	"github.com/stretchr/testify/require"
)

func TestBesginTask(t *testing.T) {
	msgServer, _, context := setupMsgServerBeginTask(t)
	createResponse, err := msgServer.BeginTask(context, &types.MsgBeginTask{
		Creator:   alice,
		Taskid:    taskId,
		Assigner:  alice,
		BeginTask: 1729687287000,
		Deadline:  1732365735000,
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgBeginTaskResponse{
		// TODO: update with a proper value when updated
		Id: 0,
	}, *createResponse)
}
func setupMsgServerBeginTask(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
	k, ctx := keepertest.LabourKeeper(t)
	labour.InitGenesis(ctx, k, *types.DefaultGenesis())
	return keeper.NewMsgServerImpl(k), k, sdk.WrapSDKContext(ctx)
}
