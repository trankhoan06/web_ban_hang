package biz

import (
	"context"
	"errors"
	"main.go/module/user/model"
)

func (biz *UserBiz) NewUpdateInformationUser(ctx context.Context, data *model.UpdateUser) error {
	if data.Email == "" {
		return errors.New("require login")

	}
	if err := biz.store.UpdateUser(ctx, data); err != nil {
		return err
	}
	return nil
}
