package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/merlingo/intravision/x/gainsharing/keeper"
	"github.com/merlingo/intravision/x/gainsharing/types"
)

func SimulateMsgSetMechanism(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	bek types.BankEscrowKeeper,
	lk types.LabourKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgSetMechanism{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the SetMechanism simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "SetMechanism simulation not implemented"), nil, nil
	}
}
