package storage

import (
	"context"
	"main.go/module/voucher/model"
)

func (s *SqlModel) DeletedVoucherVendor(ctx context.Context, VoucherId, userId int) error {
	if err := s.db.Table("voucher").Where("vendor_id=? and id=?", userId, VoucherId).Update("status", model.StatusVoucherDeleted).Error; err != nil {
		return err
	}
	return nil
}
