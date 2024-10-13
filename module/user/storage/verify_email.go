package storage

import "context"

func (s *SqlModel) VerifyEmail(ctx context.Context, userId int) error {
	if err := s.db.Table("users").Where("id", userId).Update("is_email", true).Error; err != nil {
		return err
	}
	return nil
}
