package main

import (
	// "fmt"
	"myproject/pkg/handlers"
	"net/http"
)

const myPort = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/register", handlers.Register)

	_ = http.ListenAndServe(myPort, nil)
}
