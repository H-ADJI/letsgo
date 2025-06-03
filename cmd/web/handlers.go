package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/H-ADJI/letsgo/internal/models"
	"github.com/H-ADJI/letsgo/internal/validator"
)

func (a *app) home(w http.ResponseWriter, r *http.Request) {
	snippets, err := a.snippets.Latest()
	if err != nil {
		a.serverError(w, r, err)
		return
	}
	data := a.NewTemplateData(r)
	data.Snippets = snippets
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
	data := a.NewTemplateData(r)
	data.Snippet = snippet
	a.render(w, r, http.StatusOK, "view.tmpl.html", data)
}
func (a *app) snippetCreate(w http.ResponseWriter, r *http.Request) {
	data := a.NewTemplateData(r)
	data.Form = snippetCreateForm{Expires: 1}
	a.render(
		w,
		r,
		http.StatusOK,
		"create.tmpl.html",
		data,
	)
}
func (a *app) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		a.clientError(w, http.StatusBadRequest)
		return
	}
	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")
	expires, err := strconv.Atoi(r.PostForm.Get("expires"))
	if err != nil {
		a.clientError(w, http.StatusBadRequest)
		return
	}
	form := snippetCreateForm{}
	form.CheckField(
		validator.MaxChars(title, 100),
		"title",
		"Title can not be more than 100 characters",
	)
	form.CheckField(
		validator.NotBlank(title),
		"title",
		"Title can not be blank",
	)
	form.CheckField(
		validator.NotBlank(content),
		"content",
		"Content can not be blank",
	)
	form.CheckField(
		validator.PermittedValues(expires, 1, 7, 365),
		"expires",
		"Field should be equale to 1, 7 or 365",
	)
	form.Title = title
	form.Content = content
	form.Expires = expires

	if !form.IsValid() {
		data := TemplateData{}
		data.Form = form
		a.render(w, r, http.StatusUnprocessableEntity, "create.tmpl.html", data)
		return
	}
	id, err := a.snippets.Insert(title, content, expires)
	if err != nil {
		a.serverError(w, r, err)
		return
	}
	a.sessionManager.Put(r.Context(), "flash", "Snippet successfully created !")
	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
func (a *app) userSignup(w http.ResponseWriter, r *http.Request) {
	a.render(w, r, http.StatusOK, "signup.tmpl.html", TemplateData{Form: userSignupForm{}})
}
func (a *app) userSignupPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		a.clientError(w, http.StatusBadRequest)
		return
	}
	name := r.PostForm.Get("name")
	email := r.PostForm.Get("email")
	password := r.PostForm.Get("password")
	form := userSignupForm{}
	form.Name = name
	form.Email = email
	form.Password = password

	form.CheckField(validator.NotBlank(name), "name", "this field can not be blank")
	form.CheckField(validator.NotBlank(email), "email", "this field can not be blank")
	form.CheckField(
		validator.Matches(email, validator.EmailRX),
		"email",
		"This field must be a valid email address",
	)
	form.CheckField(validator.NotBlank(password), "password", "this field can not be blank")
	form.CheckField(
		validator.MinChars(password, 8),
		"password",
		"this field can not be at least 8 characters long",
	)
	if !form.IsValid() {
		data := a.NewTemplateData(r)
		data.Form = form
		a.render(w, r, http.StatusUnprocessableEntity, "signup.tmpl.html", data)
		return
	}
	err = a.users.Insert(form.Name, form.Email, form.Password)
	if err != nil {
		if errors.Is(err, models.ErrDuplicateEmail) {
			form.AddFieldError("email", "Email already used")
			data := a.NewTemplateData(r)
			data.Form = form
			a.render(w, r, http.StatusUnprocessableEntity, "signup.tmpl.html", data)
		} else {
			a.serverError(w, r, err)
		}
		return
	}
	a.sessionManager.Put(r.Context(), "flash", "Your signup was successful. Please log in.")

	// And redirect the user to the login page.
	http.Redirect(w, r, "/user/login", http.StatusSeeOther)
}
func (a *app) userLogin(w http.ResponseWriter, r *http.Request) {
	a.render(w, r, http.StatusOK, "login.tmpl.html", TemplateData{Form: userLoginForm{}})
}
func (a *app) userLoginPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		a.clientError(w, http.StatusBadRequest)
		return
	}
	email := r.PostForm.Get("email")
	password := r.PostForm.Get("password")
	form := userLoginForm{}
	form.Email = email
	form.Password = password
	form.CheckField(validator.NotBlank(password), "password", "this field can not be blank")
	form.CheckField(validator.NotBlank(email), "email", "this field can not be blank")
	form.CheckField(
		validator.Matches(email, validator.EmailRX),
		"email",
		"This field must be a valid email address",
	)
	if !form.IsValid() {
		data := a.NewTemplateData(r)
		data.Form = form
		a.render(w, r, http.StatusUnprocessableEntity, "login.tmpl.html", data)
		return
	}
	id, err := a.users.Authenticate(email, password)
	if err != nil {
		if errors.Is(err, models.ErrInvalidCreds) {
			form.AddNonFieldError("Email or Password is incorrect")
			data := a.NewTemplateData(r)
			data.Form = form
			a.render(w, r, http.StatusUnprocessableEntity, "login.tmpl.html", data)
		} else {
			a.serverError(w, r, err)
		}
		return
	}
	err = a.sessionManager.RenewToken(r.Context())
	if err != nil {
		a.serverError(w, r, err)
	}
	a.sessionManager.Put(r.Context(), "authenticatedUserId", id)
	http.Redirect(w, r, "/snippet/create", http.StatusSeeOther)
}
func (a *app) userLogoutPost(w http.ResponseWriter, r *http.Request) {
	err := a.sessionManager.RenewToken(r.Context())
	if err != nil {
		a.serverError(w, r, err)
		return
	}
	a.sessionManager.Remove(r.Context(), "authenticatedUserId")
	a.sessionManager.Put(r.Context(), "flash", "You've been logged out")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
