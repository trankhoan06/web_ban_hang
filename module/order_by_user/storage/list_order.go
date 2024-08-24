package storage

import (
	"context"
	"main.go/module/order_by_user/model"
)

func (s *SqlModel) ListOrder(ctx context.Context, userId int, column string, moreKeyUser []string, moreKeyItem ...string) (*[]model.Order, error) {
	var orders []model.Order
	cond := map[string]interface{}{column: userId}
	db := s.db.Table("order_by_user").Where(cond)
	db = db.Where("status <> ? and status <> ? ", model.StatusOrderCancel, model.StatusOrderDone)
	for i := range moreKeyUser {
		db = db.Preload(moreKeyUser[i])
	}
	for i := range moreKeyItem {
		db = db.Preload(moreKeyItem[i])
	}
	if err := db.Find(&orders).Error; err != nil {
		return nil, err
	}
	return &orders, nil
}
