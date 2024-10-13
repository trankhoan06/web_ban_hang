package tokenProvider

import "main.go/module/user/model"

type Provider interface {
	Generate(payload Payload, expiry int) (Token, error)
	Validate(token string) (Payload, error)
	GetSecret() string
}
type Payload interface {
	GetUser() int
	GetRole() *model.RoleUser
}
type Token interface {
	Gettoken() string
}
