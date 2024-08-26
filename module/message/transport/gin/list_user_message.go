package ginMessage

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"main.go/module/message/storage"
	"net/http"
)

func ListUserMessage(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		senderId := c.MustGet(common.Current_user).(common.Requester).GetUserId()
		store := storage.NewSqlModel(db)
		result, errMess := store.ListUserMessage(c.Request.Context(), senderId, "OwnerReceiverId")
		if errMess != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errMess.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": result})

	}
}
