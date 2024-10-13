package biz

import (
	"context"
	"errors"
	"fmt"
	"main.go/common"
	emailSend "main.go/email"
	"main.go/module/user/model"
	"sync"
	"time"
)

func (biz *RegisterBiz) NewRegister(ctx context.Context, data *model.CreateUser) (*model.TokenSendEmail, error) {
	user, errUser := biz.store.FindUser(ctx, map[string]interface{}{"email": data.Email})
	if errUser == nil {
		fmt.Println(errUser)
		if *user.Status == model.StatusUserDeleted {
			data.Salt = user.Salt
			data.Password = biz.hash.Hash(user.Salt + data.Password)
			var v model.TokenSendEmail
			v.IsEmail = true
			v.Token = ""
			if err := biz.store.UpdateEmailDeleted(ctx, data); err != nil {
				return nil, err
			}
			return &v, nil

		} else if *user.Status == model.StatusUserDoing {
			return nil, errors.New("email already exists")
		}
	}
	data.Salt = common.GetSalt(50)
	data.Password = biz.hash.Hash(data.Salt + data.Password)
	if err := biz.store.RegisterAccount(ctx, data); err != nil {
		return nil, err
	}
	var verify model.TokenSendEmail
	verify.IsEmail = false
	token := common.GetSalt(30)
	verify.Token = token
	code := common.GenerateRandomCode()
	expire := time.Now().UTC().Add(-7 * time.Hour)
	expire = expire.Add(1 * time.Minute)
	var sendCode model.CreateSendCode
	sendCode.Code = code
	sendCode.Token = token
	sendCode.Email = data.Email
	sendCode.Expire = expire
	verify.Id = data.Id
	if err := biz.store.CreateSendCode(ctx, &sendCode); err != nil {
		return &verify, err
	}
	chanel := make(chan error, 1)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := emailSend.SendVerifyEmail(data.Email, code)
		if err != nil {
			fmt.Println(err)
			chanel <- err
		}

	}()
	defer close(chanel)
	return &verify, <-chanel
}
