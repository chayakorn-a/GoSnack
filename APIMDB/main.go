package main

import (
	"net/http"
	//"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"APIMDB/database"
)

// On git Bash
// DATABASE_URL=postgres://ednlnmoi:huVNUQvl1bR9puqxPOrGkJ-pPiVQ23hY@otto.db.elephantsql.com:5432/ednlnmoi go run main.go
func main() {
	r := gin.Default()

	/*r.GET("/todos", GetTodoHandler)       // QueryAllTodoTable
	r.GET("/todos/:id", getIdTodoHandler) // QueryTodoTable
	r.POST("/todos", CreateTodoHandler)   // InsertTodoTable
	r.DELETE("/todos/:id", DeleteTodoHandler)
	r.PUT("/todos/:id", updateTodoHandler) // UpdateTodoTable
*/
	r.POST("/sync", CreateSyncHandler)
	r.Run(":1234")
}

func CreateSyncHandler(c *gin.Context) {
	c.JSON(http.StatusOK, database.SyncToDb())
}