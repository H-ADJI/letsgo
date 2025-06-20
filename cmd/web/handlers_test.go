package main

import (
	"net/http"
	"testing"

	"github.com/H-ADJI/letsgo/internal/assert"
)

func TestPing(t *testing.T) {
	a := newTestApp(t)
	ts := newTestServer(t, a.routes())
	defer ts.Close()
	statusCode, _, body := ts.get(t, "/ping")
	assert.Equal(t, statusCode, http.StatusOK)
	assert.Equal(t, body, "OK")
}

func TestSnippetView(t *testing.T) {
	a := newTestApp(t)
	ts := newTestServer(t, a.routes())
	defer ts.Close()
	tests := []struct {
		name     string
		urlPath  string
		wantCode int
		wantBody string
	}{
		{
			name:     "Valid ID",
			urlPath:  "/snippet/view/1",
			wantCode: http.StatusOK,
			wantBody: "A test content ....",
		},
		{
			name:     "Non-existent ID",
			urlPath:  "/snippet/view/2",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "Negative ID",
			urlPath:  "/snippet/view/-1",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "Decimal ID",
			urlPath:  "/snippet/view/1.23",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "String ID",
			urlPath:  "/snippet/view/foo",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "Empty ID",
			urlPath:  "/snippet/view/",
			wantCode: http.StatusNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			statusCode, _, body := ts.get(t, tc.urlPath)
			assert.Equal(t, statusCode, tc.wantCode)
			if tc.wantBody != "" {
				assert.StringContains(t, body, tc.wantBody)
			}
		})
	}
}

func TestUserSignup(t *testing.T) {
	a := newTestApp(t)
	ts := newTestServer(t, a.routes())
	defer ts.Close()
	_, _, body := ts.get(t, "/user/signup")
	validCSRFToken := extractCSRFToken(t, body)

	t.Logf("token : %v", validCSRFToken)
	// const (
	// 	validName     = "Bob"
	// 	validPassword = "validPa$$word"
	// 	validEmail    = "bob@example.com"
	// 	formTag       = "<form action='/user/signup' method='POST' novalidate>"
	// )
	//
	// tests := []struct {
	// 	name         string
	// 	userName     string
	// 	userEmail    string
	// 	userPassword string
	// 	csrfToken    string
	// 	wantCode     int
	// 	wantFormTag  string
	// }{
	// 	{
	// 		name:         "Valid submission",
	// 		userName:     validName,
	// 		userEmail:    validEmail,
	// 		userPassword: validPassword,
	// 		csrfToken:    validCSRFToken,
	// 		wantCode:     http.StatusSeeOther,
	// 	},
	// 	{
	// 		name:         "Invalid CSRF Token",
	// 		userName:     validName,
	// 		userEmail:    validEmail,
	// 		userPassword: validPassword,
	// 		csrfToken:    "wrongToken",
	// 		wantCode:     http.StatusBadRequest,
	// 	},
	// 	{
	// 		name:         "Empty name",
	// 		userName:     "",
	// 		userEmail:    validEmail,
	// 		userPassword: validPassword,
	// 		csrfToken:    validCSRFToken,
	// 		wantCode:     http.StatusUnprocessableEntity,
	// 		wantFormTag:  formTag,
	// 	},
	// 	{
	// 		name:         "Empty email",
	// 		userName:     validName,
	// 		userEmail:    "",
	// 		userPassword: validPassword,
	// 		csrfToken:    validCSRFToken,
	// 		wantCode:     http.StatusUnprocessableEntity,
	// 		wantFormTag:  formTag,
	// 	},
	// 	{
	// 		name:         "Empty password",
	// 		userName:     validName,
	// 		userEmail:    validEmail,
	// 		userPassword: "",
	// 		csrfToken:    validCSRFToken,
	// 		wantCode:     http.StatusUnprocessableEntity,
	// 		wantFormTag:  formTag,
	// 	},
	// 	{
	// 		name:         "Invalid email",
	// 		userName:     validName,
	// 		userEmail:    "bob@example.",
	// 		userPassword: validPassword,
	// 		csrfToken:    validCSRFToken,
	// 		wantCode:     http.StatusUnprocessableEntity,
	// 		wantFormTag:  formTag,
	// 	},
	// 	{
	// 		name:         "Short password",
	// 		userName:     validName,
	// 		userEmail:    validEmail,
	// 		userPassword: "pa$$",
	// 		csrfToken:    validCSRFToken,
	// 		wantCode:     http.StatusUnprocessableEntity,
	// 		wantFormTag:  formTag,
	// 	},
	// 	{
	// 		name:         "Duplicate email",
	// 		userName:     validName,
	// 		userEmail:    "dupe@example.com",
	// 		userPassword: validPassword,
	// 		csrfToken:    validCSRFToken,
	// 		wantCode:     http.StatusUnprocessableEntity,
	// 		wantFormTag:  formTag,
	// 	},
	// }
	//
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		form := url.Values{}
	// 		form.Add("name", tt.userName)
	// 		form.Add("email", tt.userEmail)
	// 		form.Add("password", tt.userPassword)
	// 		form.Add("csrf_token", tt.csrfToken)
	//
	// 		code, _, body := ts.postForm(t, "/user/signup", form)
	//
	// 		assert.Equal(t, code, tt.wantCode)
	//
	// 		if tt.wantFormTag != "" {
	// 			assert.StringContains(t, body, tt.wantFormTag)
	// 		}
	// 	})
	// }
}
