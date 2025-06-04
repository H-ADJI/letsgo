package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func (a *app) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)
	a.logger.Error(err.Error(), "method", method, "uri", uri)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (a *app) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (a *app) render(
	w http.ResponseWriter,
	r *http.Request,
	status int,
	page string,
	data TemplateData,
) {

	ts, ok := a.templateCache[page]
	if !ok {
		a.serverError(w, r, fmt.Errorf("The template %s does not exist", page))
		return
	}
	buf := &bytes.Buffer{}
	err := ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		a.serverError(w, r, err)
		return
	}
	w.WriteHeader(status)
	buf.WriteTo(w)
}
func (a *app) isAuthenticated(r *http.Request) bool {
	isAuth, ok := r.Context().Value(isAuthenticatedContextKey).(bool)
	if !ok {
		return false
	}
	fmt.Println(isAuth)
	return isAuth
}
