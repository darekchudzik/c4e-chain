package e2e

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/suite"

	configurer "github.com/chain4energy/c4e-chain/tests/e2e/configurer"
)

const (
	// Environment variable name to skip the upgrade tests
	skipUpgradeEnv = "C4E_E2E_SKIP_UPGRADE"
	// Environment variable name to skip the IBC tests
	skipIBCEnv  = "C4E_E2E_SKIP_IBC"
	debugLogEnv = "C4E_E2E_DEBUG_LOG"
	// Environment variable name to skip state sync testing
	skipStateSyncEnv = "C4E_E2E_SKIP_STATE_SYNC"
	// Environment variable name to determine if this upgrade is a fork
	forkHeightEnv = "C4E_E2E_FORK_HEIGHT"
	// Environment variable name to skip cleaning up Docker resources in teardown
	skipCleanupEnv = "C4E_E2E_SKIP_CLEANUP"
	// Environment variable name to determine what version we are upgrading to
	upgradeVersionEnv = "C4E_E2E_UPGRADE_VERSION"
	// Environment variable name to skip the params change tests
	skipParamsChange = "C4E_E2E_SKIP_PARAMS_CHANGE"
)

type IntegrationTestSuite struct {
	suite.Suite

	configurer       configurer.Configurer
	skipUpgrade      bool
	skipIBC          bool
	skipStateSync    bool
	skipParamsChange bool
	forkHeight       int
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

func (s *IntegrationTestSuite) SetupSuite() {
	os.Setenv(upgradeVersionEnv, "v1.0.1")
	os.Setenv(skipStateSyncEnv, "false")
	os.Setenv(skipUpgradeEnv, "true")
	os.Setenv(skipIBCEnv, "true")
	//os.Setenv(skipCleanupEnv, "true")

	s.T().Log("setting up e2e integration test suite...")
	var (
		err             error
		upgradeSettings configurer.UpgradeSettings
	)

	if str := os.Getenv(skipUpgradeEnv); len(str) > 0 {
		s.skipUpgrade, err = strconv.ParseBool(str)
		s.Require().NoError(err)
		if s.skipUpgrade {
			s.T().Log(fmt.Sprintf("%s was true, skipping upgrade tests", skipUpgradeEnv))
		}
	}
	upgradeSettings.IsEnabled = !s.skipUpgrade

	if str := os.Getenv(forkHeightEnv); len(str) > 0 {
		upgradeSettings.ForkHeight, err = strconv.ParseInt(str, 0, 64)
		s.Require().NoError(err)
		s.T().Log(fmt.Sprintf("fork upgrade is enabled, %s was set to height %d", forkHeightEnv, upgradeSettings.ForkHeight))
	}

	if str := os.Getenv(skipIBCEnv); len(str) > 0 {
		s.skipIBC, err = strconv.ParseBool(str)
		s.Require().NoError(err)
		if s.skipIBC {
			s.T().Log(fmt.Sprintf("%s was true, skipping IBC tests", skipIBCEnv))
		}
	}

	if str := os.Getenv(skipStateSyncEnv); len(str) > 0 {
		s.skipStateSync, err = strconv.ParseBool(str)
		s.Require().NoError(err)
		if s.skipStateSync {
			s.T().Log("skipping state sync testing")
		}
	}

	if str := os.Getenv(skipParamsChange); len(str) > 0 {
		s.skipParamsChange, err = strconv.ParseBool(str)
		s.Require().NoError(err)
		if s.skipStateSync {
			s.T().Log("skipping params change testing")
		}
	}

	isDebugLogEnabled := false
	if str := os.Getenv(debugLogEnv); len(str) > 0 {
		isDebugLogEnabled, err = strconv.ParseBool(str)
		s.Require().NoError(err)
		if isDebugLogEnabled {
			s.T().Log("debug logging is enabled. container logs from running cli commands will be printed to stdout")
		}
	}

	if str := os.Getenv(upgradeVersionEnv); len(str) > 0 {
		upgradeSettings.Version = str
		s.T().Log(fmt.Sprintf("upgrade version set to %s", upgradeSettings.Version))
	}

	s.configurer, err = configurer.New(s.T(), !s.skipIBC, isDebugLogEnabled, upgradeSettings)
	s.Require().NoError(err)

	err = s.configurer.ConfigureChains()
	s.Require().NoError(err)

	err = s.configurer.RunSetup()
	s.Require().NoError(err)
}

func (s *IntegrationTestSuite) TearDownSuite() {
	if str := os.Getenv(skipCleanupEnv); len(str) > 0 {
		skipCleanup, err := strconv.ParseBool(str)
		s.Require().NoError(err)

		if skipCleanup {
			s.T().Log("skipping e2e resources clean up...")
			return
		}
	}

	err := s.configurer.ClearResources()
	s.Require().NoError(err)
}