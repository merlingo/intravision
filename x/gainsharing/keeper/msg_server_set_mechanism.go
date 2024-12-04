package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/merlingo/intravision/x/gainsharing/types"
)

func (k msgServer) SetMechanism(goCtx context.Context, msg *types.MsgSetMechanism) (*types.MsgSetMechanismResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: validate values and append Mechanism object
	err := msg.ValidateBasic()
	if err != nil {
		return nil, err
	}
	mechanism := types.Mechanism{
		Metrics:       msg.Metrics,
		Slope:         msg.Slope,
		Coefficients:  msg.Coefficients,
		ConvergeLimit: msg.ConvergeLimit,
	}
	err = mechanism.Validate()
	if err != nil {
		return nil, err
	}
	id := k.Keeper.AppendMechanism(ctx, mechanism)
	ctx.GasMeter().ConsumeGas(types.SetNewMechanism, "Create New Mechanism")

	return &types.MsgSetMechanismResponse{Id: int32(id)}, nil
}
