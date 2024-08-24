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
		var data model.UpdatePass
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		hash := common.NewSha256Hash()
		store := storage.NewSqlModel(db)
		requester := c.MustGet(common.Current_user).(common.Requester)
		bussines := biz.NewChangeBiz(store, hash)
		if err := bussines.NewChangePasswordBiz(c.Request.Context(), requester, &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": true})

	}
}
