package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (a app) routes() http.Handler {
	mux := http.NewServeMux()

	dynamic := alice.New(a.sessionManager.LoadAndSave)
	mux.Handle("GET /{$}", dynamic.ThenFunc(a.home))
	mux.Handle("GET /snippet/view/{id}", dynamic.ThenFunc(a.snippetView))
	mux.Handle("GET /snippet/create", dynamic.ThenFunc(a.snippetCreate))
	mux.Handle("POST /snippet/create", dynamic.ThenFunc(a.snippetCreatePost))

	fileserver := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("GET /static/", http.StripPrefix("/static", disableDirList(fileserver)))

	standard := alice.New(a.recoverPanic, a.logRequest, commonHeaders)
	return standard.Then(mux)
}
