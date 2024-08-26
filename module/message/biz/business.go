package biz

import (
	"context"
	"main.go/module/message/model"
)

type MessageStorage interface {
	ListMessage(ctx context.Context, sender, receiver int) (*[]model.Message, error)
	DeletedMessage(ctx context.Context, id int, column string) error
	CreateMessage(ctx context.Context, message *model.CreateMessage) error
	ListUserMessage(ctx context.Context, sender int, moreKey ...string) (*[]model.UserMessage, error)
	GetMessage(ctx context.Context, id int) (*model.Message, error)
	DeletedUserMessage(ctx context.Context, sender, receive int) error
	FindUserMessage(ctx context.Context, sender int, receive int, moreKey ...string) (*model.UserMessage, error)
}
type MessageBiz struct {
	store MessageStorage
}

func NewMessageBiz(store MessageStorage) *MessageBiz {
	return &MessageBiz{store: store}
}
