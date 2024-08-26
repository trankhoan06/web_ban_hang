package ginMessage

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"main.go/module/message/model"
	"main.go/module/message/storage"
	"net/http"
	"strconv"
	"strings"
)

func CreateMessage(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		receiverId, err := strconv.Atoi(c.Query("receiver_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var message model.CreateMessage
		if err := c.ShouldBind(&message); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		message.Message = strings.TrimSpace(message.Message)
		message.ReceiveId = receiverId
		message.SenderId = c.MustGet(common.Current_user).(common.Requester).GetUserId()
		store := storage.NewSqlModel(db)
		if err := store.CreateMessage(c.Request.Context(), &message); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": message})
	}
}
