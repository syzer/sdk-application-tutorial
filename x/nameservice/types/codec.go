package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

var ModuleCdc = codec.New()

func init() {
	RegisterCodec(ModuleCdc)
}

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgSetDAG{}, "nameservice/SetName", nil)
	cdc.RegisterConcrete(MsgBuyLocation{}, "nameservice/BuyName", nil)
}
