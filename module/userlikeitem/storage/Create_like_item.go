package storage

import (
	"context"
	"main.go/module/userlikeitem/model"
)

func (s *SqlModel) CreateUserLike(ctx context.Context, data *model.LikeItem) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}
	return nil
}
