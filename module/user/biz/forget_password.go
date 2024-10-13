package biz

import (
	"context"
	"errors"
	"main.go/common"
	emailSend "main.go/email"
	"main.go/module/user/model"
	"sync"
	"time"
)

func (biz *UserBiz) NewForgotPassWord(ctx context.Context, email string) (*model.TokenSendEmail, error) {
	user, err := biz.store.FindUser(ctx, map[string]interface{}{"email": email})
	if err != nil {
		return nil, errors.New("email doesn't exist")
	}
	token := common.GetSalt(30)
	var sendEmail model.TokenSendEmail
	sendEmail.Token = token
	sendEmail.Id = user.Id
	sendEmail.IsEmail = user.IsEmail
	var sendCode model.CreateSendCode
	expire := time.Now().UTC().Add(-7 * time.Hour)
	expire = expire.Add(1 * time.Minute)
	code := common.GenerateRandomCode()
	sendCode.Code = code
	sendCode.Token = token
	sendCode.Email = email
	sendCode.Expire = expire
	sendCode.UserId = user.Id
	if err := biz.store.CreateSendCode(ctx, &sendCode); err != nil {
		return nil, err
	}
	chanel := make(chan error, 1)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := emailSend.SendForgotPassword(email, code); err != nil {
			chanel <- err
		}
	}()
	wg.Wait()
	close(chanel)
	return &sendEmail, <-chanel
}
