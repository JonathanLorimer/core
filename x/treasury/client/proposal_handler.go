package client

import (
	"github.com/terra-project/core/x/treasury/client/cli"
	"github.com/terra-project/core/x/treasury/client/rest"

	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
)

// param change proposal handler
var (
	TaxRateUpdateProposalHandler      = govclient.NewProposalHandler(cli.GetCmdSubmitTaxRateUpdateProposal, rest.TaxRateUpdateProposalRESTHandler)
	RewardWeightUpdateProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitRewardWeightUpdateProposal, rest.RewardWeightUpdateProposalRESTHandler)
)
