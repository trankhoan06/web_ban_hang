package storage

import (
	"context"
	"main.go/module/user/model"
)

func (s *SqlModel) UpdateEmailDeleted(ctx context.Context, data *model.CreateUser) error {
	if err := s.db.Where("email=?", data.Email).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
