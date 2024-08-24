package storage

import (
	"context"
	"gorm.io/gorm"
	"main.go/module/search_item/model"
)

func (s *sqlModel) UpdateTimeSearch(ctx context.Context, keyword *model.Filter) error {
	if err := s.db.Table("search").Where("content=? and user_id=?", keyword.Keyword, keyword.UserId).Update("search_time", gorm.Expr("search_time + ?", 1)).Error; err != nil {
		return err
	}
	return nil
}
