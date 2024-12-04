package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCalculatePerformance{}

func NewMsgCalculatePerformance(creator string, mid int32, wager sdk.Coin, earner string, taskid string) *MsgCalculatePerformance {
	return &MsgCalculatePerformance{
		Creator: creator,
		Mid:     mid,
		Wager:   wager,
		Earner:  earner,
		Taskid:  taskid,
	}
}

func (msg *MsgCalculatePerformance) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
