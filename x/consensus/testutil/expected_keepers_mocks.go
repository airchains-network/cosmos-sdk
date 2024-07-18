// Code generated by MockGen. DO NOT EDIT.
// Source: x/bank/types/expected_keepers.go

// Package testutil is a generated GoMock package.
package testutil

import (
	reflect "reflect"

	types "github.com/airchains-network/cosmos-sdk/types"
	types0 "github.com/airchains-network/cosmos-sdk/x/auth/types"
	gomock "github.com/golang/mock/gomock"
)

// MockAccountKeeper is a mock of AccountKeeper interface.
type MockAccountKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockAccountKeeperMockRecorder
}

// MockAccountKeeperMockRecorder is the mock recorder for MockAccountKeeper.
type MockAccountKeeperMockRecorder struct {
	mock *MockAccountKeeper
}

// NewMockAccountKeeper creates a new mock instance.
func NewMockAccountKeeper(ctrl *gomock.Controller) *MockAccountKeeper {
	mock := &MockAccountKeeper{ctrl: ctrl}
	mock.recorder = &MockAccountKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountKeeper) EXPECT() *MockAccountKeeperMockRecorder {
	return m.recorder
}

// GetAccount mocks base method.
func (m *MockAccountKeeper) GetAccount(ctx types.Context, addr types.AccAddress) types.AccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", ctx, addr)
	ret0, _ := ret[0].(types.AccountI)
	return ret0
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockAccountKeeperMockRecorder) GetAccount(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockAccountKeeper)(nil).GetAccount), ctx, addr)
}

// GetAllAccounts mocks base method.
func (m *MockAccountKeeper) GetAllAccounts(ctx types.Context) []types.AccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllAccounts", ctx)
	ret0, _ := ret[0].([]types.AccountI)
	return ret0
}

// GetAllAccounts indicates an expected call of GetAllAccounts.
func (mr *MockAccountKeeperMockRecorder) GetAllAccounts(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllAccounts", reflect.TypeOf((*MockAccountKeeper)(nil).GetAllAccounts), ctx)
}

// GetModuleAccount mocks base method.
func (m *MockAccountKeeper) GetModuleAccount(ctx types.Context, moduleName string) types0.ModuleAccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModuleAccount", ctx, moduleName)
	ret0, _ := ret[0].(types0.ModuleAccountI)
	return ret0
}

// GetModuleAccount indicates an expected call of GetModuleAccount.
func (mr *MockAccountKeeperMockRecorder) GetModuleAccount(ctx, moduleName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModuleAccount", reflect.TypeOf((*MockAccountKeeper)(nil).GetModuleAccount), ctx, moduleName)
}

// GetModuleAccountAndPermissions mocks base method.
func (m *MockAccountKeeper) GetModuleAccountAndPermissions(ctx types.Context, moduleName string) (types0.ModuleAccountI, []string) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModuleAccountAndPermissions", ctx, moduleName)
	ret0, _ := ret[0].(types0.ModuleAccountI)
	ret1, _ := ret[1].([]string)
	return ret0, ret1
}

// GetModuleAccountAndPermissions indicates an expected call of GetModuleAccountAndPermissions.
func (mr *MockAccountKeeperMockRecorder) GetModuleAccountAndPermissions(ctx, moduleName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModuleAccountAndPermissions", reflect.TypeOf((*MockAccountKeeper)(nil).GetModuleAccountAndPermissions), ctx, moduleName)
}

// GetModuleAddress mocks base method.
func (m *MockAccountKeeper) GetModuleAddress(moduleName string) types.AccAddress {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModuleAddress", moduleName)
	ret0, _ := ret[0].(types.AccAddress)
	return ret0
}

// GetModuleAddress indicates an expected call of GetModuleAddress.
func (mr *MockAccountKeeperMockRecorder) GetModuleAddress(moduleName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModuleAddress", reflect.TypeOf((*MockAccountKeeper)(nil).GetModuleAddress), moduleName)
}

// GetModuleAddressAndPermissions mocks base method.
func (m *MockAccountKeeper) GetModuleAddressAndPermissions(moduleName string) (types.AccAddress, []string) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModuleAddressAndPermissions", moduleName)
	ret0, _ := ret[0].(types.AccAddress)
	ret1, _ := ret[1].([]string)
	return ret0, ret1
}

// GetModuleAddressAndPermissions indicates an expected call of GetModuleAddressAndPermissions.
func (mr *MockAccountKeeperMockRecorder) GetModuleAddressAndPermissions(moduleName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModuleAddressAndPermissions", reflect.TypeOf((*MockAccountKeeper)(nil).GetModuleAddressAndPermissions), moduleName)
}

// GetModulePermissions mocks base method.
func (m *MockAccountKeeper) GetModulePermissions() map[string]types0.PermissionsForAddress {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModulePermissions")
	ret0, _ := ret[0].(map[string]types0.PermissionsForAddress)
	return ret0
}

// GetModulePermissions indicates an expected call of GetModulePermissions.
func (mr *MockAccountKeeperMockRecorder) GetModulePermissions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModulePermissions", reflect.TypeOf((*MockAccountKeeper)(nil).GetModulePermissions))
}

// HasAccount mocks base method.
func (m *MockAccountKeeper) HasAccount(ctx types.Context, addr types.AccAddress) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasAccount", ctx, addr)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasAccount indicates an expected call of HasAccount.
func (mr *MockAccountKeeperMockRecorder) HasAccount(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasAccount", reflect.TypeOf((*MockAccountKeeper)(nil).HasAccount), ctx, addr)
}

// IterateAccounts mocks base method.
func (m *MockAccountKeeper) IterateAccounts(ctx types.Context, process func(types.AccountI) bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IterateAccounts", ctx, process)
}

// IterateAccounts indicates an expected call of IterateAccounts.
func (mr *MockAccountKeeperMockRecorder) IterateAccounts(ctx, process interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IterateAccounts", reflect.TypeOf((*MockAccountKeeper)(nil).IterateAccounts), ctx, process)
}

// NewAccount mocks base method.
func (m *MockAccountKeeper) NewAccount(arg0 types.Context, arg1 types.AccountI) types.AccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewAccount", arg0, arg1)
	ret0, _ := ret[0].(types.AccountI)
	return ret0
}

// NewAccount indicates an expected call of NewAccount.
func (mr *MockAccountKeeperMockRecorder) NewAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAccount", reflect.TypeOf((*MockAccountKeeper)(nil).NewAccount), arg0, arg1)
}

// NewAccountWithAddress mocks base method.
func (m *MockAccountKeeper) NewAccountWithAddress(ctx types.Context, addr types.AccAddress) types.AccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewAccountWithAddress", ctx, addr)
	ret0, _ := ret[0].(types.AccountI)
	return ret0
}

// NewAccountWithAddress indicates an expected call of NewAccountWithAddress.
func (mr *MockAccountKeeperMockRecorder) NewAccountWithAddress(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAccountWithAddress", reflect.TypeOf((*MockAccountKeeper)(nil).NewAccountWithAddress), ctx, addr)
}

// SetAccount mocks base method.
func (m *MockAccountKeeper) SetAccount(ctx types.Context, acc types.AccountI) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAccount", ctx, acc)
}

// SetAccount indicates an expected call of SetAccount.
func (mr *MockAccountKeeperMockRecorder) SetAccount(ctx, acc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAccount", reflect.TypeOf((*MockAccountKeeper)(nil).SetAccount), ctx, acc)
}

// SetModuleAccount mocks base method.
func (m *MockAccountKeeper) SetModuleAccount(ctx types.Context, macc types0.ModuleAccountI) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetModuleAccount", ctx, macc)
}

// SetModuleAccount indicates an expected call of SetModuleAccount.
func (mr *MockAccountKeeperMockRecorder) SetModuleAccount(ctx, macc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetModuleAccount", reflect.TypeOf((*MockAccountKeeper)(nil).SetModuleAccount), ctx, macc)
}

// ValidatePermissions mocks base method.
func (m *MockAccountKeeper) ValidatePermissions(macc types0.ModuleAccountI) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidatePermissions", macc)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidatePermissions indicates an expected call of ValidatePermissions.
func (mr *MockAccountKeeperMockRecorder) ValidatePermissions(macc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidatePermissions", reflect.TypeOf((*MockAccountKeeper)(nil).ValidatePermissions), macc)
}
