package storage

import (
	"context"
	"main.go/module/item/model"
)

func (s *SqlModel) ListSortItem(ctx context.Context, idItem int) (*[]model.SortItem, error) {
	var item []model.SortItem
	if err := s.db.Where("item_id=?", idItem).Find(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}
