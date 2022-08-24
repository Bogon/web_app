package userHanlder

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"webapp.io/controllers/responseCode"
	"webapp.io/controllers/responseHandler"
	"webapp.io/controllers/validatorHandler"
	"webapp.io/dao/mysql"
	"webapp.io/logic/user"
	"webapp.io/models"
)

// UserSignUpHandler is a function that takes a pointer to a gin.Context and returns nothing.
// UserSignUpHandler 注册
// @Summary 注册
// @Description 注册新用户
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query models.ParamSignUp false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} responseHandler.ResponseData
// @Router /api/v1/signup [post]
func UserSignUpHandler(c *gin.Context) {
	// 1. 获取参数和参数校验
	p := new(models.ParamSignUp)

	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("SignUp invalidate param", zap.Error(err))
		// 判断对象是不是 validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 请求参数有误，直接返回响应
			responseHandler.ResponseError(c, responseCode.CodeInvalidParam)
			return
		}
		// 请求参数有误，直接返回响应
		responseHandler.ResponseWithMsg(c, responseCode.CodeInvalidParam, validatorHandler.RemoveTopStruct(errs.Translate(validatorHandler.Trans)))
		return
	}

	// 2. 业务处理
	if err := user.SignUp(p); err != nil {
		if errors.Is(err, mysql.ErrorUserExist) {
			responseHandler.ResponseError(c, responseCode.CodeUserExist)
			return
		}
		responseHandler.ResponseError(c, responseCode.CodeServerBasy)
		return
	}

	// 3. 返回响应
	responseHandler.ResponseSuccess(c, "注册成功")
}

// UserLoginHandler is a function that takes a pointer to a gin.Context and returns nothing.
// UserLoginHandler 登录
// @Summary 登录
// @Description 登录获取token
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query models.ParamLogin false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} responseHandler.ResponseData
// @Router /api/v1/login [post]
func UserLoginHandler(c *gin.Context) {
	u := new(models.ParamLogin)
	// 1. 获取请求参数，进行参数校验
	if err := c.ShouldBindJSON(u); err != nil {
		zap.L().Error("Login invalidate param", zap.Error(err))
		// 判断对象是不是 validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 请求参数有误，直接返回响应
			responseHandler.ResponseError(c, responseCode.CodeInvalidParam)
			return
		}
		// 请求参数有误，直接返回响应
		responseHandler.ResponseWithMsg(c, responseCode.CodeInvalidParam, validatorHandler.RemoveTopStruct(errs.Translate(validatorHandler.Trans)))
		return
	}
	// 2. 业务逻辑处理
	user, err := user.Login(u)
	if err != nil {
		zap.L().Error("user.Login failed", zap.String("username", u.Username), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			responseHandler.ResponseError(c, responseCode.CodeUserNotExist)
			return
		}
		responseHandler.ResponseError(c, responseCode.CodeInvalidPassword)
		return
	}
	// 3. 返回响应
	responseHandler.ResponseSuccess(c, gin.H{
		"userID":   fmt.Sprintf("%v", user.UserID),
		"userName": user.Username,
		"token":    user.Token,
	})
}
