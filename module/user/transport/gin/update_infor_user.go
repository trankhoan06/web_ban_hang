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

func UpdateInforUser(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.UpdateUser
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		data.Email = c.MustGet(common.Current_user).(common.Requester).GetEmail()
		store := storage.NewSqlModel(db)
		business := biz.NewUserBiz(store)
		if err := business.NewUpdateInformationUser(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}
