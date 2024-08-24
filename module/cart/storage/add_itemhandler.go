package storage

import (
	"context"
	"main.go/module/cart/model"
)

func (s *SqlModel) AddItem(ctx context.Context, item *model.CartCreateUser) error {
	db := s.db.Begin()
	if err := db.Create(item).Error; err != nil {
		db.Rollback()
		return err
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return err
	}
	return nil
}
