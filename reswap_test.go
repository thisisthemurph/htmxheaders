package htmxheaders_test

import (
	hh "github.com/thisisthemurph/htmxheaders"
	"net/http/httptest"
	"testing"
)

func TestSwapString(t *testing.T) {
	tests := []struct {
		swap     hh.Swap
		expected string
	}{
		{hh.SwapInnerHTML, "innerHTML"},
		{hh.SwapOuterHTML, "outerHTML"},
		{hh.SwapBeforeBegin, "beforebegin"},
		{hh.SwapAfterBegin, "afterbegin"},
		{hh.SwapBeforeEnd, "beforeend"},
		{hh.SwapAfterEnd, "afterend"},
		{hh.SwapDelete, "delete"},
		{hh.SwapNone, "none"},
	}

	for _, test := range tests {
		result := test.swap.String()
		if result != test.expected {
			t.Errorf("Expected %s for Swap %d, got %s", test.expected, test.swap, result)
		}
	}
}

func TestSwapFromString(t *testing.T) {
	tests := []struct {
		str      string
		expected hh.Swap
	}{
		{"innerHTML", hh.SwapInnerHTML},
		{"outerHTML", hh.SwapOuterHTML},
		{"beforebegin", hh.SwapBeforeBegin},
		{"afterbegin", hh.SwapAfterBegin},
		{"beforeend", hh.SwapBeforeEnd},
		{"afterend", hh.SwapAfterEnd},
		{"delete", hh.SwapDelete},
		{"none", hh.SwapNone},
	}

	for _, test := range tests {
		result, err := hh.SwapFromString(test.str)
		if err != nil {
			t.Errorf("Unexpected error for string %s: %v", test.str, err)
		}

		if result != test.expected {
			t.Errorf("Expected %d for string %s, got %d", test.expected, test.str, result)
		}
	}
}

func TestSwapFromStringWithInvalidValid(t *testing.T) {
	tests := []struct {
		str      string
		expected hh.Swap
	}{
		{"invalid", hh.SwapInnerHTML},
		{"unsupported", hh.SwapInnerHTML},
		{"HX-Target", hh.SwapInnerHTML},
	}

	for _, test := range tests {
		result, err := hh.SwapFromString(test.str)
		if err == nil {
			t.Errorf("Expected error for string %s: got nil", test.str)
		}

		// Testing default case
		if result != test.expected {
			t.Errorf("Expected %d for string %s, got %d", test.expected, test.str, result)
		}
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
