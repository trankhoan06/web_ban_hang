package biz

import (
	"context"
	"main.go/common"
	"main.go/module/voucher/model"
)

func (biz *VoucherBiz) NewUpdateVoucher(ctx context.Context, voucherID int, data *model.UpdateVoucher) error {
	voucher, err := biz.store.GetVoucher(ctx, voucherID)
	if err != nil {
		return err
	}
	if voucher.VendorId != data.VendorId {
		return common.ErrNoPermiss
	}
	if err := biz.store.UpdateVoucher(ctx, voucherID, data); err != nil {
		return err
	}
	return nil
}
