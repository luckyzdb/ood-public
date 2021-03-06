package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func JSONP(c *gin.Context) {
	data := map[string]interface{}{
		"foo": "bar",
	}

	// /JSONP?callback=x
	// 将输出：x({\"foo\":\"bar\"})
	c.JSONP(http.StatusOK, data)
}
