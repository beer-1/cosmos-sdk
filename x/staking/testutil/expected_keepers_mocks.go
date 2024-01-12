// Code generated by MockGen. DO NOT EDIT.
// Source: x/staking/types/expected_keepers.go

// Package testutil is a generated GoMock package.
package testutil

import (
	reflect "reflect"

	math "cosmossdk.io/math"
	types "github.com/cosmos/cosmos-sdk/types"
	types0 "github.com/cosmos/cosmos-sdk/x/auth/types"
	types1 "github.com/cosmos/cosmos-sdk/x/staking/types"
	gomock "github.com/golang/mock/gomock"
)

// MockDistributionKeeper is a mock of DistributionKeeper interface.
type MockDistributionKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockDistributionKeeperMockRecorder
}

// MockDistributionKeeperMockRecorder is the mock recorder for MockDistributionKeeper.
type MockDistributionKeeperMockRecorder struct {
	mock *MockDistributionKeeper
}

// NewMockDistributionKeeper creates a new mock instance.
func NewMockDistributionKeeper(ctrl *gomock.Controller) *MockDistributionKeeper {
	mock := &MockDistributionKeeper{ctrl: ctrl}
	mock.recorder = &MockDistributionKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDistributionKeeper) EXPECT() *MockDistributionKeeperMockRecorder {
	return m.recorder
}

// GetFeePoolCommunityCoins mocks base method.
func (m *MockDistributionKeeper) GetFeePoolCommunityCoins(ctx types.Context) types.DecCoins {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFeePoolCommunityCoins", ctx)
	ret0, _ := ret[0].(types.DecCoins)
	return ret0
}

// GetFeePoolCommunityCoins indicates an expected call of GetFeePoolCommunityCoins.
func (mr *MockDistributionKeeperMockRecorder) GetFeePoolCommunityCoins(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFeePoolCommunityCoins", reflect.TypeOf((*MockDistributionKeeper)(nil).GetFeePoolCommunityCoins), ctx)
}

// GetValidatorOutstandingRewardsCoins mocks base method.
func (m *MockDistributionKeeper) GetValidatorOutstandingRewardsCoins(ctx types.Context, val types.ValAddress) types.DecCoins {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidatorOutstandingRewardsCoins", ctx, val)
	ret0, _ := ret[0].(types.DecCoins)
	return ret0
}

// GetValidatorOutstandingRewardsCoins indicates an expected call of GetValidatorOutstandingRewardsCoins.
func (mr *MockDistributionKeeperMockRecorder) GetValidatorOutstandingRewardsCoins(ctx, val interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidatorOutstandingRewardsCoins", reflect.TypeOf((*MockDistributionKeeper)(nil).GetValidatorOutstandingRewardsCoins), ctx, val)
}

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
func (m *MockAccountKeeper) GetAccount(ctx types.Context, addr types.AccAddress) types0.AccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", ctx, addr)
	ret0, _ := ret[0].(types0.AccountI)
	return ret0
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockAccountKeeperMockRecorder) GetAccount(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockAccountKeeper)(nil).GetAccount), ctx, addr)
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

// GetModuleAddress mocks base method.
func (m *MockAccountKeeper) GetModuleAddress(name string) types.AccAddress {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModuleAddress", name)
	ret0, _ := ret[0].(types.AccAddress)
	return ret0
}

// GetModuleAddress indicates an expected call of GetModuleAddress.
func (mr *MockAccountKeeperMockRecorder) GetModuleAddress(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModuleAddress", reflect.TypeOf((*MockAccountKeeper)(nil).GetModuleAddress), name)
}

// IterateAccounts mocks base method.
func (m *MockAccountKeeper) IterateAccounts(ctx types.Context, process func(types0.AccountI) bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IterateAccounts", ctx, process)
}

// IterateAccounts indicates an expected call of IterateAccounts.
func (mr *MockAccountKeeperMockRecorder) IterateAccounts(ctx, process interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IterateAccounts", reflect.TypeOf((*MockAccountKeeper)(nil).IterateAccounts), ctx, process)
}

// SetModuleAccount mocks base method.
func (m *MockAccountKeeper) SetModuleAccount(arg0 types.Context, arg1 types0.ModuleAccountI) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetModuleAccount", arg0, arg1)
}

// SetModuleAccount indicates an expected call of SetModuleAccount.
func (mr *MockAccountKeeperMockRecorder) SetModuleAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetModuleAccount", reflect.TypeOf((*MockAccountKeeper)(nil).SetModuleAccount), arg0, arg1)
}

// MockBankKeeper is a mock of BankKeeper interface.
type MockBankKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockBankKeeperMockRecorder
}

// MockBankKeeperMockRecorder is the mock recorder for MockBankKeeper.
type MockBankKeeperMockRecorder struct {
	mock *MockBankKeeper
}

// NewMockBankKeeper creates a new mock instance.
func NewMockBankKeeper(ctrl *gomock.Controller) *MockBankKeeper {
	mock := &MockBankKeeper{ctrl: ctrl}
	mock.recorder = &MockBankKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBankKeeper) EXPECT() *MockBankKeeperMockRecorder {
	return m.recorder
}

// BurnCoins mocks base method.
func (m *MockBankKeeper) BurnCoins(ctx types.Context, name string, amt types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BurnCoins", ctx, name, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// BurnCoins indicates an expected call of BurnCoins.
func (mr *MockBankKeeperMockRecorder) BurnCoins(ctx, name, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BurnCoins", reflect.TypeOf((*MockBankKeeper)(nil).BurnCoins), ctx, name, amt)
}

// DelegateCoinsFromAccountToModule mocks base method.
func (m *MockBankKeeper) DelegateCoinsFromAccountToModule(ctx types.Context, senderAddr types.AccAddress, recipientModule string, amt types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DelegateCoinsFromAccountToModule", ctx, senderAddr, recipientModule, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// DelegateCoinsFromAccountToModule indicates an expected call of DelegateCoinsFromAccountToModule.
func (mr *MockBankKeeperMockRecorder) DelegateCoinsFromAccountToModule(ctx, senderAddr, recipientModule, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DelegateCoinsFromAccountToModule", reflect.TypeOf((*MockBankKeeper)(nil).DelegateCoinsFromAccountToModule), ctx, senderAddr, recipientModule, amt)
}

// GetAllBalances mocks base method.
func (m *MockBankKeeper) GetAllBalances(ctx types.Context, addr types.AccAddress) types.Coins {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllBalances", ctx, addr)
	ret0, _ := ret[0].(types.Coins)
	return ret0
}

// GetAllBalances indicates an expected call of GetAllBalances.
func (mr *MockBankKeeperMockRecorder) GetAllBalances(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllBalances", reflect.TypeOf((*MockBankKeeper)(nil).GetAllBalances), ctx, addr)
}

// GetBalance mocks base method.
func (m *MockBankKeeper) GetBalance(ctx types.Context, addr types.AccAddress, denom string) types.Coin {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalance", ctx, addr, denom)
	ret0, _ := ret[0].(types.Coin)
	return ret0
}

// GetBalance indicates an expected call of GetBalance.
func (mr *MockBankKeeperMockRecorder) GetBalance(ctx, addr, denom interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalance", reflect.TypeOf((*MockBankKeeper)(nil).GetBalance), ctx, addr, denom)
}

// GetSupply mocks base method.
func (m *MockBankKeeper) GetSupply(ctx types.Context, denom string) types.Coin {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSupply", ctx, denom)
	ret0, _ := ret[0].(types.Coin)
	return ret0
}

// GetSupply indicates an expected call of GetSupply.
func (mr *MockBankKeeperMockRecorder) GetSupply(ctx, denom interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSupply", reflect.TypeOf((*MockBankKeeper)(nil).GetSupply), ctx, denom)
}

// LockedCoins mocks base method.
func (m *MockBankKeeper) LockedCoins(ctx types.Context, addr types.AccAddress) types.Coins {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LockedCoins", ctx, addr)
	ret0, _ := ret[0].(types.Coins)
	return ret0
}

// LockedCoins indicates an expected call of LockedCoins.
func (mr *MockBankKeeperMockRecorder) LockedCoins(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LockedCoins", reflect.TypeOf((*MockBankKeeper)(nil).LockedCoins), ctx, addr)
}

// MintCoins mocks base method.
func (m *MockBankKeeper) MintCoins(cts types.Context, name string, amt types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MintCoins", cts, name, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// MintCoins indicates an expected call of MintCoins.
func (mr *MockBankKeeperMockRecorder) MintCoins(cts, name, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MintCoins", reflect.TypeOf((*MockBankKeeper)(nil).MintCoins), cts, name, amt)
}

// SendCoins mocks base method.
func (m *MockBankKeeper) SendCoins(ctx types.Context, fromAddr, toAddr types.AccAddress, amt types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoins", ctx, fromAddr, toAddr, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoins indicates an expected call of SendCoins.
func (mr *MockBankKeeperMockRecorder) SendCoins(ctx, fromAddr, toAddr, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoins", reflect.TypeOf((*MockBankKeeper)(nil).SendCoins), ctx, fromAddr, toAddr, amt)
}

// SendCoinsFromAccountToModule mocks base method.
func (m *MockBankKeeper) SendCoinsFromAccountToModule(ctx types.Context, senderAddr types.AccAddress, recipientModule string, amt types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromAccountToModule", ctx, senderAddr, recipientModule, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromAccountToModule indicates an expected call of SendCoinsFromAccountToModule.
func (mr *MockBankKeeperMockRecorder) SendCoinsFromAccountToModule(ctx, senderAddr, recipientModule, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromAccountToModule", reflect.TypeOf((*MockBankKeeper)(nil).SendCoinsFromAccountToModule), ctx, senderAddr, recipientModule, amt)
}

// SendCoinsFromModuleToAccount mocks base method.
func (m *MockBankKeeper) SendCoinsFromModuleToAccount(ctx types.Context, senderModule string, recipientAddr types.AccAddress, amt types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromModuleToAccount", ctx, senderModule, recipientAddr, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromModuleToAccount indicates an expected call of SendCoinsFromModuleToAccount.
func (mr *MockBankKeeperMockRecorder) SendCoinsFromModuleToAccount(ctx, senderModule, recipientAddr, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromModuleToAccount", reflect.TypeOf((*MockBankKeeper)(nil).SendCoinsFromModuleToAccount), ctx, senderModule, recipientAddr, amt)
}

// SendCoinsFromModuleToModule mocks base method.
func (m *MockBankKeeper) SendCoinsFromModuleToModule(ctx types.Context, senderPool, recipientPool string, amt types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromModuleToModule", ctx, senderPool, recipientPool, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromModuleToModule indicates an expected call of SendCoinsFromModuleToModule.
func (mr *MockBankKeeperMockRecorder) SendCoinsFromModuleToModule(ctx, senderPool, recipientPool, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromModuleToModule", reflect.TypeOf((*MockBankKeeper)(nil).SendCoinsFromModuleToModule), ctx, senderPool, recipientPool, amt)
}

// SpendableCoins mocks base method.
func (m *MockBankKeeper) SpendableCoins(ctx types.Context, addr types.AccAddress) types.Coins {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SpendableCoins", ctx, addr)
	ret0, _ := ret[0].(types.Coins)
	return ret0
}

// SpendableCoins indicates an expected call of SpendableCoins.
func (mr *MockBankKeeperMockRecorder) SpendableCoins(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpendableCoins", reflect.TypeOf((*MockBankKeeper)(nil).SpendableCoins), ctx, addr)
}

// UndelegateCoinsFromModuleToAccount mocks base method.
func (m *MockBankKeeper) UndelegateCoinsFromModuleToAccount(ctx types.Context, senderModule string, recipientAddr types.AccAddress, amt types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UndelegateCoinsFromModuleToAccount", ctx, senderModule, recipientAddr, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// UndelegateCoinsFromModuleToAccount indicates an expected call of UndelegateCoinsFromModuleToAccount.
func (mr *MockBankKeeperMockRecorder) UndelegateCoinsFromModuleToAccount(ctx, senderModule, recipientAddr, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UndelegateCoinsFromModuleToAccount", reflect.TypeOf((*MockBankKeeper)(nil).UndelegateCoinsFromModuleToAccount), ctx, senderModule, recipientAddr, amt)
}

// MockValidatorSet is a mock of ValidatorSet interface.
type MockValidatorSet struct {
	ctrl     *gomock.Controller
	recorder *MockValidatorSetMockRecorder
}

// MockValidatorSetMockRecorder is the mock recorder for MockValidatorSet.
type MockValidatorSetMockRecorder struct {
	mock *MockValidatorSet
}

// NewMockValidatorSet creates a new mock instance.
func NewMockValidatorSet(ctrl *gomock.Controller) *MockValidatorSet {
	mock := &MockValidatorSet{ctrl: ctrl}
	mock.recorder = &MockValidatorSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockValidatorSet) EXPECT() *MockValidatorSetMockRecorder {
	return m.recorder
}

// Delegation mocks base method.
func (m *MockValidatorSet) Delegation(arg0 types.Context, arg1 types.AccAddress, arg2 types.ValAddress) types1.DelegationI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delegation", arg0, arg1, arg2)
	ret0, _ := ret[0].(types1.DelegationI)
	return ret0
}

// Delegation indicates an expected call of Delegation.
func (mr *MockValidatorSetMockRecorder) Delegation(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delegation", reflect.TypeOf((*MockValidatorSet)(nil).Delegation), arg0, arg1, arg2)
}

// IterateBondedValidatorsByPower mocks base method.
func (m *MockValidatorSet) IterateBondedValidatorsByPower(arg0 types.Context, arg1 func(int64, types1.ValidatorI) bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IterateBondedValidatorsByPower", arg0, arg1)
}

// IterateBondedValidatorsByPower indicates an expected call of IterateBondedValidatorsByPower.
func (mr *MockValidatorSetMockRecorder) IterateBondedValidatorsByPower(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IterateBondedValidatorsByPower", reflect.TypeOf((*MockValidatorSet)(nil).IterateBondedValidatorsByPower), arg0, arg1)
}

// IterateLastValidators mocks base method.
func (m *MockValidatorSet) IterateLastValidators(arg0 types.Context, arg1 func(int64, types1.ValidatorI) bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IterateLastValidators", arg0, arg1)
}

// IterateLastValidators indicates an expected call of IterateLastValidators.
func (mr *MockValidatorSetMockRecorder) IterateLastValidators(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IterateLastValidators", reflect.TypeOf((*MockValidatorSet)(nil).IterateLastValidators), arg0, arg1)
}

// IterateValidators mocks base method.
func (m *MockValidatorSet) IterateValidators(arg0 types.Context, arg1 func(int64, types1.ValidatorI) bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IterateValidators", arg0, arg1)
}

// IterateValidators indicates an expected call of IterateValidators.
func (mr *MockValidatorSetMockRecorder) IterateValidators(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IterateValidators", reflect.TypeOf((*MockValidatorSet)(nil).IterateValidators), arg0, arg1)
}

// Jail mocks base method.
func (m *MockValidatorSet) Jail(arg0 types.Context, arg1 types.ConsAddress) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Jail", arg0, arg1)
}

// Jail indicates an expected call of Jail.
func (mr *MockValidatorSetMockRecorder) Jail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Jail", reflect.TypeOf((*MockValidatorSet)(nil).Jail), arg0, arg1)
}

// MaxValidators mocks base method.
func (m *MockValidatorSet) MaxValidators(arg0 types.Context) uint32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaxValidators", arg0)
	ret0, _ := ret[0].(uint32)
	return ret0
}

// MaxValidators indicates an expected call of MaxValidators.
func (mr *MockValidatorSetMockRecorder) MaxValidators(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaxValidators", reflect.TypeOf((*MockValidatorSet)(nil).MaxValidators), arg0)
}

// Slash mocks base method.
func (m *MockValidatorSet) Slash(arg0 types.Context, arg1 types.ConsAddress, arg2, arg3 int64, arg4 types.Dec) math.Int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Slash", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(math.Int)
	return ret0
}

// Slash indicates an expected call of Slash.
func (mr *MockValidatorSetMockRecorder) Slash(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Slash", reflect.TypeOf((*MockValidatorSet)(nil).Slash), arg0, arg1, arg2, arg3, arg4)
}

// SlashWithInfractionReason mocks base method.
func (m *MockValidatorSet) SlashWithInfractionReason(arg0 types.Context, arg1 types.ConsAddress, arg2, arg3 int64, arg4 types.Dec, arg5 types1.Infraction) math.Int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SlashWithInfractionReason", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(math.Int)
	return ret0
}

// SlashWithInfractionReason indicates an expected call of SlashWithInfractionReason.
func (mr *MockValidatorSetMockRecorder) SlashWithInfractionReason(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SlashWithInfractionReason", reflect.TypeOf((*MockValidatorSet)(nil).SlashWithInfractionReason), arg0, arg1, arg2, arg3, arg4, arg5)
}

// StakingTokenSupply mocks base method.
func (m *MockValidatorSet) StakingTokenSupply(arg0 types.Context) math.Int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StakingTokenSupply", arg0)
	ret0, _ := ret[0].(math.Int)
	return ret0
}

// StakingTokenSupply indicates an expected call of StakingTokenSupply.
func (mr *MockValidatorSetMockRecorder) StakingTokenSupply(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StakingTokenSupply", reflect.TypeOf((*MockValidatorSet)(nil).StakingTokenSupply), arg0)
}

// TotalBondedTokens mocks base method.
func (m *MockValidatorSet) TotalBondedTokens(arg0 types.Context) math.Int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TotalBondedTokens", arg0)
	ret0, _ := ret[0].(math.Int)
	return ret0
}

// TotalBondedTokens indicates an expected call of TotalBondedTokens.
func (mr *MockValidatorSetMockRecorder) TotalBondedTokens(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TotalBondedTokens", reflect.TypeOf((*MockValidatorSet)(nil).TotalBondedTokens), arg0)
}

// Unjail mocks base method.
func (m *MockValidatorSet) Unjail(arg0 types.Context, arg1 types.ConsAddress) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Unjail", arg0, arg1)
}

// Unjail indicates an expected call of Unjail.
func (mr *MockValidatorSetMockRecorder) Unjail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unjail", reflect.TypeOf((*MockValidatorSet)(nil).Unjail), arg0, arg1)
}

// Validator mocks base method.
func (m *MockValidatorSet) Validator(arg0 types.Context, arg1 types.ValAddress) types1.ValidatorI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validator", arg0, arg1)
	ret0, _ := ret[0].(types1.ValidatorI)
	return ret0
}

// Validator indicates an expected call of Validator.
func (mr *MockValidatorSetMockRecorder) Validator(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validator", reflect.TypeOf((*MockValidatorSet)(nil).Validator), arg0, arg1)
}

// ValidatorByConsAddr mocks base method.
func (m *MockValidatorSet) ValidatorByConsAddr(arg0 types.Context, arg1 types.ConsAddress) types1.ValidatorI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidatorByConsAddr", arg0, arg1)
	ret0, _ := ret[0].(types1.ValidatorI)
	return ret0
}

// ValidatorByConsAddr indicates an expected call of ValidatorByConsAddr.
func (mr *MockValidatorSetMockRecorder) ValidatorByConsAddr(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidatorByConsAddr", reflect.TypeOf((*MockValidatorSet)(nil).ValidatorByConsAddr), arg0, arg1)
}

// MockDelegationSet is a mock of DelegationSet interface.
type MockDelegationSet struct {
	ctrl     *gomock.Controller
	recorder *MockDelegationSetMockRecorder
}

// MockDelegationSetMockRecorder is the mock recorder for MockDelegationSet.
type MockDelegationSetMockRecorder struct {
	mock *MockDelegationSet
}

// NewMockDelegationSet creates a new mock instance.
func NewMockDelegationSet(ctrl *gomock.Controller) *MockDelegationSet {
	mock := &MockDelegationSet{ctrl: ctrl}
	mock.recorder = &MockDelegationSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDelegationSet) EXPECT() *MockDelegationSetMockRecorder {
	return m.recorder
}

// GetValidatorSet mocks base method.
func (m *MockDelegationSet) GetValidatorSet() types1.ValidatorSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidatorSet")
	ret0, _ := ret[0].(types1.ValidatorSet)
	return ret0
}

// GetValidatorSet indicates an expected call of GetValidatorSet.
func (mr *MockDelegationSetMockRecorder) GetValidatorSet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidatorSet", reflect.TypeOf((*MockDelegationSet)(nil).GetValidatorSet))
}

// IterateDelegations mocks base method.
func (m *MockDelegationSet) IterateDelegations(ctx types.Context, delegator types.AccAddress, fn func(int64, types1.DelegationI) bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IterateDelegations", ctx, delegator, fn)
}

// IterateDelegations indicates an expected call of IterateDelegations.
func (mr *MockDelegationSetMockRecorder) IterateDelegations(ctx, delegator, fn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IterateDelegations", reflect.TypeOf((*MockDelegationSet)(nil).IterateDelegations), ctx, delegator, fn)
}

// MockStakingHooks is a mock of StakingHooks interface.
type MockStakingHooks struct {
	ctrl     *gomock.Controller
	recorder *MockStakingHooksMockRecorder
}

// MockStakingHooksMockRecorder is the mock recorder for MockStakingHooks.
type MockStakingHooksMockRecorder struct {
	mock *MockStakingHooks
}

// NewMockStakingHooks creates a new mock instance.
func NewMockStakingHooks(ctrl *gomock.Controller) *MockStakingHooks {
	mock := &MockStakingHooks{ctrl: ctrl}
	mock.recorder = &MockStakingHooksMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStakingHooks) EXPECT() *MockStakingHooksMockRecorder {
	return m.recorder
}

// AfterDelegationModified mocks base method.
func (m *MockStakingHooks) AfterDelegationModified(ctx types.Context, delAddr types.AccAddress, valAddr types.ValAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AfterDelegationModified", ctx, delAddr, valAddr)
	ret0, _ := ret[0].(error)
	return ret0
}

// AfterDelegationModified indicates an expected call of AfterDelegationModified.
func (mr *MockStakingHooksMockRecorder) AfterDelegationModified(ctx, delAddr, valAddr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AfterDelegationModified", reflect.TypeOf((*MockStakingHooks)(nil).AfterDelegationModified), ctx, delAddr, valAddr)
}

// AfterUnbondingInitiated mocks base method.
func (m *MockStakingHooks) AfterUnbondingInitiated(ctx types.Context, id uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AfterUnbondingInitiated", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// AfterUnbondingInitiated indicates an expected call of AfterUnbondingInitiated.
func (mr *MockStakingHooksMockRecorder) AfterUnbondingInitiated(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AfterUnbondingInitiated", reflect.TypeOf((*MockStakingHooks)(nil).AfterUnbondingInitiated), ctx, id)
}

// AfterValidatorBeginUnbonding mocks base method.
func (m *MockStakingHooks) AfterValidatorBeginUnbonding(ctx types.Context, consAddr types.ConsAddress, valAddr types.ValAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AfterValidatorBeginUnbonding", ctx, consAddr, valAddr)
	ret0, _ := ret[0].(error)
	return ret0
}

// AfterValidatorBeginUnbonding indicates an expected call of AfterValidatorBeginUnbonding.
func (mr *MockStakingHooksMockRecorder) AfterValidatorBeginUnbonding(ctx, consAddr, valAddr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AfterValidatorBeginUnbonding", reflect.TypeOf((*MockStakingHooks)(nil).AfterValidatorBeginUnbonding), ctx, consAddr, valAddr)
}

// AfterValidatorBonded mocks base method.
func (m *MockStakingHooks) AfterValidatorBonded(ctx types.Context, consAddr types.ConsAddress, valAddr types.ValAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AfterValidatorBonded", ctx, consAddr, valAddr)
	ret0, _ := ret[0].(error)
	return ret0
}

// AfterValidatorBonded indicates an expected call of AfterValidatorBonded.
func (mr *MockStakingHooksMockRecorder) AfterValidatorBonded(ctx, consAddr, valAddr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AfterValidatorBonded", reflect.TypeOf((*MockStakingHooks)(nil).AfterValidatorBonded), ctx, consAddr, valAddr)
}

// AfterValidatorCreated mocks base method.
func (m *MockStakingHooks) AfterValidatorCreated(ctx types.Context, valAddr types.ValAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AfterValidatorCreated", ctx, valAddr)
	ret0, _ := ret[0].(error)
	return ret0
}

// AfterValidatorCreated indicates an expected call of AfterValidatorCreated.
func (mr *MockStakingHooksMockRecorder) AfterValidatorCreated(ctx, valAddr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AfterValidatorCreated", reflect.TypeOf((*MockStakingHooks)(nil).AfterValidatorCreated), ctx, valAddr)
}

// AfterValidatorRemoved mocks base method.
func (m *MockStakingHooks) AfterValidatorRemoved(ctx types.Context, consAddr types.ConsAddress, valAddr types.ValAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AfterValidatorRemoved", ctx, consAddr, valAddr)
	ret0, _ := ret[0].(error)
	return ret0
}

// AfterValidatorRemoved indicates an expected call of AfterValidatorRemoved.
func (mr *MockStakingHooksMockRecorder) AfterValidatorRemoved(ctx, consAddr, valAddr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AfterValidatorRemoved", reflect.TypeOf((*MockStakingHooks)(nil).AfterValidatorRemoved), ctx, consAddr, valAddr)
}

// BeforeDelegationCreated mocks base method.
func (m *MockStakingHooks) BeforeDelegationCreated(ctx types.Context, delAddr types.AccAddress, valAddr types.ValAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeforeDelegationCreated", ctx, delAddr, valAddr)
	ret0, _ := ret[0].(error)
	return ret0
}

// BeforeDelegationCreated indicates an expected call of BeforeDelegationCreated.
func (mr *MockStakingHooksMockRecorder) BeforeDelegationCreated(ctx, delAddr, valAddr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeforeDelegationCreated", reflect.TypeOf((*MockStakingHooks)(nil).BeforeDelegationCreated), ctx, delAddr, valAddr)
}

// BeforeDelegationRemoved mocks base method.
func (m *MockStakingHooks) BeforeDelegationRemoved(ctx types.Context, delAddr types.AccAddress, valAddr types.ValAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeforeDelegationRemoved", ctx, delAddr, valAddr)
	ret0, _ := ret[0].(error)
	return ret0
}

// BeforeDelegationRemoved indicates an expected call of BeforeDelegationRemoved.
func (mr *MockStakingHooksMockRecorder) BeforeDelegationRemoved(ctx, delAddr, valAddr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeforeDelegationRemoved", reflect.TypeOf((*MockStakingHooks)(nil).BeforeDelegationRemoved), ctx, delAddr, valAddr)
}

// BeforeDelegationSharesModified mocks base method.
func (m *MockStakingHooks) BeforeDelegationSharesModified(ctx types.Context, delAddr types.AccAddress, valAddr types.ValAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeforeDelegationSharesModified", ctx, delAddr, valAddr)
	ret0, _ := ret[0].(error)
	return ret0
}

// BeforeDelegationSharesModified indicates an expected call of BeforeDelegationSharesModified.
func (mr *MockStakingHooksMockRecorder) BeforeDelegationSharesModified(ctx, delAddr, valAddr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeforeDelegationSharesModified", reflect.TypeOf((*MockStakingHooks)(nil).BeforeDelegationSharesModified), ctx, delAddr, valAddr)
}

// BeforeTokenizeShareRecordRemoved mocks base method.
func (m *MockStakingHooks) BeforeTokenizeShareRecordRemoved(ctx types.Context, recordID uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeforeTokenizeShareRecordRemoved", ctx, recordID)
	ret0, _ := ret[0].(error)
	return ret0
}

// BeforeTokenizeShareRecordRemoved indicates an expected call of BeforeTokenizeShareRecordRemoved.
func (mr *MockStakingHooksMockRecorder) BeforeTokenizeShareRecordRemoved(ctx, recordID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeforeTokenizeShareRecordRemoved", reflect.TypeOf((*MockStakingHooks)(nil).BeforeTokenizeShareRecordRemoved), ctx, recordID)
}

// BeforeValidatorModified mocks base method.
func (m *MockStakingHooks) BeforeValidatorModified(ctx types.Context, valAddr types.ValAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeforeValidatorModified", ctx, valAddr)
	ret0, _ := ret[0].(error)
	return ret0
}

// BeforeValidatorModified indicates an expected call of BeforeValidatorModified.
func (mr *MockStakingHooksMockRecorder) BeforeValidatorModified(ctx, valAddr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeforeValidatorModified", reflect.TypeOf((*MockStakingHooks)(nil).BeforeValidatorModified), ctx, valAddr)
}

// BeforeValidatorSlashed mocks base method.
func (m *MockStakingHooks) BeforeValidatorSlashed(ctx types.Context, valAddr types.ValAddress, fraction types.Dec) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeforeValidatorSlashed", ctx, valAddr, fraction)
	ret0, _ := ret[0].(error)
	return ret0
}

// BeforeValidatorSlashed indicates an expected call of BeforeValidatorSlashed.
func (mr *MockStakingHooksMockRecorder) BeforeValidatorSlashed(ctx, valAddr, fraction interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeforeValidatorSlashed", reflect.TypeOf((*MockStakingHooks)(nil).BeforeValidatorSlashed), ctx, valAddr, fraction)
}
