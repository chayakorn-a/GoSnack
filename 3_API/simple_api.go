package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	resp := []byte(`"hello GET todoss"`)
	//resp := []byte(`""{"name":"chayakorn"}""`)
	w.Write(resp)
}

func todoHandler(w http.ResponseWriter, req *http.Request) {
	method := req.Method
	// fmt.Fprintf(w, "hello %s todos", method)
	if method == "POST" {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(w, "error : %v", err)
			return
		}

		fmt.Printf("body : %s\n", body)
		fmt.Fprintf(w, "%s is OK", method)
	} else {
		fmt.Fprintf(w, "hello %s todos", method)
	}
}

func main() {
	http.HandleFunc("/todos", todoHandler)
	log.Fatal(http.ListenAndServe(":1235", nil))
}
