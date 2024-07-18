package runtime

import (
	abci "github.com/cometbft/cometbft/abci/types"

	"github.com/airchains-network/cosmos-sdk/codec"
	"github.com/airchains-network/cosmos-sdk/server/types"
	sdk "github.com/airchains-network/cosmos-sdk/types"
	"github.com/airchains-network/cosmos-sdk/types/module"
)

const ModuleName = "runtime"

// App implements the common methods for a Cosmos SDK-based application
// specific blockchain.
type AppI interface {
	// The assigned name of the app.
	Name() string

	// The application types codec.
	// NOTE: This should NOT be sealed before being returned.
	LegacyAmino() *codec.LegacyAmino

	// Application updates every begin block.
	BeginBlocker(ctx sdk.Context) (sdk.BeginBlock, error)

	// Application updates every end block.
	EndBlocker(ctx sdk.Context) (sdk.EndBlock, error)

	// Application update at chain (i.e app) initialization.
	InitChainer(ctx sdk.Context, req *abci.RequestInitChain) (*abci.ResponseInitChain, error)

	// Loads the app at a given height.
	LoadHeight(height int64) error

	// Exports the state of the application for a genesis file.
	ExportAppStateAndValidators(forZeroHeight bool, jailAllowedAddrs, modulesToExport []string) (types.ExportedApp, error)

	// Helper for the simulation framework.
	SimulationManager() *module.SimulationManager
}
