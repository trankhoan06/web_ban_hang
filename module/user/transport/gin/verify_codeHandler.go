package ginUser

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/module/user/biz"
	"main.go/module/user/storage"
	"net/http"
	"strconv"
)

func VerifyCodeForgotPassword(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		code, err := strconv.Atoi(c.Query("code"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		email := c.Query("email")
		store := storage.NewSqlModel(db)
		busines := biz.NewVerifyCodeForgotPasswordBiz(store)
		if err := busines.NewVerifyCode(c.Request.Context(), code, email); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": true})

	}
}
