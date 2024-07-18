package v2

import (
	"cosmossdk.io/core/store"

	"github.com/airchains-network/cosmos-sdk/runtime"
	sdk "github.com/airchains-network/cosmos-sdk/types"
	v1 "github.com/airchains-network/cosmos-sdk/x/distribution/migrations/v1"
)

// MigrateStore performs in-place store migrations from v0.40 to v0.43. The
// migration includes:
//
// - Change addresses to be length-prefixed.
func MigrateStore(ctx sdk.Context, storeService store.KVStoreService) error {
	store := runtime.KVStoreAdapter(storeService.OpenKVStore(ctx))
	MigratePrefixAddress(store, v1.ValidatorOutstandingRewardsPrefix)
	MigratePrefixAddress(store, v1.DelegatorWithdrawAddrPrefix)
	MigratePrefixAddressAddress(store, v1.DelegatorStartingInfoPrefix)
	MigratePrefixAddressBytes(store, v1.ValidatorHistoricalRewardsPrefix)
	MigratePrefixAddress(store, v1.ValidatorCurrentRewardsPrefix)
	MigratePrefixAddress(store, v1.ValidatorAccumulatedCommissionPrefix)
	MigratePrefixAddressBytes(store, v1.ValidatorSlashEventPrefix)

	return nil
}
