package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"webapp.io/logger"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/", func(c *gin.Context) {
		appVersion := viper.GetString("app.version")
		c.String(http.StatusOK, fmt.Sprintf("v%v", appVersion))
	})

	return r
}
