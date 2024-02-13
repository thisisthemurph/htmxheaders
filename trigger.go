package htmxheaders

import (
	"encoding/json"
	"net/http"
	"strings"
)

// TriggerDelay allows distinguishing between different types of trigger headers.
// https://htmx.org/headers/hx-trigger/
type TriggerDelay int

const (
	TriggerImmediately TriggerDelay = iota // trigger events as soon as the response is received.
	TriggerAfterSettle                     // trigger events after the settle step.
	TriggerAfterSwap                       // trigger events after the swap step.
)

// String returns the name of the header associated with when the event should be triggered.
// Possible values are HX-Trigger, HX-Trigger-After-Settle, and HX-Trigger-After-Swap.
func (td TriggerDelay) String() string {
	switch td {
	case 1:
		return "HX-Trigger-After-Settle"
	case 2:
		return "HX-Trigger-After-Swap"
	default:
		return "HX-Trigger"
	}
}

// Trigger creates a DecoratorFunction that adds a trigger header to the response.
// The when parameter specifies when the event should be triggered (e.g., immediately, after settle, after swap).
// The eventName parameter specifies the name of the event(s) to be triggered.
// Multiple event names can be provided, separated by commas.
// https://htmx.org/headers/hx-trigger/
func Trigger(when TriggerDelay, eventName ...string) DecoratorFunction {
	events := strings.Join(eventName, ", ")
	return AddCustomHeader(when.String(), events)
}

// TriggerEvent represents an event that can be triggered with additional details.
// https://htmx.org/headers/hx-trigger/
type TriggerEvent struct {
	Name   string // Name of the event.
	Detail any    // Additional details associated with the event.
}

// TriggerWithDetail creates a DecoratorFunction that adds an event JSON object to the response headers.
// The JSON object contains a mapping of event names to their corresponding details.
// The when parameter specifies when the event should be triggered (e.g., immediately, after settle, after swap).
// The events parameter specifies a slice of TriggerEvent structs, each specifying an event name and its associated details.
// https://htmx.org/headers/hx-trigger/
func TriggerWithDetail(when TriggerDelay, events ...TriggerEvent) DecoratorFunction {
	return func(w http.ResponseWriter) error {
		eventMap := map[string]interface{}{}
		for _, event := range events {
			eventMap[event.Name] = event.Detail
		}

		jsonData, err := json.Marshal(eventMap)
		if err != nil {
			return err
		}

		w.Header().Set(when.String(), string(jsonData))
		return nil
	}
}
