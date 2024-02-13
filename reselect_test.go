package htmxheaders_test

import (
	hh "github.com/thisisthemurph/htmxheaders"
	"net/http/httptest"
	"testing"
)

func TestReselect(t *testing.T) {
	w := httptest.NewRecorder()
	selector := "section.my-class"
	err := hh.SetResponseHeaders(w, hh.Reselect(selector))

	if err != nil {
		t.Errorf("Refresh returned an unexpected error: %v", err)
	}

	header := w.Header().Get("HX-Reselect")
	if header != selector {
		t.Errorf("Expected header HX-Reselect to have value %s, got %s", selector, header)
	}
}
