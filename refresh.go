package htmxheaders

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
