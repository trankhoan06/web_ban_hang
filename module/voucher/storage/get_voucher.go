package storage

import (
	"context"
	"main.go/module/voucher/model"
)

func (s *SqlModel) GetVoucher(ctx context.Context, id int) (*model.Voucher, error) {
	var voucher model.Voucher
	if err := s.db.Where("id=? and status<>?", id, model.StatusVoucherDeleted).First(&voucher).Error; err != nil {
		return nil, err
	}
	return &voucher, nil
}
