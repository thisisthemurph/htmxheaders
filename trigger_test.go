package htmxheaders_test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	hh "github.com/thisisthemurph/htmxheaders"
	"net/http/httptest"
	"testing"
)

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
