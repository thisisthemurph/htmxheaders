package htmxheaders_test

import (
	hh "github.com/thisisthemurph/htmxheaders"
	"net/http/httptest"
	"testing"
)

func TestReplaceURL(t *testing.T) {
	w := httptest.NewRecorder()
	url := "https://someurl.com"
	err := hh.SetResponseHeaders(w, hh.ReplaceURL(url))

	if err != nil {
		t.Errorf("Refresh returned an unexpected error: %v", err)
	}

	header := w.Header().Get("HX-Replace-Url")
	if header != url {
		t.Errorf("Expected header HX-Refresh to have value %s, got %s", url, header)
	}
}

func TestPreventReplaceURL(t *testing.T) {
	w := httptest.NewRecorder()
	err := hh.SetResponseHeaders(w, hh.PreventReplaceURL())

	if err != nil {
		t.Errorf("Refresh returned an unexpected error: %v", err)
	}

	want := "false"
	header := w.Header().Get("HX-Replace-Url")
	if header != want {
		t.Errorf("Expected header HX-Refresh to have value %s, got %s", want, header)
	}
}
