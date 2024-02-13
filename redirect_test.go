package htmxheaders_test

import (
	hh "github.com/thisisthemurph/htmxheaders"
	"net/http/httptest"
	"testing"
)

func TestRedirect(t *testing.T) {
	w := httptest.NewRecorder()
	url := "/some/url"
	err := hh.SetResponseHeaders(w, hh.Redirect(url))

	if err != nil {
		t.Errorf("Redirect returned an unexpected error: %v", err)
	}

	header := w.Header().Get("HX-Redirect")
	if header != url {
		t.Errorf("Expected header HX-Push-Url to have value %s, got %s", url, header)
	}
}
