package docker

// ImageConfig contains all images and their respective tags
// needed for running e2e tests.
type ImageConfig struct {
	InitRepository string
	InitTag        string

	C4eRepository string
	C4eTag        string

	RelayerRepository string
	RelayerTag        string
}

const (
	// Local osmosis repo/version.
	// It is used when skipping upgrade by setting OSMOSIS_E2E_SKIP_UPGRADE to true).
	// This image should be pre-built with `make docker-build-debug` either in CI or locally.
	LocalC4eRepository = "osmosis"
	LocalC4eTag        = "debug"
	// Local osmosis repo/version for osmosis initialization.
	// It is used when skipping upgrade by setting OSMOSIS_E2E_SKIP_UPGRADE to true).
	// This image should be pre-built with `make docker-build-e2e-chain-init` either in CI or locally.
	localInitRepository = "osmosis-e2e-chain-init"
	localInitTag        = "debug"
	// Pre-upgrade osmosis repo/tag to pull.
	// It should be uploaded to Docker Hub. OSMOSIS_E2E_SKIP_UPGRADE should be unset
	// for this functionality to be used.
	previousVersionC4eRepository = "osmolabs/osmosis-dev"
	previousVersionC4eTag        = "v8.0.0-2-debug"
	// Pre-upgrade repo/tag for osmosis initialization (this should be one version below upgradeVersion)
	previousVersionInitRepository = "osmolabs/osmosis-init"
	previousVersionInitTag        = "v8.0.0-4-osmo"
	// Hermes repo/version for relayer
	relayerRepository = "osmolabs/hermes"
	relayerTag        = "0.13.0"
)

// Returns ImageConfig needed for running e2e test.
// If isUpgrade is true, returns images for running the upgrade
// Otherwise, returns images for running non-upgrade e2e tests.
func NewImageConfig(isUpgrade bool) *ImageConfig {
	config := &ImageConfig{
		RelayerRepository: relayerRepository,
		RelayerTag:        relayerTag,
	}

	if isUpgrade {
		config.InitRepository = previousVersionInitRepository
		config.InitTag = previousVersionInitTag

		config.C4eRepository = previousVersionC4eRepository
		config.C4eTag = previousVersionC4eTag
	} else {
		config.InitRepository = localInitRepository
		config.InitTag = localInitTag

		config.C4eRepository = LocalC4eRepository
		config.C4eTag = LocalC4eTag
	}

	return config
}
