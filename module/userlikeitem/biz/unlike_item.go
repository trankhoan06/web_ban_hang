package biz

import (
	"context"
	"main.go/module/userlikeitem/model"
)

type UserUnlikestorage interface {
	GetUserLike(ctx context.Context, itemId int, userId int) (*model.LikeItem, error)
	DeletedUserLike(ctx context.Context, itemId int, userId int) error
}

type UserUnlikeBiz struct {
	store UserUnlikestorage
}

func NewUserUnlikedBiz(store UserUnlikestorage) *UserUnlikeBiz {
	return &UserUnlikeBiz{store: store}
}
func (biz *UserUnlikeBiz) NewDeleteUserUnlike(ctx context.Context, itemId int, userId int) error {
	if _, err := biz.store.GetUserLike(ctx, itemId, userId); err != nil {
		return err
	}
	if err := biz.store.DeletedUserLike(ctx, itemId, userId); err != nil {
		return err
	}
	return nil
}
