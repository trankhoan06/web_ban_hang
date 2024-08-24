package model

import (
	"time"
)

type StatusUser int

const (
	StatusUserInactive StatusUser = iota
	StatusUserActive
)

type User struct {
	UserId   int        `json:"user_id" gorm:"column:id"`
	Email    string     `json:"email" gorm:"column:email"`
	Salt     string     `json:"salt" gorm:"column:salt"`
	PassWord string     `json:"password" gorm:"column:password"`
	FirtName string     `json:"firt_name" gorm:"column:firt_name"`
	LastName string     `json:"last_name" gorm:"column:last_name"`
	Phone    string     `json:"phone" gorm:"column:phone"`
	Role     string     `json:"role" gorm:"column:role"`
	IsEmail  *bool      `json:"-" gorm:"column:is_email"`
	CreateAt *time.Time `json:"create_at" gorm:"column:create_at"`
	UpdateAt *time.Time `json:"update_at" gorm:"column:update_at"`
	Status   StatusUser `json:"status" gorm:"column:status"`
}

type SimpleUser struct {
	UserId   int    `json:"user_id" gorm:"column:id"`
	FirtName string `json:"firt_name" gorm:"column:firt_name"`
	LastName string `json:"last_name" gorm:"column:last_name"`
	Status   int    `json:"status" gorm:"column:status"`
}
type UpdateUser struct {
	UserId   int    `json:"-" gorm:"column:id"`
	FirtName string `json:"firt_name" gorm:"column:firt_name"`
	LastName string `json:"last_name" gorm:"column:last_name"`
	Phone    string `json:"phone" gorm:"column:phone"`
	Role     string `json:"role" gorm:"column:role"`
}
type LIstUserId struct {
	UserId int `json:"id" gorm:"column:id"`
}

func (u *User) GetUserId() int {
	return u.UserId
}
func (u *User) GetRole() string {
	return u.Role
}
func (u *User) GetEmail() string {
	return u.Email
}

type CreateUser struct {
	UserId   int    `json:"user_id" gorm:"column:id"`
	Email    string `json:"email" gorm:"column:email"`
	PassWord string `json:"password" gorm:"column:password"`
	Salt     string `json:"-" gorm:"column:salt"`
	Role     string `json:"-" gorm:"column:role"`
}
type UpdatePasswordForgot struct {
	PassWord string `json:"password" gorm:"column:password"`
}
type LoginUser struct {
	Email    string `json:"email" gorm:"column:email"`
	PassWord string `json:"password" gorm:"column:password"`
}
type UpdatePass struct {
	Email       string `json:"email" gorm:"column:email"`
	PassWord    string `json:"password" gorm:"column:password"`
	NewPassWord string `json:"new_password"`
}

func (UpdatePass) TableName() string           { return "users" }
func (LIstUserId) TableName() string           { return "users" }
func (UpdatePasswordForgot) TableName() string { return "users" }
func (User) TableName() string                 { return "users" }
func (CreateUser) TableName() string           { return "users" }
func (SimpleUser) TableName() string           { return "users" }
func (UpdateUser) TableName() string           { return "users" }
func (LoginUser) TableName() string            { return "users" }
