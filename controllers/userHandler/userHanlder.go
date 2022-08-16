package userHanlder

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"webapp.io/controllers/validatorHandler"
	"webapp.io/logic/user"
	"webapp.io/models"
)

// UserSignUpHandler is a function that takes a pointer to a gin.Context and returns nothing.
func UserSignUpHandler(c *gin.Context) {
	// 1. 获取参数和参数校验
	p := new(models.ParamSignUp)

	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("SignUp invalidate param", zap.Error(err))
		// 判断对象是不是 validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 请求参数有误，直接返回响应
			c.JSON(http.StatusOK, gin.H{
				"code": 10002,
				"msg":  err.Error(),
				"data": nil,
			})
		}
		// 请求参数有误，直接返回响应
		c.JSON(http.StatusOK, gin.H{
			"code": 10002,
			"msg":  validatorHandler.RemoveTopStruct(errs.Translate(validatorHandler.Trans)),
			"data": nil,
		})
		return
	}

	// 2. 业务处理
	if err := user.SignUp(p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10003,
			"msg":  "注册失败",
			"data": nil,
		})
		return
	}

	// 3. 返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"code": 0,
		"data": nil,
	})
}
