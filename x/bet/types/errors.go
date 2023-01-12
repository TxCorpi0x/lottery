package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/bet module sentinel errors
var (
	ErrInvalidAmount = sdkerrors.Register(ModuleName, 110, "bet amount is not valid")
)
