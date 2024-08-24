package storage

import (
	"context"
	"main.go/module/item/model"
)

func (s *SqlModel) UpdateItem(ctx context.Context, cond map[string]interface{}, data *model.TodoUpdateItem) error {
	s.db.Begin()
	if err := s.db.Where(cond).Updates(data).Error; err != nil {
		s.db.Rollback()
		return err
	}
	if s.db.Commit().Error != nil {
		s.db.Rollback()
		return s.db.Commit().Error
	}
	return nil
}
