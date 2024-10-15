package handlers

import (
	"myproject/pkg/render"
	"net/http"
	// "html/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html")
}

func Register(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "register.page.html")
}
