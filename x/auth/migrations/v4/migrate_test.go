package v4_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	storetypes "cosmossdk.io/store/types"

	"github.com/airchains-network/cosmos-sdk/runtime"
	"github.com/airchains-network/cosmos-sdk/testutil"
	sdk "github.com/airchains-network/cosmos-sdk/types"
	moduletestutil "github.com/airchains-network/cosmos-sdk/types/module/testutil"
	"github.com/airchains-network/cosmos-sdk/x/auth"
	"github.com/airchains-network/cosmos-sdk/x/auth/exported"
	v1 "github.com/airchains-network/cosmos-sdk/x/auth/migrations/v1"
	v4 "github.com/airchains-network/cosmos-sdk/x/auth/migrations/v4"
	"github.com/airchains-network/cosmos-sdk/x/auth/types"
)

type mockSubspace struct {
	ps types.Params
}

func newMockSubspace(ps types.Params) mockSubspace {
	return mockSubspace{ps: ps}
}

func (ms mockSubspace) GetParamSet(ctx sdk.Context, ps exported.ParamSet) {
	*ps.(*types.Params) = ms.ps
}

func TestMigrate(t *testing.T) {
	encCfg := moduletestutil.MakeTestEncodingConfig(auth.AppModuleBasic{})
	cdc := encCfg.Codec

	storeKey := storetypes.NewKVStoreKey(v1.ModuleName)
	tKey := storetypes.NewTransientStoreKey("transient_test")
	ctx := testutil.DefaultContext(storeKey, tKey)
	storeService := runtime.NewKVStoreService(storeKey)

	legacySubspace := newMockSubspace(types.DefaultParams())
	require.NoError(t, v4.Migrate(ctx, storeService, legacySubspace, cdc))

	var res types.Params
	bz, err := storeService.OpenKVStore(ctx).Get(v4.ParamsKey)
	require.NoError(t, err)
	require.NoError(t, cdc.Unmarshal(bz, &res))
	require.Equal(t, legacySubspace.ps, res)
}
