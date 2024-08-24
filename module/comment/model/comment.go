package model

import (
	"main.go/module/user/model"
	"time"
)

type CommentStatus int

const (
	CommentStatusRemove CommentStatus = iota
	CommentStatusActive
)

type CommentUser struct {
	Id            int               `json:"id" gorm:"column:id"`
	UserId        int               `json:"-" gorm:"column:user_id"`
	ItemId        int               `json:"-" gorm:"column:item_id"`
	OwnerItem     int               `json:"owner_item" gorm:"column:owner_item"`
	OwnerItemUser *model.SimpleUser `json:"owner_item_user" gorm:"foreignkey:owner_item;references:UserId"`
	Owner         *model.SimpleUser `json:"owner" gorm:"foreignkey:UserId;references:UserId "`
	Content       string            `json:"content" gorm:"column:content"`
	Status        *CommentStatus    `json:"-" gorm:"column:status"`
	IsUpdate      *bool             `json:"is_update" gorm:"column:is_update"`
	//Reply    *CommentReply     `json:"reply" gorm:"foreignKey:ParentId;references:ParentId"`
	ParentId *int       `json:"-" gorm:"column:parent_id"`
	CreateAt *time.Time `json:"create_at" gorm:"column:create_at"`
	UpdateAt *time.Time `json:"update_at" gorm:"column:update_at"`
}

//	type CommentReply struct {
//		UserId   int        `json:"-" gorm:"column:user_id"`
//		ItemId   int        `json:"-" gorm:"column:item_id"`
//		Content  string     `json:"content" gorm:"column:content"`
//		ParentId *int       `json:"-" gorm:"column:parent_id"`
//		CreateAt *time.Time `json:"create_at" gorm:"column:create_at"`
//		UpdateAt *time.Time `json:"update_at" gorm:"column:update_at"`
//	}
type CreateComment struct {
	Id        *int   `json:"-" gorm:"column:id"`
	UserId    int    `json:"-" gorm:"column:user_id"`
	ItemId    int    `json:"-" gorm:"column:item_id"`
	OwnerItem int    `json:"-" gorm:"column:owner_item"`
	Content   string `json:"content" gorm:"column:content"`
	ParentId  *int   `json:"parent_id" gorm:"column:parent_id"`
}
type OldComment struct {
	UserId     int           `json:"-" gorm:"column:user_id"`
	ItemId     int           `json:"-" gorm:"column:item_id"`
	Content    string        `json:"content" gorm:"column:content"`
	Status     CommentStatus `json:"-" gorm:"column:status"`
	IsUpdate   bool          `json:"-" gorm:"column:is_update"`
	OriginalId int           `json:"-" gorm:"column:original_id"`
	ParentId   *int          `json:"-" gorm:"column:parent_id"`
	CreateAt   *time.Time    `json:"create_at" gorm:"column:create_at"`
	UpdateAt   *time.Time    `json:"update_at" gorm:"column:update_at"`
}

type UpdateComment struct {
	Id       int    `json:"-" gorm:"column:id"`
	UserId   int    `json:"-" gorm:"column:user_id"`
	ItemId   int    `json:"-" gorm:"column:item_id"`
	Content  string `json:"content" gorm:"column:content"`
	IsUpdate bool   `json:"-" gorm:"column:is_update"`
}

func (CommentUser) TableName() string { return "comment" }
func (OldComment) TableName() string  { return "comment" }

// func (CommentReply) TableName() string  { return "comment" }
func (CreateComment) TableName() string { return "comment" }
func (UpdateComment) TableName() string { return "comment" }
