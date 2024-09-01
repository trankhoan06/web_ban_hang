package storage

import (
	"context"
	"main.go/module/voucher/model"
)

func (s *SqlModel) DeletedVoucherUser(ctx context.Context, ownerVoucher int) error {
	if err := s.db.Table("voucher_user").Where("owner_voucher=? ", ownerVoucher).Update("status", model.StatusVoucherDeleted).Error; err != nil {
		return err
	}
	return nil
}
