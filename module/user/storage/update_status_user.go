package storage

import (
	"context"
	"main.go/module/user/model"
)

func (s *SqlModel) UpdateStatusUser(ctx context.Context, userId int, status model.StatusUser) error {
	if err := s.db.Table("users").Where("id=?", userId).Update("status", 0).Error; err != nil {
		return err
	}
	return nil
}
