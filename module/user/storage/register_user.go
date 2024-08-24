package storage

import (
	"context"
	"main.go/module/user/model"
)

func (s *SqlModel) RegisterUser(ctx context.Context, user *model.CreateUser) error {
	if err := s.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
