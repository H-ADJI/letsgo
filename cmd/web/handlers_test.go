package main

import (
	"bytes"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/H-ADJI/letsgo/internal/assert"
)

func TestPing(t *testing.T) {
	a := &app{
		logger: slog.New(slog.NewTextHandler(io.Discard, nil)),
	}
	ts := httptest.NewTLSServer(a.routes())
	defer ts.Close()
	res, err := ts.Client().Get(ts.URL + "/ping")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, res.StatusCode, http.StatusOK)
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	body = bytes.TrimSpace([]byte(body))
	assert.Equal(t, string(body), "OK")

}
