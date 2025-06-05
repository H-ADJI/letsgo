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
