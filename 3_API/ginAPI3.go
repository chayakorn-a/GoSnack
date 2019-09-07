package main

import ("github.com/gin-gonic/gin"
"net/http"
"strconv"
"fmt")

type Todo struct {
	ID     string `json:id`
	Title  string `json:title`
	Status string `json:status`
}

func GetTodoHandler(c *gin.Context){
	t1 := c.DefaultQuery("name", "noname")
	c.JSON(http.StatusOK, todos)
	fmt.Println(t1)
}

func CreateTodoHandler(c *gin.Context){
	var todo Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ids++
	todo.ID = strconv.Itoa(ids)
	todos = append(todos, todo)
	//c.JSON(200, gin.H{"message": "pong post"})
	c.JSON(200, map[string]string{"message": "pong post"})
}

func getIdTodoHandler(c *gin.Context){
	idd := c.Param("id")

	for _,n := range todos{
		if n.ID == idd{
			c.JSON(http.StatusOK, n)
			return 
		}
	}
}

func DeleteTodoHandler(c *gin.Context){
	idd := c.Param("id")
	num:=Find(todos,idd)
	remove(todos,num)
	todos = todos[:(len(todos)-1)]
	c.JSON(http.StatusOK, todos)
	//fmt.Println(idd)
	//c.JSON(200, map[string]string{"message": "deleted", "id": idd})
}

func updateTodoHandler(c *gin.Context){
	var todo Todo
	idd := c.Param("id")

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	num:=Find(todos,idd)
	todos[num].Title = todo.Title
	todos[num].Status = todo.Status

	c.JSON(http.StatusOK, todos[num])
}

func Find(a []Todo, id string) int {
	for i, n := range a {
			if id == n.ID {
					return i
			}
	}
	return len(a)
}

func remove(slice []Todo, s int) []Todo {
    return append(slice[:s], slice[s+1:]...)
}

var todos []Todo
var ids int = 0

func main() {
	r := gin.Default()

	r.GET("/todos", GetTodoHandler)
	r.GET("/todos/:id", getIdTodoHandler)
	r.POST("/todos",CreateTodoHandler)
	r.DELETE("/todos/:id",DeleteTodoHandler)
	r.PUT("/todos/:id",updateTodoHandler)

	r.Run(":1234")
}

// how to build
// GOOS=linux go build -o <app name>