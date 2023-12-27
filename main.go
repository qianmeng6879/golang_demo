package main

import (
	"github.com/gin-gonic/gin"
	"mxzero.top/handles"
)

func main() {

	var server = gin.Default()

	server.POST("/token", handles.TokenCreateHandle)

	server.POST("/token/verify", handles.TokenVerifyHandle)

	server.Run()

}
