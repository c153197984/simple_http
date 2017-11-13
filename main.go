package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	// A simple http server.
	//
	// This server send "Hello user" to client based on URL.
	// Example:
	//
	// localhost:8080/foo/bar
	// > Hello foo bar!
	//
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello %s!\n",
			strings.Join(strings.Split(r.URL.String(), "/")[1:], " "))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
