package graceful

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// GracefulShutdown 设置优雅退出的逻辑
func GracefulShutdown(srv *http.Server, timeout time.Duration) {
	// 创建系统信号接收通道
	quit := make(chan os.Signal, 1)
	// 注册要接收的信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// 等待接收信号
	<-quit
	log.Println("Shutdown Server ...")

	// 创建一个超时的context
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// 优雅关闭服务器
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}

	log.Println("Server exiting")
}
