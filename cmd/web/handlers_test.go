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
