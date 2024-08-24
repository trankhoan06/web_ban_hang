package ginNotify

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	bizFollow "main.go/module/follow/biz"
	storageFollow "main.go/module/follow/storage"
	"main.go/module/notify/biz"
	"main.go/module/notify/model"
	"main.go/module/notify/storage"
	"net/http"
)

func SendNotify(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.CreateNotify
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		storeFollow := storageFollow.NewSQLModel(db)
		businessFollow := bizFollow.NewFollowBiz(storeFollow)
		userId := c.MustGet(common.Current_user).(common.Requester).GetUserId()
		user, err := businessFollow.ListUserFollow(c.Request.Context(), userId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		data.CreatorId = c.MustGet(common.Current_user).(common.Requester).GetUserId()
		data.TypeMessage = model.TypeEvent
		storeNotify := storage.NewSQLModel(db)
		businessNotify := biz.NewNotifyBiz(storeNotify)
		if err := businessNotify.SendNotify(c.Request.Context(), &data, user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": user})
	}
}
