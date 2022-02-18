package main

import (
	"wetime-go/config"
	"wetime-go/db"
	. "wetime-go/router"
)

func main() {

	// fmt.Println("hello world")

	// r := gin.Default()
	// r.GET("/hello", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{"msg": "hello"})
	// })
	// //r.Run(":8000")

	// 配置文件
	config.Process()

	// 链接数据库
	db.Init()

	// 初始化路由
	router := InitRouter()
	router.Run(":" + config.Conf.HttpPort)

}
