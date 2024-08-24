package biz

import (
	"context"
	"main.go/module/follow/model"
)

func (biz *FollowBiz) NewListFollowUser(ctx context.Context, userId int) (*[]model.Follower, error) {
	ListFollow, err := biz.store.ListFollowUser(ctx, "by_user_id=?", userId, "Owner")
	if err != nil {
		return nil, err
	}
	return ListFollow, nil

}
