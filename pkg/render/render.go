package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/direwolf707/go-web-app/pkg/config"
	"github.com/direwolf707/go-web-app/pkg/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates sets config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate renders templates using html/template
func RenderTemplate(rw http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		// get template cache
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}
	// find required template
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Template not found in cache")
	}
	td = AddDefaultData(td)
	// create empty memory buffer
	buf := new(bytes.Buffer)
	// execute template data in buffer
	_ = t.Execute(buf, td)
	// write buffer to response
	_, err := buf.WriteTo(rw)
	//
	//t.Execute(rw, nil)
	//
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}
}

// CreateTemplateCache returns a template cache
func CreateTemplateCache() (map[string]*template.Template, error) {
	// create empty cache
	myCache := map[string]*template.Template{}
	// getting list of all html (not layouts)
	pages, err := filepath.Glob("./templates/*.html")
	if err != nil {
		return myCache, err
	}
	// iterating through pages
	for _, page := range pages {
		// extracting file name (with extension)
		name := filepath.Base(page)
		// parsing template
		t, err := template.New(name).Funcs(functions).ParseFiles(page, "./templates/base.layout")
		if err != nil {
			return myCache, err
		}
		// adding to cache
		myCache[name] = t
	}
	return myCache, nil
}
