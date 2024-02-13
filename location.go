package htmxheaders

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// LocationContext represents additional optional context that can
// be provided within the LocationWithContext method
// https://htmx.org/headers/hx-location/
type LocationContext struct {
	Source  string `json:"source,omitempty"`  // the source element of the request
	Event   string `json:"event,omitempty"`   // an event that “triggered” the request
	Handler string `json:"handler,omitempty"` // a callback that will handle the response HTML
	Target  string `json:"target,omitempty"`  // the target to swap the response into
	Swap    Swap   `json:"swap,omitempty"`    // how the response will be swapped in relative to the target
	Values  string `json:"values,omitempty"`  // values to submit with the request
	Select  string `json:"select,omitempty"`  // allows you to select the content you want swapped from a response
}

type LocationContextWithPath struct {
	LocationContext
	Path string `json:"path"`
}

// Location allows you to do a client-side redirect that does not do a full page reload.
// https://htmx.org/headers/hx-location/
func Location(location string) DecoratorFunction {
	return AddCustomHeader("HX-Location", location)
}

// LocationWithContext allows you to do a client-side redirect that does not do a full page reload.
// additional options are provided in the context.
// https://htmx.org/headers/hx-location/
func LocationWithContext(path string, context LocationContext) DecoratorFunction {
	return func(w http.ResponseWriter) error {
		data, err := json.Marshal(
			LocationContextWithPath{
				LocationContext: context,
				Path:            path,
			})

		if err != nil {
			return fmt.Errorf("error marshalling context JSON: %v", err)
		}

		w.Header().Set("HX-Location", string(data))
		return nil
	}
}
