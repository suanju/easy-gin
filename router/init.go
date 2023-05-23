package router

import (
	"easy-gin/consts"
	"easy-gin/router/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"path/filepath"
)

// RoutersGroup 路由总集
type RoutersGroup struct {
	user user.RouterGroup
}

// RoutersGroupApp 实例化路由
var RoutersGroupApp = new(RoutersGroup)

func init() {
	router := gin.Default()
	//采用cors默认跨域中间件
	//详细配置  https://github.com/gin-contrib/cors
	router.Use(cors.Default())
	PrivateGroup := router.Group("")
	PrivateGroup.Use()
	{
		router.Static(consts.Resource, filepath.ToSlash(consts.ResourcePath)) //如果您的项目中需要展示静态资源
		RoutersGroupApp.user.LoginRouter.InitLoginRouter(PrivateGroup)
	}

	err := router.Run()
	if err != nil {
		log.Fatal("路由加载失败")
	}
}
