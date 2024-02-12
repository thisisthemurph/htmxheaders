package htmxheaders

// LocationContext represents additional optional context that can
// be provided within the LocationWithContext method
type LocationContext struct {
	// the source element of the request
	Source string `json:"source,omitempty"`
	// an event that “triggered” the request
	Event string `json:"event,omitempty"`
	// a callback that will handle the response HTML
	Handler string `json:"handler,omitempty"`
	// the target to swap the response into
	Target string `json:"target,omitempty"`
	// how the response will be swapped in relative to the target
	Swap string `json:"swap,omitempty"`
	// values to submit with the request
	Values string `json:"values,omitempty"`
	// headers to submit with the request
	Headers string `json:"headers,omitempty"`
	// allows you to select the content you want swapped from a response
	Select string `json:"select,omitempty"`
}

type LocationContextWithPath struct {
	LocationContext
	Path string `json:"path"`
}
