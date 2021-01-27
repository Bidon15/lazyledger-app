package lazyledgerapp

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lazyledger/lazyledger-app/x/lazyledgerapp/keeper"
	"github.com/lazyledger/lazyledger-app/x/lazyledgerapp/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
        // this line is used by starport scaffolding # 1
case *types.MsgCreatePayForMessage:
	return handleMsgCreatePayForMessage(ctx, k, msg)

case *types.MsgUpdatePayForMessage:
	return handleMsgUpdatePayForMessage(ctx, k, msg)

case *types.MsgDeletePayForMessage:
	return handleMsgDeletePayForMessage(ctx, k, msg)

		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
