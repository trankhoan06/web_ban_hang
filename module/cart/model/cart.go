package model

import (
	"main.go/module/item/model"
	"time"
)

type CartCreateUser struct {
	UserId int `json:"-" gorm:"column:user_id"`
	ItemId int `json:"item_id" gorm:"column:item_id"`
	Amount int `json:"amount" gorm:"column:amount"`
}
type CartUser struct {
	UserId   int               `json:"-" gorm:"column:user_id"`
	ItemId   int               `json:"item_id" gorm:"column:item_id"`
	Owner    *model.SimpleItem `json:"owner" gorm:"foreignkey:ItemId;references:Id"`
	Status   *statusCart       `json:"-" gorm:"column:status"`
	Amount   int               `json:"amount" gorm:"column:amount"`
	CreateAt *time.Time        `json:"create_at" gorm:"column:create_at"`
	UpdateAt *time.Time        `json:"update_at" gorm:"column:update_at"`
}
type CartUpdateUser struct {
	UserId int         `json:"-" gorm:"column:user_id"`
	ItemId int         `json:"-" gorm:"column:item_id"`
	Amount int         `json:"amount" gorm:"column:amount"`
	Status *statusCart `json:"-" gorm:"column:status"`
}

const OwnerItem string = "Owner"

func (CartUser) TableName() string       { return "cart_user" }
func (CartCreateUser) TableName() string { return "cart_user" }
func (CartUpdateUser) TableName() string { return "cart_user" }

type statusCart int

const (
	StatusRemove statusCart = iota
	StatusActive
)
