package middlewares

import (
	"github.com/gin-gonic/gin"
	"mxzero.top/rest"
	"mxzero.top/utils"
)

func JWTAuthenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var auth string = ctx.Request.Header.Get("Authorization")
		if auth == "" || auth[0:7] != "Bearer " {

			ctx.JSON(401, rest.Error("未提供身份凭证", 401))
			ctx.Abort()
			return
		}
		var token = auth[7:]
		if !utils.VerifyToken(token) {
			ctx.JSON(401, rest.Error("未提供身份凭证", 401))
			ctx.Abort()
			return
		}

		subject, err := utils.ParseToken(token)
		if err != nil {
			ctx.JSON(401, rest.Error("认证失败", 401))
			ctx.Abort()
			return
		}
		ctx.Set("x-user-id", subject["id"])
		ctx.Next()
	}
}
