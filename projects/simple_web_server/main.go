package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r "")

func main() {
	fileServer = http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err: http.ListenAndServes(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
