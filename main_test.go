package htmxheaders_test

import (
	"github.com/stretchr/testify/assert"
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

func TestSetResponseHeadersWithMultipleHeaders(t *testing.T) {
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

func TestRemoveHXHeaders(t *testing.T) {
	w := httptest.NewRecorder()
	w.Header().Set("HX-Location", "/some/location")
	w.Header().Set("HX-Push-Url", "/some/url")

	err := hh.RemoveHXHeaders(w)
	assert.NoError(t, err)

	headersToCheck := []string{"HX-Location", "HX-Push-Url", "HX-Redirect", "HX-Refresh", "HX-Replace-Url", "HX-Reswap", "HX-Retarget", "HX-Reselect"}
	for _, header := range headersToCheck {
		if w.Header().Get(header) != "" {
			t.Errorf("Expected header %s to be cleared, but it is still present", header)
		}
	}
}

func TestRemoveHXHeadersDoesNotAffectNonHXHeaders(t *testing.T) {
	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Cache-Control", "no-cache")

	err := hh.RemoveHXHeaders(w)
	assert.NoError(t, err)

	nonHXHeadersToCheck := []string{"Content-Type", "Cache-Control"}
	for _, header := range nonHXHeadersToCheck {
		if w.Header().Get(header) == "" {
			t.Errorf("Expected non-HTMX header %s to be preserved, but it is not present", header)
		}
	}
}
