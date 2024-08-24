package biz

import (
	"context"
	"errors"
	"main.go/module/user/model"
)

type GetStorage interface {
	FindUser(ctx context.Context, cond map[string]interface{}) (*model.User, error)
}
type GetUserBiz struct {
	store GetStorage
}

func NewGetUserBiz(store GetStorage) *GetUserBiz {
	return &GetUserBiz{
		store: store,
	}
}
func (biz *GetUserBiz) NewGetUser(ctx context.Context, userId int) (*model.User, error) {
	user, err := biz.store.FindUser(ctx, map[string]interface{}{"id": userId})
	if err != nil {
		return nil, err
	}
	if user.Status == model.StatusUserInactive {
		return nil, errors.New("account has been deleted")
	}
	return user, nil
}
