package router

import "github.com/gin-gonic/gin"

//路由总集
type RoutersGroup struct {
}

//实例化
var RoutersGroupApp = new(RoutersGroup)

func InitRouter() {
	router := gin.Default()
	// router.Use(middlewares.Cors())
	PrivateGroup := router.Group("")
	PrivateGroup.Use()
	{
		//规定您的静态资源目录
		router.Static("/assets", "./assets")
	}

	err := router.Run()
	if err != nil {
		return
	}
}
