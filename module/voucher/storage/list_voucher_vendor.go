package storage

import (
	"context"
	"main.go/module/voucher/model"
)

func (s *SqlModel) ListVoucherVendor(ctx context.Context, vendorId int) (*[]model.Voucher, error) {
	var vouchers []model.Voucher
	if err := s.db.Where("vendor_id=?", vendorId).Find(&vouchers).Error; err != nil {
		return nil, err
	}
	return &vouchers, nil
}
