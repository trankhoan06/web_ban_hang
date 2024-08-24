package model

import (
	"errors"
	"main.go/module/user/model"
	"time"
)

type LikeItem struct {
	UserId   int        `json:"user_id" gorm:"column:user_id"`
	ItemId   int        `json:"item_id" gorm:"column:item_id"`
	CreateAt *time.Time `json:"create_at" gorm:"column:create_at"`
}
type UserLikeItem struct {
	UserId int               `json:"-" gorm:"column:user_id"`
	Owner  *model.SimpleUser `json:"user" gorm:"foreignkey:UserId"`
}

func (LikeItem) TableName() string     { return "userlikeitem" }
func (UserLikeItem) TableName() string { return "userlikeitem" }

var (
	ErrUserLike = errors.New("user has been like item")
)
