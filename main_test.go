package htmxheaders_test

import (
	"encoding/json"
	hh "github.com/thisisthemurph/htmxheaders"
	"net/http/httptest"
	"testing"
)

func TestAddCustomHeader(t *testing.T) {
	w := httptest.NewRecorder()
	key := "Custom-Header"
	value := "TestValue"
	err := hh.AddCustomHeader(key, value)(w)

	if err != nil {
		t.Errorf("AddCustomHeader returned an unexpected error: %v", err)
	}

	header := w.Header().Get(key)
	if header != value {
		t.Errorf("Expected header %s to have value %s, got %s", key, value, header)
	}
}

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
	context := hh.LocationContext{Target: "#my-target"}
	err := hh.SetResponseHeaders(w, hh.LocationWithContext(path, context))

	if err != nil {
		t.Errorf("LocationWithContext returned an unexpected error: %v", err)
	}

	var data hh.LocationContextWithPath
	err = json.Unmarshal([]byte(w.Header().Get("HX-Location")), &data)
	if err != nil {
		t.Errorf("Error unmarshalling HX-Location JSON: %v", err)
	}

	if data.Path != path || data.Target != context.Target {
		t.Errorf("Expected path: %s, Method: %s, got path: %s, Method: %s", path, context.Target, data.Path, data.Target)
	}
}

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

func TestReswap(t *testing.T) {
	w := httptest.NewRecorder()
	swap := hh.SwapAfterBegin
	err := hh.SetResponseHeaders(w, hh.Reswap(swap))

	if err != nil {
		t.Errorf("Refresh returned an unexpected error: %v", err)
	}

	header := w.Header().Get("HX-Reswap")
	headerSwap, _ := hh.StringToSwap(header)
	if headerSwap != swap {
		t.Errorf("Expected header HX-Refresh to have value %s, got %s", swap, header)
	}
}

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

func TestAddHeadersWithMultipleHeaders(t *testing.T) {
	w := httptest.NewRecorder()
	target := "#swap-target"
	swap := hh.SwapAfterBegin

	err := hh.SetResponseHeaders(w, hh.Reswap(swap), hh.Retarget(target))

	if err != nil {
		t.Errorf("Operation returned undexpected error: %v", err)
	}

	reswapHeader := w.Header().Get("HX-Reswap")
	retargetHeader := w.Header().Get("HX-Retarget")
	if retargetHeader != target || reswapHeader != swap.String() {
		t.Errorf("Expect header HX-Retarget to have value %s and HX-Reswap to have %s, got %s and %s", target, swap.String(), retargetHeader, reswapHeader)
	}
}

func TestClearHXHeadersRemovesHXHeaders(t *testing.T) {
	w := httptest.NewRecorder()
	w.Header().Set("HX-Location", "/some/location")
	w.Header().Set("HX-Push-Url", "/some/url")

	hh.RemoveHXHeaders(w)

	headersToCheck := []string{"HX-Location", "HX-Push-Url", "HX-Redirect", "HX-Refresh", "HX-Replace-Url", "HX-Reswap", "HX-Retarget", "HX-Reselect"}
	for _, header := range headersToCheck {
		if w.Header().Get(header) != "" {
			t.Errorf("Expected header %s to be cleared, but it is still present", header)
		}
	}
}

func TestClearHXHeadersDoesNotAffectNonHXHeaders(t *testing.T) {
	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Cache-Control", "no-cache")

	hh.RemoveHXHeaders(w)

	nonHXHeadersToCheck := []string{"Content-Type", "Cache-Control"}
	for _, header := range nonHXHeadersToCheck {
		if w.Header().Get(header) == "" {
			t.Errorf("Expected non-HTMX header %s to be preserved, but it is not present", header)
		}
	}
}
