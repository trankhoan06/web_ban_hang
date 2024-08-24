package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"main.go/component/middleware"
	"main.go/component/tokenProvider/jwt"
	gincart "main.go/module/cart/transport/gin"
	ginComment "main.go/module/comment/transport/gin"
	ginFollow "main.go/module/follow/transport/gin"
	"main.go/module/item/transport/gin"
	ginNotify "main.go/module/notify/transport/gin"
	ginOrder "main.go/module/order_by_user/transport/gin"
	ginSearch "main.go/module/search_item/transport/gin"
	"main.go/module/upload"
	"main.go/module/user/storage"
	ginUser2 "main.go/module/user/transport/gin"
	ginUserlikeItem "main.go/module/userlikeitem/transport/gin"
	"os"
)

func main() {

	dsn := os.Getenv("DOMAIN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	token := jwt.NewJwtProvider("jwt", "Khoandz123@")
	author := storage.NewSqlModel(db)
	middlewareAuthor := middleware.RequesMiddleware(author, token)
	r := gin.Default()
	r.Static("/static", "./static")
	v1 := r.Group("/v1")
	{
		v1.PATCH("/updateUser", middlewareAuthor, ginUser2.UpdateUser(db))
		v1.GET("/profile", middlewareAuthor, ginUser2.Profile(db))
		v1.PUT("/upload", upload.Upload_image(db))
		v1.POST("/register", ginUser2.CreateUser(db))
		v1.PATCH("/verify_email", ginUser2.VerifyEmailAccount(db))
		v1.POST("/login", ginUser2.LoginUser(db, token))
		v1.PATCH("/changepass", middlewareAuthor, ginUser2.ChangePassword(db))
		v1.POST("/forgotPassword", ginUser2.ForgotPassword(db))
		v1.GET("/forgotPassword/verifyCode", ginUser2.VerifyCodeForgotPassword(db))
		v1.PATCH("/update_password_forgot", ginUser2.UpdatePasswordForgot(db))
		v1.DELETE("/deleted_user", middlewareAuthor, ginUser2.DeletedUser(db))
		v1.GET("/user", ginUser2.GetUser(db))
		cart := v1.Group("/cart", middlewareAuthor)
		{
			cart.POST("", gincart.AddItem(db))
			cart.GET("/:itemId", gincart.GetItem(db))
			cart.GET("/list_item", gincart.ListItemCart(db))
			cart.PATCH("/update/:id", gincart.UpdateItemCart(db))
			cart.DELETE("/deleted/:id", gincart.DeletedItemCart(db))
		}
		comment := v1.Group("/comment", middlewareAuthor)
		{
			comment.POST("/:item_id", ginComment.CreateComment(db))
			comment.POST("/email/:email", ginComment.SendEmail(db))
			comment.GET("/:item_id/:id", ginComment.GetListParentAndChild(db))
			comment.GET("/list_comment/:item_id", ginComment.ListComment(db))
			comment.PATCH("/update_comment/:item_id/:id", ginComment.UpdateComment(db))
			comment.GET("/old_comment/:item_id/:id", ginComment.GetOldComment(db))
			comment.DELETE("/deleted_comment/:item_id/:id", ginComment.DeletedComment(db))
		}
		order := v1.Group("/order", middlewareAuthor)
		{
			order.POST("/create_order", ginOrder.CreateOrder(db))
			order.POST("/appreciate", ginOrder.AppreciateItem(db))
			order.GET("", ginOrder.GetOrder(db))
			order.GET("/get_order_sell", ginOrder.GetOrderSell(db))
			order.GET("list_order", ginOrder.ListOrder(db))
			order.GET("list_order_cancel_and_done", ginOrder.ListOrderCancelAndDone(db))
			order.GET("list_order_cancel_and_done_sell", ginOrder.ListOrderCancelAndDoneSell(db))
			order.GET("list_order_sell", ginOrder.ListOrderSell(db))
			order.PATCH("/update", ginOrder.UpdateOrder(db))
			order.PATCH("/update_status_sell", ginOrder.UpdateStatusSell(db))
			order.DELETE("/cancel", ginOrder.CancelOrder(db))
		}
		item := v1.Group("/items", middlewareAuthor)
		{
			item.GET("/:id", ginItem.Getitem(db))
			item.GET("/own/:id", ginItem.GetOwnitem(db))
			item.PATCH("/:id", ginItem.UpdateItem(db))
			item.POST("", ginItem.CreateItem(db))
			item.GET("", ginItem.ListItem(db))
			item.GET("/own", ginItem.ListOwnItem(db))
			item.DELETE("/:id", ginItem.DeletedItem(db))

			item.POST("/:id/like", ginUserlikeItem.CreateLikeItem(db))
			item.DELETE("/:id/unlike", ginUserlikeItem.UnLikeItem(db))
			item.GET("/:id/listlikeitem", ginUserlikeItem.ListLike(db))
		}
		search := v1.Group("/search", middlewareAuthor)
		{
			search.GET("/result", ginSearch.SearchItem(db))
			search.GET("/list_keyword", ginSearch.ListKeyword(db))
			search.DELETE("/deleted_keyword", ginSearch.DeletedKeyword(db))
			search.DELETE("/deleted_all_keyword", ginSearch.DeletedAllKeyword(db))
		}
		follow := v1.Group("/follow", middlewareAuthor)
		{
			follow.POST("/create", ginFollow.CreateFollow(db))
			follow.GET("/list_follow", ginFollow.ListUserFollow(db))
			follow.GET("/list_follow_user", ginFollow.ListFollowUser(db))
			follow.GET("/amount_follow_user", ginFollow.GetAmountFollowUser(db))
			follow.GET("/amount_user_follow", ginFollow.GetAmountUserFollow(db))
			follow.DELETE("/Unfollow", ginFollow.UnFollow(db))
		}
		notify := v1.Group("/notify", middlewareAuthor)
		{
			notify.POST("/send_notify", ginNotify.SendNotify(db))
			notify.POST("/send_all_notify", ginNotify.SendAllNotify(db))
			notify.GET("/list_notify", ginNotify.ListNotify(db))
			notify.GET("/list_notify_of_creator", ginNotify.ListNotifyOfCreator(db))
			notify.PATCH("/read_notify", ginNotify.ReadNotify(db))
			notify.DELETE("/deleted_notify", ginNotify.DeletedNotify(db))
			notify.DELETE("/deleted_notify_of_creator", ginNotify.DeletedNotifyOfCreator(db))
		}
	}

	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
