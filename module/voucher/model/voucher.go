package model

import "time"

type TypeVoucher int
type StatusVoucher int

const (
	TypeVoucherFashion TypeVoucher = iota
	TypeVoucherCook
	TypeVoucherShip
)
const (
	StatusVoucherExpire StatusVoucher = iota
	StatusVoucherDoing
)

type Voucher struct {
	Id            int           `json:"id" gorm:"column:id"`
	VendorId      int           `json:"vendor_id" gorm:"column:vendor_id"`
	Name          string        `json:"name" gorm:"column:name"`
	Type          TypeVoucher   `json:"type" gorm:"column:type"`
	Status        StatusVoucher `json:"status" gorm:"column:status"`
	Discount      int           `json:"discount" gorm:"column:discount"`
	MaxDiscount   int           `json:"max_discount" gorm:"column:max_discount"`
	Amount        int           `json:"amount" gorm:"column:amount"`
	MinimumSingle int           `json:"minimum_single" gorm:"column:minimum_single"`
	ApplyAll      bool          `json:"apply_all" gorm:"column:apply_all"`
	Effective     time.Time     `json:"effective" gorm:"column:effective"`
	Expire        time.Time     `json:"expire" gorm:"column:expire"`
}

type CreateVoucher struct {
	VendorId      int         `json:"-" gorm:"column:vendor_id"`
	Name          string      `json:"name" gorm:"column:name"`
	Type          TypeVoucher `json:"type" gorm:"column:type"`
	Discount      int         `json:"discount" gorm:"column:discount"`
	MaxDiscount   int         `json:"max_discount" gorm:"column:max_discount"`
	MinimumSingle int         `json:"minimum_single" gorm:"column:minimum_single"`
	Amount        int         `json:"amount" gorm:"column:amount"`
	ApplyAll      bool        `json:"apply_all" gorm:"column:apply_all"`
	Effective     time.Time   `json:"effective" gorm:"column:effective"`
	Expire        time.Time   `json:"expire" gorm:"column:expire"`
}
type UpdateVoucher struct {
	VendorId      int          `json:"-" gorm:"column:vendor_id"`
	Name          *string      `json:"name" gorm:"column:name"`
	Type          *TypeVoucher `json:"type" gorm:"column:type"`
	Discount      *int         `json:"discount" gorm:"column:discount"`
	MaxDiscount   *int         `json:"max_discount" gorm:"column:max_discount"`
	MinimumSingle *int         `json:"minimum_single" gorm:"column:minimum_single"`
	Amount        *int         `json:"amount" gorm:"column:amount"`
	Effective     time.Time    `json:"effective" gorm:"column:effective"`
	Expire        time.Time    `json:"expire" gorm:"column:expire"`
}
type VoucherUser struct {
	UserId       int       `json:"user_id" gorm:"column:user_id"`
	VoucherId    int       `json:"voucher_id" gorm:"column:voucher_id"`
	OwnerVoucher *Voucher  `json:"owner_voucher" gorm:"column:foreignkey:VoucherId;references:Id"`
	Effective    time.Time `json:"effective" gorm:"column:effective"`
	Expire       time.Time `json:"expire" gorm:"column:expire"`
}

func (CreateVoucher) TableName() string { return "voucher" }
func (VoucherUser) TableName() string   { return "voucher_user" }
func (UpdateVoucher) TableName() string { return "voucher" }
func (Voucher) TableName() string       { return "voucher" }
