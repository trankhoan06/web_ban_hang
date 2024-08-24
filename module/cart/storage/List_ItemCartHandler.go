package storage

import (
	"context"
	"main.go/module/cart/model"
)

func (s *SqlModel) ListItemCart(ctx context.Context, userId int, moreKey ...string) (*[]model.CartUser, error) {
	var result []model.CartUser
	db := s.db.Where("user_id = ? and status<> ?", userId, model.StatusRemove)
	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}
	if err := db.Find(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}
