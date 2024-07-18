package authz

import (
	"github.com/airchains-network/cosmos-sdk/client"
	addresscodec "github.com/airchains-network/cosmos-sdk/codec/address"
	"github.com/airchains-network/cosmos-sdk/testutil"
	clitestutil "github.com/airchains-network/cosmos-sdk/testutil/cli"
	"github.com/airchains-network/cosmos-sdk/x/authz/client/cli"
)

func CreateGrant(clientCtx client.Context, args []string) (testutil.BufferWriter, error) {
	cmd := cli.NewCmdGrantAuthorization(addresscodec.NewBech32Codec("cosmos"))
	return clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
}
