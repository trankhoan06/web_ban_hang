package model

type Filter struct {
	UserId  *int   `json:"-" gorm:"column:user_id" `
	Keyword string `json:"keyword" gorm:"column:content" form:"keyword"`
}
