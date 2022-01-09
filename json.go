package main

import "github.com/gin-gonic/gin"

func json(c *gin.Context) {
	c.JSON(200, gin.H{
		"html": "<b>Hello, world!</b>",
	})
}
