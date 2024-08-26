package ginMessage

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"main.go/module/message/biz"
	"main.go/module/message/storage"
	"net/http"
	"strconv"
)

func DeletedUserMessage(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		receive, err := strconv.Atoi(c.Query("receiver_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		senderId := c.MustGet(common.Current_user).(common.Requester).GetUserId()
		store := storage.NewSqlModel(db)
		business := biz.NewMessageBiz(store)
		if err := business.NewDeletedUserMessage(c.Request.Context(), senderId, receive); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
