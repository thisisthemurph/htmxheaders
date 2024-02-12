package htmxheaders

import (
	"encoding/json"
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

// PushURL pushes a new url into the history stack.
// https://htmx.org/headers/hx-push-url/
func PushURL(url string) DecoratorFunction {
	return AddCustomHeader("HX-Push-Url", url)
}

// PreventPushURL prevents the browser’s history from being updated.
// https://htmx.org/headers/hx-push-url/
func PreventPushURL() DecoratorFunction {
	return PushURL("false")
}

// Redirect can be used to do a client-side redirect to a new location.
// https://htmx.org/reference/#response_headers
func Redirect(path string) DecoratorFunction {
	return AddCustomHeader("HX-Redirect", path)
}

// Refresh forces the client-side to do a full refresh of the page.
// https://htmx.org/reference/#response_headers
func Refresh() DecoratorFunction {
	return AddCustomHeader("HX-Refresh", "true")
}

// PreventRefresh forces the client-side to do a full refresh of the page.
// https://htmx.org/reference/#response_headers
func PreventRefresh() DecoratorFunction {
	return AddCustomHeader("HX-Refresh", "false")
}

// ReplaceURL replaces the current URL in the location bar.
// https://htmx.org/headers/hx-replace-url/
func ReplaceURL(url string) DecoratorFunction {
	return AddCustomHeader("HX-Replace-Url", url)
}

// PreventReplaceURL replaces the current URL in the location bar.
// https://htmx.org/headers/hx-replace-url/
func PreventReplaceURL() DecoratorFunction {
	return ReplaceURL("false")
}

// Reswap allows you to override how the response will be swapped.
// https://htmx.org/reference/#response_headers
func Reswap(swapMethod Swap) DecoratorFunction {
	return AddCustomHeader("HX-Reswap", swapMethod.String())
}

// Retarget a CSS selector that overrides the target of the content update to
// a different element on the page.
// https://htmx.org/reference/#response_headers
func Retarget(target string) DecoratorFunction {
	return AddCustomHeader("HX-Retarget", target)
}

// Reselect a CSS selector that allows you to choose which part of the response is used to be swapped in.
// Overrides an existing hx-select on the triggering element
// https://htmx.org/reference/#response_headers
func Reselect(selector string) DecoratorFunction {
	return AddCustomHeader("HX-Reselect", selector)
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
