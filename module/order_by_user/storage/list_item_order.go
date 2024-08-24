package storage

import (
	"context"
	"main.go/module/order_by_user/model"
)

func (s *SqlModel) GetListItemOrder(ctx context.Context, itemId int) (*[]model.AppreciateOrder, error) {
	var appreciate []model.AppreciateOrder
	if err := s.db.Table("order_by_user").Where("item_id = ?", itemId).Find(&appreciate).Error; err != nil {
		return nil, err
	}
	return &appreciate, nil
}
