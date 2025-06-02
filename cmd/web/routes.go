package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (a app) routes() http.Handler {
	mux := http.NewServeMux()

	dynamic := alice.New(a.sessionManager.LoadAndSave)

	mux.Handle("GET /{$}", dynamic.ThenFunc(a.home))

	snippetMux := http.NewServeMux()
	snippetMux.HandleFunc("GET /view/{id}", a.snippetView)
	snippetMux.HandleFunc("POST /create", a.snippetCreatePost)
	snippetMux.HandleFunc("GET /create", a.snippetCreate)
	mux.Handle("/snippet/", http.StripPrefix("/snippet", dynamic.Then(snippetMux)))

	fileserver := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("GET /static/", http.StripPrefix("/static", disableDirList(fileserver)))

	standard := alice.New(a.recoverPanic, a.logRequest, commonHeaders)
	return standard.Then(mux)
}
