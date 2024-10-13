package storage

import (
	"context"
	"main.go/module/user/model"
)

func (s *SqlModel) ListUserId(ctx context.Context) (*[]model.ListUserId, error) {
	var result []model.ListUserId
	if err := s.db.Table("users").Where("status<>?", 0).Find(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}
