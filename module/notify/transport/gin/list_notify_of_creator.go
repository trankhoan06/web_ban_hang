package ginNotify

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"main.go/module/notify/storage"
	"net/http"
)

func ListNotifyOfCreator(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		creatorId := c.MustGet(common.Current_user).(common.Requester).GetUserId()
		store := storage.NewSQLModel(db)
		result, err := store.ListNotifyOfCreator(c.Request.Context(), creatorId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": result})
	}
}
