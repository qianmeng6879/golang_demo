package handles

import (
	"fmt"

	"github.com/cristalhq/jwt/v5"
	"github.com/gin-gonic/gin"
	"mxzero.top/models"
	"mxzero.top/rest"
	"mxzero.top/utils"
)

var key = []byte(`secret`)
var signer, _ = jwt.NewSignerHS(jwt.HS256, key)
var verifier, _ = jwt.NewVerifierHS(jwt.HS256, key)

type UsernamePasswordParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Code     int    `json:"code"`
}

type TokenParam struct {
	Token string `json:"token"`
}

// 解析token
func TokenVerifyHandle(c *gin.Context) {
	var param TokenParam
	c.Bind(&param)
	if param.Token == "" {
		c.JSON(400, rest.Error("token is null", 400))
		return
	}

	subject, err := utils.ParseToken(param.Token)
	if err != nil {
		c.JSON(400, rest.Error(err.Error(), 400))
		return

	}
	c.JSON(200, rest.Success(subject))
}

// 获取token
func TokenCreateHandle(c *gin.Context) {
	var param UsernamePasswordParam
	c.Bind(&param)

	if param.Username != "" && param.Password != "" {
		fmt.Printf("username=%v password=%v\n", param.Username, param.Password)

		var user models.User
		models.Db.Where("username = ?", param.Username).First(&user)

		if user.ID == 0 || user.Password != param.Password {
			c.JSON(400, rest.Error("用户名与密码不匹配", 400))
		} else {
			var token = utils.CreateToken(map[string]interface{}{
				"id":       user.ID,
				"username": user.Username,
				"is_admin": user.IsAdmin,
			})
			c.JSON(200, rest.Success(map[string]interface{}{
				"access_token": token,
				"express":      7200,
			}))
		}
	} else {
		c.JSON(400, rest.Error("用户名或密码为空", 400))
	}

}
