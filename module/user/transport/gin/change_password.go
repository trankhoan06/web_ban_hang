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

func ChangePassword(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var changePassword model.ChangePassWord
		if err := c.ShouldBind(&changePassword); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
			return
		}
		userId := c.MustGet(common.Current_user).(common.Requester).GetUserId()
		store := storage.NewSqlModel(db)
		hash := common.NewSha256Hash()
		business := biz.NewRegisterBiz(store, hash)
		if err := business.NewChangePassword(c.Request.Context(), userId, changePassword.Password, changePassword.NewPassword); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
