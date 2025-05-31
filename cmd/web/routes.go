package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (a app) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", a.home)
	mux.HandleFunc("GET /snippet/view/{id}", a.snippetView)
	mux.HandleFunc("GET /snippet/create", a.snippetCreate)
	mux.HandleFunc("POST /snippet/create", a.snippetCreatePost)
	fileserver := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("GET /static/", http.StripPrefix("/static", disableDirList(fileserver)))
	standard := alice.New(a.recoverPanic, a.logRequest, commonHeaders)
	return standard.Then(mux)
}
