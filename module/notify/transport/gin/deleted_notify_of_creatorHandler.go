package ginNotify

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"main.go/module/notify/biz"
	"main.go/module/notify/model"
	"main.go/module/notify/storage"
	"net/http"
)

func DeletedNotifyOfCreator(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var deletedNotify model.CreatorNotify

		if err := c.ShouldBind(&deletedNotify); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		CreatorId := c.MustGet(common.Current_user).(common.Requester).GetUserId()
		store := storage.NewSQLModel(db)
		business := biz.NewNotifyBiz(store)
		if err := business.DeletedNotifyOfCreator(c.Request.Context(), CreatorId, deletedNotify.Message, deletedNotify.CreateAt); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
