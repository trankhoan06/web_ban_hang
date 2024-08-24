package model

import "time"

type TypeMessage int

const (
	TypeEvent TypeMessage = iota + 1
	TypeItem
	TypeComment
	TypeLikeItem
	TypeOrder
	TypeAppreciate
)

type Notify struct {
	Id          int         `json:"id" gorm:"column:id"`
	UserId      int         `json:"user_id" gorm:"column:user_id"`
	Creator     int         `json:"creator" gorm:"column:creator"`
	Message     string      `json:"message" gorm:"column:message"`
	TypeMessage TypeMessage `json:"type_message" gorm:"column:type_message"`
	IsRead      bool        `json:"is_read" gorm:"column:is_read"`
	CreateAt    time.Time   `json:"create_at" gorm:"column:create_at"`
	UpdateAt    time.Time   `json:"update_at" gorm:"column:update_at"`
}

type CreateNotify struct {
	UserId      int         `json:"user_id" gorm:"column:user_id"`
	ItemId      *int        `json:"item_id" gorm:"column:item_id"`
	TypeMessage TypeMessage `json:"type_message" gorm:"column:type_message"`
	CommentId   *int        `json:"comment_id" gorm:"column:comment_id"`
	Message     string      `json:"message" gorm:"column:message"`
	CreatorId   int         `json:"creator_id" gorm:"column:creator_id"`
}
type CreatorNotify struct {
	Message  string    `json:"message" gorm:"column:message"`
	CreateAt time.Time `json:"create_at" gorm:"column:create_at"`
}

func (Notify) TableName() string        { return "notify" }
func (CreatorNotify) TableName() string { return "notify" }
func (CreateNotify) TableName() string  { return "notify" }
