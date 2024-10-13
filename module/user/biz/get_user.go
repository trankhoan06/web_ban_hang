package biz

import (
	"context"
	"main.go/module/user/model"
)

func (biz *UserBiz) NewGetUser(ctx context.Context, userId int) (*model.User, error) {
	user, err := biz.store.FindUser(ctx, map[string]interface{}{"user_id": userId})
	if err != nil {
		return nil, err
	}
	return user, nil
}
