package storage

import (
	"context"
	"main.go/common"
	"main.go/module/item/model"
)

func (s *SqlModel) GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoList, error) {
	var data model.TodoList
	db := s.db.Preload("Owner")
	if err := db.Table("todo_items").Where(cond).First(&data).Error; err != nil {
		return nil, common.ErrFindItem
	}
	if *data.Status == model.StatusItemDeleted {
		return nil, common.ErrFindItem
	}
	return &data, nil
}
