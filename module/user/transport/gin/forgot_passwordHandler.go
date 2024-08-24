package ginUser

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	emailSend "main.go/email"
	"main.go/module/user/biz"
	"main.go/module/user/model"
	"main.go/module/user/storage"
	"net/http"
	"time"
)

func ForgotPassword(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		email := c.Query("email")
		expire := time.Now().UTC().Add(1 * time.Minute)
		if email == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "email is required"})
			return
		}
		store := storage.NewSqlModel(db)
		code := common.GenerateRandomCode()
		busines := biz.NewFindUserBiz(store)
		codeEmail := &model.CreateSendCode{
			Email:    email,
			Code:     code,
			ExpireAt: expire,
			CreateAt: time.Now(),
		}
		err := busines.NewSendCode(c.Request.Context(), map[string]interface{}{"email": email}, codeEmail)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		title := "Forgot password"
		content := fmt.Sprintf(`
<h1>Hello</h1>
<p>Web send code: %d</p>
`, code)
		to := []string{email}
		err = emailSend.SendEmail(title, content, to, nil, nil, nil)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": true})

	}
}
