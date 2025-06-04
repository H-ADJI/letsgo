package main

import (
	"html/template"
	"io/fs"
	"net/http"
	"path/filepath"
	"time"

	"github.com/H-ADJI/letsgo/internal/models"
	"github.com/H-ADJI/letsgo/internal/validator"
	"github.com/H-ADJI/letsgo/ui"
	"github.com/justinas/nosurf"
)

var funcMap = template.FuncMap{"humanDate": humanDate}

type userLoginForm struct {
	Email    string
	Password string
	validator.Validator
}
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
	Snippet   models.Snippet
	Snippets  []models.Snippet
	Form      any
	Flash     string
	IsAuth    bool
	CSRFToken string
}

func (a *app) NewTemplateData(r *http.Request) TemplateData {
	return TemplateData{
		Flash:     a.sessionManager.PopString(r.Context(), "flash"),
		IsAuth:    a.isAuthenticated(r),
		CSRFToken: nosurf.Token(r),
	}
}
func NewTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := fs.Glob(ui.Files, "html/pages/*.tmpl.html")
	if err != nil {
		return nil, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		files := []string{
			"html/base.tmpl.html",
			"html/partials/*.tmpl.html",
			page,
		}
		ts, err := template.New(name).Funcs(funcMap).ParseFS(ui.Files, files...)
		if err != nil {
			return nil, err
		}
		cache[name] = ts
	}
	return cache, nil
}
func humanDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.UTC().Format("02 Jan 2006 at 15:04")

}
