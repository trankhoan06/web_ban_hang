package ginUser

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/module/user/biz"
	"main.go/module/user/storage"
	"net/http"
)

func ForgotPassword(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		email := c.Query("email")
		store := storage.NewSqlModel(db)
		business := biz.NewUserBiz(store)
		verify, err := business.NewForgotPassWord(c.Request.Context(), email)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": verify})
	}
}
