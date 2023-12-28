package app

import (
	"github.com/gin-gonic/gin"
	"mxzero.top/handles"
	"mxzero.top/middlewares"
)

func CreateApp() *gin.Engine {
	var server = gin.Default()
	var api_v1 = server.Group("/v1")

	{
		api_v1.POST("/token", handles.TokenCreateHandle)
		api_v1.POST("/token/verify", handles.TokenVerifyHandle)
	}

	var api_v1_auth_require = api_v1.Group("", middlewares.JWTAuthenticate())
	{
		api_v1_auth_require.GET("/users/userinfo", handles.UserInfoService)
	}

	return server
}
