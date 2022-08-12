package userHanlder

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"webapp.io/models"
)

// UserSignUpHandler is a function that takes a pointer to a gin.Context and returns nothing.
func UserSignUpHandler(c *gin.Context) {
	// 1. 获取参数和参数校验
	var p models.ParamSignUp

	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("SignUp invalidate param", zap.Error(err))
		// 请求参数有误，直接返回响应
		c.JSON(http.StatusOK, gin.H{
			"code": 10002,
			"msg":  "参数有误",
			"data": nil,
		})
		return
	}
	fmt.Println(p)

	// 手动对请求参数进行业务规则校验
	if len(p.Username) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 || p.Password != p.RePassword {
		zap.L().Error("SignUp invalidate param")
		// 请求参数有误，直接返回响应
		c.JSON(http.StatusOK, gin.H{
			"code": 10002,
			"msg":  "参数有误",
			"data": nil,
		})
		return
	}
	// 2. 业务处理
	//logic.SignUp()
	//user.SignUp()

	// 3. 返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"code": 0,
		"data": gin.H{
			"username": "夏利",
			"nickname": "溜溜没",
			"age":      18,
			"birthday": "1990-02-14",
		},
	})
}
