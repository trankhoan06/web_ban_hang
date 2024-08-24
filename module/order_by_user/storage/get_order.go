package storage

import (
	"context"
	"errors"
	"main.go/module/order_by_user/model"
)

func (s *SqlModel) GetOrder(ctx context.Context, addressColumn string, id int, userId int, keyUser, keyItem string) (*model.Order, error) {
	var orders model.Order
	db := s.db.Table("order_by_user").Where(addressColumn, id, userId)
	db = db.Preload(keyUser)
	db = db.Preload(keyItem)
	if err := db.First(&orders).Error; err != nil {
		return nil, errors.New("You don't this order")
	}
	return &orders, nil
}
