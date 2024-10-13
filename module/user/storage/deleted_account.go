package storage

import (
	"context"
	"main.go/module/user/model"
)

func (s *SqlModel) DeletedAccount(ctx context.Context, id int) error {
	if err := s.db.Table("users").Where("id=?", id).Update("status", model.StatusUserDeleted).Error; err != nil {
		return err
	}
	return nil
}
