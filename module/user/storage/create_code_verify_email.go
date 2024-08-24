package storage

import (
	"context"
	"main.go/module/user/model"
)

func (s *SqlModel) CreateCodeVerifyEmail(ctx context.Context, code *model.CreateSendCode) error {
	if err := s.db.Create(&code).Error; err != nil {
		return err
	}
	return nil
}
