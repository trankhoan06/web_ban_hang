package biz

import (
	"context"
	"main.go/module/user/model"
)

type FindUserStorage interface {
	FindUser(ctx context.Context, cond map[string]interface{}) (*model.User, error)
	CreateCodeVerifyEmail(ctx context.Context, code *model.CreateSendCode) error
}
type findUserBiz struct {
	store FindUserStorage
}

func NewFindUserBiz(store FindUserStorage) *findUserBiz {
	return &findUserBiz{store: store}
}
func (biz *findUserBiz) NewSendCode(ctx context.Context, cond map[string]interface{}, code *model.CreateSendCode) error {
	//now := time.Now()
	//if now.Before(code.ExpireAt) {
	//	remainingSeconds := code.ExpireAt.Sub(now).Seconds()
	//	return fmt.Errorf("resend the code after %f seconds", remainingSeconds)
	//}
	_, err := biz.store.FindUser(ctx, cond)
	if err != nil {
		return err
	}
	if err := biz.store.CreateCodeVerifyEmail(ctx, code); err != nil {
		return err
	}
	return nil

}
