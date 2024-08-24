package ginFollow

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/module/follow/biz"
	"main.go/module/follow/storage"
	"net/http"
	"strconv"
)

func GetAmountUserFollow(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		userId, err := strconv.Atoi(c.Query("user_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := storage.NewSQLModel(db)
		business := biz.NewFollowBiz(store)
		listFollow, err := business.NewGetAmountFollow(c.Request.Context(), userId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"listFollow": listFollow})
	}
}
