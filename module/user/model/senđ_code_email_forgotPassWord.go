package model

import "time"

type StatusCode int

type SendCode struct {
	Id       int       `json:"-" gorm:"column:id"`
	Email    string    `json:"email" gorm:"column:email"`
	Code     int       `json:"code" gorm:"column:code"`
	Token    string    `json:"token" gorm:"column:token"`
	CreateAt time.Time `json:"create_at" gorm:"column:create_at"`
	ExpireAt time.Time `json:"expire_at" gorm:"column:expire_at"`
}
type CreateSendCode struct {
	Email    string    `json:"email" gorm:"column:email"`
	Code     int       `json:"-" gorm:"column:code"`
	CreateAt time.Time `json:"create_at" gorm:"column:create_at"`
	ExpireAt time.Time `json:"expire_at" gorm:"column:expire_at"`
}

func (SendCode) TableName() string       { return "send_code_email" }
func (CreateSendCode) TableName() string { return "send_code_email" }
