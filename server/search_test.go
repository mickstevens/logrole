package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

const mms = "MM89a8c4a6891c53054e9cd604922bfb61"

var searchTests = []struct {
	in       string
	code     int
	location string
}{
	{"/search?q=" + mms, 301, "/messages/" + mms},
	{"/search?", 302, "/"},
	{"/search?q=unknown", 302, "/"},
}

func TestSearchRedirects(t *testing.T) {
	t.Parallel()
	s := &searchServer{}
	for _, tt := range searchTests {
		req, _ := http.NewRequest("GET", tt.in, nil)
		w := httptest.NewRecorder()
		s.ServeHTTP(w, req)
		if w.Code != tt.code {
			t.Errorf("expected Code to equal %d, got %d", tt.code, w.Code)
		}
		if location := w.Header().Get("Location"); location != tt.location {
			t.Errorf("expected Location to equal %s, got %s", tt.location, location)
		}
	}
}