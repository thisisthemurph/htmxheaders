package htmxheaders

import (
	"fmt"
	"net/http"
)

type DecoratorFunction func(w http.ResponseWriter) error

// SetResponseHeaders sets custom HTTP headers in the provided http.ResponseWriter.
//
// This function applies one or more custom header decorators to the given response writer `w`.
// A decorator function modifies the response by adding or modifying HTTP headers as per the provided configuration.
//
// Parameters:
//
//	w: http.ResponseWriter - The response writer to which custom headers will be applied.
//	decorators: []DecoratorFunction - A slice of DecoratorFunction types representing header decorators.
//	            Each decorator function is responsible for setting specific HTTP headers.
//
// Returns:
//
//	error: An error if any of the decorator functions encounter an issue while setting the headers.
//	       It returns nil if all the decorators are applied successfully.
//
// Example usage:
//
//	err := SetResponseHeaders(w, PushURL("/new-url"), AddCustomHeader("Custom-Header", "Value"))
//	if err != nil {
//	    // Handle error
//	}
//
// Note:
//
//	If an error is returned, the function will not add any of the remaining headers, but will leave all
//	previously set headers, it is your responsibility to remove these headers.
//
//	The order of decorators matters. Headers set by decorators earlier in the slice may be overwritten
//	by subsequent decorators.
//
//	It's the responsibility of the caller to ensure that the provided response writer 'w' is not nil.
//	Passing a nil response writer will result in a panic.
func SetResponseHeaders(w http.ResponseWriter, decorators ...DecoratorFunction) error {
	for _, decorator := range decorators {
		if err := decorator(w); err != nil {
			return err
		}
	}
	return nil
}

func AddCustomHeader(key, value string) DecoratorFunction {
	return func(w http.ResponseWriter) error {
		w.Header().Set(key, value)
		return nil
	}
}

// RemoveHXHeaders removes all HTMX-related headers from the provided http.ResponseWriter.
//
// HTMX headers are special HTTP headers used to control HTMX behavior in web applications,
// such as client-side updates and navigation. This function clears headers commonly used
// by HTMX, including "HX-Location", "HX-Push-Url", "HX-Redirect", "HX-Refresh", "HX-Replace-Url",
// "HX-Reswap", "HX-Retarget", and "HX-Reselect".
//
// Parameters:
//
//	w: http.ResponseWriter - The response writer from which HTMX headers will be removed.
//
// Returns:
//
//	error: An error if the provided response writer 'w' is nil.
//
// Example usage:
//
//	w := httptest.NewRecorder()
//	w.Header().Set("HX-Location", "/some/location")
//	w.Header().Set("HX-Push-Url", "/some/url")
//	err := hh.RemoveHXHeaders(w)
//	if err != nil {
//	    // Handle the error
//	}
//
//	// Now, the HTMX headers are removed from the response recorder 'w'.
//
// Note:
//
//	This function does not clear non-HTMX headers.
func RemoveHXHeaders(w http.ResponseWriter) error {
	if w == nil {
		return fmt.Errorf("cannot clear HX headers for nil http.ResponseWriter")
	}

	headers := []string{
		"HX-Location",
		"HX-Push-Url",
		"HX-Redirect",
		"HX-Refresh",
		"HX-Replace-Url",
		"HX-Reswap",
		"HX-Retarget",
		"HX-Reselect",
	}

	// Loop through each header and remove it from the response
	for _, header := range headers {
		w.Header().Del(header)
	}

	return nil
}
