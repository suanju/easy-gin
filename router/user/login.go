package user

import "github.com/gin-gonic/gin"

type LoginRouter struct {
}

func (s *LoginRouter) InitLoginRouter(Router *gin.RouterGroup) {
	loginRouter := Router.Group("login").Use()
	{
		loginControllers := new(users.LoginControllers)
		loginRouter.POST("/register", loginControllers.Register)
		loginRouter.POST("/login", loginControllers.Login)
	}
}
