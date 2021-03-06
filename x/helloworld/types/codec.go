package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgGet{}, "hello/GetMessage", nil)
	cdc.RegisterConcrete(MsgSet{}, "hello/SetMessage", nil)
	cdc.RegisterConcrete(MsgPropose{}, "hello/BuyMessage", nil)
	cdc.RegisterConcrete(MsgSell{}, "hello/SellMessage", nil)
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
