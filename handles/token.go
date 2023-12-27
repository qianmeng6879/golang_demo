package handles

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/cristalhq/jwt/v5"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"mxzero.top/models"
	"mxzero.top/rest"
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

	newToken, _ := jwt.Parse([]byte(param.Token), verifier)

	err := verifier.Verify(newToken)
	if err != nil {
		c.JSON(400, rest.Error("token无效", 400))
		return
	}
	var newClaims jwt.RegisteredClaims
	err = json.Unmarshal(newToken.Claims(), &newClaims)
	if err != nil {
		c.JSON(400, rest.Error("token无效", 400))
		return
	}

	var subjectData map[string]interface{}
	err = json.Unmarshal([]byte(newClaims.Subject), &subjectData)
	if err != nil {
		c.JSON(400, rest.Error("解析subject失败", 400))
		return
	}
	c.JSON(200, rest.Success(subjectData))

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
			subject, _ := json.Marshal(map[string]interface{}{
				"id":       user.ID,
				"username": user.Username,
				"is_admin": user.IsAdmin,
			})

			var current time.Time = time.Now()
			var express time.Time = current.Add(2 * time.Hour)

			claims := &jwt.RegisteredClaims{
				ID:        uuid.New().String(),
				Subject:   string(subject),
				Issuer:    "open.mxzero.top",
				IssuedAt:  jwt.NewNumericDate(current),
				ExpiresAt: jwt.NewNumericDate(express),
			}
			token, _ := jwt.NewBuilder(signer).Build(claims)

			var accessToken string = token.String()
			c.JSON(200, rest.Success(map[string]interface{}{
				"access_token": accessToken,
				"express":      7200,
			}))
		}
	} else {
		c.JSON(400, rest.Error("用户名或密码为空", 400))
	}

}
