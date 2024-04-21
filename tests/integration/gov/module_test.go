package gov_test

import (
	"testing"

	"gotest.tools/v3/assert"

	"cosmossdk.io/depinject"
	"cosmossdk.io/log"
	_ "cosmossdk.io/x/accounts"
	authkeeper "cosmossdk.io/x/auth/keeper"
	authtypes "cosmossdk.io/x/auth/types"
	"cosmossdk.io/x/gov/types"
	_ "cosmossdk.io/x/mint"
	_ "cosmossdk.io/x/protocolpool"

	"github.com/cosmos/cosmos-sdk/testutil/configurator"
	"cosmossdk.io/simapp/sims"
)

func TestItCreatesModuleAccountOnInitBlock(t *testing.T) {
	var accountKeeper authkeeper.AccountKeeper
	app, err := sims.SetupAtGenesis(
		depinject.Configs(
			configurator.NewAppConfig(
				configurator.AccountsModule(),
				configurator.AuthModule(),
				configurator.StakingModule(),
				configurator.BankModule(),
				configurator.GovModule(),
				configurator.ConsensusModule(),
				configurator.ProtocolPoolModule(),
			),
			depinject.Supply(log.NewNopLogger()),
		),
		&accountKeeper,
	)
	assert.NilError(t, err)

	ctx := app.BaseApp.NewContext(false)
	acc := accountKeeper.GetAccount(ctx, authtypes.NewModuleAddress(types.ModuleName))
	assert.Assert(t, acc != nil)
}
