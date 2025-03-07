package parameter

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/elys-network/elys/x/parameter/keeper"
	"github.com/elys-network/elys/x/parameter/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the anteHandlerParam
	for _, elem := range genState.AnteHandlerParamList {
		k.SetAnteHandlerParam(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.AnteHandlerParamList = make([]types.AnteHandlerParam, 0)
	anteParam, bfound := k.GetAnteHandlerParam(ctx)
	if bfound {
		genesis.AnteHandlerParamList = append(genesis.AnteHandlerParamList, anteParam)
	}

	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
