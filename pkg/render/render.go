package render

import (
	"fmt"
	"github.com/koleaby4/go-bookings/pkg/config"
	"github.com/koleaby4/go-bookings/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}
func CreateTemplateCache() (map[string]*template.Template, error) {
	result := map[string]*template.Template{}

	templates, err := filepath.Glob("./templates/*.html")
	if err != nil {
		log.Fatalf("error reading templates dir. err=%v", err)
	}

	layouts, err := filepath.Glob("./templates/*.layout")
	if err != nil {
		return result, err
	}

	for _, file := range templates {
		fileName := filepath.Base(file)
		fmt.Println("parsing", file)
		t, err := template.New(fileName).ParseFiles(file)
		if err != nil {
			return result, err
		}
		t, err = t.ParseFiles(layouts...)
		if err != nil {
			return result, err
		}

		result[fileName] = t
	}
	return result, nil
}

var cache, _ = CreateTemplateCache()

func addDefaultData(data *models.TemplateData) *models.TemplateData {
	return data
}
func RenderTemplate(w http.ResponseWriter, templateName string, data *models.TemplateData) {
	tc := app.TemplateCache
	t, ok := tc[templateName]
	if !ok {
		log.Fatalf("could not get template=%v from the cache=%v", templateName, tc)
	}

	data = addDefaultData(data)

	if err := t.Execute(w, data); err != nil {
		log.Fatalln("error executing template", t, err)
	}
}
