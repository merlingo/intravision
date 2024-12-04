package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/merlingo/intravision/x/gainsharing/types"
)

func (k msgServer) CalculatePerformance(goCtx context.Context, msg *types.MsgCalculatePerformance) (*types.MsgCalculatePerformanceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: validate values, get related mechanism, calculate the reward,
	//create and append the performance, send reward to worker, return id and reward value
	err := msg.ValidateBasic()
	if err != nil {
		return nil, err
	}
	mechanism, found := k.Keeper.GetMechanism(ctx, uint64(msg.Mid))
	if !found {
		return nil, types.ErrMechanismNotFound
	}
	activities := k.labourKeeper.GetTaskActivities(ctx, msg.Earner, msg.Taskid)
	if len(activities) == 0 {
		return nil, types.ErrActivitiesNotFound
	}
	tid, err := strconv.ParseUint(msg.Taskid, 10, 64)
	if err != nil {
		return nil, types.ErrActivitiesNotFound
	}
	task, found := k.labourKeeper.GetTask(ctx, tid)
	if !found {
		return nil, types.ErrActivitiesNotFound
	}
	rewards := mechanism.CalculateRewards(msg.Wager, activities, task)

	//pay rewards to earner account
	//rewards := msg.Wager
	performance := types.Performance{
		Reward: rewards,
		Mid:    msg.Mid,
		Tid:    msg.Taskid,
		Wager:  msg.Wager,
		Earner: msg.Earner,
	}

	//err = performance.Validate()
	//if err != nil {
	//	return nil, err
	//}

	id := k.Keeper.AppendPerformance(ctx, performance)
	//send token to earner account
	err = k.bankEscrowKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.AccAddress(performance.Earner), sdk.NewCoins(rewards))
	if err != nil {
		return nil, types.ErrRewardFunding
	}
	ctx.GasMeter().ConsumeGas(types.PerformanceGas, "Calculate Performance")

	return &types.MsgCalculatePerformanceResponse{Id: int32(id), Reward: rewards}, nil
}
