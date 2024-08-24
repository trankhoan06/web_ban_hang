package ginUser

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"main.go/component/tokenProvider"
	emailSend "main.go/email"
	"main.go/module/user/biz"
	"main.go/module/user/model"
	"main.go/module/user/storage"
	"net/http"
	"time"
)

func LoginUser(db *gorm.DB, provider tokenProvider.Provider) func(*gin.Context) {
	return func(c *gin.Context) {
		var login model.LoginUser
		if err := c.ShouldBind(&login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		expire := time.Now().UTC().Add(1 * time.Minute)
		store := storage.NewSqlModel(db)
		code := common.GenerateRandomCode()
		codeEmail := &model.CreateSendCode{
			Email:    login.Email,
			Code:     code,
			ExpireAt: expire,
			CreateAt: time.Now(),
		}
		hash := common.NewSha256Hash()
		business := biz.NewLoginUserBiz(store, hash, provider, 60*60*7*24)
		token, err := business.LoginUser(c.Request.Context(), &login, codeEmail)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		title := "Verify Email"
		content := fmt.Sprintf(`
<h1>Hello</h1>
<p>Web send code: %d</p>
`, code)
		to := []string{login.Email}
		err = emailSend.SendEmail(title, content, to, nil, nil, nil)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token})

	}
}
