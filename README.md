# HTMX Headers

A Go package for manipulating HTMX headers in HTTP responses.

## Overview

HTMX Headers Go Package provides functionality to add, modify, and remove HTMX headers in HTTP responses.
HTMX headers are special HTTP headers used to control HTMX behavior in web applications, such as client-side updates 
and navigation.

For more information on HTMX headers, read the official HTMX Response Headers Reference at [htmx.org](https://htmx.org/reference/#response_headers).

## Installation

To install HTMX Headers Go Package, use `go get`:

```bash
go get github.com/thisisthemurph/htmxheaders
```

## Usage

Import the htmxheaders package in your Go code:

```go
package main

import "github.com/thisisthemurph/htmxheaders"
```

or use an alias:

```go
package main

import hh "github.com/thisisthemurph/htmxheaders"
```

### Adding headers

Adding headers uses the `SetResponseHeaders` function, which takes a `http.ResponseWriter` followed by any number of
decorator functions which will add the relevant header with the appropriate value.

The following example tells HTMX to change the original swap and target values set in the `hx-swap` and `hx-target` attributes:

```go
w := // Your http.ResponseWriter
err := hh.SetResponseHeaders(w, hh.Reswap(hh.SwapOuterHTML), hh.Retarget("#my-component"))
if err != nil {
    // Handle error
}
```

### Custom headers

If you have a custom header you would like to add, this is possible with the `AddCustomHeader` function:

```go
hh.SetResponseHeaders(w, hh.AddCustomHeader("key", "value"))
```

### Removing HX headers

If you want to remove any headers before returning your response, you can use the `RemoveHXHeaders` method:
This will remove all HX headers, but will not affect any of your own headers.

```go
w := // Your http.ResponseWriter
err := hh.RemoveHXHeaders(w)
if err != nil {
    // Handle error
}
```