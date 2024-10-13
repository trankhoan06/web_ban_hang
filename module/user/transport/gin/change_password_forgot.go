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

func ChangePasswordForgot(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var token model.ChangePasswordForgot
		if err := c.ShouldBind(&token); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := storage.NewSqlModel(db)
		hash := common.NewSha256Hash()
		business := biz.NewRegisterBiz(store, hash)
		if err := business.ChangeForgotPassword(c.Request.Context(), token.Token, token.Password, token.UserId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
