package storage

import (
	"context"
	"main.go/module/user/model"
)

func (s *SqlModel) UpdateUser(ctx context.Context, data *model.UpdateUser, cond map[string]interface{}) error {
	if err := s.db.Where(cond).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
