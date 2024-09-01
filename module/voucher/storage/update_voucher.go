package storage

import (
	"context"
	"main.go/module/voucher/model"
)

func (s *SqlModel) UpdateVoucher(ctx context.Context, voucherId int, data *model.UpdateVoucher) error {
	if err := s.db.Where("id=? and vendor_id=?", voucherId, data.VendorId).Updates(data).Error; err != nil {
	}
	return nil
}
