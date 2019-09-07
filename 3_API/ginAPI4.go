package main

import ("github.com/gin-gonic/gin"
"net/http"
"fmt"
)

type Todo struct {
	ID     string `json:id`
	Title  string `json:title`
	Status string `json:status`
}

func authMiddleware(c *gin.Context){
	fmt.Println("This is a middleware")
	token := c.GetHeader("Authorization")
	if token != "Bearer token123"{
		c.JSON(http.StatusUnauthorized, gin.H{"error":"Unauthorized"})
		c.Abort()
		return
	}

	c.Next()
	fmt.Println("after in middleware")
}

func GetTodoHandler(c *gin.Context){
	//t1 := c.DefaultQuery("name", "noname")
	fmt.Println("In GetTodoHandler()")
	//c.JSON(http.StatusOK, todos)
	//fmt.Println(t1)
	//fmt.Println("Token : ", token)
}

func main() {
	r := gin.Default()

	r.Use(authMiddleware)
	r.GET("/todos", GetTodoHandler)

	r.Run(":1234")
}

// how to build
// GOOS=linux go build -o <app name>