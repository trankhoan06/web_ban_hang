package ginUser

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"main.go/module/user/biz"
	"main.go/module/user/model"
	"main.go/module/user/storage"
	"net/http"
)

func UpdatePasswordForgot(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		token := c.Query("token")
		if token == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "token required"})
			return
		}
		var password model.UpdatePasswordForgot
		if err := c.ShouldBindJSON(&password); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := storage.NewSqlModel(db)
		hash := common.NewSha256Hash()
		business := biz.NewUpdatePasswordForgotBiz(store, hash)
		if err := business.NewUpdatePasswordForgot(c.Request.Context(), token, password.PassWord); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
