package creata_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	creata "github.com/creatachain/gaia/v4/app"

	"github.com/creatachain/gaia/v4/app/helpers"
	"github.com/stretchr/testify/require"
	"github.com/creatachain/augusteum/libs/log"
	"github.com/creatachain/augusteum/libs/rand"
	dbm "github.com/creatachain/tm-db"

	"github.com/creatachain/creata-sdk/baseapp"
	"github.com/creatachain/creata-sdk/creataapp"
	"github.com/creatachain/creata-sdk/store"
	simulation2 "github.com/creatachain/creata-sdk/types/simulation"
	"github.com/creatachain/creata-sdk/x/simulation"
)

func init() {
	creataapp.GetSimulatorFlags()
}

// Profile with:
// /usr/local/go/bin/go test -benchmem -run=^$ github.com/creatachain/creata-sdk/CreataApp -bench ^BenchmarkFullAppSimulation$ -Commit=true -cpuprofile cpu.out
func BenchmarkFullAppSimulation(b *testing.B) {
	config, db, dir, logger, _, err := creataapp.SetupSimulation("goleveldb-app-sim", "Simulation")
	if err != nil {
		b.Fatalf("simulation setup failed: %s", err.Error())
	}

	defer func() {
		db.Close()
		err = os.RemoveAll(dir)
		if err != nil {
			b.Fatal(err)
		}
	}()

	app := creata.NewCreataApp(logger, db, nil, true, map[int64]bool{}, creata.DefaultNodeHome, creataapp.FlagPeriodValue, creata.MakeEncodingConfig(), creataapp.EmptyAppOptions{}, interBlockCacheOpt())

	// Run randomized simulation:w
	_, simParams, simErr := simulation.SimulateFromSeed(
		b,
		os.Stdout,
		app.BaseApp,
		creataapp.AppStateFn(app.AppCodec(), app.SimulationManager()),
		simulation2.RandomAccounts, // Replace with own random account function if using keys other than secp256k1
		creataapp.SimulationOperations(app, app.AppCodec(), config),
		app.ModuleAccountAddrs(),
		config,
		app.AppCodec(),
	)

	// export state and simParams before the simulation error is checked
	if err = creataapp.CheckExportSimulation(app, config, simParams); err != nil {
		b.Fatal(err)
	}

	if simErr != nil {
		b.Fatal(simErr)
	}

	if config.Commit {
		creataapp.PrintStats(db)
	}
}

// interBlockCacheOpt returns a BaseApp option function that sets the persistent
// inter-block write-through cache.
func interBlockCacheOpt() func(*baseapp.BaseApp) {
	return baseapp.SetInterBlockCache(store.NewCommitKVStoreCacheManager())
}

//// TODO: Make another test for the fuzzer itself, which just has noOp txs
//// and doesn't depend on the application.
func TestAppStateDeterminism(t *testing.T) {
	if !creataapp.FlagEnabledValue {
		t.Skip("skipping application simulation")
	}

	config := creataapp.NewConfigFromFlags()
	config.InitialBlockHeight = 1
	config.ExportParamsPath = ""
	config.OnOperation = false
	config.AllInvariants = false
	config.ChainID = helpers.CREATAAPPChainID

	numSeeds := 3
	numTimesToRunPerSeed := 5
	appHashList := make([]json.RawMessage, numTimesToRunPerSeed)

	for i := 0; i < numSeeds; i++ {
		config.Seed = rand.Int63()

		for j := 0; j < numTimesToRunPerSeed; j++ {
			var logger log.Logger
			if creataapp.FlagVerboseValue {
				logger = log.TestingLogger()
			} else {
				logger = log.NewNopLogger()
			}

			db := dbm.NewMemDB()
			app := creata.NewCreataApp(logger, db, nil, true, map[int64]bool{}, creata.DefaultNodeHome, creataapp.FlagPeriodValue, creata.MakeEncodingConfig(), creataapp.EmptyAppOptions{}, interBlockCacheOpt())

			fmt.Printf(
				"running non-determinism simulation; seed %d: %d/%d, attempt: %d/%d\n",
				config.Seed, i+1, numSeeds, j+1, numTimesToRunPerSeed,
			)

			_, _, err := simulation.SimulateFromSeed(
				t,
				os.Stdout,
				app.BaseApp,
				creataapp.AppStateFn(app.AppCodec(), app.SimulationManager()),
				simulation2.RandomAccounts, // Replace with own random account function if using keys other than secp256k1
				creataapp.SimulationOperations(app, app.AppCodec(), config),
				app.ModuleAccountAddrs(),
				config,
				app.AppCodec(),
			)
			require.NoError(t, err)

			if config.Commit {
				creataapp.PrintStats(db)
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
