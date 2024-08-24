package biz

import (
	"context"
	"main.go/module/user/model"
)

type UpdatePasswordForgotModuleStorage interface {
	ChangePasswordsto(ctx context.Context, change string, cond map[string]interface{}) error
	FindEmailForgotPassword(ctx context.Context, token string) (*model.SendCode, error)
	FindUser(ctx context.Context, cond map[string]interface{}) (*model.User, error)
	UpdateTokenCodeVerifyEmail(ctx context.Context, id int, token string) error
}
type UpdatePasswordForgotBiz struct {
	store UpdatePasswordForgotModuleStorage
	hash  Hasher
}

func NewUpdatePasswordForgotBiz(store UpdatePasswordForgotModuleStorage, hash Hasher) *UpdatePasswordForgotBiz {
	return &UpdatePasswordForgotBiz{store: store, hash: hash}
}
func (biz *UpdatePasswordForgotBiz) NewUpdatePasswordForgot(ctx context.Context, token string, password string) error {
	emailCode, err := biz.store.FindEmailForgotPassword(ctx, token)
	email := emailCode.Email
	if err != nil {
		return err
	}
	user, err := biz.store.FindUser(ctx, map[string]interface{}{"email": email})
	if err != nil {
		return err
	}
	password = biz.hash.Hash(password + user.Salt)
	if err := biz.store.ChangePasswordsto(ctx, password, map[string]interface{}{"email": email}); err != nil {
		return err
	}
	if err := biz.store.UpdateTokenCodeVerifyEmail(ctx, emailCode.Id, ""); err != nil {
		return err
	}
	return nil

}
