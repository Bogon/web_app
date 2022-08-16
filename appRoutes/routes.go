package appRoutes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"webapp.io/controllers/userHanlder"
	"webapp.io/logger"
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

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg":    "pong",
			"status": 0,
			"data": gin.H{
				"tips": "success",
			},
		})
	})

	return r
}
