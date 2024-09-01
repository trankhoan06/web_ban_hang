package biz

import (
	"context"
	"main.go/module/voucher/model"
)

func (biz *VoucherBiz) AddVoucher(ctx context.Context, VoucherId int, userId int) (*model.Voucher, error) {
	voucher, err := biz.store.GetVoucher(ctx, VoucherId)
	if err != nil {
		return nil, err
	}
	voucherUser := &model.VoucherUser{
		UserId:       userId,
		OwnerVoucher: voucher.VendorId,
		VoucherId:    VoucherId,
		Effective:    voucher.Effective,
		Expire:       voucher.Expire,
	}
	if err := biz.store.AddVoucher(ctx, voucherUser); err != nil {
		return nil, err
	}
	return voucher, nil
}
