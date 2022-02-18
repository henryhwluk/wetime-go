package router

import (
	. "wetime-go/apis"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", IndexApi)

	router.POST("/user", AddUserApi)

	router.POST("/login", LoginUserApi)

	router.POST("/upload", UploadImage)

	router.POST("/post", PostFeedApi)

	router.GET("/user/:id", GetUserApi)

	router.POST("/follow/:id/user/:uid", FollowUserApi)

	router.GET("/follow/:id", GetFollowApi)

	router.GET("/post/:id", GetPostApi)

	router.GET("/token/:id", GetRCTokenApi)

	return router
}
