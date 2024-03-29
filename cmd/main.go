package main

import (
	"github.com/alexedwards/scs/v2"
	"github.com/koleaby4/go-bookings/pkg/config"
	"github.com/koleaby4/go-bookings/pkg/handlers"
	"github.com/koleaby4/go-bookings/pkg/render"
	"log"
	"net/http"
	"time"
)

var app config.AppConfig
var session *scs.SessionManager = scs.New()

func main() {

	app.InProduction = false
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatalln("error populating template cache.", err)
	}

	app.TemplateCache = tc
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	address := "localhost:8080"

	//http.HandleFunc("/", repo.Home)
	//http.HandleFunc("/home", repo.Home)
	//http.HandleFunc("/about", repo.About)
	//err = http.ListenAndServe(address, nil)

	srv := &http.Server{
		Addr:    address,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalln("error while listening", err)
	}
}
