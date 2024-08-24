package storage

import (
	"context"
	"main.go/module/notify/model"
)

func (s *SqlModel) SendNotify(ctx context.Context, data *model.CreateNotify) error {
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
