package storage

import (
	"context"
	"main.go/module/user/model"
)

func (s *SqlModel) FindEmailForgotPassword(ctx context.Context, token string) (*model.SendCode, error) {
	var emailCode model.SendCode
	if err := s.db.Table("send_code_email").Where("token=?", token).First(&emailCode).Error; err != nil {
		return nil, err
	}
	return &emailCode, nil
}
