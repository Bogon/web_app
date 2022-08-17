package appRoutes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"webapp.io/controllers/responseHandler"
	"webapp.io/controllers/userHanlder"
	"webapp.io/logger"
	"webapp.io/middlewares/jwtauth"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin 设置成发布模式
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 注册业务路由 - 注册
	r.POST("signup", userHanlder.UserSignUpHandler)
	// 注册业务路由 - 登录
	r.POST("login", userHanlder.UserLoginHandler)

	r.GET("/", func(c *gin.Context) {
		appVersion := viper.GetString("app.version")
		c.String(http.StatusOK, fmt.Sprintf("v%v", appVersion))
	})

	// 需求：该接口只有在登录情况下才可以调用
	r.GET("/ping", jwtauth.JWTAuthMiddleware(), func(c *gin.Context) {
		// 已经登录 正常显示
		responseHandler.ResponseSuccess(c, gin.H{
			"tips": "success",
		})
	})

	return r
}
