package biz

import (
	"context"
	"errors"
)

func (biz *RegisterBiz) NewChangePassword(ctx context.Context, userId int, password, newPassword string) error {
	user, err := biz.store.FindUser(ctx, map[string]interface{}{"id": userId})
	if err != nil {
		return err
	}
	password = biz.hash.Hash(user.Salt + password)
	if password != user.Password {
		return errors.New("password error")
	}
	newPassword = biz.hash.Hash(user.Salt + newPassword)
	if err := biz.store.ChangePassword(ctx, userId, password); err != nil {
		return err
	}
	return nil
}
