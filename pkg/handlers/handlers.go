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

func (repo *Repository) Index(w http.ResponseWriter, r *http.Request) {
	repo.App.Session.Put(r.Context(), "remote_ip", r.RemoteAddr)
	render.RenderTemplate(w, "index.html", &models.TemplateData{})
}

func (repo *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "majors.html", &models.TemplateData{})
}

func (repo *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "generals.html", &models.TemplateData{})
}
func (repo *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "reservation.html", &models.TemplateData{})
}
func (repo *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "make-reservation.html", &models.TemplateData{})
}

func (repo *Repository) About(w http.ResponseWriter, r *http.Request) {
	data := models.TemplateData{Strings: map[string]string{}}

	remoteIp := repo.App.Session.GetString(r.Context(), "remote_ip")
	if remoteIp == "" {
		repo.App.Session.Put(r.Context(), "remote_ip", r.RemoteAddr)
		remoteIp = r.RemoteAddr
	}
	data.Strings["remote_ip"] = remoteIp
	render.RenderTemplate(w, "about.html", &data)
}
