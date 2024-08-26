package storage

import (
	"context"
	"main.go/module/message/model"
)

func (s *SqlModel) CreateMessage(ctx context.Context, message *model.CreateMessage) error {
	db := s.db.Begin()
	if err := db.Create(message).Error; err != nil {
		db.Rollback()
		return err
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return err
	}
	return nil
}
