package jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"main.go/common"
	"main.go/component/tokenProvider"
	"time"
)

type JwtProvider struct {
	Prefix string
	Secret string
}

func NewJwtProvider(prefix string, secret string) *JwtProvider {
	return &JwtProvider{Prefix: prefix, Secret: secret}
}

type MyClaims struct {
	Payload common.Payload `json:"payload"`
	jwt.StandardClaims
}
type token struct {
	Token   string    `json:"token"`
	Expiry  int       `json:"expiry"`
	Created time.Time `json:"created"`
}

func (t *token) Gettoken() string {
	return t.Token
}

func (j *JwtProvider) GetSecret() string {
	return j.Secret
}
func (j *JwtProvider) Generate(data tokenProvider.Payload, expiry int) (tokenProvider.Token, error) {
	now := time.Now()
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, &MyClaims{
		common.Payload{
			URole: data.GetRole(),
			UId:   data.GetUser(),
		},
		jwt.StandardClaims{
			ExpiresAt: now.Add(time.Duration(expiry) * time.Second).Unix(),
			IssuedAt:  now.Unix(),
			Id:        fmt.Sprint(now.UnixNano()),
		},
	})
	myToken, err := t.SignedString([]byte(j.Secret))

	if err != nil {
		return nil, err
	}
	return &token{
		Token:   myToken,
		Expiry:  expiry,
		Created: now,
	}, nil
}
func (j *JwtProvider) Validate(token string) (tokenProvider.Payload, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &MyClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	if !jwtToken.Valid {
		return nil, errors.New("invalid token")
	}
	claims, ok := jwtToken.Claims.(*MyClaims)
	if !ok {
		return nil, errors.New("invalid token")
	}
	return &claims.Payload, nil
}
