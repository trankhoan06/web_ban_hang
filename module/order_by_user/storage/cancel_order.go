package storage

import (
	"context"
	"main.go/module/order_by_user/model"
)

func (s *SqlModel) CancelOrder(ctx context.Context, id int) error {
	if err := s.db.Table("order_by_user").Where("id=?", id).Update("status", model.StatusOrderCancel).Error; err != nil {
		return err
	}
	return nil
}
