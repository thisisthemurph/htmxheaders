package main

import (
	"fmt"
	"html/template"
	"net/http"

	hh "github.com/thisisthemurph/htmxheaders"
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

	http.HandleFunc("/send", func(w http.ResponseWriter, r *http.Request) {
		subject := r.FormValue("subject")
		message := r.FormValue("message")

		if subject == "" || message == "" {
			// Set HX-Retarget and HX-Reswap headers to show error message.
			_ = hh.SetResponseHeaders(w, hh.Retarget("#error"), hh.Reswap(hh.SwapOuterHTML))

			// Render the partial, which will be injected into the correct place with HTMX.
			errMsg := struct{ ErrorMessage string }{ErrorMessage: "You must provide a subject and a message."}
			renderPartial(w, "error.html", errMsg)
			return
		}

		// Handle form submission (send email, save to database, etc.)
		// For demonstration, just print the received values.
		fmt.Printf("Subject: %s\nMessage: %s\n", subject, message)
		_ = hh.SetResponseHeaders(w, hh.Redirect("/thankyou"))
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
