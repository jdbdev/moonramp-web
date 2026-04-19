package main

import (
	"fmt"
	"net/http"
)

func main() {

	// Handler routes
	http.HandleFunc("/", homeHandler)
	// Server
	http.ListenAndServe(":8080", nil)

}

// Hanlder functions
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "temporary home page for moonramp\n")
	fmt.Fprintf(w, "request path: %s", r.URL.Path)
}
