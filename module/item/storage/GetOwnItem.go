package storage

import (
	"context"
	"main.go/common"
	"main.go/module/item/model"
)

func (s *SqlModel) GetOwnItem(ctx context.Context, cond map[string]interface{}, user map[string]interface{}) (*model.TodoList, error) {
	var data model.TodoList
	db := s.db.Table("todo_items").Where(user)
	if err := db.Where(cond).First(&data).Error; err != nil {
		return nil, common.ErrFoundItem
	}
	if *data.Status == model.StatusItemDeleted {
		return nil, common.ErrFindItem
	}
	return &data, nil
}
