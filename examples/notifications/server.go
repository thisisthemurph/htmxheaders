package main

import (
	"fmt"
	hh "github.com/thisisthemurph/htmxheaders"
	"html/template"
	"net/http"
	"strconv"
)

// In this example, we have a form that will redirect you to the thank-you page on success using
// the Redirect function, which sets the HX-Redirect header.
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
		data := struct{ Value string }{strconv.Itoa(val + 1)}

		if (val+1)%3 == 0 {
			fmt.Printf("%v is divisible by 3.\n", val+1)
			hh.SetResponseHeaders(hh.)
		}

		renderPartial(w, "input.html", data)
		return
	})

	http.HandleFunc("/thankyou", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "thankyou.html")
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
