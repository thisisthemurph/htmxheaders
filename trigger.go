package htmxheaders

import (
	"encoding/json"
	"net/http"
	"strings"
)

func Trigger(eventName ...string) DecoratorFunction {
	events := strings.Join(eventName, ", ")
	return AddCustomHeader("HX-Trigger", events)
}

// TriggerEvent represents an event that can be triggered with additional details.
// https://htmx.org/headers/hx-trigger/
type TriggerEvent struct {
	Name   string // Name of the event.
	Detail any    // Additional details associated with the event.
}

// TriggerWithDetail adds an event JSON object to the response headers.
// The JSON object contains a mapping of event names to their corresponding details.
// Each TriggerEvent in the arguments specifies an event name and its associated details.
// https://htmx.org/headers/hx-trigger/
func TriggerWithDetail(events ...TriggerEvent) DecoratorFunction {
	return func(w http.ResponseWriter) error {
		eventMap := map[string]interface{}{}
		for _, event := range events {
			eventMap[event.Name] = event.Detail
		}

		jsonData, err := json.Marshal(eventMap)
		if err != nil {
			return err
		}

		w.Header().Set("HX-Trigger", string(jsonData))
		return nil
	}
}
