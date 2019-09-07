package main

import ("github.com/gin-gonic/gin"
"net/http")

type Todo struct {
	ID     string `json:id`
	Title  string `json:title`
	Status string `json:status`
}

var jsons []Todo

func main() {
	r := gin.Default()

	r.GET("/todos", func(c *gin.Context) {
		c.JSON(http.StatusOK, jsons)
		//c.JSON(200, gin.H{"message": "pong get"})
	})
	r.POST("/todos", func(c *gin.Context) {
		var json Todo

		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		jsons = append(jsons, json)
		c.JSON(200, gin.H{"message": "pong post"})
	})
	r.Run(":1234")
}
