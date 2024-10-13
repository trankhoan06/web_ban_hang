package model

import "time"

type SendCode struct {
	UserId   int       `json:"user_id" gorm:"column:user_id"`
	Email    string    `json:"email" gorm:"column:email"`
	Code     int       `json:"code" gorm:"column:code"`
	Token    string    `json:"token" gorm:"column:token"`
	Verify   bool      `json:"verify" gorm:"column:verify"`
	CreateAt time.Time `json:"create_at" gorm:"column:create_at"`
	Expire   time.Time `json:"expire_at" gorm:"column:expire_at"`
}
type CreateSendCode struct {
	UserId int       `json:"user_id" gorm:"column:user_id"`
	Email  string    `json:"email" gorm:"column:email"`
	Code   int       `json:"code" gorm:"column:code"`
	Token  string    `json:"token" gorm:"column:token"`
	Expire time.Time `json:"expire_at" gorm:"column:expire_at"`
}
type TokenSendEmail struct {
	Id      int    `json:"id" `
	Token   string `json:"token"`
	IsEmail bool   `json:"is_email"`
}
type VerifySendEmail struct {
	Token  string `json:"token"`
	Code   int    `json:"code"`
	UserId int    `json:"user_id"`
}
type ChangePasswordForgot struct {
	Token    string `json:"token"`
	UserId   int    `json:"user_id"`
	Password string `json:"password"`
}

func (SendCode) TableName() string       { return "send_code_email" }
func (CreateSendCode) TableName() string { return "send_code_email" }
