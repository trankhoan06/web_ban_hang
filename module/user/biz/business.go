package biz

import (
	"context"
	"main.go/component/tokenProvider"
	"main.go/module/user/model"
)

type BusinessUser interface {
	ListUserId(ctx context.Context) (*[]model.ListUserId, error)
	UpdateUser(ctx context.Context, data *model.UpdateUser) error
	RegisterAccount(ctx context.Context, data *model.CreateUser) error
	FindUser(ctx context.Context, cond map[string]interface{}) (*model.User, error)
	DeletedAccount(ctx context.Context, id int) error
	CreateSendCode(ctx context.Context, data *model.CreateSendCode) error
	FindSendCode(ctx context.Context, cond map[string]interface{}) (*model.SendCode, error)
	UpdateEmailDeleted(ctx context.Context, data *model.CreateUser) error
	ChangePassword(ctx context.Context, userId int, password string) error
	VerifyEmail(ctx context.Context, userId int) error
	UpdateForgot(ctx context.Context, token string, userId int, cond map[string]interface{}) error
}
type Hasher interface {
	Hash(str string) string
}
type RegisterBiz struct {
	store BusinessUser
	hash  Hasher
}

func NewRegisterBiz(store BusinessUser, hash Hasher) *RegisterBiz {
	return &RegisterBiz{store: store, hash: hash}
}

type UserBiz struct {
	store BusinessUser
}

func NewUserBiz(store BusinessUser) *UserBiz {
	return &UserBiz{store: store}
}

type LoginBiz struct {
	store    BusinessUser
	hash     Hasher
	provider tokenProvider.Provider
}

func NewLoginBiz(store BusinessUser, hash Hasher, provider tokenProvider.Provider) *LoginBiz {
	return &LoginBiz{store: store, hash: hash, provider: provider}
}
