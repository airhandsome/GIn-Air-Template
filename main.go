package main

import (
	"fmt"
	"gin-air-template/config"
	"gin-air-template/graceful"
	"gin-air-template/middleware"
	"gin-air-template/models"
	"gin-air-template/routes"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"log"
	"net/http"
	"time"
)

func main() {
	// 加载配置
	config.InitConfig()

	// 设置Gin运行模式
	gin.SetMode(config.Config.RunMode)

	// 创建Gin实例
	r := gin.New()

	// 使用中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.CircuitBreakerMiddleware(5, 5*time.Minute))
	r.Use(middleware.RateLimiterMiddleware(rate.NewLimiter(rate.Every(time.Minute), 10)))

	// 初始化路由
	routes.InitRoutes(r)

	// 连接数据库
	db, err := config.InitDatabase(config.DatabaseConfig{
		Dialect:     "sqlite",
		DSN:         "file:test.db?cache=shared&mode=memory",
		Migrate:     true,
		AutoMigrate: true,
	})
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	// 初始化模型
	models.InitModels(db)

	// 启动服务
	// 创建HTTP服务器
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Config.HttpPort),
		Handler: r,
	}

	// 在goroutine中启动服务器
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 设置优雅退出的逻辑
	graceful.GracefulShutdown(srv, 5*time.Second)
}
