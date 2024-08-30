package model

import (
	"main.go/module/user/model"
	"time"
)

const NameItem string = "id"

type TodoList struct {
	Id          int               `json:"id" gorm:"column:id"`
	UserId      int               `json:"-" gorm:"column:user_id"`
	Owner       *model.SimpleUser `json:"owner" gorm:"foreignkey:UserId;references:UserId"`
	Title       string            `json:"title" gorm:"column:title"`
	Description string            `json:"description" gorm:"column:description"`
	Price       int               `json:"price" gorm:"column:price"`
	AmountItem  int               `json:"amount_item" gorm:"column:amount_item"`
	AmountSold  int               `json:"amount_sold" gorm:"column:amount_sold"`
	Image       *Image            `json:"image" gorm:"column:image"`
	Status      *StatusItem       `json:"status" gorm:"column:status"`
	CreatedAt   time.Time         `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   time.Time         `json:"updated_at" gorm:"column:updated_at"`
}
type TodoUpdateAmountItem struct {
	AmountItem int `json:"amount_item" gorm:"column:amount_item"`
}
type TodoCreateItem struct {
	UserId      int         `json:"-" gorm:"column:user_id"`
	Title       string      `json:"title" gorm:"column:title"`
	Price       int         `json:"price" gorm:"column:price"`
	Category    string      `json:"category" gorm:"column:category"`
	Description string      `json:"description" gorm:"column:description"`
	AmountItem  int         `json:"amount_item" gorm:"column:amount_item"`
	Status      *StatusItem `json:"status" gorm:"column:status"`
	Image       *Image      `json:"image" gorm:"column:image"`
}
type TodoUpdateItem struct {
	UserId      int         `json:"-" gorm:"column:user_id"`
	Title       *string     `json:"title" gorm:"column:title"`
	Description *string     `json:"description" gorm:"column:description"`
	AmountItem  *int        `json:"amount_item" gorm:"column:amount_item"`
	Status      *StatusItem `json:"status" gorm:"column:status"`
	Image       *Image      `json:"image" gorm:"column:image"`
}
type SimpleItem struct {
	Id          int    `json:"-" gorm:"column:id"`
	UserId      int    `json:"user_id" gorm:"column:user_id"`
	Title       string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
	Image       *Image `json:"image" gorm:"column:image"`
}

func (SimpleItem) TableName() string           { return "todo_items" }
func (TodoUpdateAmountItem) TableName() string { return "todo_items" }
func (TodoList) TableName() string             { return "todo_items" }
func (TodoCreateItem) TableName() string       { return "todo_items" }
func (TodoUpdateItem) TableName() string       { return "todo_items" }
