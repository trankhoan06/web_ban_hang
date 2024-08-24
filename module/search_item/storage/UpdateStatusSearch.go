package storage

import (
	"context"
	"main.go/module/search_item/model"
)

func (s *sqlModel) UpdateStatusSearch(ctx context.Context, keyword *model.Filter, status model.StatusSearch) error {
	if err := s.db.Table("search").Where("content=? and user_id= ?", keyword.Keyword, keyword.UserId).Update("status", status).Error; err != nil {
		return err
	}
	return nil
}
