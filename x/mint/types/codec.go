package types

import (
	"github.com/airchains-network/cosmos-sdk/codec"
	"github.com/airchains-network/cosmos-sdk/codec/legacy"
	"github.com/airchains-network/cosmos-sdk/codec/types"
	sdk "github.com/airchains-network/cosmos-sdk/types"
	"github.com/airchains-network/cosmos-sdk/types/msgservice"
)

// RegisterLegacyAminoCodec registers concrete types on the LegacyAmino codec
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(Params{}, "cosmos-sdk/x/mint/Params", nil)
	legacy.RegisterAminoMsg(cdc, &MsgUpdateParams{}, "cosmos-sdk/x/mint/MsgUpdateParams")
}

// RegisterInterfaces registers the interfaces types with the interface registry.
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgUpdateParams{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
