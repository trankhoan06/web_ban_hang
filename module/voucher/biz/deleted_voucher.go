package biz

import (
	"context"
	"errors"
	"main.go/common"
)

func (biz *VoucherBiz) NewDeletedVoucher(ctx context.Context, voucherId, userId int) error {
	voucher, err := biz.store.GetVoucher(ctx, voucherId)
	if err != nil {
		return errors.New("Voucher has been deleted or no exist")
	}
	if voucher.VendorId != userId {
		return common.ErrNoPermiss
	}
	if err := biz.store.DeletedVoucherVendor(ctx, voucherId, userId); err != nil {
		return err
	}
	if err := biz.store.DeletedVoucherUser(ctx, userId); err != nil {
		return err
	}
	return nil
}
