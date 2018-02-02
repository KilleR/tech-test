package main

import (
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "markup.html")
}
