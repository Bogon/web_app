package main

import (
	"context"
	"flag"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"webapp.io/appRoutes"
	"webapp.io/dao/mysql"
	"webapp.io/dao/redis"
	"webapp.io/logger"
	"webapp.io/pkg/snowflakeID"
	"webapp.io/settings"
)

// Go Web开发较通用的脚手架模板

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
	if err := logger.Init(settings.Conf.LogConf); err != nil {
		zap.L().Error("logger init failed, error:", zap.Error(err))
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success…")

	// 7. 分布式ID生成演示
	if err := snowflakeID.Init(settings.Conf.SnowflakeConf.StartTime, settings.Conf.SnowflakeConf.MachineID); err != nil {
		zap.L().Fatal("snowflake uuid failed: ", zap.Error(err))
		return
	}
	uuid := snowflakeID.GetId()
	zap.L().Debug(fmt.Sprintf("uuid: %v \n", uuid))

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

	// 5. 注册路由
	r := routes.Setup()
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
