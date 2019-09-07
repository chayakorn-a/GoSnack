package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Todo struct {
	ID     string `json:id`
	Title  string `json:title`
	Status string `json:status`
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	resp := []byte(`"hello GET todoss"`)
	//resp := []byte(`""{"name":"chayakorn"}""`)
	w.Write(resp)
}

func AllocateHandler() func(http.ResponseWriter, *http.Request) {
	//var t Todo
	var todos []Todo

	return func(w http.ResponseWriter, req *http.Request) {
		method := req.Method
		// fmt.Fprintf(w, "hello %s todos", method)
		if method == "POST" {
			body, err := ioutil.ReadAll(req.Body)
			if err != nil {
				fmt.Fprintf(w, "error : %v", err)
				return
			}
			var t Todo
			err = json.Unmarshal(body, &t)
			if err != nil {
				fmt.Fprintf(w, "error: ", err)
				return
			}
			todos = append(todos, t)
			fmt.Printf("body POST Struct: % #v\n", t)
			fmt.Fprintf(w, "%s is OK", method)
		} else {
			fmt.Printf("body GET Struct: % #v\n", todos)
			//var t Todo
			b, err := json.Marshal(todos)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application.json")
			w.Write(b)
			//fmt.Fprintf(w, "hello %s todos", method)
		}
	}
}

func main() {
	todoHandler := AllocateHandler()
	http.HandleFunc("/todos", todoHandler)
	log.Fatal(http.ListenAndServe(":1235", nil))
}
