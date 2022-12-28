package genesis

import (
	testcommon "github.com/chain4energy/c4e-chain/testutil/common"
	distributortypes "github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	mintertypes "github.com/chain4energy/c4e-chain/x/cfeminter/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

var (
	accountsAddresses, _ = testcommon.CreateAccounts(2, 0)
	DevelopmentFundAddr  = accountsAddresses[0]
	LpAccountAddr        = accountsAddresses[1]
)

var CfeDistributorParams = distributortypes.Params{
	SubDistributors: []distributortypes.SubDistributor{
		{
			Name: "tx_fee_distributor",
			Destinations: distributortypes.Destinations{
				PrimaryShare: distributortypes.Account{
					Id:   "c4e_distributor",
					Type: distributortypes.Main,
				},
				BurnShare: sdk.ZeroDec(),
			},
			Sources: []*distributortypes.Account{
				{
					Id:   "fee_collector",
					Type: distributortypes.ModuleAccount,
				},
			},
		},
		{
			Name: "inflation_and_fee_distributor",
			Destinations: distributortypes.Destinations{
				PrimaryShare: distributortypes.Account{
					Id:   distributortypes.ValidatorsRewardsCollector,
					Type: distributortypes.ModuleAccount,
				},
				BurnShare: sdk.ZeroDec(),
				Shares: []*distributortypes.DestinationShare{
					{
						Name:  "development_fund",
						Share: sdk.MustNewDecFromStr("0.05"),
						Destination: distributortypes.Account{
							Id:   DevelopmentFundAddr.String(),
							Type: distributortypes.BaseAccount,
						},
					},
					{
						Name:  "usage_incentives",
						Share: sdk.MustNewDecFromStr("0.35"),
						Destination: distributortypes.Account{
							Id:   "usage_incentives_collector",
							Type: distributortypes.InternalAccount,
						},
					},
				},
			},
			Sources: []*distributortypes.Account{
				{
					Id:   "c4e_distributor",
					Type: distributortypes.Main,
				},
			},
		},
		{
			Name: "usage_incentives_distributor",
			Destinations: distributortypes.Destinations{
				PrimaryShare: distributortypes.Account{
					Id:   LpAccountAddr.String(),
					Type: distributortypes.BaseAccount,
				},
				BurnShare: sdk.ZeroDec(),
				Shares: []*distributortypes.DestinationShare{
					{
						Name:  "green_energy_booster",
						Share: sdk.MustNewDecFromStr("0.34"),
						Destination: distributortypes.Account{
							Id:   "green_energy_booster_collector",
							Type: distributortypes.ModuleAccount,
						},
					},
					{
						Name:  "governance_booster",
						Share: sdk.MustNewDecFromStr("0.33"),
						Destination: distributortypes.Account{
							Id:   "governance_booster_collector",
							Type: distributortypes.ModuleAccount,
						},
					},
				},
			},
			Sources: []*distributortypes.Account{
				{
					Id:   "usage_incentives_collector",
					Type: distributortypes.InternalAccount,
				},
			},
		},
	},
}

var CfeMinterrParams = mintertypes.Params{
	MintDenom: "uc4e",
	MinterConfig: mintertypes.MinterConfig{
		StartTime: time.Now(),
		Minters: []*mintertypes.Minter{
			{
				SequenceId: 1,
				Type:       mintertypes.ExponentialStepMintingType,
				ExponentialStepMinting: &mintertypes.ExponentialStepMinting{
					StepDuration:     time.Hour * 24 * 365 * 4, // 4 years
					Amount:           sdk.NewInt(160000000000000),
					AmountMultiplier: sdk.MustNewDecFromStr("0.5"),
				},
			},
		},
	},
}
