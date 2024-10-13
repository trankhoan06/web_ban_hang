package model

import "time"

type StatusUser int

const (
	StatusUserDeleted StatusUser = iota
	StatusUserDoing
)

type RoleUser int

const (
	RoleUserUser RoleUser = iota + 1
	RoleUserAdmin
)

type User struct {
	Id          int         `json:"id" gorm:"column:id"`
	Email       string      `json:"email" gorm:"column:email"`
	Salt        string      `json:"salt" gorm:"column:salt"`
	Password    string      `json:"password" gorm:"column:password"`
	FirstName   string      `json:"first_name" gorm:"column:first_name"`
	LastName    string      `json:"last_name" gorm:"column:last_name"`
	Description string      `json:"description" gorm:"column:description"`
	Phone       string      `json:"phone" gorm:"column:phone"`
	Role        *RoleUser   `json:"role" gorm:"column:role"`
	Status      *StatusUser `json:"status" gorm:"column:status"`
	IsEmail     bool        `json:"is_email" gorm:"column:is_email"`
	CreateAt    time.Time   `json:"create_at" gorm:"column:create_at"`
	UpdateAt    time.Time   `json:"update_at" gorm:"column:update_at"`
}

func (u *User) GetUserId() int {
	return u.Id
}
func (u *User) GetEmail() string {
	return u.Email
}
func (u *User) GetRole() *RoleUser {
	return u.Role
}

type CreateUser struct {
	Id          int         `json:"id" gorm:"column:id"`
	Email       string      `json:"email" gorm:"column:email"`
	Status      *StatusUser `json:"-" gorm:"column:status"`
	Salt        string      `json:"-" gorm:"column:salt"`
	Password    string      `json:"password" gorm:"column:password"`
	FirstName   string      `json:"first_name" gorm:"column:first_name"`
	LastName    string      `json:"last_name" gorm:"column:last_name"`
	Address     string      `json:"address" gorm:"column:address"`
	Description string      `json:"description" gorm:"column:description"`
	Phone       string      `json:"phone" gorm:"column:phone"`
}
type UpdateUser struct {
	Email       string  `json:"-" gorm:"column:email"`
	FirstName   *string `json:"first_name" gorm:"column:first_name"`
	LastName    *string `json:"last_name" gorm:"column:last_name"`
	Description *string `json:"description" gorm:"column:description"`
	Address     *string `json:"address" gorm:"column:address"`
	Phone       *string `json:"phone" gorm:"column:phone"`
}
type SimpleUser struct {
	Id        int    `json:"id" gorm:"column:id"`
	FirstName string `json:"first_name" gorm:"column:first_name"`
	LastName  string `json:"last_name" gorm:"column:last_name"`
}
type ListUserId struct {
	UserId int `json:"id" gorm:"column:id"`
}
type LoginUser struct {
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"password" gorm:"column:password"`
}
type ChangePassWord struct {
	Password    string `json:"password" gorm:"column:password"`
	NewPassword string `json:"new_password" gorm:"column:new_password"`
}

func (User) TableName() string       { return "users" }
func (CreateUser) TableName() string { return "users" }
func (UpdateUser) TableName() string { return "users" }
func (SimpleUser) TableName() string { return "users" }
