package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
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
	fmt.Fprintf(w, "Display a snippet with id : %d ...", id)
}
func (a *app) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a snippet using a form"))
}
func (a *app) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	w.Header().Add("Cache-Control", "public")
	w.Header().Add("Cache-Control", "max-age=31536000")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Save new snippet"))
}
