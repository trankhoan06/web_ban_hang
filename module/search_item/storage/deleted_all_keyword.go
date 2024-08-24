package storage

import (
	"context"
	"main.go/module/search_item/model"
)

func (s *sqlModel) DeletedAllSearch(ctx context.Context, userId int) error {
	if err := s.db.Table("search").Where("user_id=?", userId).Update("status", model.StatusSearchInactive).Error; err != nil {
		return err
	}
	return nil
}
