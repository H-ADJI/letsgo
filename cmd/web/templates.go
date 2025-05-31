package main

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/H-ADJI/letsgo/internal/models"
)

var funcMap = template.FuncMap{"humanDate": humanDate}

type TemplateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}

func NewTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := filepath.Glob("./ui/html/pages/*.tmpl.html")
	if err != nil {
		return nil, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		files := []string{
			"./ui/html/base.tmpl.html",
			page,
		}
		ts, err := template.New(name).Funcs(funcMap).ParseFiles(files...)
		if err != nil {
			return nil, err
		}
		ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl.html")
		if err != nil {
			return nil, err
		}
		cache[name] = ts

	}
	return cache, nil
}
func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}
