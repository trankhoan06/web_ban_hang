package storage

import (
	"context"
	"main.go/module/comment/model"
)

func (s *SqlModel) UpdateOldComment(ctx context.Context, data *model.OldComment) error {
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
