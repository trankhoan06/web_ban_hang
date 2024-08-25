package biz

import (
	"context"
	"errors"
	"main.go/module/follow/model"
	modelUser "main.go/module/user/model"
)

func (biz *FollowUserBiz) CreateFollowUser(ctx context.Context, user *model.CreateFollower) (*modelUser.User, error) {
	userAccount, err := biz.store1.FindUser(ctx, map[string]interface{}{"id": user.UserId})
	if err == nil {
		if userAccount.Status == 0 {
			return nil, errors.New("account has been deleted")
		}
	}
	if err != nil {
		return nil, err
	}
	if _, err := biz.store.GetUserFollow(ctx, user); err == nil {
		return nil, errors.New("user is already follow")
	}
	if err := biz.store.CreateFollowUser(ctx, user); err != nil {
		return nil, err
	}
	return userAccount, nil

}
