package main

import (
	"log"
	"myproject/pkg/config"
	"myproject/pkg/render"

	// "fmt"
	"myproject/pkg/handlers"
	"net/http"
)

const myPort = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Error creating template cache:")
	}
	app.TemplateCache = tc

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/register", handlers.Register)

	_ = http.ListenAndServe(myPort, nil)
}

// 34
