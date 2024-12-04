package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetMechanism{}

func NewMsgSetMechanism(creator string, metrics string, coefficients string, convergeLimit string, slope string) *MsgSetMechanism {
	return &MsgSetMechanism{
		Creator:       creator,
		Metrics:       metrics,
		Coefficients:  coefficients,
		ConvergeLimit: convergeLimit,
		Slope:         slope,
	}
}

func (msg *MsgSetMechanism) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
