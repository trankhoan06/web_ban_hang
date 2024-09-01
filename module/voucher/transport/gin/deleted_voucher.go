package ginVoucher

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"main.go/module/voucher/biz"
	"main.go/module/voucher/storage"
	"net/http"
	"strconv"
)

func DeletedVoucher(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		voucherId, err := strconv.Atoi(c.Query("voucher_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		userId := c.MustGet(common.Current_user).(common.Requester).GetUserId()

		store := storage.NewSqlModel(db)
		business := biz.NewVoucherBiz(store)
		if err := business.NewDeletedVoucher(c.Request.Context(), voucherId, userId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
