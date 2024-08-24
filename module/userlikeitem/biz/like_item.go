package biz

import (
	"context"
	"main.go/module/userlikeitem/model"
)

type UserLikestorage interface {
	GetUserLike(ctx context.Context, itemId int, userId int) (*model.LikeItem, error)
	CreateUserLike(ctx context.Context, data *model.LikeItem) error
}

type UserLikeBiz struct {
	store UserLikestorage
}

func NewUserLikeBiz(store UserLikestorage) *UserLikeBiz {
	return &UserLikeBiz{store: store}
}
func (biz *UserLikeBiz) NewCreateUserLike(ctx context.Context, data *model.LikeItem) error {
	if _, err := biz.store.GetUserLike(ctx, data.ItemId, data.UserId); err == nil {
		return model.ErrUserLike
	}
	if err := biz.store.CreateUserLike(ctx, data); err != nil {
		return err
	}
	return nil
}
