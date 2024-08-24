package ginNotify

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"main.go/module/notify/storage"
	"net/http"
	"strconv"
)

func DeletedNotify(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		userID := c.MustGet(common.Current_user).(common.Requester).GetUserId()
		store := storage.NewSQLModel(db)
		if err := store.DeletedNotify(c.Request.Context(), id, userID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": true})
		return
	}
}
