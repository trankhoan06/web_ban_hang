package ginUser

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/module/user/biz"
	"main.go/module/user/storage"
	"net/http"
	"strconv"
)

func VerifyEmailAccount(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		email := c.Query("email")
		if email == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "email required"})
			return
		}
		code, err := strconv.Atoi(c.Query("code"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "code required"})
			return
		}
		store := storage.NewSqlModel(db)
		busines := biz.NewVerifyEmailAccountBiz(store)
		if err := busines.NewVerifyEmailAccount(c.Request.Context(), email, code); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
