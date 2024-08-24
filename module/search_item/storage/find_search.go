package storage

import (
	"context"
	"main.go/module/search_item/model"
)

func (s *sqlModel) FindSearch(ctx context.Context, keyword *model.Filter) (*model.Search, error) {
	var search model.Search
	if err := s.db.Table("search").Where("content=? and user_id=?", keyword.Keyword, keyword.UserId).First(&search).Error; err != nil {
		return nil, err
	}
	return &search, nil
}
