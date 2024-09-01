package ginVoucher

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"main.go/module/voucher/biz"
	"main.go/module/voucher/model"
	"main.go/module/voucher/storage"
	"net/http"
	"strconv"
)

func UpdateVoucher(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		voucherId, err := strconv.Atoi(c.Query("voucher_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var data model.UpdateVoucher
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		data.VendorId = c.MustGet(common.Current_user).(common.Requester).GetUserId()
		store := storage.NewSqlModel(db)
		business := biz.NewVoucherBiz(store)
		if err := business.NewUpdateVoucher(c.Request.Context(), voucherId, &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}
