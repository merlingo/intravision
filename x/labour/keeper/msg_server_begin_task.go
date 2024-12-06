package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/merlingo/intravision/x/labour/types"
)

func (k msgServer) BeginTask(goCtx context.Context, msg *types.MsgBeginTask) (*types.MsgBeginTaskResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// 1- build Task object  2. validate it  3. store the object 4. emit the event
	err := msg.ValidateBasic()
	if err != nil {
		return nil, err
	}
	task := types.Task{
		TaskId:    msg.GetTaskid(),
		Assigner:  msg.GetAssigner(),
		State:     1,
		BeginTask: msg.GetBeginTask(),
		Deadline:  msg.GetDeadline(),
		Wager:     msg.GetWager(),
	}
	err = task.Validate()
	if err != nil {
		return nil, err
	}
	wager, _ := sdk.ParseCoinsNormalized(msg.Wager)
	if !wager.IsValid() {
		return nil, errorsmod.Wrap(types.ErrSample, "amount is not a valid Coins object")
	}
	id := k.Keeper.AppendTask(ctx, task)
	ctx.GasMeter().ConsumeGas(types.BeginTaskGas, "Create New Task")

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, sdk.AccAddress(task.Assigner), types.ModuleName, wager)
	if err != nil {
		return nil, err
	}

	return &types.MsgBeginTaskResponse{Id: id}, nil
}
