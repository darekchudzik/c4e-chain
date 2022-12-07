package e2e

import (
	"encoding/json"
	"github.com/chain4energy/c4e-chain/tests/e2e/configurer/chain"
	"github.com/chain4energy/c4e-chain/tests/e2e/configurer/config"
	appparams "github.com/chain4energy/c4e-chain/tests/e2e/encoding/params"
	"github.com/chain4energy/c4e-chain/tests/e2e/initialization"
	cfedistributortypes "github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	cfemintertypes "github.com/chain4energy/c4e-chain/x/cfeminter/types"
	cfevestingtypes "github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	paramsutils "github.com/cosmos/cosmos-sdk/x/params/client/utils"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type ParamsSetupSuite struct {
	BaseSetupSuite
}

func TestParamsChangeSuite(t *testing.T) {
	suite.Run(t, new(ParamsSetupSuite))
}

func (s *ParamsSetupSuite) SetupSuite() {
	s.BaseSetupSuite.SetupSuite(true, false)
}

func (s *ParamsSetupSuite) TestMinterAndDistributorParamsChangeProposal() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()
	s.NoError(err)

	endTime := time.Now().Add(10 * time.Minute).UTC()
	newMinter := cfemintertypes.MinterConfig{
		StartTime: time.Now().UTC(),
		Minters: []*cfemintertypes.Minter{
			{
				SequenceId: 1,
				Type:       cfemintertypes.LINEAR_MINTING,
				LinearMinting: &cfemintertypes.LinearMinting{
					Amount: sdk.NewInt(100000),
				},
				EndTime: &endTime,
			},
			{
				SequenceId: 2,
				Type:       cfemintertypes.NO_MINTING,
			},
		},
	}
	newDenom := "newTestDenom"
	s.cfeminterParamsChange(node, chainA, newDenom, newMinter)
	s.validateTotalSupply(node, newDenom, true, 15)

	newSubDistributors := []cfedistributortypes.SubDistributor{
		{
			Name: "New subdistributor",
			Sources: []*cfedistributortypes.Account{
				{
					Id:   cfedistributortypes.DistributorMainAccount,
					Type: cfedistributortypes.MAIN,
				},
			},
			Destinations: cfedistributortypes.Destinations{
				PrimaryShare: cfedistributortypes.Account{
					Id:   cfedistributortypes.GreenEnergyBoosterCollector,
					Type: cfedistributortypes.MODULE_ACCOUNT,
				},
				BurnShare: sdk.ZeroDec(),
			},
		},
	}
	s.cfedistributorParamsChange(node, chainA, newSubDistributors)

	accAddress := authtypes.NewModuleAddress(cfedistributortypes.GreenEnergyBoosterCollector).String()
	s.validateBalanceOfAccount(node, newDenom, accAddress, true, 15)
}

func (s *ParamsSetupSuite) TestCfeminterParamsProposalNoMinting() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()
	s.NoError(err)

	newMinter := cfemintertypes.MinterConfig{
		StartTime: time.Now().UTC(),
		Minters: []*cfemintertypes.Minter{
			{
				SequenceId: 1,
				Type:       cfemintertypes.NO_MINTING,
			},
		},
	}

	newDenom := "newDenom"
	s.cfeminterParamsChange(node, chainA, newDenom, newMinter)
	s.validateTotalSupply(node, newDenom, false, 15)
}

func (s *ParamsSetupSuite) TestCfeVestingProposal() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()

	newVestingDenom, err := json.Marshal("newDenom")
	s.NoError(err)

	proposal := paramsutils.ParamChangeProposalJSON{
		Title:       "Cfevesting module params change",
		Description: "Change cfevesting params",
		Changes: paramsutils.ParamChangesJSON{
			paramsutils.ParamChangeJSON{
				Subspace: cfevestingtypes.ModuleName,
				Key:      string(cfevestingtypes.KeyDenom),
				Value:    newVestingDenom,
			},
		},
		Deposit: sdk.NewCoin(appparams.CoinDenom, config.MinDepositValue).String(),
	}

	proposalJSON, err := json.Marshal(proposal)
	s.NoError(err)
	node.SubmitParamChangeProposal(string(proposalJSON), initialization.ValidatorWalletName)
	chainA.LatestProposalNumber += 1

	for _, n := range chainA.NodeConfigs {
		n.VoteYesProposal(initialization.ValidatorWalletName, chainA.LatestProposalNumber)
	}

	s.Eventually(
		func() bool {
			return node.ValidateParams(newVestingDenom, cfevestingtypes.ModuleName, string(cfevestingtypes.KeyDenom))
		},
		time.Minute,
		time.Second*5,
		"C4e node failed to validate params",
	)
}

func (s *ParamsSetupSuite) cfedistributorParamsChange(node *chain.NodeConfig, chainConfig *chain.Config, newSubDistributors []cfedistributortypes.SubDistributor) {
	newSubDistributorsJSON, err := json.Marshal(newSubDistributors)
	s.NoError(err)

	proposal := paramsutils.ParamChangeProposalJSON{
		Title:       "CfeDistributor module params change",
		Description: "Change cfedistributor params",
		Changes: paramsutils.ParamChangesJSON{
			paramsutils.ParamChangeJSON{
				Subspace: cfedistributortypes.ModuleName,
				Key:      string(cfedistributortypes.KeySubDistributors),
				Value:    newSubDistributorsJSON,
			},
		},
		Deposit: sdk.NewCoin(appparams.CoinDenom, config.MinDepositValue).String(),
	}

	proposalJSON, err := json.Marshal(proposal)
	s.NoError(err)
	node.SubmitParamChangeProposal(string(proposalJSON), initialization.ValidatorWalletName)
	chainConfig.LatestProposalNumber += 1

	for _, n := range chainConfig.NodeConfigs {
		n.VoteYesProposal(initialization.ValidatorWalletName, chainConfig.LatestProposalNumber)
	}

	s.Eventually(
		func() bool {
			return node.ValidateParams(newSubDistributorsJSON, cfedistributortypes.ModuleName, string(cfedistributortypes.KeySubDistributors))
		},
		time.Minute,
		time.Second*5,
		"C4e node failed to validate params",
	)
}

func (s *ParamsSetupSuite) cfeminterParamsChange(node *chain.NodeConfig, chainConfig *chain.Config, newDenom string, newMinterConfig cfemintertypes.MinterConfig) {
	newMinterJSON, err := json.Marshal(newMinterConfig)
	s.NoError(err)
	newDenomJSON, err := json.Marshal(newDenom)
	s.NoError(err)

	proposal := paramsutils.ParamChangeProposalJSON{
		Title:       "Cfeminter module params change",
		Description: "Change cfeminter params",
		Changes: paramsutils.ParamChangesJSON{
			paramsutils.ParamChangeJSON{
				Subspace: cfemintertypes.ModuleName,
				Key:      string(cfemintertypes.KeyMinterConfig),
				Value:    newMinterJSON,
			},
			paramsutils.ParamChangeJSON{
				Subspace: cfemintertypes.ModuleName,
				Key:      string(cfemintertypes.KeyMintDenom),
				Value:    newDenomJSON,
			},
		},
		Deposit: sdk.NewCoin(appparams.CoinDenom, config.MinDepositValue).String(),
	}

	proposalJSON, err := json.Marshal(proposal)
	s.NoError(err)
	node.SubmitParamChangeProposal(string(proposalJSON), initialization.ValidatorWalletName)

	chainConfig.LatestProposalNumber += 1

	for _, n := range chainConfig.NodeConfigs {
		n.VoteYesProposal(initialization.ValidatorWalletName, chainConfig.LatestProposalNumber)
	}

	s.Eventually(
		func() bool {
			return node.ValidateParams(newDenomJSON, cfemintertypes.ModuleName, string(cfemintertypes.KeyMintDenom)) &&
				node.ValidateParams(newMinterJSON, cfemintertypes.ModuleName, string(cfemintertypes.KeyMinterConfig))
		},
		time.Minute,
		time.Second*5,
		"C4e node failed to retrieve params",
	)
}