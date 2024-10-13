package storage

import "context"

func (s *SqlModel) ChangePassword(ctx context.Context, userId int, password string) error {
	if err := s.db.Table("users").Where("id", userId).Update("password", password).Error; err != nil {
		return err
	}
	return nil
}
