package biz

import (
	"context"
	"errors"
	"main.go/module/follow/model"
)

func (biz *FollowUserBiz) CreateFollowUser(ctx context.Context, user *model.CreateFollower) error {
	userAccount, err := biz.store1.FindUser(ctx, map[string]interface{}{"id": user.UserId})
	if err == nil {
		if userAccount.Status == 0 {
			return errors.New("account has been deleted")
		}
	}
	if err != nil {
		return err
	}
	if _, err := biz.store.GetUserFollow(ctx, user); err == nil {
		return errors.New("user is already follow")
	}
	if err := biz.store.CreateFollowUser(ctx, user); err != nil {
		return err
	}
	return nil

}
