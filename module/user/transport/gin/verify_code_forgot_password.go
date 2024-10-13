package ginUser

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/module/user/biz"
	"main.go/module/user/model"
	"main.go/module/user/storage"
	"net/http"
)

func VerifyCodeForgotPassword(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var verifyForgot model.VerifySendEmail
		if err := c.ShouldBindJSON(&verifyForgot); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := storage.NewSqlModel(db)
		business := biz.NewUserBiz(store)
		if err := business.NewVerifyCodeForgotPassword(c.Request.Context(), verifyForgot.Code, verifyForgot.Token, verifyForgot.UserId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
