package handlers

import (
	"github.com/koleaby4/go-bookings/pkg/config"
	"github.com/koleaby4/go-bookings/pkg/models"
	"github.com/koleaby4/go-bookings/pkg/render"
	"net/http"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{App: a}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (repo *Repository) Home(w http.ResponseWriter, r *http.Request) {
	repo.App.Session.Put(r.Context(), "remote_ip", r.RemoteAddr)
	render.RenderTemplate(w, "home.html", &models.TemplateData{})
}

func (repo *Repository) About(w http.ResponseWriter, r *http.Request) {
	data := models.TemplateData{Strings: map[string]string{"FirstName": "James", "LastName": "Bond"}}
	data.Strings["remote_ip"] = repo.App.Session.Get(r.Context(), "remote_ip").(string)
	render.RenderTemplate(w, "about.html", &data)
}
