package storage

import (
	"context"
	"main.go/module/item/model"
)

func (s *SqlModel) CreateSortItem(ctx context.Context, data *model.CreateSortItem) error {
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
