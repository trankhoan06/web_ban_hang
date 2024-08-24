package storage

import (
	"context"
)

func (s *SqlModel) ChangePasswordsto(ctx context.Context, change string, cond map[string]interface{}) error {
	if err := s.db.Table("users").Where(cond).Update("password", change).Error; err != nil {
		return err
	}
	return nil
}
