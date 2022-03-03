package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("ping", pingpongHandler)
	r.Run(":8080")
}

func pingpongHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
