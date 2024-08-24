package biz

import (
	"context"
	"errors"
	"fmt"
	"main.go/common"
	"main.go/component/tokenProvider"
	"main.go/module/user/model"
)

type loginStorage interface {
	FindUser(ctx context.Context, cond map[string]interface{}) (*model.User, error)
	CreateCodeVerifyEmail(ctx context.Context, code *model.CreateSendCode) error
}
type LoginUserBiz struct {
	store    loginStorage
	hash     Hasher
	provider tokenProvider.Provider
	expiry   int
}

func NewLoginUserBiz(store loginStorage, hash Hasher, provider tokenProvider.Provider, expiry int) *LoginUserBiz {
	return &LoginUserBiz{
		store:    store,
		hash:     hash,
		provider: provider,
		expiry:   expiry,
	}
}
func (biz *LoginUserBiz) LoginUser(ctx context.Context, data *model.LoginUser, emailSend *model.CreateSendCode) (tokenProvider.Token, error) {
	user, err := biz.store.FindUser(ctx, map[string]interface{}{"email": data.Email})
	if err != nil {
		return nil, errors.New("email of password error")
	}
	data.PassWord = biz.hash.Hash(data.PassWord + user.Salt)
	if user.PassWord != data.PassWord {
		return nil, errors.New("email of password error")
	}
	var payload = &common.Payload{
		URole: user.Role,
		UId:   user.UserId,
	}
	if !*user.IsEmail {
		return nil, fmt.Errorf("please verify email")
	}
	token, err := biz.provider.Generate(payload, biz.expiry)
	if err != nil {
		return nil, err
	}
	if err := biz.store.CreateCodeVerifyEmail(ctx, emailSend); err != nil {
		return nil, err
	}
	return token, nil

}
