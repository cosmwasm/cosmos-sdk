package upgrade

import "fmt"

// Software Upgrade Proposals
// TODO: We have to add fields for SUP specific arguments e.g. commit hash,
// upgrade date, etc.
type SoftwareUpgradeProposal struct {
	upgrade.Plan
	Description string `json:"description"`
}

func NewSoftwareUpgradeProposal(plan upgrade.Plan, description string) Content {
	return SoftwareUpgradeProposal{plan, description}
}

// Implements Proposal Interface
var _ Content = SoftwareUpgradeProposal{}

// nolint
func (sup SoftwareUpgradeProposal) GetTitle() string       { return sup.Name }
func (sup SoftwareUpgradeProposal) GetDescription() string { return sup.Description }
func (sup SoftwareUpgradeProposal) ProposalRoute() string  { return RouterKey }
func (sup SoftwareUpgradeProposal) ProposalType() string   { return ProposalTypeSoftwareUpgrade }
func (sup SoftwareUpgradeProposal) ValidateBasic() sdk.Error {
	return ValidateAbstract(DefaultCodespace, sup)
}

func (sup SoftwareUpgradeProposal) String() string {
	return fmt.Sprintf(`Software Upgrade Proposal:
  Title:       %s
  Description: %s
`, sup.Name, sup.Description)
}
