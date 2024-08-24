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

func CreateUser(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.CreateUser
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		expire := time.Now().UTC().Add(1 * time.Minute)
		code := common.GenerateRandomCode()
		codeEmail := &model.CreateSendCode{
			Email:    data.Email,
			Code:     code,
			ExpireAt: expire,
			CreateAt: time.Now(),
		}
		store := storage.NewSqlModel(db)
		hash := common.NewSha256Hash()
		business := biz.NewRegisterbiz(store, hash)
		if err := business.NewRegisterUser(c.Request.Context(), &data, codeEmail); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		title := "Verify Email"
		content := fmt.Sprintf(`
<h1>Hello</h1>
<p>Web send code: %d</p>
`, code)
		to := []string{data.Email}
		err := emailSend.SendEmail(title, content, to, nil, nil, nil)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
