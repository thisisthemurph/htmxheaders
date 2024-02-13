package htmxheaders

// Reselect a CSS selector that allows you to choose which part of the response is used to be swapped in.
// Overrides an existing hx-select on the triggering element
// https://htmx.org/reference/#response_headers
func Reselect(selector string) DecoratorFunction {
	return AddCustomHeader("HX-Reselect", selector)
}
