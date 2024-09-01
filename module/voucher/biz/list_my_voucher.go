package biz

import (
	"context"
	"main.go/module/voucher/model"
)

func (biz *VoucherBiz) NewListMyVoucher(ctx context.Context, userId int) (*[]model.VoucherUser, error) {
	vouchers, err := biz.store.ListMyVoucher(ctx, userId)
	if err != nil {
		return nil, err
	}
	return vouchers, nil
}
