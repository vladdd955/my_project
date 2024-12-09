package handlers

import (
	"myproject/pkg/config"
	"myproject/pkg/render"
	"net/http"
	// "html/template"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// NewRepo create a new one repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers set the repository
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html")
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html")
}

func (m *Repository) Register(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "register.page.html")
}
