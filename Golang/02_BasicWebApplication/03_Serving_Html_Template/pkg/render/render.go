package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/bishal7679/Serving-html-template/pkg/config"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates set the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// this function is used to render the template
func RenderTemplate(w http.ResponseWriter, tmpl string) {

	tc := app.TemplateCache
	// // get the template cache from app config
	// tc, err := CreateTemplateCache()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Couldn't get the template from template cache")
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, nil)

	_, err := buf.WriteTo(w)

	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}
	// parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	// err = parsedTemplate.Execute(w, nil)
	// if err != nil {
	// 	fmt.Println("error parsing template", err)
	// 	return
	// }
}

// CreateTemplateCache creates a template cache
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		match, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(match) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts

	}

	return myCache, nil
}
