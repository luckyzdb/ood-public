package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 提供 unicode 实体
	r.GET("/json", json)

	// 提供字面字符
	r.GET("/purejson", purejson)

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}
