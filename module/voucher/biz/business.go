package biz

import (
	"context"
	"main.go/module/voucher/model"
)

type VoucherStorage interface {
	CreateVoucher(ctx context.Context, data *model.CreateVoucher) error
	GetVoucher(ctx context.Context, id int) (*model.Voucher, error)
	ListVoucherVendor(ctx context.Context, vendorId int) (*[]model.Voucher, error)
	ListMyVoucher(ctx context.Context, userId int) (*[]model.VoucherUser, error)
	DeletedVoucherVendor(ctx context.Context, VoucherId, userId int) error
	DeletedVoucherUser(ctx context.Context, VoucherId int) error
}

type VoucherBiz struct {
	store VoucherStorage
}

func NewVoucherBiz(store VoucherStorage) *VoucherBiz {
	return &VoucherBiz{store}
}
