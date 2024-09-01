package biz

import (
	"context"
	"main.go/module/voucher/model"
)

func (biz *VoucherBiz) NewCreateVoucher(ctx context.Context, data *model.CreateVoucher) error {
	if err := biz.store.CreateVoucher(ctx, data); err != nil {
		return err
	}
	return nil
}
