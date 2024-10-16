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

func (biz *LoginBiz) NewLogin(ctx context.Context, login *model.LoginUser, expire int) (*model.TokenSendEmail, error) {
	var verify model.TokenSendEmail
	user, errUser := biz.store.FindUser(ctx, map[string]interface{}{"email": login.Email})
	if errUser != nil {
		return nil, errors.New("email of password is incorrect")
	}
	if user.Password != biz.hash.Hash(user.Salt+login.Password) {
		return nil, errors.New("email of password is incorrect")
	}
	if *user.Status == model.StatusUserDeleted {
		return nil, errors.New("account is deleted")
	}
	if !user.IsEmail {

		verify.IsEmail = false
		token := common.GetSalt(30)
		verify.Token = token
		verify.Id = user.Id
		code := common.GenerateRandomCode()
		expire1 := time.Now().UTC().Add(-7 * time.Hour)
		expire1 = expire1.Add(1 * time.Minute)
		var sendCode model.CreateSendCode
		sendCode.Code = code
		sendCode.Token = token
		sendCode.Email = login.Email
		sendCode.Expire = expire1
		if err := biz.store.CreateSendCode(ctx, &sendCode); err != nil {
			return &verify, err
		}
		chanel := make(chan error, 1)
		wg := new(sync.WaitGroup)
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := emailSend.SendVerifyEmail(login.Email, code)
			if err != nil {
				fmt.Println(err)
				chanel <- err
			}

		}()
		wg.Wait()
		close(chanel)
		return &verify, <-chanel
	}
	verify.IsEmail = true
	var payload = &common.Payload{
		URole: user.Role,
		UId:   user.Id,
	}
	token, err := biz.provider.Generate(payload, expire)
	if err != nil {
		return nil, err
	}
	verify.Token = token.Gettoken()
	return &verify, nil

}
