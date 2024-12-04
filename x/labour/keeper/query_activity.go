package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/merlingo/intravision/x/labour/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ActivityAll(ctx context.Context, req *types.QueryAllActivityRequest) (*types.QueryAllActivityResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var activitys []types.Activity

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	activityStore := prefix.NewStore(store, types.KeyPrefix(types.ActivityKey))

	pageRes, err := query.Paginate(activityStore, req.Pagination, func(key []byte, value []byte) error {
		var activity types.Activity
		if err := k.cdc.Unmarshal(value, &activity); err != nil {
			return err
		}
		if activity.TaskId == req.Taskid && activity.Worker == req.Worker {
			activitys = append(activitys, activity)
		}
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	//pageRes.Total = uint64(len(activitys))
	return &types.QueryAllActivityResponse{Activity: activitys, Pagination: pageRes}, nil
}

func (k Keeper) Activity(ctx context.Context, req *types.QueryGetActivityRequest) (*types.QueryGetActivityResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	activity, found := k.GetActivity(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetActivityResponse{Activity: activity}, nil
}

func (k Keeper) LastActivity(ctx context.Context, req *types.QueryLastActivityRequest) (*types.QueryLastActivityResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	activities := k.GetAllActivity(ctx)
	if len(activities) == 0 {
		return nil, sdkerrors.ErrKeyNotFound
	}

	var activity types.Activity
	activity = types.Activity{FinishWork: 0}
	for i := range activities {
		if activities[i].Worker == req.Worker && activity.FinishWork < activities[i].FinishWork {
			activity = activities[i]
		}
	}
	if activity.FinishWork == 0 {
		return nil, sdkerrors.ErrInvalidAddress
	}
	return &types.QueryLastActivityResponse{Activity: activity}, nil
}
