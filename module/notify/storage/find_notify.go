package storage

import (
	"context"
	"main.go/module/notify/model"
)

func (s *SqlModel) FindNotify(ctx context.Context, id int) (*model.Notify, error) {
	var data model.Notify
	if err := s.db.Table("notify").Where("id=?", id).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
