package storage

import (
	"context"
	"main.go/module/search_item/model"
)

func (s *sqlModel) CreateSearch(ctx context.Context, keyword *model.Filter) error {
	db := s.db.Begin()
	if err := db.Table("search").Create(keyword).Error; err != nil {
		db.Rollback()
		return err
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return err
	}
	return nil
}
