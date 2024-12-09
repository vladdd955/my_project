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
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	http.HandleFunc("/register", handlers.Repo.Register)

	_ = http.ListenAndServe(myPort, nil)
}

// 34
