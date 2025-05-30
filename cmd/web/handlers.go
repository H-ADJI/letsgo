package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/H-ADJI/letsgo/internal/models"
)

func (a *app) home(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles(
		"./ui/html/pages/home.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/base.tmpl.html",
	)
	if err != nil {
		a.serverError(w, r, err)
		return
	}
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		a.serverError(w, r, err)
		return
	}
}
func (a *app) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	snippet, err := a.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecords) {
			http.NotFound(w, r)
		} else {
			a.serverError(w, r, err)
		}
		return
	}

	fmt.Fprintf(w, "%+v", snippet)
}
func (a *app) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a snippet using a form"))
}
func (a *app) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	title := "O, snail"
	content := "O, snail\n Climb very fast\n or you'll get crushed"
	expires := 7

	id, err := a.snippets.Insert(title, content, expires)
	if err != nil {
		a.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
