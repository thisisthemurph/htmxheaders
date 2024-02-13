package htmxheaders_test

import (
	hh "github.com/thisisthemurph/htmxheaders"
	"net/http/httptest"
	"testing"
)

func TestPushURL(t *testing.T) {
	w := httptest.NewRecorder()
	url := "/some/url"
	err := hh.SetResponseHeaders(w, hh.PushURL(url))

	if err != nil {
		t.Errorf("PushURL returned an unexpected error: %v", err)
	}

	header := w.Header().Get("HX-Push-Url")
	if header != url {
		t.Errorf("Expected header HX-Push-Url to have value %s, got %s", url, header)
	}
}

func TestPreventPushURL(t *testing.T) {
	w := httptest.NewRecorder()
	err := hh.SetResponseHeaders(w, hh.PreventPushURL())

	if err != nil {
		t.Errorf("PreventPushURL returned an unexpected error: %v", err)
	}

	want := "false"
	header := w.Header().Get("HX-Push-Url")
	if header != want {
		t.Errorf("Expected header HX-Push-Url to have value %s, got %s", want, header)
	}
}
