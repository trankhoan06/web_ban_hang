package storage

import (
	"context"
	"errors"
	modelitem "main.go/module/item/model"
	"main.go/module/search_item/model"
)

func (s *sqlModel) SearchItem(ctx context.Context, keyword string, category *model.CategorySearch, moreKey ...string) (*[]modelitem.TodoList, error) {
	db := s.db.Table("todo_items").Where("title LIKE ? or description LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}
	category.Process()
	if category.Name != "" {
		order := category.Name + " " + category.Arrangement
		db = db.Order(order)
	} else if category.Name == "" {
		return nil, errors.New("category name is empty")
	}
	var items []modelitem.TodoList
	if err := db.Where("status<>?", "Deleted").Order("amount_sold DESC, update_at DESC").Find(&items).Error; err != nil {
		return nil, err
	}
	return &items, nil
}
