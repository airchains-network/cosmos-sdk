package keeper_test

import (
	gocontext "context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"

	storetypes "cosmossdk.io/store/types"

	"github.com/airchains-network/cosmos-sdk/baseapp"
	"github.com/airchains-network/cosmos-sdk/runtime"
	"github.com/airchains-network/cosmos-sdk/testutil"
	sdk "github.com/airchains-network/cosmos-sdk/types"
	moduletestutil "github.com/airchains-network/cosmos-sdk/types/module/testutil"
	authtypes "github.com/airchains-network/cosmos-sdk/x/auth/types"
	govtypes "github.com/airchains-network/cosmos-sdk/x/gov/types"
	"github.com/airchains-network/cosmos-sdk/x/mint"
	"github.com/airchains-network/cosmos-sdk/x/mint/keeper"
	minttestutil "github.com/airchains-network/cosmos-sdk/x/mint/testutil"
	"github.com/airchains-network/cosmos-sdk/x/mint/types"
)

type MintTestSuite struct {
	suite.Suite

	ctx         sdk.Context
	queryClient types.QueryClient
	mintKeeper  keeper.Keeper
}

func (suite *MintTestSuite) SetupTest() {
	encCfg := moduletestutil.MakeTestEncodingConfig(mint.AppModuleBasic{})
	key := storetypes.NewKVStoreKey(types.StoreKey)
	storeService := runtime.NewKVStoreService(key)
	testCtx := testutil.DefaultContextWithDB(suite.T(), key, storetypes.NewTransientStoreKey("transient_test"))
	suite.ctx = testCtx.Ctx

	// gomock initializations
	ctrl := gomock.NewController(suite.T())
	accountKeeper := minttestutil.NewMockAccountKeeper(ctrl)
	bankKeeper := minttestutil.NewMockBankKeeper(ctrl)
	stakingKeeper := minttestutil.NewMockStakingKeeper(ctrl)

	accountKeeper.EXPECT().GetModuleAddress("mint").Return(sdk.AccAddress{})

	suite.mintKeeper = keeper.NewKeeper(
		encCfg.Codec,
		storeService,
		stakingKeeper,
		accountKeeper,
		bankKeeper,
		authtypes.FeeCollectorName,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	err := suite.mintKeeper.Params.Set(suite.ctx, types.DefaultParams())
	suite.Require().NoError(err)
	suite.Require().NoError(suite.mintKeeper.Minter.Set(suite.ctx, types.DefaultInitialMinter()))

	queryHelper := baseapp.NewQueryServerTestHelper(testCtx.Ctx, encCfg.InterfaceRegistry)
	types.RegisterQueryServer(queryHelper, keeper.NewQueryServerImpl(suite.mintKeeper))

	suite.queryClient = types.NewQueryClient(queryHelper)
}

func (suite *MintTestSuite) TestGRPCParams() {
	params, err := suite.queryClient.Params(gocontext.Background(), &types.QueryParamsRequest{})
	suite.Require().NoError(err)
	kparams, err := suite.mintKeeper.Params.Get(suite.ctx)
	suite.Require().NoError(err)
	suite.Require().Equal(params.Params, kparams)

	inflation, err := suite.queryClient.Inflation(gocontext.Background(), &types.QueryInflationRequest{})
	suite.Require().NoError(err)
	minter, err := suite.mintKeeper.Minter.Get(suite.ctx)
	suite.Require().NoError(err)
	suite.Require().Equal(inflation.Inflation, minter.Inflation)

	annualProvisions, err := suite.queryClient.AnnualProvisions(gocontext.Background(), &types.QueryAnnualProvisionsRequest{})
	suite.Require().NoError(err)
	suite.Require().Equal(annualProvisions.AnnualProvisions, minter.AnnualProvisions)
}

func TestMintTestSuite(t *testing.T) {
	suite.Run(t, new(MintTestSuite))
}
