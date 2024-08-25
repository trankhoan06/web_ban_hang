package storage

import "context"

func (s *SqlModel) UpdateRole(ctx context.Context, userID int, role string) error {
	if err := s.db.Table("users").Where("id=?", userID).Update("role", role).Error; err != nil {
		return err
	}
	return nil
}
