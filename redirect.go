package htmxheaders

// Redirect can be used to do a client-side redirect to a new location.
// https://htmx.org/reference/#response_headers
func Redirect(path string) DecoratorFunction {
	return AddCustomHeader("HX-Redirect", path)
}
