package storage

import (
	"context"
	"main.go/module/voucher/model"
)

func (s *SqlModel) ListMyVoucher(ctx context.Context, userId int) (*[]model.VoucherUser, error) {
	var voucherUsers []model.VoucherUser
	if err := s.db.Where("user_id = ? and status<>?", userId, model.StatusVoucherExpire).Find(&voucherUsers).Error; err != nil {
		return nil, err
	}
	return &voucherUsers, nil
}
