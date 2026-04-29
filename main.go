package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./ui/build"))
	http.Handle("/", fs)
	http.ListenAndServe(":8080", nil)
	fmt.Println("Starting server on 8080")
}
