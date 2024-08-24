package biz

import (
	"context"
	"errors"
	"main.go/common"
	"main.go/module/user/model"
	"strings"
)

type RegisterStorage interface {
	FindUser(ctx context.Context, cond map[string]interface{}) (*model.User, error)
	RegisterUser(ctx context.Context, user *model.CreateUser) error
	CreateCodeVerifyEmail(ctx context.Context, code *model.CreateSendCode) error
	UpdateStatusUser(ctx context.Context, userId int, status model.StatusUser) error
}
type Hasher interface {
	Hash(str string) string
}
type RegisterBiz struct {
	store RegisterStorage
	hash  Hasher
}

func NewRegisterbiz(store RegisterStorage, hash Hasher) *RegisterBiz {
	return &RegisterBiz{store: store, hash: hash}
}
func (biz *RegisterBiz) NewRegisterUser(ctx context.Context, user *model.CreateUser, emailSend *model.CreateSendCode) error {
	user.Email = strings.ToLower(user.Email)
	data, _ := biz.store.FindUser(ctx, map[string]interface{}{"email": user.Email})
	if data.Status == 1 {
		return errors.New("Email has been registered")

	} else if data.Status == 0 {
		if err := biz.store.UpdateStatusUser(ctx, user.UserId, model.StatusUserActive); err != nil {
			return err
		}
		return nil
	}
	salt := common.GetSalt(50)
	user.PassWord = biz.hash.Hash(user.PassWord + salt)
	user.Salt = salt
	user.Role = "user"
	if err := biz.store.RegisterUser(ctx, user); err != nil {
		return err

	}
	if err := biz.store.CreateCodeVerifyEmail(ctx, emailSend); err != nil {
		return err
	}
	return nil
}
