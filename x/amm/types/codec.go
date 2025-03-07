package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreatePool{}, "amm/CreatePool", nil)
	cdc.RegisterConcrete(&MsgJoinPool{}, "amm/JoinPool", nil)
	cdc.RegisterConcrete(&MsgExitPool{}, "amm/ExitPool", nil)
	cdc.RegisterConcrete(&MsgSwapExactAmountIn{}, "amm/SwapExactAmountIn", nil)
	cdc.RegisterConcrete(&MsgSwapExactAmountOut{}, "amm/SwapExactAmountOut", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreatePool{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgJoinPool{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgExitPool{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSwapExactAmountIn{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSwapExactAmountOut{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
