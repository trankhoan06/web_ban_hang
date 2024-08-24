package storage

import (
	"context"
	"main.go/module/order_by_user/model"
)

func (s *SqlModel) UpdateOrder(ctx context.Context, data *model.UpdateOrder) error {
	if err := s.db.Where("id=?", &data.Id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
