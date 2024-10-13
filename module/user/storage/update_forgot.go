package storage

import (
	"context"
)

func (s *SqlModel) UpdateForgot(ctx context.Context, token string, userId int, cond map[string]interface{}) error {
	if err := s.db.Table("send_code_email").Where("token=? and user_id=?", token, userId).
		Updates(cond).Error; err != nil {
		return err
	}
	return nil
}
