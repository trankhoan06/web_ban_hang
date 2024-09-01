package biz

import (
	"context"
	"main.go/module/voucher/model"
)

func (biz *VoucherBiz) NewListVoucherVendor(ctx context.Context, vendorId int) (*[]model.Voucher, error) {
	vouchers, err := biz.store.ListVoucherVendor(ctx, vendorId)
	if err != nil {
		return nil, err
	}
	return vouchers, nil
}
