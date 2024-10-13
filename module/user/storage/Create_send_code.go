package storage

import (
	"context"
	"main.go/module/user/model"
)

func (s *SqlModel) CreateSendCode(ctx context.Context, data *model.CreateSendCode) error {
	db := s.db.Begin()
	if err := db.Create(data).Error; err != nil {
		db.Rollback()
		return err
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return err
	}
	return nil
}
