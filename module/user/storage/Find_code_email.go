package storage

import (
	"context"
	"main.go/module/user/model"
)

func (s *SqlModel) FindVerifyVerifyCode(ctx context.Context, email string) (*model.SendCode, error) {
	var result []model.SendCode
	if err := s.db.Table("send_code_email").Where("email = ?", email).Find(&result).Error; err != nil {
		return nil, err
	}
	length := len(result)
	return &result[length-1], nil
}
