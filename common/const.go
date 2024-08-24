package common

import (
	"fmt"
	modelComment "main.go/module/comment/model"
)

func Recovery() {
	if err := recover(); err != nil {
		fmt.Println("recovered:", err)
	}
}

type Payload struct {
	UId   int    `json:"user_id"`
	URole string `json:"role"`
}

func (p *Payload) GetUser() int {
	return p.UId
}
func (p *Payload) GetRole() string {
	return p.URole
}

const Current_user = "current_user"

type Requester interface {
	GetUserId() int
	GetRole() string
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
