package testutil

import (
	"testing"

	"github.com/stretchr/testify/require"

	cryptotypes "github.com/airchains-network/cosmos-sdk/crypto/types"
	sdk "github.com/airchains-network/cosmos-sdk/types"
	"github.com/airchains-network/cosmos-sdk/x/staking/types"
)

// NewValidator is a testing helper method to create validators in tests
func NewValidator(tb testing.TB, operator sdk.ValAddress, pubKey cryptotypes.PubKey) types.Validator {
	tb.Helper()
	v, err := types.NewValidator(operator.String(), pubKey, types.Description{})
	require.NoError(tb, err)
	return v
}
