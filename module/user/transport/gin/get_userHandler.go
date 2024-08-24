package ginUser

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/module/user/biz"
	"main.go/module/user/storage"
	"net/http"
	"strconv"
)

func GetUser(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		userId, err := strconv.Atoi(c.Query("user_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := storage.NewSqlModel(db)
		busines := biz.NewGetUserBiz(store)
		user, err := busines.NewGetUser(c.Request.Context(), userId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"user": user})
	}
}
