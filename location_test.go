package htmxheaders_test

import (
	"encoding/json"
	hh "github.com/thisisthemurph/htmxheaders"
	"net/http/httptest"
	"testing"
)

func TestLocation(t *testing.T) {
	w := httptest.NewRecorder()
	location := "/some/path"
	err := hh.SetResponseHeaders(w, hh.Location(location))

	if err != nil {
		t.Errorf("Location returned an unexpected error: %v", err)
	}

	header := w.Header().Get("HX-Location")
	if header != location {
		t.Errorf("Expected header HX-Location to have value %s, got %s", location, header)
	}
}

func TestLocationWithContext(t *testing.T) {
	w := httptest.NewRecorder()
	path := "/some/path"
	context := hh.LocationContext{Target: "#my-target", Swap: hh.SwapAfterBegin}
	err := hh.SetResponseHeaders(w, hh.LocationWithContext(path, context))

	if err != nil {
		t.Errorf("LocationWithContext returned an unexpected error: %v", err)
	}

	var data hh.LocationContextWithPath
	err = json.Unmarshal([]byte(w.Header().Get("HX-Location")), &data)
	if err != nil {
		t.Errorf("Error unmarshalling HX-Location JSON: %v", err)
	}

	if data.Path != path {
		t.Errorf("Expected path: %s, got path: %s", path, data.Path)
	}

	if data.Target != context.Target {
		t.Errorf("Expected target: %s, got target: %s", context.Target, data.Target)
	}

	if data.Swap != context.Swap {
		t.Errorf("Expected swap: %s, got swap: $%s", context.Swap, data.Swap)
	}
}
