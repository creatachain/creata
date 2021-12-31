package params

import (
	"github.com/creatachain/creata-sdk/client"
	"github.com/creatachain/creata-sdk/codec"
	"github.com/creatachain/creata-sdk/codec/types"
)

// EncodingConfig specifies the concrete encoding types to use for a given app.
// This is provided for compatibility between protobuf and amino implementations.
type EncodingConfig struct {
	InterfaceRegistry types.InterfaceRegistry
	Marshaler         codec.Marshaler
	TxConfig          client.TxConfig
	Amino             *codec.LegacyAmino
}
