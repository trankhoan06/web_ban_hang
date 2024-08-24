package storage

import (
	"context"
)

func (s *SqlModel) UpdateTokenCodeVerifyEmail(ctx context.Context, id int, token string) error {
	if err := s.db.Table("send_code_email").Where("id=?", id).Update("token", token).Error; err != nil {
		return err
	}
	return nil
}
