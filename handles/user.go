package handles

import (
	"github.com/gin-gonic/gin"
	"mxzero.top/rest"
)

func UserInfoService(c *gin.Context) {
	c.JSON(200, rest.Success(map[string]interface{}{
		"id":       10,
		"username": "zero",
	}))
}
