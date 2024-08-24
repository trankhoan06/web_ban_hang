package ginUser

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"main.go/module/user/biz"
	"main.go/module/user/model"
	"main.go/module/user/storage"
	"net/http"
)

func UpdateUser(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.UpdateUser
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		request := c.MustGet(common.Current_user).(common.Requester)
		store := storage.NewSqlModel(db)
		business := biz.NewUpdateUserBiz(store)
		if err := business.NewUpdateUser(c.Request.Context(), &data, request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"update": true})

	}
}
