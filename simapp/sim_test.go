package simapp

import (
	"encoding/json"
	"flag"
	"math/rand"
	"os"
	"runtime/debug"
	"strings"
	"testing"

	abci "github.com/cometbft/cometbft/abci/types"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"cosmossdk.io/log"
	"cosmossdk.io/store"
	storetypes "cosmossdk.io/store/types"
	authzkeeper "cosmossdk.io/x/authz/keeper"
	"cosmossdk.io/x/feegrant"
	slashingtypes "cosmossdk.io/x/slashing/types"
	stakingtypes "cosmossdk.io/x/staking/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client/flags"
	codectestutil "github.com/cosmos/cosmos-sdk/codec/testutil"
	"github.com/cosmos/cosmos-sdk/server"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	simcli "github.com/cosmos/cosmos-sdk/x/simulation/client/cli"
	"github.com/cosmos/cosmos-sdk/x/simulation/helper"
	simtestutil "github.com/cosmos/cosmos-sdk/x/simulation/helper"
)

// SimAppChainID hardcoded chainID for simulation
const SimAppChainID = "simulation-app"

var FlagEnableStreamingValue bool

// Get flags every time the simulator is run
func init() {
	simcli.GetSimulatorFlags()
	flag.BoolVar(&FlagEnableStreamingValue, "EnableStreaming", false, "Enable streaming service")
}

// fauxMerkleModeOpt returns a BaseApp option to use a dbStoreAdapter instead of
// an IAVLStore for faster simulation speed.
func fauxMerkleModeOpt(bapp *baseapp.BaseApp) {
	bapp.SetFauxMerkleMode()
}

// interBlockCacheOpt returns a BaseApp option function that sets the persistent
// inter-block write-through cache.
func interBlockCacheOpt() func(*baseapp.BaseApp) {
	return baseapp.SetInterBlockCache(store.NewCommitKVStoreCacheManager())
}

func TestFullAppSimulation(t *testing.T) {
	config, db, _, app := setupSimulationApp(t, "skipping application simulation")
	// run randomized simulation
	_, simParams, simErr := simulation.SimulateFromSeed(
		t,
		os.Stdout,
		app.BaseApp,
		helper.AppStateFn(app.AppCodec(), app.AuthKeeper.AddressCodec(), app.StakingKeeper.ValidatorAddressCodec(), app.SimulationManager(), app.DefaultGenesis()),
		simtypes.RandomAccounts, // Replace with own random account function if using keys other than secp256k1
		helper.SimulationOperations(app, app.AppCodec(), config),
		BlockedAddresses(),
		config,
		app.AppCodec(),
		codectestutil.CodecOptions{}.GetAddressCodec(),
	)

	// export state and simParams before the simulation error is checked
	err := helper.CheckExportSimulation(app, config, simParams)
	require.NoError(t, err)
	require.NoError(t, simErr)

	if config.Commit {
		helper.PrintStats(db)
	}
}

func TestAppImportExport(t *testing.T) {
	config, db, appOptions, app := setupSimulationApp(t, "skipping application import/export simulation")

	// Run randomized simulation
	_, simParams, simErr := simulation.SimulateFromSeed(
		t,
		os.Stdout,
		app.BaseApp,
		helper.AppStateFn(app.AppCodec(), app.AuthKeeper.AddressCodec(), app.StakingKeeper.ValidatorAddressCodec(), app.SimulationManager(), app.DefaultGenesis()),
		simtypes.RandomAccounts, // Replace with own random account function if using keys other than secp256k1
		helper.SimulationOperations(app, app.AppCodec(), config),
		BlockedAddresses(),
		config,
		app.AppCodec(),
		codectestutil.CodecOptions{}.GetAddressCodec(),
	)

	// export state and simParams before the simulation error is checked
	err := helper.CheckExportSimulation(app, config, simParams)
	require.NoError(t, err)
	require.NoError(t, simErr)

	if config.Commit {
		helper.PrintStats(db)
	}

	t.Log("exporting genesis...\n")

	exported, err := app.ExportAppStateAndValidators(false, []string{}, []string{})
	require.NoError(t, err)

	t.Log("importing genesis...\n")

	newDB, newDir, _, _, err := helper.SetupSimulation(config, "leveldb-app-sim-2", "Simulation-2", simcli.FlagVerboseValue, simcli.FlagEnabledValue)
	require.NoError(t, err, "simulation setup failed")

	defer func() {
		require.NoError(t, newDB.Close())
		require.NoError(t, os.RemoveAll(newDir))
	}()

	appOptions[flags.FlagHome] = newDir // ensure a unique folder for the new app
	newApp := NewSimApp(log.NewNopLogger(), newDB, nil, true, appOptions, fauxMerkleModeOpt, baseapp.SetChainID(SimAppChainID))
	require.Equal(t, "SimApp", newApp.Name())

	var genesisState GenesisState
	err = json.Unmarshal(exported.AppState, &genesisState)
	require.NoError(t, err)

	ctxA := app.NewContextLegacy(true, cmtproto.Header{Height: app.LastBlockHeight()})
	ctxB := newApp.NewContextLegacy(true, cmtproto.Header{Height: app.LastBlockHeight()})
	_, err = newApp.ModuleManager.InitGenesis(ctxB, genesisState)
	if err != nil {
		if strings.Contains(err.Error(), "validator set is empty after InitGenesis") {
			t.Log("Skipping simulation as all validators have been unbonded")
			t.Logf("err: %s stacktrace: %s\n", err, string(debug.Stack()))
			return
		}
	}

	require.NoError(t, err)
	err = newApp.StoreConsensusParams(ctxB, exported.ConsensusParams)
	require.NoError(t, err)

	t.Log("comparing stores...")
	// skip certain prefixes
	skipPrefixes := map[string][][]byte{
		stakingtypes.StoreKey: {
			stakingtypes.UnbondingQueueKey, stakingtypes.RedelegationQueueKey, stakingtypes.ValidatorQueueKey,
			stakingtypes.HistoricalInfoKey, stakingtypes.UnbondingIDKey, stakingtypes.UnbondingIndexKey,
			stakingtypes.UnbondingTypeKey,
		},
		authzkeeper.StoreKey:   {authzkeeper.GrantQueuePrefix},
		feegrant.StoreKey:      {feegrant.FeeAllowanceQueueKeyPrefix},
		slashingtypes.StoreKey: {slashingtypes.ValidatorMissedBlockBitmapKeyPrefix},
	}

	storeKeys := app.GetStoreKeys()
	require.NotEmpty(t, storeKeys)

	for _, appKeyA := range storeKeys {
		// only compare kvstores
		if _, ok := appKeyA.(*storetypes.KVStoreKey); !ok {
			continue
		}

		keyName := appKeyA.Name()
		appKeyB := newApp.GetKey(keyName)

		storeA := ctxA.KVStore(appKeyA)
		storeB := ctxB.KVStore(appKeyB)

		failedKVAs, failedKVBs := helper.DiffKVStores(storeA, storeB, skipPrefixes[keyName])
		require.Equal(t, len(failedKVAs), len(failedKVBs), "unequal sets of key-values to compare %s, key stores %s and %s", keyName, appKeyA, appKeyB)

		t.Logf("compared %d different key/value pairs between %s and %s\n", len(failedKVAs), appKeyA, appKeyB)
		if !assert.Equal(t, 0, len(failedKVAs), helper.GetSimulationLog(keyName, app.SimulationManager().StoreDecoders, failedKVAs, failedKVBs)) {
			for _, v := range failedKVAs {
				t.Logf("store mismatch: %q\n", v)
			}
			t.FailNow()
		}
	}
}

func TestAppSimulationAfterImport(t *testing.T) {
	config, db, appOptions, app := setupSimulationApp(t, "skipping application simulation after import")

	// Run randomized simulation
	stopEarly, simParams, simErr := simulation.SimulateFromSeed(
		t,
		os.Stdout,
		app.BaseApp,
		helper.AppStateFn(app.AppCodec(), app.AuthKeeper.AddressCodec(), app.StakingKeeper.ValidatorAddressCodec(), app.SimulationManager(), app.DefaultGenesis()),
		simtypes.RandomAccounts, // Replace with own random account function if using keys other than secp256k1
		helper.SimulationOperations(app, app.AppCodec(), config),
		BlockedAddresses(),
		config,
		app.AppCodec(),
		codectestutil.CodecOptions{}.GetAddressCodec(),
	)
	require.NoError(t, simErr)

	// export state and simParams before the simulation error is checked
	err := helper.CheckExportSimulation(app, config, simParams)
	require.NoError(t, err)

	if config.Commit {
		helper.PrintStats(db)
	}

	if stopEarly {
		t.Log("can't export or import a zero-validator genesis, exiting test...")
		return
	}

	t.Logf("exporting genesis...\n")

	exported, err := app.ExportAppStateAndValidators(true, []string{}, []string{})
	require.NoError(t, err)

	t.Logf("importing genesis...\n")

	newDB, newDir, _, _, err := helper.SetupSimulation(config, "leveldb-app-sim-2", "Simulation-2", simcli.FlagVerboseValue, simcli.FlagEnabledValue)
	require.NoError(t, err, "simulation setup failed")

	defer func() {
		require.NoError(t, newDB.Close())
		require.NoError(t, os.RemoveAll(newDir))
	}()

	newApp := NewSimApp(log.NewNopLogger(), newDB, nil, true, appOptions, fauxMerkleModeOpt, baseapp.SetChainID(SimAppChainID))
	if !simcli.FlagSigverifyTxValue {
		newApp.SetNotSigverifyTx()
	}
	require.Equal(t, "SimApp", newApp.Name())

	_, err = newApp.InitChain(&abci.RequestInitChain{
		AppStateBytes: exported.AppState,
		ChainId:       SimAppChainID,
	})
	require.NoError(t, err)
	_, _, err = simulation.SimulateFromSeed(
		t,
		os.Stdout,
		newApp.BaseApp,
		helper.AppStateFn(app.AppCodec(), app.AuthKeeper.AddressCodec(), app.StakingKeeper.ValidatorAddressCodec(), app.SimulationManager(), app.DefaultGenesis()),
		simtypes.RandomAccounts, // Replace with own random account function if using keys other than secp256k1
		helper.SimulationOperations(newApp, newApp.AppCodec(), config),
		BlockedAddresses(),
		config,
		app.AppCodec(),
		codectestutil.CodecOptions{}.GetAddressCodec(),
	)
	require.NoError(t, err)
}

func setupSimulationApp(t *testing.T, msg string) (simtypes.Config, dbm.DB, simtestutil.AppOptionsMap, *SimApp) {
	t.Helper()
	config := simcli.NewConfigFromFlags()
	config.ChainID = SimAppChainID

	db, dir, logger, skip, err := helper.SetupSimulation(config, "leveldb-app-sim", "Simulation", simcli.FlagVerboseValue, simcli.FlagEnabledValue)
	if skip {
		t.Skip(msg)
	}
	require.NoError(t, err, "simulation setup failed")

	t.Cleanup(func() {
		require.NoError(t, db.Close())
		require.NoError(t, os.RemoveAll(dir))
	})

	appOptions := make(helper.AppOptionsMap, 0)
	appOptions[flags.FlagHome] = dir // ensure a unique folder
	appOptions[server.FlagInvCheckPeriod] = simcli.FlagPeriodValue

	app := NewSimApp(logger, db, nil, true, appOptions, fauxMerkleModeOpt, baseapp.SetChainID(SimAppChainID))
	if !simcli.FlagSigverifyTxValue {
		app.SetNotSigverifyTx()
	}
	require.Equal(t, "SimApp", app.Name())
	return config, db, appOptions, app
}

// TODO: Make another test for the fuzzer itself, which just has noOp txs
// and doesn't depend on the application.
func TestAppStateDeterminism(t *testing.T) {
	if !simcli.FlagEnabledValue {
		t.Skip("skipping application simulation")
	}

	config := simcli.NewConfigFromFlags()
	config.InitialBlockHeight = 1
	config.ExportParamsPath = ""
	config.OnOperation = false
	config.AllInvariants = false
	config.ChainID = SimAppChainID

	numSeeds := 3
	numTimesToRunPerSeed := 3 // This used to be set to 5, but we've temporarily reduced it to 3 for the sake of faster CI.
	appHashList := make([]json.RawMessage, numTimesToRunPerSeed)

	// We will be overriding the random seed and just run a single simulation on the provided seed value
	if config.Seed != simcli.DefaultSeedValue {
		numSeeds = 1
	}

	appOptions := viper.New()
	if FlagEnableStreamingValue {
		m := make(map[string]interface{})
		m["streaming.abci.keys"] = []string{"*"}
		m["streaming.abci.plugin"] = "abci_v1"
		m["streaming.abci.stop-node-on-err"] = true
		for key, value := range m {
			appOptions.SetDefault(key, value)
		}
	}
	appOptions.SetDefault(server.FlagInvCheckPeriod, simcli.FlagPeriodValue)
	if simcli.FlagVerboseValue {
		appOptions.SetDefault(flags.FlagLogLevel, "debug")
	}

	for i := 0; i < numSeeds; i++ {
		if config.Seed == simcli.DefaultSeedValue {
			config.Seed = rand.Int63()
		}

		t.Log("config.Seed: ", config.Seed)

		for j := 0; j < numTimesToRunPerSeed; j++ {
			var logger log.Logger
			if simcli.FlagVerboseValue {
				logger = log.NewTestLogger(t)
			} else {
				logger = log.NewNopLogger()
			}

			appOptions.SetDefault(flags.FlagHome, t.TempDir())
			db := dbm.NewMemDB()
			app := NewSimApp(logger, db, nil, true, appOptions, interBlockCacheOpt(), baseapp.SetChainID(SimAppChainID))
			if !simcli.FlagSigverifyTxValue {
				app.SetNotSigverifyTx()
			}

			t.Logf(
				"running non-determinism simulation; seed %d: %d/%d, attempt: %d/%d\n",
				config.Seed, i+1, numSeeds, j+1, numTimesToRunPerSeed,
			)

			_, _, err := simulation.SimulateFromSeed(
				t,
				os.Stdout,
				app.BaseApp,
				helper.AppStateFn(app.AppCodec(), app.AuthKeeper.AddressCodec(), app.StakingKeeper.ValidatorAddressCodec(), app.SimulationManager(), app.DefaultGenesis()),
				simtypes.RandomAccounts, // Replace with own random account function if using keys other than secp256k1
				helper.SimulationOperations(app, app.AppCodec(), config),
				BlockedAddresses(),
				config,
				app.AppCodec(),
				codectestutil.CodecOptions{}.GetAddressCodec(),
			)
			require.NoError(t, err)

			if config.Commit {
				helper.PrintStats(db)
			}

			appHash := app.LastCommitID().Hash
			appHashList[j] = appHash

			if j != 0 {
				require.Equal(
					t, string(appHashList[0]), string(appHashList[j]),
					"non-determinism in seed %d: %d/%d, attempt: %d/%d\n", config.Seed, i+1, numSeeds, j+1, numTimesToRunPerSeed,
				)
			}
		}
	}
}
