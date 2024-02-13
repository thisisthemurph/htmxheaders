package htmxheaders_test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestTrigger(t *testing.T) {
	w := httptest.NewRecorder()
	eventName := "testEvent"

	err := hh.SetResponseHeaders(w, hh.Trigger(eventName))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	actualHeader := w.Header().Get("HX-Trigger")
	if actualHeader != eventName {
		t.Errorf("Expected header HX-Trigger to have value: %s, got: %s instead", eventName, actualHeader)
	}
}

func TestTriggerWitMultipleEvents(t *testing.T) {
	w := httptest.NewRecorder()
	event1 := "event1"
	event2 := "event2"
	event3 := "event3"

	err := hh.SetResponseHeaders(w, hh.Trigger(event1, event2, event3))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expectedHeader := "event1, event2, event3"
	actualHeader := w.Header().Get("HX-Trigger")
	if actualHeader != expectedHeader {
		t.Errorf("Expected header HX-Trigger to have value: %q, got: %q instead", expectedHeader, actualHeader)
	}
}

func TestTriggerWithDetails(t *testing.T) {
	testCases := []struct {
		Name  string
		Event hh.TriggerEvent
	}{
		{
			Name: "Single event with string details",
			Event: hh.TriggerEvent{
				Name: "event1", Detail: "details1",
			},
		},
		{
			Name: "Single event with int details",
			Event: hh.TriggerEvent{
				Name: "event2", Detail: 123,
			},
		},
		{
			Name: "Single event with map details",
			Event: hh.TriggerEvent{
				Name: "event3", Detail: map[string]interface{}{"key": "value"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			// Prepare a dummy HTTP response writer
			w := httptest.NewRecorder()

			// Create a DecoratorFunction with TriggerWithDetail
			err := hh.SetResponseHeaders(w, hh.TriggerWithDetail(tc.Event))
			assert.NoError(t, err)

			// Check if the HX-Trigger header is set correctly
			expectedHeader := map[string]interface{}{tc.Event.Name: tc.Event.Detail}
			expectedHeaderValue, err := json.Marshal(expectedHeader)
			assert.NoError(t, err)
			actualHeaderValue := w.Header().Get("HX-Trigger")

			require.JSONEq(t, string(expectedHeaderValue), actualHeaderValue)

		})
	}
}

func TestTriggerWithDetailsHandlesMultipleEvents(t *testing.T) {
	events := []hh.TriggerEvent{
		hh.TriggerEvent{
			Name:   "event1",
			Detail: "details1",
		},
		hh.TriggerEvent{
			Name:   "event2",
			Detail: 123,
		},
		hh.TriggerEvent{
			Name:   "event3",
			Detail: map[string]interface{}{"key": "value"},
		},
	}

	w := httptest.NewRecorder()
	err := hh.SetResponseHeaders(w, hh.TriggerWithDetail(events...))
	assert.NoError(t, err)

	expectedJSON := `{"event1": "details1", "event2": 123, "event3": {"key": "value"}}`
	actualHeader := w.Header().Get("HX-Trigger")
	require.JSONEq(t, expectedJSON, actualHeader)
}

func TestRemoveHXHeadersRemovesHXHeaders(t *testing.T) {
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
