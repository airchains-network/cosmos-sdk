package cmd

import (
	"context"
	"fmt"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"

	"cosmossdk.io/log"

	"github.com/airchains-network/cosmos-sdk/client"
	"github.com/airchains-network/cosmos-sdk/client/flags"
	"github.com/airchains-network/cosmos-sdk/server"
	"github.com/airchains-network/cosmos-sdk/types/module"
	moduletestutil "github.com/airchains-network/cosmos-sdk/types/module/testutil"
	"github.com/airchains-network/cosmos-sdk/x/auth"
	"github.com/airchains-network/cosmos-sdk/x/bank"
	banktypes "github.com/airchains-network/cosmos-sdk/x/bank/types"
	"github.com/airchains-network/cosmos-sdk/x/consensus"
	"github.com/airchains-network/cosmos-sdk/x/distribution"
	"github.com/airchains-network/cosmos-sdk/x/genutil"
	genutiltest "github.com/airchains-network/cosmos-sdk/x/genutil/client/testutil"
	genutiltypes "github.com/airchains-network/cosmos-sdk/x/genutil/types"
	"github.com/airchains-network/cosmos-sdk/x/mint"
	"github.com/airchains-network/cosmos-sdk/x/params"
	"github.com/airchains-network/cosmos-sdk/x/staking"
)

func Test_TestnetCmd(t *testing.T) {
	moduleBasic := module.NewBasicManager(
		auth.AppModuleBasic{},
		genutil.NewAppModuleBasic(genutiltypes.DefaultMessageValidator),
		bank.AppModuleBasic{},
		staking.AppModuleBasic{},
		mint.AppModuleBasic{},
		distribution.AppModuleBasic{},
		params.AppModuleBasic{},
		consensus.AppModuleBasic{},
	)

	home := t.TempDir()
	encodingConfig := moduletestutil.MakeTestEncodingConfig(auth.AppModuleBasic{}, staking.AppModuleBasic{})
	logger := log.NewNopLogger()
	cfg, err := genutiltest.CreateDefaultCometConfig(home)
	require.NoError(t, err)

	err = genutiltest.ExecInitCmd(moduleBasic, home, encodingConfig.Codec)
	require.NoError(t, err)

	serverCtx := server.NewContext(viper.New(), cfg, logger)
	clientCtx := client.Context{}.
		WithCodec(encodingConfig.Codec).
		WithHomeDir(home).
		WithTxConfig(encodingConfig.TxConfig)

	ctx := context.Background()
	ctx = context.WithValue(ctx, server.ServerContextKey, serverCtx)
	ctx = context.WithValue(ctx, client.ClientContextKey, &clientCtx)
	cmd := testnetInitFilesCmd(moduleBasic, banktypes.GenesisBalancesIterator{})
	cmd.SetArgs([]string{fmt.Sprintf("--%s=test", flags.FlagKeyringBackend), fmt.Sprintf("--output-dir=%s", home)})
	err = cmd.ExecuteContext(ctx)
	require.NoError(t, err)

	genFile := cfg.GenesisFile()
	appState, _, err := genutiltypes.GenesisStateFromGenFile(genFile)
	require.NoError(t, err)

	bankGenState := banktypes.GetGenesisStateFromAppState(encodingConfig.Codec, appState)
	require.NotEmpty(t, bankGenState.Supply.String())
}
