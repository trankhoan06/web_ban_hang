package model

import (
	modelUser "main.go/module/user/model"
	"time"
)

type StatusMessage int

const (
	StatusDeleted StatusMessage = iota
	StatusDoing
	StatusRemove
)

type Message struct {
	Id              int           `json:"id" gorm:"column:id"`
	SenderId        int           `json:"sender_id" gorm:"column:sender_id"`
	ReceiverId      int           `json:"receiver_id" gorm:"column:receiver_id"`
	Message         string        `json:"message" gorm:"column:message"`
	IsStatusSender  StatusMessage `json:"is_status_sender" gorm:"column:is_status_sender"`
	IsStatusReceive StatusMessage `json:"is_status_receive" gorm:"column:is_status_receive"`
	CreateAt        time.Time     `json:"create_at" gorm:"column:create_at"`
	UpdateAt        time.Time     `json:"update_at" gorm:"column:update_at"`
}
type CreateMessage struct {
	Id        *int   `json:"id" gorm:"column:id"`
	SenderId  int    `json:"sender_id" gorm:"column:sender_id"`
	ReceiveId int    `json:"receiver_id" gorm:"column:receiver_id"`
	Message   string `json:"message" gorm:"column:message"`
}
type UserMessage struct {
	SenderId        int                   `json:"sender_id" gorm:"column:sender_id"`
	ReceiverId      int                   `json:"receiver_id" gorm:"column:receiver_id"`
	OwnerReceiverId *modelUser.SimpleUser `json:"owner_receiver_id" gorm:"foreignkey:ReceiverId;references:UserId"`
}

func (CreateMessage) TableName() string { return "message" }
func (UserMessage) TableName() string   { return "message" }
func (Message) TableName() string       { return "message" }
