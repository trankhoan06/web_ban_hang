package storage

import (
	"context"
	"main.go/module/item/model"
)

func (s *SqlModel) ListItemCategory(ctx context.Context, category string) (*[]model.TodoList, error) {
	var result []model.TodoList
	if err := s.db.Table("todo_items").Where("category=?", category).Find(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}
