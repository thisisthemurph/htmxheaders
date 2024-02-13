package htmxheaders

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
