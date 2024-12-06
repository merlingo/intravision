package labour

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/merlingo/intravision/api/intravision/labour"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "TaskAll",
					Use:       "list-task",
					Short:     "List all task",
				},
				{
					RpcMethod:      "Task",
					Use:            "show-task [id]",
					Short:          "Shows a task by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "ActivityAll",
					Use:            "list-activity [worker] [taskid]",
					Short:          "List all activity",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "worker"}, {ProtoField: "taskid"}},
				},
				{
					RpcMethod:      "Activity",
					Use:            "show-activity [id]",
					Short:          "Shows a activity by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "Work",
					Use:            "work [worker] [taskid] [begin-work] [finish-work] [working-time]",
					Short:          "Send a work tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "worker"}, {ProtoField: "taskid"}, {ProtoField: "beginWork"}, {ProtoField: "finishWork"}, {ProtoField: "workingTime"}},
				},
				{
					RpcMethod:      "BeginTask",
					Use:            "begin-task [taskid] [assigner] [begin-task] [deadline] [wager]",
					Short:          "Send a beginTask tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "taskid"}, {ProtoField: "assigner"}, {ProtoField: "beginTask"}, {ProtoField: "deadline"}, {ProtoField: "wager"}},
				},
				{
					RpcMethod:      "FinishTask",
					Use:            "finish-task [taskid] [finish-task]",
					Short:          "Send a finishTask tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "taskid"}, {ProtoField: "finishTask"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
