package main

import "github.com/gin-gonic/gin"

func purejson(c *gin.Context) {
	c.PureJSON(200, gin.H{
		"html": "<b>Hello, world!</b>",
	})
}
