package middleware

import (
	"main.go/component/tokenProvider"
	"main.go/module/user/storage"
)

type MiddlewareManager struct {
	token  tokenProvider.Provider
	authen *storage.SqlModel
}

func NewMiddlewareManager(token tokenProvider.Provider, authen *storage.SqlModel) *MiddlewareManager {
	return &MiddlewareManager{token: token, authen: authen}
}
