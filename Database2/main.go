package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// On git Bash
// DATABASE_URL=postgres://ednlnmoi:huVNUQvl1bR9puqxPOrGkJ-pPiVQ23hY@otto.db.elephantsql.com:5432/ednlnmoi go run main.go
func main() {
	url := os.Getenv("DATABASE_URL")
	fmt.Println(url)
	db, err := sql.Open("postgres", url)

	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer db.Close()

	//CreateTodoTable(db)
	//InsertTodoTable(db)
	//QueryTodoTable(db)
	//UpdateTodoTable(db)
	// QueryTodoTable(db)

	QueryAllTodoTable(db)
}
