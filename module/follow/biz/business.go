package biz

import (
	"context"
	"main.go/module/follow/model"
	modelUser "main.go/module/user/model"
)

type FollowStorage interface {
	CreateFollowUser(ctx context.Context, user *model.CreateFollower) error
	Unfollow(ctx context.Context, user *model.CreateFollower) error
	GetAmountFollow(ctx context.Context, column string, userId int) (int, error)
	ListFollowUser(ctx context.Context, column string, userID int, moreKey ...string) (*[]model.Follower, error)
	GetUserFollow(ctx context.Context, user *model.CreateFollower) (*model.Follower, error)
}
type UserStorage interface {
	FindUser(ctx context.Context, cond map[string]interface{}) (*modelUser.User, error)
}
type FollowBiz struct {
	store FollowStorage
}
type FollowUserBiz struct {
	store  FollowStorage
	store1 UserStorage
}

func NewFollowBiz(store FollowStorage) *FollowBiz {
	return &FollowBiz{store: store}
}
func NewFollowUserBiz(store FollowStorage, store1 UserStorage) *FollowUserBiz {
	return &FollowUserBiz{store: store, store1: store1}
}
