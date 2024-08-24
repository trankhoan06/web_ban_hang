package model

import (
	"main.go/module/user/model"
	"time"
)

type Follower struct {
	Id       int               `json:"id" gorm:"column:id"`
	UserId   int               `json:"user_id" gorm:"column:user_id"`
	ByUserId int               `json:"by_user_id" gorm:"column:by_user_id"`
	Owner    *model.SimpleUser `json:"owner" gorm:"foreignkey:UserId;references:UserId"`
	CreateAt time.Time         `json:"create_at" gorm:"column:create_at"`
}
type CreateFollower struct {
	UserId   int `json:"user_id" gorm:"column:user_id"`
	ByUserId int `json:"by_user_id" gorm:"column:by_user_id"`
}

func (Follower) TableName() string       { return "follow" }
func (CreateFollower) TableName() string { return "follow" }
