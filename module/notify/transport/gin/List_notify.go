package ginNotify

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"main.go/module/notify/biz"
	"main.go/module/notify/storage"
	"net/http"
)

func ListNotify(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		userID := c.MustGet(common.Current_user).(common.Requester).GetUserId()
		store := storage.NewSQLModel(db)
		business := biz.NewNotifyBiz(store)
		notifys, err := business.ListNotify(c.Request.Context(), userID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": notifys})
	}
}
