package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateBet = "create_bet"
	TypeMsgUpdateBet = "update_bet"
	TypeMsgDeleteBet = "delete_bet"
)

var _ sdk.Msg = &MsgCreateBet{}

func NewMsgCreateBet(
	creator string,
	amount uint64,

) *MsgCreateBet {
	return &MsgCreateBet{
		Creator: creator,
		Amount:  amount,
	}
}

func (msg *MsgCreateBet) Route() string {
	return RouterKey
}

func (msg *MsgCreateBet) Type() string {
	return TypeMsgCreateBet
}

func (msg *MsgCreateBet) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateBet) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateBet) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
