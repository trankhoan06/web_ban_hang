package ginVoucher

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"main.go/module/voucher/biz"
	"main.go/module/voucher/storage"
	"net/http"
)

func ListVoucherVendor(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		userId := c.MustGet(common.Current_user).(common.Requester).GetUserId()
		store := storage.NewSqlModel(db)
		business := biz.NewVoucherBiz(store)
		voucher, err := business.NewListVoucherVendor(c.Request.Context(), userId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": voucher})
	}
}
