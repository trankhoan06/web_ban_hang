package storage

import (
	"context"
	"main.go/module/voucher/model"
)

func (s *SqlModel) AddVoucher(ctx context.Context, data *model.VoucherUser) error {
	db := s.db.Begin()
	if err := db.Create(data).Error; err != nil {
		db.Rollback()
		return err
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return err
	}
	return nil
}
