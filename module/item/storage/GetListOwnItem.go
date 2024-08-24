package storage

import (
	"context"
	"main.go/module/item/model"
)

func (s *SqlModel) ListOwnItem(ctx context.Context, userId int, moreKey ...string) (*[]model.TodoList, error) {
	var result []model.TodoList
	db := s.db.Table("todo_items").Where("status<>? and user_id=?", "Deleted", userId)
	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}
	if err := db.Find(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}
