package client

import (
	govclient "github.com/airchains-network/cosmos-sdk/x/gov/client"
	"github.com/airchains-network/cosmos-sdk/x/params/client/cli"
)

// ProposalHandler is the param change proposal handler.
var ProposalHandler = govclient.NewProposalHandler(cli.NewSubmitParamChangeProposalTxCmd)
