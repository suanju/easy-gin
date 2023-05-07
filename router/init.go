package router

import "github.com/gin-gonic/gin"

//路由总集
type RoutersGroup struct {
}

//实例化
var RoutersGroupApp = new(RoutersGroup)

func InitRouter() {
	router := gin.Default()
	//跨域中间件
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
