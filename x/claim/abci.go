package claim

import (
	"github.com/rebuschain/rebus.core/v1/x/claim/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// EndBlocker called every block, process inflation, update validator set.
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {

	params, err := k.GetParams(ctx)
	if err != nil {
		panic(err)
	}

	// End Airdrop
	goneTime := ctx.BlockTime().Sub(params.AirdropStartTime)
	if goneTime > params.DurationUntilDecay {
		// airdrop time passed
		err := k.EndAirdrop(ctx)
		if err != nil {
			panic(err)
		}
	}
}