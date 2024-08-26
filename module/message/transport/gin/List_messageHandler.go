package ginMessage

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"main.go/module/message/storage"
	"net/http"
	"strconv"
)

func ListMessage(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		receiverId, err := strconv.Atoi(c.Query("receiver_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		senderId := c.MustGet(common.Current_user).(common.Requester).GetUserId()
		store := storage.NewSqlModel(db)
		message, err := store.ListMessage(c.Request.Context(), senderId, receiverId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": message})

	}
}
