package model

import (
	modelitem "main.go/module/item/model"
	"main.go/module/user/model"
	"time"
)

type StatusOrder int

const (
	StatusOrderCancel StatusOrder = iota
	StatusOrderPrepare
	StatusOrderDoing
	StatusOrderDone
)

type StatusOrderItem struct {
	Status StatusOrder `json:"status"`
}
type Order struct {
	Id        int                   `json:"id" gorm:"column:id"`
	SellId    int                   `json:"-" gorm:"column:sell_id"`
	UserId    int                   `json:"-" gorm:"column:user_id"`
	OwnerUser *model.SimpleUser     `json:"owner_user" gorm:"foreignkey:SellId;references:id"`
	ItemId    int                   `json:"-" gorm:"column:item_id"`
	OwnerItem *modelitem.SimpleItem `json:"owner_item" gorm:"foreignkey:ItemId;references:id"`
	Status    StatusOrder           `json:"status" gorm:"column:status"`
	Amount    int                   `json:"amount" gorm:"column:amount"`
	Address   string                `json:"address" gorm:"column:address"`
	Telephone string                `json:"telephone" gorm:"column:telephone"`
	CreateAt  time.Time             `json:"create_at" gorm:"column:create_at"`
	UpdateAt  time.Time             `json:"update_at" gorm:"column:update_at"`
}
type CreateOrder struct {
	UserId    int    `json:"-" gorm:"column:user_id"`
	SellId    int    `json:"-" gorm:"column:sell_id"`
	ItemId    int    `json:"item_id" gorm:"column:item_id"`
	Amount    int    `json:"amount" gorm:"column:amount"`
	Address   string `json:"address" gorm:"column:address"`
	Telephone string `json:"telephone" gorm:"column:telephone"`
}
type UpdateOrder struct {
	Id         int     `json:"id" gorm:"column:id"`
	UserId     int     `json:"-" gorm:"column:user_id"`
	Appreciate int     `json:"appreciate" gorm:"column:appreciate"`
	Address    *string `json:"address" gorm:"column:address"`
	Telephone  *string `json:"telephone" gorm:"column:telephone"`
}

type AmountSoldOrder struct {
	AmountSold int `json:"amount_sold" gorm:"column:amount_sold"`
}

type AppreciateOrder struct {
	Appreciate int `json:"appreciate" gorm:"column:appreciate"`
}

func (AmountSoldOrder) TableName() string { return "todo_items" }

func (AppreciateOrder) TableName() string { return "todo_items" }
func (Order) TableName() string           { return "order_by_user" }
func (CreateOrder) TableName() string     { return "order_by_user" }
func (UpdateOrder) TableName() string     { return "order_by_user" }
