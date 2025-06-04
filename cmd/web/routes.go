package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (a *app) routes() http.Handler {
	mux := http.NewServeMux()

	dynamic := alice.New(a.sessionManager.LoadAndSave, noSurf, a.authenticate)
	protected := dynamic.Append(a.requireAuth)

	mux.Handle("GET /{$}", dynamic.ThenFunc(a.home))

	snippetMux := http.NewServeMux()
	snippetMux.HandleFunc("GET /view/{id}", a.snippetView)
	snippetMux.Handle("POST /create", protected.ThenFunc(a.snippetCreatePost))
	snippetMux.Handle("GET /create", protected.ThenFunc(a.snippetCreate))
	mux.Handle("/snippet/", http.StripPrefix("/snippet", dynamic.Then(snippetMux)))

	userMux := http.NewServeMux()
	userMux.HandleFunc("GET /signup", a.userSignup)
	userMux.HandleFunc("POST /signup", a.userSignupPost)
	userMux.HandleFunc("GET /login", a.userLogin)
	userMux.HandleFunc("POST /login", a.userLoginPost)
	userMux.Handle("POST /logout", protected.ThenFunc(a.userLogoutPost))
	mux.Handle("/user/", http.StripPrefix("/user", dynamic.Then(userMux)))

	fileserver := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("GET /static/", http.StripPrefix("/static", disableDirList(fileserver)))

	standard := alice.New(a.recoverPanic, a.logRequest, commonHeaders)
	return standard.Then(mux)
}
