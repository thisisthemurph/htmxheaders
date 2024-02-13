package htmxheaders_test

import (
	hh "github.com/thisisthemurph/htmxheaders"
	"net/http/httptest"
	"testing"
)

func TestRefresh(t *testing.T) {
	w := httptest.NewRecorder()
	err := hh.SetResponseHeaders(w, hh.Refresh())

	if err != nil {
		t.Errorf("Refresh returned an unexpected error: %v", err)
	}

	want := "true"
	header := w.Header().Get("HX-Refresh")
	if header != want {
		t.Errorf("Expected header HX-Refresh to have value %s, got %s", want, header)
	}
}

func TestPreventRefresh(t *testing.T) {
	w := httptest.NewRecorder()
	err := hh.SetResponseHeaders(w, hh.PreventRefresh())

	if err != nil {
		t.Errorf("Refresh returned an unexpected error: %v", err)
	}

	want := "false"
	header := w.Header().Get("HX-Refresh")
	if header != want {
		t.Errorf("Expected header HX-Refresh to have value %s, got %s", want, header)
	}
}
