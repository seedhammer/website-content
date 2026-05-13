package main

import (
	"net/http"

	"seedhammer.com/website/content"
)

func main() {
	mux := new(http.ServeMux)
	content.Register(mux)
	http.ListenAndServe(":8080", mux)
}
