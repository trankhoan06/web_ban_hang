package biz

import (
	"context"
	"errors"
	"time"
)

func (biz *UserBiz) NewVerifyCodeForgotPassword(ctx context.Context, code int, token string, userId int) error {
	sendCode, err := biz.store.FindSendCode(ctx, map[string]interface{}{"token": token, "user_id": userId})
	if err != nil {
		return err
	}
	now := time.Now().Add(-7 * time.Hour)

	if now.After(sendCode.Expire) {
		return errors.New("code is expire")
	}
	if code != sendCode.Code {
		return errors.New("code is wrong")
	}
	expire := time.Now().Add(-7 * time.Hour)
	expire = expire.Add(30 * time.Minute)
	if err := biz.store.UpdateForgot(ctx, token, userId, map[string]interface{}{"verify": true,
		"expire_at": expire}); err != nil {
		return err
	}
	return nil
}
