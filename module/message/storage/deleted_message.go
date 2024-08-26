package storage

import (
	"context"
	"main.go/module/message/model"
)

func (s *SqlModel) DeletedMessage(ctx context.Context, id int, column string) error {
	db := s.db.Table("message").Where("id=?", id)
	if err := db.Update(column, model.StatusRemove).Error; err != nil {
		return err
	}
	return nil
}
