package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// RoutersGroup 路由总集
type RoutersGroup struct {
}

// RoutersGroupApp 实例化
var RoutersGroupApp = new(RoutersGroup)

func InitRouter() {
	router := gin.Default()
	// 添加 cors 中间件
	config := cors.DefaultConfig()
	router.Use(cors.New(config))
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
