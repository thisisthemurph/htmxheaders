# HTMX Headers Documentation

A Go package for manipulating HTMX headers in HTTP responses.

HTMX Headers Go Package provides functionality to add, modify, and remove HTMX headers in HTTP responses.
HTMX headers are special HTTP headers used to control HTMX behavior in web applications, such as client-side updates
and navigation.

For more information on HTMX headers, read the official HTMX Response Headers Reference at [htmx.org](https://htmx.org/reference/#response_headers).

## SetResponseHeaders

This is the primary function that applies custom HTTP headers to the provided response writer (w) using one or more 
header decorators. Each decorator function modifies the response by adding or modifying HTTP headers as per the 
provided configuration.

**Parameters:**

- `w`: `http.ResponseWriter` - The response writer to which custom headers will be applied.
- `decorators`: `[]DecoratorFunction` - A slice of DecoratorFunction types representing header decorators. 
Each decorator function is responsible for setting specific HTTP headers.

**Returns:**

- `error`: An error if any of the decorator functions encounter an issue while setting the headers. 
Returns nil if all decorators are applied successfully.

**Example usage:**

```go
err := hh.SetResponseHeaders(w, hh.Reswap(hh.SwapOuterHTML), hh.Retarget("#new-target"))
if err != nil {
    // Handle error
}
```

## AddCustomHeader

This function returns a decorator function that adds a custom HTTP header to the response writer.
You should not generally need to use this function, but it serves as a useful way to add HTMX headers that may be
more recent than this package or that may be bespoke to your solution.

**Parameters:**

- `key`: `string` - The name of the custom HTTP header.
- `value`: `string` - The value to set for the custom HTTP header.

**Returns:**

- `DecoratorFunction`: A decorator function that sets the provided custom HTTP header in the response writer.
  - `error`: Always `nil`. The decorator function returned by `AddCustomHeader` does not return any errors.

**Example usage:**

```go
// Add a custom header "Authorization: Bearer token123"
_ = hh.SetResponseHeaders(w, hh.AddCustomHeader("Authorization", "Bearer token123"))
```

## Location

This function returns a decorator function that sets the `HX-Location` header, allowing client-side redirect without a 
full page reload.

**Parameters:**

- `location`: `string` - The URL to which the client will be redirected.

**Returns:**

- `DecoratorFunction`: A decorator function that sets the `HX-Location` header with the provided URL in the response writer.
  - `error`: Always `nil`. The decorator function returned by `Location` does not return any errors.

**Example usage:**

```go
// Apply the decorator to set the client-side redirect location using SetResponseHeaders
_ = hh.SetResponseHeaders(w, hh.Location("/new-url"))
```

## LocationWithContext

This function returns a decorator function that sets the `HX-Location` header with additional options provided in the context.

**Parameters:**

- `path`: `string` - The URL path to which the client will be redirected.
- `context`: `LocationContext` - Additional options provided in the context for the client-side redirect.

**Returns:**

- `DecoratorFunction`: A decorator function that sets the `HX-Location` header with the provided URL and context in the response writer.
    - `error`: May return an error if there is an issue marshalling the context JSON. However, in most cases, this error is always `nil`.

**Example usage:**

```go
// Define a LocationContext
ctx := hh.LocationContext{
    Target: "#my-new-target",
	Swap:   hh.SwapOuterHTML
}

// Apply the decorator to set the client-side redirect location with context using SetResponseHeaders
_ = hh.SetResponseHeaders(w, hh.LocationWithContext("/new-url", ctx))
```

## PushURL

This function returns a decorator function that sets the `HX-Push-Url` header, pushing a new URL into the history stack.

**Parameters:**

- `url`: `string` - The URL to push into the history stack.

**Returns:**

- `DecoratorFunction`: A decorator function that sets the `HX-Push-Url` header with the provided URL in the response writer.
    - `error`: Always `nil`. The decorator function returned by `PushURL` does not return any errors.

**Example usage:**

```go
// Apply the decorator to push a new URL into the history stack using SetResponseHeaders
_ = hh.SetResponseHeaders(w, hh.PushURL("/new-url"))
```

## PreventPushURL

This function returns a decorator function sets the `HX-Push-URL` header with a value of `false` preventing updating 
the browser's history.

**Parameters:**

None

**Returns:**

- `DecoratorFunction`: A decorator function that sets the `HX-Push-Url` header to `false` in the response writer, 
preventing the browser's history from being updated.
    - `error`: Always `nil`. The decorator function returned by `PreventPushURL` does not return any errors.

**Example usage:**

```go
// Apply the decorator to prevent updating the browser's history using SetResponseHeaders
_ = hh.SetResponseHeaders(w, hh.PreventPushURL())
```

## Redirect

This function returns a decorator function that sets the `HX-Redirect` header for a client-side redirect to a new location.

**Parameters:**

- `path`: `string` - The URL path to which the client will be redirected.

**Returns:**

- `DecoratorFunction`: A decorator function that sets the `HX-Redirect` header with the provided URL path in the 
response writer.
    - `error`: Always `nil`. The decorator function returned by `Redirect` does not return any errors.

**Example usage:**

```go
// Apply the decorator to redirect the client to a new location using SetResponseHeaders
_ = hh.SetResponseHeaders(w, hh.Redirect("/new-url"))
```

## Refresh

This function returns a decorator function that sets the `HX-Refresh` header with the value of `true` to forces a 
full refresh of the page on the client-side.

**Parameters:**

None

**Returns:**

- `DecoratorFunction`: A decorator function that sets the `HX-Refresh` header to `true` in the response writer, 
forcing a full refresh of the page.
  - `error`: Always `nil`. The decorator function returned by `Refresh` does not return any errors.

**Example usage:**

```go
// Apply the decorator to force a full refresh of the page using SetResponseHeaders
_ = hh.SetResponseHeaders(w, hh.Refresh())
```

## PreventRefresh

This function returns a decorator function that sets the `HX-Refresh` header to `false` preventing a full refresh 
of the page on the client-side.

**Parameters:**

None

**Returns:**

- `DecoratorFunction`: A decorator function that sets the `HX-Refresh` header to `false` in the response writer, 
preventing a full refresh of the page.
  - `error`: Always `nil`. The decorator function returned by `PreventRefresh` does not return any errors.

**Example usage:**

```go
// Apply the decorator to prevent a full refresh of the page using SetResponseHeaders
_ = hh.SetResponseHeaders(w, hh.PreventRefresh())
```

## ReplaceURL

This function returns a decorator function that sets the `HX-Replace-Url` header, replacing the current URL in 
the location bar.

**Parameters:**

- `url`: `string` - The URL to replace the current URL in the location bar.

**Returns:**

- `DecoratorFunction`: A decorator function that sets the `HX-Replace-Url` header with the provided URL in the response writer.
  - `error`: Always `nil`. The decorator function returned by `ReplaceURL` does not return any errors.

**Example usage:**

```go
// Apply the decorator to replace the current URL in the location bar using SetResponseHeaders
_ = hh.SetResponseHeaders(w, hh.ReplaceURL("/new-url"))
```

## PreventReplaceURL

This function returns a decorator function that prevents replacing the current URL in the location bar.

**Parameters:**

None

**Returns:**

- `DecoratorFunction`: A decorator function that sets the `HX-Replace-Url` header to `"false"` in the response writer, preventing the replacement of the current URL in the location bar.

    - `error`: Always `nil`. The decorator function returned by `PreventReplaceURL` does not return any errors.

**Example usage:**

```go
// Apply the decorator to prevent replacing the current URL in the location bar using SetResponseHeaders
_ = hh.SetResponseHeaders(w, hh.PreventReplaceURL())
```

## PreventReplaceURL

This function returns a decorator function that sets the `HX-Replace-Url` header to `false`, preventing the 
replacement of the current URL in the location bar.

**Parameters:**

None

**Returns:**

- `DecoratorFunction`: A decorator function that sets the `HX-Replace-Url` header to `false` in the response writer, 
preventing the replacement of the current URL in the location bar.

    - `error`: Always `nil`. The decorator function returned by `PreventReplaceURL` does not return any errors.

**Example usage:**

```go
// Apply the decorator to prevent replacing the current URL in the location bar using SetResponseHeaders
_ = hh.SetResponseHeaders(w, hh.PreventReplaceURL())
```

## Reswap

This function returns a decorator function that sets the `HX-Reswap` header, allowing you to override how the 
response will be swapped.

**Parameters:**

- `swapMethod`: `Swap` - The method for swapping the response.

Possible options for the `swapMethod` are:

- `SwapInnerHTML`
- `SwapOuterHTML`
- `SwapBeforeBegin`
- `SwapAfterBegin`
- `SwapBeforeEnd`
- `SwapAfterEnd`
- `SwapDelete`
- `SwapNone`

**Returns:**

- `DecoratorFunction`: A decorator function that sets the `HX-Reswap` header with the provided swap method in the response writer.

    - `error`: Always `nil`. The decorator function returned by `Reswap` does not return any errors.

**Example usage:**

```go
// Apply the decorator to override how the response will be swapped using SetResponseHeaders
_ = hh.SetResponseHeaders(w, hh.Reswap(hh.AppendSwap))
```

## Retarget

This function returns a decorator function that sets the `HX-Retarget` header, allowing you to override the target 
of the content update to a different element on the page.

**Parameters:**

- `target`: `string` - The CSS selector representing the element to which the content update will be targeted.

**Returns:**

- `DecoratorFunction`: A decorator function that sets the `HX-Retarget` header with the provided target selector in the response writer.
    - `error`: Always `nil`. The decorator function returned by `Retarget` does not return any errors.

**Example usage:**

```go
// Apply the decorator to override the target element for content update using SetResponseHeaders
_ = hh.SetResponseHeaders(w, hh.Retarget("#update-div"))
```

## Reselect

This function returns a decorator function that sets the `HX-Reselect` header, allowing you to choose which part 
of the response is used to be swapped in, overriding an existing `hx-select` on the triggering element.

**Parameters:**

- `selector`: `string` - The CSS selector representing the part of the response to be swapped in.

**Returns:**

- `DecoratorFunction`: A decorator function that sets the `HX-Reselect` header with the provided selector in the 
response writer.
  - `error`: Always `nil`. The decorator function returned by `Reselect` does not return any errors.

**Example usage:**

```go
// Apply the decorator to choose which part of the response to swap in using SetResponseHeaders
_ = hh.SetResponseHeaders(w, hh.Reselect("#content"))
```

## Trigger

This function returns a decorator function that sets the `HX-Trigger` header, allowing you to trigger one or more 
events on the client side. For more information, including how to listen for the event in JavaScript, view the
official [HTMX documentation](https://htmx.org/headers/hx-trigger).

**Parameters:**

- `eventName`: `string` - The name(s) of the event(s) to trigger.

**Returns:**

- `DecoratorFunction`: A decorator function that sets the `HX-Trigger` header with the provided event name(s) in 
the response writer.
  - `error`: Always `nil`. The decorator function returned by `Trigger` does not return any errors.

**Example usage:**

```go
// Apply the decorator to trigger an event using SetResponseHeaders
_ = hh.SetResponseHeaders(w, hh.Trigger("customEvent"))
```

Or you can do multiple events like:

```go
_ = hh.SetResponseHeaders(w, hh.Trigger("customEvent", "myOtherEvent", "yetAnotherEvent"))
```

If your events are in a slice:

```go
events := []string{"customEvent1", "customEvent2"}
_ = hh.SetResponseHeaders(w, hh.Trigger(events...))
```

## TriggerWithDetail

This function returns a decorator function that adds an event JSON object to the `HX-Trigger` header within the 
response headers, containing a mapping of event names to their corresponding details.

**Parameters:**

- `events`: `...TriggerEvent` - Variable number of TriggerEvent structs, each specifying an event name and its 
associated details.

**Returns:**

- `DecoratorFunction`: A decorator function that sets the `HX-Trigger` header with the JSON representation of the 
provided event details in the response writer.
  - `error`: May return an error if there is an issue marshalling the event details into JSON. However, in most cases, this error is always `nil`.

**Example usage:**

```go
// Define TriggerEvent structs
events := []hh.TriggerEvent{
    {Name: "customEvent1", Detail: "detail1"},
    {Name: "customEvent2", Detail: "detail2"},
}

// Apply the decorator to add event details to the response headers using SetResponseHeaders
_ = hh.SetResponseHeaders(w, hh.TriggerWithDetail(events...))
```

The detail can be any object that can be marshalled into JSON. For example, you may have an event that provides toast
notifications to your users.

```go
type Toast struct {
	LogLevel int    `json:"level"`
	Title    string `json:"title"`
	Message  string `json:"message"`
}

events := []hh.TriggerEvent{
	{Name: "toast", Detail: Toast{LogLevel: 2, Title: "Warning", Message: "Incorrect email format."}},
	{Name: "toast", Detail: Toast{LogLevel: 1, Title: "Info", Message: "Account not created."}},
}

_ = hh.SetResponseHeaders(w, hh.TriggerWithDetail(events...))
```

## RemoveHXHeaders

This function returns a decorator function that removes all HTMX-related headers from the response writer, 
ensuring that they are not included in the HTTP response.

**Parameters:**

- `w`: `http.ResponseWriter` - The response writer from which HTMX headers will be removed.

**Returns:**

- `error`: An error if the provided response writer `w` is nil.

**Example usage:**

```go
// Assuming some headers have been set as detailed above.

// Remove HTMX headers
err := hh.RemoveHXHeaders(w)
if err != nil {
    // Handle the error
}
```