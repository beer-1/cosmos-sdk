package bank_test

import (
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"github.com/cosmos/cosmos-sdk/x/bank/testutil"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/stretchr/testify/require"

	"cosmossdk.io/depinject"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/testutil/configurator"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/x/auth"
	_ "github.com/cosmos/cosmos-sdk/x/auth/tx/config"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	_ "github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
	_ "github.com/cosmos/cosmos-sdk/x/consensus"
	_ "github.com/cosmos/cosmos-sdk/x/distribution"
	_ "github.com/cosmos/cosmos-sdk/x/gov"
	_ "github.com/cosmos/cosmos-sdk/x/params"
	_ "github.com/cosmos/cosmos-sdk/x/staking"
)

func Test20685(t *testing.T) {
	wg := sync.WaitGroup{}

	numAccounts := 100
	s, _, _, txsBytes := setupSUT(t, numAccounts)

	for i := 0; i < numAccounts/2-1; i++ {
		wg.Add(1)
		go func(i int) {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			_, _, err := s.App.Simulate(txsBytes[i])

			wg.Done()
			require.NoError(t, err)
		}(i)
	}

	execUnit := 10
	for i := numAccounts/2 - 1; i < numAccounts-1; i += execUnit {
		_, err := s.App.FinalizeBlock(&abci.RequestFinalizeBlock{Txs: txsBytes[i : i+execUnit], Height: s.App.LastBlockHeight() + 1})
		require.NoError(t, err)

		_, err = s.App.Commit()
		require.NoError(t, err)
	}

	wg.Wait()
}

func Test20685_WithCheckTxSimulate(t *testing.T) {
	wg := sync.WaitGroup{}

	numAccounts := 100
	s, _, _, txsBytes := setupSUT(t, numAccounts)

	for i := 0; i < numAccounts/2-1; i++ {
		wg.Add(1)
		go func(i int) {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			_, _, err := s.App.Simulate(txsBytes[i])

			wg.Done()
			require.NoError(t, err)
		}(i)
	}

	for i := numAccounts/2 - 1; i < numAccounts; i++ {
		_, err := s.App.CheckTx(&abci.RequestCheckTx{Tx: txsBytes[i], Type: abci.CheckTxType_New})
		require.NoError(t, err)
	}

	wg.Wait()
}

func setupSUT(t *testing.T, numAccounts int) (*suite, *baseapp.BaseApp, sdk.Context, [][]byte) {
	senderKeys := make([]cryptotypes.PrivKey, numAccounts)
	for i := 0; i < numAccounts; i++ {
		senderKeys[i] = secp256k1.GenPrivKey()
	}

	ga := make([]authtypes.GenesisAccount, len(senderKeys))
	for i, senderKey := range senderKeys {
		senderAddr := sdk.AccAddress(senderKey.PubKey().Address())
		ga[i] = &authtypes.BaseAccount{Address: senderAddr.String()}
	}
	s := createTestSuite(t, ga)
	ctx := s.App.NewContext(false)
	// fund
	for _, senderKey := range senderKeys {
		senderAddr := sdk.AccAddress(senderKey.PubKey().Address())
		require.NoError(t, testutil.FundAccount(ctx, s.BankKeeper, senderAddr, sdk.NewCoins(sdk.NewInt64Coin("foocoin", 100_000_000))))
	}
	nextBlock(t, s.App)

	txsBz := make([][]byte, len(senderKeys))
	ctx = s.App.NewContext(true)
	for i, senderKey := range senderKeys {
		senderAddr := sdk.AccAddress(senderKey.PubKey().Address())
		senderAccount := s.AccountKeeper.GetAccount(ctx, senderAddr)
		require.NotNil(t, senderAccount)

		accNum, seq := senderAccount.GetAccountNumber(), senderAccount.GetSequence()

		// encode an example msg
		receiverAddr := make([]byte, 32)
		_, err := rand.Read(receiverAddr)
		require.NoError(t, err)

		sendMsg := types.NewMsgSend(senderAddr, receiverAddr, sdk.Coins{sdk.NewInt64Coin("foocoin", 100)})
		txBytes := encodeAndSign(t, s.TxConfig, []sdk.Msg{sendMsg}, "", []uint64{accNum}, []uint64{seq}, senderKey)
		txsBz[i] = txBytes
	}

	return &s, nil, ctx, txsBz
}

func createTestSuite(t *testing.T, genesisAccounts []authtypes.GenesisAccount) suite {
	res := suite{}

	var genAccounts []simtestutil.GenesisAccount
	for _, acc := range genesisAccounts {
		genAccounts = append(genAccounts, simtestutil.GenesisAccount{GenesisAccount: acc})
	}

	startupCfg := simtestutil.DefaultStartUpConfig()
	startupCfg.GenesisAccounts = genAccounts

	app, err := simtestutil.SetupWithConfiguration(
		depinject.Configs(
			configurator.NewAppConfig(
				configurator.ParamsModule(),
				configurator.AuthModule(),
				configurator.StakingModule(),
				configurator.TxModule(),
				configurator.ConsensusModule(),
				configurator.BankModule(),
				configurator.GovModule(),
				configurator.DistributionModule(),
			),
			depinject.Supply(log.NewNopLogger()),
		),
		startupCfg, &res.BankKeeper, &res.AccountKeeper, &res.TxConfig)

	res.App = app

	require.NoError(t, err)
	return res
}

func nextBlock(t *testing.T, app *runtime.App) {
	t.Helper()
	_, err := app.FinalizeBlock(&abci.RequestFinalizeBlock{Height: app.LastBlockHeight() + 1})
	require.NoError(t, err)
	_, err = app.Commit()
	require.NoError(t, err)
}

func encodeAndSign(
	t *testing.T,
	txConfig client.TxConfig,
	msgs []sdk.Msg,
	chainID string,
	accNums, accSeqs []uint64,
	priv ...cryptotypes.PrivKey,
) []byte {
	t.Helper()
	tx, err := simtestutil.GenSignedMockTx(
		rand.New(rand.NewSource(time.Now().UnixNano())),
		txConfig,
		msgs,
		sdk.Coins{sdk.NewInt64Coin(sdk.DefaultBondDenom, 0)},
		simtestutil.DefaultGenTxGas,
		chainID,
		accNums,
		accSeqs,
		priv...,
	)
	require.NoError(t, err)
	txBytes, err := txConfig.TxEncoder()(tx)
	require.Nil(t, err)
	return txBytes
}

type suite struct {
	TxConfig      client.TxConfig
	App           *runtime.App
	AccountKeeper authkeeper.AccountKeeper
	BankKeeper    bankkeeper.Keeper
}
