package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/lib/pq"
)

type Todo struct {
	ID     string `json:id`
	Title  string `json:title`
	Status string `json:status`
}

func CreateTodoTable(db *sql.DB) {
	createTb := `
	CREATE TABLE IF NOT EXISTS todos (
	id SERIAL PRIMARY KEY,
	title TEXT,
	status TEXT
	);
	`
	_, err := db.Exec(createTb)

	if err != nil {
		log.Fatal("can't create table ", err)
	} else {
		fmt.Println("create table success")
	}
}

func InsertTodoTable(db *sql.DB) {
	row := db.QueryRow("INSERT INTO todos (title, status) values ($1, $2) RETURNING id", "buy bmw", "ac")
	var id int
	err := row.Scan(&id)
	if err != nil {
		fmt.Println("can't scan id", err)
		return
	}
	fmt.Println("insert todo scccess id : ", id)
}

func QueryTodoTable(db *sql.DB) {
	stmt, err := db.Prepare("SELECT id, title, status FROM todos where id=$1")
	if err != nil {
		log.Fatal("can't prepare query one row statement", err)
	}
	rowId := 1
	row := stmt.QueryRow(rowId)
	var id int
	var title, status string

	err = row.Scan(&id, &title, &status)
	if err != nil {
		log.Fatal("can't Scan row into variables", err)
	}
	fmt.Println("one row", id, title, status)
}

func UpdateTodoTable(db *sql.DB) {
	stmt, err := db.Prepare("UPDATE todos SET status=$2 WHERE id=$1;")
	if err != nil {
		log.Fatal("can't prepare statement update", err)
	}

	if _, err := stmt.Exec(1, "inactive"); err != nil {
		log.Fatal("error execute update", err)
	}
	fmt.Println("update success")
}

func QueryAllTodoTable(db *sql.DB) {
	stmt, err := db.Prepare("SELECT id, title, status FROM todos")
	if err != nil {
		log.Fatal("can't prepare query all todos statement", err)
	}

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal("can't query all todos", err)
	}
	var todos []Todo
	for rows.Next() {
		var id int
		var title, status string
		err := rows.Scan(&id, &title, &status)
		if err != nil {
			log.Fatal("can't scan row into variable", err)
		}
		todo := Todo{strconv.Itoa(id), title, status}
		todos = append(todos, todo)
		//fmt.Println(id, title, status)
	}
	fmt.Println(todos)
	fmt.Println("query all todos success")
}
