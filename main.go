package main

import (
	"context"
	"flag"
	"fmt"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"webapp.io/appRoutes"
	"webapp.io/controllers/validatorHandler"
	"webapp.io/dao/mysql"
	"webapp.io/dao/redis"
	_ "webapp.io/docs"
	"webapp.io/logger"
	"webapp.io/pkg/snowflakeID"
	"webapp.io/settings"
)

// Go Web开发较通用的脚手架模板

// @title webapp.io 帖子发布系统
// @version 1.0
// @description 帖子发布系统
// @termsOfService http://swagger.io/terms/

// @contact.name Senyas
// @contact.url http://www.swagger.io/support
// @contact.email sains@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host http://127.0.0.1:8001
// @BasePath /api/v1
func main() {
	var filePath string
	flag.StringVar(&filePath, "f", "./conf/config.yaml", "配置文件相对路径")
	//if len(os.Args) < 2 {
	//	fmt.Println("need config file. eg: web_app.io ./conf/config.yaml")
	//	return
	//}

	// 1. 加载配置
	if err := settings.Init(filePath); err != nil {
		fmt.Println("settings init failed, error:", err)
		return
	}

	// 2. 初始化日志
	if err := logger.Init(settings.Conf.LogConf, settings.Conf.AppConf.Mode); err != nil {
		zap.L().Error("logger init failed, error:", zap.Error(err))
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success…")

	// 3. 初始化 MySQL 连接
	if err := mysql.Init(settings.Conf.MySQLConf); err != nil {
		zap.L().Error("mysql init failed, error:", zap.Error(err))
		return
	}
	defer mysql.Close()
	zap.L().Debug("mysql init success…")

	// 4. 初始化 Redis 连接
	if err := redis.Init(settings.Conf.RedisConf); err != nil {
		zap.L().Error("redis init failed, error:", zap.Error(err))
		return
	}
	defer redis.Close()
	zap.L().Debug("redis init success…")

	// 初始化 gin 框架内置的校验器使用的翻译器
	if err := validatorHandler.InitTrans("zh"); err != nil {
		zap.L().Error("validator trans init failed, error:", zap.Error(err))
		return
	}

	// 5. 注册路由
	r := appRoutes.Setup(settings.Conf.AppConf.Mode)

	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	// 5.1 初始化分布式ID生成框架
	if err := snowflakeID.Init("2022-08-16", 1); err != nil {
		zap.L().Error("snowflakeID init failed, error:", zap.Error(err))
		return
	}

	// 6. 启动服务（优雅关机）
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", settings.Conf.AppConf.Port),
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("listen: %s\n", zap.Error(err))
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown: ", zap.Error(err))
	}

	zap.L().Info("Server exiting")

}
