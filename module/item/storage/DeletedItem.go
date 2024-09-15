package storage

import (
	"context"
	"main.go/module/item/model"
)

func (s *SqlModel) DeletedItem(ctx context.Context, cond map[string]interface{}) error {
	if err := s.db.Table("todo_items").Where(cond).Updates(map[string]interface{}{
		"status": model.StatusItemDeleted,
	}).Error; err != nil {
		return err
	}
	return nil
}
