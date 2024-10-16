package biz

import (
	"context"
	"errors"
	"main.go/common"
	"main.go/module/user/model"
	"time"
)

func (biz *LoginBiz) NewVerifyCodeEmail(ctx context.Context, code int, token string, userId int, expire int) (*model.TokenSendEmail, error) {
	sendCode, err := biz.store.FindSendCode(ctx, map[string]interface{}{"token": token, "userId": userId})
	if err != nil {
		return nil, err
	}
	now := time.Now().Add(-7 * time.Hour)

	if now.After(sendCode.Expire) {
		return nil, errors.New("code is expire")
	}
	if code != sendCode.Code {
		return nil, errors.New("code is wrong")
	}
	if err := biz.store.VerifyEmail(ctx, userId); err != nil {
		return nil, err
	}
	var tokenEmail model.TokenSendEmail
	tokenEmail.IsEmail = true
	tokenEmail.Id = userId
	user:=model.RoleUserUser
	var payload=&common.Payload{
		URole: &user,
		UId: userId,
	}
	token1, err1:=biz.provider.Generate( payload,expire)
	if err1 != nil {
		return nil, err1
	}
	tokenEmail.Token=token1.Gettoken()
	return &tokenEmail, nil
}
