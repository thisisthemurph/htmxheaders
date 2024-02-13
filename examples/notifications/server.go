package main

import (
	"fmt"
	hh "github.com/thisisthemurph/htmxheaders"
	"html/template"
	"net/http"
	"strconv"
)

// In this example we have a form that will send the current numeric value.
// The handler will then increment the value and replace the input element with
// a partial containing the new value.
// If the value is divisible by 3, the `HX-Trigger` header will be set with a message
// to be presented to the user using the JS alert method; see the `index.html`.
//
// If there is an issue with the form data, an error message partial will instead be returned
// and presented above the form using the `Reswap` and `Retarget` functions to set the `Re-Swap`
// and `Re-Target` headers.
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/inc", func(w http.ResponseWriter, r *http.Request) {
		val, _ := strconv.Atoi(r.FormValue("value"))

		if (val+1)%3 == 0 {
			// Create a message and add it to the detail of the `showAlert` event.
			// The details can be anything that can be serialized into JSON.
			// Here we are just using a simple string, but it could be anything.
			message := fmt.Sprintf("Number %v is divisible by 3!", val+1)
			event := hh.TriggerEvent{Name: "showAlert", Detail: message}
			_ = hh.SetResponseHeaders(w, hh.TriggerWithDetail(hh.TriggerAfterSwap, event))
		}

		// We continue with our normal rendering in this example, but you may wish to do something
		// else depending on your use case.
		partialData := struct{ Value string }{strconv.Itoa(val + 1)}
		renderPartial(w, "input.html", partialData)
		return
	})

	http.ListenAndServe(":3000", nil)
}

// renderPartial is a helper method to render a partial HTML response.
// This is just for demonstration, there will be better ways to do this.
func renderPartial(w http.ResponseWriter, name string, data interface{}) {
	tmpl, err := template.ParseFiles(name)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing template %q: %v", name, err), http.StatusInternalServerError)
		return
	}

	if err = tmpl.Execute(w, data); err != nil {
		http.Error(w, fmt.Sprintf("Error executing template %q: %v", name, err), http.StatusInternalServerError)
	}
}
