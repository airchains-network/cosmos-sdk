package types

import (
	"github.com/airchains-network/cosmos-sdk/codec"
	"github.com/airchains-network/cosmos-sdk/codec/legacy"
	"github.com/airchains-network/cosmos-sdk/codec/types"
	sdk "github.com/airchains-network/cosmos-sdk/types"
	"github.com/airchains-network/cosmos-sdk/types/msgservice"
)

// RegisterLegacyAminoCodec registers concrete types on LegacyAmino codec
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(Params{}, "cosmos-sdk/x/slashing/Params", nil)
	legacy.RegisterAminoMsg(cdc, &MsgUnjail{}, "cosmos-sdk/MsgUnjail")
	legacy.RegisterAminoMsg(cdc, &MsgUpdateParams{}, "cosmos-sdk/x/slashing/MsgUpdateParams")
}

// RegisterInterfaces registers the interfaces types with the Interface Registry.
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUnjail{},
		&MsgUpdateParams{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
