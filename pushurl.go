package htmxheaders

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
