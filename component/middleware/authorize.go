package middleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"main.go/common"
	"main.go/component/tokenProvider"
	"main.go/module/user/model"
	"net/http"
	"strings"
)

type Authorize interface {
	FindUser(ctx context.Context, cond map[string]interface{}) (*model.User, error)
}

func Extractoken(s string) (string, error) {
	str := strings.Split(s, " ")
	if str[0] != "Bearer" || len(str) < 2 || strings.TrimSpace(str[1]) == "" {
		return "", errors.New("token has been fault")
	}
	return str[1], nil
}

func RequesMiddleware(authorize Authorize, provider tokenProvider.Provider) func(*gin.Context) {
	return func(c *gin.Context) {
		s, err := Extractoken(c.GetHeader("Authorization"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		payLoad, err := provider.Validate(s)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user, err := authorize.FindUser(c.Request.Context(), map[string]interface{}{"id": payLoad.GetUser()})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if user.Status == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprint("user has been deleted of banned"),
			})
			return
		}
		c.Set(common.Current_user, user)
		c.Next()

	}
}
