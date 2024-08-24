package ginUserlikeItem

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	storageNotify "main.go/module/notify/storage"
	"main.go/module/userlikeitem/biz"
	"main.go/module/userlikeitem/storage"
	"net/http"
	"strconv"
)

func UnLikeItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		request := c.MustGet(common.Current_user).(common.Requester)

		store := storage.NewSqlModel(db)
		business := biz.NewUserUnlikedBiz(store)
		if err := business.NewDeleteUserUnlike(c.Request.Context(), id, request.GetUserId()); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		storeNotify := storageNotify.NewSQLModel(db)
		errItem := storeNotify.DeletedNotifyLikeItem(c.Request.Context(), request.GetUserId(), id)
		if errItem != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
