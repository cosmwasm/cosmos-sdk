package upgrade

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

func NewParamChangeProposalHandler(k Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) sdk.Error {
		switch c := content.(type) {
		case govtypes.SoftwareUpgradeProposal:
			return handleSoftwareUpgradeProposal(ctx, k, c)

		default:
			errMsg := fmt.Sprintf("unrecognized software upgrade proposal content type: %T", c)
			return sdk.ErrUnknownRequest(errMsg)
		}
	}
}

func handleSoftwareUpgradeProposal(ctx sdk.Context, k Keeper, p govtypes.SoftwareUpgradeProposal) sdk.Error {
	return k.ScheduleUpgrade(ctx, p.Plan)
}
