package storage

import (
	"context"
	"main.go/module/search_item/model"
)

func (s *sqlModel) ListSearch(ctx context.Context, user int) (*[]model.Search, error) {
	var searchs []model.Search
	if err := s.db.Table("search").Where("user_id=? and status <> ?", user, model.StatusSearchInactive).Order("update_at desc").Limit(10).Find(&searchs).Error; err != nil {
		return nil, err
	}
	return &searchs, nil
}
