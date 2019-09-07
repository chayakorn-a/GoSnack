package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong get"})
	})
	r.POST("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong post"})
	})
	r.Run(":1234")
}
