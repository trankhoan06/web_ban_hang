package ginUser

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"main.go/module/user/biz"
	"main.go/module/user/storage"
	"net/http"
)

func RegisterRole(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		role := c.Query("role")
		userId := c.MustGet(common.Current_user).(common.Requester).GetUserId()
		store := storage.NewSqlModel(db)
		business := biz.NewRegisterRoleBiz(store)
		if err := business.NewRegisterRole(c.Request.Context(), userId, role); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
