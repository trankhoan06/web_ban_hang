package storage

import (
	"context"
	"main.go/module/item/model"
)

func (s *SqlModel) ListItem(ctx context.Context, paging *model.Paging, filter *model.Filter, moreKey ...string) (*[]model.TodoList, error) {
	var result []model.TodoList
	db := s.db.Table("todo_items").Where("status<>?", "Deleted")
	if f := filter; f != nil {
		if v := f.Status; v != "" {
			db = db.Where("status=?", v)
		}
	}
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}
	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}
	if err := db.Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).Find(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}
