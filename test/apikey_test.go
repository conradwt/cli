package test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/confluentinc/cli/internal/pkg/config"
	"github.com/confluentinc/cli/internal/pkg/config/load"
	v3 "github.com/confluentinc/cli/internal/pkg/config/v3"
	"github.com/confluentinc/cli/internal/pkg/log"
)

func (s *CLITestSuite) TestAPIKeyCommands() {
	kafkaAPIURL := serveKafkaAPI(s.T()).URL
	loginURL := serve(s.T(), kafkaAPIURL).URL

	// TODO: add --config flag to all commands or ENVVAR instead of using standard config file location
	tests := []CLITest{
		{args: "api-key create --resource lkc-bob", login: "default", fixture: "apikey1.golden"}, // MYKEY3
		{args: "api-key list --resource lkc-bob", fixture: "apikey2.golden"},
		{args: "api-key list --resource lkc-abc", fixture: "apikey3.golden"},
		{args: "api-key update MYKEY1 --description first-key", fixture: "apikey40.golden"},
		{args: "api-key list --resource lkc-bob", fixture: "apikey41.golden"},

		// list json and yaml output
		{args: "api-key list", fixture: "apikey28.golden"},
		{args: "api-key list -o json", fixture: "apikey29.golden"},
		{args: "api-key list -o yaml", fixture: "apikey30.golden"},

		// create api key for kafka cluster
		{args: "api-key list --resource lkc-cool1", fixture: "apikey4.golden"},
		{args: "api-key create --description my-cool-app --resource lkc-cool1", fixture: "apikey5.golden"}, // MYKEY4
		{args: "api-key list --resource lkc-cool1", fixture: "apikey6.golden"},

		// create api key for other kafka cluster
		{args: "api-key create --description my-other-app --resource lkc-other1", fixture: "apikey7.golden"}, // MYKEY5
		{args: "api-key list --resource lkc-cool1", fixture: "apikey6.golden"},
		{args: "api-key list --resource lkc-other1", fixture: "apikey8.golden"},

		// create api key for ksql cluster
		{args: "api-key create --description my-ksql-app --resource lksqlc-ksql1", fixture: "apikey9.golden"}, // MYKEY6
		{args: "api-key list --resource lkc-cool1", fixture: "apikey6.golden"},
		{args: "api-key list --resource lksqlc-ksql1", fixture: "apikey10.golden"},

		// create api key for schema registry cluster
		{args: "api-key create --resource lsrc-1", fixture: "apikey20.golden"}, // MYKEY7
		{args: "api-key list --resource lsrc-1", fixture: "apikey21.golden"},

		// create cloud api key
		{args: "api-key create --resource cloud", fixture: "apikey34.golden"}, // MYKEY8
		{args: "api-key list --resource cloud", fixture: "apikey35.golden"},

		// use an api key for kafka cluster
		{args: "api-key use MYKEY4 --resource lkc-cool1", fixture: "apikey45.golden"},
		{args: "api-key list --resource lkc-cool1", fixture: "apikey11.golden"},

		// use an api key for other kafka cluster
		{args: "api-key use MYKEY5 --resource lkc-other1", fixture: "apikey46.golden"},
		{args: "api-key list --resource lkc-cool1", fixture: "apikey11.golden"},
		{args: "api-key list --resource lkc-other1", fixture: "apikey12.golden"},

		// delete api key that is in use
		{args: "api-key delete MYKEY5", fixture: "apikey42.golden"},
		{args: "api-key list --resource lkc-other1", fixture: "apikey43.golden"},

		// store an api-key for kafka cluster
		{args: "api-key store UIAPIKEY100 @test/fixtures/input/UIAPISECRET100.txt --resource lkc-cool1", fixture: "apikey47.golden"},
		{args: "api-key list --resource lkc-cool1", fixture: "apikey11.golden"},

		// store an api-key for other kafka cluster
		{args: "api-key store UIAPIKEY101 @test/fixtures/input/UIAPISECRET101.txt --resource lkc-other1", fixture: "apikey48.golden"},
		{args: "api-key list --resource lkc-cool1", fixture: "apikey11.golden"},
		{args: "api-key list --resource lkc-other1", fixture: "apikey44.golden"},

		// store exists already error
		{args: "api-key store UIAPIKEY101 @test/fixtures/input/UIAPISECRET101.txt --resource lkc-other1", fixture: "apikey-override-error.golden", wantErrCode: 1},

		// store an api-key for ksql cluster (not yet supported)
		//{args: "api-key store UIAPIKEY103 UIAPISECRET103 --resource lksqlc-ksql1", fixture: "empty.golden"},
		//{args: "api-key list --resource lksqlc-ksql1", fixture: "apikey10.golden"},
		// TODO: change test back once api-key store and use command allows for non kafka clusters
		{args: "api-key store UIAPIKEY103 UIAPISECRET103 --resource lksqlc-ksql1", fixture: "apikey36.golden", wantErrCode: 1},
		{args: "api-key use UIAPIKEY103 --resource lksqlc-ksql1", fixture: "apikey36.golden", wantErrCode: 1},

		// list all api-keys
		{args: "api-key list", fixture: "apikey22.golden"},

		// list api-keys belonging to currently logged in user
		{args: "api-key list --current-user", fixture: "apikey23.golden"},

		// create api-key for a service account
		{args: "api-key create --resource lkc-cool1 --service-account 99", fixture: "apikey24.golden"},
		{args: "api-key list --current-user", fixture: "apikey23.golden"},
		{args: "api-key list", fixture: "apikey25.golden"},
		{args: "api-key list --service-account 99", fixture: "apikey26.golden"},
		{args: "api-key list --resource lkc-cool1", fixture: "apikey27.golden"},
		{args: "api-key list --resource lkc-cool1 --service-account 99", fixture: "apikey26.golden"},

		// create json yaml output
		{args: "api-key create --description human-output --resource lkc-other1", fixture: "apikey31.golden"},
		{args: "api-key create --description json-output --resource lkc-other1 -o json", fixture: "apikey32.golden"},
		{args: "api-key create --description yaml-output --resource lkc-other1 -o yaml", fixture: "apikey33.golden"},

		// store: error handling
		{name: "error if storing unknown api key", args: "api-key store UNKNOWN @test/fixtures/input/UIAPISECRET100.txt --resource lkc-cool1", fixture: "apikey15.golden"},
		{name: "error if storing api key with existing secret", args: "api-key store UIAPIKEY100 NEWSECRET --resource lkc-cool1", fixture: "apikey16.golden"},
		{name: "succeed if forced to overwrite existing secret", args: "api-key store -f UIAPIKEY100 NEWSECRET --resource lkc-cool1", fixture: "apikey49.golden",
			wantFunc: func(t *testing.T) {
				logger := log.New()
				cfg := v3.New(&config.Params{
					CLIName:    "ccloud",
					MetricSink: nil,
					Logger:     logger,
				})
				cfg, err := load.LoadAndMigrate(cfg)
				require.NoError(t, err)
				ctx := cfg.Context()
				require.NotNil(t, ctx)
				kcc := ctx.KafkaClusterContext.GetKafkaClusterConfig("lkc-cool1")
				pair := kcc.APIKeys["UIAPIKEY100"]
				require.NotNil(t, pair)
				require.Equal(t, "NEWSECRET", pair.Secret)

			}},

		// use: error handling
		{name: "error if using non-existent api-key", args: "api-key use UNKNOWN --resource lkc-cool1", fixture: "apikey17.golden"},
		{name: "error if using api-key for wrong cluster", args: "api-key use MYKEY2 --resource lkc-cool1", fixture: "apikey18.golden"},
		{name: "error if using api-key without existing secret", args: "api-key use UIAPIKEY103 --resource lkc-cool1", fixture: "apikey19.golden"},

		// more errors
		{args: "api-key use UIAPIKEY103", fixture: "apikey37.golden", wantErrCode: 1},
		{args: "api-key create", fixture: "apikey38.golden", wantErrCode: 1},
		{args: "api-key use UIAPIKEY103 --resource lkc-unknown", fixture: "apikey-resource-unknown-error.golden", wantErrCode: 1},
		{args: "api-key create --resource lkc-unknown", fixture: "apikey-resource-unknown-error.golden", wantErrCode: 1},
	}
	resetConfiguration(s.T(), "ccloud")
	for _, tt := range tests {
		if tt.name == "" {
			tt.name = tt.args
		}
		tt.workflow = true
		s.runCcloudTest(tt, loginURL)
	}
}
