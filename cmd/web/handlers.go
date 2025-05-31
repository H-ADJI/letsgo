package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/H-ADJI/letsgo/internal/models"
)

func (a *app) home(w http.ResponseWriter, r *http.Request) {
	snippets, err := a.snippets.Latest()
	if err != nil {
		a.serverError(w, r, err)
		return
	}
	data := TemplateData{Snippets: snippets}
	a.render(w, r, http.StatusOK, "home.tmpl.html", data)
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
	data := TemplateData{Snippet: snippet}
	a.render(w, r, http.StatusOK, "view.tmpl.html", data)
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
