package ginUser

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"main.go/module/user/biz"
	"main.go/module/user/storage"
	"net/http"
)

func DeletedAccount(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		userId := c.MustGet(common.Current_user).(common.Requester).GetUserId()
		store := storage.NewSqlModel(db)
		business := biz.NewUserBiz(store)
		if err := business.NewDeletedAccount(c.Request.Context(), userId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
