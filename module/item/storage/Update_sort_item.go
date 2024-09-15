package storage

import (
	"context"
	"main.go/module/item/model"
)

func (s *SqlModel) UpdateSortItem(ctx context.Context, data *model.UpdateSortItem) error {
	if err := s.db.Where("id=?", data.Id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
