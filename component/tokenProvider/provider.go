package tokenProvider

type Provider interface {
	Generate(payload Payload, expiry int) (Token, error)
	Validate(token string) (Payload, error)
	GetSecret() string
}
type Payload interface {
	GetUser() int
	GetRole() string
}
type Token interface {
	Gettoken() string
}
