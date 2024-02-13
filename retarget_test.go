package htmxheaders_test

import (
	hh "github.com/thisisthemurph/htmxheaders"
	"net/http/httptest"
	"testing"
)

func TestRetarget(t *testing.T) {
	w := httptest.NewRecorder()
	target := "#swap-target"
	err := hh.SetResponseHeaders(w, hh.Retarget(target))

	if err != nil {
		t.Errorf("Refresh returned an unexpected error: %v", err)
	}

	header := w.Header().Get("HX-Retarget")
	if header != target {
		t.Errorf("Expected header HX-Refresh to have value %s, got %s", target, header)
	}
}
