package biz

import (
	"context"
	"errors"
	"time"
)

func (biz *UserBiz) NewVerifyCodeEmail(ctx context.Context, code int, token string, userId int) error {
	sendCode, err := biz.store.FindSendCode(ctx, map[string]interface{}{"token": token, "userId": userId})
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
	if err := biz.store.VerifyEmail(ctx, userId); err != nil {
		return err
	}
	return nil
}
