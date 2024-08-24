package storage

import (
	"context"
	"main.go/module/user/model"
)

func (s *SqlModel) ListUserId(ctx context.Context) (*[]model.LIstUserId, error) {
	var result []model.LIstUserId
	if err := s.db.Table("users").Where("status<>?", 0).Find(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}
