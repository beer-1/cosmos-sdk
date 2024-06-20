package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

type AccountCreatorDecorator struct {
	ak AccountKeeper
}

func NewAccountCreatorDecorator(ak AccountKeeper) AccountCreatorDecorator {
	return AccountCreatorDecorator{ak}
}

func (vbd AccountCreatorDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	if !ctx.IsCheckTx() || simulate {
		return next(ctx, tx, simulate)
	}

	// only create at check tx not simulate and not finalize tx
	msgs := tx.GetMsgs()
	for _, msg := range msgs {
		if msg, ok := msg.(*banktypes.MsgSend); ok {
			addr, err := vbd.ak.AddressCodec().StringToBytes(msg.ToAddress)
			if err != nil {
				return ctx, err
			}

			acc := vbd.ak.NewAccountWithAddress(ctx, addr)
			vbd.ak.SetAccount(ctx, acc)
		}
	}

	return next(ctx, tx, simulate)
}
