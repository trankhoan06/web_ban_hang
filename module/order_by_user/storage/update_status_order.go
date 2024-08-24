package storage

import (
	"context"
	"main.go/module/order_by_user/model"
)

func (s *SqlModel) UpdateStatusOrder(ctx context.Context, id int, status model.StatusOrder) error {
	if err := s.db.Table("order_by_user").Where("id=?", id).Update("status", status).Error; err != nil {
		return err
	}
	return nil
}
