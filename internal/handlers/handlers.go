package handlers

import (
	"fmt"
	"net/http"

	"github.com/direwolf707/go-web-app/internal/config"
	"github.com/direwolf707/go-web-app/internal/models"
	"github.com/direwolf707/go-web-app/internal/render"
)

// Repository is repository type
type Repository struct {
	App *config.AppConfig
}

// Repo is the repository used by handlers
var Repo *Repository

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("hello request initiated")
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(rw, r, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("about request initiated")
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	render.RenderTemplate(rw, r, "about.page.tmpl", &models.TemplateData{
		StringMap: map[string]string{
			"test":      "Hello,again!!!",
			"remote_ip": remoteIP,
		},
	})
}

func (m *Repository) Reservations(rw http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(rw, r, "make-reservation.page.tmpl", &models.TemplateData{})
}
func (m *Repository) Generals(rw http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(rw, r, "generals.page.tmpl", &models.TemplateData{})
}
func (m *Repository) Majors(rw http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(rw, r, "majors.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Availability(rw http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(rw, r, "search-availability.page.tmpl", &models.TemplateData{})
}

func (m *Repository) PostAvailability(rw http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	rw.Write([]byte(fmt.Sprintf("start: %s end %s", start, end)))
	//render.RenderTemplate(rw, "search-availability.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Contact(rw http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(rw, r, "contact.page.tmpl", &models.TemplateData{})
}
