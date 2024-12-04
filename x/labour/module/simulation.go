package labour

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/merlingo/intravision/testutil/sample"
	laboursimulation "github.com/merlingo/intravision/x/labour/simulation"
	"github.com/merlingo/intravision/x/labour/types"
)

// avoid unused import issue
var (
	_ = laboursimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgWork = "op_weight_msg_work"
	// TODO: Determine the simulation weight value
	defaultWeightMsgWork int = 100

	opWeightMsgBeginTask = "op_weight_msg_begin_task"
	// TODO: Determine the simulation weight value
	defaultWeightMsgBeginTask int = 100

	opWeightMsgFinishTask = "op_weight_msg_finish_task"
	// TODO: Determine the simulation weight value
	defaultWeightMsgFinishTask int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	labourGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&labourGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgWork int
	simState.AppParams.GetOrGenerate(opWeightMsgWork, &weightMsgWork, nil,
		func(_ *rand.Rand) {
			weightMsgWork = defaultWeightMsgWork
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgWork,
		laboursimulation.SimulateMsgWork(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgBeginTask int
	simState.AppParams.GetOrGenerate(opWeightMsgBeginTask, &weightMsgBeginTask, nil,
		func(_ *rand.Rand) {
			weightMsgBeginTask = defaultWeightMsgBeginTask
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgBeginTask,
		laboursimulation.SimulateMsgBeginTask(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgFinishTask int
	simState.AppParams.GetOrGenerate(opWeightMsgFinishTask, &weightMsgFinishTask, nil,
		func(_ *rand.Rand) {
			weightMsgFinishTask = defaultWeightMsgFinishTask
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgFinishTask,
		laboursimulation.SimulateMsgFinishTask(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgWork,
			defaultWeightMsgWork,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				laboursimulation.SimulateMsgWork(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgBeginTask,
			defaultWeightMsgBeginTask,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				laboursimulation.SimulateMsgBeginTask(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgFinishTask,
			defaultWeightMsgFinishTask,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				laboursimulation.SimulateMsgFinishTask(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
