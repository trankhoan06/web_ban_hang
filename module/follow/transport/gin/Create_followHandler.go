package ginFollow

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"main.go/module/follow/biz"
	"main.go/module/follow/model"
	"main.go/module/follow/storage"
	storageUser "main.go/module/user/storage"
	"net/http"
	"strconv"
)

func CreateFollow(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		userId, err := strconv.Atoi(c.Query("user_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := storage.NewSQLModel(db)
		store1 := storageUser.NewSqlModel(db)
		business := biz.NewFollowUserBiz(store, store1)
		var user model.CreateFollower
		user.ByUserId = c.MustGet(common.Current_user).(common.Requester).GetUserId()
		user.UserId = userId
		if user.UserId == user.ByUserId {
			c.JSON(http.StatusBadRequest, gin.H{"error": "you can't follow yourself"})
			return
		}
		if err := business.CreateFollowUser(c.Request.Context(), &user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
