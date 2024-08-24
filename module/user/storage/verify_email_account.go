package storage

import "context"

func (s *SqlModel) VerifyEmailAccount(ctx context.Context, email string) error {
	if err := s.db.Table("users").Where("email=?", email).Update("is_email", true).Error; err != nil {
		return err
	}
	return nil
}
