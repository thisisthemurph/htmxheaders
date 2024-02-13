package htmxheaders

// Retarget a CSS selector that overrides the target of the content update to
// a different element on the page.
// https://htmx.org/reference/#response_headers
func Retarget(target string) DecoratorFunction {
	return AddCustomHeader("HX-Retarget", target)
}
