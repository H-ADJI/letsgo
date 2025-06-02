package main

import (
	"html/template"
	"net/http"
	"path/filepath"
	"time"

	"github.com/H-ADJI/letsgo/internal/models"
	"github.com/H-ADJI/letsgo/internal/validator"
)

var funcMap = template.FuncMap{"humanDate": humanDate}

type userSignupForm struct {
	Name     string
	Email    string
	Password string
	validator.Validator
}
type snippetCreateForm struct {
	Title   string
	Content string
	Expires int
	validator.Validator
}
type TemplateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
	Form     any
	Flash    string
}

func (a *app) NewTemplateData(r *http.Request) TemplateData {
	return TemplateData{Flash: a.sessionManager.PopString(r.Context(), "flash")}
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
