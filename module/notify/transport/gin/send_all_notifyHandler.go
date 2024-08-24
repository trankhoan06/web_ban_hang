package ginNotify

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"main.go/module/notify/biz"
	"main.go/module/notify/model"
	"main.go/module/notify/storage"
	bizUser "main.go/module/user/biz"
	storageUser "main.go/module/user/storage"
	"net/http"
)

func SendAllNotify(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var notify model.CreateNotify
		if err := c.ShouldBind(&notify); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		notify.CreatorId = c.MustGet(common.Current_user).(common.Requester).GetUserId()
		notify.TypeMessage = model.TypeEvent
		storeUser := storageUser.NewSqlModel(db)
		store := storage.NewSQLModel(db)
		businessUser := bizUser.NewListAllUserIDBiz(storeUser)
		business := biz.NewNotifyUserBiz(store, storeUser)
		result, err := businessUser.NewListAllUserID(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := business.NewUserSendAllNotify(c.Request.Context(), result, &notify); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": result})

	}
}
