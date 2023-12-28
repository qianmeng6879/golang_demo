package main

import (
	"github.com/gin-gonic/gin"
	"mxzero.top/handles"
	"mxzero.top/middlewares"
	"mxzero.top/rest"
)

func main() {
	var server = gin.Default()
	var api_v1 = server.Group("/v1")
	api_v1.Use(middlewares.JWTAuthenticate())
	{
		api_v1.GET("/users", func(ctx *gin.Context) {
			ctx.JSON(200, rest.Success("test"))
		})
	}

	server.POST("/token", handles.TokenCreateHandle)

	server.POST("/token/verify", handles.TokenVerifyHandle)

	server.GET("/users/userinfo", handles.UserInfoService)
	server.Run()

}
