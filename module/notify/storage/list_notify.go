package storage

import (
	"context"
	"main.go/module/notify/model"
)

func (s *SqlModel) ListNotify(ctx context.Context, userId int) (*[]model.Notify, error) {
	var data []model.Notify
	if err := s.db.Table("notify").Where("user_id=?", userId).Find(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
