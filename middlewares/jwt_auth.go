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

		if !utils.VerifyToken(auth[7:]) {
			ctx.JSON(401, rest.Error("未提供身份凭证", 401))
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
