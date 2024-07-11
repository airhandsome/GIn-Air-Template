package controllers

import (
	"gin-air-template/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping 返回一个简单的 pong 响应
func Ping(c *gin.Context) {
	// 这里可以添加任何业务逻辑，目前仅返回一个固定的响应
	response := utils.Response{
		Code:    http.StatusOK,
		Message: "pong",
	}

	c.JSON(http.StatusOK, response)
}
