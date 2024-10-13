package common

import (
	modelComment "main.go/module/comment/model"
	"main.go/module/user/model"
)

type Payload struct {
	UId   int             `json:"user_id"`
	URole *model.RoleUser `json:"role"`
}

func (p *Payload) GetUser() int {
	return p.UId
}
func (p *Payload) GetRole() *model.RoleUser {
	return p.URole
}

const Current_user = "current_user"

type Requester interface {
	GetUserId() int
	GetRole() *model.RoleUser
	GetEmail() string
}
type TreeComment struct {
	Val   interface{}
	Child []*TreeComment
}

func NewNode(value modelComment.CommentUser) *TreeComment {
	return &TreeComment{
		Val:   value,
		Child: []*TreeComment{},
	}
}
