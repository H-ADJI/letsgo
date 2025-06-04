package main

import (
	"net/http"

	"github.com/H-ADJI/letsgo/ui"
	"github.com/justinas/alice"
)

func (a *app) routes() http.Handler {
	mux := http.NewServeMux()

	dynamic := alice.New(a.sessionManager.LoadAndSave, noSurf, a.authenticate)
	protected := dynamic.Append(a.requireAuth)

	mux.Handle("GET /{$}", dynamic.ThenFunc(a.home))

	snippetMux := http.NewServeMux()
	snippetMux.Handle("GET /view/{id}", dynamic.ThenFunc(a.snippetView))
	snippetMux.Handle("POST /create", protected.ThenFunc(a.snippetCreatePost))
	snippetMux.Handle("GET /create", protected.ThenFunc(a.snippetCreate))
	mux.Handle("/snippet/", http.StripPrefix("/snippet", snippetMux))

	userMux := http.NewServeMux()
	userMux.Handle("GET /signup", dynamic.ThenFunc(a.userSignup))
	userMux.Handle("POST /signup", dynamic.ThenFunc(a.userSignupPost))
	userMux.Handle("GET /login", dynamic.ThenFunc(a.userLogin))
	userMux.Handle("POST /login", dynamic.ThenFunc(a.userLoginPost))
	userMux.Handle("POST /logout", protected.ThenFunc(a.userLogoutPost))
	mux.Handle("/user/", http.StripPrefix("/user", userMux))

	fileserver := http.FileServerFS(ui.Files)
	mux.Handle("GET /static/", disableDirList(fileserver))

	standard := alice.New(a.recoverPanic, a.logRequest, commonHeaders)
	return standard.Then(mux)
}
