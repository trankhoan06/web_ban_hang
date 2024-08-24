package biz

import (
	"context"
	"fmt"
	"main.go/module/user/model"
)

type VerifyEmailAccountStorage interface {
	FindVerifyVerifyCode(ctx context.Context, email string) (*model.SendCode, error)
	VerifyEmailAccount(ctx context.Context, email string) error
}
type VerifyEmailAccountBiz struct {
	store VerifyEmailAccountStorage
}

func NewVerifyEmailAccountBiz(store VerifyEmailAccountStorage) *VerifyEmailAccountBiz {
	return &VerifyEmailAccountBiz{store: store}
}
func (biz *VerifyEmailAccountBiz) NewVerifyEmailAccount(ctx context.Context, email string, code int) error {
	sendCode, err := biz.store.FindVerifyVerifyCode(ctx, email)
	if err != nil {
		return err
	}
	if sendCode.Code != code {
		return fmt.Errorf("verify email code not match")
	}
	if err := biz.store.VerifyEmailAccount(ctx, email); err != nil {
		return err
	}
	return nil
}
