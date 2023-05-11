package middlewares

import (
	"easy-gin/models/user"
	"easy-gin/utils/jwt"
	ControllersCommon "easy-gin/utils/response"
	"github.com/gin-gonic/gin"
)

//VerificationToken 请求头中携带token
func VerificationToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		claim, err := jwt.ParseToken(token)
		if err != nil {
			ControllersCommon.NotLogin(c, "Token过期")
			c.Abort()
			return
		}
		u := new(users.User)
		if !u.IsExistByField("id", claim.UserID) {
			ControllersCommon.NotLogin(c, "用户异常")
			c.Abort()
			return
		}
		c.Set("uid", u.ID)
		c.Set("currentUserName", u.Username)
		c.Next()
	}
}
