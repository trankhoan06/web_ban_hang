package storage

import (
	"context"
	"main.go/module/order_by_user/model"
)

func (s *SqlModel) CreateOrder(ctx context.Context, order *model.CreateOrder) error {
	if err := s.db.Create(order).Error; err != nil {
		return err
	}
	return nil
}
