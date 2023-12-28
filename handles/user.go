package handles

import (
	"github.com/gin-gonic/gin"
	"mxzero.top/models"
	"mxzero.top/rest"
)

func UserInfoService(c *gin.Context) {
	var userId = c.MustGet("x-user-id").(float64)
	var user models.User
	models.Db.Where("id = ?", int(userId)).First(&user)
	c.JSON(200, rest.Success(user))
}
