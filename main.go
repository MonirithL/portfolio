package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./ui/build"))
	http.Handle("/", fs)
	fmt.Println("Starting server on 8080")
	http.ListenAndServe("localhost:8080", nil)
}
