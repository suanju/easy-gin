package user

import (
	"easy-gin/controllers"
	receive "easy-gin/interaction/receive/user"
	"easy-gin/service/user"
	"github.com/gin-gonic/gin"
)

type LoginControllers struct {
	controllers.BaseControllers
}

//Login 登入
func (lg LoginControllers) Login(ctx *gin.Context) {
	if rec, err := controllers.ShouldBind(ctx, new(receive.LoginReceive)); err == nil {
		results, err := user.Login(rec)
		lg.Response(ctx, results, err)
	}
}

//Register 注册
func (lg LoginControllers) Register(ctx *gin.Context) {
	if rec, err := controllers.ShouldBind(ctx, new(receive.RegisterReceive)); err == nil {
		results, err := user.Register(rec)
		lg.Response(ctx, results, err)
	}
}
