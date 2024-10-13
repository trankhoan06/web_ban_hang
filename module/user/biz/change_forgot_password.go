package biz

import (
	"context"
	"errors"
	"time"
)

func (biz *RegisterBiz) ChangeForgotPassword(ctx context.Context, token, password string, userId int) error {
	sendCode, err := biz.store.FindSendCode(ctx, map[string]interface{}{"token": token, "user_id": userId})
	if err != nil {
		return err
	}
	if sendCode.Verify == false {
		return errors.New("you don't verify code")
	}
	now := time.Now().Add(-7 * time.Hour)
	if now.After(sendCode.Expire) {
		return errors.New("expire")
	}
	user, errUser := biz.store.FindUser(ctx, map[string]interface{}{"id": userId})
	if errUser != nil {
		return errUser
	}
	password = biz.hash.Hash(user.Salt + password)
	if err := biz.store.ChangePassword(ctx, userId, password); err != nil {
		return err
	}
	expire := time.Now().Add(-7 * time.Hour)
	if err := biz.store.UpdateForgot(ctx, token, userId, map[string]interface{}{"expire_at	": expire}); err != nil {
		return err
	}
	return nil
}
