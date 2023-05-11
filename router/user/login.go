package user

import (
	"easy-gin/controllers/user"
	"github.com/gin-gonic/gin"
)

type LoginRouter struct {
}

func (s *LoginRouter) InitLoginRouter(Router *gin.RouterGroup) {
	loginRouter := Router.Group("login").Use()
	{
		loginControllers := new(user.LoginControllers)
		loginRouter.POST("/register", loginControllers.Register)
		loginRouter.POST("/login", loginControllers.Login)
	}
}
