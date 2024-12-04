package gainsharing

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/merlingo/intravision/testutil/sample"
	gainsharingsimulation "github.com/merlingo/intravision/x/gainsharing/simulation"
	"github.com/merlingo/intravision/x/gainsharing/types"
)

// avoid unused import issue
var (
	_ = gainsharingsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgSetMechanism = "op_weight_msg_set_mechanism"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetMechanism int = 100

	opWeightMsgCalculatePerformance = "op_weight_msg_calculate_performance"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCalculatePerformance int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	gainsharingGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&gainsharingGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgSetMechanism int
	simState.AppParams.GetOrGenerate(opWeightMsgSetMechanism, &weightMsgSetMechanism, nil,
		func(_ *rand.Rand) {
			weightMsgSetMechanism = defaultWeightMsgSetMechanism
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetMechanism,
		gainsharingsimulation.SimulateMsgSetMechanism(am.accountKeeper,
			am.bankKeeper,
			am.bankEscrowKeeper,
			am.labourKeeper,
			am.keeper),
	))

	var weightMsgCalculatePerformance int
	simState.AppParams.GetOrGenerate(opWeightMsgCalculatePerformance, &weightMsgCalculatePerformance, nil,
		func(_ *rand.Rand) {
			weightMsgCalculatePerformance = defaultWeightMsgCalculatePerformance
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCalculatePerformance,
		gainsharingsimulation.SimulateMsgCalculatePerformance(am.accountKeeper,
			am.bankKeeper,
			am.bankEscrowKeeper,
			am.labourKeeper,
			am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgSetMechanism,
			defaultWeightMsgSetMechanism,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				gainsharingsimulation.SimulateMsgSetMechanism(am.accountKeeper,
					am.bankKeeper,
					am.bankEscrowKeeper,
					am.labourKeeper,
					am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCalculatePerformance,
			defaultWeightMsgCalculatePerformance,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				gainsharingsimulation.SimulateMsgCalculatePerformance(am.accountKeeper,
					am.bankKeeper,
					am.bankEscrowKeeper,
					am.labourKeeper,
					am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
