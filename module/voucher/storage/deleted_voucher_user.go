package storage

import (
	"context"
	"main.go/module/voucher/model"
)

func (s *SqlModel) DeletedVoucherUser(ctx context.Context, VoucherId int) error {
	if err := s.db.Table("voucher_user").Where("voucher_id=? ", VoucherId).Update("status", model.StatusVoucherExpire).Error; err != nil {
		return err
	}
	return nil
}
