package biz

import (
	"context"
	"errors"
	"main.go/module/follow/model"
)

func (biz *FollowBiz) Unfollow(ctx context.Context, user *model.CreateFollower) error {
	if _, err := biz.store.GetUserFollow(ctx, user); err != nil {
		return errors.New("user don't follow this user")
	}
	if err := biz.store.Unfollow(ctx, user); err != nil {
		return err
	}
	return nil
}
