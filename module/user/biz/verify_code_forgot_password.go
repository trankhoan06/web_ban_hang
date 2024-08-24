package biz

import (
	"context"
	"fmt"
	"main.go/common"
	"main.go/module/user/model"
	"time"
)

type VerifyCodeForgotPasswordStorage interface {
	FindVerifyVerifyCode(ctx context.Context, email string) (*model.SendCode, error)
	UpdateTokenCodeVerifyEmail(ctx context.Context, id int, token string) error
}
type VerifyCodeForgotPasswordBiz struct {
	store VerifyCodeForgotPasswordStorage
}

func NewVerifyCodeForgotPasswordBiz(store VerifyCodeForgotPasswordStorage) *VerifyCodeForgotPasswordBiz {
	return &VerifyCodeForgotPasswordBiz{store: store}
}
func (biz *VerifyCodeForgotPasswordBiz) NewVerifyCode(ctx context.Context, code int, email string) error {
	verifyCode, err := biz.store.FindVerifyVerifyCode(ctx, email)
	if err != nil {
		return err
	}
	if code != verifyCode.Code {
		return fmt.Errorf("code not equal")
	}
	now := time.Now()
	if now.After(verifyCode.ExpireAt) {
		return fmt.Errorf("code has been expire")
	}
	token := common.GetSalt(50)
	_ = biz.store.UpdateTokenCodeVerifyEmail(ctx, verifyCode.Id, token)
	return nil
}
