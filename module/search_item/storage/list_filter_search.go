package storage

import (
	"context"
	"main.go/module/search_item/model"
)

func (s *sqlModel) ListFilterSearch(ctx context.Context, keyword string) (*[]model.Search, error) {
	var data []model.Search
	if err := s.db.Table("search").Where("content like ?", "%"+keyword+"%").
		Select("content, SUM(search_time) AS total_time,  MAX(update_at) AS last_update").
		Group("content").
		Order("total_time desc, last_update desc").
		Limit(10).
		Find(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
